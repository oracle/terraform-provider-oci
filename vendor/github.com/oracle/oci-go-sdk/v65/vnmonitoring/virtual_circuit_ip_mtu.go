// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"strings"
)

// VirtualCircuitIpMtuEnum Enum with underlying type: string
type VirtualCircuitIpMtuEnum string

// Set of constants representing the allowable values for VirtualCircuitIpMtuEnum
const (
	VirtualCircuitIpMtuMtu1500 VirtualCircuitIpMtuEnum = "MTU_1500"
	VirtualCircuitIpMtuMtu9000 VirtualCircuitIpMtuEnum = "MTU_9000"
)

var mappingVirtualCircuitIpMtuEnum = map[string]VirtualCircuitIpMtuEnum{
	"MTU_1500": VirtualCircuitIpMtuMtu1500,
	"MTU_9000": VirtualCircuitIpMtuMtu9000,
}

var mappingVirtualCircuitIpMtuEnumLowerCase = map[string]VirtualCircuitIpMtuEnum{
	"mtu_1500": VirtualCircuitIpMtuMtu1500,
	"mtu_9000": VirtualCircuitIpMtuMtu9000,
}

// GetVirtualCircuitIpMtuEnumValues Enumerates the set of values for VirtualCircuitIpMtuEnum
func GetVirtualCircuitIpMtuEnumValues() []VirtualCircuitIpMtuEnum {
	values := make([]VirtualCircuitIpMtuEnum, 0)
	for _, v := range mappingVirtualCircuitIpMtuEnum {
		values = append(values, v)
	}
	return values
}

// GetVirtualCircuitIpMtuEnumStringValues Enumerates the set of values in String for VirtualCircuitIpMtuEnum
func GetVirtualCircuitIpMtuEnumStringValues() []string {
	return []string{
		"MTU_1500",
		"MTU_9000",
	}
}

// GetMappingVirtualCircuitIpMtuEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVirtualCircuitIpMtuEnum(val string) (VirtualCircuitIpMtuEnum, bool) {
	enum, ok := mappingVirtualCircuitIpMtuEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
