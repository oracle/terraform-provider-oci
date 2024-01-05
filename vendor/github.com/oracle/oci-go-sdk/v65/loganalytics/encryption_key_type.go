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

// EncryptionKeyTypeEnum Enum with underlying type: string
type EncryptionKeyTypeEnum string

// Set of constants representing the allowable values for EncryptionKeyTypeEnum
const (
	EncryptionKeyTypeActiveData   EncryptionKeyTypeEnum = "ACTIVE_DATA"
	EncryptionKeyTypeArchivalData EncryptionKeyTypeEnum = "ARCHIVAL_DATA"
	EncryptionKeyTypeAll          EncryptionKeyTypeEnum = "ALL"
)

var mappingEncryptionKeyTypeEnum = map[string]EncryptionKeyTypeEnum{
	"ACTIVE_DATA":   EncryptionKeyTypeActiveData,
	"ARCHIVAL_DATA": EncryptionKeyTypeArchivalData,
	"ALL":           EncryptionKeyTypeAll,
}

var mappingEncryptionKeyTypeEnumLowerCase = map[string]EncryptionKeyTypeEnum{
	"active_data":   EncryptionKeyTypeActiveData,
	"archival_data": EncryptionKeyTypeArchivalData,
	"all":           EncryptionKeyTypeAll,
}

// GetEncryptionKeyTypeEnumValues Enumerates the set of values for EncryptionKeyTypeEnum
func GetEncryptionKeyTypeEnumValues() []EncryptionKeyTypeEnum {
	values := make([]EncryptionKeyTypeEnum, 0)
	for _, v := range mappingEncryptionKeyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetEncryptionKeyTypeEnumStringValues Enumerates the set of values in String for EncryptionKeyTypeEnum
func GetEncryptionKeyTypeEnumStringValues() []string {
	return []string{
		"ACTIVE_DATA",
		"ARCHIVAL_DATA",
		"ALL",
	}
}

// GetMappingEncryptionKeyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEncryptionKeyTypeEnum(val string) (EncryptionKeyTypeEnum, bool) {
	enum, ok := mappingEncryptionKeyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
