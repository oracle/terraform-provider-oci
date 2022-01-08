// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard API
//
// Use the Cloud Guard API to automate processes that you would otherwise perform through the Cloud Guard Console.
// **Note:** You can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

// ConditionFilterTypeEnum Enum with underlying type: string
type ConditionFilterTypeEnum string

// Set of constants representing the allowable values for ConditionFilterTypeEnum
const (
	ConditionFilterTypeCondition ConditionFilterTypeEnum = "CONDITION"
	ConditionFilterTypeConfig    ConditionFilterTypeEnum = "CONFIG"
)

var mappingConditionFilterTypeEnum = map[string]ConditionFilterTypeEnum{
	"CONDITION": ConditionFilterTypeCondition,
	"CONFIG":    ConditionFilterTypeConfig,
}

// GetConditionFilterTypeEnumValues Enumerates the set of values for ConditionFilterTypeEnum
func GetConditionFilterTypeEnumValues() []ConditionFilterTypeEnum {
	values := make([]ConditionFilterTypeEnum, 0)
	for _, v := range mappingConditionFilterTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetConditionFilterTypeEnumStringValues Enumerates the set of values in String for ConditionFilterTypeEnum
func GetConditionFilterTypeEnumStringValues() []string {
	return []string{
		"CONDITION",
		"CONFIG",
	}
}
