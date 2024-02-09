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

// PhoneNumberTypeEnum Enum with underlying type: string
type PhoneNumberTypeEnum string

// Set of constants representing the allowable values for PhoneNumberTypeEnum
const (
	PhoneNumberTypeShortCode PhoneNumberTypeEnum = "SHORT_CODE"
	PhoneNumberTypeTollFree  PhoneNumberTypeEnum = "TOLL_FREE"
)

var mappingPhoneNumberTypeEnum = map[string]PhoneNumberTypeEnum{
	"SHORT_CODE": PhoneNumberTypeShortCode,
	"TOLL_FREE":  PhoneNumberTypeTollFree,
}

var mappingPhoneNumberTypeEnumLowerCase = map[string]PhoneNumberTypeEnum{
	"short_code": PhoneNumberTypeShortCode,
	"toll_free":  PhoneNumberTypeTollFree,
}

// GetPhoneNumberTypeEnumValues Enumerates the set of values for PhoneNumberTypeEnum
func GetPhoneNumberTypeEnumValues() []PhoneNumberTypeEnum {
	values := make([]PhoneNumberTypeEnum, 0)
	for _, v := range mappingPhoneNumberTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPhoneNumberTypeEnumStringValues Enumerates the set of values in String for PhoneNumberTypeEnum
func GetPhoneNumberTypeEnumStringValues() []string {
	return []string{
		"SHORT_CODE",
		"TOLL_FREE",
	}
}

// GetMappingPhoneNumberTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPhoneNumberTypeEnum(val string) (PhoneNumberTypeEnum, bool) {
	enum, ok := mappingPhoneNumberTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
