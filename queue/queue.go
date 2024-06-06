package queue

import (
	"context"
	"fmt"
	"github.com/zack-sys/log/enum"
	"github.com/zack-sys/log/es"
	"time"
)

var msgChannel chan []byte

func init() {
	msgChannel = make(chan []byte, enum.QueueLen)
	go Consumption()
}
func Push(msg []byte) {
	msgChannel <- msg
}

func Consumption() {
	timer := time.Tick(time.Duration(enum.MsgTimeOut) * time.Second)

	msgQueue := make([][]byte, 0)
	for {
		select {
		case <-timer:
			// 数据写入es
			if len(msgQueue) == 0 {
				continue
			}
			temp := make([][]byte, len(msgQueue))
			copy(temp, msgQueue)
			go func() {
				err := es.PushEs(context.Background(), temp)
				if err != nil {
					fmt.Println("es.PushEs err:", err)
				}
			}()
			msgQueue = make([][]byte, 0)
		case msg := <-msgChannel:
			msgQueue = append(msgQueue, msg)
			if len(msgQueue) >= enum.MsgLen {
				// 数据写入es
				temp := make([][]byte, len(msgQueue))
				copy(temp, msgQueue)
				go func() {
					err := es.PushEs(context.Background(), temp)
					if err != nil {
						fmt.Println("es.PushEs err:", err)
					}
				}()
				msgQueue = make([][]byte, 0)
			}
		}
	}
}
