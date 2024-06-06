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
