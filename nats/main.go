package main

import (
	"context"
	"log"
	"nats/config"
	"nats/consumer"
	"nats/producer"
	"time"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())
	config.CreateServer()
	go consumer.Consumer(ctx)
	time.Sleep(2 * time.Second)

	go producer.Producer(ctx)

	time.Sleep(2 * time.Second)
	log.Println("server shutdown completed")
}

// package main

// import (
// 	"log"
// 	"nats/config"
// 	"time"

// 	"github.com/nats-io/nats.go"
// )

// func main() {

// 	config.CreateServer()
// 	nc, err := nats.Connect("nats://127.0.0.1:15000")

// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	defer nc.Close()

// 	nc.Subscribe("foo", func(msg *nats.Msg) {
// 		log.Println("Subscriber 1:", string(msg.Data))
// 	})

// 	nc.Subscribe("foo", func(msg *nats.Msg) {
// 		log.Println("Subscriber 2:", string(msg.Data))
// 	})

// 	nc.Subscribe("foo", func(msg *nats.Msg) {
// 		log.Println("Subscriber 3:", string(msg.Data))
// 	})

// 	if err := nc.Publish("foo", []byte("Message")); err != nil {
// 		log.Fatalln(err)
// 	}

// 	time.Sleep(2 * time.Second)
// }
