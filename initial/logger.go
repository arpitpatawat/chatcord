package initial

import "go.uber.org/zap"

// Using Zap logging 
func InitLogger() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
}
