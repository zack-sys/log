package util

import (
	"net"
	"strings"
	"time"
)

func GetString(a any) string {
	e, ok := a.(string)
	if ok {
		return e
	}
	return ""
}
func GetIndex(index string) string {
	return index + "-" + time.Now().Format("2006-01-02")
}
func GetUseIp() string {
	dial, err := net.Dial("udp", "8.8.8.8:80") // Google的公共DNS服务器
	if err != nil {
		return "127.0.0.1"
	}
	addr := dial.LocalAddr().String()

	index := strings.LastIndex(addr, ":")
	return addr[:index]
}
