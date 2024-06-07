package log

import (
	"context"
	uuid "github.com/satori/go.uuid"
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
	log.SetOutput(write.NewLog())
	log.SetLevel(log.InfoLevel)
	skipCall = 2
}
func SetIndex(i string) {
	index = i
	es.SetEsIndex(i)
}
func SetConsolePrint(flag bool) {
	log.SetOutput(write.NewLogCfg(&write.Log{Console: flag}))
}

func SkipCall(skip int) {
	skipCall = skip
}

func getBasicFields(ctx context.Context) log.Fields {
	if index == "" {
		panic("索引不能为空")
	}

	hostname, _ := os.Hostname()
	basicfields := log.Fields{
		"caller":       zapcore.NewEntryCaller(runtime.Caller(skipCall)).TrimmedPath(),
		"@fluentd_tag": util.GetIndex(index),
		"Trace_Id":     util.GetString(ctx.Value(enum.TraceId)),
		"os":           runtime.GOOS,
		"hostname":     hostname,
		"pid":          os.Getpid(),
		"server_host":  util.GetUseIp(),
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
func Trace(ctx context.Context, msg string, ext ...log.Fields) {
	entry := log.WithFields(getBasicFields(ctx))
	for _, v := range ext {
		entry = entry.WithFields(v)
	}
	entry.Trace(msg)
}
func Debug(ctx context.Context, msg string, ext ...log.Fields) {
	entry := log.WithFields(getBasicFields(ctx))
	for _, v := range ext {
		entry = entry.WithFields(v)
	}
	entry.Debug(msg)
}
func Warn(ctx context.Context, msg string, ext ...log.Fields) {
	entry := log.WithFields(getBasicFields(ctx))
	for _, v := range ext {
		entry = entry.WithFields(v)
	}
	entry.Warn(msg)
}
func Error(ctx context.Context, msg string, ext ...log.Fields) {
	entry := log.WithFields(getBasicFields(ctx))
	for _, v := range ext {
		entry = entry.WithFields(v)
	}
	entry.Error(msg)
}
func Fatal(ctx context.Context, msg string, ext ...log.Fields) {
	entry := log.WithFields(getBasicFields(ctx))
	for _, v := range ext {
		entry = entry.WithFields(v)
	}
	entry.Fatal(msg)
}
func Panic(ctx context.Context, msg string, ext ...log.Fields) {
	entry := log.WithFields(getBasicFields(ctx))
	for _, v := range ext {
		entry = entry.WithFields(v)
	}
	entry.Panic(msg)
}

func NewContext() context.Context {
	ctx := context.Background()
	context.WithValue(ctx, enum.TraceId, uuid.NewV4().String())
	return ctx
}
