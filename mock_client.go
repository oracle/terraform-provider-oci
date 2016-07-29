package main

import "github.com/stretchr/testify/mock"

type MockClient struct {
	mock.Mock
}

func (m *MockClient) CreateUser(name, description string) (*BareMetalIdentity, error) {
	args := m.Called(name, description)
	return args.Get(0).(*BareMetalIdentity), args.Error(1)
}
