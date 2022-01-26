// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

// TargetResourceTypeEnum Enum with underlying type: string
type TargetResourceTypeEnum string

// Set of constants representing the allowable values for TargetResourceTypeEnum
const (
	TargetResourceTypeCompartment TargetResourceTypeEnum = "COMPARTMENT"
	TargetResourceTypeErpcloud    TargetResourceTypeEnum = "ERPCLOUD"
	TargetResourceTypeHcmcloud    TargetResourceTypeEnum = "HCMCLOUD"
)

var mappingTargetResourceType = map[string]TargetResourceTypeEnum{
	"COMPARTMENT": TargetResourceTypeCompartment,
	"ERPCLOUD":    TargetResourceTypeErpcloud,
	"HCMCLOUD":    TargetResourceTypeHcmcloud,
}

// GetTargetResourceTypeEnumValues Enumerates the set of values for TargetResourceTypeEnum
func GetTargetResourceTypeEnumValues() []TargetResourceTypeEnum {
	values := make([]TargetResourceTypeEnum, 0)
	for _, v := range mappingTargetResourceType {
		values = append(values, v)
	}
	return values
}
