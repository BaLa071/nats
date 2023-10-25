package main

import (
	"context"
	"fmt"
	"log"

	pro "nats_grpc/proto"

	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	} else {
		log.Println("connected")
	}

	client := pro.NewFileTransferServiceClient(conn)
	Request := &pro.Request{
		FileName: "balaji",
	}
	getFileResponse, err := client.TransferFile(context.Background(), Request)
	if err != nil {
		log.Fatalf("GetFile error: %v", err)
	}
	log.Println(getFileResponse.Id)





	//fetching value from jetstream
	nc, err := nats.Connect("nats://0.0.0.0:4222")
	if err != nil {
		log.Fatal("Failed to connect to NATS server:", err)
	}
	defer nc.Close()

	fmt.Println("Connected to NATS server on port 15000")

	js, _ := nc.JetStream()
	var kv nats.KeyValue
	if stream, _ := js.StreamInfo("Sample"); stream == nil {
		kv, _ = js.CreateKeyValue(&nats.KeyValueConfig{
			Bucket: "test_transfer",
		})
	} else {
		kv, _ = js.KeyValue("test_transfer")
	}

	entry,_:=kv.Get(getFileResponse.Id)
	// fmt.Println("File is :"res.Value())
	fmt.Printf("%s @ %d -> %q\n", entry.Key(), entry.Revision(), string(entry.Value()))
	// fmt.Println(entry.Value())
}
