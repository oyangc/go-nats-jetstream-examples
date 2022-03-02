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
	nc.Subscribe("*.>", func(m *nats.Msg) {
		result, _ := json.Marshal(m)
		fmt.Printf("<=== 1【%s】%s \r\n %s", m.Subject, string(m.Data), string(result))
		nc.Publish(m.Reply, []byte("ok-2"))
	})
	nc.Subscribe("NATS_ROOM.>", func(m *nats.Msg) {
		result, _ := json.Marshal(m)
		fmt.Printf("<=== 2【%s】%s \r\n %s", m.Subject, string(m.Data), string(result))
		nc.Publish(m.Reply, []byte("ok-3"))
	})
	message, err := nc.Request("NATS_ROOM.xxx.yyy", []byte("data 测试数据"), 1*time.Second)
	if err != nil {
		log.Fatal("Request timeout error", err)
	}
	result, _ := json.Marshal(message)
	fmt.Printf("<<<=== 【request response】data=%s     message=%s", string(message.Data), string(result))
	//nc.Publish("NATS_ROOM.xxx", []byte("测试数据"))
	//nc.Publish("data", []byte("data 测试数据"))
	runtime.Goexit()
}
