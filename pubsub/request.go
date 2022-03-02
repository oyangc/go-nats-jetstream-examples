package main

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"runtime"
	"time"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal("connect error")
	}
	nc.Subscribe("data", func(m *nats.Msg) {
		result, _ := json.Marshal(m)
		fmt.Printf("<=== 【%s】%s \r\n %s", m.Subject, string(m.Data), string(result))
		nc.Publish(m.Reply, []byte("ok"))
	})
	message, err := nc.Request("data", []byte("data 测试数据"), 1*time.Second)
	if err != nil {
		log.Fatal("Request timeout error", err)
	}
	result, _ := json.Marshal(message)
	fmt.Println("message=", string(result))
	runtime.Goexit()
}
