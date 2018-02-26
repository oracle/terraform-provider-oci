// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import (
	"net/http"
)

// The health check policy configuration.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/HealthChecker/
// Also https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/requests/HealthCheckerDetails
type HealthChecker struct {
	OPCRequestIDUnmarshaller
	OPCWorkRequestIDUnmarshaller
	Protocol string `url:"-" header:"-" json:"protocol"` // TODO: add validation in provider, must be in {"HTTP","TCP"}
	URLPath  string `url:"-" header:"-" json:"urlPath"`
	// Optional
	IntervalInMS      int    `url:"-" header:"-" json:"intervalInMillis,omitempty"`  // Default: 10000
	Port              int    `url:"-" header:"-" json:"port,omitempty"`              // Default: 0
	ResponseBodyRegex string `url:"-" header:"-" json:"responseBodyRegex,omitempty"` // Default: ".*",
	Retries           int    `url:"-" header:"-" json:"retries,omitempty"`           // Default: 3
	ReturnCode        int    `url:"-" header:"-" json:"returnCode,omitempty"`        // Default: 200
	TimeoutInMS       int    `url:"-" header:"-" json:"timeoutInMillis,omitempty"`   // Default: 3000,

}

// GetHealthChecker Gets the health check policy information for a given load balancer and backend set.
//
// See: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/HealthChecker/GetHealthChecker
func (c *Client) GetHealthChecker(
	loadBalancerID string,
	backendSetName string,
	opts *ClientRequestOptions,
) (healthChecker *HealthChecker, e error) {
	details := &requestDetails{
		ids: urlParts{resourceLoadBalancers, loadBalancerID,
			resourceBackendSets, backendSetName, resourceHealthChecker},
		optional: opts,
	}

	var resp *response
	if resp, e = c.loadBalancerApi.getRequest(details); e != nil {
		return
	}

	healthChecker = &HealthChecker{}
	e = resp.unmarshal(healthChecker)
	return
}

// UpdateHealthChecker Updates the health check policy for a given load balancer and backend set.
//
// See: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/HealthChecker/UpdateHealthChecker
func (c *Client) UpdateHealthChecker(
	loadBalancerID string,
	backendSetName string,
	healthCheckerOptions HealthChecker, // TODO: confirm that all fields are optional
	opts *LoadBalancerOptions,
) (workRequestID string, e error) {

	details := &requestDetails{
		ids: urlParts{resourceLoadBalancers, loadBalancerID,
			resourceBackendSets, backendSetName},
		required: healthCheckerOptions,
		optional: opts,
	}

	var resp *response
	if resp, e = c.objectStorageApi.request(http.MethodPut, details); e != nil {
		return
	}

	healthChecker := &HealthChecker{}
	e = resp.unmarshal(healthChecker)
	if e == nil {
		workRequestID = healthChecker.WorkRequestID
	}
	return
}
