package main

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"runtime"
	"strconv"
	"time"
)

func main() {
	// Connect to NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal("connect error")
	}

	// Create JetStream Context
	js, _ := nc.JetStream(nats.PublishAsyncMaxPending(256))
	//js, _ := nc.JetStream()

	js.DeleteConsumer("ORDERS", "MONITOR")
	js.DeleteStream("ORDERS")

	js.AddStream(&nats.StreamConfig{
		Name:     "ORDERS",
		Subjects: []string{"ORDERS.scratch"},
		//Subjects: []string{"ORDERS.*"},//jetstream不支持通配符
		Retention: nats.WorkQueuePolicy,
	})
	js.UpdateStream(&nats.StreamConfig{
		Name:     "ORDERS",
		MaxBytes: 8,
	})
	js.AddConsumer("ORDERS", &nats.ConsumerConfig{ //存消息
		Durable: "MONITOR",
	})
	//打印信息
	info, _ := js.StreamInfo("ORDERS")
	marshal, _ := json.Marshal(info)
	fmt.Println("===> StreamInfo ", string(marshal))

	consumerInfo, _ := js.ConsumerInfo("ORDERS", "MONITOR")
	marshal2, _ := json.Marshal(consumerInfo)
	fmt.Println("===> ConsumerInfo ", string(marshal2))

	// Simple Stream Publisher
	js.Publish("ORDERS.scratch", []byte("hello"))

	// Simple Async Stream Publisher
	max := 50
	for i := 0; i < max; i++ {
		js.PublishAsync("ORDERS.scratch", []byte("hello "+strconv.Itoa(i)))
		time.Sleep(time.Duration(500) * time.Millisecond)
	}

	runtime.Goexit()
}
