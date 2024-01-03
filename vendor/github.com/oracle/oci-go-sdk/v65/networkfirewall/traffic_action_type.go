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

// TrafficActionTypeEnum Enum with underlying type: string
type TrafficActionTypeEnum string

// Set of constants representing the allowable values for TrafficActionTypeEnum
const (
	TrafficActionTypeAllow   TrafficActionTypeEnum = "ALLOW"
	TrafficActionTypeDrop    TrafficActionTypeEnum = "DROP"
	TrafficActionTypeReject  TrafficActionTypeEnum = "REJECT"
	TrafficActionTypeInspect TrafficActionTypeEnum = "INSPECT"
)

var mappingTrafficActionTypeEnum = map[string]TrafficActionTypeEnum{
	"ALLOW":   TrafficActionTypeAllow,
	"DROP":    TrafficActionTypeDrop,
	"REJECT":  TrafficActionTypeReject,
	"INSPECT": TrafficActionTypeInspect,
}

var mappingTrafficActionTypeEnumLowerCase = map[string]TrafficActionTypeEnum{
	"allow":   TrafficActionTypeAllow,
	"drop":    TrafficActionTypeDrop,
	"reject":  TrafficActionTypeReject,
	"inspect": TrafficActionTypeInspect,
}

// GetTrafficActionTypeEnumValues Enumerates the set of values for TrafficActionTypeEnum
func GetTrafficActionTypeEnumValues() []TrafficActionTypeEnum {
	values := make([]TrafficActionTypeEnum, 0)
	for _, v := range mappingTrafficActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTrafficActionTypeEnumStringValues Enumerates the set of values in String for TrafficActionTypeEnum
func GetTrafficActionTypeEnumStringValues() []string {
	return []string{
		"ALLOW",
		"DROP",
		"REJECT",
		"INSPECT",
	}
}

// GetMappingTrafficActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTrafficActionTypeEnum(val string) (TrafficActionTypeEnum, bool) {
	enum, ok := mappingTrafficActionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
