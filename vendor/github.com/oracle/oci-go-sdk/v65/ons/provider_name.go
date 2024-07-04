// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Notifications API
//
// Use the Notifications API to broadcast messages to distributed components by topic, using a publish-subscribe pattern.
// For information about managing topics, subscriptions, and messages, see the Notifications documentation (https://docs.cloud.oracle.com/iaas/Content/Notification/home.htm).
//

package ons

import (
	"strings"
)

// ProviderNameEnum Enum with underlying type: string
type ProviderNameEnum string

// Set of constants representing the allowable values for ProviderNameEnum
const (
	ProviderNameEmail         ProviderNameEnum = "EMAIL"
	ProviderNameSms           ProviderNameEnum = "SMS"
	ProviderNameSlack         ProviderNameEnum = "SLACK"
	ProviderNameOceanInternal ProviderNameEnum = "OCEAN_INTERNAL"
)

var mappingProviderNameEnum = map[string]ProviderNameEnum{
	"EMAIL":          ProviderNameEmail,
	"SMS":            ProviderNameSms,
	"SLACK":          ProviderNameSlack,
	"OCEAN_INTERNAL": ProviderNameOceanInternal,
}

var mappingProviderNameEnumLowerCase = map[string]ProviderNameEnum{
	"email":          ProviderNameEmail,
	"sms":            ProviderNameSms,
	"slack":          ProviderNameSlack,
	"ocean_internal": ProviderNameOceanInternal,
}

// GetProviderNameEnumValues Enumerates the set of values for ProviderNameEnum
func GetProviderNameEnumValues() []ProviderNameEnum {
	values := make([]ProviderNameEnum, 0)
	for _, v := range mappingProviderNameEnum {
		values = append(values, v)
	}
	return values
}

// GetProviderNameEnumStringValues Enumerates the set of values in String for ProviderNameEnum
func GetProviderNameEnumStringValues() []string {
	return []string{
		"EMAIL",
		"SMS",
		"SLACK",
		"OCEAN_INTERNAL",
	}
}

// GetMappingProviderNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProviderNameEnum(val string) (ProviderNameEnum, bool) {
	enum, ok := mappingProviderNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
