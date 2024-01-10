// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Firewall API
//
// Use the Network Firewall API to create network firewalls and configure policies that regulates network traffic in and across VCNs.
//

package networkfirewall

import (
	"strings"
)

// MappedSecretTypeEnum Enum with underlying type: string
type MappedSecretTypeEnum string

// Set of constants representing the allowable values for MappedSecretTypeEnum
const (
	MappedSecretTypeOciVault MappedSecretTypeEnum = "OCI_VAULT"
)

var mappingMappedSecretTypeEnum = map[string]MappedSecretTypeEnum{
	"OCI_VAULT": MappedSecretTypeOciVault,
}

var mappingMappedSecretTypeEnumLowerCase = map[string]MappedSecretTypeEnum{
	"oci_vault": MappedSecretTypeOciVault,
}

// GetMappedSecretTypeEnumValues Enumerates the set of values for MappedSecretTypeEnum
func GetMappedSecretTypeEnumValues() []MappedSecretTypeEnum {
	values := make([]MappedSecretTypeEnum, 0)
	for _, v := range mappingMappedSecretTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMappedSecretTypeEnumStringValues Enumerates the set of values in String for MappedSecretTypeEnum
func GetMappedSecretTypeEnumStringValues() []string {
	return []string{
		"OCI_VAULT",
	}
}

// GetMappingMappedSecretTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMappedSecretTypeEnum(val string) (MappedSecretTypeEnum, bool) {
	enum, ok := mappingMappedSecretTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
