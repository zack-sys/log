package log

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/zack-sys/log/es"
	"testing"
	"time"
)

func TestInfo(t *testing.T) {
	SetIndex("test")
	es.InitEsClient("http://127.0.0.1:9200")
	for {
		for i := 0; i < 500; i++ {
			Info(context.Background(),
				fmt.Sprint("你好", i),
				logrus.Fields{
					"user": "dddd",
				},
			)
		}
		time.Sleep(10 * time.Second)
	}

}
