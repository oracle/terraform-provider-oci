// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Bastion API
//
// Oracle Cloud Infrastructure Bastion provides restricted and time-limited access to target resources that don't have public endpoints. Through the configuration of a bastion, you can let authorized users connect from specific IP addresses to target resources by way of Secure Shell (SSH) sessions hosted on the bastion.
//

package bastion

import (
	"strings"
)

// SessionLifecycleStateEnum Enum with underlying type: string
type SessionLifecycleStateEnum string

// Set of constants representing the allowable values for SessionLifecycleStateEnum
const (
	SessionLifecycleStateCreating SessionLifecycleStateEnum = "CREATING"
	SessionLifecycleStateActive   SessionLifecycleStateEnum = "ACTIVE"
	SessionLifecycleStateDeleting SessionLifecycleStateEnum = "DELETING"
	SessionLifecycleStateDeleted  SessionLifecycleStateEnum = "DELETED"
	SessionLifecycleStateFailed   SessionLifecycleStateEnum = "FAILED"
)

var mappingSessionLifecycleStateEnum = map[string]SessionLifecycleStateEnum{
	"CREATING": SessionLifecycleStateCreating,
	"ACTIVE":   SessionLifecycleStateActive,
	"DELETING": SessionLifecycleStateDeleting,
	"DELETED":  SessionLifecycleStateDeleted,
	"FAILED":   SessionLifecycleStateFailed,
}

// GetSessionLifecycleStateEnumValues Enumerates the set of values for SessionLifecycleStateEnum
func GetSessionLifecycleStateEnumValues() []SessionLifecycleStateEnum {
	values := make([]SessionLifecycleStateEnum, 0)
	for _, v := range mappingSessionLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSessionLifecycleStateEnumStringValues Enumerates the set of values in String for SessionLifecycleStateEnum
func GetSessionLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingSessionLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSessionLifecycleStateEnum(val string) (SessionLifecycleStateEnum, bool) {
	mappingSessionLifecycleStateEnumIgnoreCase := make(map[string]SessionLifecycleStateEnum)
	for k, v := range mappingSessionLifecycleStateEnum {
		mappingSessionLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSessionLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
