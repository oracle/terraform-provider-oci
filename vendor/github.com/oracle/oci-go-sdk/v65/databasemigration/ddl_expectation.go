// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// DdlExpectationEnum Enum with underlying type: string
type DdlExpectationEnum string

// Set of constants representing the allowable values for DdlExpectationEnum
const (
	DdlExpectationDdlExpected    DdlExpectationEnum = "DDL_EXPECTED"
	DdlExpectationDdlNotExpected DdlExpectationEnum = "DDL_NOT_EXPECTED"
)

var mappingDdlExpectationEnum = map[string]DdlExpectationEnum{
	"DDL_EXPECTED":     DdlExpectationDdlExpected,
	"DDL_NOT_EXPECTED": DdlExpectationDdlNotExpected,
}

var mappingDdlExpectationEnumLowerCase = map[string]DdlExpectationEnum{
	"ddl_expected":     DdlExpectationDdlExpected,
	"ddl_not_expected": DdlExpectationDdlNotExpected,
}

// GetDdlExpectationEnumValues Enumerates the set of values for DdlExpectationEnum
func GetDdlExpectationEnumValues() []DdlExpectationEnum {
	values := make([]DdlExpectationEnum, 0)
	for _, v := range mappingDdlExpectationEnum {
		values = append(values, v)
	}
	return values
}

// GetDdlExpectationEnumStringValues Enumerates the set of values in String for DdlExpectationEnum
func GetDdlExpectationEnumStringValues() []string {
	return []string{
		"DDL_EXPECTED",
		"DDL_NOT_EXPECTED",
	}
}

// GetMappingDdlExpectationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDdlExpectationEnum(val string) (DdlExpectationEnum, bool) {
	enum, ok := mappingDdlExpectationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
