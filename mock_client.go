package main

import (
	"github.com/MustWin/baremtlclient"
	"github.com/stretchr/testify/mock"
)

type MockClient struct {
	mock.Mock
}

func (m *MockClient) CreateUser(name, description string, options ...baremtlsdk.Options) (*baremtlsdk.Resource, error) {
	args := m.Called(name, description)
	return args.Get(0).(*baremtlsdk.Resource), args.Error(1)
}

func (m *MockClient) GetUser(userID string) (*baremtlsdk.Resource, error) {
	args := m.Called(userID)
	u, _ := args.Get(0).(*baremtlsdk.Resource)
	return u, args.Error(1)
}

func (m *MockClient) UpdateUser(userID, userDescription string, opts ...baremtlsdk.Options) (*baremtlsdk.Resource, error) {
	args := m.Called(userID, userDescription)
	u, _ := args.Get(0).(*baremtlsdk.Resource)
	return u, args.Error(1)
}

func (m *MockClient) DeleteUser(userID string, opts ...baremtlsdk.Options) error {
	args := m.Called(userID)
	return args.Error(0)
}
