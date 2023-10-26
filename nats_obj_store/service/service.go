package service

import (
	"bufio"
	"context"
	"fmt"
	"log"
	pro "obj_store/proto"
	"os"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

type Server struct {
	pro.UnimplementedObjGrpcServiceServer
}


var metadata nats.ObjectMeta

func CreateByteSlice(file *os.File) []byte {

	//getting the file size
	stat, err := file.Stat()
	if err != nil {
		log.Println("error getting file size: ", err)
	}

	bs := make([]byte, stat.Size())
	//reading file into byte slice
	_, err1 := bufio.NewReader(file).Read(bs)
	if err1 != nil {
		log.Println("error converting byte slice: ", err)
	}
	return bs
}

func (s *Server) GetFile(ctx context.Context, req *pro.GetFileRequest) (*pro.GetFileResponse, error) {
	nc, err := nats.Connect("nats://0.0.0.0:4222")
	if err != nil {
		log.Fatal("Failed to connect to NATS server:", err)
	}
	defer nc.Close()

	js, err := nc.JetStream()
	if err != nil {
		log.Fatal("Failed to connect jetstream ", err)
	}

	filePath := req.FileName
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("can't open local file: ", err)
	}
	defer file.Close()

	//creating file into byte slice
	data := CreateByteSlice(file)
	var obj nats.ObjectStore

	if stream, _ := js.StreamInfo("testos"); stream == nil {
		obj, _ = js.CreateObjectStore(&nats.ObjectStoreConfig{
			Bucket: "newos",
		})
		fmt.Println("1")
	} else {
		obj, _ = js.ObjectStore("newos")
		fmt.Println("2")
	}

	metadata = nats.ObjectMeta{
		Name: uuid.New().String(),
	}
	// var temp *nats.ObjectInfo
	info, err := obj.PutBytes(metadata.Name, data)
	if err != nil {
		log.Println("Can't put file in object store: ", err)
	}
	// info, err := obj.Put(&metadata, reader)
	// if err != nil {
	// 	log.Println("Can't put file in object store: ", err)
	// }
	log.Println("obj info: ", info.ObjectMeta)
	return &pro.GetFileResponse{
		FileID: metadata.Name,
	}, nil

}
