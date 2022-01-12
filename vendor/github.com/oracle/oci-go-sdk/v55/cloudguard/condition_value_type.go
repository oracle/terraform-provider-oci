// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

// ConditionValueTypeEnum Enum with underlying type: string
type ConditionValueTypeEnum string

// Set of constants representing the allowable values for ConditionValueTypeEnum
const (
	ConditionValueTypeManaged ConditionValueTypeEnum = "MANAGED"
	ConditionValueTypeCustom  ConditionValueTypeEnum = "CUSTOM"
)

var mappingConditionValueType = map[string]ConditionValueTypeEnum{
	"MANAGED": ConditionValueTypeManaged,
	"CUSTOM":  ConditionValueTypeCustom,
}

// GetConditionValueTypeEnumValues Enumerates the set of values for ConditionValueTypeEnum
func GetConditionValueTypeEnumValues() []ConditionValueTypeEnum {
	values := make([]ConditionValueTypeEnum, 0)
	for _, v := range mappingConditionValueType {
		values = append(values, v)
	}
	return values
}
