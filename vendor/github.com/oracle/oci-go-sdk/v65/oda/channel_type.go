// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Digital Assistant Service Instance API
//
// API to create and maintain Oracle Digital Assistant service instances.
//

package oda

import (
	"strings"
)

// ChannelTypeEnum Enum with underlying type: string
type ChannelTypeEnum string

// Set of constants representing the allowable values for ChannelTypeEnum
const (
	ChannelTypeAndroid      ChannelTypeEnum = "ANDROID"
	ChannelTypeAppevent     ChannelTypeEnum = "APPEVENT"
	ChannelTypeApplication  ChannelTypeEnum = "APPLICATION"
	ChannelTypeCortana      ChannelTypeEnum = "CORTANA"
	ChannelTypeFacebook     ChannelTypeEnum = "FACEBOOK"
	ChannelTypeIos          ChannelTypeEnum = "IOS"
	ChannelTypeMsteams      ChannelTypeEnum = "MSTEAMS"
	ChannelTypeOss          ChannelTypeEnum = "OSS"
	ChannelTypeOsvc         ChannelTypeEnum = "OSVC"
	ChannelTypeServicecloud ChannelTypeEnum = "SERVICECLOUD"
	ChannelTypeSlack        ChannelTypeEnum = "SLACK"
	ChannelTypeTest         ChannelTypeEnum = "TEST"
	ChannelTypeTwilio       ChannelTypeEnum = "TWILIO"
	ChannelTypeWeb          ChannelTypeEnum = "WEB"
	ChannelTypeWebhook      ChannelTypeEnum = "WEBHOOK"
)

var mappingChannelTypeEnum = map[string]ChannelTypeEnum{
	"ANDROID":      ChannelTypeAndroid,
	"APPEVENT":     ChannelTypeAppevent,
	"APPLICATION":  ChannelTypeApplication,
	"CORTANA":      ChannelTypeCortana,
	"FACEBOOK":     ChannelTypeFacebook,
	"IOS":          ChannelTypeIos,
	"MSTEAMS":      ChannelTypeMsteams,
	"OSS":          ChannelTypeOss,
	"OSVC":         ChannelTypeOsvc,
	"SERVICECLOUD": ChannelTypeServicecloud,
	"SLACK":        ChannelTypeSlack,
	"TEST":         ChannelTypeTest,
	"TWILIO":       ChannelTypeTwilio,
	"WEB":          ChannelTypeWeb,
	"WEBHOOK":      ChannelTypeWebhook,
}

var mappingChannelTypeEnumLowerCase = map[string]ChannelTypeEnum{
	"android":      ChannelTypeAndroid,
	"appevent":     ChannelTypeAppevent,
	"application":  ChannelTypeApplication,
	"cortana":      ChannelTypeCortana,
	"facebook":     ChannelTypeFacebook,
	"ios":          ChannelTypeIos,
	"msteams":      ChannelTypeMsteams,
	"oss":          ChannelTypeOss,
	"osvc":         ChannelTypeOsvc,
	"servicecloud": ChannelTypeServicecloud,
	"slack":        ChannelTypeSlack,
	"test":         ChannelTypeTest,
	"twilio":       ChannelTypeTwilio,
	"web":          ChannelTypeWeb,
	"webhook":      ChannelTypeWebhook,
}

// GetChannelTypeEnumValues Enumerates the set of values for ChannelTypeEnum
func GetChannelTypeEnumValues() []ChannelTypeEnum {
	values := make([]ChannelTypeEnum, 0)
	for _, v := range mappingChannelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetChannelTypeEnumStringValues Enumerates the set of values in String for ChannelTypeEnum
func GetChannelTypeEnumStringValues() []string {
	return []string{
		"ANDROID",
		"APPEVENT",
		"APPLICATION",
		"CORTANA",
		"FACEBOOK",
		"IOS",
		"MSTEAMS",
		"OSS",
		"OSVC",
		"SERVICECLOUD",
		"SLACK",
		"TEST",
		"TWILIO",
		"WEB",
		"WEBHOOK",
	}
}

// GetMappingChannelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingChannelTypeEnum(val string) (ChannelTypeEnum, bool) {
	enum, ok := mappingChannelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
