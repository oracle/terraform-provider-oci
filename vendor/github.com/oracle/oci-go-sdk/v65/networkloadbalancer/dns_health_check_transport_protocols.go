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

// DnsHealthCheckTransportProtocolsEnum Enum with underlying type: string
type DnsHealthCheckTransportProtocolsEnum string

// Set of constants representing the allowable values for DnsHealthCheckTransportProtocolsEnum
const (
	DnsHealthCheckTransportProtocolsUdp DnsHealthCheckTransportProtocolsEnum = "UDP"
	DnsHealthCheckTransportProtocolsTcp DnsHealthCheckTransportProtocolsEnum = "TCP"
)

var mappingDnsHealthCheckTransportProtocolsEnum = map[string]DnsHealthCheckTransportProtocolsEnum{
	"UDP": DnsHealthCheckTransportProtocolsUdp,
	"TCP": DnsHealthCheckTransportProtocolsTcp,
}

var mappingDnsHealthCheckTransportProtocolsEnumLowerCase = map[string]DnsHealthCheckTransportProtocolsEnum{
	"udp": DnsHealthCheckTransportProtocolsUdp,
	"tcp": DnsHealthCheckTransportProtocolsTcp,
}

// GetDnsHealthCheckTransportProtocolsEnumValues Enumerates the set of values for DnsHealthCheckTransportProtocolsEnum
func GetDnsHealthCheckTransportProtocolsEnumValues() []DnsHealthCheckTransportProtocolsEnum {
	values := make([]DnsHealthCheckTransportProtocolsEnum, 0)
	for _, v := range mappingDnsHealthCheckTransportProtocolsEnum {
		values = append(values, v)
	}
	return values
}

// GetDnsHealthCheckTransportProtocolsEnumStringValues Enumerates the set of values in String for DnsHealthCheckTransportProtocolsEnum
func GetDnsHealthCheckTransportProtocolsEnumStringValues() []string {
	return []string{
		"UDP",
		"TCP",
	}
}

// GetMappingDnsHealthCheckTransportProtocolsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDnsHealthCheckTransportProtocolsEnum(val string) (DnsHealthCheckTransportProtocolsEnum, bool) {
	enum, ok := mappingDnsHealthCheckTransportProtocolsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
