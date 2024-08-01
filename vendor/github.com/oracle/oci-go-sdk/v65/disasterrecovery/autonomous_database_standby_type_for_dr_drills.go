// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (DR) API to manage disaster recovery for business applications.
// Full Stack DR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster
// recovery capabilities for all layers of an application stack, including infrastructure, middleware, database,
// and application.
//

package disasterrecovery

import (
	"strings"
)

// AutonomousDatabaseStandbyTypeForDrDrillsEnum Enum with underlying type: string
type AutonomousDatabaseStandbyTypeForDrDrillsEnum string

// Set of constants representing the allowable values for AutonomousDatabaseStandbyTypeForDrDrillsEnum
const (
	AutonomousDatabaseStandbyTypeForDrDrillsFullClone        AutonomousDatabaseStandbyTypeForDrDrillsEnum = "FULL_CLONE"
	AutonomousDatabaseStandbyTypeForDrDrillsRefreshableClone AutonomousDatabaseStandbyTypeForDrDrillsEnum = "REFRESHABLE_CLONE"
	AutonomousDatabaseStandbyTypeForDrDrillsSnapshotStandby  AutonomousDatabaseStandbyTypeForDrDrillsEnum = "SNAPSHOT_STANDBY"
)

var mappingAutonomousDatabaseStandbyTypeForDrDrillsEnum = map[string]AutonomousDatabaseStandbyTypeForDrDrillsEnum{
	"FULL_CLONE":        AutonomousDatabaseStandbyTypeForDrDrillsFullClone,
	"REFRESHABLE_CLONE": AutonomousDatabaseStandbyTypeForDrDrillsRefreshableClone,
	"SNAPSHOT_STANDBY":  AutonomousDatabaseStandbyTypeForDrDrillsSnapshotStandby,
}

var mappingAutonomousDatabaseStandbyTypeForDrDrillsEnumLowerCase = map[string]AutonomousDatabaseStandbyTypeForDrDrillsEnum{
	"full_clone":        AutonomousDatabaseStandbyTypeForDrDrillsFullClone,
	"refreshable_clone": AutonomousDatabaseStandbyTypeForDrDrillsRefreshableClone,
	"snapshot_standby":  AutonomousDatabaseStandbyTypeForDrDrillsSnapshotStandby,
}

// GetAutonomousDatabaseStandbyTypeForDrDrillsEnumValues Enumerates the set of values for AutonomousDatabaseStandbyTypeForDrDrillsEnum
func GetAutonomousDatabaseStandbyTypeForDrDrillsEnumValues() []AutonomousDatabaseStandbyTypeForDrDrillsEnum {
	values := make([]AutonomousDatabaseStandbyTypeForDrDrillsEnum, 0)
	for _, v := range mappingAutonomousDatabaseStandbyTypeForDrDrillsEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseStandbyTypeForDrDrillsEnumStringValues Enumerates the set of values in String for AutonomousDatabaseStandbyTypeForDrDrillsEnum
func GetAutonomousDatabaseStandbyTypeForDrDrillsEnumStringValues() []string {
	return []string{
		"FULL_CLONE",
		"REFRESHABLE_CLONE",
		"SNAPSHOT_STANDBY",
	}
}

// GetMappingAutonomousDatabaseStandbyTypeForDrDrillsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseStandbyTypeForDrDrillsEnum(val string) (AutonomousDatabaseStandbyTypeForDrDrillsEnum, bool) {
	enum, ok := mappingAutonomousDatabaseStandbyTypeForDrDrillsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
