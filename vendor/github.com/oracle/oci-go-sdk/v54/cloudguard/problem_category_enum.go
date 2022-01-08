// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard API
//
// Use the Cloud Guard API to automate processes that you would otherwise perform through the Cloud Guard Console.
// **Note:** You can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

// ProblemCategoryEnumEnum Enum with underlying type: string
type ProblemCategoryEnumEnum string

// Set of constants representing the allowable values for ProblemCategoryEnumEnum
const (
	ProblemCategoryEnumSecurityZone ProblemCategoryEnumEnum = "SECURITY_ZONE"
)

var mappingProblemCategoryEnumEnum = map[string]ProblemCategoryEnumEnum{
	"SECURITY_ZONE": ProblemCategoryEnumSecurityZone,
}

// GetProblemCategoryEnumEnumValues Enumerates the set of values for ProblemCategoryEnumEnum
func GetProblemCategoryEnumEnumValues() []ProblemCategoryEnumEnum {
	values := make([]ProblemCategoryEnumEnum, 0)
	for _, v := range mappingProblemCategoryEnumEnum {
		values = append(values, v)
	}
	return values
}

// GetProblemCategoryEnumEnumStringValues Enumerates the set of values in String for ProblemCategoryEnumEnum
func GetProblemCategoryEnumEnumStringValues() []string {
	return []string{
		"SECURITY_ZONE",
	}
}
