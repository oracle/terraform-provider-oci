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

// SeverityTypeEnum Enum with underlying type: string
type SeverityTypeEnum string

// Set of constants representing the allowable values for SeverityTypeEnum
const (
	SeverityTypeInfo    SeverityTypeEnum = "INFO"
	SeverityTypeError   SeverityTypeEnum = "ERROR"
	SeverityTypeWarning SeverityTypeEnum = "WARNING"
)

var mappingSeverityTypeEnum = map[string]SeverityTypeEnum{
	"INFO":    SeverityTypeInfo,
	"ERROR":   SeverityTypeError,
	"WARNING": SeverityTypeWarning,
}

var mappingSeverityTypeEnumLowerCase = map[string]SeverityTypeEnum{
	"info":    SeverityTypeInfo,
	"error":   SeverityTypeError,
	"warning": SeverityTypeWarning,
}

// GetSeverityTypeEnumValues Enumerates the set of values for SeverityTypeEnum
func GetSeverityTypeEnumValues() []SeverityTypeEnum {
	values := make([]SeverityTypeEnum, 0)
	for _, v := range mappingSeverityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSeverityTypeEnumStringValues Enumerates the set of values in String for SeverityTypeEnum
func GetSeverityTypeEnumStringValues() []string {
	return []string{
		"INFO",
		"ERROR",
		"WARNING",
	}
}

// GetMappingSeverityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSeverityTypeEnum(val string) (SeverityTypeEnum, bool) {
	enum, ok := mappingSeverityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
