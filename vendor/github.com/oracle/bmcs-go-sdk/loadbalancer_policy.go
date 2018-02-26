// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

// The properties that define a load balancer.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/LoadBalancerPolicies/

type LoadBalancerPolicy struct {
	Name string `json:"name"`
}

// ListLoadBalancerPolicies contains a list of backend policies
//
type ListLoadBalancerPolicies struct {
	OPCRequestIDUnmarshaller
	LoadBalancerPolicies []LoadBalancerPolicy
}

func (l *ListLoadBalancerPolicies) GetList() interface{} {
	return &l.LoadBalancerPolicies
}

// ListLoadBalancerPolicies Lists the available load balancer policies.
//
// See: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/LoadBalancerPolicy/ListPolicies
func (c *Client) ListLoadBalancerPolicies(
	compartmentID string,
	opts *ListLoadBalancerPolicyOptions,
) (loadbalancerPolicies *ListLoadBalancerPolicies, e error) {
	required := struct {
		CompartmentID string `header:"-" json:"-" url:"compartmentId,omitempty"`
	}{
		CompartmentID: compartmentID,
	}
	details := &requestDetails{
		name:     resourceLoadBalancerPolicies,
		required: required,
		optional: opts,
	}

	var resp *response
	if resp, e = c.loadBalancerApi.getRequest(details); e != nil {
		return
	}

	loadbalancerPolicies = &ListLoadBalancerPolicies{}
	e = resp.unmarshal(loadbalancerPolicies)
	return
}
