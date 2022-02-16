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

// DiskGroupEnum Enum with underlying type: string
type DiskGroupEnum string

// Set of constants representing the allowable values for DiskGroupEnum
const (
	DiskGroupStorage DiskGroupEnum = "STORAGE"
)

var mappingDiskGroupEnum = map[string]DiskGroupEnum{
	"STORAGE": DiskGroupStorage,
}

// GetDiskGroupEnumValues Enumerates the set of values for DiskGroupEnum
func GetDiskGroupEnumValues() []DiskGroupEnum {
	values := make([]DiskGroupEnum, 0)
	for _, v := range mappingDiskGroupEnum {
		values = append(values, v)
	}
	return values
}

// GetDiskGroupEnumStringValues Enumerates the set of values in String for DiskGroupEnum
func GetDiskGroupEnumStringValues() []string {
	return []string{
		"STORAGE",
	}
}

// GetMappingDiskGroupEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDiskGroupEnum(val string) (DiskGroupEnum, bool) {
	mappingDiskGroupEnumIgnoreCase := make(map[string]DiskGroupEnum)
	for k, v := range mappingDiskGroupEnum {
		mappingDiskGroupEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingDiskGroupEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
