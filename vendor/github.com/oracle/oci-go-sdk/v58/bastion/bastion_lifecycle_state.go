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
	mappingBastionLifecycleStateEnumIgnoreCase := make(map[string]BastionLifecycleStateEnum)
	for k, v := range mappingBastionLifecycleStateEnum {
		mappingBastionLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingBastionLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
