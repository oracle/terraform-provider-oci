// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmsynthetics

import (
	"strings"
)

// DnsTransportProtocolEnum Enum with underlying type: string
type DnsTransportProtocolEnum string

// Set of constants representing the allowable values for DnsTransportProtocolEnum
const (
	DnsTransportProtocolTcp DnsTransportProtocolEnum = "TCP"
	DnsTransportProtocolUdp DnsTransportProtocolEnum = "UDP"
)

var mappingDnsTransportProtocolEnum = map[string]DnsTransportProtocolEnum{
	"TCP": DnsTransportProtocolTcp,
	"UDP": DnsTransportProtocolUdp,
}

var mappingDnsTransportProtocolEnumLowerCase = map[string]DnsTransportProtocolEnum{
	"tcp": DnsTransportProtocolTcp,
	"udp": DnsTransportProtocolUdp,
}

// GetDnsTransportProtocolEnumValues Enumerates the set of values for DnsTransportProtocolEnum
func GetDnsTransportProtocolEnumValues() []DnsTransportProtocolEnum {
	values := make([]DnsTransportProtocolEnum, 0)
	for _, v := range mappingDnsTransportProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetDnsTransportProtocolEnumStringValues Enumerates the set of values in String for DnsTransportProtocolEnum
func GetDnsTransportProtocolEnumStringValues() []string {
	return []string{
		"TCP",
		"UDP",
	}
}

// GetMappingDnsTransportProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDnsTransportProtocolEnum(val string) (DnsTransportProtocolEnum, bool) {
	enum, ok := mappingDnsTransportProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
