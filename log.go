package log

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/zack-sys/log/enum"
	"github.com/zack-sys/log/es"
	"github.com/zack-sys/log/util"
	"github.com/zack-sys/log/write"
	"go.uber.org/zap/zapcore"
	"os"
	"runtime"
)

var index string
var skipCall int

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetOutput(write.NewLog())
	log.SetLevel(log.InfoLevel)
	skipCall = 2
}
func SetIndex(i string) {
	index = i
	es.SetEsIndex(i)
}

func SkipCall(skip int) {
	skipCall = skip
}

func getBasicFields(ctx context.Context) log.Fields {
	if index == "" {
		panic("索引不能为空")
	}

	basicfields := log.Fields{
		"caller":       zapcore.NewEntryCaller(runtime.Caller(skipCall)).TrimmedPath(),
		"@fluentd_tag": index,
		"Trace_Id":     util.GetString(ctx.Value(enum.TraceId)),
	}
	return basicfields
}
func Info(ctx context.Context, msg string, ext ...log.Fields) {
	entry := log.WithFields(getBasicFields(ctx))
	for _, v := range ext {
		entry = entry.WithFields(v)
	}
	entry.Info(msg)
}
