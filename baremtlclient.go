package main

import "github.com/MustWin/baremtlclient"

type BareMetalClient interface {
	CreateUser(name, description string, options ...baremtlclient.Options) (*baremtlclient.Resource, error)
	GetUser(userID string) (*baremtlclient.Resource, error)
	UpdateUser(userID string)
}

type Client struct{}

func (c *Client) CreateUser(name, description string, options ...baremtlclient.Options) (*baremtlclient.Resource, error) {
	panic("")
}

func (c *Client) GetUser(userID string) (*baremtlclient.Resource, error) {
	panic("")
}

func (c *Client) UpdateUser(userID string) {
	panic("")
}
