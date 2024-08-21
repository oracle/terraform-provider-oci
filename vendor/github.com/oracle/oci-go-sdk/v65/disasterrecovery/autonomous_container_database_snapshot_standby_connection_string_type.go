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

// AutonomousContainerDatabaseSnapshotStandbyConnectionStringTypeEnum Enum with underlying type: string
type AutonomousContainerDatabaseSnapshotStandbyConnectionStringTypeEnum string

// Set of constants representing the allowable values for AutonomousContainerDatabaseSnapshotStandbyConnectionStringTypeEnum
const (
	AutonomousContainerDatabaseSnapshotStandbyConnectionStringTypeSnapshotService AutonomousContainerDatabaseSnapshotStandbyConnectionStringTypeEnum = "SNAPSHOT_SERVICE"
	AutonomousContainerDatabaseSnapshotStandbyConnectionStringTypePrimaryService  AutonomousContainerDatabaseSnapshotStandbyConnectionStringTypeEnum = "PRIMARY_SERVICE"
)

var mappingAutonomousContainerDatabaseSnapshotStandbyConnectionStringTypeEnum = map[string]AutonomousContainerDatabaseSnapshotStandbyConnectionStringTypeEnum{
	"SNAPSHOT_SERVICE": AutonomousContainerDatabaseSnapshotStandbyConnectionStringTypeSnapshotService,
	"PRIMARY_SERVICE":  AutonomousContainerDatabaseSnapshotStandbyConnectionStringTypePrimaryService,
}

var mappingAutonomousContainerDatabaseSnapshotStandbyConnectionStringTypeEnumLowerCase = map[string]AutonomousContainerDatabaseSnapshotStandbyConnectionStringTypeEnum{
	"snapshot_service": AutonomousContainerDatabaseSnapshotStandbyConnectionStringTypeSnapshotService,
	"primary_service":  AutonomousContainerDatabaseSnapshotStandbyConnectionStringTypePrimaryService,
}

// GetAutonomousContainerDatabaseSnapshotStandbyConnectionStringTypeEnumValues Enumerates the set of values for AutonomousContainerDatabaseSnapshotStandbyConnectionStringTypeEnum
func GetAutonomousContainerDatabaseSnapshotStandbyConnectionStringTypeEnumValues() []AutonomousContainerDatabaseSnapshotStandbyConnectionStringTypeEnum {
	values := make([]AutonomousContainerDatabaseSnapshotStandbyConnectionStringTypeEnum, 0)
	for _, v := range mappingAutonomousContainerDatabaseSnapshotStandbyConnectionStringTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousContainerDatabaseSnapshotStandbyConnectionStringTypeEnumStringValues Enumerates the set of values in String for AutonomousContainerDatabaseSnapshotStandbyConnectionStringTypeEnum
func GetAutonomousContainerDatabaseSnapshotStandbyConnectionStringTypeEnumStringValues() []string {
	return []string{
		"SNAPSHOT_SERVICE",
		"PRIMARY_SERVICE",
	}
}

// GetMappingAutonomousContainerDatabaseSnapshotStandbyConnectionStringTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousContainerDatabaseSnapshotStandbyConnectionStringTypeEnum(val string) (AutonomousContainerDatabaseSnapshotStandbyConnectionStringTypeEnum, bool) {
	enum, ok := mappingAutonomousContainerDatabaseSnapshotStandbyConnectionStringTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
