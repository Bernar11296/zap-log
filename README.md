# Zap-log
## Information 

### Basic log
```
package main

import (
	"log"
	"os"
)

func main() {
	logFile, _ := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer logFile.Close()
	log.SetOutput(logFile)
	log.Println("Hello, world!")
}

```

### Zap log
```
$ go get -u go.uber.org/zap
```
package utils

import "go.uber.org/zap"

var Logger *zap.Logger

func InitializeLogger() {
	Logger, _ = zap.NewProduction()
}
```
utils.InitializeLogger()
	utils.Logger.Info("Hello, world!")
	utils.Logger.Error("Not able to reach website", zap.String("url", "google.com"))

```
Result JSON:

{
    "level": "error",
    "ts": 1648382422.5089228,
    "caller": "utils/logger.go:22",
    "msg": "Not able to reach blog.",
    "url": "codewithmukesh.com",
    "attempt": 3,
    "stacktrace": "github.com/iammukeshm/structured-logging-golang-zap/utils.Error\n\tD:/go-repo/structured-logging-golang-zap/utils/logger.go:22\nmain.main\n\tD:/go-repo/structured-logging-golang-zap/main.go:10\nruntime.main\n\tC:/Program Files/Go/src/runtime/proc.go:250"
}

### 
```
func Initialize() {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(config)
	logFile, _ := os.OpenFile("log.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	writer := zapcore.AddSync(logFile)
	defaultLogLevel := zapcore.DebugLevel
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
	)
	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}
```
Result : log.json file
```

