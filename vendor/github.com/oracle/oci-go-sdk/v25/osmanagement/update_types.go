// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management API
//
// API for the OS Management service. Use these API operations for working
// with Managed instances and Managed instance groups.
//

package osmanagement

// UpdateTypesEnum Enum with underlying type: string
type UpdateTypesEnum string

// Set of constants representing the allowable values for UpdateTypesEnum
const (
	UpdateTypesSecurity    UpdateTypesEnum = "SECURITY"
	UpdateTypesBug         UpdateTypesEnum = "BUG"
	UpdateTypesEnhancement UpdateTypesEnum = "ENHANCEMENT"
	UpdateTypesOther       UpdateTypesEnum = "OTHER"
)

var mappingUpdateTypes = map[string]UpdateTypesEnum{
	"SECURITY":    UpdateTypesSecurity,
	"BUG":         UpdateTypesBug,
	"ENHANCEMENT": UpdateTypesEnhancement,
	"OTHER":       UpdateTypesOther,
}

// GetUpdateTypesEnumValues Enumerates the set of values for UpdateTypesEnum
func GetUpdateTypesEnumValues() []UpdateTypesEnum {
	values := make([]UpdateTypesEnum, 0)
	for _, v := range mappingUpdateTypes {
		values = append(values, v)
	}
	return values
}
