package main

type BareMetalClient interface {
	CreateUser(name, description string) (id string, e error)
}

type Client struct{}

func (c *Client) CreateUser(name, description string) (id string, e error) {
	panic("CreateUser called on temporary BareMetalClient stub!")
}
