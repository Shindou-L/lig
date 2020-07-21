package writer

import (
	"io"
	"log"
)

type WriteFunc func(string)

// calls Output to print to the standard logger
func Standard(msg string) {
	log.Println(msg)
}

func Blank(_ string) {}

func NewWriter(writer io.Writer) WriteFunc {
	return func(s string) {
		_, _ = writer.Write([]byte(s))
	}
}

func WithLogger(logger *log.Logger) WriteFunc {
	return func(s string) {
		logger.Println(s)
	}
}
