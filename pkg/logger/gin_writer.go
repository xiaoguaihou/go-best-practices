package logger

import (
	"io"
)

type GinWriter struct {
	io.Writer
}

func (s GinWriter) Write(p []byte) (n int, err error) {

	logger.Info(string(p))

	return len(p), nil
}
