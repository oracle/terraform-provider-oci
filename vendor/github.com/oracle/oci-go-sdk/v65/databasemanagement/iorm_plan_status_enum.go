// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"strings"
)

// IormPlanStatusEnumEnum Enum with underlying type: string
type IormPlanStatusEnumEnum string

// Set of constants representing the allowable values for IormPlanStatusEnumEnum
const (
	IormPlanStatusEnumActive   IormPlanStatusEnumEnum = "ACTIVE"
	IormPlanStatusEnumInactive IormPlanStatusEnumEnum = "INACTIVE"
	IormPlanStatusEnumOther    IormPlanStatusEnumEnum = "OTHER"
)

var mappingIormPlanStatusEnumEnum = map[string]IormPlanStatusEnumEnum{
	"ACTIVE":   IormPlanStatusEnumActive,
	"INACTIVE": IormPlanStatusEnumInactive,
	"OTHER":    IormPlanStatusEnumOther,
}

var mappingIormPlanStatusEnumEnumLowerCase = map[string]IormPlanStatusEnumEnum{
	"active":   IormPlanStatusEnumActive,
	"inactive": IormPlanStatusEnumInactive,
	"other":    IormPlanStatusEnumOther,
}

// GetIormPlanStatusEnumEnumValues Enumerates the set of values for IormPlanStatusEnumEnum
func GetIormPlanStatusEnumEnumValues() []IormPlanStatusEnumEnum {
	values := make([]IormPlanStatusEnumEnum, 0)
	for _, v := range mappingIormPlanStatusEnumEnum {
		values = append(values, v)
	}
	return values
}

// GetIormPlanStatusEnumEnumStringValues Enumerates the set of values in String for IormPlanStatusEnumEnum
func GetIormPlanStatusEnumEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"OTHER",
	}
}

// GetMappingIormPlanStatusEnumEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIormPlanStatusEnumEnum(val string) (IormPlanStatusEnumEnum, bool) {
	enum, ok := mappingIormPlanStatusEnumEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
