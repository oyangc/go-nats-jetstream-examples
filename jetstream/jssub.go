package main

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"runtime"
	"time"
)

func printnatsinfo(js nats.JetStreamContext) {
	info, _ := js.StreamInfo("ORDERS")
	marshal, _ := json.Marshal(info)
	fmt.Println("===> StreamInfo ", string(marshal))

	consumerInfo, _ := js.ConsumerInfo("ORDERS", "MONITOR")
	marshal2, _ := json.Marshal(consumerInfo)
	fmt.Println("===> ConsumerInfo ", string(marshal2))
}

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal("connect error")
	}
	js, _ := nc.JetStream()
	printnatsinfo(js)
	//sub端如何delstream的话 消息会丢失
	//js.DeleteConsumer("ORDERS", "MONITOR")
	//js.DeleteStream("ORDERS")

	js.AddStream(&nats.StreamConfig{
		Name:     "ORDERS",
		Subjects: []string{"ORDERS.scratch"},
		//Subjects:  []string{"ORDERS.*"}, //jetstream不支持通配符
		Retention: nats.WorkQueuePolicy,
	})
	js.UpdateStream(&nats.StreamConfig{
		Name:     "ORDERS",
		MaxBytes: 8,
	})
	js.AddConsumer("ORDERS", &nats.ConsumerConfig{
		Durable: "MONITOR",
	})

	printnatsinfo(js)

	//2、三种订阅（选第一或第二种即可）
	/**
	第一种：js.Subscribe、这2种方式都可以"ORDERS.*"  "ORDERS.scratch"
	第二种：js.SubscribeSync可以使用("ORDERS.scratch") 不能使用("ORDERS.*")
	第三种：PullSubscribe无法订阅消息
	*/
	//第一种订阅
	//js.Subscribe、这2种方式都可以"ORDERS.*"  "ORDERS.scratch"
	sub, _ := js.Subscribe("ORDERS.*", func(m *nats.Msg) {
		//js.Subscribe("ORDERS.scratch", func(m *nats.Msg) {
		fmt.Printf("Received a JetStream message: %s\n", string(m.Data))
		m.Ack()
	})

	//第二种订阅
	//js.SubscribeSync可以使用("ORDERS.scratch") 不能使用("ORDERS.*")
	// Simple Sync Durable Consumer (optional SubOpts at the end)
	/*sub, _ := js.SubscribeSync("ORDERS.scratch", nats.Durable("MONITOR"), nats.MaxDeliver(3))
	for {
		m, _ := sub.NextMsg(3 * time.Second)
		if m != nil {
			fmt.Printf("<===  Received a JetStream message: %s\n", string(m.Data))
			m.Ack()
		}
	}*/

	//第三种订阅
	/**
	PullSubscribe无法订阅消息
	*/
	/*subscribe, err := js.PullSubscribe("ORDERS.scratch", "MONITOR")
	for {
		msgs, _ := subscribe.Fetch(3)
		fmt.Println("<=== ", msgs)
		for i, x := range msgs {
			fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
		}
		time.Sleep(time.Duration(3) * time.Second)
	}*/

	time.Sleep(time.Duration(30) * time.Second)
	// Unsubscribe
	sub.Unsubscribe()
	sub.Drain()

	runtime.Goexit()
}
