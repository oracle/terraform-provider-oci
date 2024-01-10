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

// ChannelCategoryEnum Enum with underlying type: string
type ChannelCategoryEnum string

// Set of constants representing the allowable values for ChannelCategoryEnum
const (
	ChannelCategoryAgent       ChannelCategoryEnum = "AGENT"
	ChannelCategoryApplication ChannelCategoryEnum = "APPLICATION"
	ChannelCategoryBot         ChannelCategoryEnum = "BOT"
	ChannelCategoryBotAsAgent  ChannelCategoryEnum = "BOT_AS_AGENT"
	ChannelCategorySystem      ChannelCategoryEnum = "SYSTEM"
	ChannelCategoryEvent       ChannelCategoryEnum = "EVENT"
)

var mappingChannelCategoryEnum = map[string]ChannelCategoryEnum{
	"AGENT":        ChannelCategoryAgent,
	"APPLICATION":  ChannelCategoryApplication,
	"BOT":          ChannelCategoryBot,
	"BOT_AS_AGENT": ChannelCategoryBotAsAgent,
	"SYSTEM":       ChannelCategorySystem,
	"EVENT":        ChannelCategoryEvent,
}

var mappingChannelCategoryEnumLowerCase = map[string]ChannelCategoryEnum{
	"agent":        ChannelCategoryAgent,
	"application":  ChannelCategoryApplication,
	"bot":          ChannelCategoryBot,
	"bot_as_agent": ChannelCategoryBotAsAgent,
	"system":       ChannelCategorySystem,
	"event":        ChannelCategoryEvent,
}

// GetChannelCategoryEnumValues Enumerates the set of values for ChannelCategoryEnum
func GetChannelCategoryEnumValues() []ChannelCategoryEnum {
	values := make([]ChannelCategoryEnum, 0)
	for _, v := range mappingChannelCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetChannelCategoryEnumStringValues Enumerates the set of values in String for ChannelCategoryEnum
func GetChannelCategoryEnumStringValues() []string {
	return []string{
		"AGENT",
		"APPLICATION",
		"BOT",
		"BOT_AS_AGENT",
		"SYSTEM",
		"EVENT",
	}
}

// GetMappingChannelCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingChannelCategoryEnum(val string) (ChannelCategoryEnum, bool) {
	enum, ok := mappingChannelCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
