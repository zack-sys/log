package log

import (
	"context"
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/zack-sys/log/enum"
	"github.com/zack-sys/log/es"
	"github.com/zack-sys/log/util"
	"github.com/zack-sys/log/write"
	"go.uber.org/zap/zapcore"
	"os"
	"runtime"
	"strings"
	"time"
)

var index string
var skipCall int
var nowIp string

func init() {
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000Z07:00",
	})
	log.SetOutput(write.NewLog())
	log.SetLevel(log.InfoLevel)
	skipCall = 2
	nowIp = util.GetUseIp()
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
func SetLogUseIp(ip string) {
	nowIp = ip
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
		"server_host":  nowIp,
		"timestamp":    time.Now().UnixMilli(),
	}
	return basicfields
}
func PreSend(ctx context.Context, ext ...log.Fields) *log.Entry {
	entry := log.WithFields(getBasicFields(ctx))
	for _, v := range ext {
		entry = entry.WithFields(v)
	}
	marshal, _ := json.Marshal(entry.Data)

	entry = entry.WithFields(log.Fields{
		"size": len(strings.Split(string(marshal), "")),
	})
	return entry
}
func Info(ctx context.Context, msg string, ext ...log.Fields) {
	entry := PreSend(ctx, ext...)
	entry.Info(msg)
}
func Trace(ctx context.Context, msg string, ext ...log.Fields) {
	entry := PreSend(ctx, ext...)
	entry.Trace(msg)
}
func Debug(ctx context.Context, msg string, ext ...log.Fields) {
	entry := PreSend(ctx, ext...)
	entry.Debug(msg)
}
func Warn(ctx context.Context, msg string, ext ...log.Fields) {
	entry := PreSend(ctx, ext...)
	entry.Warn(msg)
}
func Error(ctx context.Context, msg string, ext ...log.Fields) {
	entry := PreSend(ctx, ext...)
	entry.Error(msg)
}
func Fatal(ctx context.Context, msg string, ext ...log.Fields) {
	entry := PreSend(ctx, ext...)
	entry.Fatal(msg)
}
func Panic(ctx context.Context, msg string, ext ...log.Fields) {
	entry := PreSend(ctx, ext...)
	entry.Panic(msg)
}

func NewContext() context.Context {
	ctx := context.Background()
	context.WithValue(ctx, enum.TraceId, uuid.NewV4().String())
	return ctx
}
