// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Firewall (WAF) API
//
// API for the Web Application Firewall service.
// Use this API to manage regional Web App Firewalls and corresponding policies for protecting HTTP services.
//

package waf

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

var mappingWorkRequestOperationTypeEnum = map[string]WorkRequestOperationTypeEnum{
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

var mappingWorkRequestOperationTypeEnumLowerCase = map[string]WorkRequestOperationTypeEnum{
	"create_waf_policy":           WorkRequestOperationTypeCreateWafPolicy,
	"update_waf_policy":           WorkRequestOperationTypeUpdateWafPolicy,
	"delete_waf_policy":           WorkRequestOperationTypeDeleteWafPolicy,
	"move_waf_policy":             WorkRequestOperationTypeMoveWafPolicy,
	"create_network_address_list": WorkRequestOperationTypeCreateNetworkAddressList,
	"update_network_address_list": WorkRequestOperationTypeUpdateNetworkAddressList,
	"delete_network_address_list": WorkRequestOperationTypeDeleteNetworkAddressList,
	"move_network_address_list":   WorkRequestOperationTypeMoveNetworkAddressList,
	"create_web_app_firewall":     WorkRequestOperationTypeCreateWebAppFirewall,
	"update_web_app_firewall":     WorkRequestOperationTypeUpdateWebAppFirewall,
	"delete_web_app_firewall":     WorkRequestOperationTypeDeleteWebAppFirewall,
	"move_web_app_firewall":       WorkRequestOperationTypeMoveWebAppFirewall,
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
		"CREATE_WAF_POLICY",
		"UPDATE_WAF_POLICY",
		"DELETE_WAF_POLICY",
		"MOVE_WAF_POLICY",
		"CREATE_NETWORK_ADDRESS_LIST",
		"UPDATE_NETWORK_ADDRESS_LIST",
		"DELETE_NETWORK_ADDRESS_LIST",
		"MOVE_NETWORK_ADDRESS_LIST",
		"CREATE_WEB_APP_FIREWALL",
		"UPDATE_WEB_APP_FIREWALL",
		"DELETE_WEB_APP_FIREWALL",
		"MOVE_WEB_APP_FIREWALL",
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
