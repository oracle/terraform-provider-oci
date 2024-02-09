// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// PhoneNumberLifecycleDetailsEnum Enum with underlying type: string
type PhoneNumberLifecycleDetailsEnum string

// Set of constants representing the allowable values for PhoneNumberLifecycleDetailsEnum
const (
	PhoneNumberLifecycleDetailsPendingApproval PhoneNumberLifecycleDetailsEnum = "PENDING_APPROVAL"
	PhoneNumberLifecycleDetailsApproved        PhoneNumberLifecycleDetailsEnum = "APPROVED"
	PhoneNumberLifecycleDetailsRejected        PhoneNumberLifecycleDetailsEnum = "REJECTED"
	PhoneNumberLifecycleDetailsInactive        PhoneNumberLifecycleDetailsEnum = "INACTIVE"
)

var mappingPhoneNumberLifecycleDetailsEnum = map[string]PhoneNumberLifecycleDetailsEnum{
	"PENDING_APPROVAL": PhoneNumberLifecycleDetailsPendingApproval,
	"APPROVED":         PhoneNumberLifecycleDetailsApproved,
	"REJECTED":         PhoneNumberLifecycleDetailsRejected,
	"INACTIVE":         PhoneNumberLifecycleDetailsInactive,
}

var mappingPhoneNumberLifecycleDetailsEnumLowerCase = map[string]PhoneNumberLifecycleDetailsEnum{
	"pending_approval": PhoneNumberLifecycleDetailsPendingApproval,
	"approved":         PhoneNumberLifecycleDetailsApproved,
	"rejected":         PhoneNumberLifecycleDetailsRejected,
	"inactive":         PhoneNumberLifecycleDetailsInactive,
}

// GetPhoneNumberLifecycleDetailsEnumValues Enumerates the set of values for PhoneNumberLifecycleDetailsEnum
func GetPhoneNumberLifecycleDetailsEnumValues() []PhoneNumberLifecycleDetailsEnum {
	values := make([]PhoneNumberLifecycleDetailsEnum, 0)
	for _, v := range mappingPhoneNumberLifecycleDetailsEnum {
		values = append(values, v)
	}
	return values
}

// GetPhoneNumberLifecycleDetailsEnumStringValues Enumerates the set of values in String for PhoneNumberLifecycleDetailsEnum
func GetPhoneNumberLifecycleDetailsEnumStringValues() []string {
	return []string{
		"PENDING_APPROVAL",
		"APPROVED",
		"REJECTED",
		"INACTIVE",
	}
}

// GetMappingPhoneNumberLifecycleDetailsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPhoneNumberLifecycleDetailsEnum(val string) (PhoneNumberLifecycleDetailsEnum, bool) {
	enum, ok := mappingPhoneNumberLifecycleDetailsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
