// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management API
//
// API for the OS Management service. Use these API operations for working
// with Managed instances and Managed instance groups.
//

package osmanagement

import (
	"strings"
)

// UpdateTypesEnum Enum with underlying type: string
type UpdateTypesEnum string

// Set of constants representing the allowable values for UpdateTypesEnum
const (
	UpdateTypesSecurity    UpdateTypesEnum = "SECURITY"
	UpdateTypesBug         UpdateTypesEnum = "BUG"
	UpdateTypesEnhancement UpdateTypesEnum = "ENHANCEMENT"
	UpdateTypesOther       UpdateTypesEnum = "OTHER"
)

var mappingUpdateTypesEnum = map[string]UpdateTypesEnum{
	"SECURITY":    UpdateTypesSecurity,
	"BUG":         UpdateTypesBug,
	"ENHANCEMENT": UpdateTypesEnhancement,
	"OTHER":       UpdateTypesOther,
}

// GetUpdateTypesEnumValues Enumerates the set of values for UpdateTypesEnum
func GetUpdateTypesEnumValues() []UpdateTypesEnum {
	values := make([]UpdateTypesEnum, 0)
	for _, v := range mappingUpdateTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateTypesEnumStringValues Enumerates the set of values in String for UpdateTypesEnum
func GetUpdateTypesEnumStringValues() []string {
	return []string{
		"SECURITY",
		"BUG",
		"ENHANCEMENT",
		"OTHER",
	}
}

// GetMappingUpdateTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateTypesEnum(val string) (UpdateTypesEnum, bool) {
	mappingUpdateTypesEnumIgnoreCase := make(map[string]UpdateTypesEnum)
	for k, v := range mappingUpdateTypesEnum {
		mappingUpdateTypesEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingUpdateTypesEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
