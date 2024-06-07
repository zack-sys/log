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
	es.InitEsClient("", "", "")
	for {
		for i := 0; i < 10; i++ {
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
