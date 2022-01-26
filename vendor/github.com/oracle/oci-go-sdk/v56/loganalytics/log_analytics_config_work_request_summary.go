// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// LogAnalyticsConfigWorkRequestSummary LogAnalyticsConfigWorkRequestSummary
type LogAnalyticsConfigWorkRequestSummary struct {

	// The workrequest unique identifier.
	Id *string `mandatory:"false" json:"id"`

	// The compartment unique identifier.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The operation type.  There are two classes of operations, association operations and
	// lookup operations.  Associations may be created or deleted, and lookup operations include
	// append, update and delete.
	OperationType LogAnalyticsConfigWorkRequestSummaryOperationTypeEnum `mandatory:"false" json:"operationType,omitempty"`

	// The completion percentage.
	PercentComplete *int64 `mandatory:"false" json:"percentComplete"`

	// The time at which the work request finished.
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// The time at which the work request was accepted.
	TimeAccepted *common.SDKTime `mandatory:"false" json:"timeAccepted"`

	// The lifecycle status.  Valid values are ACCEPTED, IN_PROGRESS, SUCCEEDED
	// or FAILED.
	LifecycleState LogAnalyticsConfigWorkRequestSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m LogAnalyticsConfigWorkRequestSummary) String() string {
	return common.PointerString(m)
}

// LogAnalyticsConfigWorkRequestSummaryOperationTypeEnum Enum with underlying type: string
type LogAnalyticsConfigWorkRequestSummaryOperationTypeEnum string

// Set of constants representing the allowable values for LogAnalyticsConfigWorkRequestSummaryOperationTypeEnum
const (
	LogAnalyticsConfigWorkRequestSummaryOperationTypeCreateAssociations LogAnalyticsConfigWorkRequestSummaryOperationTypeEnum = "CREATE_ASSOCIATIONS"
	LogAnalyticsConfigWorkRequestSummaryOperationTypeDeleteAssociations LogAnalyticsConfigWorkRequestSummaryOperationTypeEnum = "DELETE_ASSOCIATIONS"
	LogAnalyticsConfigWorkRequestSummaryOperationTypeAppendLookupData   LogAnalyticsConfigWorkRequestSummaryOperationTypeEnum = "APPEND_LOOKUP_DATA"
	LogAnalyticsConfigWorkRequestSummaryOperationTypeUpdateLookupData   LogAnalyticsConfigWorkRequestSummaryOperationTypeEnum = "UPDATE_LOOKUP_DATA"
	LogAnalyticsConfigWorkRequestSummaryOperationTypeDeleteLookup       LogAnalyticsConfigWorkRequestSummaryOperationTypeEnum = "DELETE_LOOKUP"
)

var mappingLogAnalyticsConfigWorkRequestSummaryOperationType = map[string]LogAnalyticsConfigWorkRequestSummaryOperationTypeEnum{
	"CREATE_ASSOCIATIONS": LogAnalyticsConfigWorkRequestSummaryOperationTypeCreateAssociations,
	"DELETE_ASSOCIATIONS": LogAnalyticsConfigWorkRequestSummaryOperationTypeDeleteAssociations,
	"APPEND_LOOKUP_DATA":  LogAnalyticsConfigWorkRequestSummaryOperationTypeAppendLookupData,
	"UPDATE_LOOKUP_DATA":  LogAnalyticsConfigWorkRequestSummaryOperationTypeUpdateLookupData,
	"DELETE_LOOKUP":       LogAnalyticsConfigWorkRequestSummaryOperationTypeDeleteLookup,
}

// GetLogAnalyticsConfigWorkRequestSummaryOperationTypeEnumValues Enumerates the set of values for LogAnalyticsConfigWorkRequestSummaryOperationTypeEnum
func GetLogAnalyticsConfigWorkRequestSummaryOperationTypeEnumValues() []LogAnalyticsConfigWorkRequestSummaryOperationTypeEnum {
	values := make([]LogAnalyticsConfigWorkRequestSummaryOperationTypeEnum, 0)
	for _, v := range mappingLogAnalyticsConfigWorkRequestSummaryOperationType {
		values = append(values, v)
	}
	return values
}

// LogAnalyticsConfigWorkRequestSummaryLifecycleStateEnum Enum with underlying type: string
type LogAnalyticsConfigWorkRequestSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for LogAnalyticsConfigWorkRequestSummaryLifecycleStateEnum
const (
	LogAnalyticsConfigWorkRequestSummaryLifecycleStateAccepted   LogAnalyticsConfigWorkRequestSummaryLifecycleStateEnum = "ACCEPTED"
	LogAnalyticsConfigWorkRequestSummaryLifecycleStateInProgress LogAnalyticsConfigWorkRequestSummaryLifecycleStateEnum = "IN_PROGRESS"
	LogAnalyticsConfigWorkRequestSummaryLifecycleStateSucceeded  LogAnalyticsConfigWorkRequestSummaryLifecycleStateEnum = "SUCCEEDED"
	LogAnalyticsConfigWorkRequestSummaryLifecycleStateFailed     LogAnalyticsConfigWorkRequestSummaryLifecycleStateEnum = "FAILED"
)

var mappingLogAnalyticsConfigWorkRequestSummaryLifecycleState = map[string]LogAnalyticsConfigWorkRequestSummaryLifecycleStateEnum{
	"ACCEPTED":    LogAnalyticsConfigWorkRequestSummaryLifecycleStateAccepted,
	"IN_PROGRESS": LogAnalyticsConfigWorkRequestSummaryLifecycleStateInProgress,
	"SUCCEEDED":   LogAnalyticsConfigWorkRequestSummaryLifecycleStateSucceeded,
	"FAILED":      LogAnalyticsConfigWorkRequestSummaryLifecycleStateFailed,
}

// GetLogAnalyticsConfigWorkRequestSummaryLifecycleStateEnumValues Enumerates the set of values for LogAnalyticsConfigWorkRequestSummaryLifecycleStateEnum
func GetLogAnalyticsConfigWorkRequestSummaryLifecycleStateEnumValues() []LogAnalyticsConfigWorkRequestSummaryLifecycleStateEnum {
	values := make([]LogAnalyticsConfigWorkRequestSummaryLifecycleStateEnum, 0)
	for _, v := range mappingLogAnalyticsConfigWorkRequestSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}
