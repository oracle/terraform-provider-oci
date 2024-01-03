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

// BotMultilingualModeEnum Enum with underlying type: string
type BotMultilingualModeEnum string

// Set of constants representing the allowable values for BotMultilingualModeEnum
const (
	BotMultilingualModeNative      BotMultilingualModeEnum = "NATIVE"
	BotMultilingualModeTranslation BotMultilingualModeEnum = "TRANSLATION"
)

var mappingBotMultilingualModeEnum = map[string]BotMultilingualModeEnum{
	"NATIVE":      BotMultilingualModeNative,
	"TRANSLATION": BotMultilingualModeTranslation,
}

var mappingBotMultilingualModeEnumLowerCase = map[string]BotMultilingualModeEnum{
	"native":      BotMultilingualModeNative,
	"translation": BotMultilingualModeTranslation,
}

// GetBotMultilingualModeEnumValues Enumerates the set of values for BotMultilingualModeEnum
func GetBotMultilingualModeEnumValues() []BotMultilingualModeEnum {
	values := make([]BotMultilingualModeEnum, 0)
	for _, v := range mappingBotMultilingualModeEnum {
		values = append(values, v)
	}
	return values
}

// GetBotMultilingualModeEnumStringValues Enumerates the set of values in String for BotMultilingualModeEnum
func GetBotMultilingualModeEnumStringValues() []string {
	return []string{
		"NATIVE",
		"TRANSLATION",
	}
}

// GetMappingBotMultilingualModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBotMultilingualModeEnum(val string) (BotMultilingualModeEnum, bool) {
	enum, ok := mappingBotMultilingualModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
