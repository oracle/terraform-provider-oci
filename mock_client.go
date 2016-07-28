package main

import "github.com/stretchr/testify/mock"

type MockClient struct {
	mock.Mock
}

func (m *MockClient) CreateUser(name, description string) (id string, e error) {
	args := m.Called(name, description)
	return args.String(0), args.Error(1)
}
