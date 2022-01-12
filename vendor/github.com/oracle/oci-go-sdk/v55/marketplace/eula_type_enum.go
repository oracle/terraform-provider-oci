// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Manage applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplace

// EulaTypeEnumEnum Enum with underlying type: string
type EulaTypeEnumEnum string

// Set of constants representing the allowable values for EulaTypeEnumEnum
const (
	EulaTypeEnumText EulaTypeEnumEnum = "TEXT"
)

var mappingEulaTypeEnum = map[string]EulaTypeEnumEnum{
	"TEXT": EulaTypeEnumText,
}

// GetEulaTypeEnumEnumValues Enumerates the set of values for EulaTypeEnumEnum
func GetEulaTypeEnumEnumValues() []EulaTypeEnumEnum {
	values := make([]EulaTypeEnumEnum, 0)
	for _, v := range mappingEulaTypeEnum {
		values = append(values, v)
	}
	return values
}
