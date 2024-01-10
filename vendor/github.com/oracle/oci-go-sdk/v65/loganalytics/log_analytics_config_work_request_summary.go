// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LogAnalyticsConfigWorkRequestSummary LogAnalyticsConfigWorkRequestSummary
type LogAnalyticsConfigWorkRequestSummary struct {

	// The workrequest unique identifier.
	Id *string `mandatory:"false" json:"id"`

	// The compartment unique identifier.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The operation type
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogAnalyticsConfigWorkRequestSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLogAnalyticsConfigWorkRequestSummaryOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetLogAnalyticsConfigWorkRequestSummaryOperationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLogAnalyticsConfigWorkRequestSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLogAnalyticsConfigWorkRequestSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LogAnalyticsConfigWorkRequestSummaryOperationTypeEnum Enum with underlying type: string
type LogAnalyticsConfigWorkRequestSummaryOperationTypeEnum string

// Set of constants representing the allowable values for LogAnalyticsConfigWorkRequestSummaryOperationTypeEnum
const (
	LogAnalyticsConfigWorkRequestSummaryOperationTypeCreateAssociations    LogAnalyticsConfigWorkRequestSummaryOperationTypeEnum = "CREATE_ASSOCIATIONS"
	LogAnalyticsConfigWorkRequestSummaryOperationTypeDeleteAssociations    LogAnalyticsConfigWorkRequestSummaryOperationTypeEnum = "DELETE_ASSOCIATIONS"
	LogAnalyticsConfigWorkRequestSummaryOperationTypeAppendLookupData      LogAnalyticsConfigWorkRequestSummaryOperationTypeEnum = "APPEND_LOOKUP_DATA"
	LogAnalyticsConfigWorkRequestSummaryOperationTypeUpdateLookupData      LogAnalyticsConfigWorkRequestSummaryOperationTypeEnum = "UPDATE_LOOKUP_DATA"
	LogAnalyticsConfigWorkRequestSummaryOperationTypeDeleteLookup          LogAnalyticsConfigWorkRequestSummaryOperationTypeEnum = "DELETE_LOOKUP"
	LogAnalyticsConfigWorkRequestSummaryOperationTypeEnableIngestTimeRule  LogAnalyticsConfigWorkRequestSummaryOperationTypeEnum = "ENABLE_INGEST_TIME_RULE"
	LogAnalyticsConfigWorkRequestSummaryOperationTypeDisableIngestTimeRule LogAnalyticsConfigWorkRequestSummaryOperationTypeEnum = "DISABLE_INGEST_TIME_RULE"
)

var mappingLogAnalyticsConfigWorkRequestSummaryOperationTypeEnum = map[string]LogAnalyticsConfigWorkRequestSummaryOperationTypeEnum{
	"CREATE_ASSOCIATIONS":      LogAnalyticsConfigWorkRequestSummaryOperationTypeCreateAssociations,
	"DELETE_ASSOCIATIONS":      LogAnalyticsConfigWorkRequestSummaryOperationTypeDeleteAssociations,
	"APPEND_LOOKUP_DATA":       LogAnalyticsConfigWorkRequestSummaryOperationTypeAppendLookupData,
	"UPDATE_LOOKUP_DATA":       LogAnalyticsConfigWorkRequestSummaryOperationTypeUpdateLookupData,
	"DELETE_LOOKUP":            LogAnalyticsConfigWorkRequestSummaryOperationTypeDeleteLookup,
	"ENABLE_INGEST_TIME_RULE":  LogAnalyticsConfigWorkRequestSummaryOperationTypeEnableIngestTimeRule,
	"DISABLE_INGEST_TIME_RULE": LogAnalyticsConfigWorkRequestSummaryOperationTypeDisableIngestTimeRule,
}

var mappingLogAnalyticsConfigWorkRequestSummaryOperationTypeEnumLowerCase = map[string]LogAnalyticsConfigWorkRequestSummaryOperationTypeEnum{
	"create_associations":      LogAnalyticsConfigWorkRequestSummaryOperationTypeCreateAssociations,
	"delete_associations":      LogAnalyticsConfigWorkRequestSummaryOperationTypeDeleteAssociations,
	"append_lookup_data":       LogAnalyticsConfigWorkRequestSummaryOperationTypeAppendLookupData,
	"update_lookup_data":       LogAnalyticsConfigWorkRequestSummaryOperationTypeUpdateLookupData,
	"delete_lookup":            LogAnalyticsConfigWorkRequestSummaryOperationTypeDeleteLookup,
	"enable_ingest_time_rule":  LogAnalyticsConfigWorkRequestSummaryOperationTypeEnableIngestTimeRule,
	"disable_ingest_time_rule": LogAnalyticsConfigWorkRequestSummaryOperationTypeDisableIngestTimeRule,
}

// GetLogAnalyticsConfigWorkRequestSummaryOperationTypeEnumValues Enumerates the set of values for LogAnalyticsConfigWorkRequestSummaryOperationTypeEnum
func GetLogAnalyticsConfigWorkRequestSummaryOperationTypeEnumValues() []LogAnalyticsConfigWorkRequestSummaryOperationTypeEnum {
	values := make([]LogAnalyticsConfigWorkRequestSummaryOperationTypeEnum, 0)
	for _, v := range mappingLogAnalyticsConfigWorkRequestSummaryOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetLogAnalyticsConfigWorkRequestSummaryOperationTypeEnumStringValues Enumerates the set of values in String for LogAnalyticsConfigWorkRequestSummaryOperationTypeEnum
func GetLogAnalyticsConfigWorkRequestSummaryOperationTypeEnumStringValues() []string {
	return []string{
		"CREATE_ASSOCIATIONS",
		"DELETE_ASSOCIATIONS",
		"APPEND_LOOKUP_DATA",
		"UPDATE_LOOKUP_DATA",
		"DELETE_LOOKUP",
		"ENABLE_INGEST_TIME_RULE",
		"DISABLE_INGEST_TIME_RULE",
	}
}

// GetMappingLogAnalyticsConfigWorkRequestSummaryOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLogAnalyticsConfigWorkRequestSummaryOperationTypeEnum(val string) (LogAnalyticsConfigWorkRequestSummaryOperationTypeEnum, bool) {
	enum, ok := mappingLogAnalyticsConfigWorkRequestSummaryOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
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

var mappingLogAnalyticsConfigWorkRequestSummaryLifecycleStateEnum = map[string]LogAnalyticsConfigWorkRequestSummaryLifecycleStateEnum{
	"ACCEPTED":    LogAnalyticsConfigWorkRequestSummaryLifecycleStateAccepted,
	"IN_PROGRESS": LogAnalyticsConfigWorkRequestSummaryLifecycleStateInProgress,
	"SUCCEEDED":   LogAnalyticsConfigWorkRequestSummaryLifecycleStateSucceeded,
	"FAILED":      LogAnalyticsConfigWorkRequestSummaryLifecycleStateFailed,
}

var mappingLogAnalyticsConfigWorkRequestSummaryLifecycleStateEnumLowerCase = map[string]LogAnalyticsConfigWorkRequestSummaryLifecycleStateEnum{
	"accepted":    LogAnalyticsConfigWorkRequestSummaryLifecycleStateAccepted,
	"in_progress": LogAnalyticsConfigWorkRequestSummaryLifecycleStateInProgress,
	"succeeded":   LogAnalyticsConfigWorkRequestSummaryLifecycleStateSucceeded,
	"failed":      LogAnalyticsConfigWorkRequestSummaryLifecycleStateFailed,
}

// GetLogAnalyticsConfigWorkRequestSummaryLifecycleStateEnumValues Enumerates the set of values for LogAnalyticsConfigWorkRequestSummaryLifecycleStateEnum
func GetLogAnalyticsConfigWorkRequestSummaryLifecycleStateEnumValues() []LogAnalyticsConfigWorkRequestSummaryLifecycleStateEnum {
	values := make([]LogAnalyticsConfigWorkRequestSummaryLifecycleStateEnum, 0)
	for _, v := range mappingLogAnalyticsConfigWorkRequestSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetLogAnalyticsConfigWorkRequestSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for LogAnalyticsConfigWorkRequestSummaryLifecycleStateEnum
func GetLogAnalyticsConfigWorkRequestSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
	}
}

// GetMappingLogAnalyticsConfigWorkRequestSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLogAnalyticsConfigWorkRequestSummaryLifecycleStateEnum(val string) (LogAnalyticsConfigWorkRequestSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingLogAnalyticsConfigWorkRequestSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
