package logger

import (
	"go.uber.org/zap"
)

var Log *zap.Logger

func InitLogger() {
	var err error
	// Production preset is good for JSON logs
	Log, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}