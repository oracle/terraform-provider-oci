package main

import (
	"github.com/MustWin/baremtlclient"
	"github.com/stretchr/testify/mock"
)

type MockClient struct {
	mock.Mock
}

func (m *MockClient) CreateUser(name, description string, options ...baremtlclient.Options) (*baremtlclient.Resource, error) {
	args := m.Called(name, description)
	return args.Get(0).(*baremtlclient.Resource), args.Error(1)
}
