// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
//

package datacatalog

// HarvestStatusEnum Enum with underlying type: string
type HarvestStatusEnum string

// Set of constants representing the allowable values for HarvestStatusEnum
const (
	HarvestStatusComplete   HarvestStatusEnum = "COMPLETE"
	HarvestStatusError      HarvestStatusEnum = "ERROR"
	HarvestStatusInProgress HarvestStatusEnum = "IN_PROGRESS"
	HarvestStatusDeferred   HarvestStatusEnum = "DEFERRED"
)

var mappingHarvestStatus = map[string]HarvestStatusEnum{
	"COMPLETE":    HarvestStatusComplete,
	"ERROR":       HarvestStatusError,
	"IN_PROGRESS": HarvestStatusInProgress,
	"DEFERRED":    HarvestStatusDeferred,
}

// GetHarvestStatusEnumValues Enumerates the set of values for HarvestStatusEnum
func GetHarvestStatusEnumValues() []HarvestStatusEnum {
	values := make([]HarvestStatusEnum, 0)
	for _, v := range mappingHarvestStatus {
		values = append(values, v)
	}
	return values
}
