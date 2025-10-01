// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Exadata Fleet Update service API
//
// Use the Exadata Fleet Update service to patch large collections of components directly,
// as a single entity, orchestrating the maintenance actions to update all chosen components in the stack in a single cycle.
//

package fleetsoftwareupdate

import (
	"strings"
)

// ExadataReleaseVersionsEnum Enum with underlying type: string
type ExadataReleaseVersionsEnum string

// Set of constants representing the allowable values for ExadataReleaseVersionsEnum
const (
	ExadataReleaseVersionsExaOl5111 ExadataReleaseVersionsEnum = "EXA_OL_5_11_1"
	ExadataReleaseVersionsExaOl5112 ExadataReleaseVersionsEnum = "EXA_OL_5_11_2"
	ExadataReleaseVersionsExaOl5121 ExadataReleaseVersionsEnum = "EXA_OL_5_12_1"
	ExadataReleaseVersionsExaOl6122 ExadataReleaseVersionsEnum = "EXA_OL_6_12_2"
	ExadataReleaseVersionsExaOl6181 ExadataReleaseVersionsEnum = "EXA_OL_6_18_1"
	ExadataReleaseVersionsExaOl7191 ExadataReleaseVersionsEnum = "EXA_OL_7_19_1"
	ExadataReleaseVersionsExaOl7192 ExadataReleaseVersionsEnum = "EXA_OL_7_19_2"
	ExadataReleaseVersionsExaOl7193 ExadataReleaseVersionsEnum = "EXA_OL_7_19_3"
	ExadataReleaseVersionsExaOl7201 ExadataReleaseVersionsEnum = "EXA_OL_7_20_1"
	ExadataReleaseVersionsExaOl7212 ExadataReleaseVersionsEnum = "EXA_OL_7_21_2"
	ExadataReleaseVersionsExaOl7221 ExadataReleaseVersionsEnum = "EXA_OL_7_22_1"
	ExadataReleaseVersionsExaOl8231 ExadataReleaseVersionsEnum = "EXA_OL_8_23_1"
	ExadataReleaseVersionsExaOl8241 ExadataReleaseVersionsEnum = "EXA_OL_8_24_1"
)

var mappingExadataReleaseVersionsEnum = map[string]ExadataReleaseVersionsEnum{
	"EXA_OL_5_11_1": ExadataReleaseVersionsExaOl5111,
	"EXA_OL_5_11_2": ExadataReleaseVersionsExaOl5112,
	"EXA_OL_5_12_1": ExadataReleaseVersionsExaOl5121,
	"EXA_OL_6_12_2": ExadataReleaseVersionsExaOl6122,
	"EXA_OL_6_18_1": ExadataReleaseVersionsExaOl6181,
	"EXA_OL_7_19_1": ExadataReleaseVersionsExaOl7191,
	"EXA_OL_7_19_2": ExadataReleaseVersionsExaOl7192,
	"EXA_OL_7_19_3": ExadataReleaseVersionsExaOl7193,
	"EXA_OL_7_20_1": ExadataReleaseVersionsExaOl7201,
	"EXA_OL_7_21_2": ExadataReleaseVersionsExaOl7212,
	"EXA_OL_7_22_1": ExadataReleaseVersionsExaOl7221,
	"EXA_OL_8_23_1": ExadataReleaseVersionsExaOl8231,
	"EXA_OL_8_24_1": ExadataReleaseVersionsExaOl8241,
}

var mappingExadataReleaseVersionsEnumLowerCase = map[string]ExadataReleaseVersionsEnum{
	"exa_ol_5_11_1": ExadataReleaseVersionsExaOl5111,
	"exa_ol_5_11_2": ExadataReleaseVersionsExaOl5112,
	"exa_ol_5_12_1": ExadataReleaseVersionsExaOl5121,
	"exa_ol_6_12_2": ExadataReleaseVersionsExaOl6122,
	"exa_ol_6_18_1": ExadataReleaseVersionsExaOl6181,
	"exa_ol_7_19_1": ExadataReleaseVersionsExaOl7191,
	"exa_ol_7_19_2": ExadataReleaseVersionsExaOl7192,
	"exa_ol_7_19_3": ExadataReleaseVersionsExaOl7193,
	"exa_ol_7_20_1": ExadataReleaseVersionsExaOl7201,
	"exa_ol_7_21_2": ExadataReleaseVersionsExaOl7212,
	"exa_ol_7_22_1": ExadataReleaseVersionsExaOl7221,
	"exa_ol_8_23_1": ExadataReleaseVersionsExaOl8231,
	"exa_ol_8_24_1": ExadataReleaseVersionsExaOl8241,
}

// GetExadataReleaseVersionsEnumValues Enumerates the set of values for ExadataReleaseVersionsEnum
func GetExadataReleaseVersionsEnumValues() []ExadataReleaseVersionsEnum {
	values := make([]ExadataReleaseVersionsEnum, 0)
	for _, v := range mappingExadataReleaseVersionsEnum {
		values = append(values, v)
	}
	return values
}

// GetExadataReleaseVersionsEnumStringValues Enumerates the set of values in String for ExadataReleaseVersionsEnum
func GetExadataReleaseVersionsEnumStringValues() []string {
	return []string{
		"EXA_OL_5_11_1",
		"EXA_OL_5_11_2",
		"EXA_OL_5_12_1",
		"EXA_OL_6_12_2",
		"EXA_OL_6_18_1",
		"EXA_OL_7_19_1",
		"EXA_OL_7_19_2",
		"EXA_OL_7_19_3",
		"EXA_OL_7_20_1",
		"EXA_OL_7_21_2",
		"EXA_OL_7_22_1",
		"EXA_OL_8_23_1",
		"EXA_OL_8_24_1",
	}
}

// GetMappingExadataReleaseVersionsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadataReleaseVersionsEnum(val string) (ExadataReleaseVersionsEnum, bool) {
	enum, ok := mappingExadataReleaseVersionsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
