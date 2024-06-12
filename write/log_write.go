package write

import (
	"github.com/zack-sys/log/queue"
	"os"
)

type Log struct {
	Console bool
}

func NewLog() *Log {
	return &Log{Console: true}
}

func NewLogCfg(l *Log) *Log {
	return l
}

func (l Log) Write(p []byte) (n int, err error) {
	go queue.Push(string(p))
	if !l.Console {
		return 0, nil
	}
	return os.Stdout.Write(p)
}
