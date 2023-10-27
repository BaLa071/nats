package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)
//this file is not need for basic nats pushlisg and subscribe .....
func main() {
	// config.CreateServer()
	nc, err := nats.Connect("nats://0.0.0.0:4222")
	if err != nil {
		log.Fatal("Failed to connect to NATS server:", err)
	}
	defer nc.Close()

	fmt.Println("Connected to NATS server on port 15000")

	js, _ := nc.JetStream()
	var kv nats.KeyValue
	if stream, _ := js.StreamInfo("KV_discovery"); stream == nil {
		kv, _ = js.CreateKeyValue(&nats.KeyValueConfig{
			Bucket: "discovery",
		})
	} else {
		kv, _ = js.KeyValue("discovery")
	}
	
	kv.Put("services.orders", []byte("bala"))
	entry, _ := kv.Get("services.orders")
	fmt.Printf("%s @ %d -> %q\n", entry.Key(), entry.Revision(), string(entry.Value()))
}
