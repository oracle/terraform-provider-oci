// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"strings"
)

// CompatibilityOptionEnum Enum with underlying type: string
type CompatibilityOptionEnum string

// Set of constants representing the allowable values for CompatibilityOptionEnum
const (
	CompatibilityOptionForceInnodb           CompatibilityOptionEnum = "FORCE_INNODB"
	CompatibilityOptionSkipInvalidAccounts   CompatibilityOptionEnum = "SKIP_INVALID_ACCOUNTS"
	CompatibilityOptionStripDefiners         CompatibilityOptionEnum = "STRIP_DEFINERS"
	CompatibilityOptionStripRestrictedGrants CompatibilityOptionEnum = "STRIP_RESTRICTED_GRANTS"
	CompatibilityOptionStripTablespaces      CompatibilityOptionEnum = "STRIP_TABLESPACES"
	CompatibilityOptionIgnoreWildcardGrants  CompatibilityOptionEnum = "IGNORE_WILDCARD_GRANTS"
	CompatibilityOptionStripInvalidGrants    CompatibilityOptionEnum = "STRIP_INVALID_GRANTS"
)

var mappingCompatibilityOptionEnum = map[string]CompatibilityOptionEnum{
	"FORCE_INNODB":            CompatibilityOptionForceInnodb,
	"SKIP_INVALID_ACCOUNTS":   CompatibilityOptionSkipInvalidAccounts,
	"STRIP_DEFINERS":          CompatibilityOptionStripDefiners,
	"STRIP_RESTRICTED_GRANTS": CompatibilityOptionStripRestrictedGrants,
	"STRIP_TABLESPACES":       CompatibilityOptionStripTablespaces,
	"IGNORE_WILDCARD_GRANTS":  CompatibilityOptionIgnoreWildcardGrants,
	"STRIP_INVALID_GRANTS":    CompatibilityOptionStripInvalidGrants,
}

var mappingCompatibilityOptionEnumLowerCase = map[string]CompatibilityOptionEnum{
	"force_innodb":            CompatibilityOptionForceInnodb,
	"skip_invalid_accounts":   CompatibilityOptionSkipInvalidAccounts,
	"strip_definers":          CompatibilityOptionStripDefiners,
	"strip_restricted_grants": CompatibilityOptionStripRestrictedGrants,
	"strip_tablespaces":       CompatibilityOptionStripTablespaces,
	"ignore_wildcard_grants":  CompatibilityOptionIgnoreWildcardGrants,
	"strip_invalid_grants":    CompatibilityOptionStripInvalidGrants,
}

// GetCompatibilityOptionEnumValues Enumerates the set of values for CompatibilityOptionEnum
func GetCompatibilityOptionEnumValues() []CompatibilityOptionEnum {
	values := make([]CompatibilityOptionEnum, 0)
	for _, v := range mappingCompatibilityOptionEnum {
		values = append(values, v)
	}
	return values
}

// GetCompatibilityOptionEnumStringValues Enumerates the set of values in String for CompatibilityOptionEnum
func GetCompatibilityOptionEnumStringValues() []string {
	return []string{
		"FORCE_INNODB",
		"SKIP_INVALID_ACCOUNTS",
		"STRIP_DEFINERS",
		"STRIP_RESTRICTED_GRANTS",
		"STRIP_TABLESPACES",
		"IGNORE_WILDCARD_GRANTS",
		"STRIP_INVALID_GRANTS",
	}
}

// GetMappingCompatibilityOptionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCompatibilityOptionEnum(val string) (CompatibilityOptionEnum, bool) {
	enum, ok := mappingCompatibilityOptionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
