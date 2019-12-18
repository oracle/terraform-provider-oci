// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// OSMS
//
// OS Management as a Service API definition
//

package osmanagement

// UpdateTypesEnum Enum with underlying type: string
type UpdateTypesEnum string

// Set of constants representing the allowable values for UpdateTypesEnum
const (
	UpdateTypesSecurity    UpdateTypesEnum = "SECURITY"
	UpdateTypesBug         UpdateTypesEnum = "BUG"
	UpdateTypesEnhancement UpdateTypesEnum = "ENHANCEMENT"
)

var mappingUpdateTypes = map[string]UpdateTypesEnum{
	"SECURITY":    UpdateTypesSecurity,
	"BUG":         UpdateTypesBug,
	"ENHANCEMENT": UpdateTypesEnhancement,
}

// GetUpdateTypesEnumValues Enumerates the set of values for UpdateTypesEnum
func GetUpdateTypesEnumValues() []UpdateTypesEnum {
	values := make([]UpdateTypesEnum, 0)
	for _, v := range mappingUpdateTypes {
		values = append(values, v)
	}
	return values
}
