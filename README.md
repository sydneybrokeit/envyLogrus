# ENVy Logrus
Small helper library to turn an arbitrary environment variable into a Logrus log level.

## how to use
```
import github.com/hmschreck/envyLogrus
...

logrus.SetLevel(envyLogrus.GetLogLevel("env", "fallback"))
```

If no environment variable is found, it will use the fallback; if the fallback doesn't
exist in the map, it will default to `logrus.DebugLevel`

## using custom log maps
custom log level maps can be used.

```
var customLogLevels = map[string]logrus.Level{
    "X": logrus.DebugLevel,
    "Y": logrus.ErrorLevel,
}

...

logrus.SetLevel(envyLogrus.GetCustomLogLevel("env_key", "fallback", customLogLevels))
```