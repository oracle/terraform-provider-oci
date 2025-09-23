// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Fleets API
//
// The APIs for the Fleet Management (https://docs.oracle.com/en-us/iaas/jms/doc/fleet-management.html) feature of Java Management Service to monitor and manage the usage of Java in your enterprise. Use these APIs to manage fleets, configure managed instances to report to fleets, and gain insights into the Java workloads running on these instances by carrying out basic and advanced features.
//

package jms

import (
	"strings"
)

// ConfidenceLevelEnum Enum with underlying type: string
type ConfidenceLevelEnum string

// Set of constants representing the allowable values for ConfidenceLevelEnum
const (
	ConfidenceLevelHigh   ConfidenceLevelEnum = "HIGH"
	ConfidenceLevelMedium ConfidenceLevelEnum = "MEDIUM"
	ConfidenceLevelLow    ConfidenceLevelEnum = "LOW"
)

var mappingConfidenceLevelEnum = map[string]ConfidenceLevelEnum{
	"HIGH":   ConfidenceLevelHigh,
	"MEDIUM": ConfidenceLevelMedium,
	"LOW":    ConfidenceLevelLow,
}

var mappingConfidenceLevelEnumLowerCase = map[string]ConfidenceLevelEnum{
	"high":   ConfidenceLevelHigh,
	"medium": ConfidenceLevelMedium,
	"low":    ConfidenceLevelLow,
}

// GetConfidenceLevelEnumValues Enumerates the set of values for ConfidenceLevelEnum
func GetConfidenceLevelEnumValues() []ConfidenceLevelEnum {
	values := make([]ConfidenceLevelEnum, 0)
	for _, v := range mappingConfidenceLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetConfidenceLevelEnumStringValues Enumerates the set of values in String for ConfidenceLevelEnum
func GetConfidenceLevelEnumStringValues() []string {
	return []string{
		"HIGH",
		"MEDIUM",
		"LOW",
	}
}

// GetMappingConfidenceLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConfidenceLevelEnum(val string) (ConfidenceLevelEnum, bool) {
	enum, ok := mappingConfidenceLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
