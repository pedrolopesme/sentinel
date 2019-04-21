package client

import (
	"fmt"
	"os"
	"time"

	stan "github.com/nats-io/go-nats-streaming"
)

type NATSServer interface {
	GetConnection() stan.Conn
}

type BasicNATSServer struct {
	Conn stan.Conn
}

func (n *BasicNATSServer) GetConnection() stan.Conn {
	return n.Conn
}

// TODO add some logging
// TODO add some tests
func NewNATSServer(clusterID string, clientID string, uri string) (server *BasicNATSServer, err error) {
	fmt.Printf("Trying to connect to NATS server %v - %v\n", clusterID, uri)
	beforeConnect := time.Now()
	natsConnection, err := stan.Connect(clusterID, clientID, stan.NatsURL(uri))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("Connected to NATS server %v\n", time.Since(beforeConnect).String())

	return &BasicNATSServer{
		Conn: natsConnection,
	}, nil
}
