// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"strings"
)

// StorageServerEnum Enum with underlying type: string
type StorageServerEnum string

// Set of constants representing the allowable values for StorageServerEnum
const (
	StorageServerStorage    StorageServerEnum = "STORAGE"
	StorageServerIops       StorageServerEnum = "IOPS"
	StorageServerThroughput StorageServerEnum = "THROUGHPUT"
)

var mappingStorageServerEnum = map[string]StorageServerEnum{
	"STORAGE":    StorageServerStorage,
	"IOPS":       StorageServerIops,
	"THROUGHPUT": StorageServerThroughput,
}

// GetStorageServerEnumValues Enumerates the set of values for StorageServerEnum
func GetStorageServerEnumValues() []StorageServerEnum {
	values := make([]StorageServerEnum, 0)
	for _, v := range mappingStorageServerEnum {
		values = append(values, v)
	}
	return values
}

// GetStorageServerEnumStringValues Enumerates the set of values in String for StorageServerEnum
func GetStorageServerEnumStringValues() []string {
	return []string{
		"STORAGE",
		"IOPS",
		"THROUGHPUT",
	}
}

// GetMappingStorageServerEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStorageServerEnum(val string) (StorageServerEnum, bool) {
	mappingStorageServerEnumIgnoreCase := make(map[string]StorageServerEnum)
	for k, v := range mappingStorageServerEnum {
		mappingStorageServerEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingStorageServerEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
