// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Firewall API
//
// Use the Network Firewall API to create network firewalls and configure policies that regulates network traffic in and across VCNs.
//

package networkfirewall

import (
	"strings"
)

// ServiceTypeEnum Enum with underlying type: string
type ServiceTypeEnum string

// Set of constants representing the allowable values for ServiceTypeEnum
const (
	ServiceTypeTcpService ServiceTypeEnum = "TCP_SERVICE"
	ServiceTypeUdpService ServiceTypeEnum = "UDP_SERVICE"
)

var mappingServiceTypeEnum = map[string]ServiceTypeEnum{
	"TCP_SERVICE": ServiceTypeTcpService,
	"UDP_SERVICE": ServiceTypeUdpService,
}

var mappingServiceTypeEnumLowerCase = map[string]ServiceTypeEnum{
	"tcp_service": ServiceTypeTcpService,
	"udp_service": ServiceTypeUdpService,
}

// GetServiceTypeEnumValues Enumerates the set of values for ServiceTypeEnum
func GetServiceTypeEnumValues() []ServiceTypeEnum {
	values := make([]ServiceTypeEnum, 0)
	for _, v := range mappingServiceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetServiceTypeEnumStringValues Enumerates the set of values in String for ServiceTypeEnum
func GetServiceTypeEnumStringValues() []string {
	return []string{
		"TCP_SERVICE",
		"UDP_SERVICE",
	}
}

// GetMappingServiceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingServiceTypeEnum(val string) (ServiceTypeEnum, bool) {
	enum, ok := mappingServiceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
