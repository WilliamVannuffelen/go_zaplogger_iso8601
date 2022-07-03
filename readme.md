# go_zaplogger_iso8601

A simple wrapper package around [go.uber.org/zap](go.uber.org/zap).

Limited zapLogger API exposure through an interface. Intended for targeted use with CLI applications.

- Timestamps in ISO8601 format. 
- Writes unstructured human-readable output to stdout/stderr and a file of choice.
- Only exposes Debug(), Info(), Warn(), Error(), Panic() and Fatal() methods.

# Usage

## Preconfigured zap.SugaredLogger gets exposed via an interface:

```go
type Logger interface{ ... }
```


## Init logger by calling:

```go
func InitLogger(filePath string, logLevel string) (Logger, error)
// valid log levels are 'debug', 'info', 'warn', 'error'
```

Example:

```go
package main

import (
  logger "github.com/williamvannuffelen/go_zaplogger_iso8601"
)

func main() {
  log, invalidLevel := logger.InitLogger("log.txt", "debug")

  log.Info("Foo!")
  log.Warn("Bar!")
}
```
Results in:
```
2022-06-28T22:49:42.426+0200 - INFO - go_zaplogger_iso8601@v0.1.2/zaplogger_iso8601.go:56 - github.com/williamvannuffelen/go_zaplogger_iso8601.InitLogger - Logger init successful.
2022-06-28T22:49:42.446+0200 - INFO - test/test.go:10 - main.main - Foo!
2022-06-28T22:49:42.446+0200 - WARN - test/test.go:11 - main.main - Bar!
```

Logger can be passed to other packages if they import zap.

```go
package main

import (
  logger "github.com/williamvannuffelen/go_zaplogger_iso8601"
)

func main() {
  log, err := logger.InitLogger("log.txt", "debug")
	if err != nil {
		log.Warn(err)
	}
	log.Info("Logger initialized successfully.")  

  // pass the interface as an explicit dependency
  childpackage.ShowExample(log, "example")
}

//##################

package childpackage

import (
  "fmt"
  logger "github.com/williamvannuffelen/go_zaplogger_iso8601"
)

func ShowExample(log logger.Logger, example string) {
  log.Info(fmt.Sprintf("This is an %s.", example))
}
```
Results in 
```
2022-06-28T22:29:56.183+0200 - INFO - main/main.go:10 - main.main - Logger initialized successfully.
2022-06-28T22:29:56.202+0200 - INFO - childpackage/childpackage.go:10 - childpackage/childpackage.ShowExample - This is an example.

```

# Error handling

Logger panics when there are issues with the provided filePath.
```
panic: couldn't open sink "/log.txt": open /log.txt: read-only file system

goroutine 1 [running]:
zaplogger_iso8601.InitLogger({0x104a0270f, 0x8}, {0x104a0248a, 0x7})
        /Users/me/go/pkg/mod/github.com/williamvannuffelen/go_zaplogger_iso8601@v0.1.2/zaplogger_iso8601.go:52 +0x388
main.main()
        /Users/me/devel/go/example/example.go:11 +0x38
exit status 2
```
Logger returns an error object when provided with an invalid logLevel and defaults to 'info'.
```
2022-06-28T22:17:47.050+0200 - INFO - main/main.go:10 - main.main - invalid value provided for logLevel. Defaulting to 'info'
```