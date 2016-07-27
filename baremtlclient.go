package main

import (
	"crypto/rsa"
)

type BareMetalClient interface {
	New(userOCID, tenancyOCID, keyFingerPrint string, privateKey *rsa.PrivateKey) BareMetalClient
	UserCreate(name, description string) (id string, err error)
}

type Client struct{}

func (c *Client) New(userOCID, tenancyOCID, keyFingerPrint string, privateKey *rsa.PrivateKey) BareMetalClient {
	return nil
}

func (c *Client) UserCreate(name, description string) (id string, err error) {
	// TODO: return a random string
	return "SOME_USER_ID", nil
}
