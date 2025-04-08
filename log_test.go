package log

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestInfo(t *testing.T) {
	SetIndex("test")
	SetConsolePrint(true)
	ext := make(map[string]interface{}, 0)
	ext["kkk"] = "vvv"
	ext["asdasd"] = "asd"
	SetExtField(ext)
	logrus.SetLevel(logrus.WarnLevel)
	//es.InitEsClient("http://121.36.229.98:9200", "elastic", "canda_4006889967")
	for {
		for i := 0; i < 10; i++ {
			Info(context.Background(),
				fmt.Sprint("你好", i),
				logrus.Fields{
					"user":   "dddd",
					"user33": "ddd12133a啊啊1ddd,3d",
				},
			)
			Error(context.Background(),
				fmt.Sprint("你好", i),
				logrus.Fields{
					"user":   "dddd",
					"user33": "ddd12133a啊啊1ddd,3d",
				},
			)
		}
		time.Sleep(10 * time.Second)
	}

}
