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

// LogAnalyticsConfigWorkRequest LogAnalyticsConfigWorkRequest
type LogAnalyticsConfigWorkRequest struct {

	// The workrequest unique identifier.
	Id *string `mandatory:"false" json:"id"`

	// The compartment unique identifier.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The operation type
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogAnalyticsConfigWorkRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLogAnalyticsConfigWorkRequestOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetLogAnalyticsConfigWorkRequestOperationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLogAnalyticsConfigWorkRequestLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLogAnalyticsConfigWorkRequestLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LogAnalyticsConfigWorkRequestOperationTypeEnum Enum with underlying type: string
type LogAnalyticsConfigWorkRequestOperationTypeEnum string

// Set of constants representing the allowable values for LogAnalyticsConfigWorkRequestOperationTypeEnum
const (
	LogAnalyticsConfigWorkRequestOperationTypeCreateAssociations    LogAnalyticsConfigWorkRequestOperationTypeEnum = "CREATE_ASSOCIATIONS"
	LogAnalyticsConfigWorkRequestOperationTypeDeleteAssociations    LogAnalyticsConfigWorkRequestOperationTypeEnum = "DELETE_ASSOCIATIONS"
	LogAnalyticsConfigWorkRequestOperationTypeAppendLookupData      LogAnalyticsConfigWorkRequestOperationTypeEnum = "APPEND_LOOKUP_DATA"
	LogAnalyticsConfigWorkRequestOperationTypeUpdateLookupData      LogAnalyticsConfigWorkRequestOperationTypeEnum = "UPDATE_LOOKUP_DATA"
	LogAnalyticsConfigWorkRequestOperationTypeDeleteLookup          LogAnalyticsConfigWorkRequestOperationTypeEnum = "DELETE_LOOKUP"
	LogAnalyticsConfigWorkRequestOperationTypeEnableIngestTimeRule  LogAnalyticsConfigWorkRequestOperationTypeEnum = "ENABLE_INGEST_TIME_RULE"
	LogAnalyticsConfigWorkRequestOperationTypeDisableIngestTimeRule LogAnalyticsConfigWorkRequestOperationTypeEnum = "DISABLE_INGEST_TIME_RULE"
)

var mappingLogAnalyticsConfigWorkRequestOperationTypeEnum = map[string]LogAnalyticsConfigWorkRequestOperationTypeEnum{
	"CREATE_ASSOCIATIONS":      LogAnalyticsConfigWorkRequestOperationTypeCreateAssociations,
	"DELETE_ASSOCIATIONS":      LogAnalyticsConfigWorkRequestOperationTypeDeleteAssociations,
	"APPEND_LOOKUP_DATA":       LogAnalyticsConfigWorkRequestOperationTypeAppendLookupData,
	"UPDATE_LOOKUP_DATA":       LogAnalyticsConfigWorkRequestOperationTypeUpdateLookupData,
	"DELETE_LOOKUP":            LogAnalyticsConfigWorkRequestOperationTypeDeleteLookup,
	"ENABLE_INGEST_TIME_RULE":  LogAnalyticsConfigWorkRequestOperationTypeEnableIngestTimeRule,
	"DISABLE_INGEST_TIME_RULE": LogAnalyticsConfigWorkRequestOperationTypeDisableIngestTimeRule,
}

var mappingLogAnalyticsConfigWorkRequestOperationTypeEnumLowerCase = map[string]LogAnalyticsConfigWorkRequestOperationTypeEnum{
	"create_associations":      LogAnalyticsConfigWorkRequestOperationTypeCreateAssociations,
	"delete_associations":      LogAnalyticsConfigWorkRequestOperationTypeDeleteAssociations,
	"append_lookup_data":       LogAnalyticsConfigWorkRequestOperationTypeAppendLookupData,
	"update_lookup_data":       LogAnalyticsConfigWorkRequestOperationTypeUpdateLookupData,
	"delete_lookup":            LogAnalyticsConfigWorkRequestOperationTypeDeleteLookup,
	"enable_ingest_time_rule":  LogAnalyticsConfigWorkRequestOperationTypeEnableIngestTimeRule,
	"disable_ingest_time_rule": LogAnalyticsConfigWorkRequestOperationTypeDisableIngestTimeRule,
}

// GetLogAnalyticsConfigWorkRequestOperationTypeEnumValues Enumerates the set of values for LogAnalyticsConfigWorkRequestOperationTypeEnum
func GetLogAnalyticsConfigWorkRequestOperationTypeEnumValues() []LogAnalyticsConfigWorkRequestOperationTypeEnum {
	values := make([]LogAnalyticsConfigWorkRequestOperationTypeEnum, 0)
	for _, v := range mappingLogAnalyticsConfigWorkRequestOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetLogAnalyticsConfigWorkRequestOperationTypeEnumStringValues Enumerates the set of values in String for LogAnalyticsConfigWorkRequestOperationTypeEnum
func GetLogAnalyticsConfigWorkRequestOperationTypeEnumStringValues() []string {
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

// GetMappingLogAnalyticsConfigWorkRequestOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLogAnalyticsConfigWorkRequestOperationTypeEnum(val string) (LogAnalyticsConfigWorkRequestOperationTypeEnum, bool) {
	enum, ok := mappingLogAnalyticsConfigWorkRequestOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
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

var mappingLogAnalyticsConfigWorkRequestLifecycleStateEnum = map[string]LogAnalyticsConfigWorkRequestLifecycleStateEnum{
	"ACCEPTED":    LogAnalyticsConfigWorkRequestLifecycleStateAccepted,
	"IN_PROGRESS": LogAnalyticsConfigWorkRequestLifecycleStateInProgress,
	"SUCCEEDED":   LogAnalyticsConfigWorkRequestLifecycleStateSucceeded,
	"FAILED":      LogAnalyticsConfigWorkRequestLifecycleStateFailed,
}

var mappingLogAnalyticsConfigWorkRequestLifecycleStateEnumLowerCase = map[string]LogAnalyticsConfigWorkRequestLifecycleStateEnum{
	"accepted":    LogAnalyticsConfigWorkRequestLifecycleStateAccepted,
	"in_progress": LogAnalyticsConfigWorkRequestLifecycleStateInProgress,
	"succeeded":   LogAnalyticsConfigWorkRequestLifecycleStateSucceeded,
	"failed":      LogAnalyticsConfigWorkRequestLifecycleStateFailed,
}

// GetLogAnalyticsConfigWorkRequestLifecycleStateEnumValues Enumerates the set of values for LogAnalyticsConfigWorkRequestLifecycleStateEnum
func GetLogAnalyticsConfigWorkRequestLifecycleStateEnumValues() []LogAnalyticsConfigWorkRequestLifecycleStateEnum {
	values := make([]LogAnalyticsConfigWorkRequestLifecycleStateEnum, 0)
	for _, v := range mappingLogAnalyticsConfigWorkRequestLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetLogAnalyticsConfigWorkRequestLifecycleStateEnumStringValues Enumerates the set of values in String for LogAnalyticsConfigWorkRequestLifecycleStateEnum
func GetLogAnalyticsConfigWorkRequestLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
	}
}

// GetMappingLogAnalyticsConfigWorkRequestLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLogAnalyticsConfigWorkRequestLifecycleStateEnum(val string) (LogAnalyticsConfigWorkRequestLifecycleStateEnum, bool) {
	enum, ok := mappingLogAnalyticsConfigWorkRequestLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
