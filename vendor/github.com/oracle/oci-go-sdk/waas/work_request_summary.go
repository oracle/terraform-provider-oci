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

// WorkRequestSummary The summarized details of a work request.
type WorkRequestSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the work request.
	Id *string `mandatory:"true" json:"id"`

	// A description of the operation requested by the work request.
	OperationType WorkRequestOperationTypesEnum `mandatory:"true" json:"operationType"`

	// The current status of the work request.
	Status WorkRequestStatusValuesEnum `mandatory:"true" json:"status"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment that contains the work request.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time the work request was created, expressed in RFC 3339 timestamp format.
	TimeAccepted *common.SDKTime `mandatory:"true" json:"timeAccepted"`

	// The date and time the work request moved from the `ACCEPTED` state to the `IN_PROGRESS` state, expressed in RFC 3339 timestamp format.
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// The date and time the work request was fulfilled or terminated, in the format defined by RFC3339.
	TimeFinished *common.SDKTime `mandatory:"true" json:"timeFinished"`

	// The resources being used to complete the work request operation.
	Resources []WorkRequestResource `mandatory:"false" json:"resources"`

	// The percentage of work completed by the work request.
	PercentComplete *int `mandatory:"false" json:"percentComplete"`
}

func (m WorkRequestSummary) String() string {
	return common.PointerString(m)
}

// WorkRequestSummaryOperationTypeEnum is an alias to type: WorkRequestOperationTypesEnum
// Consider using WorkRequestOperationTypesEnum instead
// Deprecated
type WorkRequestSummaryOperationTypeEnum = WorkRequestOperationTypesEnum

// Set of constants representing the allowable values for WorkRequestOperationTypesEnum
// Deprecated
const (
	WorkRequestSummaryOperationTypeCreateWaasPolicy           WorkRequestOperationTypesEnum = "CREATE_WAAS_POLICY"
	WorkRequestSummaryOperationTypeUpdateWaasPolicy           WorkRequestOperationTypesEnum = "UPDATE_WAAS_POLICY"
	WorkRequestSummaryOperationTypeDeleteWaasPolicy           WorkRequestOperationTypesEnum = "DELETE_WAAS_POLICY"
	WorkRequestSummaryOperationTypeCreateHttpRedirect         WorkRequestOperationTypesEnum = "CREATE_HTTP_REDIRECT"
	WorkRequestSummaryOperationTypeUpdateHttpRedirect         WorkRequestOperationTypesEnum = "UPDATE_HTTP_REDIRECT"
	WorkRequestSummaryOperationTypeDeleteHttpRedirect         WorkRequestOperationTypesEnum = "DELETE_HTTP_REDIRECT"
	WorkRequestSummaryOperationTypePurgeWaasPolicyCache       WorkRequestOperationTypesEnum = "PURGE_WAAS_POLICY_CACHE"
	WorkRequestSummaryOperationTypeCreateCustomProtectionRule WorkRequestOperationTypesEnum = "CREATE_CUSTOM_PROTECTION_RULE"
	WorkRequestSummaryOperationTypeUpdateCustomProtectionRule WorkRequestOperationTypesEnum = "UPDATE_CUSTOM_PROTECTION_RULE"
	WorkRequestSummaryOperationTypeDeleteCustomProtectionRule WorkRequestOperationTypesEnum = "DELETE_CUSTOM_PROTECTION_RULE"
)

// GetWorkRequestSummaryOperationTypeEnumValues Enumerates the set of values for WorkRequestOperationTypesEnum
// Consider using GetWorkRequestOperationTypesEnumValue
// Deprecated
var GetWorkRequestSummaryOperationTypeEnumValues = GetWorkRequestOperationTypesEnumValues

// WorkRequestSummaryStatusEnum is an alias to type: WorkRequestStatusValuesEnum
// Consider using WorkRequestStatusValuesEnum instead
// Deprecated
type WorkRequestSummaryStatusEnum = WorkRequestStatusValuesEnum

// Set of constants representing the allowable values for WorkRequestStatusValuesEnum
// Deprecated
const (
	WorkRequestSummaryStatusAccepted   WorkRequestStatusValuesEnum = "ACCEPTED"
	WorkRequestSummaryStatusInProgress WorkRequestStatusValuesEnum = "IN_PROGRESS"
	WorkRequestSummaryStatusFailed     WorkRequestStatusValuesEnum = "FAILED"
	WorkRequestSummaryStatusSucceeded  WorkRequestStatusValuesEnum = "SUCCEEDED"
	WorkRequestSummaryStatusCanceling  WorkRequestStatusValuesEnum = "CANCELING"
	WorkRequestSummaryStatusCanceled   WorkRequestStatusValuesEnum = "CANCELED"
)

// GetWorkRequestSummaryStatusEnumValues Enumerates the set of values for WorkRequestStatusValuesEnum
// Consider using GetWorkRequestStatusValuesEnumValue
// Deprecated
var GetWorkRequestSummaryStatusEnumValues = GetWorkRequestStatusValuesEnumValues
