// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Container Engine for Kubernetes API
//
// API for the Container Engine for Kubernetes service. Use this API to build, deploy,
// and manage cloud-native applications. For more information, see
// Overview of Container Engine for Kubernetes (https://docs.cloud.oracle.com/iaas/Content/ContEng/Concepts/contengoverview.htm).
//

package containerengine

// SortOrderEnum Enum with underlying type: string
type SortOrderEnum string

// Set of constants representing the allowable values for SortOrderEnum
const (
	SortOrderAsc  SortOrderEnum = "ASC"
	SortOrderDesc SortOrderEnum = "DESC"
)

var mappingSortOrder = map[string]SortOrderEnum{
	"ASC":  SortOrderAsc,
	"DESC": SortOrderDesc,
}

// GetSortOrderEnumValues Enumerates the set of values for SortOrderEnum
func GetSortOrderEnumValues() []SortOrderEnum {
	values := make([]SortOrderEnum, 0)
	for _, v := range mappingSortOrder {
		values = append(values, v)
	}
	return values
}
