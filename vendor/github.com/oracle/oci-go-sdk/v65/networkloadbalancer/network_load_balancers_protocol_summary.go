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

// NetworkLoadBalancersProtocolSummaryEnum Enum with underlying type: string
type NetworkLoadBalancersProtocolSummaryEnum string

// Set of constants representing the allowable values for NetworkLoadBalancersProtocolSummaryEnum
const (
	NetworkLoadBalancersProtocolSummaryAny       NetworkLoadBalancersProtocolSummaryEnum = "ANY"
	NetworkLoadBalancersProtocolSummaryTcp       NetworkLoadBalancersProtocolSummaryEnum = "TCP"
	NetworkLoadBalancersProtocolSummaryUdp       NetworkLoadBalancersProtocolSummaryEnum = "UDP"
	NetworkLoadBalancersProtocolSummaryTcpAndUdp NetworkLoadBalancersProtocolSummaryEnum = "TCP_AND_UDP"
	NetworkLoadBalancersProtocolSummaryL3Ip      NetworkLoadBalancersProtocolSummaryEnum = "L3IP"
)

var mappingNetworkLoadBalancersProtocolSummaryEnum = map[string]NetworkLoadBalancersProtocolSummaryEnum{
	"ANY":         NetworkLoadBalancersProtocolSummaryAny,
	"TCP":         NetworkLoadBalancersProtocolSummaryTcp,
	"UDP":         NetworkLoadBalancersProtocolSummaryUdp,
	"TCP_AND_UDP": NetworkLoadBalancersProtocolSummaryTcpAndUdp,
	"L3IP":        NetworkLoadBalancersProtocolSummaryL3Ip,
}

var mappingNetworkLoadBalancersProtocolSummaryEnumLowerCase = map[string]NetworkLoadBalancersProtocolSummaryEnum{
	"any":         NetworkLoadBalancersProtocolSummaryAny,
	"tcp":         NetworkLoadBalancersProtocolSummaryTcp,
	"udp":         NetworkLoadBalancersProtocolSummaryUdp,
	"tcp_and_udp": NetworkLoadBalancersProtocolSummaryTcpAndUdp,
	"l3ip":        NetworkLoadBalancersProtocolSummaryL3Ip,
}

// GetNetworkLoadBalancersProtocolSummaryEnumValues Enumerates the set of values for NetworkLoadBalancersProtocolSummaryEnum
func GetNetworkLoadBalancersProtocolSummaryEnumValues() []NetworkLoadBalancersProtocolSummaryEnum {
	values := make([]NetworkLoadBalancersProtocolSummaryEnum, 0)
	for _, v := range mappingNetworkLoadBalancersProtocolSummaryEnum {
		values = append(values, v)
	}
	return values
}

// GetNetworkLoadBalancersProtocolSummaryEnumStringValues Enumerates the set of values in String for NetworkLoadBalancersProtocolSummaryEnum
func GetNetworkLoadBalancersProtocolSummaryEnumStringValues() []string {
	return []string{
		"ANY",
		"TCP",
		"UDP",
		"TCP_AND_UDP",
		"L3IP",
	}
}

// GetMappingNetworkLoadBalancersProtocolSummaryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNetworkLoadBalancersProtocolSummaryEnum(val string) (NetworkLoadBalancersProtocolSummaryEnum, bool) {
	enum, ok := mappingNetworkLoadBalancersProtocolSummaryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
