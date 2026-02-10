// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// AssessmentMigrationTypesEnum Enum with underlying type: string
type AssessmentMigrationTypesEnum string

// Set of constants representing the allowable values for AssessmentMigrationTypesEnum
const (
	AssessmentMigrationTypesOnline        AssessmentMigrationTypesEnum = "ONLINE"
	AssessmentMigrationTypesOnlineStandby AssessmentMigrationTypesEnum = "ONLINE_STANDBY"
	AssessmentMigrationTypesOffline       AssessmentMigrationTypesEnum = "OFFLINE"
)

var mappingAssessmentMigrationTypesEnum = map[string]AssessmentMigrationTypesEnum{
	"ONLINE":         AssessmentMigrationTypesOnline,
	"ONLINE_STANDBY": AssessmentMigrationTypesOnlineStandby,
	"OFFLINE":        AssessmentMigrationTypesOffline,
}

var mappingAssessmentMigrationTypesEnumLowerCase = map[string]AssessmentMigrationTypesEnum{
	"online":         AssessmentMigrationTypesOnline,
	"online_standby": AssessmentMigrationTypesOnlineStandby,
	"offline":        AssessmentMigrationTypesOffline,
}

// GetAssessmentMigrationTypesEnumValues Enumerates the set of values for AssessmentMigrationTypesEnum
func GetAssessmentMigrationTypesEnumValues() []AssessmentMigrationTypesEnum {
	values := make([]AssessmentMigrationTypesEnum, 0)
	for _, v := range mappingAssessmentMigrationTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetAssessmentMigrationTypesEnumStringValues Enumerates the set of values in String for AssessmentMigrationTypesEnum
func GetAssessmentMigrationTypesEnumStringValues() []string {
	return []string{
		"ONLINE",
		"ONLINE_STANDBY",
		"OFFLINE",
	}
}

// GetMappingAssessmentMigrationTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAssessmentMigrationTypesEnum(val string) (AssessmentMigrationTypesEnum, bool) {
	enum, ok := mappingAssessmentMigrationTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
