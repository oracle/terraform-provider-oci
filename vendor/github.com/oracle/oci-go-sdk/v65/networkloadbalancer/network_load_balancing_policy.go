// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// NetworkLoadBalancer API
//
// This describes the network load balancer API.
//

package networkloadbalancer

import (
	"strings"
)

// NetworkLoadBalancingPolicyEnum Enum with underlying type: string
type NetworkLoadBalancingPolicyEnum string

// Set of constants representing the allowable values for NetworkLoadBalancingPolicyEnum
const (
	NetworkLoadBalancingPolicyTwoTuple   NetworkLoadBalancingPolicyEnum = "TWO_TUPLE"
	NetworkLoadBalancingPolicyThreeTuple NetworkLoadBalancingPolicyEnum = "THREE_TUPLE"
	NetworkLoadBalancingPolicyFiveTuple  NetworkLoadBalancingPolicyEnum = "FIVE_TUPLE"
)

var mappingNetworkLoadBalancingPolicyEnum = map[string]NetworkLoadBalancingPolicyEnum{
	"TWO_TUPLE":   NetworkLoadBalancingPolicyTwoTuple,
	"THREE_TUPLE": NetworkLoadBalancingPolicyThreeTuple,
	"FIVE_TUPLE":  NetworkLoadBalancingPolicyFiveTuple,
}

var mappingNetworkLoadBalancingPolicyEnumLowerCase = map[string]NetworkLoadBalancingPolicyEnum{
	"two_tuple":   NetworkLoadBalancingPolicyTwoTuple,
	"three_tuple": NetworkLoadBalancingPolicyThreeTuple,
	"five_tuple":  NetworkLoadBalancingPolicyFiveTuple,
}

// GetNetworkLoadBalancingPolicyEnumValues Enumerates the set of values for NetworkLoadBalancingPolicyEnum
func GetNetworkLoadBalancingPolicyEnumValues() []NetworkLoadBalancingPolicyEnum {
	values := make([]NetworkLoadBalancingPolicyEnum, 0)
	for _, v := range mappingNetworkLoadBalancingPolicyEnum {
		values = append(values, v)
	}
	return values
}

// GetNetworkLoadBalancingPolicyEnumStringValues Enumerates the set of values in String for NetworkLoadBalancingPolicyEnum
func GetNetworkLoadBalancingPolicyEnumStringValues() []string {
	return []string{
		"TWO_TUPLE",
		"THREE_TUPLE",
		"FIVE_TUPLE",
	}
}

// GetMappingNetworkLoadBalancingPolicyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNetworkLoadBalancingPolicyEnum(val string) (NetworkLoadBalancingPolicyEnum, bool) {
	enum, ok := mappingNetworkLoadBalancingPolicyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
