// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"strings"
)

// PayloadTypeEnum Enum with underlying type: string
type PayloadTypeEnum string

// Set of constants representing the allowable values for PayloadTypeEnum
const (
	PayloadTypeJson PayloadTypeEnum = "JSON"
	PayloadTypeGzip PayloadTypeEnum = "GZIP"
	PayloadTypeZip  PayloadTypeEnum = "ZIP"
)

var mappingPayloadTypeEnum = map[string]PayloadTypeEnum{
	"JSON": PayloadTypeJson,
	"GZIP": PayloadTypeGzip,
	"ZIP":  PayloadTypeZip,
}

var mappingPayloadTypeEnumLowerCase = map[string]PayloadTypeEnum{
	"json": PayloadTypeJson,
	"gzip": PayloadTypeGzip,
	"zip":  PayloadTypeZip,
}

// GetPayloadTypeEnumValues Enumerates the set of values for PayloadTypeEnum
func GetPayloadTypeEnumValues() []PayloadTypeEnum {
	values := make([]PayloadTypeEnum, 0)
	for _, v := range mappingPayloadTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPayloadTypeEnumStringValues Enumerates the set of values in String for PayloadTypeEnum
func GetPayloadTypeEnumStringValues() []string {
	return []string{
		"JSON",
		"GZIP",
		"ZIP",
	}
}

// GetMappingPayloadTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPayloadTypeEnum(val string) (PayloadTypeEnum, bool) {
	enum, ok := mappingPayloadTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
