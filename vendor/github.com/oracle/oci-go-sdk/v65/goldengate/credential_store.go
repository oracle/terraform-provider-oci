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

// CredentialStoreEnum Enum with underlying type: string
type CredentialStoreEnum string

// Set of constants representing the allowable values for CredentialStoreEnum
const (
	CredentialStoreGoldengate CredentialStoreEnum = "GOLDENGATE"
	CredentialStoreIam        CredentialStoreEnum = "IAM"
)

var mappingCredentialStoreEnum = map[string]CredentialStoreEnum{
	"GOLDENGATE": CredentialStoreGoldengate,
	"IAM":        CredentialStoreIam,
}

var mappingCredentialStoreEnumLowerCase = map[string]CredentialStoreEnum{
	"goldengate": CredentialStoreGoldengate,
	"iam":        CredentialStoreIam,
}

// GetCredentialStoreEnumValues Enumerates the set of values for CredentialStoreEnum
func GetCredentialStoreEnumValues() []CredentialStoreEnum {
	values := make([]CredentialStoreEnum, 0)
	for _, v := range mappingCredentialStoreEnum {
		values = append(values, v)
	}
	return values
}

// GetCredentialStoreEnumStringValues Enumerates the set of values in String for CredentialStoreEnum
func GetCredentialStoreEnumStringValues() []string {
	return []string{
		"GOLDENGATE",
		"IAM",
	}
}

// GetMappingCredentialStoreEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCredentialStoreEnum(val string) (CredentialStoreEnum, bool) {
	enum, ok := mappingCredentialStoreEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
