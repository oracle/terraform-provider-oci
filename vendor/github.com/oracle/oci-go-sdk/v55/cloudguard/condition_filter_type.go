// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

// ConditionFilterTypeEnum Enum with underlying type: string
type ConditionFilterTypeEnum string

// Set of constants representing the allowable values for ConditionFilterTypeEnum
const (
	ConditionFilterTypeCondition ConditionFilterTypeEnum = "CONDITION"
	ConditionFilterTypeConfig    ConditionFilterTypeEnum = "CONFIG"
)

var mappingConditionFilterType = map[string]ConditionFilterTypeEnum{
	"CONDITION": ConditionFilterTypeCondition,
	"CONFIG":    ConditionFilterTypeConfig,
}

// GetConditionFilterTypeEnumValues Enumerates the set of values for ConditionFilterTypeEnum
func GetConditionFilterTypeEnumValues() []ConditionFilterTypeEnum {
	values := make([]ConditionFilterTypeEnum, 0)
	for _, v := range mappingConditionFilterType {
		values = append(values, v)
	}
	return values
}
