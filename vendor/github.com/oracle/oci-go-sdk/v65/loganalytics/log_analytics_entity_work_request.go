// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// LogAnalyticsEntityWorkRequest LogAnalyticsEntityWorkRequest
type LogAnalyticsEntityWorkRequest struct {

	// The work request unique identifier.
	Id *string `mandatory:"true" json:"id"`

	// The compartment unique identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// An enum description of the type of work the work request is doing.
	OperationType LogAnalyticsEntityWorkRequestOperationTypeEnum `mandatory:"true" json:"operationType"`

	// The completion percentage.
	PercentComplete *int64 `mandatory:"true" json:"percentComplete"`

	// The time at which the work request was updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The time at which the work request was accepted.
	TimeAccepted *common.SDKTime `mandatory:"true" json:"timeAccepted"`

	// The current status of the work request. Valid values are ACCEPTED, IN_PROGRESS, SUCCEEDED or FAILED
	Status LogAnalyticsEntityWorkRequestStatusEnum `mandatory:"true" json:"status"`

	// The time at which the work request was started.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The time at which the work request was finished.
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// Information on the number of entities deleted if work request is finished or in-progress. Failure reason if the work request failed.
	WorkRequestInfo *string `mandatory:"false" json:"workRequestInfo"`
}

func (m LogAnalyticsEntityWorkRequest) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogAnalyticsEntityWorkRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLogAnalyticsEntityWorkRequestOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetLogAnalyticsEntityWorkRequestOperationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLogAnalyticsEntityWorkRequestStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetLogAnalyticsEntityWorkRequestStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LogAnalyticsEntityWorkRequestOperationTypeEnum Enum with underlying type: string
type LogAnalyticsEntityWorkRequestOperationTypeEnum string

// Set of constants representing the allowable values for LogAnalyticsEntityWorkRequestOperationTypeEnum
const (
	LogAnalyticsEntityWorkRequestOperationTypeDeleteEntities LogAnalyticsEntityWorkRequestOperationTypeEnum = "DELETE_ENTITIES"
)

var mappingLogAnalyticsEntityWorkRequestOperationTypeEnum = map[string]LogAnalyticsEntityWorkRequestOperationTypeEnum{
	"DELETE_ENTITIES": LogAnalyticsEntityWorkRequestOperationTypeDeleteEntities,
}

var mappingLogAnalyticsEntityWorkRequestOperationTypeEnumLowerCase = map[string]LogAnalyticsEntityWorkRequestOperationTypeEnum{
	"delete_entities": LogAnalyticsEntityWorkRequestOperationTypeDeleteEntities,
}

// GetLogAnalyticsEntityWorkRequestOperationTypeEnumValues Enumerates the set of values for LogAnalyticsEntityWorkRequestOperationTypeEnum
func GetLogAnalyticsEntityWorkRequestOperationTypeEnumValues() []LogAnalyticsEntityWorkRequestOperationTypeEnum {
	values := make([]LogAnalyticsEntityWorkRequestOperationTypeEnum, 0)
	for _, v := range mappingLogAnalyticsEntityWorkRequestOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetLogAnalyticsEntityWorkRequestOperationTypeEnumStringValues Enumerates the set of values in String for LogAnalyticsEntityWorkRequestOperationTypeEnum
func GetLogAnalyticsEntityWorkRequestOperationTypeEnumStringValues() []string {
	return []string{
		"DELETE_ENTITIES",
	}
}

// GetMappingLogAnalyticsEntityWorkRequestOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLogAnalyticsEntityWorkRequestOperationTypeEnum(val string) (LogAnalyticsEntityWorkRequestOperationTypeEnum, bool) {
	enum, ok := mappingLogAnalyticsEntityWorkRequestOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// LogAnalyticsEntityWorkRequestStatusEnum Enum with underlying type: string
type LogAnalyticsEntityWorkRequestStatusEnum string

// Set of constants representing the allowable values for LogAnalyticsEntityWorkRequestStatusEnum
const (
	LogAnalyticsEntityWorkRequestStatusAccepted   LogAnalyticsEntityWorkRequestStatusEnum = "ACCEPTED"
	LogAnalyticsEntityWorkRequestStatusInProgress LogAnalyticsEntityWorkRequestStatusEnum = "IN_PROGRESS"
	LogAnalyticsEntityWorkRequestStatusSucceeded  LogAnalyticsEntityWorkRequestStatusEnum = "SUCCEEDED"
	LogAnalyticsEntityWorkRequestStatusFailed     LogAnalyticsEntityWorkRequestStatusEnum = "FAILED"
)

var mappingLogAnalyticsEntityWorkRequestStatusEnum = map[string]LogAnalyticsEntityWorkRequestStatusEnum{
	"ACCEPTED":    LogAnalyticsEntityWorkRequestStatusAccepted,
	"IN_PROGRESS": LogAnalyticsEntityWorkRequestStatusInProgress,
	"SUCCEEDED":   LogAnalyticsEntityWorkRequestStatusSucceeded,
	"FAILED":      LogAnalyticsEntityWorkRequestStatusFailed,
}

var mappingLogAnalyticsEntityWorkRequestStatusEnumLowerCase = map[string]LogAnalyticsEntityWorkRequestStatusEnum{
	"accepted":    LogAnalyticsEntityWorkRequestStatusAccepted,
	"in_progress": LogAnalyticsEntityWorkRequestStatusInProgress,
	"succeeded":   LogAnalyticsEntityWorkRequestStatusSucceeded,
	"failed":      LogAnalyticsEntityWorkRequestStatusFailed,
}

// GetLogAnalyticsEntityWorkRequestStatusEnumValues Enumerates the set of values for LogAnalyticsEntityWorkRequestStatusEnum
func GetLogAnalyticsEntityWorkRequestStatusEnumValues() []LogAnalyticsEntityWorkRequestStatusEnum {
	values := make([]LogAnalyticsEntityWorkRequestStatusEnum, 0)
	for _, v := range mappingLogAnalyticsEntityWorkRequestStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetLogAnalyticsEntityWorkRequestStatusEnumStringValues Enumerates the set of values in String for LogAnalyticsEntityWorkRequestStatusEnum
func GetLogAnalyticsEntityWorkRequestStatusEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
	}
}

// GetMappingLogAnalyticsEntityWorkRequestStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLogAnalyticsEntityWorkRequestStatusEnum(val string) (LogAnalyticsEntityWorkRequestStatusEnum, bool) {
	enum, ok := mappingLogAnalyticsEntityWorkRequestStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
