package util

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
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

func DeepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}
func DeepCopyJson(dst, src interface{}) error {
	marshal, err := json.Marshal(src)
	if err != nil {
		return err
	}
	err = json.Unmarshal(marshal, dst)
	if err != nil {
		return err
	}
	return nil
}
