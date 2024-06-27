// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Delegate Access Control API
//
// Oracle Delegate Access Control allows ExaCC and ExaCS customers to delegate management of their Exadata resources operators outside their tenancies.
// With Delegate Access Control, Support Providers can deliver managed services using comprehensive and robust tooling built on the OCI platform.
// Customers maintain control over who has access to the delegated resources in their tenancy and what actions can be taken.
// Enterprises managing resources across multiple tenants can use Delegate Access Control to streamline management tasks.
// Using logging service, customers can view a near real-time audit report of all actions performed by a Service Provider operator.
//

package delegateaccesscontrol

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
)

var mappingWorkRequestResourceMetadataKeyEnum = map[string]WorkRequestResourceMetadataKeyEnum{
	"VM_CLUSTER_ID":   WorkRequestResourceMetadataKeyVmClusterId,
	"HOSTNAMES":       WorkRequestResourceMetadataKeyHostnames,
	"COMMANDS":        WorkRequestResourceMetadataKeyCommands,
	"RESULT_LOCATION": WorkRequestResourceMetadataKeyResultLocation,
}

var mappingWorkRequestResourceMetadataKeyEnumLowerCase = map[string]WorkRequestResourceMetadataKeyEnum{
	"vm_cluster_id":   WorkRequestResourceMetadataKeyVmClusterId,
	"hostnames":       WorkRequestResourceMetadataKeyHostnames,
	"commands":        WorkRequestResourceMetadataKeyCommands,
	"result_location": WorkRequestResourceMetadataKeyResultLocation,
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
	}
}

// GetMappingWorkRequestResourceMetadataKeyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestResourceMetadataKeyEnum(val string) (WorkRequestResourceMetadataKeyEnum, bool) {
	enum, ok := mappingWorkRequestResourceMetadataKeyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
