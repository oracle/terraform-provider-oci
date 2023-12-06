// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// StorageTierEnum Enum with underlying type: string
type StorageTierEnum string

// Set of constants representing the allowable values for StorageTierEnum
const (
	StorageTierStandard         StorageTierEnum = "STANDARD"
	StorageTierInfrequentaccess StorageTierEnum = "INFREQUENTACCESS"
	StorageTierArchive          StorageTierEnum = "ARCHIVE"
)

var mappingStorageTierEnum = map[string]StorageTierEnum{
	"STANDARD":         StorageTierStandard,
	"INFREQUENTACCESS": StorageTierInfrequentaccess,
	"ARCHIVE":          StorageTierArchive,
}

var mappingStorageTierEnumLowerCase = map[string]StorageTierEnum{
	"standard":         StorageTierStandard,
	"infrequentaccess": StorageTierInfrequentaccess,
	"archive":          StorageTierArchive,
}

// GetStorageTierEnumValues Enumerates the set of values for StorageTierEnum
func GetStorageTierEnumValues() []StorageTierEnum {
	values := make([]StorageTierEnum, 0)
	for _, v := range mappingStorageTierEnum {
		values = append(values, v)
	}
	return values
}

// GetStorageTierEnumStringValues Enumerates the set of values in String for StorageTierEnum
func GetStorageTierEnumStringValues() []string {
	return []string{
		"STANDARD",
		"INFREQUENTACCESS",
		"ARCHIVE",
	}
}

// GetMappingStorageTierEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStorageTierEnum(val string) (StorageTierEnum, bool) {
	enum, ok := mappingStorageTierEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
