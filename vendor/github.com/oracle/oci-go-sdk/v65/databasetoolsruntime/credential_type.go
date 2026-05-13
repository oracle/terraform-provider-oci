// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools Runtime API
//
// Use the Database Tools Runtime API to connect to databases through Database Tools Connections.
//

package databasetoolsruntime

import (
	"strings"
)

// CredentialTypeEnum Enum with underlying type: string
type CredentialTypeEnum string

// Set of constants representing the allowable values for CredentialTypeEnum
const (
	CredentialTypeBasic CredentialTypeEnum = "BASIC"
)

var mappingCredentialTypeEnum = map[string]CredentialTypeEnum{
	"BASIC": CredentialTypeBasic,
}

var mappingCredentialTypeEnumLowerCase = map[string]CredentialTypeEnum{
	"basic": CredentialTypeBasic,
}

// GetCredentialTypeEnumValues Enumerates the set of values for CredentialTypeEnum
func GetCredentialTypeEnumValues() []CredentialTypeEnum {
	values := make([]CredentialTypeEnum, 0)
	for _, v := range mappingCredentialTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCredentialTypeEnumStringValues Enumerates the set of values in String for CredentialTypeEnum
func GetCredentialTypeEnumStringValues() []string {
	return []string{
		"BASIC",
	}
}

// GetMappingCredentialTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCredentialTypeEnum(val string) (CredentialTypeEnum, bool) {
	enum, ok := mappingCredentialTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
