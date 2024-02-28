// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Connector Hub API
//
// Use the Connector Hub API to transfer data between services in Oracle Cloud Infrastructure.
// For more information about Connector Hub, see
// the Connector Hub documentation (https://docs.cloud.oracle.com/iaas/Content/connector-hub/home.htm).
// Connector Hub is formerly known as Service Connector Hub.
//

package sch

import (
	"strings"
)

// EstimatedThroughputEnum Enum with underlying type: string
type EstimatedThroughputEnum string

// Set of constants representing the allowable values for EstimatedThroughputEnum
const (
	EstimatedThroughputLow     EstimatedThroughputEnum = "LOW"
	EstimatedThroughputMedium  EstimatedThroughputEnum = "MEDIUM"
	EstimatedThroughputHigh    EstimatedThroughputEnum = "HIGH"
	EstimatedThroughputUnknown EstimatedThroughputEnum = "UNKNOWN"
)

var mappingEstimatedThroughputEnum = map[string]EstimatedThroughputEnum{
	"LOW":     EstimatedThroughputLow,
	"MEDIUM":  EstimatedThroughputMedium,
	"HIGH":    EstimatedThroughputHigh,
	"UNKNOWN": EstimatedThroughputUnknown,
}

var mappingEstimatedThroughputEnumLowerCase = map[string]EstimatedThroughputEnum{
	"low":     EstimatedThroughputLow,
	"medium":  EstimatedThroughputMedium,
	"high":    EstimatedThroughputHigh,
	"unknown": EstimatedThroughputUnknown,
}

// GetEstimatedThroughputEnumValues Enumerates the set of values for EstimatedThroughputEnum
func GetEstimatedThroughputEnumValues() []EstimatedThroughputEnum {
	values := make([]EstimatedThroughputEnum, 0)
	for _, v := range mappingEstimatedThroughputEnum {
		values = append(values, v)
	}
	return values
}

// GetEstimatedThroughputEnumStringValues Enumerates the set of values in String for EstimatedThroughputEnum
func GetEstimatedThroughputEnumStringValues() []string {
	return []string{
		"LOW",
		"MEDIUM",
		"HIGH",
		"UNKNOWN",
	}
}

// GetMappingEstimatedThroughputEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEstimatedThroughputEnum(val string) (EstimatedThroughputEnum, bool) {
	enum, ok := mappingEstimatedThroughputEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
