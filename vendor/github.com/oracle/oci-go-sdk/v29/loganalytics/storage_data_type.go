// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

// StorageDataTypeEnum Enum with underlying type: string
type StorageDataTypeEnum string

// Set of constants representing the allowable values for StorageDataTypeEnum
const (
	StorageDataTypeLog    StorageDataTypeEnum = "LOG"
	StorageDataTypeLookup StorageDataTypeEnum = "LOOKUP"
)

var mappingStorageDataType = map[string]StorageDataTypeEnum{
	"LOG":    StorageDataTypeLog,
	"LOOKUP": StorageDataTypeLookup,
}

// GetStorageDataTypeEnumValues Enumerates the set of values for StorageDataTypeEnum
func GetStorageDataTypeEnumValues() []StorageDataTypeEnum {
	values := make([]StorageDataTypeEnum, 0)
	for _, v := range mappingStorageDataType {
		values = append(values, v)
	}
	return values
}
