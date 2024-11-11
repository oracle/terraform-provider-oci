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

// RefreshConnectionTypeEnum Enum with underlying type: string
type RefreshConnectionTypeEnum string

// Set of constants representing the allowable values for RefreshConnectionTypeEnum
const (
	RefreshConnectionTypeDefault RefreshConnectionTypeEnum = "DEFAULT"
)

var mappingRefreshConnectionTypeEnum = map[string]RefreshConnectionTypeEnum{
	"DEFAULT": RefreshConnectionTypeDefault,
}

var mappingRefreshConnectionTypeEnumLowerCase = map[string]RefreshConnectionTypeEnum{
	"default": RefreshConnectionTypeDefault,
}

// GetRefreshConnectionTypeEnumValues Enumerates the set of values for RefreshConnectionTypeEnum
func GetRefreshConnectionTypeEnumValues() []RefreshConnectionTypeEnum {
	values := make([]RefreshConnectionTypeEnum, 0)
	for _, v := range mappingRefreshConnectionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRefreshConnectionTypeEnumStringValues Enumerates the set of values in String for RefreshConnectionTypeEnum
func GetRefreshConnectionTypeEnumStringValues() []string {
	return []string{
		"DEFAULT",
	}
}

// GetMappingRefreshConnectionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRefreshConnectionTypeEnum(val string) (RefreshConnectionTypeEnum, bool) {
	enum, ok := mappingRefreshConnectionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
