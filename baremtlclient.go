package main

import "time"

type BareMetalClient interface {
	CreateUser(name, description string) (*BareMetalIdentity, error)
}

type BareMetalIdentity struct {
	ID            string
	Name          string
	Description   string
	CompartmentID string
	State         string
	TimeModified  time.Time
	TimeCreated   time.Time
}

type Client struct{}

func (c *Client) CreateUser(name, description string) (*BareMetalIdentity, error) {
	panic("CreateUser called on temporary BareMetalClient stub!")
}
