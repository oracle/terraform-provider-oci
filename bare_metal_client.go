package main

import (
	"crypto/rsa"
)

// This will be replaced with the real client at some point?
type BareMetalClient interface {
	New(userOCID, tenancyOCID, keyFingerPrint string, privateKey *rsa.PrivateKey) BareMetalClient
	UserCreate(name, description string) (id string, err error)
}

type ProductionBareMetalClient struct{}

func (c *ProductionBareMetalClient) New(userOCID, tenancyOCID, keyFingerPrint string, privateKey *rsa.PrivateKey) BareMetalClient {
	return nil
}

func (c *ProductionBareMetalClient) UserCreate(name, description string) (id string, err error) {
	// TODO: return a random string
	return "SOME_USER_ID", nil
}
