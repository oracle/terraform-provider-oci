// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

// SensitiveTypeEntityEnum Enum with underlying type: string
type SensitiveTypeEntityEnum string

// Set of constants representing the allowable values for SensitiveTypeEntityEnum
const (
	SensitiveTypeEntitySensitiveType     SensitiveTypeEntityEnum = "SENSITIVE_TYPE"
	SensitiveTypeEntitySensitiveCategory SensitiveTypeEntityEnum = "SENSITIVE_CATEGORY"
)

var mappingSensitiveTypeEntityEnum = map[string]SensitiveTypeEntityEnum{
	"SENSITIVE_TYPE":     SensitiveTypeEntitySensitiveType,
	"SENSITIVE_CATEGORY": SensitiveTypeEntitySensitiveCategory,
}

// GetSensitiveTypeEntityEnumValues Enumerates the set of values for SensitiveTypeEntityEnum
func GetSensitiveTypeEntityEnumValues() []SensitiveTypeEntityEnum {
	values := make([]SensitiveTypeEntityEnum, 0)
	for _, v := range mappingSensitiveTypeEntityEnum {
		values = append(values, v)
	}
	return values
}

// GetSensitiveTypeEntityEnumStringValues Enumerates the set of values in String for SensitiveTypeEntityEnum
func GetSensitiveTypeEntityEnumStringValues() []string {
	return []string{
		"SENSITIVE_TYPE",
		"SENSITIVE_CATEGORY",
	}
}
