// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (FSDR) API to manage disaster recovery for business applications.
// FSDR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster recovery
// capabilities for all layers of an application stack, including infrastructure, middleware, database, and application.
//

package disasterrecovery

import (
	"strings"
)

// DrProtectionGroupMemberTypeEnum Enum with underlying type: string
type DrProtectionGroupMemberTypeEnum string

// Set of constants representing the allowable values for DrProtectionGroupMemberTypeEnum
const (
	DrProtectionGroupMemberTypeComputeInstance    DrProtectionGroupMemberTypeEnum = "COMPUTE_INSTANCE"
	DrProtectionGroupMemberTypeVolumeGroup        DrProtectionGroupMemberTypeEnum = "VOLUME_GROUP"
	DrProtectionGroupMemberTypeDatabase           DrProtectionGroupMemberTypeEnum = "DATABASE"
	DrProtectionGroupMemberTypeAutonomousDatabase DrProtectionGroupMemberTypeEnum = "AUTONOMOUS_DATABASE"
)

var mappingDrProtectionGroupMemberTypeEnum = map[string]DrProtectionGroupMemberTypeEnum{
	"COMPUTE_INSTANCE":    DrProtectionGroupMemberTypeComputeInstance,
	"VOLUME_GROUP":        DrProtectionGroupMemberTypeVolumeGroup,
	"DATABASE":            DrProtectionGroupMemberTypeDatabase,
	"AUTONOMOUS_DATABASE": DrProtectionGroupMemberTypeAutonomousDatabase,
}

var mappingDrProtectionGroupMemberTypeEnumLowerCase = map[string]DrProtectionGroupMemberTypeEnum{
	"compute_instance":    DrProtectionGroupMemberTypeComputeInstance,
	"volume_group":        DrProtectionGroupMemberTypeVolumeGroup,
	"database":            DrProtectionGroupMemberTypeDatabase,
	"autonomous_database": DrProtectionGroupMemberTypeAutonomousDatabase,
}

// GetDrProtectionGroupMemberTypeEnumValues Enumerates the set of values for DrProtectionGroupMemberTypeEnum
func GetDrProtectionGroupMemberTypeEnumValues() []DrProtectionGroupMemberTypeEnum {
	values := make([]DrProtectionGroupMemberTypeEnum, 0)
	for _, v := range mappingDrProtectionGroupMemberTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDrProtectionGroupMemberTypeEnumStringValues Enumerates the set of values in String for DrProtectionGroupMemberTypeEnum
func GetDrProtectionGroupMemberTypeEnumStringValues() []string {
	return []string{
		"COMPUTE_INSTANCE",
		"VOLUME_GROUP",
		"DATABASE",
		"AUTONOMOUS_DATABASE",
	}
}

// GetMappingDrProtectionGroupMemberTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrProtectionGroupMemberTypeEnum(val string) (DrProtectionGroupMemberTypeEnum, bool) {
	enum, ok := mappingDrProtectionGroupMemberTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
