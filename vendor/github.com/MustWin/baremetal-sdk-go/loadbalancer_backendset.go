// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import (
	"net/http"
)

// BackendSet defines the configuration of a load balancer backend set.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/BackendSet/
type BackendSet struct {
	OPCRequestIDUnmarshaller
	OPCWorkRequestIDUnmarshaller
	Backends                 []Backend                        `json:"backends" url:"-"`
	HealthChecker            *HealthChecker                   `json:"healthChecker" url:"-"`
	Name                     string                           `json:"name,omitempty" url:"-"`             // Only on create
	Policy                   string                           `json:"policy" url:"-"`                     // FIXME: supposedly has default: "ROUND_ROBIN" but then raises error when null. For valid values see ListPolicies()
	SSLConfig                *SSLConfiguration                `json:"sslConfiguration,omitempty" url:"-"` // TODO: acc test, waiting on CreateCertificate() tests
	SessionPersistenceConfig *SessionPersistenceConfiguration `json:"sessionPersistenceConfiguration,omitempty" url:"-"`
}

type SSLConfiguration struct {
	CertificateName       string `json:"certificateName"`
	VerifyDepth           int    `json:"verifyDepth"`
	VerifyPeerCertificate bool   `json:"verifyPeerCertificate"`
}

type SessionPersistenceConfiguration struct {
	CookieName      string `json:"cookieName"`
	DisableFallback bool   `json:"disableFallback"`
}

// ListBackendSets contains a list of backend Sets
//
type ListBackendSets struct {
	OPCRequestIDUnmarshaller
	BackendSets []BackendSet
}

func (l *ListBackendSets) GetList() interface{} {
	return &l.BackendSets
}

// CreateBackendSet Adds a backend set to a load balancer.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/BackendSet/CreateBackendSet
func (c *Client) CreateBackendSet(
	loadBalancerID string,
	name string,
	policy string,
	backends []Backend,
	healthChecker *HealthChecker,
	sslConfig *SSLConfiguration,
	sessionPersistenceConfig *SessionPersistenceConfiguration,
	opts *LoadBalancerOptions,
) (workRequestID string, e error) {
	required := BackendSet{
		Name:                     name,
		Policy:                   policy,
		SSLConfig:                sslConfig,
		SessionPersistenceConfig: sessionPersistenceConfig,
		HealthChecker:            healthChecker,
		Backends:                 backends,
	}

	details := &requestDetails{
		name: resourceLoadBalancers,
		ids: urlParts{
			loadBalancerID,
			resourceBackendSets,
		},
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.loadBalancerApi.postRequest(details); e != nil {
		return
	}

	backendset := &BackendSet{}
	e = resp.unmarshal(backendset)
	if e == nil {
		workRequestID = backendset.WorkRequestID
	}
	return
}

// GetBackendSet Gets the specified backend set's configuration information.
//
// See: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/BackendSet/GetBackendSet
func (c *Client) GetBackendSet(
	loadBalancerID string,
	backendSetName string,
	opts *ClientRequestOptions,
) (backendset *BackendSet, e error) {
	details := &requestDetails{
		name: resourceLoadBalancers,
		ids: urlParts{
			loadBalancerID,
			resourceBackendSets,
			backendSetName,
		},
		optional: opts,
	}

	var resp *response
	if resp, e = c.loadBalancerApi.getRequest(details); e != nil {
		return
	}

	backendset = &BackendSet{}
	e = resp.unmarshal(backendset)
	return
}

// ListBackendSets Lists all backend sets associated with a given load balancer.
//
// See: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/BackendSet/ListBackendSets
func (c *Client) ListBackendSets(
	loadBalancerID string,
	opts *ClientRequestOptions,
) (backends *ListBackendSets, e error) {
	details := &requestDetails{
		name: resourceLoadBalancers,
		ids: urlParts{
			loadBalancerID,
			resourceBackendSets,
		},
	}

	var resp *response
	if resp, e = c.loadBalancerApi.getRequest(details); e != nil {
		return
	}

	backends = &ListBackendSets{}
	e = resp.unmarshal(backends)
	return
}

// TODO: Determine if any parameters to the load balancer API are optional.

// UpdateBackendSet Updates a backend set.
//
// See: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/BackendSet/UpdateBackendSet
func (c *Client) UpdateBackendSet(
	loadBalancerID string,
	backendSetName string,
	opts *UpdateLoadBalancerBackendSetOptions,
) (workRequestID string, e error) {

	details := &requestDetails{
		name: resourceLoadBalancers,
		ids: urlParts{
			loadBalancerID,
			resourceBackendSets,
			backendSetName,
		},
		optional: opts,
	}

	var resp *response
	if resp, e = c.loadBalancerApi.request(http.MethodPut, details); e != nil {
		return
	}

	backendset := &BackendSet{}
	e = resp.unmarshal(backendset)
	if e == nil {
		workRequestID = backendset.WorkRequestID
	}
	return
}

// Deletes the specified backend set. Note that deleting a backend set removes its backend servers from the load balancer.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/BackendSet/DeleteBackendSet
func (c *Client) DeleteBackendSet(
	loadBalancerID string,
	backendSetName string,
	opts *ClientRequestOptions,
) (workRequestID string, e error) {

	details := &requestDetails{
		name: resourceLoadBalancers,
		ids: urlParts{
			loadBalancerID,
			resourceBackendSets,
			backendSetName,
		},
		optional: opts,
	}

	var resp *response
	if resp, e = c.loadBalancerApi.request(http.MethodDelete, details); e != nil {
		return
	}

	backendset := &BackendSet{}
	e = resp.unmarshal(backendset)

	if e == nil {
		workRequestID = backendset.WorkRequestID
	}
	return
}
