// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"strings"
)

// NamedCredentialContentTypeEnum Enum with underlying type: string
type NamedCredentialContentTypeEnum string

// Set of constants representing the allowable values for NamedCredentialContentTypeEnum
const (
	NamedCredentialContentTypeBasic NamedCredentialContentTypeEnum = "BASIC"
)

var mappingNamedCredentialContentTypeEnum = map[string]NamedCredentialContentTypeEnum{
	"BASIC": NamedCredentialContentTypeBasic,
}

var mappingNamedCredentialContentTypeEnumLowerCase = map[string]NamedCredentialContentTypeEnum{
	"basic": NamedCredentialContentTypeBasic,
}

// GetNamedCredentialContentTypeEnumValues Enumerates the set of values for NamedCredentialContentTypeEnum
func GetNamedCredentialContentTypeEnumValues() []NamedCredentialContentTypeEnum {
	values := make([]NamedCredentialContentTypeEnum, 0)
	for _, v := range mappingNamedCredentialContentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetNamedCredentialContentTypeEnumStringValues Enumerates the set of values in String for NamedCredentialContentTypeEnum
func GetNamedCredentialContentTypeEnumStringValues() []string {
	return []string{
		"BASIC",
	}
}

// GetMappingNamedCredentialContentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNamedCredentialContentTypeEnum(val string) (NamedCredentialContentTypeEnum, bool) {
	enum, ok := mappingNamedCredentialContentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
