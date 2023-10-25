package config

import (
	"log"

	natsServer "github.com/nats-io/nats-server/v2/server"
)

func CreateServer() {

	opts := &natsServer.Options{
		ServerName:     "local_nats_server",
		Host:           "localhost",
		Port:           15000,
		NoLog:          false,
		NoSigs:         false,
		MaxControlLine: 4096,
		MaxPayload:     65536,
	}

	server, err := natsServer.NewServer(opts)
	if err != nil {
		log.Fatal(err)
	}
	err = natsServer.Run(server)
	if err != nil {
		log.Fatal("Failed to start NATS server:", err)
	}
	log.Println("NATS server started!")
}
