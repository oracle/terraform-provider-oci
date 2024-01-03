// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"strings"
)

// PrivilegeGrantableOptionEnum Enum with underlying type: string
type PrivilegeGrantableOptionEnum string

// Set of constants representing the allowable values for PrivilegeGrantableOptionEnum
const (
	PrivilegeGrantableOptionAdminOption PrivilegeGrantableOptionEnum = "ADMIN_OPTION"
	PrivilegeGrantableOptionGrantOption PrivilegeGrantableOptionEnum = "GRANT_OPTION"
)

var mappingPrivilegeGrantableOptionEnum = map[string]PrivilegeGrantableOptionEnum{
	"ADMIN_OPTION": PrivilegeGrantableOptionAdminOption,
	"GRANT_OPTION": PrivilegeGrantableOptionGrantOption,
}

var mappingPrivilegeGrantableOptionEnumLowerCase = map[string]PrivilegeGrantableOptionEnum{
	"admin_option": PrivilegeGrantableOptionAdminOption,
	"grant_option": PrivilegeGrantableOptionGrantOption,
}

// GetPrivilegeGrantableOptionEnumValues Enumerates the set of values for PrivilegeGrantableOptionEnum
func GetPrivilegeGrantableOptionEnumValues() []PrivilegeGrantableOptionEnum {
	values := make([]PrivilegeGrantableOptionEnum, 0)
	for _, v := range mappingPrivilegeGrantableOptionEnum {
		values = append(values, v)
	}
	return values
}

// GetPrivilegeGrantableOptionEnumStringValues Enumerates the set of values in String for PrivilegeGrantableOptionEnum
func GetPrivilegeGrantableOptionEnumStringValues() []string {
	return []string{
		"ADMIN_OPTION",
		"GRANT_OPTION",
	}
}

// GetMappingPrivilegeGrantableOptionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPrivilegeGrantableOptionEnum(val string) (PrivilegeGrantableOptionEnum, bool) {
	enum, ok := mappingPrivilegeGrantableOptionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
