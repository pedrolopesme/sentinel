package client

import (
	stan "github.com/nats-io/go-nats-streaming"
	"github.com/stretchr/testify/mock"
)

type MockedNATSServer struct {
	mock.Mock
}

func (ns *MockedNATSServer) GetConnection() stan.Conn {
	args := ns.Called()
	return args.Get(0).(stan.Conn)
}
