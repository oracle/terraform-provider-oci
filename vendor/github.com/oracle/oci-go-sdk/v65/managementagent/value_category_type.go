// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// Use the Management Agent API to manage your infrastructure's management agents, including their plugins and install keys.
// For more information, see Management Agent (https://docs.oracle.com/iaas/management-agents/index.html).
//

package managementagent

import (
	"strings"
)

// ValueCategoryTypeEnum Enum with underlying type: string
type ValueCategoryTypeEnum string

// Set of constants representing the allowable values for ValueCategoryTypeEnum
const (
	ValueCategoryTypeClearText        ValueCategoryTypeEnum = "CLEAR_TEXT"
	ValueCategoryTypeSecretIdentifier ValueCategoryTypeEnum = "SECRET_IDENTIFIER"
	ValueCategoryTypeAdbIdentifier    ValueCategoryTypeEnum = "ADB_IDENTIFIER"
	ValueCategoryTypeAllowedValue     ValueCategoryTypeEnum = "ALLOWED_VALUE"
)

var mappingValueCategoryTypeEnum = map[string]ValueCategoryTypeEnum{
	"CLEAR_TEXT":        ValueCategoryTypeClearText,
	"SECRET_IDENTIFIER": ValueCategoryTypeSecretIdentifier,
	"ADB_IDENTIFIER":    ValueCategoryTypeAdbIdentifier,
	"ALLOWED_VALUE":     ValueCategoryTypeAllowedValue,
}

var mappingValueCategoryTypeEnumLowerCase = map[string]ValueCategoryTypeEnum{
	"clear_text":        ValueCategoryTypeClearText,
	"secret_identifier": ValueCategoryTypeSecretIdentifier,
	"adb_identifier":    ValueCategoryTypeAdbIdentifier,
	"allowed_value":     ValueCategoryTypeAllowedValue,
}

// GetValueCategoryTypeEnumValues Enumerates the set of values for ValueCategoryTypeEnum
func GetValueCategoryTypeEnumValues() []ValueCategoryTypeEnum {
	values := make([]ValueCategoryTypeEnum, 0)
	for _, v := range mappingValueCategoryTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetValueCategoryTypeEnumStringValues Enumerates the set of values in String for ValueCategoryTypeEnum
func GetValueCategoryTypeEnumStringValues() []string {
	return []string{
		"CLEAR_TEXT",
		"SECRET_IDENTIFIER",
		"ADB_IDENTIFIER",
		"ALLOWED_VALUE",
	}
}

// GetMappingValueCategoryTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingValueCategoryTypeEnum(val string) (ValueCategoryTypeEnum, bool) {
	enum, ok := mappingValueCategoryTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
