// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.cloud.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.cloud.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"strings"
)

// SecurityRatingEnum Enum with underlying type: string
type SecurityRatingEnum string

// Set of constants representing the allowable values for SecurityRatingEnum
const (
	SecurityRatingExcellent SecurityRatingEnum = "EXCELLENT"
	SecurityRatingGood      SecurityRatingEnum = "GOOD"
	SecurityRatingFair      SecurityRatingEnum = "FAIR"
	SecurityRatingPoor      SecurityRatingEnum = "POOR"
	SecurityRatingNa        SecurityRatingEnum = "NA"
)

var mappingSecurityRatingEnum = map[string]SecurityRatingEnum{
	"EXCELLENT": SecurityRatingExcellent,
	"GOOD":      SecurityRatingGood,
	"FAIR":      SecurityRatingFair,
	"POOR":      SecurityRatingPoor,
	"NA":        SecurityRatingNa,
}

var mappingSecurityRatingEnumLowerCase = map[string]SecurityRatingEnum{
	"excellent": SecurityRatingExcellent,
	"good":      SecurityRatingGood,
	"fair":      SecurityRatingFair,
	"poor":      SecurityRatingPoor,
	"na":        SecurityRatingNa,
}

// GetSecurityRatingEnumValues Enumerates the set of values for SecurityRatingEnum
func GetSecurityRatingEnumValues() []SecurityRatingEnum {
	values := make([]SecurityRatingEnum, 0)
	for _, v := range mappingSecurityRatingEnum {
		values = append(values, v)
	}
	return values
}

// GetSecurityRatingEnumStringValues Enumerates the set of values in String for SecurityRatingEnum
func GetSecurityRatingEnumStringValues() []string {
	return []string{
		"EXCELLENT",
		"GOOD",
		"FAIR",
		"POOR",
		"NA",
	}
}

// GetMappingSecurityRatingEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecurityRatingEnum(val string) (SecurityRatingEnum, bool) {
	enum, ok := mappingSecurityRatingEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
