// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

// TA shape is a template that determines the total pre-provisioned bandwidth (ingress plus egress) for the load balancer.
// Note that the pre-provisioned maximum capacity applies to aggregated connections, not to a single client attempting to use the full bandwidth.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/LoadBalancerShape/

type LoadBalancerShape struct {
	Name string `json:"name"`
}

// ListLoadBalancerShapes contains a list of shapes
//
type ListLoadBalancerShapes struct {
	OPCRequestIDUnmarshaller
	LoadBalancerShapes []LoadBalancerShape
}

func (l *ListLoadBalancerShapes) GetList() interface{} {
	return &l.LoadBalancerShapes
}

// ListLoadBalancerShapes Lists the valid load balancer shapes.
//
// See: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/LoadBalancerShape/ListShapes
func (c *Client) ListLoadBalancerShapes(
	compartmentID string,
	opts *ListLoadBalancerPolicyOptions,
) (loadbalancerShapes *ListLoadBalancerShapes, e error) {
	required := struct {
		CompartmentID string `header:"-" json:"-" url:"compartmentId,omitempty"`
	}{
		CompartmentID: compartmentID,
	}
	details := &requestDetails{
		name:     resourceLoadBalancerShapes,
		required: required,
		optional: opts,
	}

	var resp *response
	if resp, e = c.loadBalancerApi.getRequest(details); e != nil {
		return
	}

	loadbalancerShapes = &ListLoadBalancerShapes{}
	e = resp.unmarshal(loadbalancerShapes)
	return
}
