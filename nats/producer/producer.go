package producer

import (
	"context"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func Producer(ctx context.Context) {
	nc, err := nats.Connect("nats://127.0.0.1:15000")
	if err != nil {
		log.Fatal("Failed to connect to NATS server", err)
	}
	subject := "logs"

	if err := nc.Publish(subject, []byte("Hi I am bala")); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("done")
	}
}
