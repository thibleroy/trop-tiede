package utils

import (
	"math/rand"

	"go.uber.org/zap"
)

func RandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
func RandomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(RandInt(65, 90))
	}
	return string(bytes)
}

func HandleError(err error, msg string) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	if err != nil {
		logger.Error("Error",
			zap.Error(err),
			zap.String("Error message", msg),
		)
	}
}
