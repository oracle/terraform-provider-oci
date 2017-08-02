// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import (
	"net/http"
)

// Backend defines a backend server that is a member of a load balancer backend set.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/Backend/
// Also https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/requests/BackendDetails
type Backend struct {
	OPCRequestIDUnmarshaller
	OPCWorkRequestIDUnmarshaller
	Backup    bool   `json:"backup"`
	Drain     bool   `json:"drain"`
	IPAddress string `json:"ipAddress"`
	Offline   bool   `json:"offline"`
	Port      int    `json:"port"`
	Weight    int    `json:"weight"`
}

// ListBackends contains a list of backends
//
type ListBackends struct {
	OPCRequestIDUnmarshaller
	Backends []Backend
}

func (l *ListBackends) GetList() interface{} {
	return &l.Backends
}

// CreateBackend Adds a backend server to a backend set.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/Backend/CreateBackend
func (c *Client) CreateBackend(
	loadBalancerID string,
	backendSetName string,
	ipAddr string,
	port int,
	opts *CreateLoadBalancerBackendOptions,
) (workRequestID string, e error) {

	required := struct {
		IPAddr string `header:"-" json:"ipAddress" url:"-"`
		Port   int    `header:"-" json:"port" url:"-"`
	}{
		IPAddr: ipAddr,
		Port:   port,
	}

	details := &requestDetails{
		name:     resourceLoadBalancers,
		ids:      urlParts{loadBalancerID, resourceBackendSets, backendSetName, resourceBackends},
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.loadBalancerApi.postRequest(details); e != nil {
		return
	}

	backend := &Backend{}
	e = resp.unmarshal(backend)
	if e == nil {
		workRequestID = backend.WorkRequestID
	}
	return
}

// GetBackend Gets the specified backend server's configuration information.
//
// See: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/Backend/GetBackend
func (c *Client) GetBackend(
	loadBalancerID string,
	backendSetName string,
	backendName string,
	opts *ClientRequestOptions,
) (backend *Backend, e error) {
	details := &requestDetails{
		name: resourceLoadBalancers,
		ids: urlParts{loadBalancerID,
			resourceBackendSets, backendSetName,
			resourceBackends, backendName},
		optional: opts,
	}

	var resp *response
	if resp, e = c.loadBalancerApi.getRequest(details); e != nil {
		return
	}

	backend = &Backend{}
	e = resp.unmarshal(backend)
	return
}

// ListBackend Lists the backend servers for a given load balancer and backend set.
//
// See: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/Backend/ListBackends
func (c *Client) ListBackends(
	loadBalancerID string,
	backendSetName string,
) (backends *ListBackends, e error) {
	details := &requestDetails{
		name: resourceLoadBalancers,
		ids: urlParts{loadBalancerID,
			resourceBackendSets, backendSetName, resourceBackends},
	}

	var resp *response
	if resp, e = c.loadBalancerApi.getRequest(details); e != nil {
		return
	}

	backends = &ListBackends{}
	e = resp.unmarshal(backends)
	return
}

// UpdateBackend Updates the configuration of a backend server within the specified backend set.
//
// See: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/Backend/UpdateBackend
func (c *Client) UpdateBackend(
	loadBalancerID string,
	backendSetName string,
	backendName string,
	opts *UpdateLoadBalancerBackendOptions,
) (workRequestID string, e error) {

	details := &requestDetails{
		name: resourceLoadBalancers,
		ids: urlParts{loadBalancerID,
			resourceBackendSets, backendSetName,
			resourceBackends, backendName},
		optional: opts,
	}

	var resp *response
	if resp, e = c.loadBalancerApi.request(http.MethodPut, details); e != nil {
		return
	}

	backend := &Backend{}
	e = resp.unmarshal(backend)
	if e == nil {
		workRequestID = backend.WorkRequestID
	}
	return
}

// DeleteBackend Removes a backend server from a given load balancer and backend set.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/Backend/DeleteBackend
func (c *Client) DeleteBackend(
	loadBalancerID string,
	backendSetName string,
	backendName string,
	opts *ClientRequestOptions,
) (workRequestID string, e error) {

	details := &requestDetails{
		name: resourceLoadBalancers,
		ids: urlParts{loadBalancerID,
			resourceBackendSets, backendSetName,
			resourceBackends, backendName},
		optional: opts,
	}

	var resp *response
	if resp, e = c.loadBalancerApi.request(http.MethodDelete, details); e != nil {
		return
	}

	backend := &Backend{}
	e = resp.unmarshal(backend)
	if e == nil {
		workRequestID = backend.WorkRequestID
	}
	return
}
