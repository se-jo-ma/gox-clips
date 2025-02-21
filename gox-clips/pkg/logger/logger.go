package logger

import (
	"log"
)

func Logger() *log.Logger {
	return log.Default()
}

func Logger_test() *log.Logger {
	return log.Default()
}
