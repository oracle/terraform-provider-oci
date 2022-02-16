// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Advisor API
//
// Use the Cloud Advisor API to find potential inefficiencies in your tenancy and address them.
// Cloud Advisor can help you save money, improve performance, strengthen system resilience, and improve security.
// For more information, see Cloud Advisor (https://docs.cloud.oracle.com/Content/CloudAdvisor/Concepts/cloudadvisoroverview.htm).
//

package optimizer

import (
	"strings"
)

// TagValueTypeEnum Enum with underlying type: string
type TagValueTypeEnum string

// Set of constants representing the allowable values for TagValueTypeEnum
const (
	TagValueTypeValue TagValueTypeEnum = "VALUE"
	TagValueTypeAny   TagValueTypeEnum = "ANY"
)

var mappingTagValueTypeEnum = map[string]TagValueTypeEnum{
	"VALUE": TagValueTypeValue,
	"ANY":   TagValueTypeAny,
}

// GetTagValueTypeEnumValues Enumerates the set of values for TagValueTypeEnum
func GetTagValueTypeEnumValues() []TagValueTypeEnum {
	values := make([]TagValueTypeEnum, 0)
	for _, v := range mappingTagValueTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTagValueTypeEnumStringValues Enumerates the set of values in String for TagValueTypeEnum
func GetTagValueTypeEnumStringValues() []string {
	return []string{
		"VALUE",
		"ANY",
	}
}

// GetMappingTagValueTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTagValueTypeEnum(val string) (TagValueTypeEnum, bool) {
	mappingTagValueTypeEnumIgnoreCase := make(map[string]TagValueTypeEnum)
	for k, v := range mappingTagValueTypeEnum {
		mappingTagValueTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingTagValueTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
