// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"strings"
)

// HandleGrantErrorsEnum Enum with underlying type: string
type HandleGrantErrorsEnum string

// Set of constants representing the allowable values for HandleGrantErrorsEnum
const (
	HandleGrantErrorsAbort       HandleGrantErrorsEnum = "ABORT"
	HandleGrantErrorsDropAccount HandleGrantErrorsEnum = "DROP_ACCOUNT"
	HandleGrantErrorsIgnore      HandleGrantErrorsEnum = "IGNORE"
)

var mappingHandleGrantErrorsEnum = map[string]HandleGrantErrorsEnum{
	"ABORT":        HandleGrantErrorsAbort,
	"DROP_ACCOUNT": HandleGrantErrorsDropAccount,
	"IGNORE":       HandleGrantErrorsIgnore,
}

var mappingHandleGrantErrorsEnumLowerCase = map[string]HandleGrantErrorsEnum{
	"abort":        HandleGrantErrorsAbort,
	"drop_account": HandleGrantErrorsDropAccount,
	"ignore":       HandleGrantErrorsIgnore,
}

// GetHandleGrantErrorsEnumValues Enumerates the set of values for HandleGrantErrorsEnum
func GetHandleGrantErrorsEnumValues() []HandleGrantErrorsEnum {
	values := make([]HandleGrantErrorsEnum, 0)
	for _, v := range mappingHandleGrantErrorsEnum {
		values = append(values, v)
	}
	return values
}

// GetHandleGrantErrorsEnumStringValues Enumerates the set of values in String for HandleGrantErrorsEnum
func GetHandleGrantErrorsEnumStringValues() []string {
	return []string{
		"ABORT",
		"DROP_ACCOUNT",
		"IGNORE",
	}
}

// GetMappingHandleGrantErrorsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHandleGrantErrorsEnum(val string) (HandleGrantErrorsEnum, bool) {
	enum, ok := mappingHandleGrantErrorsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
