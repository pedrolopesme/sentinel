package client

import (
	"fmt"
	"github.com/nats-io/go-nats"
	"os"
)

type NATSServer interface {
	GetConnection() *nats.Conn
}

type BasicNATSServer struct {
	Conn *nats.Conn
}

func (n *BasicNATSServer) GetConnection() *nats.Conn {
	return n.Conn
}

// TODO add some logging
// TODO add some tests
func NewNATSServer(uri string) (server *BasicNATSServer, err error) {
	fmt.Printf("Trying to connect nats server %v\n", uri)
	natsConnection, err := nats.Connect(uri)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return &BasicNATSServer{
		Conn: natsConnection,
	}, nil
}
