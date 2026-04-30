// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools API
//
// Use the Database Tools API to manage connections, private endpoints, and work requests in the Database Tools service.
//

package databasetools

import (
	"strings"
)

// IdentityProviderTypeEnum Enum with underlying type: string
type IdentityProviderTypeEnum string

// Set of constants representing the allowable values for IdentityProviderTypeEnum
const (
	IdentityProviderTypeOciIam  IdentityProviderTypeEnum = "OCI_IAM"
	IdentityProviderTypeAzureAd IdentityProviderTypeEnum = "AZURE_AD"
	IdentityProviderTypeNone    IdentityProviderTypeEnum = "NONE"
)

var mappingIdentityProviderTypeEnum = map[string]IdentityProviderTypeEnum{
	"OCI_IAM":  IdentityProviderTypeOciIam,
	"AZURE_AD": IdentityProviderTypeAzureAd,
	"NONE":     IdentityProviderTypeNone,
}

var mappingIdentityProviderTypeEnumLowerCase = map[string]IdentityProviderTypeEnum{
	"oci_iam":  IdentityProviderTypeOciIam,
	"azure_ad": IdentityProviderTypeAzureAd,
	"none":     IdentityProviderTypeNone,
}

// GetIdentityProviderTypeEnumValues Enumerates the set of values for IdentityProviderTypeEnum
func GetIdentityProviderTypeEnumValues() []IdentityProviderTypeEnum {
	values := make([]IdentityProviderTypeEnum, 0)
	for _, v := range mappingIdentityProviderTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetIdentityProviderTypeEnumStringValues Enumerates the set of values in String for IdentityProviderTypeEnum
func GetIdentityProviderTypeEnumStringValues() []string {
	return []string{
		"OCI_IAM",
		"AZURE_AD",
		"NONE",
	}
}

// GetMappingIdentityProviderTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIdentityProviderTypeEnum(val string) (IdentityProviderTypeEnum, bool) {
	enum, ok := mappingIdentityProviderTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
