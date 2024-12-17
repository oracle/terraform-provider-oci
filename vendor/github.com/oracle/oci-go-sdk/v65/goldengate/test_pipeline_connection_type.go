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

// TestPipelineConnectionTypeEnum Enum with underlying type: string
type TestPipelineConnectionTypeEnum string

// Set of constants representing the allowable values for TestPipelineConnectionTypeEnum
const (
	TestPipelineConnectionTypeDefault TestPipelineConnectionTypeEnum = "DEFAULT"
)

var mappingTestPipelineConnectionTypeEnum = map[string]TestPipelineConnectionTypeEnum{
	"DEFAULT": TestPipelineConnectionTypeDefault,
}

var mappingTestPipelineConnectionTypeEnumLowerCase = map[string]TestPipelineConnectionTypeEnum{
	"default": TestPipelineConnectionTypeDefault,
}

// GetTestPipelineConnectionTypeEnumValues Enumerates the set of values for TestPipelineConnectionTypeEnum
func GetTestPipelineConnectionTypeEnumValues() []TestPipelineConnectionTypeEnum {
	values := make([]TestPipelineConnectionTypeEnum, 0)
	for _, v := range mappingTestPipelineConnectionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTestPipelineConnectionTypeEnumStringValues Enumerates the set of values in String for TestPipelineConnectionTypeEnum
func GetTestPipelineConnectionTypeEnumStringValues() []string {
	return []string{
		"DEFAULT",
	}
}

// GetMappingTestPipelineConnectionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTestPipelineConnectionTypeEnum(val string) (TestPipelineConnectionTypeEnum, bool) {
	enum, ok := mappingTestPipelineConnectionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
