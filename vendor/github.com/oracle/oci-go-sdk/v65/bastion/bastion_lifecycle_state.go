// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Bastion API
//
// Use the Bastion API to provide restricted and time-limited access to target resources that don't have public endpoints. Bastions let authorized users connect from specific IP addresses to target resources using Secure Shell (SSH) sessions. For more information, see the Bastion documentation (https://docs.cloud.oracle.com/iaas/Content/Bastion/home.htm).
//

package bastion

import (
	"strings"
)

// BastionLifecycleStateEnum Enum with underlying type: string
type BastionLifecycleStateEnum string

// Set of constants representing the allowable values for BastionLifecycleStateEnum
const (
	BastionLifecycleStateCreating BastionLifecycleStateEnum = "CREATING"
	BastionLifecycleStateUpdating BastionLifecycleStateEnum = "UPDATING"
	BastionLifecycleStateActive   BastionLifecycleStateEnum = "ACTIVE"
	BastionLifecycleStateDeleting BastionLifecycleStateEnum = "DELETING"
	BastionLifecycleStateDeleted  BastionLifecycleStateEnum = "DELETED"
	BastionLifecycleStateFailed   BastionLifecycleStateEnum = "FAILED"
)

var mappingBastionLifecycleStateEnum = map[string]BastionLifecycleStateEnum{
	"CREATING": BastionLifecycleStateCreating,
	"UPDATING": BastionLifecycleStateUpdating,
	"ACTIVE":   BastionLifecycleStateActive,
	"DELETING": BastionLifecycleStateDeleting,
	"DELETED":  BastionLifecycleStateDeleted,
	"FAILED":   BastionLifecycleStateFailed,
}

var mappingBastionLifecycleStateEnumLowerCase = map[string]BastionLifecycleStateEnum{
	"creating": BastionLifecycleStateCreating,
	"updating": BastionLifecycleStateUpdating,
	"active":   BastionLifecycleStateActive,
	"deleting": BastionLifecycleStateDeleting,
	"deleted":  BastionLifecycleStateDeleted,
	"failed":   BastionLifecycleStateFailed,
}

// GetBastionLifecycleStateEnumValues Enumerates the set of values for BastionLifecycleStateEnum
func GetBastionLifecycleStateEnumValues() []BastionLifecycleStateEnum {
	values := make([]BastionLifecycleStateEnum, 0)
	for _, v := range mappingBastionLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetBastionLifecycleStateEnumStringValues Enumerates the set of values in String for BastionLifecycleStateEnum
func GetBastionLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingBastionLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBastionLifecycleStateEnum(val string) (BastionLifecycleStateEnum, bool) {
	enum, ok := mappingBastionLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
