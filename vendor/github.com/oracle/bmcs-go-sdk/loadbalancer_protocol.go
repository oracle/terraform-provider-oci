// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

// The properties that define a load balancer.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/LoadBalancerProtocols/

type LoadBalancerProtocol struct {
	Name string `json:"name"`
}

// ListLoadBalancerProtocols contains a list of protocols
//
type ListLoadBalancerProtocols struct {
	OPCRequestIDUnmarshaller
	LoadBalancerProtocols []LoadBalancerProtocol
}

func (l *ListLoadBalancerProtocols) GetList() interface{} {
	return &l.LoadBalancerProtocols
}

// ListLoadBalancerProtocols Lists the available load balancer policies.
//
// See: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/LoadBalancerPolicy/ListProtocols
func (c *Client) ListLoadBalancerProtocols(
	compartmentID string,
	opts *ListLoadBalancerPolicyOptions,
) (loadbalancerProtocols *ListLoadBalancerProtocols, e error) {
	required := struct {
		CompartmentID string `header:"-" json:"-" url:"compartmentId,omitempty"`
	}{
		CompartmentID: compartmentID,
	}
	details := &requestDetails{
		name:     resourceLoadBalancerProtocols,
		required: required,
		optional: opts,
	}

	var resp *response
	if resp, e = c.loadBalancerApi.getRequest(details); e != nil {
		return
	}

	loadbalancerProtocols = &ListLoadBalancerProtocols{}
	e = resp.unmarshal(loadbalancerProtocols)
	return
}
