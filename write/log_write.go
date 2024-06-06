package write

import (
	"os"
	"testlog/queue"
)

type Log struct {
}

func NewLog() *Log {
	return &Log{}
}

func (l Log) Write(p []byte) (n int, err error) {
	go queue.Push(p)
	return os.Stdout.Write(p)
}
