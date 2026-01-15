// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// HandlerTypeEnum Enum with underlying type: string
type HandlerTypeEnum string

// Set of constants representing the allowable values for HandlerTypeEnum
const (
	HandlerTypeTelegraf HandlerTypeEnum = "TELEGRAF"
	HandlerTypeCollectd HandlerTypeEnum = "COLLECTD"
)

var mappingHandlerTypeEnum = map[string]HandlerTypeEnum{
	"TELEGRAF": HandlerTypeTelegraf,
	"COLLECTD": HandlerTypeCollectd,
}

var mappingHandlerTypeEnumLowerCase = map[string]HandlerTypeEnum{
	"telegraf": HandlerTypeTelegraf,
	"collectd": HandlerTypeCollectd,
}

// GetHandlerTypeEnumValues Enumerates the set of values for HandlerTypeEnum
func GetHandlerTypeEnumValues() []HandlerTypeEnum {
	values := make([]HandlerTypeEnum, 0)
	for _, v := range mappingHandlerTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetHandlerTypeEnumStringValues Enumerates the set of values in String for HandlerTypeEnum
func GetHandlerTypeEnumStringValues() []string {
	return []string{
		"TELEGRAF",
		"COLLECTD",
	}
}

// GetMappingHandlerTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHandlerTypeEnum(val string) (HandlerTypeEnum, bool) {
	enum, ok := mappingHandlerTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
