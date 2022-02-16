// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"strings"
)

// AlertStatusEnum Enum with underlying type: string
type AlertStatusEnum string

// Set of constants representing the allowable values for AlertStatusEnum
const (
	AlertStatusOpen   AlertStatusEnum = "OPEN"
	AlertStatusClosed AlertStatusEnum = "CLOSED"
)

var mappingAlertStatusEnum = map[string]AlertStatusEnum{
	"OPEN":   AlertStatusOpen,
	"CLOSED": AlertStatusClosed,
}

// GetAlertStatusEnumValues Enumerates the set of values for AlertStatusEnum
func GetAlertStatusEnumValues() []AlertStatusEnum {
	values := make([]AlertStatusEnum, 0)
	for _, v := range mappingAlertStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetAlertStatusEnumStringValues Enumerates the set of values in String for AlertStatusEnum
func GetAlertStatusEnumStringValues() []string {
	return []string{
		"OPEN",
		"CLOSED",
	}
}

// GetMappingAlertStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAlertStatusEnum(val string) (AlertStatusEnum, bool) {
	mappingAlertStatusEnumIgnoreCase := make(map[string]AlertStatusEnum)
	for k, v := range mappingAlertStatusEnum {
		mappingAlertStatusEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingAlertStatusEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
