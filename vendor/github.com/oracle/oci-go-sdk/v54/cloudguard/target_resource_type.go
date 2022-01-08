// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard API
//
// Use the Cloud Guard API to automate processes that you would otherwise perform through the Cloud Guard Console.
// **Note:** You can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

// TargetResourceTypeEnum Enum with underlying type: string
type TargetResourceTypeEnum string

// Set of constants representing the allowable values for TargetResourceTypeEnum
const (
	TargetResourceTypeCompartment  TargetResourceTypeEnum = "COMPARTMENT"
	TargetResourceTypeErpcloud     TargetResourceTypeEnum = "ERPCLOUD"
	TargetResourceTypeHcmcloud     TargetResourceTypeEnum = "HCMCLOUD"
	TargetResourceTypeSecurityZone TargetResourceTypeEnum = "SECURITY_ZONE"
)

var mappingTargetResourceTypeEnum = map[string]TargetResourceTypeEnum{
	"COMPARTMENT":   TargetResourceTypeCompartment,
	"ERPCLOUD":      TargetResourceTypeErpcloud,
	"HCMCLOUD":      TargetResourceTypeHcmcloud,
	"SECURITY_ZONE": TargetResourceTypeSecurityZone,
}

// GetTargetResourceTypeEnumValues Enumerates the set of values for TargetResourceTypeEnum
func GetTargetResourceTypeEnumValues() []TargetResourceTypeEnum {
	values := make([]TargetResourceTypeEnum, 0)
	for _, v := range mappingTargetResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTargetResourceTypeEnumStringValues Enumerates the set of values in String for TargetResourceTypeEnum
func GetTargetResourceTypeEnumStringValues() []string {
	return []string{
		"COMPARTMENT",
		"ERPCLOUD",
		"HCMCLOUD",
		"SECURITY_ZONE",
	}
}
