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

// SecurityAssessmentTargetTypeEnum Enum with underlying type: string
type SecurityAssessmentTargetTypeEnum string

// Set of constants representing the allowable values for SecurityAssessmentTargetTypeEnum
const (
	SecurityAssessmentTargetTypeTargetDatabase      SecurityAssessmentTargetTypeEnum = "TARGET_DATABASE"
	SecurityAssessmentTargetTypeTargetDatabaseGroup SecurityAssessmentTargetTypeEnum = "TARGET_DATABASE_GROUP"
)

var mappingSecurityAssessmentTargetTypeEnum = map[string]SecurityAssessmentTargetTypeEnum{
	"TARGET_DATABASE":       SecurityAssessmentTargetTypeTargetDatabase,
	"TARGET_DATABASE_GROUP": SecurityAssessmentTargetTypeTargetDatabaseGroup,
}

var mappingSecurityAssessmentTargetTypeEnumLowerCase = map[string]SecurityAssessmentTargetTypeEnum{
	"target_database":       SecurityAssessmentTargetTypeTargetDatabase,
	"target_database_group": SecurityAssessmentTargetTypeTargetDatabaseGroup,
}

// GetSecurityAssessmentTargetTypeEnumValues Enumerates the set of values for SecurityAssessmentTargetTypeEnum
func GetSecurityAssessmentTargetTypeEnumValues() []SecurityAssessmentTargetTypeEnum {
	values := make([]SecurityAssessmentTargetTypeEnum, 0)
	for _, v := range mappingSecurityAssessmentTargetTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSecurityAssessmentTargetTypeEnumStringValues Enumerates the set of values in String for SecurityAssessmentTargetTypeEnum
func GetSecurityAssessmentTargetTypeEnumStringValues() []string {
	return []string{
		"TARGET_DATABASE",
		"TARGET_DATABASE_GROUP",
	}
}

// GetMappingSecurityAssessmentTargetTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecurityAssessmentTargetTypeEnum(val string) (SecurityAssessmentTargetTypeEnum, bool) {
	enum, ok := mappingSecurityAssessmentTargetTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
