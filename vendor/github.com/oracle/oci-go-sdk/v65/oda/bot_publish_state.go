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

// BotPublishStateEnum Enum with underlying type: string
type BotPublishStateEnum string

// Set of constants representing the allowable values for BotPublishStateEnum
const (
	BotPublishStatePublished BotPublishStateEnum = "PUBLISHED"
	BotPublishStateDraft     BotPublishStateEnum = "DRAFT"
)

var mappingBotPublishStateEnum = map[string]BotPublishStateEnum{
	"PUBLISHED": BotPublishStatePublished,
	"DRAFT":     BotPublishStateDraft,
}

var mappingBotPublishStateEnumLowerCase = map[string]BotPublishStateEnum{
	"published": BotPublishStatePublished,
	"draft":     BotPublishStateDraft,
}

// GetBotPublishStateEnumValues Enumerates the set of values for BotPublishStateEnum
func GetBotPublishStateEnumValues() []BotPublishStateEnum {
	values := make([]BotPublishStateEnum, 0)
	for _, v := range mappingBotPublishStateEnum {
		values = append(values, v)
	}
	return values
}

// GetBotPublishStateEnumStringValues Enumerates the set of values in String for BotPublishStateEnum
func GetBotPublishStateEnumStringValues() []string {
	return []string{
		"PUBLISHED",
		"DRAFT",
	}
}

// GetMappingBotPublishStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBotPublishStateEnum(val string) (BotPublishStateEnum, bool) {
	enum, ok := mappingBotPublishStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
