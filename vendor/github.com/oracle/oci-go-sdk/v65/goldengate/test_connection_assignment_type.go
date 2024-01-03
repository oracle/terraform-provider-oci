// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"strings"
)

// TestConnectionAssignmentTypeEnum Enum with underlying type: string
type TestConnectionAssignmentTypeEnum string

// Set of constants representing the allowable values for TestConnectionAssignmentTypeEnum
const (
	TestConnectionAssignmentTypeDefault TestConnectionAssignmentTypeEnum = "DEFAULT"
)

var mappingTestConnectionAssignmentTypeEnum = map[string]TestConnectionAssignmentTypeEnum{
	"DEFAULT": TestConnectionAssignmentTypeDefault,
}

var mappingTestConnectionAssignmentTypeEnumLowerCase = map[string]TestConnectionAssignmentTypeEnum{
	"default": TestConnectionAssignmentTypeDefault,
}

// GetTestConnectionAssignmentTypeEnumValues Enumerates the set of values for TestConnectionAssignmentTypeEnum
func GetTestConnectionAssignmentTypeEnumValues() []TestConnectionAssignmentTypeEnum {
	values := make([]TestConnectionAssignmentTypeEnum, 0)
	for _, v := range mappingTestConnectionAssignmentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTestConnectionAssignmentTypeEnumStringValues Enumerates the set of values in String for TestConnectionAssignmentTypeEnum
func GetTestConnectionAssignmentTypeEnumStringValues() []string {
	return []string{
		"DEFAULT",
	}
}

// GetMappingTestConnectionAssignmentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTestConnectionAssignmentTypeEnum(val string) (TestConnectionAssignmentTypeEnum, bool) {
	enum, ok := mappingTestConnectionAssignmentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
