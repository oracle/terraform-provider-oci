// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// NetworkLoadBalancer API
//
// A description of the network load balancer API
//

package networkloadbalancer

// NetworkLoadBalancersProtocolSummaryEnum Enum with underlying type: string
type NetworkLoadBalancersProtocolSummaryEnum string

// Set of constants representing the allowable values for NetworkLoadBalancersProtocolSummaryEnum
const (
	NetworkLoadBalancersProtocolSummaryAny NetworkLoadBalancersProtocolSummaryEnum = "ANY"
	NetworkLoadBalancersProtocolSummaryTcp NetworkLoadBalancersProtocolSummaryEnum = "TCP"
	NetworkLoadBalancersProtocolSummaryUdp NetworkLoadBalancersProtocolSummaryEnum = "UDP"
)

var mappingNetworkLoadBalancersProtocolSummary = map[string]NetworkLoadBalancersProtocolSummaryEnum{
	"ANY": NetworkLoadBalancersProtocolSummaryAny,
	"TCP": NetworkLoadBalancersProtocolSummaryTcp,
	"UDP": NetworkLoadBalancersProtocolSummaryUdp,
}

// GetNetworkLoadBalancersProtocolSummaryEnumValues Enumerates the set of values for NetworkLoadBalancersProtocolSummaryEnum
func GetNetworkLoadBalancersProtocolSummaryEnumValues() []NetworkLoadBalancersProtocolSummaryEnum {
	values := make([]NetworkLoadBalancersProtocolSummaryEnum, 0)
	for _, v := range mappingNetworkLoadBalancersProtocolSummary {
		values = append(values, v)
	}
	return values
}
