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

// PreferredCredentialStatusEnum Enum with underlying type: string
type PreferredCredentialStatusEnum string

// Set of constants representing the allowable values for PreferredCredentialStatusEnum
const (
	PreferredCredentialStatusSet    PreferredCredentialStatusEnum = "SET"
	PreferredCredentialStatusNotSet PreferredCredentialStatusEnum = "NOT_SET"
)

var mappingPreferredCredentialStatusEnum = map[string]PreferredCredentialStatusEnum{
	"SET":     PreferredCredentialStatusSet,
	"NOT_SET": PreferredCredentialStatusNotSet,
}

var mappingPreferredCredentialStatusEnumLowerCase = map[string]PreferredCredentialStatusEnum{
	"set":     PreferredCredentialStatusSet,
	"not_set": PreferredCredentialStatusNotSet,
}

// GetPreferredCredentialStatusEnumValues Enumerates the set of values for PreferredCredentialStatusEnum
func GetPreferredCredentialStatusEnumValues() []PreferredCredentialStatusEnum {
	values := make([]PreferredCredentialStatusEnum, 0)
	for _, v := range mappingPreferredCredentialStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetPreferredCredentialStatusEnumStringValues Enumerates the set of values in String for PreferredCredentialStatusEnum
func GetPreferredCredentialStatusEnumStringValues() []string {
	return []string{
		"SET",
		"NOT_SET",
	}
}

// GetMappingPreferredCredentialStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPreferredCredentialStatusEnum(val string) (PreferredCredentialStatusEnum, bool) {
	enum, ok := mappingPreferredCredentialStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
