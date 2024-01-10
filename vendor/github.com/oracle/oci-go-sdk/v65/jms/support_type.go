// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"strings"
)

// SupportTypeEnum Enum with underlying type: string
type SupportTypeEnum string

// Set of constants representing the allowable values for SupportTypeEnum
const (
	SupportTypeLts    SupportTypeEnum = "LTS"
	SupportTypeNonLts SupportTypeEnum = "NON_LTS"
)

var mappingSupportTypeEnum = map[string]SupportTypeEnum{
	"LTS":     SupportTypeLts,
	"NON_LTS": SupportTypeNonLts,
}

var mappingSupportTypeEnumLowerCase = map[string]SupportTypeEnum{
	"lts":     SupportTypeLts,
	"non_lts": SupportTypeNonLts,
}

// GetSupportTypeEnumValues Enumerates the set of values for SupportTypeEnum
func GetSupportTypeEnumValues() []SupportTypeEnum {
	values := make([]SupportTypeEnum, 0)
	for _, v := range mappingSupportTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSupportTypeEnumStringValues Enumerates the set of values in String for SupportTypeEnum
func GetSupportTypeEnumStringValues() []string {
	return []string{
		"LTS",
		"NON_LTS",
	}
}

// GetMappingSupportTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSupportTypeEnum(val string) (SupportTypeEnum, bool) {
	enum, ok := mappingSupportTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
