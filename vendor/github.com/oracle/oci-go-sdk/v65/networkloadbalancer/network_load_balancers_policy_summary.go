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

// NetworkLoadBalancersPolicySummaryEnum Enum with underlying type: string
type NetworkLoadBalancersPolicySummaryEnum string

// Set of constants representing the allowable values for NetworkLoadBalancersPolicySummaryEnum
const (
	NetworkLoadBalancersPolicySummaryTwoTuple   NetworkLoadBalancersPolicySummaryEnum = "TWO_TUPLE"
	NetworkLoadBalancersPolicySummaryThreeTuple NetworkLoadBalancersPolicySummaryEnum = "THREE_TUPLE"
	NetworkLoadBalancersPolicySummaryFiveTuple  NetworkLoadBalancersPolicySummaryEnum = "FIVE_TUPLE"
)

var mappingNetworkLoadBalancersPolicySummaryEnum = map[string]NetworkLoadBalancersPolicySummaryEnum{
	"TWO_TUPLE":   NetworkLoadBalancersPolicySummaryTwoTuple,
	"THREE_TUPLE": NetworkLoadBalancersPolicySummaryThreeTuple,
	"FIVE_TUPLE":  NetworkLoadBalancersPolicySummaryFiveTuple,
}

var mappingNetworkLoadBalancersPolicySummaryEnumLowerCase = map[string]NetworkLoadBalancersPolicySummaryEnum{
	"two_tuple":   NetworkLoadBalancersPolicySummaryTwoTuple,
	"three_tuple": NetworkLoadBalancersPolicySummaryThreeTuple,
	"five_tuple":  NetworkLoadBalancersPolicySummaryFiveTuple,
}

// GetNetworkLoadBalancersPolicySummaryEnumValues Enumerates the set of values for NetworkLoadBalancersPolicySummaryEnum
func GetNetworkLoadBalancersPolicySummaryEnumValues() []NetworkLoadBalancersPolicySummaryEnum {
	values := make([]NetworkLoadBalancersPolicySummaryEnum, 0)
	for _, v := range mappingNetworkLoadBalancersPolicySummaryEnum {
		values = append(values, v)
	}
	return values
}

// GetNetworkLoadBalancersPolicySummaryEnumStringValues Enumerates the set of values in String for NetworkLoadBalancersPolicySummaryEnum
func GetNetworkLoadBalancersPolicySummaryEnumStringValues() []string {
	return []string{
		"TWO_TUPLE",
		"THREE_TUPLE",
		"FIVE_TUPLE",
	}
}

// GetMappingNetworkLoadBalancersPolicySummaryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNetworkLoadBalancersPolicySummaryEnum(val string) (NetworkLoadBalancersPolicySummaryEnum, bool) {
	enum, ok := mappingNetworkLoadBalancersPolicySummaryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
