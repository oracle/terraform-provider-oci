// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// UserAssessmentTargetTypeEnum Enum with underlying type: string
type UserAssessmentTargetTypeEnum string

// Set of constants representing the allowable values for UserAssessmentTargetTypeEnum
const (
	UserAssessmentTargetTypeTargetDatabase      UserAssessmentTargetTypeEnum = "TARGET_DATABASE"
	UserAssessmentTargetTypeTargetDatabaseGroup UserAssessmentTargetTypeEnum = "TARGET_DATABASE_GROUP"
)

var mappingUserAssessmentTargetTypeEnum = map[string]UserAssessmentTargetTypeEnum{
	"TARGET_DATABASE":       UserAssessmentTargetTypeTargetDatabase,
	"TARGET_DATABASE_GROUP": UserAssessmentTargetTypeTargetDatabaseGroup,
}

var mappingUserAssessmentTargetTypeEnumLowerCase = map[string]UserAssessmentTargetTypeEnum{
	"target_database":       UserAssessmentTargetTypeTargetDatabase,
	"target_database_group": UserAssessmentTargetTypeTargetDatabaseGroup,
}

// GetUserAssessmentTargetTypeEnumValues Enumerates the set of values for UserAssessmentTargetTypeEnum
func GetUserAssessmentTargetTypeEnumValues() []UserAssessmentTargetTypeEnum {
	values := make([]UserAssessmentTargetTypeEnum, 0)
	for _, v := range mappingUserAssessmentTargetTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUserAssessmentTargetTypeEnumStringValues Enumerates the set of values in String for UserAssessmentTargetTypeEnum
func GetUserAssessmentTargetTypeEnumStringValues() []string {
	return []string{
		"TARGET_DATABASE",
		"TARGET_DATABASE_GROUP",
	}
}

// GetMappingUserAssessmentTargetTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUserAssessmentTargetTypeEnum(val string) (UserAssessmentTargetTypeEnum, bool) {
	enum, ok := mappingUserAssessmentTargetTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
