// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// WorkRequestSummary A work request summary object.
type WorkRequestSummary struct {

	// The asynchronous operation tracked by this work request.
	OperationType WorkRequestSummaryOperationTypeEnum `mandatory:"true" json:"operationType"`

	// The status of this work request.
	Status WorkRequestSummaryStatusEnum `mandatory:"true" json:"status"`

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

func (m WorkRequestSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WorkRequestSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingWorkRequestSummaryOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetWorkRequestSummaryOperationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingWorkRequestSummaryStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetWorkRequestSummaryStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// WorkRequestSummaryOperationTypeEnum Enum with underlying type: string
type WorkRequestSummaryOperationTypeEnum string

// Set of constants representing the allowable values for WorkRequestSummaryOperationTypeEnum
const (
	WorkRequestSummaryOperationTypeCreateWorkspace WorkRequestSummaryOperationTypeEnum = "CREATE_WORKSPACE"
	WorkRequestSummaryOperationTypeUpdateWorkspace WorkRequestSummaryOperationTypeEnum = "UPDATE_WORKSPACE"
	WorkRequestSummaryOperationTypeDeleteWorkspace WorkRequestSummaryOperationTypeEnum = "DELETE_WORKSPACE"
	WorkRequestSummaryOperationTypeMoveWorkspace   WorkRequestSummaryOperationTypeEnum = "MOVE_WORKSPACE"
)

var mappingWorkRequestSummaryOperationTypeEnum = map[string]WorkRequestSummaryOperationTypeEnum{
	"CREATE_WORKSPACE": WorkRequestSummaryOperationTypeCreateWorkspace,
	"UPDATE_WORKSPACE": WorkRequestSummaryOperationTypeUpdateWorkspace,
	"DELETE_WORKSPACE": WorkRequestSummaryOperationTypeDeleteWorkspace,
	"MOVE_WORKSPACE":   WorkRequestSummaryOperationTypeMoveWorkspace,
}

var mappingWorkRequestSummaryOperationTypeEnumLowerCase = map[string]WorkRequestSummaryOperationTypeEnum{
	"create_workspace": WorkRequestSummaryOperationTypeCreateWorkspace,
	"update_workspace": WorkRequestSummaryOperationTypeUpdateWorkspace,
	"delete_workspace": WorkRequestSummaryOperationTypeDeleteWorkspace,
	"move_workspace":   WorkRequestSummaryOperationTypeMoveWorkspace,
}

// GetWorkRequestSummaryOperationTypeEnumValues Enumerates the set of values for WorkRequestSummaryOperationTypeEnum
func GetWorkRequestSummaryOperationTypeEnumValues() []WorkRequestSummaryOperationTypeEnum {
	values := make([]WorkRequestSummaryOperationTypeEnum, 0)
	for _, v := range mappingWorkRequestSummaryOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestSummaryOperationTypeEnumStringValues Enumerates the set of values in String for WorkRequestSummaryOperationTypeEnum
func GetWorkRequestSummaryOperationTypeEnumStringValues() []string {
	return []string{
		"CREATE_WORKSPACE",
		"UPDATE_WORKSPACE",
		"DELETE_WORKSPACE",
		"MOVE_WORKSPACE",
	}
}

// GetMappingWorkRequestSummaryOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestSummaryOperationTypeEnum(val string) (WorkRequestSummaryOperationTypeEnum, bool) {
	enum, ok := mappingWorkRequestSummaryOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// WorkRequestSummaryStatusEnum Enum with underlying type: string
type WorkRequestSummaryStatusEnum string

// Set of constants representing the allowable values for WorkRequestSummaryStatusEnum
const (
	WorkRequestSummaryStatusAccepted   WorkRequestSummaryStatusEnum = "ACCEPTED"
	WorkRequestSummaryStatusInProgress WorkRequestSummaryStatusEnum = "IN_PROGRESS"
	WorkRequestSummaryStatusFailed     WorkRequestSummaryStatusEnum = "FAILED"
	WorkRequestSummaryStatusSucceeded  WorkRequestSummaryStatusEnum = "SUCCEEDED"
	WorkRequestSummaryStatusCanceling  WorkRequestSummaryStatusEnum = "CANCELING"
	WorkRequestSummaryStatusCanceled   WorkRequestSummaryStatusEnum = "CANCELED"
)

var mappingWorkRequestSummaryStatusEnum = map[string]WorkRequestSummaryStatusEnum{
	"ACCEPTED":    WorkRequestSummaryStatusAccepted,
	"IN_PROGRESS": WorkRequestSummaryStatusInProgress,
	"FAILED":      WorkRequestSummaryStatusFailed,
	"SUCCEEDED":   WorkRequestSummaryStatusSucceeded,
	"CANCELING":   WorkRequestSummaryStatusCanceling,
	"CANCELED":    WorkRequestSummaryStatusCanceled,
}

var mappingWorkRequestSummaryStatusEnumLowerCase = map[string]WorkRequestSummaryStatusEnum{
	"accepted":    WorkRequestSummaryStatusAccepted,
	"in_progress": WorkRequestSummaryStatusInProgress,
	"failed":      WorkRequestSummaryStatusFailed,
	"succeeded":   WorkRequestSummaryStatusSucceeded,
	"canceling":   WorkRequestSummaryStatusCanceling,
	"canceled":    WorkRequestSummaryStatusCanceled,
}

// GetWorkRequestSummaryStatusEnumValues Enumerates the set of values for WorkRequestSummaryStatusEnum
func GetWorkRequestSummaryStatusEnumValues() []WorkRequestSummaryStatusEnum {
	values := make([]WorkRequestSummaryStatusEnum, 0)
	for _, v := range mappingWorkRequestSummaryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestSummaryStatusEnumStringValues Enumerates the set of values in String for WorkRequestSummaryStatusEnum
func GetWorkRequestSummaryStatusEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingWorkRequestSummaryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestSummaryStatusEnum(val string) (WorkRequestSummaryStatusEnum, bool) {
	enum, ok := mappingWorkRequestSummaryStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
