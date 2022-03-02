package main

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"runtime"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal("connect error")
	}
	nc.Subscribe("data", func(m *nats.Msg) {
		result, _ := json.Marshal(m)
		fmt.Printf("<=== 【%s】%s \r\n %s", m.Subject, string(m.Data), string(result))
		//time.Sleep(5 * time.Second)
		nc.Publish(m.Reply, []byte("success"))
	})
	runtime.Goexit()
}
