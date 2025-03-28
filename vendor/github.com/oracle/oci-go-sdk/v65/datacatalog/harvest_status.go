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

// HarvestStatusEnum Enum with underlying type: string
type HarvestStatusEnum string

// Set of constants representing the allowable values for HarvestStatusEnum
const (
	HarvestStatusComplete   HarvestStatusEnum = "COMPLETE"
	HarvestStatusError      HarvestStatusEnum = "ERROR"
	HarvestStatusInProgress HarvestStatusEnum = "IN_PROGRESS"
	HarvestStatusDeferred   HarvestStatusEnum = "DEFERRED"
)

var mappingHarvestStatusEnum = map[string]HarvestStatusEnum{
	"COMPLETE":    HarvestStatusComplete,
	"ERROR":       HarvestStatusError,
	"IN_PROGRESS": HarvestStatusInProgress,
	"DEFERRED":    HarvestStatusDeferred,
}

var mappingHarvestStatusEnumLowerCase = map[string]HarvestStatusEnum{
	"complete":    HarvestStatusComplete,
	"error":       HarvestStatusError,
	"in_progress": HarvestStatusInProgress,
	"deferred":    HarvestStatusDeferred,
}

// GetHarvestStatusEnumValues Enumerates the set of values for HarvestStatusEnum
func GetHarvestStatusEnumValues() []HarvestStatusEnum {
	values := make([]HarvestStatusEnum, 0)
	for _, v := range mappingHarvestStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetHarvestStatusEnumStringValues Enumerates the set of values in String for HarvestStatusEnum
func GetHarvestStatusEnumStringValues() []string {
	return []string{
		"COMPLETE",
		"ERROR",
		"IN_PROGRESS",
		"DEFERRED",
	}
}

// GetMappingHarvestStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHarvestStatusEnum(val string) (HarvestStatusEnum, bool) {
	enum, ok := mappingHarvestStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
