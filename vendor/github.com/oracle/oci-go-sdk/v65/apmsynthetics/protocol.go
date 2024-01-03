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

// ProtocolEnum Enum with underlying type: string
type ProtocolEnum string

// Set of constants representing the allowable values for ProtocolEnum
const (
	ProtocolIcmp ProtocolEnum = "ICMP"
	ProtocolTcp  ProtocolEnum = "TCP"
)

var mappingProtocolEnum = map[string]ProtocolEnum{
	"ICMP": ProtocolIcmp,
	"TCP":  ProtocolTcp,
}

var mappingProtocolEnumLowerCase = map[string]ProtocolEnum{
	"icmp": ProtocolIcmp,
	"tcp":  ProtocolTcp,
}

// GetProtocolEnumValues Enumerates the set of values for ProtocolEnum
func GetProtocolEnumValues() []ProtocolEnum {
	values := make([]ProtocolEnum, 0)
	for _, v := range mappingProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetProtocolEnumStringValues Enumerates the set of values in String for ProtocolEnum
func GetProtocolEnumStringValues() []string {
	return []string{
		"ICMP",
		"TCP",
	}
}

// GetMappingProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProtocolEnum(val string) (ProtocolEnum, bool) {
	enum, ok := mappingProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
