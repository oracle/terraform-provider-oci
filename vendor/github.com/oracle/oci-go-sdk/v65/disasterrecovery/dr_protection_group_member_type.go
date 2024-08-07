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

// DrProtectionGroupMemberTypeEnum Enum with underlying type: string
type DrProtectionGroupMemberTypeEnum string

// Set of constants representing the allowable values for DrProtectionGroupMemberTypeEnum
const (
	DrProtectionGroupMemberTypeComputeInstance             DrProtectionGroupMemberTypeEnum = "COMPUTE_INSTANCE"
	DrProtectionGroupMemberTypeComputeInstanceMovable      DrProtectionGroupMemberTypeEnum = "COMPUTE_INSTANCE_MOVABLE"
	DrProtectionGroupMemberTypeComputeInstanceNonMovable   DrProtectionGroupMemberTypeEnum = "COMPUTE_INSTANCE_NON_MOVABLE"
	DrProtectionGroupMemberTypeVolumeGroup                 DrProtectionGroupMemberTypeEnum = "VOLUME_GROUP"
	DrProtectionGroupMemberTypeDatabase                    DrProtectionGroupMemberTypeEnum = "DATABASE"
	DrProtectionGroupMemberTypeAutonomousDatabase          DrProtectionGroupMemberTypeEnum = "AUTONOMOUS_DATABASE"
	DrProtectionGroupMemberTypeAutonomousContainerDatabase DrProtectionGroupMemberTypeEnum = "AUTONOMOUS_CONTAINER_DATABASE"
	DrProtectionGroupMemberTypeLoadBalancer                DrProtectionGroupMemberTypeEnum = "LOAD_BALANCER"
	DrProtectionGroupMemberTypeNetworkLoadBalancer         DrProtectionGroupMemberTypeEnum = "NETWORK_LOAD_BALANCER"
	DrProtectionGroupMemberTypeFileSystem                  DrProtectionGroupMemberTypeEnum = "FILE_SYSTEM"
	DrProtectionGroupMemberTypeObjectStorageBucket         DrProtectionGroupMemberTypeEnum = "OBJECT_STORAGE_BUCKET"
)

var mappingDrProtectionGroupMemberTypeEnum = map[string]DrProtectionGroupMemberTypeEnum{
	"COMPUTE_INSTANCE":              DrProtectionGroupMemberTypeComputeInstance,
	"COMPUTE_INSTANCE_MOVABLE":      DrProtectionGroupMemberTypeComputeInstanceMovable,
	"COMPUTE_INSTANCE_NON_MOVABLE":  DrProtectionGroupMemberTypeComputeInstanceNonMovable,
	"VOLUME_GROUP":                  DrProtectionGroupMemberTypeVolumeGroup,
	"DATABASE":                      DrProtectionGroupMemberTypeDatabase,
	"AUTONOMOUS_DATABASE":           DrProtectionGroupMemberTypeAutonomousDatabase,
	"AUTONOMOUS_CONTAINER_DATABASE": DrProtectionGroupMemberTypeAutonomousContainerDatabase,
	"LOAD_BALANCER":                 DrProtectionGroupMemberTypeLoadBalancer,
	"NETWORK_LOAD_BALANCER":         DrProtectionGroupMemberTypeNetworkLoadBalancer,
	"FILE_SYSTEM":                   DrProtectionGroupMemberTypeFileSystem,
	"OBJECT_STORAGE_BUCKET":         DrProtectionGroupMemberTypeObjectStorageBucket,
}

var mappingDrProtectionGroupMemberTypeEnumLowerCase = map[string]DrProtectionGroupMemberTypeEnum{
	"compute_instance":              DrProtectionGroupMemberTypeComputeInstance,
	"compute_instance_movable":      DrProtectionGroupMemberTypeComputeInstanceMovable,
	"compute_instance_non_movable":  DrProtectionGroupMemberTypeComputeInstanceNonMovable,
	"volume_group":                  DrProtectionGroupMemberTypeVolumeGroup,
	"database":                      DrProtectionGroupMemberTypeDatabase,
	"autonomous_database":           DrProtectionGroupMemberTypeAutonomousDatabase,
	"autonomous_container_database": DrProtectionGroupMemberTypeAutonomousContainerDatabase,
	"load_balancer":                 DrProtectionGroupMemberTypeLoadBalancer,
	"network_load_balancer":         DrProtectionGroupMemberTypeNetworkLoadBalancer,
	"file_system":                   DrProtectionGroupMemberTypeFileSystem,
	"object_storage_bucket":         DrProtectionGroupMemberTypeObjectStorageBucket,
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
		"COMPUTE_INSTANCE_MOVABLE",
		"COMPUTE_INSTANCE_NON_MOVABLE",
		"VOLUME_GROUP",
		"DATABASE",
		"AUTONOMOUS_DATABASE",
		"AUTONOMOUS_CONTAINER_DATABASE",
		"LOAD_BALANCER",
		"NETWORK_LOAD_BALANCER",
		"FILE_SYSTEM",
		"OBJECT_STORAGE_BUCKET",
	}
}

// GetMappingDrProtectionGroupMemberTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrProtectionGroupMemberTypeEnum(val string) (DrProtectionGroupMemberTypeEnum, bool) {
	enum, ok := mappingDrProtectionGroupMemberTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
