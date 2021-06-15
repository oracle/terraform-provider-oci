// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// NetworkLoadBalancer API
//
// A description of the network load balancer API
//

package networkloadbalancer

// NetworkLoadBalancersPolicySummaryEnum Enum with underlying type: string
type NetworkLoadBalancersPolicySummaryEnum string

// Set of constants representing the allowable values for NetworkLoadBalancersPolicySummaryEnum
const (
	NetworkLoadBalancersPolicySummaryTwoTuple   NetworkLoadBalancersPolicySummaryEnum = "TWO_TUPLE"
	NetworkLoadBalancersPolicySummaryThreeTuple NetworkLoadBalancersPolicySummaryEnum = "THREE_TUPLE"
	NetworkLoadBalancersPolicySummaryFiveTuple  NetworkLoadBalancersPolicySummaryEnum = "FIVE_TUPLE"
)

var mappingNetworkLoadBalancersPolicySummary = map[string]NetworkLoadBalancersPolicySummaryEnum{
	"TWO_TUPLE":   NetworkLoadBalancersPolicySummaryTwoTuple,
	"THREE_TUPLE": NetworkLoadBalancersPolicySummaryThreeTuple,
	"FIVE_TUPLE":  NetworkLoadBalancersPolicySummaryFiveTuple,
}

// GetNetworkLoadBalancersPolicySummaryEnumValues Enumerates the set of values for NetworkLoadBalancersPolicySummaryEnum
func GetNetworkLoadBalancersPolicySummaryEnumValues() []NetworkLoadBalancersPolicySummaryEnum {
	values := make([]NetworkLoadBalancersPolicySummaryEnum, 0)
	for _, v := range mappingNetworkLoadBalancersPolicySummary {
		values = append(values, v)
	}
	return values
}
