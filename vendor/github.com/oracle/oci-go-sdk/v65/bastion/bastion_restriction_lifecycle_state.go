// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// BastionRestrictionLifecycleStateEnum Enum with underlying type: string
type BastionRestrictionLifecycleStateEnum string

// Set of constants representing the allowable values for BastionRestrictionLifecycleStateEnum
const (
	BastionRestrictionLifecycleStateCreating BastionRestrictionLifecycleStateEnum = "CREATING"
	BastionRestrictionLifecycleStateActive   BastionRestrictionLifecycleStateEnum = "ACTIVE"
	BastionRestrictionLifecycleStateDeleting BastionRestrictionLifecycleStateEnum = "DELETING"
	BastionRestrictionLifecycleStateDeleted  BastionRestrictionLifecycleStateEnum = "DELETED"
	BastionRestrictionLifecycleStateFailed   BastionRestrictionLifecycleStateEnum = "FAILED"
)

var mappingBastionRestrictionLifecycleStateEnum = map[string]BastionRestrictionLifecycleStateEnum{
	"CREATING": BastionRestrictionLifecycleStateCreating,
	"ACTIVE":   BastionRestrictionLifecycleStateActive,
	"DELETING": BastionRestrictionLifecycleStateDeleting,
	"DELETED":  BastionRestrictionLifecycleStateDeleted,
	"FAILED":   BastionRestrictionLifecycleStateFailed,
}

var mappingBastionRestrictionLifecycleStateEnumLowerCase = map[string]BastionRestrictionLifecycleStateEnum{
	"creating": BastionRestrictionLifecycleStateCreating,
	"active":   BastionRestrictionLifecycleStateActive,
	"deleting": BastionRestrictionLifecycleStateDeleting,
	"deleted":  BastionRestrictionLifecycleStateDeleted,
	"failed":   BastionRestrictionLifecycleStateFailed,
}

// GetBastionRestrictionLifecycleStateEnumValues Enumerates the set of values for BastionRestrictionLifecycleStateEnum
func GetBastionRestrictionLifecycleStateEnumValues() []BastionRestrictionLifecycleStateEnum {
	values := make([]BastionRestrictionLifecycleStateEnum, 0)
	for _, v := range mappingBastionRestrictionLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetBastionRestrictionLifecycleStateEnumStringValues Enumerates the set of values in String for BastionRestrictionLifecycleStateEnum
func GetBastionRestrictionLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingBastionRestrictionLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBastionRestrictionLifecycleStateEnum(val string) (BastionRestrictionLifecycleStateEnum, bool) {
	enum, ok := mappingBastionRestrictionLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
