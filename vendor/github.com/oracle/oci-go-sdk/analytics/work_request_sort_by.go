// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

// WorkRequestSortByEnum Enum with underlying type: string
type WorkRequestSortByEnum string

// Set of constants representing the allowable values for WorkRequestSortByEnum
const (
	WorkRequestSortById            WorkRequestSortByEnum = "id"
	WorkRequestSortByOperationType WorkRequestSortByEnum = "operationType"
	WorkRequestSortByStatus        WorkRequestSortByEnum = "status"
	WorkRequestSortByTimeAccepted  WorkRequestSortByEnum = "timeAccepted"
	WorkRequestSortByTimeStarted   WorkRequestSortByEnum = "timeStarted"
	WorkRequestSortByTimeFinished  WorkRequestSortByEnum = "timeFinished"
)

var mappingWorkRequestSortBy = map[string]WorkRequestSortByEnum{
	"id":            WorkRequestSortById,
	"operationType": WorkRequestSortByOperationType,
	"status":        WorkRequestSortByStatus,
	"timeAccepted":  WorkRequestSortByTimeAccepted,
	"timeStarted":   WorkRequestSortByTimeStarted,
	"timeFinished":  WorkRequestSortByTimeFinished,
}

// GetWorkRequestSortByEnumValues Enumerates the set of values for WorkRequestSortByEnum
func GetWorkRequestSortByEnumValues() []WorkRequestSortByEnum {
	values := make([]WorkRequestSortByEnum, 0)
	for _, v := range mappingWorkRequestSortBy {
		values = append(values, v)
	}
	return values
}
