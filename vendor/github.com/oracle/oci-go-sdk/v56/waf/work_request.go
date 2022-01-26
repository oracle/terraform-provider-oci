// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Firewall (WAF) API
//
// API for the Web Application Firewall service.
// Use this API to manage regional Web App Firewalls and corresponding policies for protecting HTTP services.
//

package waf

import (
	"github.com/oracle/oci-go-sdk/v56/common"
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

// WorkRequestOperationTypeEnum Enum with underlying type: string
type WorkRequestOperationTypeEnum string

// Set of constants representing the allowable values for WorkRequestOperationTypeEnum
const (
	WorkRequestOperationTypeCreateWafPolicy          WorkRequestOperationTypeEnum = "CREATE_WAF_POLICY"
	WorkRequestOperationTypeUpdateWafPolicy          WorkRequestOperationTypeEnum = "UPDATE_WAF_POLICY"
	WorkRequestOperationTypeDeleteWafPolicy          WorkRequestOperationTypeEnum = "DELETE_WAF_POLICY"
	WorkRequestOperationTypeMoveWafPolicy            WorkRequestOperationTypeEnum = "MOVE_WAF_POLICY"
	WorkRequestOperationTypeCreateNetworkAddressList WorkRequestOperationTypeEnum = "CREATE_NETWORK_ADDRESS_LIST"
	WorkRequestOperationTypeUpdateNetworkAddressList WorkRequestOperationTypeEnum = "UPDATE_NETWORK_ADDRESS_LIST"
	WorkRequestOperationTypeDeleteNetworkAddressList WorkRequestOperationTypeEnum = "DELETE_NETWORK_ADDRESS_LIST"
	WorkRequestOperationTypeMoveNetworkAddressList   WorkRequestOperationTypeEnum = "MOVE_NETWORK_ADDRESS_LIST"
	WorkRequestOperationTypeCreateWebAppFirewall     WorkRequestOperationTypeEnum = "CREATE_WEB_APP_FIREWALL"
	WorkRequestOperationTypeUpdateWebAppFirewall     WorkRequestOperationTypeEnum = "UPDATE_WEB_APP_FIREWALL"
	WorkRequestOperationTypeDeleteWebAppFirewall     WorkRequestOperationTypeEnum = "DELETE_WEB_APP_FIREWALL"
	WorkRequestOperationTypeMoveWebAppFirewall       WorkRequestOperationTypeEnum = "MOVE_WEB_APP_FIREWALL"
)

var mappingWorkRequestOperationType = map[string]WorkRequestOperationTypeEnum{
	"CREATE_WAF_POLICY":           WorkRequestOperationTypeCreateWafPolicy,
	"UPDATE_WAF_POLICY":           WorkRequestOperationTypeUpdateWafPolicy,
	"DELETE_WAF_POLICY":           WorkRequestOperationTypeDeleteWafPolicy,
	"MOVE_WAF_POLICY":             WorkRequestOperationTypeMoveWafPolicy,
	"CREATE_NETWORK_ADDRESS_LIST": WorkRequestOperationTypeCreateNetworkAddressList,
	"UPDATE_NETWORK_ADDRESS_LIST": WorkRequestOperationTypeUpdateNetworkAddressList,
	"DELETE_NETWORK_ADDRESS_LIST": WorkRequestOperationTypeDeleteNetworkAddressList,
	"MOVE_NETWORK_ADDRESS_LIST":   WorkRequestOperationTypeMoveNetworkAddressList,
	"CREATE_WEB_APP_FIREWALL":     WorkRequestOperationTypeCreateWebAppFirewall,
	"UPDATE_WEB_APP_FIREWALL":     WorkRequestOperationTypeUpdateWebAppFirewall,
	"DELETE_WEB_APP_FIREWALL":     WorkRequestOperationTypeDeleteWebAppFirewall,
	"MOVE_WEB_APP_FIREWALL":       WorkRequestOperationTypeMoveWebAppFirewall,
}

// GetWorkRequestOperationTypeEnumValues Enumerates the set of values for WorkRequestOperationTypeEnum
func GetWorkRequestOperationTypeEnumValues() []WorkRequestOperationTypeEnum {
	values := make([]WorkRequestOperationTypeEnum, 0)
	for _, v := range mappingWorkRequestOperationType {
		values = append(values, v)
	}
	return values
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

var mappingWorkRequestStatus = map[string]WorkRequestStatusEnum{
	"ACCEPTED":    WorkRequestStatusAccepted,
	"IN_PROGRESS": WorkRequestStatusInProgress,
	"FAILED":      WorkRequestStatusFailed,
	"SUCCEEDED":   WorkRequestStatusSucceeded,
	"CANCELING":   WorkRequestStatusCanceling,
	"CANCELED":    WorkRequestStatusCanceled,
}

// GetWorkRequestStatusEnumValues Enumerates the set of values for WorkRequestStatusEnum
func GetWorkRequestStatusEnumValues() []WorkRequestStatusEnum {
	values := make([]WorkRequestStatusEnum, 0)
	for _, v := range mappingWorkRequestStatus {
		values = append(values, v)
	}
	return values
}
