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

// LogAnalyticsConfigWorkRequest LogAnalyticsConfigWorkRequest
type LogAnalyticsConfigWorkRequest struct {

	// The workrequest unique identifier.
	Id *string `mandatory:"false" json:"id"`

	// The compartment unique identifier.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The operation type.  There are two classes of operations, association operations and
	// lookup operations.  Associations may be created or deleted, and lookup operations include
	// append, update and delete.
	OperationType LogAnalyticsConfigWorkRequestOperationTypeEnum `mandatory:"false" json:"operationType,omitempty"`

	// The list of config work request responses.
	Payload []LogAnalyticsConfigWorkRequestPayload `mandatory:"false" json:"payload"`

	// The completion percentage.
	PercentComplete *int64 `mandatory:"false" json:"percentComplete"`

	// The time at which the work request was started.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The time at which the work request was accepted.
	TimeAccepted *common.SDKTime `mandatory:"false" json:"timeAccepted"`

	// The time at which the work request was finished.
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// The lifecycle status.  Valid values are ACCEPTED, IN_PROGRESS, SUCCEEDED
	// or FAILED
	LifecycleState LogAnalyticsConfigWorkRequestLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m LogAnalyticsConfigWorkRequest) String() string {
	return common.PointerString(m)
}

// LogAnalyticsConfigWorkRequestOperationTypeEnum Enum with underlying type: string
type LogAnalyticsConfigWorkRequestOperationTypeEnum string

// Set of constants representing the allowable values for LogAnalyticsConfigWorkRequestOperationTypeEnum
const (
	LogAnalyticsConfigWorkRequestOperationTypeCreateAssociations LogAnalyticsConfigWorkRequestOperationTypeEnum = "CREATE_ASSOCIATIONS"
	LogAnalyticsConfigWorkRequestOperationTypeDeleteAssociations LogAnalyticsConfigWorkRequestOperationTypeEnum = "DELETE_ASSOCIATIONS"
	LogAnalyticsConfigWorkRequestOperationTypeAppendLookupData   LogAnalyticsConfigWorkRequestOperationTypeEnum = "APPEND_LOOKUP_DATA"
	LogAnalyticsConfigWorkRequestOperationTypeUpdateLookupData   LogAnalyticsConfigWorkRequestOperationTypeEnum = "UPDATE_LOOKUP_DATA"
	LogAnalyticsConfigWorkRequestOperationTypeDeleteLookup       LogAnalyticsConfigWorkRequestOperationTypeEnum = "DELETE_LOOKUP"
)

var mappingLogAnalyticsConfigWorkRequestOperationType = map[string]LogAnalyticsConfigWorkRequestOperationTypeEnum{
	"CREATE_ASSOCIATIONS": LogAnalyticsConfigWorkRequestOperationTypeCreateAssociations,
	"DELETE_ASSOCIATIONS": LogAnalyticsConfigWorkRequestOperationTypeDeleteAssociations,
	"APPEND_LOOKUP_DATA":  LogAnalyticsConfigWorkRequestOperationTypeAppendLookupData,
	"UPDATE_LOOKUP_DATA":  LogAnalyticsConfigWorkRequestOperationTypeUpdateLookupData,
	"DELETE_LOOKUP":       LogAnalyticsConfigWorkRequestOperationTypeDeleteLookup,
}

// GetLogAnalyticsConfigWorkRequestOperationTypeEnumValues Enumerates the set of values for LogAnalyticsConfigWorkRequestOperationTypeEnum
func GetLogAnalyticsConfigWorkRequestOperationTypeEnumValues() []LogAnalyticsConfigWorkRequestOperationTypeEnum {
	values := make([]LogAnalyticsConfigWorkRequestOperationTypeEnum, 0)
	for _, v := range mappingLogAnalyticsConfigWorkRequestOperationType {
		values = append(values, v)
	}
	return values
}

// LogAnalyticsConfigWorkRequestLifecycleStateEnum Enum with underlying type: string
type LogAnalyticsConfigWorkRequestLifecycleStateEnum string

// Set of constants representing the allowable values for LogAnalyticsConfigWorkRequestLifecycleStateEnum
const (
	LogAnalyticsConfigWorkRequestLifecycleStateAccepted   LogAnalyticsConfigWorkRequestLifecycleStateEnum = "ACCEPTED"
	LogAnalyticsConfigWorkRequestLifecycleStateInProgress LogAnalyticsConfigWorkRequestLifecycleStateEnum = "IN_PROGRESS"
	LogAnalyticsConfigWorkRequestLifecycleStateSucceeded  LogAnalyticsConfigWorkRequestLifecycleStateEnum = "SUCCEEDED"
	LogAnalyticsConfigWorkRequestLifecycleStateFailed     LogAnalyticsConfigWorkRequestLifecycleStateEnum = "FAILED"
)

var mappingLogAnalyticsConfigWorkRequestLifecycleState = map[string]LogAnalyticsConfigWorkRequestLifecycleStateEnum{
	"ACCEPTED":    LogAnalyticsConfigWorkRequestLifecycleStateAccepted,
	"IN_PROGRESS": LogAnalyticsConfigWorkRequestLifecycleStateInProgress,
	"SUCCEEDED":   LogAnalyticsConfigWorkRequestLifecycleStateSucceeded,
	"FAILED":      LogAnalyticsConfigWorkRequestLifecycleStateFailed,
}

// GetLogAnalyticsConfigWorkRequestLifecycleStateEnumValues Enumerates the set of values for LogAnalyticsConfigWorkRequestLifecycleStateEnum
func GetLogAnalyticsConfigWorkRequestLifecycleStateEnumValues() []LogAnalyticsConfigWorkRequestLifecycleStateEnum {
	values := make([]LogAnalyticsConfigWorkRequestLifecycleStateEnum, 0)
	for _, v := range mappingLogAnalyticsConfigWorkRequestLifecycleState {
		values = append(values, v)
	}
	return values
}
