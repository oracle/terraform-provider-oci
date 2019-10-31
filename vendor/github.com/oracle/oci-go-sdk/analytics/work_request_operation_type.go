// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

// WorkRequestOperationTypeEnum Enum with underlying type: string
type WorkRequestOperationTypeEnum string

// Set of constants representing the allowable values for WorkRequestOperationTypeEnum
const (
	WorkRequestOperationTypeCreateAnalyticsInstance            WorkRequestOperationTypeEnum = "CREATE_ANALYTICS_INSTANCE"
	WorkRequestOperationTypeDeleteAnalyticsInstance            WorkRequestOperationTypeEnum = "DELETE_ANALYTICS_INSTANCE"
	WorkRequestOperationTypeStartAnalyticsInstance             WorkRequestOperationTypeEnum = "START_ANALYTICS_INSTANCE"
	WorkRequestOperationTypeStopAnalyticsInstance              WorkRequestOperationTypeEnum = "STOP_ANALYTICS_INSTANCE"
	WorkRequestOperationTypeScaleAnalyticsInstance             WorkRequestOperationTypeEnum = "SCALE_ANALYTICS_INSTANCE"
	WorkRequestOperationTypeChangeAnalyticsInstanceCompartment WorkRequestOperationTypeEnum = "CHANGE_ANALYTICS_INSTANCE_COMPARTMENT"
)

var mappingWorkRequestOperationType = map[string]WorkRequestOperationTypeEnum{
	"CREATE_ANALYTICS_INSTANCE":             WorkRequestOperationTypeCreateAnalyticsInstance,
	"DELETE_ANALYTICS_INSTANCE":             WorkRequestOperationTypeDeleteAnalyticsInstance,
	"START_ANALYTICS_INSTANCE":              WorkRequestOperationTypeStartAnalyticsInstance,
	"STOP_ANALYTICS_INSTANCE":               WorkRequestOperationTypeStopAnalyticsInstance,
	"SCALE_ANALYTICS_INSTANCE":              WorkRequestOperationTypeScaleAnalyticsInstance,
	"CHANGE_ANALYTICS_INSTANCE_COMPARTMENT": WorkRequestOperationTypeChangeAnalyticsInstanceCompartment,
}

// GetWorkRequestOperationTypeEnumValues Enumerates the set of values for WorkRequestOperationTypeEnum
func GetWorkRequestOperationTypeEnumValues() []WorkRequestOperationTypeEnum {
	values := make([]WorkRequestOperationTypeEnum, 0)
	for _, v := range mappingWorkRequestOperationType {
		values = append(values, v)
	}
	return values
}
