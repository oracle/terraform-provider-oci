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

// AlertTypeEnumEnum Enum with underlying type: string
type AlertTypeEnumEnum string

// Set of constants representing the allowable values for AlertTypeEnumEnum
const (
	AlertTypeEnumStateful  AlertTypeEnumEnum = "STATEFUL"
	AlertTypeEnumStateless AlertTypeEnumEnum = "STATELESS"
)

var mappingAlertTypeEnumEnum = map[string]AlertTypeEnumEnum{
	"STATEFUL":  AlertTypeEnumStateful,
	"STATELESS": AlertTypeEnumStateless,
}

var mappingAlertTypeEnumEnumLowerCase = map[string]AlertTypeEnumEnum{
	"stateful":  AlertTypeEnumStateful,
	"stateless": AlertTypeEnumStateless,
}

// GetAlertTypeEnumEnumValues Enumerates the set of values for AlertTypeEnumEnum
func GetAlertTypeEnumEnumValues() []AlertTypeEnumEnum {
	values := make([]AlertTypeEnumEnum, 0)
	for _, v := range mappingAlertTypeEnumEnum {
		values = append(values, v)
	}
	return values
}

// GetAlertTypeEnumEnumStringValues Enumerates the set of values in String for AlertTypeEnumEnum
func GetAlertTypeEnumEnumStringValues() []string {
	return []string{
		"STATEFUL",
		"STATELESS",
	}
}

// GetMappingAlertTypeEnumEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAlertTypeEnumEnum(val string) (AlertTypeEnumEnum, bool) {
	enum, ok := mappingAlertTypeEnumEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
