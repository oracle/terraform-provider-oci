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

// InspectActionTypeEnum Enum with underlying type: string
type InspectActionTypeEnum string

// Set of constants representing the allowable values for InspectActionTypeEnum
const (
	InspectActionTypeInspect              InspectActionTypeEnum = "INSPECT"
	InspectActionTypeInspectAndCaptureLog InspectActionTypeEnum = "INSPECT_AND_CAPTURE_LOG"
)

var mappingInspectActionTypeEnum = map[string]InspectActionTypeEnum{
	"INSPECT":                 InspectActionTypeInspect,
	"INSPECT_AND_CAPTURE_LOG": InspectActionTypeInspectAndCaptureLog,
}

var mappingInspectActionTypeEnumLowerCase = map[string]InspectActionTypeEnum{
	"inspect":                 InspectActionTypeInspect,
	"inspect_and_capture_log": InspectActionTypeInspectAndCaptureLog,
}

// GetInspectActionTypeEnumValues Enumerates the set of values for InspectActionTypeEnum
func GetInspectActionTypeEnumValues() []InspectActionTypeEnum {
	values := make([]InspectActionTypeEnum, 0)
	for _, v := range mappingInspectActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetInspectActionTypeEnumStringValues Enumerates the set of values in String for InspectActionTypeEnum
func GetInspectActionTypeEnumStringValues() []string {
	return []string{
		"INSPECT",
		"INSPECT_AND_CAPTURE_LOG",
	}
}

// GetMappingInspectActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInspectActionTypeEnum(val string) (InspectActionTypeEnum, bool) {
	enum, ok := mappingInspectActionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
