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

// OnPremConnectorLifecycleStateEnum Enum with underlying type: string
type OnPremConnectorLifecycleStateEnum string

// Set of constants representing the allowable values for OnPremConnectorLifecycleStateEnum
const (
	OnPremConnectorLifecycleStateCreating OnPremConnectorLifecycleStateEnum = "CREATING"
	OnPremConnectorLifecycleStateUpdating OnPremConnectorLifecycleStateEnum = "UPDATING"
	OnPremConnectorLifecycleStateActive   OnPremConnectorLifecycleStateEnum = "ACTIVE"
	OnPremConnectorLifecycleStateInactive OnPremConnectorLifecycleStateEnum = "INACTIVE"
	OnPremConnectorLifecycleStateDeleting OnPremConnectorLifecycleStateEnum = "DELETING"
	OnPremConnectorLifecycleStateDeleted  OnPremConnectorLifecycleStateEnum = "DELETED"
	OnPremConnectorLifecycleStateFailed   OnPremConnectorLifecycleStateEnum = "FAILED"
)

var mappingOnPremConnectorLifecycleStateEnum = map[string]OnPremConnectorLifecycleStateEnum{
	"CREATING": OnPremConnectorLifecycleStateCreating,
	"UPDATING": OnPremConnectorLifecycleStateUpdating,
	"ACTIVE":   OnPremConnectorLifecycleStateActive,
	"INACTIVE": OnPremConnectorLifecycleStateInactive,
	"DELETING": OnPremConnectorLifecycleStateDeleting,
	"DELETED":  OnPremConnectorLifecycleStateDeleted,
	"FAILED":   OnPremConnectorLifecycleStateFailed,
}

// GetOnPremConnectorLifecycleStateEnumValues Enumerates the set of values for OnPremConnectorLifecycleStateEnum
func GetOnPremConnectorLifecycleStateEnumValues() []OnPremConnectorLifecycleStateEnum {
	values := make([]OnPremConnectorLifecycleStateEnum, 0)
	for _, v := range mappingOnPremConnectorLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOnPremConnectorLifecycleStateEnumStringValues Enumerates the set of values in String for OnPremConnectorLifecycleStateEnum
func GetOnPremConnectorLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingOnPremConnectorLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOnPremConnectorLifecycleStateEnum(val string) (OnPremConnectorLifecycleStateEnum, bool) {
	mappingOnPremConnectorLifecycleStateEnumIgnoreCase := make(map[string]OnPremConnectorLifecycleStateEnum)
	for k, v := range mappingOnPremConnectorLifecycleStateEnum {
		mappingOnPremConnectorLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingOnPremConnectorLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
