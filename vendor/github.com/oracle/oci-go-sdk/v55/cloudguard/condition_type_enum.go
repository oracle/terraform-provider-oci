// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard API
//
// Use the Cloud Guard API to automate processes that you would otherwise perform through the Cloud Guard Console.
// **Note:** You can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

// ConditionTypeEnumEnum Enum with underlying type: string
type ConditionTypeEnumEnum string

// Set of constants representing the allowable values for ConditionTypeEnumEnum
const (
	ConditionTypeEnumActivityCondition   ConditionTypeEnumEnum = "ActivityCondition"
	ConditionTypeEnumSecurityCondition   ConditionTypeEnumEnum = "SecurityCondition"
	ConditionTypeEnumCloudGuardCondition ConditionTypeEnumEnum = "CloudGuardCondition"
	ConditionTypeEnumThreatCondition     ConditionTypeEnumEnum = "ThreatCondition"
)

var mappingConditionTypeEnumEnum = map[string]ConditionTypeEnumEnum{
	"ActivityCondition":   ConditionTypeEnumActivityCondition,
	"SecurityCondition":   ConditionTypeEnumSecurityCondition,
	"CloudGuardCondition": ConditionTypeEnumCloudGuardCondition,
	"ThreatCondition":     ConditionTypeEnumThreatCondition,
}

// GetConditionTypeEnumEnumValues Enumerates the set of values for ConditionTypeEnumEnum
func GetConditionTypeEnumEnumValues() []ConditionTypeEnumEnum {
	values := make([]ConditionTypeEnumEnum, 0)
	for _, v := range mappingConditionTypeEnumEnum {
		values = append(values, v)
	}
	return values
}

// GetConditionTypeEnumEnumStringValues Enumerates the set of values in String for ConditionTypeEnumEnum
func GetConditionTypeEnumEnumStringValues() []string {
	return []string{
		"ActivityCondition",
		"SecurityCondition",
		"CloudGuardCondition",
		"ThreatCondition",
	}
}
