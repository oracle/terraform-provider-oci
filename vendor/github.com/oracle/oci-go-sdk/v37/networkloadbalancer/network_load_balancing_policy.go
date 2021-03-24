// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// NetworkLoadBalancer API
//
// A description of the network load balancer API
//

package networkloadbalancer

// NetworkLoadBalancingPolicyEnum Enum with underlying type: string
type NetworkLoadBalancingPolicyEnum string

// Set of constants representing the allowable values for NetworkLoadBalancingPolicyEnum
const (
	NetworkLoadBalancingPolicyTwoTuple   NetworkLoadBalancingPolicyEnum = "TWO_TUPLE"
	NetworkLoadBalancingPolicyThreeTuple NetworkLoadBalancingPolicyEnum = "THREE_TUPLE"
	NetworkLoadBalancingPolicyFiveTuple  NetworkLoadBalancingPolicyEnum = "FIVE_TUPLE"
)

var mappingNetworkLoadBalancingPolicy = map[string]NetworkLoadBalancingPolicyEnum{
	"TWO_TUPLE":   NetworkLoadBalancingPolicyTwoTuple,
	"THREE_TUPLE": NetworkLoadBalancingPolicyThreeTuple,
	"FIVE_TUPLE":  NetworkLoadBalancingPolicyFiveTuple,
}

// GetNetworkLoadBalancingPolicyEnumValues Enumerates the set of values for NetworkLoadBalancingPolicyEnum
func GetNetworkLoadBalancingPolicyEnumValues() []NetworkLoadBalancingPolicyEnum {
	values := make([]NetworkLoadBalancingPolicyEnum, 0)
	for _, v := range mappingNetworkLoadBalancingPolicy {
		values = append(values, v)
	}
	return values
}
