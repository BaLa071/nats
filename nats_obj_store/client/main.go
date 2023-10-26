package main

import (
	"context"
	"fmt"
	"log"
	pro "obj_store/proto"
	"os"

	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Println("Can't dial grpc client: ", err)
	}

	client := pro.NewObjGrpcServiceClient(conn)
	getFileRequest := &pro.GetFileRequest{
		FileName: "/home/balaji/go/src/testint/test3.jpg",
	}

	getFileResponse, err := client.GetFile(context.Background(), getFileRequest)
	if err != nil {
		log.Fatalf("GetFile error: %v", err)
	}
	log.Println(getFileResponse)

	GetFileFromNats(getFileResponse.FileID)

}

func GetFileFromNats(FileID string) {
	nc, err := nats.Connect("nats://0.0.0.0:4222")
	if err != nil {
		log.Fatal("can't connect ", err)
	}
	defer nc.Close()

	js, err := nc.JetStream()
	if err != nil {
		log.Fatal("can't connect to js ", err)
	}

	var obj nats.ObjectStore
	if stream, _ := js.StreamInfo("testos"); stream == nil {
		obj, _ = js.CreateObjectStore(&nats.ObjectStoreConfig{
			Bucket: "newos",
		})
		fmt.Println("1")
	} else {
		log.Println("obj store not exists: ", err)
	}

	res, err := obj.GetBytes(FileID)
	if err != nil {
		log.Println("error getting file in object store: ", err)
	}

	localfile, err1 := os.Create("/home/balaji/go/src/fun.jpg")
	if err1 != nil {
		log.Println("err ", err)
	}
	_, err2 := localfile.Write(res)
	if err2 != nil {
		log.Println("err2 ", err2)
	}

	fmt.Println("\nGot File, Content: \n", string(res))

}
