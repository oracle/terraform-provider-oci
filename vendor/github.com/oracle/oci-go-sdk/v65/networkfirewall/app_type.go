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

// AppTypeEnum Enum with underlying type: string
type AppTypeEnum string

// Set of constants representing the allowable values for AppTypeEnum
const (
	AppTypeIcmp   AppTypeEnum = "ICMP"
	AppTypeIcmpV6 AppTypeEnum = "ICMP_V6"
)

var mappingAppTypeEnum = map[string]AppTypeEnum{
	"ICMP":    AppTypeIcmp,
	"ICMP_V6": AppTypeIcmpV6,
}

var mappingAppTypeEnumLowerCase = map[string]AppTypeEnum{
	"icmp":    AppTypeIcmp,
	"icmp_v6": AppTypeIcmpV6,
}

// GetAppTypeEnumValues Enumerates the set of values for AppTypeEnum
func GetAppTypeEnumValues() []AppTypeEnum {
	values := make([]AppTypeEnum, 0)
	for _, v := range mappingAppTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAppTypeEnumStringValues Enumerates the set of values in String for AppTypeEnum
func GetAppTypeEnumStringValues() []string {
	return []string{
		"ICMP",
		"ICMP_V6",
	}
}

// GetMappingAppTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAppTypeEnum(val string) (AppTypeEnum, bool) {
	enum, ok := mappingAppTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
