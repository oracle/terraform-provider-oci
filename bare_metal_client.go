package main

import (
	"crypto/rsa"
)

// Client base struct for the baremetal clients
type Client struct {
}

// BareMetalClient is an interface used to injecting a mock instance for unit tests
type BareMetalClient interface {
	New(userOCID, tenancyOCID, keyFingerPrint string, privateKey *rsa.PrivateKey) (c *Client)
	UserCreate(name, description string) (id string, err error)
}

// MockClient is used to access Oracle BareMetal Services during tests
type MockClient struct {
}

// New creates a new BareMetalClient instance
func (mock MockClient) New(userOCID, tenancyOCID, keyFingerPrint string, privateKey *rsa.PrivateKey) (c *Client) {
	return nil
}

// UserCreate method to create an user
func (mock MockClient) UserCreate(name, description string) (id string, err error) {
	// TODO: return a random string
	return "SOME_USER_ID", nil
}
