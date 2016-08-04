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

func (m *MockClient) GetUser(id string) (*baremtlsdk.Resource, error) {
	args := m.Called(id)
	u, _ := args.Get(0).(*baremtlsdk.Resource)
	return u, args.Error(1)
}

func (m *MockClient) UpdateUser(id, description string, opts ...baremtlsdk.Options) (*baremtlsdk.Resource, error) {
	args := m.Called(id, description)
	u, _ := args.Get(0).(*baremtlsdk.Resource)
	return u, args.Error(1)
}

func (m *MockClient) DeleteUser(id string, opts ...baremtlsdk.Options) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockClient) CreateGroup(name, description string, options ...baremtlsdk.Options) (*baremtlsdk.Resource, error) {
	args := m.Called(name, description)
	return args.Get(0).(*baremtlsdk.Resource), args.Error(1)
}

func (m *MockClient) GetGroup(id string) (*baremtlsdk.Resource, error) {
	args := m.Called(id)
	u, _ := args.Get(0).(*baremtlsdk.Resource)
	return u, args.Error(1)
}

func (m *MockClient) UpdateGroup(id, description string, opts ...baremtlsdk.Options) (*baremtlsdk.Resource, error) {
	args := m.Called(id, description)
	u, _ := args.Get(0).(*baremtlsdk.Resource)
	return u, args.Error(1)
}

func (m *MockClient) DeleteGroup(id string, opts ...baremtlsdk.Options) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockClient) CreatePolicy(name, description string, statements []string, options ...baremtlsdk.Options) (*baremtlsdk.Policy, error) {
	args := m.Called(name, description, statements)
	return args.Get(0).(*baremtlsdk.Policy), args.Error(1)
}

func (m *MockClient) GetPolicy(id string) (*baremtlsdk.Policy, error) {
	args := m.Called(id)
	u, _ := args.Get(0).(*baremtlsdk.Policy)
	return u, args.Error(1)
}

func (m *MockClient) UpdatePolicy(id, description string, statements []string, opts ...baremtlsdk.Options) (*baremtlsdk.Policy, error) {
	args := m.Called(id, description, statements)
	u, _ := args.Get(0).(*baremtlsdk.Policy)
	return u, args.Error(1)
}

func (m *MockClient) DeletePolicy(id string, opts ...baremtlsdk.Options) error {
	args := m.Called(id)
	return args.Error(0)
}
