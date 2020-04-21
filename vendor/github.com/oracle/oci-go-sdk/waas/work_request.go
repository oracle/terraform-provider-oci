// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"github.com/oracle/oci-go-sdk/common"
)

// WorkRequest Many of the API requests you use to create and configure WAAS policies do not take effect immediately. In these cases, the request spawns an asynchronous work flow to fulfill the request. `WorkRequest` objects provide visibility for in-progress work flows. For more information about work requests, see Viewing the State of a Work Request (https://docs.cloud.oracle.com/Content/Balance/Tasks/viewingworkrequest.htm).
type WorkRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the work request.
	Id *string `mandatory:"true" json:"id"`

	// A description of the operation requested by the work request.
	OperationType WorkRequestOperationTypesEnum `mandatory:"true" json:"operationType"`

	// The current status of the work request.
	Status WorkRequestStatusValuesEnum `mandatory:"true" json:"status"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment that contains the work request.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time the work request was created, in the format defined by RFC3339.
	TimeAccepted *common.SDKTime `mandatory:"true" json:"timeAccepted"`

	// The date and time the work request moved from the `ACCEPTED` state to the `IN_PROGRESS` state, expressed in RFC 3339 timestamp format.
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// The date and time the work request was fulfilled or terminated, expressed in RFC 3339 timestamp format.
	TimeFinished *common.SDKTime `mandatory:"true" json:"timeFinished"`

	// The resources being used to complete the work request operation.
	Resources []WorkRequestResource `mandatory:"false" json:"resources"`

	// The percentage of work completed by the work request.
	PercentComplete *int `mandatory:"false" json:"percentComplete"`

	// The list of log entries from the work request workflow.
	Logs []WorkRequestLogEntry `mandatory:"false" json:"logs"`

	// The list of errors that occurred while fulfilling the work request.
	Errors []WorkRequestError `mandatory:"false" json:"errors"`
}

func (m WorkRequest) String() string {
	return common.PointerString(m)
}

// WorkRequestOperationTypeEnum is an alias to type: WorkRequestOperationTypesEnum
// Consider using WorkRequestOperationTypesEnum instead
// Deprecated
type WorkRequestOperationTypeEnum = WorkRequestOperationTypesEnum

// Set of constants representing the allowable values for WorkRequestOperationTypesEnum
// Deprecated
const (
	WorkRequestOperationTypeCreateWaasPolicy           WorkRequestOperationTypesEnum = "CREATE_WAAS_POLICY"
	WorkRequestOperationTypeUpdateWaasPolicy           WorkRequestOperationTypesEnum = "UPDATE_WAAS_POLICY"
	WorkRequestOperationTypeDeleteWaasPolicy           WorkRequestOperationTypesEnum = "DELETE_WAAS_POLICY"
	WorkRequestOperationTypeCreateHttpRedirect         WorkRequestOperationTypesEnum = "CREATE_HTTP_REDIRECT"
	WorkRequestOperationTypeUpdateHttpRedirect         WorkRequestOperationTypesEnum = "UPDATE_HTTP_REDIRECT"
	WorkRequestOperationTypeDeleteHttpRedirect         WorkRequestOperationTypesEnum = "DELETE_HTTP_REDIRECT"
	WorkRequestOperationTypePurgeWaasPolicyCache       WorkRequestOperationTypesEnum = "PURGE_WAAS_POLICY_CACHE"
	WorkRequestOperationTypeCreateCustomProtectionRule WorkRequestOperationTypesEnum = "CREATE_CUSTOM_PROTECTION_RULE"
	WorkRequestOperationTypeUpdateCustomProtectionRule WorkRequestOperationTypesEnum = "UPDATE_CUSTOM_PROTECTION_RULE"
	WorkRequestOperationTypeDeleteCustomProtectionRule WorkRequestOperationTypesEnum = "DELETE_CUSTOM_PROTECTION_RULE"
)

// GetWorkRequestOperationTypeEnumValues Enumerates the set of values for WorkRequestOperationTypesEnum
// Consider using GetWorkRequestOperationTypesEnumValue
// Deprecated
var GetWorkRequestOperationTypeEnumValues = GetWorkRequestOperationTypesEnumValues

// WorkRequestStatusEnum is an alias to type: WorkRequestStatusValuesEnum
// Consider using WorkRequestStatusValuesEnum instead
// Deprecated
type WorkRequestStatusEnum = WorkRequestStatusValuesEnum

// Set of constants representing the allowable values for WorkRequestStatusValuesEnum
// Deprecated
const (
	WorkRequestStatusAccepted   WorkRequestStatusValuesEnum = "ACCEPTED"
	WorkRequestStatusInProgress WorkRequestStatusValuesEnum = "IN_PROGRESS"
	WorkRequestStatusFailed     WorkRequestStatusValuesEnum = "FAILED"
	WorkRequestStatusSucceeded  WorkRequestStatusValuesEnum = "SUCCEEDED"
	WorkRequestStatusCanceling  WorkRequestStatusValuesEnum = "CANCELING"
	WorkRequestStatusCanceled   WorkRequestStatusValuesEnum = "CANCELED"
)

// GetWorkRequestStatusEnumValues Enumerates the set of values for WorkRequestStatusValuesEnum
// Consider using GetWorkRequestStatusValuesEnumValue
// Deprecated
var GetWorkRequestStatusEnumValues = GetWorkRequestStatusValuesEnumValues
