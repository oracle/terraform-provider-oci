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

// TunnelInspectionProtocolEnum Enum with underlying type: string
type TunnelInspectionProtocolEnum string

// Set of constants representing the allowable values for TunnelInspectionProtocolEnum
const (
	TunnelInspectionProtocolVxlan TunnelInspectionProtocolEnum = "VXLAN"
)

var mappingTunnelInspectionProtocolEnum = map[string]TunnelInspectionProtocolEnum{
	"VXLAN": TunnelInspectionProtocolVxlan,
}

var mappingTunnelInspectionProtocolEnumLowerCase = map[string]TunnelInspectionProtocolEnum{
	"vxlan": TunnelInspectionProtocolVxlan,
}

// GetTunnelInspectionProtocolEnumValues Enumerates the set of values for TunnelInspectionProtocolEnum
func GetTunnelInspectionProtocolEnumValues() []TunnelInspectionProtocolEnum {
	values := make([]TunnelInspectionProtocolEnum, 0)
	for _, v := range mappingTunnelInspectionProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetTunnelInspectionProtocolEnumStringValues Enumerates the set of values in String for TunnelInspectionProtocolEnum
func GetTunnelInspectionProtocolEnumStringValues() []string {
	return []string{
		"VXLAN",
	}
}

// GetMappingTunnelInspectionProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTunnelInspectionProtocolEnum(val string) (TunnelInspectionProtocolEnum, bool) {
	enum, ok := mappingTunnelInspectionProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
