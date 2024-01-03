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

// IdcsPreventedOperationsEnum Enum with underlying type: string
type IdcsPreventedOperationsEnum string

// Set of constants representing the allowable values for IdcsPreventedOperationsEnum
const (
	IdcsPreventedOperationsReplace IdcsPreventedOperationsEnum = "replace"
	IdcsPreventedOperationsUpdate  IdcsPreventedOperationsEnum = "update"
	IdcsPreventedOperationsDelete  IdcsPreventedOperationsEnum = "delete"
)

var mappingIdcsPreventedOperationsEnum = map[string]IdcsPreventedOperationsEnum{
	"replace": IdcsPreventedOperationsReplace,
	"update":  IdcsPreventedOperationsUpdate,
	"delete":  IdcsPreventedOperationsDelete,
}

var mappingIdcsPreventedOperationsEnumLowerCase = map[string]IdcsPreventedOperationsEnum{
	"replace": IdcsPreventedOperationsReplace,
	"update":  IdcsPreventedOperationsUpdate,
	"delete":  IdcsPreventedOperationsDelete,
}

// GetIdcsPreventedOperationsEnumValues Enumerates the set of values for IdcsPreventedOperationsEnum
func GetIdcsPreventedOperationsEnumValues() []IdcsPreventedOperationsEnum {
	values := make([]IdcsPreventedOperationsEnum, 0)
	for _, v := range mappingIdcsPreventedOperationsEnum {
		values = append(values, v)
	}
	return values
}

// GetIdcsPreventedOperationsEnumStringValues Enumerates the set of values in String for IdcsPreventedOperationsEnum
func GetIdcsPreventedOperationsEnumStringValues() []string {
	return []string{
		"replace",
		"update",
		"delete",
	}
}

// GetMappingIdcsPreventedOperationsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIdcsPreventedOperationsEnum(val string) (IdcsPreventedOperationsEnum, bool) {
	enum, ok := mappingIdcsPreventedOperationsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
