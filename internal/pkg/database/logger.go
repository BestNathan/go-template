package database

import (
	"os"
	"strings"

	gl "gorm.io/gorm/logger"
)

type Logger struct{}

func NewLogger() *Logger {
	return &Logger{}
}

func (l Logger) GormLogger() gl.Interface {
	env := os.Getenv("GO_ENV")
	if strings.ToLower(env) == "production" {
		return gl.Default
	}

	return gl.Default.LogMode(gl.Info)
}
