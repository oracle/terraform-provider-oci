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

// InspectionTypeEnum Enum with underlying type: string
type InspectionTypeEnum string

// Set of constants representing the allowable values for InspectionTypeEnum
const (
	InspectionTypeSslInboundInspection InspectionTypeEnum = "SSL_INBOUND_INSPECTION"
	InspectionTypeSslForwardProxy      InspectionTypeEnum = "SSL_FORWARD_PROXY"
)

var mappingInspectionTypeEnum = map[string]InspectionTypeEnum{
	"SSL_INBOUND_INSPECTION": InspectionTypeSslInboundInspection,
	"SSL_FORWARD_PROXY":      InspectionTypeSslForwardProxy,
}

var mappingInspectionTypeEnumLowerCase = map[string]InspectionTypeEnum{
	"ssl_inbound_inspection": InspectionTypeSslInboundInspection,
	"ssl_forward_proxy":      InspectionTypeSslForwardProxy,
}

// GetInspectionTypeEnumValues Enumerates the set of values for InspectionTypeEnum
func GetInspectionTypeEnumValues() []InspectionTypeEnum {
	values := make([]InspectionTypeEnum, 0)
	for _, v := range mappingInspectionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetInspectionTypeEnumStringValues Enumerates the set of values in String for InspectionTypeEnum
func GetInspectionTypeEnumStringValues() []string {
	return []string{
		"SSL_INBOUND_INSPECTION",
		"SSL_FORWARD_PROXY",
	}
}

// GetMappingInspectionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInspectionTypeEnum(val string) (InspectionTypeEnum, bool) {
	enum, ok := mappingInspectionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
