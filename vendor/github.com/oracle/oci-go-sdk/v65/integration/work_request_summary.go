// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Integration API
//
// Oracle Integration API.
//

package integration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// WorkRequestSummary A description of work request status.
type WorkRequestSummary struct {

	// Type of the work request.
	OperationType WorkRequestSummaryOperationTypeEnum `mandatory:"true" json:"operationType"`

	// Status of current work request.
	Status WorkRequestSummaryStatusEnum `mandatory:"true" json:"status"`

	// The id of the work request.
	Id *string `mandatory:"true" json:"id"`

	// The ocid of the compartment that contains the work request. Work
	// requests should be scoped to the same compartment as the resource the
	// work request affects. If the work request affects multiple resources,
	// and those resources are not in the same compartment, it is up to the
	// service team to pick the primary resource whose compartment should be
	// used.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The resources affected by this work request.
	Resources []WorkRequestResource `mandatory:"true" json:"resources"`

	// Percentage of the request completed.
	PercentComplete *float32 `mandatory:"true" json:"percentComplete"`

	// The date and time the request was created, as described in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeAccepted *common.SDKTime `mandatory:"true" json:"timeAccepted"`

	// The date and time the request was started, as described in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339),
	// section 14.29.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time the object was finished, as described in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
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
	WorkRequestSummaryOperationTypeCreateIntegrationInstance               WorkRequestSummaryOperationTypeEnum = "CREATE_INTEGRATION_INSTANCE"
	WorkRequestSummaryOperationTypeUpdateIntegrationInstance               WorkRequestSummaryOperationTypeEnum = "UPDATE_INTEGRATION_INSTANCE"
	WorkRequestSummaryOperationTypeStopIntegrationInstance                 WorkRequestSummaryOperationTypeEnum = "STOP_INTEGRATION_INSTANCE"
	WorkRequestSummaryOperationTypeStartIntegrationInstance                WorkRequestSummaryOperationTypeEnum = "START_INTEGRATION_INSTANCE"
	WorkRequestSummaryOperationTypeDeleteIntegrationInstance               WorkRequestSummaryOperationTypeEnum = "DELETE_INTEGRATION_INSTANCE"
	WorkRequestSummaryOperationTypeChangePrivateEndpointOutboundConnection WorkRequestSummaryOperationTypeEnum = "CHANGE_PRIVATE_ENDPOINT_OUTBOUND_CONNECTION"
	WorkRequestSummaryOperationTypeEnableProcessAutomation                 WorkRequestSummaryOperationTypeEnum = "ENABLE_PROCESS_AUTOMATION"
)

var mappingWorkRequestSummaryOperationTypeEnum = map[string]WorkRequestSummaryOperationTypeEnum{
	"CREATE_INTEGRATION_INSTANCE":                 WorkRequestSummaryOperationTypeCreateIntegrationInstance,
	"UPDATE_INTEGRATION_INSTANCE":                 WorkRequestSummaryOperationTypeUpdateIntegrationInstance,
	"STOP_INTEGRATION_INSTANCE":                   WorkRequestSummaryOperationTypeStopIntegrationInstance,
	"START_INTEGRATION_INSTANCE":                  WorkRequestSummaryOperationTypeStartIntegrationInstance,
	"DELETE_INTEGRATION_INSTANCE":                 WorkRequestSummaryOperationTypeDeleteIntegrationInstance,
	"CHANGE_PRIVATE_ENDPOINT_OUTBOUND_CONNECTION": WorkRequestSummaryOperationTypeChangePrivateEndpointOutboundConnection,
	"ENABLE_PROCESS_AUTOMATION":                   WorkRequestSummaryOperationTypeEnableProcessAutomation,
}

var mappingWorkRequestSummaryOperationTypeEnumLowerCase = map[string]WorkRequestSummaryOperationTypeEnum{
	"create_integration_instance":                 WorkRequestSummaryOperationTypeCreateIntegrationInstance,
	"update_integration_instance":                 WorkRequestSummaryOperationTypeUpdateIntegrationInstance,
	"stop_integration_instance":                   WorkRequestSummaryOperationTypeStopIntegrationInstance,
	"start_integration_instance":                  WorkRequestSummaryOperationTypeStartIntegrationInstance,
	"delete_integration_instance":                 WorkRequestSummaryOperationTypeDeleteIntegrationInstance,
	"change_private_endpoint_outbound_connection": WorkRequestSummaryOperationTypeChangePrivateEndpointOutboundConnection,
	"enable_process_automation":                   WorkRequestSummaryOperationTypeEnableProcessAutomation,
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
		"CREATE_INTEGRATION_INSTANCE",
		"UPDATE_INTEGRATION_INSTANCE",
		"STOP_INTEGRATION_INSTANCE",
		"START_INTEGRATION_INSTANCE",
		"DELETE_INTEGRATION_INSTANCE",
		"CHANGE_PRIVATE_ENDPOINT_OUTBOUND_CONNECTION",
		"ENABLE_PROCESS_AUTOMATION",
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
