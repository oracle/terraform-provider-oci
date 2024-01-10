// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity Domains API
//
// Use the Identity Domains API to manage resources within an identity domain, for example, users, dynamic resource groups, groups, and identity providers. For information about managing resources within identity domains, see Identity and Access Management (with identity domains) (https://docs.oracle.com/iaas/Content/Identity/home.htm). This REST API is SCIM compliant.
// Use the table of contents and search tool to explore the Identity Domains API.
//

package identitydomains

import (
	"strings"
)

// AttributeSetsEnum Enum with underlying type: string
type AttributeSetsEnum string

// Set of constants representing the allowable values for AttributeSetsEnum
const (
	AttributeSetsAll     AttributeSetsEnum = "all"
	AttributeSetsAlways  AttributeSetsEnum = "always"
	AttributeSetsNever   AttributeSetsEnum = "never"
	AttributeSetsRequest AttributeSetsEnum = "request"
	AttributeSetsDefault AttributeSetsEnum = "default"
)

var mappingAttributeSetsEnum = map[string]AttributeSetsEnum{
	"all":     AttributeSetsAll,
	"always":  AttributeSetsAlways,
	"never":   AttributeSetsNever,
	"request": AttributeSetsRequest,
	"default": AttributeSetsDefault,
}

var mappingAttributeSetsEnumLowerCase = map[string]AttributeSetsEnum{
	"all":     AttributeSetsAll,
	"always":  AttributeSetsAlways,
	"never":   AttributeSetsNever,
	"request": AttributeSetsRequest,
	"default": AttributeSetsDefault,
}

// GetAttributeSetsEnumValues Enumerates the set of values for AttributeSetsEnum
func GetAttributeSetsEnumValues() []AttributeSetsEnum {
	values := make([]AttributeSetsEnum, 0)
	for _, v := range mappingAttributeSetsEnum {
		values = append(values, v)
	}
	return values
}

// GetAttributeSetsEnumStringValues Enumerates the set of values in String for AttributeSetsEnum
func GetAttributeSetsEnumStringValues() []string {
	return []string{
		"all",
		"always",
		"never",
		"request",
		"default",
	}
}

// GetMappingAttributeSetsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAttributeSetsEnum(val string) (AttributeSetsEnum, bool) {
	enum, ok := mappingAttributeSetsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
