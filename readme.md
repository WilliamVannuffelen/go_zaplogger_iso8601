```go
package main

import (
	logger "zaplogger_iso8601"
	zap "go.uber.org/zap"
)


func initLogger(filePath string, logLevel string) (*zap.Logger, error) {
	log, err := logger.InitLogger(filePath, logLevel)
	if err != nil {
		log.Fatal("fatality")
	}
	return log, err
}


func main() {
	log, err := initLogger("log.txt", "debug")
  if err != nil {
		log.Fatal("mimimi")
	}

	log.Info("Success!")
```