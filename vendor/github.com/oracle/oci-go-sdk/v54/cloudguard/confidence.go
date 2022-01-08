// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard API
//
// Use the Cloud Guard API to automate processes that you would otherwise perform through the Cloud Guard Console.
// **Note:** You can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

// ConfidenceEnum Enum with underlying type: string
type ConfidenceEnum string

// Set of constants representing the allowable values for ConfidenceEnum
const (
	ConfidenceCritical ConfidenceEnum = "CRITICAL"
	ConfidenceHigh     ConfidenceEnum = "HIGH"
	ConfidenceMedium   ConfidenceEnum = "MEDIUM"
	ConfidenceLow      ConfidenceEnum = "LOW"
	ConfidenceMinor    ConfidenceEnum = "MINOR"
)

var mappingConfidenceEnum = map[string]ConfidenceEnum{
	"CRITICAL": ConfidenceCritical,
	"HIGH":     ConfidenceHigh,
	"MEDIUM":   ConfidenceMedium,
	"LOW":      ConfidenceLow,
	"MINOR":    ConfidenceMinor,
}

// GetConfidenceEnumValues Enumerates the set of values for ConfidenceEnum
func GetConfidenceEnumValues() []ConfidenceEnum {
	values := make([]ConfidenceEnum, 0)
	for _, v := range mappingConfidenceEnum {
		values = append(values, v)
	}
	return values
}

// GetConfidenceEnumStringValues Enumerates the set of values in String for ConfidenceEnum
func GetConfidenceEnumStringValues() []string {
	return []string{
		"CRITICAL",
		"HIGH",
		"MEDIUM",
		"LOW",
		"MINOR",
	}
}
