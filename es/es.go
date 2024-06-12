package es

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/olivere/elastic/v7/config"
	"github.com/zack-sys/log/util"
)

var (
	defaultEsClient *elastic.Client
	index           string
)

func SetEsIndex(i string) {
	index = i
}

func InitEsClient(host string, name, pass string) {
	if host == "" {
		fmt.Println("es host is empty")
		return
	}
	// 创建es连接
	var sniff = false
	cfg := &config.Config{
		URL:      host,
		Username: name,
		Password: pass,
		Sniff:    &sniff,
	}
	var err error
	defaultEsClient, err = elastic.NewClientFromConfig(cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println("es client init success")
}

// es数据推送
func PushEs(ctx context.Context, buffer []string) (err error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("es client panic err:", err)
		}
	}()
	if defaultEsClient == nil {
		return nil
	}
	if len(buffer) == 0 {
		return nil
	}
	bulk := defaultEsClient.Bulk()
	for _, v := range buffer {
		if len(v) == 0 {
			continue
		}
		var a map[string]interface{}
		err := json.Unmarshal([]byte(v), &a)
		if err != nil {
			fmt.Println("es bulk unmarshal err:", err.Error())
			continue
		}
		result, err := json.Marshal(a)
		if err != nil {
			fmt.Println("es bulk Marshal err:", err.Error())
			continue
		}

		req := elastic.NewBulkIndexRequest().Doc(string(result))
		bulk = bulk.Add(req)
	}

	_, err = bulk.Index(util.GetIndex(index)).Type("_doc").Do(ctx)
	if err != nil {
		return err
	}
	return nil
}
