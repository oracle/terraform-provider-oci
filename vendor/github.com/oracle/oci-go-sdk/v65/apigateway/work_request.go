// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// API Gateway API
//
// API for the API Gateway service. Use this API to manage gateways, deployments, and related items.
// For more information, see
// Overview of API Gateway (https://docs.cloud.oracle.com/iaas/Content/APIGateway/Concepts/apigatewayoverview.htm).
//

package apigateway

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// WorkRequest A description of the work request status.
type WorkRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the resource.
	Id *string `mandatory:"true" json:"id"`

	// The type of the work request.
	OperationType WorkRequestOperationTypeEnum `mandatory:"true" json:"operationType"`

	// The status of the work request.
	Status WorkRequestStatusEnum `mandatory:"true" json:"status"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment in which the
	// resource is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The resources affected by this work request.
	Resources []WorkRequestResource `mandatory:"true" json:"resources"`

	// Percentage of the request completed.
	PercentComplete *float32 `mandatory:"true" json:"percentComplete"`

	// The date and time the request was created, as described in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeAccepted *common.SDKTime `mandatory:"true" json:"timeAccepted"`

	// The date and time the request was started, as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339),
	// section 14.29.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time the request was finished, as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
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
	WorkRequestOperationTypeCreateGateway     WorkRequestOperationTypeEnum = "CREATE_GATEWAY"
	WorkRequestOperationTypeUpdateGateway     WorkRequestOperationTypeEnum = "UPDATE_GATEWAY"
	WorkRequestOperationTypeDeleteGateway     WorkRequestOperationTypeEnum = "DELETE_GATEWAY"
	WorkRequestOperationTypeCreateDeployment  WorkRequestOperationTypeEnum = "CREATE_DEPLOYMENT"
	WorkRequestOperationTypeUpdateDeployment  WorkRequestOperationTypeEnum = "UPDATE_DEPLOYMENT"
	WorkRequestOperationTypeDeleteDeployment  WorkRequestOperationTypeEnum = "DELETE_DEPLOYMENT"
	WorkRequestOperationTypeCreateCertificate WorkRequestOperationTypeEnum = "CREATE_CERTIFICATE"
	WorkRequestOperationTypeUpdateCertificate WorkRequestOperationTypeEnum = "UPDATE_CERTIFICATE"
	WorkRequestOperationTypeDeleteCertificate WorkRequestOperationTypeEnum = "DELETE_CERTIFICATE"
	WorkRequestOperationTypeCreateApi         WorkRequestOperationTypeEnum = "CREATE_API"
	WorkRequestOperationTypeUpdateApi         WorkRequestOperationTypeEnum = "UPDATE_API"
	WorkRequestOperationTypeDeleteApi         WorkRequestOperationTypeEnum = "DELETE_API"
	WorkRequestOperationTypeValidateApi       WorkRequestOperationTypeEnum = "VALIDATE_API"
	WorkRequestOperationTypeCreateSdk         WorkRequestOperationTypeEnum = "CREATE_SDK"
	WorkRequestOperationTypeDeleteSdk         WorkRequestOperationTypeEnum = "DELETE_SDK"
	WorkRequestOperationTypeCreateUsagePlan   WorkRequestOperationTypeEnum = "CREATE_USAGE_PLAN"
	WorkRequestOperationTypeUpdateUsagePlan   WorkRequestOperationTypeEnum = "UPDATE_USAGE_PLAN"
	WorkRequestOperationTypeDeleteUsagePlan   WorkRequestOperationTypeEnum = "DELETE_USAGE_PLAN"
	WorkRequestOperationTypeCreateSubscriber  WorkRequestOperationTypeEnum = "CREATE_SUBSCRIBER"
	WorkRequestOperationTypeUpdateSubscriber  WorkRequestOperationTypeEnum = "UPDATE_SUBSCRIBER"
	WorkRequestOperationTypeDeleteSubscriber  WorkRequestOperationTypeEnum = "DELETE_SUBSCRIBER"
)

var mappingWorkRequestOperationTypeEnum = map[string]WorkRequestOperationTypeEnum{
	"CREATE_GATEWAY":     WorkRequestOperationTypeCreateGateway,
	"UPDATE_GATEWAY":     WorkRequestOperationTypeUpdateGateway,
	"DELETE_GATEWAY":     WorkRequestOperationTypeDeleteGateway,
	"CREATE_DEPLOYMENT":  WorkRequestOperationTypeCreateDeployment,
	"UPDATE_DEPLOYMENT":  WorkRequestOperationTypeUpdateDeployment,
	"DELETE_DEPLOYMENT":  WorkRequestOperationTypeDeleteDeployment,
	"CREATE_CERTIFICATE": WorkRequestOperationTypeCreateCertificate,
	"UPDATE_CERTIFICATE": WorkRequestOperationTypeUpdateCertificate,
	"DELETE_CERTIFICATE": WorkRequestOperationTypeDeleteCertificate,
	"CREATE_API":         WorkRequestOperationTypeCreateApi,
	"UPDATE_API":         WorkRequestOperationTypeUpdateApi,
	"DELETE_API":         WorkRequestOperationTypeDeleteApi,
	"VALIDATE_API":       WorkRequestOperationTypeValidateApi,
	"CREATE_SDK":         WorkRequestOperationTypeCreateSdk,
	"DELETE_SDK":         WorkRequestOperationTypeDeleteSdk,
	"CREATE_USAGE_PLAN":  WorkRequestOperationTypeCreateUsagePlan,
	"UPDATE_USAGE_PLAN":  WorkRequestOperationTypeUpdateUsagePlan,
	"DELETE_USAGE_PLAN":  WorkRequestOperationTypeDeleteUsagePlan,
	"CREATE_SUBSCRIBER":  WorkRequestOperationTypeCreateSubscriber,
	"UPDATE_SUBSCRIBER":  WorkRequestOperationTypeUpdateSubscriber,
	"DELETE_SUBSCRIBER":  WorkRequestOperationTypeDeleteSubscriber,
}

var mappingWorkRequestOperationTypeEnumLowerCase = map[string]WorkRequestOperationTypeEnum{
	"create_gateway":     WorkRequestOperationTypeCreateGateway,
	"update_gateway":     WorkRequestOperationTypeUpdateGateway,
	"delete_gateway":     WorkRequestOperationTypeDeleteGateway,
	"create_deployment":  WorkRequestOperationTypeCreateDeployment,
	"update_deployment":  WorkRequestOperationTypeUpdateDeployment,
	"delete_deployment":  WorkRequestOperationTypeDeleteDeployment,
	"create_certificate": WorkRequestOperationTypeCreateCertificate,
	"update_certificate": WorkRequestOperationTypeUpdateCertificate,
	"delete_certificate": WorkRequestOperationTypeDeleteCertificate,
	"create_api":         WorkRequestOperationTypeCreateApi,
	"update_api":         WorkRequestOperationTypeUpdateApi,
	"delete_api":         WorkRequestOperationTypeDeleteApi,
	"validate_api":       WorkRequestOperationTypeValidateApi,
	"create_sdk":         WorkRequestOperationTypeCreateSdk,
	"delete_sdk":         WorkRequestOperationTypeDeleteSdk,
	"create_usage_plan":  WorkRequestOperationTypeCreateUsagePlan,
	"update_usage_plan":  WorkRequestOperationTypeUpdateUsagePlan,
	"delete_usage_plan":  WorkRequestOperationTypeDeleteUsagePlan,
	"create_subscriber":  WorkRequestOperationTypeCreateSubscriber,
	"update_subscriber":  WorkRequestOperationTypeUpdateSubscriber,
	"delete_subscriber":  WorkRequestOperationTypeDeleteSubscriber,
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
		"CREATE_GATEWAY",
		"UPDATE_GATEWAY",
		"DELETE_GATEWAY",
		"CREATE_DEPLOYMENT",
		"UPDATE_DEPLOYMENT",
		"DELETE_DEPLOYMENT",
		"CREATE_CERTIFICATE",
		"UPDATE_CERTIFICATE",
		"DELETE_CERTIFICATE",
		"CREATE_API",
		"UPDATE_API",
		"DELETE_API",
		"VALIDATE_API",
		"CREATE_SDK",
		"DELETE_SDK",
		"CREATE_USAGE_PLAN",
		"UPDATE_USAGE_PLAN",
		"DELETE_USAGE_PLAN",
		"CREATE_SUBSCRIBER",
		"UPDATE_SUBSCRIBER",
		"DELETE_SUBSCRIBER",
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
