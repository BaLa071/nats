package consumer

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func Consumer(ctx context.Context) {
	nc, err := nats.Connect("nats://127.0.0.1:15000")
	if err != nil {
		log.Fatal("Failed to connect to NATS server:", err)
	}
	defer nc.Close()

	fmt.Println("Connected to NATS server on port 15000")

	subject := "logs"
	// messages := make(chan *nats.Msg, 1000)
	fmt.Println("1")
	_, err1 := nc.Subscribe(subject, func(msg *nats.Msg) {
		fmt.Println("2")
		log.Println("Subscriber :", string(msg.Data))
	})
	fmt.Println("3")
	if err1 != nil {
		log.Fatal("Failed to subscribe to subject:", err)
	}
	log.Println("Subscribed to", subject)

	time.Sleep(10 * time.Second)

	// for {
	// 	select {
	// 	case <-ctx.Done():
	// 		log.Println("exiting from consumer")
	// 	case msg := <-messages:
	// 		log.Println("received", string(msg.Data))
	// 		return
	// 	}
	// }

}
