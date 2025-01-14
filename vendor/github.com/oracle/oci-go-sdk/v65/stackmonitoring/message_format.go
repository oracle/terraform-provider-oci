// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"strings"
)

// MessageFormatEnum Enum with underlying type: string
type MessageFormatEnum string

// Set of constants representing the allowable values for MessageFormatEnum
const (
	MessageFormatRaw          MessageFormatEnum = "RAW"
	MessageFormatPrettyJson   MessageFormatEnum = "PRETTY_JSON"
	MessageFormatOnsOptimized MessageFormatEnum = "ONS_OPTIMIZED"
)

var mappingMessageFormatEnum = map[string]MessageFormatEnum{
	"RAW":           MessageFormatRaw,
	"PRETTY_JSON":   MessageFormatPrettyJson,
	"ONS_OPTIMIZED": MessageFormatOnsOptimized,
}

var mappingMessageFormatEnumLowerCase = map[string]MessageFormatEnum{
	"raw":           MessageFormatRaw,
	"pretty_json":   MessageFormatPrettyJson,
	"ons_optimized": MessageFormatOnsOptimized,
}

// GetMessageFormatEnumValues Enumerates the set of values for MessageFormatEnum
func GetMessageFormatEnumValues() []MessageFormatEnum {
	values := make([]MessageFormatEnum, 0)
	for _, v := range mappingMessageFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetMessageFormatEnumStringValues Enumerates the set of values in String for MessageFormatEnum
func GetMessageFormatEnumStringValues() []string {
	return []string{
		"RAW",
		"PRETTY_JSON",
		"ONS_OPTIMIZED",
	}
}

// GetMappingMessageFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMessageFormatEnum(val string) (MessageFormatEnum, bool) {
	enum, ok := mappingMessageFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
