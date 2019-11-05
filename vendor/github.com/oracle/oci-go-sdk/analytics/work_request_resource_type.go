// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

// WorkRequestResourceTypeEnum Enum with underlying type: string
type WorkRequestResourceTypeEnum string

// Set of constants representing the allowable values for WorkRequestResourceTypeEnum
const (
	WorkRequestResourceTypeAnalyticsInstance WorkRequestResourceTypeEnum = "ANALYTICS_INSTANCE"
)

var mappingWorkRequestResourceType = map[string]WorkRequestResourceTypeEnum{
	"ANALYTICS_INSTANCE": WorkRequestResourceTypeAnalyticsInstance,
}

// GetWorkRequestResourceTypeEnumValues Enumerates the set of values for WorkRequestResourceTypeEnum
func GetWorkRequestResourceTypeEnumValues() []WorkRequestResourceTypeEnum {
	values := make([]WorkRequestResourceTypeEnum, 0)
	for _, v := range mappingWorkRequestResourceType {
		values = append(values, v)
	}
	return values
}
