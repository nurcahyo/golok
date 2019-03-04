# Golok
Golok is a logger for golang, for easy logging in your programs.
Golok here to help you learn more about what's happening within your application, Golok provides robust logging services that allow you to log messages to files, the system error log, and [sentry](https://sentry.io)

Golok inspired by laravel logger facades.

## Install
```
$go get -u -v github.com/nurcahyo/golok
```

Use go get -u to update the package.

## Configuration

To config you just need a map. 
You can use json file and unmarshall it to map[string]interface{} or just use map[string]interface{} on your code.
Example config:
```
var config = map[string]interface{}{
	"default": "stack",
	"channels": map[string]interface{}{
		"stack": map[string]interface{}{
			"channels": []string{"system", "file"},
			"driver":   "stack",
		},
		"system": map[string]interface{}{
			"level":  "debug",
			"driver": "syslog",
		},
		"file": map[string]interface{}{
			"level":    "debug",
			"daily":    false,
			"filename": "log",
			"path":     "./",
			"driver":   "file",
		},
		"sentry": map[string]interface{}{
			"level":       "error",
			"driver":      "sentry",
			"environment": "golok-test",
			"wait":        true,
			"dsn":         "https://772b29d6912f4efba93f82716ecfbb6f:e6c87cac79e448859c21ffe53ccee741@sentry.io/1407424",
		},
	},
}
```


Available Channel Drivers

| Name  | Description |
|-------|-------------|
| stack | A wrapper to facilitate creating "multi-channel" channels |
| syslog| A handler to print error log to stderr or to your console output |
| file  | A handler to print error log to file |
| sentry| A handler to print error log to [sentry](https://sentry.io) |

### Configuring stack channel
As previously mentioned, the stack driver allows you to combine multiple channels into a single log channel. To illustrate how to use log stacks, let's take a look at an example configuration that you might see in a production application:

On the example config above you can see this code part.

```
...
"stack": map[string]interface{}{
  "channels": []string{"system", "file"},
  "driver":   "stack",
}
...
```
Let's dissect this configuration. First, notice our stack channel aggregates two other channels via its channels option: system and file. So, when logging messages, both of these channels will have the opportunity to log the message.

## Writing Log Message

You may write information to the logs using golok package. The logger provides the eight logging levels defined:
- Debug
- Info
- Warning
- Error
- Critical

To write the log message you need to import package first
```
import github.com/nurcahyo/golok
```
After import write the config for example use the example config on **Configuring Slack channel** sections above.
And then initialize the config

```
golok.Initialize(config)
```

After that you can call log level function directly from any source that import the golok package.
```
golok.Debug("test log")
golok.Error(errors.New("Iam test log"))
golok.Stack([]string{"sentry"}).Debug("Test stack with sentry log below config level")
golok.Stack([]string{"sentry"}).Error(errors.New("Test stack with sentry log pushed"))
```

For full example see test file [golok_test](https://github.com/nurcahyo/golok/blob/master/golok_test.go)


## Future Development Plan

- Decorate syslog output
- Add ability to write log to Specific Channels for example: `golok.Only("file").Debug("Write to file")`
- Add New Relic Driver
- Add Slack Notification Driver
- Add contextual information for example: `golok.Info("Create post failed", map[string]interface{"user_id": 1}))


