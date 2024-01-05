// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration (WAA) API
//
// API for the Web Application Acceleration service.
// Use this API to manage regional Web App Acceleration policies such as Caching and Compression
// for accelerating HTTP services.
//

package waa

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// WorkRequest A description of WorkRequest status
type WorkRequest struct {

	// Type of the WorkRequest
	OperationType WorkRequestOperationTypeEnum `mandatory:"true" json:"operationType"`

	// Status of current work request.
	Status WorkRequestStatusEnum `mandatory:"true" json:"status"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the WorkRequest.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment that contains the WorkRequest.
	// WorkRequests should be scoped to the same compartment as the resource the work request affects.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The resources affected by this WorkRequest.
	Resources []WorkRequestResource `mandatory:"true" json:"resources"`

	// Percentage of the request completed.
	PercentComplete *float32 `mandatory:"true" json:"percentComplete"`

	// The date and time the request was created, as described in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeAccepted *common.SDKTime `mandatory:"true" json:"timeAccepted"`

	// The date and time the request was started, as described in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time the object was finished, as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`
}

func (m WorkRequest) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WorkRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingWorkRequestOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetWorkRequestOperationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingWorkRequestStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetWorkRequestStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// WorkRequestOperationTypeEnum Enum with underlying type: string
type WorkRequestOperationTypeEnum string

// Set of constants representing the allowable values for WorkRequestOperationTypeEnum
const (
	WorkRequestOperationTypeCreateWaaPolicy              WorkRequestOperationTypeEnum = "CREATE_WAA_POLICY"
	WorkRequestOperationTypeUpdateWaaPolicy              WorkRequestOperationTypeEnum = "UPDATE_WAA_POLICY"
	WorkRequestOperationTypeDeleteWaaPolicy              WorkRequestOperationTypeEnum = "DELETE_WAA_POLICY"
	WorkRequestOperationTypeMoveWaaPolicy                WorkRequestOperationTypeEnum = "MOVE_WAA_POLICY"
	WorkRequestOperationTypeCreateWebAppAcceleration     WorkRequestOperationTypeEnum = "CREATE_WEB_APP_ACCELERATION"
	WorkRequestOperationTypeUpdateWebAppAcceleration     WorkRequestOperationTypeEnum = "UPDATE_WEB_APP_ACCELERATION"
	WorkRequestOperationTypeDeleteWebAppAcceleration     WorkRequestOperationTypeEnum = "DELETE_WEB_APP_ACCELERATION"
	WorkRequestOperationTypeMoveWebAppAcceleration       WorkRequestOperationTypeEnum = "MOVE_WEB_APP_ACCELERATION"
	WorkRequestOperationTypePurgeWebAppAccelerationCache WorkRequestOperationTypeEnum = "PURGE_WEB_APP_ACCELERATION_CACHE"
)

var mappingWorkRequestOperationTypeEnum = map[string]WorkRequestOperationTypeEnum{
	"CREATE_WAA_POLICY":                WorkRequestOperationTypeCreateWaaPolicy,
	"UPDATE_WAA_POLICY":                WorkRequestOperationTypeUpdateWaaPolicy,
	"DELETE_WAA_POLICY":                WorkRequestOperationTypeDeleteWaaPolicy,
	"MOVE_WAA_POLICY":                  WorkRequestOperationTypeMoveWaaPolicy,
	"CREATE_WEB_APP_ACCELERATION":      WorkRequestOperationTypeCreateWebAppAcceleration,
	"UPDATE_WEB_APP_ACCELERATION":      WorkRequestOperationTypeUpdateWebAppAcceleration,
	"DELETE_WEB_APP_ACCELERATION":      WorkRequestOperationTypeDeleteWebAppAcceleration,
	"MOVE_WEB_APP_ACCELERATION":        WorkRequestOperationTypeMoveWebAppAcceleration,
	"PURGE_WEB_APP_ACCELERATION_CACHE": WorkRequestOperationTypePurgeWebAppAccelerationCache,
}

var mappingWorkRequestOperationTypeEnumLowerCase = map[string]WorkRequestOperationTypeEnum{
	"create_waa_policy":                WorkRequestOperationTypeCreateWaaPolicy,
	"update_waa_policy":                WorkRequestOperationTypeUpdateWaaPolicy,
	"delete_waa_policy":                WorkRequestOperationTypeDeleteWaaPolicy,
	"move_waa_policy":                  WorkRequestOperationTypeMoveWaaPolicy,
	"create_web_app_acceleration":      WorkRequestOperationTypeCreateWebAppAcceleration,
	"update_web_app_acceleration":      WorkRequestOperationTypeUpdateWebAppAcceleration,
	"delete_web_app_acceleration":      WorkRequestOperationTypeDeleteWebAppAcceleration,
	"move_web_app_acceleration":        WorkRequestOperationTypeMoveWebAppAcceleration,
	"purge_web_app_acceleration_cache": WorkRequestOperationTypePurgeWebAppAccelerationCache,
}

// GetWorkRequestOperationTypeEnumValues Enumerates the set of values for WorkRequestOperationTypeEnum
func GetWorkRequestOperationTypeEnumValues() []WorkRequestOperationTypeEnum {
	values := make([]WorkRequestOperationTypeEnum, 0)
	for _, v := range mappingWorkRequestOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestOperationTypeEnumStringValues Enumerates the set of values in String for WorkRequestOperationTypeEnum
func GetWorkRequestOperationTypeEnumStringValues() []string {
	return []string{
		"CREATE_WAA_POLICY",
		"UPDATE_WAA_POLICY",
		"DELETE_WAA_POLICY",
		"MOVE_WAA_POLICY",
		"CREATE_WEB_APP_ACCELERATION",
		"UPDATE_WEB_APP_ACCELERATION",
		"DELETE_WEB_APP_ACCELERATION",
		"MOVE_WEB_APP_ACCELERATION",
		"PURGE_WEB_APP_ACCELERATION_CACHE",
	}
}

// GetMappingWorkRequestOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestOperationTypeEnum(val string) (WorkRequestOperationTypeEnum, bool) {
	enum, ok := mappingWorkRequestOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// WorkRequestStatusEnum Enum with underlying type: string
type WorkRequestStatusEnum string

// Set of constants representing the allowable values for WorkRequestStatusEnum
const (
	WorkRequestStatusAccepted   WorkRequestStatusEnum = "ACCEPTED"
	WorkRequestStatusInProgress WorkRequestStatusEnum = "IN_PROGRESS"
	WorkRequestStatusFailed     WorkRequestStatusEnum = "FAILED"
	WorkRequestStatusSucceeded  WorkRequestStatusEnum = "SUCCEEDED"
	WorkRequestStatusCanceling  WorkRequestStatusEnum = "CANCELING"
	WorkRequestStatusCanceled   WorkRequestStatusEnum = "CANCELED"
)

var mappingWorkRequestStatusEnum = map[string]WorkRequestStatusEnum{
	"ACCEPTED":    WorkRequestStatusAccepted,
	"IN_PROGRESS": WorkRequestStatusInProgress,
	"FAILED":      WorkRequestStatusFailed,
	"SUCCEEDED":   WorkRequestStatusSucceeded,
	"CANCELING":   WorkRequestStatusCanceling,
	"CANCELED":    WorkRequestStatusCanceled,
}

var mappingWorkRequestStatusEnumLowerCase = map[string]WorkRequestStatusEnum{
	"accepted":    WorkRequestStatusAccepted,
	"in_progress": WorkRequestStatusInProgress,
	"failed":      WorkRequestStatusFailed,
	"succeeded":   WorkRequestStatusSucceeded,
	"canceling":   WorkRequestStatusCanceling,
	"canceled":    WorkRequestStatusCanceled,
}

// GetWorkRequestStatusEnumValues Enumerates the set of values for WorkRequestStatusEnum
func GetWorkRequestStatusEnumValues() []WorkRequestStatusEnum {
	values := make([]WorkRequestStatusEnum, 0)
	for _, v := range mappingWorkRequestStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestStatusEnumStringValues Enumerates the set of values in String for WorkRequestStatusEnum
func GetWorkRequestStatusEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingWorkRequestStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestStatusEnum(val string) (WorkRequestStatusEnum, bool) {
	enum, ok := mappingWorkRequestStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
