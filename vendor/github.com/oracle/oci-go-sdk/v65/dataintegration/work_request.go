// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// WorkRequest The API operations used to create and configure Data Integration resources do not take effect immediately. In these cases, the operation spawns an asynchronous workflow to fulfill the request. Work requests provide visibility into the status of these in-progress, long-running asynchronous workflows.
type WorkRequest struct {

	// The asynchronous operation tracked by this work request.
	OperationType WorkRequestOperationTypeEnum `mandatory:"true" json:"operationType"`

	// The status of this work request.
	Status WorkRequestStatusEnum `mandatory:"true" json:"status"`

	// The ID of the work request.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that contains this work request. Work requests should be scoped to
	// the same compartment as the resource the work request affects. If the work request affects multiple resources that are not in the same compartment, then the system picks a primary
	// resource whose compartment should be used.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The resources affected by this work request.
	Resources []WorkRequestResource `mandatory:"true" json:"resources"`

	// The completed percentage of the operation tracked by this work request.
	PercentComplete *float32 `mandatory:"true" json:"percentComplete"`

	// The date and time this work request was accepted, in the timestamp format defined by
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeAccepted *common.SDKTime `mandatory:"true" json:"timeAccepted"`

	// The date and time the work request transitioned from `ACCEPTED` to `IN_PROGRESS`, in the timestamp format defined by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time the work request reached a terminal state, either `FAILED` or `SUCCEEDED`, in the timestamp format defined by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
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
	WorkRequestOperationTypeCreateWorkspace WorkRequestOperationTypeEnum = "CREATE_WORKSPACE"
	WorkRequestOperationTypeUpdateWorkspace WorkRequestOperationTypeEnum = "UPDATE_WORKSPACE"
	WorkRequestOperationTypeDeleteWorkspace WorkRequestOperationTypeEnum = "DELETE_WORKSPACE"
	WorkRequestOperationTypeMoveWorkspace   WorkRequestOperationTypeEnum = "MOVE_WORKSPACE"
)

var mappingWorkRequestOperationTypeEnum = map[string]WorkRequestOperationTypeEnum{
	"CREATE_WORKSPACE": WorkRequestOperationTypeCreateWorkspace,
	"UPDATE_WORKSPACE": WorkRequestOperationTypeUpdateWorkspace,
	"DELETE_WORKSPACE": WorkRequestOperationTypeDeleteWorkspace,
	"MOVE_WORKSPACE":   WorkRequestOperationTypeMoveWorkspace,
}

var mappingWorkRequestOperationTypeEnumLowerCase = map[string]WorkRequestOperationTypeEnum{
	"create_workspace": WorkRequestOperationTypeCreateWorkspace,
	"update_workspace": WorkRequestOperationTypeUpdateWorkspace,
	"delete_workspace": WorkRequestOperationTypeDeleteWorkspace,
	"move_workspace":   WorkRequestOperationTypeMoveWorkspace,
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
		"CREATE_WORKSPACE",
		"UPDATE_WORKSPACE",
		"DELETE_WORKSPACE",
		"MOVE_WORKSPACE",
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
