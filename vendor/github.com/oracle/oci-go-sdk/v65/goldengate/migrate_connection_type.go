// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// MigrateConnectionTypeEnum Enum with underlying type: string
type MigrateConnectionTypeEnum string

// Set of constants representing the allowable values for MigrateConnectionTypeEnum
const (
	MigrateConnectionTypeSecret MigrateConnectionTypeEnum = "SECRET"
)

var mappingMigrateConnectionTypeEnum = map[string]MigrateConnectionTypeEnum{
	"SECRET": MigrateConnectionTypeSecret,
}

var mappingMigrateConnectionTypeEnumLowerCase = map[string]MigrateConnectionTypeEnum{
	"secret": MigrateConnectionTypeSecret,
}

// GetMigrateConnectionTypeEnumValues Enumerates the set of values for MigrateConnectionTypeEnum
func GetMigrateConnectionTypeEnumValues() []MigrateConnectionTypeEnum {
	values := make([]MigrateConnectionTypeEnum, 0)
	for _, v := range mappingMigrateConnectionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMigrateConnectionTypeEnumStringValues Enumerates the set of values in String for MigrateConnectionTypeEnum
func GetMigrateConnectionTypeEnumStringValues() []string {
	return []string{
		"SECRET",
	}
}

// GetMappingMigrateConnectionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMigrateConnectionTypeEnum(val string) (MigrateConnectionTypeEnum, bool) {
	enum, ok := mappingMigrateConnectionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
