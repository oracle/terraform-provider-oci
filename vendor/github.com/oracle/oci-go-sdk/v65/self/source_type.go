// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// SELF Service API
//
// Use the SELF Service API to manage Subscriptions in Oracle Cloud Infrastructure Marketplace. For more information, see Overview of Marketplace (https://docs.oracle.com/iaas/Content/Marketplace/Concepts/marketoverview.htm)
//

package self

import (
	"strings"
)

// SourceTypeEnum Enum with underlying type: string
type SourceTypeEnum string

// Set of constants representing the allowable values for SourceTypeEnum
const (
	SourceTypeOciNative  SourceTypeEnum = "OCI_NATIVE"
	SourceTypeThirdParty SourceTypeEnum = "THIRD_PARTY"
)

var mappingSourceTypeEnum = map[string]SourceTypeEnum{
	"OCI_NATIVE":  SourceTypeOciNative,
	"THIRD_PARTY": SourceTypeThirdParty,
}

var mappingSourceTypeEnumLowerCase = map[string]SourceTypeEnum{
	"oci_native":  SourceTypeOciNative,
	"third_party": SourceTypeThirdParty,
}

// GetSourceTypeEnumValues Enumerates the set of values for SourceTypeEnum
func GetSourceTypeEnumValues() []SourceTypeEnum {
	values := make([]SourceTypeEnum, 0)
	for _, v := range mappingSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSourceTypeEnumStringValues Enumerates the set of values in String for SourceTypeEnum
func GetSourceTypeEnumStringValues() []string {
	return []string{
		"OCI_NATIVE",
		"THIRD_PARTY",
	}
}

// GetMappingSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSourceTypeEnum(val string) (SourceTypeEnum, bool) {
	enum, ok := mappingSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
