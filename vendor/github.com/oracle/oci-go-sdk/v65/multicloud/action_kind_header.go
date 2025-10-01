// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Multicloud API
//
// Use the Oracle Multicloud API to retrieve resource anchors and network anchors, and the metadata mappings related a Cloud Service Provider. For more information, see <link to docs>.
//

package multicloud

import (
	"strings"
)

// ActionKindHeaderEnum Enum with underlying type: string
type ActionKindHeaderEnum string

// Set of constants representing the allowable values for ActionKindHeaderEnum
const (
	ActionKindHeaderCreate ActionKindHeaderEnum = "create"
	ActionKindHeaderUpdate ActionKindHeaderEnum = "update"
)

var mappingActionKindHeaderEnum = map[string]ActionKindHeaderEnum{
	"create": ActionKindHeaderCreate,
	"update": ActionKindHeaderUpdate,
}

var mappingActionKindHeaderEnumLowerCase = map[string]ActionKindHeaderEnum{
	"create": ActionKindHeaderCreate,
	"update": ActionKindHeaderUpdate,
}

// GetActionKindHeaderEnumValues Enumerates the set of values for ActionKindHeaderEnum
func GetActionKindHeaderEnumValues() []ActionKindHeaderEnum {
	values := make([]ActionKindHeaderEnum, 0)
	for _, v := range mappingActionKindHeaderEnum {
		values = append(values, v)
	}
	return values
}

// GetActionKindHeaderEnumStringValues Enumerates the set of values in String for ActionKindHeaderEnum
func GetActionKindHeaderEnumStringValues() []string {
	return []string{
		"create",
		"update",
	}
}

// GetMappingActionKindHeaderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingActionKindHeaderEnum(val string) (ActionKindHeaderEnum, bool) {
	enum, ok := mappingActionKindHeaderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
