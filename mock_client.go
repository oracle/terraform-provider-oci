package main

// MockClient is used to access Oracle BareMetal Services during tests
import "crypto/rsa"

type MockClient struct {
}

// New creates a new BareMetalClient instance
func (mock MockClient) New(userOCID, tenancyOCID, keyFingerPrint string, privateKey *rsa.PrivateKey) BareMetalClient {
	return nil
}

// UserCreate method to create an user
func (mock MockClient) UserCreate(name, description string) (id string, err error) {
	// TODO: return a random string
	return "SOME_USER_ID", nil
}
