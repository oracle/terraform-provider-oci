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

func (m *MockClient) GetUser(userID string) (*baremtlclient.Resource, error) {
	args := m.Called(userID)
	u, _ := args.Get(0).(*baremtlclient.Resource)
	return u, args.Error(1)
}

func (m *MockClient) UpdateUser(userID string) {
	m.Called(userID)
}
