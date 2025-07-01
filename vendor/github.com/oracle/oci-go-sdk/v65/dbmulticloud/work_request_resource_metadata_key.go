// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Database MultiCloud Data plane Integration
//
// 1. Oracle Azure Connector Resource: This is for installing Azure Arc Server in ExaCS VM Cluster.
//   There are two way to install Azure Arc Server (Azure Identity) in ExaCS VMCluster.
//     a. Using Bearer Access Token or
//     b. By providing Authentication token
// 2. Oracle Azure Blob Container Resource: This is for to capture Azure Container details
//    and same will be used in multiple ExaCS VMCluster to mount the Azure Container.
// 3. Oracle Azure Blob Mount Resource: This is for to mount Azure Container in ExaCS VMCluster
//    using Oracle Azure Connector and Oracle Azure Blob Container Resource.
//

package dbmulticloud

import (
	"strings"
)

// WorkRequestResourceMetadataKeyEnum Enum with underlying type: string
type WorkRequestResourceMetadataKeyEnum string

// Set of constants representing the allowable values for WorkRequestResourceMetadataKeyEnum
const (
	WorkRequestResourceMetadataKeyVmClusterId    WorkRequestResourceMetadataKeyEnum = "VM_CLUSTER_ID"
	WorkRequestResourceMetadataKeyHostnames      WorkRequestResourceMetadataKeyEnum = "HOSTNAMES"
	WorkRequestResourceMetadataKeyCommands       WorkRequestResourceMetadataKeyEnum = "COMMANDS"
	WorkRequestResourceMetadataKeyResultLocation WorkRequestResourceMetadataKeyEnum = "RESULT_LOCATION"
	WorkRequestResourceMetadataKeyIsDryRun       WorkRequestResourceMetadataKeyEnum = "IS_DRY_RUN"
)

var mappingWorkRequestResourceMetadataKeyEnum = map[string]WorkRequestResourceMetadataKeyEnum{
	"VM_CLUSTER_ID":   WorkRequestResourceMetadataKeyVmClusterId,
	"HOSTNAMES":       WorkRequestResourceMetadataKeyHostnames,
	"COMMANDS":        WorkRequestResourceMetadataKeyCommands,
	"RESULT_LOCATION": WorkRequestResourceMetadataKeyResultLocation,
	"IS_DRY_RUN":      WorkRequestResourceMetadataKeyIsDryRun,
}

var mappingWorkRequestResourceMetadataKeyEnumLowerCase = map[string]WorkRequestResourceMetadataKeyEnum{
	"vm_cluster_id":   WorkRequestResourceMetadataKeyVmClusterId,
	"hostnames":       WorkRequestResourceMetadataKeyHostnames,
	"commands":        WorkRequestResourceMetadataKeyCommands,
	"result_location": WorkRequestResourceMetadataKeyResultLocation,
	"is_dry_run":      WorkRequestResourceMetadataKeyIsDryRun,
}

// GetWorkRequestResourceMetadataKeyEnumValues Enumerates the set of values for WorkRequestResourceMetadataKeyEnum
func GetWorkRequestResourceMetadataKeyEnumValues() []WorkRequestResourceMetadataKeyEnum {
	values := make([]WorkRequestResourceMetadataKeyEnum, 0)
	for _, v := range mappingWorkRequestResourceMetadataKeyEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestResourceMetadataKeyEnumStringValues Enumerates the set of values in String for WorkRequestResourceMetadataKeyEnum
func GetWorkRequestResourceMetadataKeyEnumStringValues() []string {
	return []string{
		"VM_CLUSTER_ID",
		"HOSTNAMES",
		"COMMANDS",
		"RESULT_LOCATION",
		"IS_DRY_RUN",
	}
}

// GetMappingWorkRequestResourceMetadataKeyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestResourceMetadataKeyEnum(val string) (WorkRequestResourceMetadataKeyEnum, bool) {
	enum, ok := mappingWorkRequestResourceMetadataKeyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
