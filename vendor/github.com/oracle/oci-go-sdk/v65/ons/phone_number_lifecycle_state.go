// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Notifications API
//
// Use the Notifications API to broadcast messages to distributed components by topic, using a publish-subscribe pattern.
// For information about managing topics, subscriptions, and messages, see Notifications Overview (https://docs.cloud.oracle.com/iaas/Content/Notification/Concepts/notificationoverview.htm).
//

package ons

import (
	"strings"
)

// PhoneNumberLifecycleStateEnum Enum with underlying type: string
type PhoneNumberLifecycleStateEnum string

// Set of constants representing the allowable values for PhoneNumberLifecycleStateEnum
const (
	PhoneNumberLifecycleStateCreating       PhoneNumberLifecycleStateEnum = "CREATING"
	PhoneNumberLifecycleStateActive         PhoneNumberLifecycleStateEnum = "ACTIVE"
	PhoneNumberLifecycleStateUpdating       PhoneNumberLifecycleStateEnum = "UPDATING"
	PhoneNumberLifecycleStateDeleting       PhoneNumberLifecycleStateEnum = "DELETING"
	PhoneNumberLifecycleStateDeleted        PhoneNumberLifecycleStateEnum = "DELETED"
	PhoneNumberLifecycleStateFailed         PhoneNumberLifecycleStateEnum = "FAILED"
	PhoneNumberLifecycleStateNeedsAttention PhoneNumberLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingPhoneNumberLifecycleStateEnum = map[string]PhoneNumberLifecycleStateEnum{
	"CREATING":        PhoneNumberLifecycleStateCreating,
	"ACTIVE":          PhoneNumberLifecycleStateActive,
	"UPDATING":        PhoneNumberLifecycleStateUpdating,
	"DELETING":        PhoneNumberLifecycleStateDeleting,
	"DELETED":         PhoneNumberLifecycleStateDeleted,
	"FAILED":          PhoneNumberLifecycleStateFailed,
	"NEEDS_ATTENTION": PhoneNumberLifecycleStateNeedsAttention,
}

var mappingPhoneNumberLifecycleStateEnumLowerCase = map[string]PhoneNumberLifecycleStateEnum{
	"creating":        PhoneNumberLifecycleStateCreating,
	"active":          PhoneNumberLifecycleStateActive,
	"updating":        PhoneNumberLifecycleStateUpdating,
	"deleting":        PhoneNumberLifecycleStateDeleting,
	"deleted":         PhoneNumberLifecycleStateDeleted,
	"failed":          PhoneNumberLifecycleStateFailed,
	"needs_attention": PhoneNumberLifecycleStateNeedsAttention,
}

// GetPhoneNumberLifecycleStateEnumValues Enumerates the set of values for PhoneNumberLifecycleStateEnum
func GetPhoneNumberLifecycleStateEnumValues() []PhoneNumberLifecycleStateEnum {
	values := make([]PhoneNumberLifecycleStateEnum, 0)
	for _, v := range mappingPhoneNumberLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPhoneNumberLifecycleStateEnumStringValues Enumerates the set of values in String for PhoneNumberLifecycleStateEnum
func GetPhoneNumberLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingPhoneNumberLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPhoneNumberLifecycleStateEnum(val string) (PhoneNumberLifecycleStateEnum, bool) {
	enum, ok := mappingPhoneNumberLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
