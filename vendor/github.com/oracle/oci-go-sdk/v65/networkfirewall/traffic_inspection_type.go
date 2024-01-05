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

// TrafficInspectionTypeEnum Enum with underlying type: string
type TrafficInspectionTypeEnum string

// Set of constants representing the allowable values for TrafficInspectionTypeEnum
const (
	TrafficInspectionTypeIntrusionDetection  TrafficInspectionTypeEnum = "INTRUSION_DETECTION"
	TrafficInspectionTypeIntrusionPrevention TrafficInspectionTypeEnum = "INTRUSION_PREVENTION"
)

var mappingTrafficInspectionTypeEnum = map[string]TrafficInspectionTypeEnum{
	"INTRUSION_DETECTION":  TrafficInspectionTypeIntrusionDetection,
	"INTRUSION_PREVENTION": TrafficInspectionTypeIntrusionPrevention,
}

var mappingTrafficInspectionTypeEnumLowerCase = map[string]TrafficInspectionTypeEnum{
	"intrusion_detection":  TrafficInspectionTypeIntrusionDetection,
	"intrusion_prevention": TrafficInspectionTypeIntrusionPrevention,
}

// GetTrafficInspectionTypeEnumValues Enumerates the set of values for TrafficInspectionTypeEnum
func GetTrafficInspectionTypeEnumValues() []TrafficInspectionTypeEnum {
	values := make([]TrafficInspectionTypeEnum, 0)
	for _, v := range mappingTrafficInspectionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTrafficInspectionTypeEnumStringValues Enumerates the set of values in String for TrafficInspectionTypeEnum
func GetTrafficInspectionTypeEnumStringValues() []string {
	return []string{
		"INTRUSION_DETECTION",
		"INTRUSION_PREVENTION",
	}
}

// GetMappingTrafficInspectionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTrafficInspectionTypeEnum(val string) (TrafficInspectionTypeEnum, bool) {
	enum, ok := mappingTrafficInspectionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
