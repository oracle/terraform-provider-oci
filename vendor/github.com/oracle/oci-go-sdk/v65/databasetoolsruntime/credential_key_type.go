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

// CredentialKeyTypeEnum Enum with underlying type: string
type CredentialKeyTypeEnum string

// Set of constants representing the allowable values for CredentialKeyTypeEnum
const (
	CredentialKeyTypeCredentialName CredentialKeyTypeEnum = "CREDENTIAL_NAME"
	CredentialKeyTypePublicSynonym  CredentialKeyTypeEnum = "PUBLIC_SYNONYM"
)

var mappingCredentialKeyTypeEnum = map[string]CredentialKeyTypeEnum{
	"CREDENTIAL_NAME": CredentialKeyTypeCredentialName,
	"PUBLIC_SYNONYM":  CredentialKeyTypePublicSynonym,
}

var mappingCredentialKeyTypeEnumLowerCase = map[string]CredentialKeyTypeEnum{
	"credential_name": CredentialKeyTypeCredentialName,
	"public_synonym":  CredentialKeyTypePublicSynonym,
}

// GetCredentialKeyTypeEnumValues Enumerates the set of values for CredentialKeyTypeEnum
func GetCredentialKeyTypeEnumValues() []CredentialKeyTypeEnum {
	values := make([]CredentialKeyTypeEnum, 0)
	for _, v := range mappingCredentialKeyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCredentialKeyTypeEnumStringValues Enumerates the set of values in String for CredentialKeyTypeEnum
func GetCredentialKeyTypeEnumStringValues() []string {
	return []string{
		"CREDENTIAL_NAME",
		"PUBLIC_SYNONYM",
	}
}

// GetMappingCredentialKeyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCredentialKeyTypeEnum(val string) (CredentialKeyTypeEnum, bool) {
	enum, ok := mappingCredentialKeyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
