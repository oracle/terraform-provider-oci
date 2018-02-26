// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import (
	"net/http"
)

// Listener defines a listener's configuration details.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/requests/ListenerDetails
type Listener struct {
	OPCRequestIDUnmarshaller
	OPCWorkRequestIDUnmarshaller
	DefaultBackendSetName string            `header:"-" url:"-" json:"defaultBackendSetName"`
	Name                  string            `header:"-" url:"-" json:"name,omitempty"` // Only for create
	Port                  int               `header:"-" url:"-" json:"port"`
	Protocol              string            `header:"-" url:"-" json:"protocol"` // TODO: add validation in provider, For valid values see ListProtocols()
	SSLConfig             *SSLConfiguration `header:"-" url:"-" json:"sslConfiguration,omitempty"`
}

// CreateListener Adds a listener to a load balancer.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/Listener/CreateListener
func (c *Client) CreateListener(
	loadBalancerID string,
	name string,
	defaultBackendSetName string,
	protocol string,
	port int,
	sslConfig *SSLConfiguration,
	opts *LoadBalancerOptions,
) (workRequestID string, e error) {

	required := Listener{
		Name: name,
		DefaultBackendSetName: defaultBackendSetName,
		Protocol:              protocol,
		Port:                  port,
		SSLConfig:             sslConfig,
	}

	details := &requestDetails{
		name: resourceLoadBalancers,
		ids: urlParts{
			loadBalancerID,
			resourceListeners,
		},
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.loadBalancerApi.postRequest(details); e != nil {
		return
	}

	listener := &Listener{}
	e = resp.unmarshal(listener)
	if e == nil {
		workRequestID = listener.WorkRequestID
	}
	return
}

// TODO: Determine if any parameters to the load balancer API are optional.

// UpdateListener Updates a listener for a given load balancer.
//
// See: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/Listener/UpdateListener
func (c *Client) UpdateListener(
	loadBalancerID string,
	listenerName string,
	opts *UpdateLoadBalancerListenerOptions,
) (workRequestID string, e error) {

	details := &requestDetails{
		name: resourceLoadBalancers,
		ids: urlParts{
			loadBalancerID,
			resourceListeners,
			listenerName,
		},
		optional: opts,
	}

	var resp *response
	if resp, e = c.loadBalancerApi.request(http.MethodPut, details); e != nil {
		return
	}

	listener := &Listener{}
	e = resp.unmarshal(listener)
	if e == nil {
		workRequestID = listener.WorkRequestID
	}
	return
}

// Deletes a listener from a load balancer.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/Listener/DeleteListener
func (c *Client) DeleteListener(
	loadBalancerID string,
	listenerName string,
	opts *ClientRequestOptions,
) (workRequestID string, e error) {

	details := &requestDetails{
		name: resourceLoadBalancers,
		ids: urlParts{
			loadBalancerID,
			resourceListeners,
			listenerName,
		},
		optional: opts,
	}

	var resp *response
	if resp, e = c.loadBalancerApi.request(http.MethodDelete, details); e != nil {
		return
	}

	listener := &Listener{}
	e = resp.unmarshal(listener)
	if e == nil {
		workRequestID = listener.WorkRequestID
	}
	return
}
