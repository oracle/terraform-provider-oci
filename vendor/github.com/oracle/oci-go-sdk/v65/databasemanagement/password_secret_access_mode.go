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

// PasswordSecretAccessModeEnum Enum with underlying type: string
type PasswordSecretAccessModeEnum string

// Set of constants representing the allowable values for PasswordSecretAccessModeEnum
const (
	PasswordSecretAccessModeUserPrincipal     PasswordSecretAccessModeEnum = "USER_PRINCIPAL"
	PasswordSecretAccessModeResourcePrincipal PasswordSecretAccessModeEnum = "RESOURCE_PRINCIPAL"
)

var mappingPasswordSecretAccessModeEnum = map[string]PasswordSecretAccessModeEnum{
	"USER_PRINCIPAL":     PasswordSecretAccessModeUserPrincipal,
	"RESOURCE_PRINCIPAL": PasswordSecretAccessModeResourcePrincipal,
}

var mappingPasswordSecretAccessModeEnumLowerCase = map[string]PasswordSecretAccessModeEnum{
	"user_principal":     PasswordSecretAccessModeUserPrincipal,
	"resource_principal": PasswordSecretAccessModeResourcePrincipal,
}

// GetPasswordSecretAccessModeEnumValues Enumerates the set of values for PasswordSecretAccessModeEnum
func GetPasswordSecretAccessModeEnumValues() []PasswordSecretAccessModeEnum {
	values := make([]PasswordSecretAccessModeEnum, 0)
	for _, v := range mappingPasswordSecretAccessModeEnum {
		values = append(values, v)
	}
	return values
}

// GetPasswordSecretAccessModeEnumStringValues Enumerates the set of values in String for PasswordSecretAccessModeEnum
func GetPasswordSecretAccessModeEnumStringValues() []string {
	return []string{
		"USER_PRINCIPAL",
		"RESOURCE_PRINCIPAL",
	}
}

// GetMappingPasswordSecretAccessModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPasswordSecretAccessModeEnum(val string) (PasswordSecretAccessModeEnum, bool) {
	enum, ok := mappingPasswordSecretAccessModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
