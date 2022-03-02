package main

import (
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

	nc.Subscribe("ORDERS.*", func(m *nats.Msg) {
		fmt.Printf("Received a JetStream message: %s\n", string(m.Data))
	})
	runtime.Goexit()
}
