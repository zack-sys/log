package queue

import (
	"context"
	"fmt"
	"github.com/zack-sys/log/enum"
	"github.com/zack-sys/log/es"
	"github.com/zack-sys/log/util"
	"sync"
	"time"
)

var msgChannel chan string

func init() {
	msgChannel = make(chan string, enum.QueueLen)
	go Consumption()
}
func Push(msg string) {
	msgChannel <- msg
}

func Consumption() {
	timer := time.Tick(time.Duration(enum.MsgTimeOut) * time.Second)

	msgQueue := make([]string, 0)
	lock := sync.RWMutex{}
	for {
		select {
		case <-timer:
			// 数据写入es
			if len(msgQueue) == 0 {
				continue
			}
			lock.Lock()
			temp := make([]string, len(msgQueue))
			util.DeepCopyJson(&temp, msgQueue)
			msgQueue = make([]string, 0)
			lock.Unlock()

			go func() {
				err := es.PushEs(context.Background(), temp)
				if err != nil {
					fmt.Println("es.PushEs err:", err, "len:", len(temp))
				}
			}()
		case msg := <-msgChannel:
			lock.Lock()
			msgQueue = append(msgQueue, msg)

			if len(msgQueue) >= enum.MsgLen {
				// 数据写入es
				temp := make([]string, len(msgQueue))
				util.DeepCopyJson(&temp, msgQueue)
				msgQueue = make([]string, 0)

				go func() {
					err := es.PushEs(context.Background(), temp)
					if err != nil {
						fmt.Println("es.PushEs err:", err, "len:", len(temp))
					}
				}()
			}
			lock.Unlock()
		}
	}
}
