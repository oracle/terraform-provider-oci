// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// AuditProfileTargetTypeEnum Enum with underlying type: string
type AuditProfileTargetTypeEnum string

// Set of constants representing the allowable values for AuditProfileTargetTypeEnum
const (
	AuditProfileTargetTypeTargetDatabase      AuditProfileTargetTypeEnum = "TARGET_DATABASE"
	AuditProfileTargetTypeTargetDatabaseGroup AuditProfileTargetTypeEnum = "TARGET_DATABASE_GROUP"
)

var mappingAuditProfileTargetTypeEnum = map[string]AuditProfileTargetTypeEnum{
	"TARGET_DATABASE":       AuditProfileTargetTypeTargetDatabase,
	"TARGET_DATABASE_GROUP": AuditProfileTargetTypeTargetDatabaseGroup,
}

var mappingAuditProfileTargetTypeEnumLowerCase = map[string]AuditProfileTargetTypeEnum{
	"target_database":       AuditProfileTargetTypeTargetDatabase,
	"target_database_group": AuditProfileTargetTypeTargetDatabaseGroup,
}

// GetAuditProfileTargetTypeEnumValues Enumerates the set of values for AuditProfileTargetTypeEnum
func GetAuditProfileTargetTypeEnumValues() []AuditProfileTargetTypeEnum {
	values := make([]AuditProfileTargetTypeEnum, 0)
	for _, v := range mappingAuditProfileTargetTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAuditProfileTargetTypeEnumStringValues Enumerates the set of values in String for AuditProfileTargetTypeEnum
func GetAuditProfileTargetTypeEnumStringValues() []string {
	return []string{
		"TARGET_DATABASE",
		"TARGET_DATABASE_GROUP",
	}
}

// GetMappingAuditProfileTargetTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAuditProfileTargetTypeEnum(val string) (AuditProfileTargetTypeEnum, bool) {
	enum, ok := mappingAuditProfileTargetTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
