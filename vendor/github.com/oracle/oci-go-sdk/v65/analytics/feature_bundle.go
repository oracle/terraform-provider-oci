// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

import (
	"strings"
)

// FeatureBundleEnum Enum with underlying type: string
type FeatureBundleEnum string

// Set of constants representing the allowable values for FeatureBundleEnum
const (
	FeatureBundleFawPaid    FeatureBundleEnum = "FAW_PAID"
	FeatureBundleFawFree    FeatureBundleEnum = "FAW_FREE"
	FeatureBundleEeEmbedded FeatureBundleEnum = "EE_EMBEDDED"
	FeatureBundleSeEmbedded FeatureBundleEnum = "SE_EMBEDDED"
)

var mappingFeatureBundleEnum = map[string]FeatureBundleEnum{
	"FAW_PAID":    FeatureBundleFawPaid,
	"FAW_FREE":    FeatureBundleFawFree,
	"EE_EMBEDDED": FeatureBundleEeEmbedded,
	"SE_EMBEDDED": FeatureBundleSeEmbedded,
}

var mappingFeatureBundleEnumLowerCase = map[string]FeatureBundleEnum{
	"faw_paid":    FeatureBundleFawPaid,
	"faw_free":    FeatureBundleFawFree,
	"ee_embedded": FeatureBundleEeEmbedded,
	"se_embedded": FeatureBundleSeEmbedded,
}

// GetFeatureBundleEnumValues Enumerates the set of values for FeatureBundleEnum
func GetFeatureBundleEnumValues() []FeatureBundleEnum {
	values := make([]FeatureBundleEnum, 0)
	for _, v := range mappingFeatureBundleEnum {
		values = append(values, v)
	}
	return values
}

// GetFeatureBundleEnumStringValues Enumerates the set of values in String for FeatureBundleEnum
func GetFeatureBundleEnumStringValues() []string {
	return []string{
		"FAW_PAID",
		"FAW_FREE",
		"EE_EMBEDDED",
		"SE_EMBEDDED",
	}
}

// GetMappingFeatureBundleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFeatureBundleEnum(val string) (FeatureBundleEnum, bool) {
	enum, ok := mappingFeatureBundleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
