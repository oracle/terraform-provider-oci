// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
// For more information, see Data Catalog (https://docs.oracle.com/iaas/data-catalog/home.htm).
//

package datacatalog

import (
	"strings"
)

// EventConfigStatusEnum Enum with underlying type: string
type EventConfigStatusEnum string

// Set of constants representing the allowable values for EventConfigStatusEnum
const (
	EventConfigStatusEnabled  EventConfigStatusEnum = "ENABLED"
	EventConfigStatusDisabled EventConfigStatusEnum = "DISABLED"
)

var mappingEventConfigStatusEnum = map[string]EventConfigStatusEnum{
	"ENABLED":  EventConfigStatusEnabled,
	"DISABLED": EventConfigStatusDisabled,
}

var mappingEventConfigStatusEnumLowerCase = map[string]EventConfigStatusEnum{
	"enabled":  EventConfigStatusEnabled,
	"disabled": EventConfigStatusDisabled,
}

// GetEventConfigStatusEnumValues Enumerates the set of values for EventConfigStatusEnum
func GetEventConfigStatusEnumValues() []EventConfigStatusEnum {
	values := make([]EventConfigStatusEnum, 0)
	for _, v := range mappingEventConfigStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetEventConfigStatusEnumStringValues Enumerates the set of values in String for EventConfigStatusEnum
func GetEventConfigStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingEventConfigStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEventConfigStatusEnum(val string) (EventConfigStatusEnum, bool) {
	enum, ok := mappingEventConfigStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
