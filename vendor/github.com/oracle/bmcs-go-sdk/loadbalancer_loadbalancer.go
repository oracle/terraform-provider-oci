// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import (
	"net/http"
)

// LoadBalancer defines a Load Balancer.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/LoadBalancer/
type LoadBalancer struct {
	OPCRequestIDUnmarshaller
	OPCWorkRequestIDUnmarshaller
	CompartmentID string                 `json:"compartmentId"`
	DisplayName   string                 `json:"displayName"`
	ID            string                 `json:"id"`
	IPAddresses   []IPAddress            `json:"ipAddresses"` // TODO: is there a better way?
	IsPrivate     bool                   `json:"isPrivate"`
	Shape         string                 `json:"shapeName"`
	State         string                 `json:"lifecycleState"`
	SubnetIDs     []string               `json:"subnetIds"`
	TimeCreated   Time                   `json:"timeCreated"`
	BackendSets   map[string]BackendSet  `json:"backendSets"`
	Certificates  map[string]Certificate `json:"certificates"`
	Listeners     map[string]Listener    `json:"listeners"`
}

type IPAddress struct {
	IPAddress string `json:"ipAddress"`
}

// ListLoadBalancers contains a list of load balancers.
type ListLoadBalancers struct {
	OPCRequestIDUnmarshaller
	NextPageUnmarshaller
	LoadBalancers []LoadBalancer
}

func (l *ListLoadBalancers) GetList() interface{} {
	return &l.LoadBalancers
}

// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/requests/CreateLoadBalancerDetails
type CreateLoadBalancerDetails struct {
	ocidRequirement
	BackendSets  *BackendSet  `header:"-" json:"backendSets,omitempty" url:"-"`
	Certificates *Certificate `header:"-" json:"certificates,omitempty" url:"-"`
	IsPrivate    bool         `header:"-" json:"isPrivate,omitempty" url:"-"`
	Listeners    *Listener    `header:"-" json:"listeners,omitempty" url:"-"`
	Shape        string       `header:"-" json:"shapeName,omitempty" url:"-"`
	SubnetIDs    []string     `header:"-" json:"subnetIds,omitempty" url:"-"`
}

// CreateLoadBalancer creates a new load balancer in the specified compartment.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/LoadBalancer/CreateLoadBalancer
func (c *Client) CreateLoadBalancer(
	backendSets *BackendSet,
	certificates *Certificate,
	compartmentID string,
	listeners *Listener,
	shape string,
	subnetIDs []string,
	opts *CreateLoadBalancerOptions) (workRequestID string, e error) {

	required := CreateLoadBalancerDetails{
		BackendSets:  backendSets,
		Certificates: certificates,
		Listeners:    listeners,
		Shape:        shape,
		SubnetIDs:    subnetIDs,
	}
	required.CompartmentID = compartmentID

	details := &requestDetails{
		name:     resourceLoadBalancers,
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.loadBalancerApi.postRequest(details); e != nil {
		return
	}

	loadbalancer := &LoadBalancer{}
	e = resp.unmarshal(loadbalancer)
	if e == nil {
		workRequestID = loadbalancer.WorkRequestID
	}
	return
}

// GetLoadBalancer gets the specified load balancer's configuration information.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/LoadBalancer/GetLoadBalancer
func (c *Client) GetLoadBalancer(id string, opts *ClientRequestOptions) (loadbalancer *LoadBalancer, e error) {
	details := &requestDetails{
		name:     resourceLoadBalancers,
		ids:      urlParts{id},
		optional: opts,
	}

	var resp *response
	if resp, e = c.loadBalancerApi.getRequest(details); e != nil {
		return
	}

	loadbalancer = &LoadBalancer{}
	e = resp.unmarshal(loadbalancer)
	return
}

// ListLoadBalancers lists all load balancers in the specified compartment.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/LoadBalancer/ListLoadBalancers
func (c *Client) ListLoadBalancers(compartmentID string, opts *ListOptions) (loadbalancers *ListLoadBalancers, e error) {
	details := &requestDetails{
		name:     resourceLoadBalancers,
		required: listOCIDRequirement{CompartmentID: compartmentID},
		optional: opts,
	}

	var resp *response
	if resp, e = c.loadBalancerApi.getRequest(details); e != nil {
		return
	}

	loadbalancers = &ListLoadBalancers{}
	e = resp.unmarshal(loadbalancers)
	return
}

// UpdateLoadBalancer updates a load balancer's configuration.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/LoadBalancer/UpdateLoadBalancer
func (c *Client) UpdateLoadBalancer(id string, opts *UpdateLoadBalancerOptions) (workRequestID string, e error) {
	// TODO: Determine if any parameters to the load balancer API are optional.
	details := &requestDetails{
		name:     resourceLoadBalancers,
		ids:      urlParts{id},
		optional: opts,
	}

	var resp *response
	if resp, e = c.loadBalancerApi.request(http.MethodPut, details); e != nil {
		return
	}

	loadbalancer := &LoadBalancer{}
	e = resp.unmarshal(loadbalancer)
	if e == nil {
		workRequestID = loadbalancer.WorkRequestID
	}
	return
}

// DeleteLoadBalancer stops a load balancer and removes it from service.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/LoadBalancer/DeleteLoadBalancer
func (c *Client) DeleteLoadBalancer(id string, opts *ClientRequestOptions) (workRequestID string, e error) {
	details := &requestDetails{
		name:     resourceLoadBalancers,
		ids:      urlParts{id},
		optional: opts,
	}

	var resp *response
	if resp, e = c.loadBalancerApi.request(http.MethodDelete, details); e != nil {
		return
	}

	loadbalancer := &LoadBalancer{}
	e = resp.unmarshal(loadbalancer)
	if e == nil {
		workRequestID = loadbalancer.WorkRequestID
	}
	return
}
