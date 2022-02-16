// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// AlertPolicyLifecycleStateEnum Enum with underlying type: string
type AlertPolicyLifecycleStateEnum string

// Set of constants representing the allowable values for AlertPolicyLifecycleStateEnum
const (
	AlertPolicyLifecycleStateCreating AlertPolicyLifecycleStateEnum = "CREATING"
	AlertPolicyLifecycleStateUpdating AlertPolicyLifecycleStateEnum = "UPDATING"
	AlertPolicyLifecycleStateActive   AlertPolicyLifecycleStateEnum = "ACTIVE"
	AlertPolicyLifecycleStateDeleting AlertPolicyLifecycleStateEnum = "DELETING"
	AlertPolicyLifecycleStateDeleted  AlertPolicyLifecycleStateEnum = "DELETED"
	AlertPolicyLifecycleStateFailed   AlertPolicyLifecycleStateEnum = "FAILED"
)

var mappingAlertPolicyLifecycleStateEnum = map[string]AlertPolicyLifecycleStateEnum{
	"CREATING": AlertPolicyLifecycleStateCreating,
	"UPDATING": AlertPolicyLifecycleStateUpdating,
	"ACTIVE":   AlertPolicyLifecycleStateActive,
	"DELETING": AlertPolicyLifecycleStateDeleting,
	"DELETED":  AlertPolicyLifecycleStateDeleted,
	"FAILED":   AlertPolicyLifecycleStateFailed,
}

// GetAlertPolicyLifecycleStateEnumValues Enumerates the set of values for AlertPolicyLifecycleStateEnum
func GetAlertPolicyLifecycleStateEnumValues() []AlertPolicyLifecycleStateEnum {
	values := make([]AlertPolicyLifecycleStateEnum, 0)
	for _, v := range mappingAlertPolicyLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAlertPolicyLifecycleStateEnumStringValues Enumerates the set of values in String for AlertPolicyLifecycleStateEnum
func GetAlertPolicyLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingAlertPolicyLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAlertPolicyLifecycleStateEnum(val string) (AlertPolicyLifecycleStateEnum, bool) {
	mappingAlertPolicyLifecycleStateEnumIgnoreCase := make(map[string]AlertPolicyLifecycleStateEnum)
	for k, v := range mappingAlertPolicyLifecycleStateEnum {
		mappingAlertPolicyLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingAlertPolicyLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
