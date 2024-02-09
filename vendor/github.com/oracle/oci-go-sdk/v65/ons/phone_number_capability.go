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

// PhoneNumberCapabilityEnum Enum with underlying type: string
type PhoneNumberCapabilityEnum string

// Set of constants representing the allowable values for PhoneNumberCapabilityEnum
const (
	PhoneNumberCapabilitySms PhoneNumberCapabilityEnum = "SMS"
)

var mappingPhoneNumberCapabilityEnum = map[string]PhoneNumberCapabilityEnum{
	"SMS": PhoneNumberCapabilitySms,
}

var mappingPhoneNumberCapabilityEnumLowerCase = map[string]PhoneNumberCapabilityEnum{
	"sms": PhoneNumberCapabilitySms,
}

// GetPhoneNumberCapabilityEnumValues Enumerates the set of values for PhoneNumberCapabilityEnum
func GetPhoneNumberCapabilityEnumValues() []PhoneNumberCapabilityEnum {
	values := make([]PhoneNumberCapabilityEnum, 0)
	for _, v := range mappingPhoneNumberCapabilityEnum {
		values = append(values, v)
	}
	return values
}

// GetPhoneNumberCapabilityEnumStringValues Enumerates the set of values in String for PhoneNumberCapabilityEnum
func GetPhoneNumberCapabilityEnumStringValues() []string {
	return []string{
		"SMS",
	}
}

// GetMappingPhoneNumberCapabilityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPhoneNumberCapabilityEnum(val string) (PhoneNumberCapabilityEnum, bool) {
	enum, ok := mappingPhoneNumberCapabilityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
