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
	//es.InitEsClient("http://121.36.229.98:9200", "elastic", "canda_4006889967")
	for {
		for i := 0; i < 10; i++ {
			Info(context.Background(),
				fmt.Sprint("你好", i),
				logrus.Fields{
					"user":   "dddd",
					"user33": "ddd123d",
				},
			)
		}
		time.Sleep(10 * time.Second)
	}

}
