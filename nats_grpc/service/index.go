package service

import (
	"context"
	"fmt"
	"log"
	pro "nats_grpc/proto"

	"github.com/nats-io/nats.go"
)

type Server struct {
	pro.UnimplementedFileTransferServiceServer
}

func (s *Server) TransferFile(ctx context.Context, req *pro.Request) (*pro.Response, error) {
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
	ids:="1"
	kv.Put(ids, []byte(req.FileName))

	return &pro.Response{Id:ids},nil
}
