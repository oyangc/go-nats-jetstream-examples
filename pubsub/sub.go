package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"runtime"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal("connect error")
	}
	i := 1
	nc.Subscribe("*.>", func(m *nats.Msg) {
		i++
		log.Printf("<<<=== %d %s %s", i, m.Subject, string(m.Data))
		//nc.Publish(m.Reply, []byte("success "+strconv.Itoa(i)))
	})
	runtime.Goexit()
}
