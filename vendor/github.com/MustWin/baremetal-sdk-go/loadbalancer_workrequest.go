// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import (
	"time"
)

// Many of the API requests you use to create and configure load balancing do not take effect
// immediately. In these cases, the request spawns an asynchronous work flow to fulfill the request.
// WorkRequest objects provide visibility for in-progress work flows.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/WorkRequest/

type WorkRequest struct {
	OPCRequestIDUnmarshaller
	OPCWorkRequestIDUnmarshaller
	ID             string `json:"id"`
	ErrorDetails   []WorkRequestError
	State          string    `json:"lifecycleState"`
	LoadBalancerID string    `json:"loadBalancerId"`
	Message        string    `json:"message"`
	TimeAccepted   time.Time `json:"timeAccepted"`
	TimeFinished   time.Time `json:"timeFinished"`
	Type           string    `json:"type"`
}

type WorkRequestError struct {
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
}

// ListWorkRequest contains a list of backend Sets
//
type ListWorkRequests struct {
	OPCRequestIDUnmarshaller
	NextPageUnmarshaller
	WorkRequests []WorkRequest
}

func (l *ListWorkRequests) GetList() interface{} {
	return &l.WorkRequests
}

// GetWorkRequest Gets the details of a work request.
//
// See: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/WorkRequest/GetWorkRequest
func (c *Client) GetWorkRequest(
	workRequestID string,
	opts *ClientRequestOptions,
) (workRequest *WorkRequest, e error) {
	details := &requestDetails{
		name:     resourceLoadBalancerWorkRequests,
		ids:      urlParts{workRequestID},
		optional: opts,
	}

	var resp *response
	if resp, e = c.loadBalancerApi.getRequest(details); e != nil {
		return
	}

	workRequest = &WorkRequest{}
	e = resp.unmarshal(workRequest)
	return
}

// ListWorkRequests Lists the work requests for a given load balancer.
//
// See: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/BackendSet/ListWorkRequest
func (c *Client) ListWorkRequests(
	loadBalancerID string,
	opts *ListLoadBalancerPolicyOptions,
) (workRequests *ListWorkRequests, e error) {
	details := &requestDetails{
		ids: urlParts{resourceLoadBalancers, loadBalancerID,
			resourceWorkRequests},
		optional: opts,
	}

	var resp *response
	if resp, e = c.loadBalancerApi.getRequest(details); e != nil {
		return
	}

	workRequests = &ListWorkRequests{}
	e = resp.unmarshal(workRequests)
	return
}
