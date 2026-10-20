// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"strings"
)

// SubsettingPolicyFormatEnum Enum with underlying type: string
type SubsettingPolicyFormatEnum string

// Set of constants representing the allowable values for SubsettingPolicyFormatEnum
const (
	SubsettingPolicyFormatXml SubsettingPolicyFormatEnum = "XML"
)

var mappingSubsettingPolicyFormatEnum = map[string]SubsettingPolicyFormatEnum{
	"XML": SubsettingPolicyFormatXml,
}

var mappingSubsettingPolicyFormatEnumLowerCase = map[string]SubsettingPolicyFormatEnum{
	"xml": SubsettingPolicyFormatXml,
}

// GetSubsettingPolicyFormatEnumValues Enumerates the set of values for SubsettingPolicyFormatEnum
func GetSubsettingPolicyFormatEnumValues() []SubsettingPolicyFormatEnum {
	values := make([]SubsettingPolicyFormatEnum, 0)
	for _, v := range mappingSubsettingPolicyFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetSubsettingPolicyFormatEnumStringValues Enumerates the set of values in String for SubsettingPolicyFormatEnum
func GetSubsettingPolicyFormatEnumStringValues() []string {
	return []string{
		"XML",
	}
}

// GetMappingSubsettingPolicyFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSubsettingPolicyFormatEnum(val string) (SubsettingPolicyFormatEnum, bool) {
	enum, ok := mappingSubsettingPolicyFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
