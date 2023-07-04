package log

import (
	"encoding/json"
	"os"
	"strconv"

	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	isLogTypeDevelopment, err := strconv.ParseBool(os.Getenv("IS_LOG_TYPE_DEV"))
	if err != nil {
		isLogTypeDevelopment = false
	}

	callerSkip := zap.AddCallerSkip(1)

	if isLogTypeDevelopment {
		logger, _ = zap.NewDevelopment(callerSkip)
	} else {
		logger, _ = zap.NewProduction(callerSkip)
	}
}

func String(key string, value string) zap.Field {
	return zap.String(key, value)
}

func Int(key string, value int) zap.Field {
	return zap.Int(key, value)
}

func Bool(key string, value bool) zap.Field {
	return zap.Bool(key, value)
}

func Object(objectKey string, object interface{}) zap.Field {
	jsonByte, _ := json.Marshal(object)
	return zap.String(objectKey, string(jsonByte))
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}
