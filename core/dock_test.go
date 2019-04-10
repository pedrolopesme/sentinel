package core

import (
	"github.com/stretchr/testify/mock"
)

type MockedDock struct {
	mock.Mock
}

func (md *MockedDock) GetId() string {
	args := md.Called()
	return args.String(0)
}

func (md *MockedDock) Watch() (err error) {
	args := md.Called()
	return args.Error(0)
}
