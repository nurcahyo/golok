package sentry

import (
	"log"
	"os"
	"runtime"
	"strconv"

	raven "github.com/getsentry/raven-go"
	"github.com/nurcahyo/golok/contract"
	"github.com/nurcahyo/golok/util"
	pkgErrors "github.com/pkg/errors"
)

var (
	mutateTagMiddlewares []contract.MutateSentryTagsMiddleware
)

type SentryHandler struct {
	DSN   string
	level string
	env   string
	wait  bool
}

func NewHandler(cfg map[string]interface{}) contract.Loggable {
	if _, ok := cfg["dsn"]; !ok {
		log.Fatal("Can't find sentry dsn")
	}
	env := util.MapGet(cfg, "environment", os.Getenv("APP_ENV")).(string)
	if env == "" {
		env = os.Getenv("ENVIRONMENT")
	}
	raven.SetDefaultLoggerName("github.com/nurcahyo/golok")
	return &SentryHandler{
		DSN:   cfg["dsn"].(string),
		level: util.MapGet(cfg, "level", "error").(string),
		env:   env,
		wait:  util.MapGet(cfg, "wait", false).(bool),
	}
}

type causer interface {
	Cause() error
}

// Iteratively fetches all the Extra data added to an error,
// and it's underlying errors. Extra data defined first is
// respected, and is not overridden when extracting.
func extractExtra(err error) raven.Extra {
	extra := raven.Extra{}

	currentErr := err
	for currentErr != nil {
		if errWithExtra, ok := currentErr.(raven.ErrWithExtra); ok {
			for k, v := range errWithExtra.ExtraInfo() {
				extra[k] = v
			}
		}

		if errWithCause, ok := currentErr.(causer); ok {
			currentErr = errWithCause.Cause()
		} else {
			currentErr = nil
		}
	}

	return extra
}

func (handler *SentryHandler) write(level string, msg interface{}) {
	if util.LevelLte(handler.level, level) {
		raven.SetDSN(handler.DSN)
		_, file, line, ok := runtime.Caller(3)
		if !ok {
			file = "???"
			line = 0
		}

		tags := map[string]string{
			"file":        file,
			"line":        strconv.Itoa(line),
			"level":       level,
			"environment": handler.env,
			"os":          runtime.GOOS,
			"os.Arch":     runtime.GOARCH,
		}
		client := raven.DefaultClient

		if err, ok := msg.(error); ok {
			var interfaces []raven.Interface
			if client == nil {
				return
			}

			extra := extractExtra(err)
			cause := pkgErrors.Cause(err)

			packet := raven.NewPacketWithExtra(
				err.Error(),
				extra,
				append(interfaces,
					raven.NewException(cause, raven.GetOrNewStacktrace(cause, 3, 3, client.IncludePaths())))...)

			if handler.wait {
				eventID, ch := client.Capture(packet, tags)
				if eventID != "" {
					<-ch
				}
				return
			}
			client.Capture(packet, tags)
			return
		}
		if handler.wait {
			client.CaptureMessageAndWait(
				msg.(string),
				tags,
			)
			return
		}

		client.CaptureMessage(
			msg.(string),
			tags,
		)
	}

}

func (handler *SentryHandler) Debug(msg string) {
	handler.write("debug", msg)
}

func (handler *SentryHandler) Info(msg string) {
	handler.write("info", msg)
}

func (handler *SentryHandler) Warning(msg string) {
	handler.write("warning", msg)
}

func (handler *SentryHandler) Error(err error) {
	handler.write("error", err)
}

func (handler *SentryHandler) Critical(err error) {
	handler.write("critical", err)
}
