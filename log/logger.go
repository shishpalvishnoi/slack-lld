package log

import "go.uber.org/zap"

func GetNewLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		logger.Error("Error in logger: ", zap.Error(err))
	}
	return logger
}
