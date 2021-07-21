// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"github.com/oracle/oci-go-sdk/v45/common"
)

// WorkRequest An asynchronous work request.
type WorkRequest struct {

	// The resources that are affected by the work request.
	OperationType WorkRequestOperationTypeEnum `mandatory:"true" json:"operationType"`

	// The current status of the work request.
	Status WorkRequestStatusEnum `mandatory:"true" json:"status"`

	// The OCID of the work request.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that contains the work request.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The resources that are affected by this work request.
	Resources []WorkRequestResource `mandatory:"true" json:"resources"`

	// Progress of the work request in percentage.
	PercentComplete *float32 `mandatory:"true" json:"percentComplete"`

	// The date and time the work request was accepted, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeAccepted *common.SDKTime `mandatory:"true" json:"timeAccepted"`

	// The date and time the work request transitioned from ACCEPTED to IN_PROGRESS, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time the work request reached a terminal state, either FAILED or SUCCEEDED. Format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`
}

func (m WorkRequest) String() string {
	return common.PointerString(m)
}

// WorkRequestOperationTypeEnum Enum with underlying type: string
type WorkRequestOperationTypeEnum string

// Set of constants representing the allowable values for WorkRequestOperationTypeEnum
const (
	WorkRequestOperationTypeEnableDataSafeConfiguration      WorkRequestOperationTypeEnum = "ENABLE_DATA_SAFE_CONFIGURATION"
	WorkRequestOperationTypeCreatePrivateEndpoint            WorkRequestOperationTypeEnum = "CREATE_PRIVATE_ENDPOINT"
	WorkRequestOperationTypeUpdatePrivateEndpoint            WorkRequestOperationTypeEnum = "UPDATE_PRIVATE_ENDPOINT"
	WorkRequestOperationTypeDeletePrivateEndpoint            WorkRequestOperationTypeEnum = "DELETE_PRIVATE_ENDPOINT"
	WorkRequestOperationTypeChangePrivateEndpointCompartment WorkRequestOperationTypeEnum = "CHANGE_PRIVATE_ENDPOINT_COMPARTMENT"
	WorkRequestOperationTypeCreateOnpremConnector            WorkRequestOperationTypeEnum = "CREATE_ONPREM_CONNECTOR"
	WorkRequestOperationTypeUpdateOnpremConnector            WorkRequestOperationTypeEnum = "UPDATE_ONPREM_CONNECTOR"
	WorkRequestOperationTypeDeleteOnpremConnector            WorkRequestOperationTypeEnum = "DELETE_ONPREM_CONNECTOR"
	WorkRequestOperationTypeUpdateOnpremConnectorWallet      WorkRequestOperationTypeEnum = "UPDATE_ONPREM_CONNECTOR_WALLET"
	WorkRequestOperationTypeChangeOnpremConnectorCompartment WorkRequestOperationTypeEnum = "CHANGE_ONPREM_CONNECTOR_COMPARTMENT"
	WorkRequestOperationTypeCreateTargetDatabase             WorkRequestOperationTypeEnum = "CREATE_TARGET_DATABASE"
	WorkRequestOperationTypeUpdateTargetDatabase             WorkRequestOperationTypeEnum = "UPDATE_TARGET_DATABASE"
	WorkRequestOperationTypeActivateTargetDatabase           WorkRequestOperationTypeEnum = "ACTIVATE_TARGET_DATABASE"
	WorkRequestOperationTypeDeactivateTargetDatabase         WorkRequestOperationTypeEnum = "DEACTIVATE_TARGET_DATABASE"
	WorkRequestOperationTypeDeleteTargetDatabase             WorkRequestOperationTypeEnum = "DELETE_TARGET_DATABASE"
	WorkRequestOperationTypeChangeTargetDatabaseCompartment  WorkRequestOperationTypeEnum = "CHANGE_TARGET_DATABASE_COMPARTMENT"
)

var mappingWorkRequestOperationType = map[string]WorkRequestOperationTypeEnum{
	"ENABLE_DATA_SAFE_CONFIGURATION":      WorkRequestOperationTypeEnableDataSafeConfiguration,
	"CREATE_PRIVATE_ENDPOINT":             WorkRequestOperationTypeCreatePrivateEndpoint,
	"UPDATE_PRIVATE_ENDPOINT":             WorkRequestOperationTypeUpdatePrivateEndpoint,
	"DELETE_PRIVATE_ENDPOINT":             WorkRequestOperationTypeDeletePrivateEndpoint,
	"CHANGE_PRIVATE_ENDPOINT_COMPARTMENT": WorkRequestOperationTypeChangePrivateEndpointCompartment,
	"CREATE_ONPREM_CONNECTOR":             WorkRequestOperationTypeCreateOnpremConnector,
	"UPDATE_ONPREM_CONNECTOR":             WorkRequestOperationTypeUpdateOnpremConnector,
	"DELETE_ONPREM_CONNECTOR":             WorkRequestOperationTypeDeleteOnpremConnector,
	"UPDATE_ONPREM_CONNECTOR_WALLET":      WorkRequestOperationTypeUpdateOnpremConnectorWallet,
	"CHANGE_ONPREM_CONNECTOR_COMPARTMENT": WorkRequestOperationTypeChangeOnpremConnectorCompartment,
	"CREATE_TARGET_DATABASE":              WorkRequestOperationTypeCreateTargetDatabase,
	"UPDATE_TARGET_DATABASE":              WorkRequestOperationTypeUpdateTargetDatabase,
	"ACTIVATE_TARGET_DATABASE":            WorkRequestOperationTypeActivateTargetDatabase,
	"DEACTIVATE_TARGET_DATABASE":          WorkRequestOperationTypeDeactivateTargetDatabase,
	"DELETE_TARGET_DATABASE":              WorkRequestOperationTypeDeleteTargetDatabase,
	"CHANGE_TARGET_DATABASE_COMPARTMENT":  WorkRequestOperationTypeChangeTargetDatabaseCompartment,
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
)

var mappingWorkRequestStatus = map[string]WorkRequestStatusEnum{
	"ACCEPTED":    WorkRequestStatusAccepted,
	"IN_PROGRESS": WorkRequestStatusInProgress,
	"FAILED":      WorkRequestStatusFailed,
	"SUCCEEDED":   WorkRequestStatusSucceeded,
}

// GetWorkRequestStatusEnumValues Enumerates the set of values for WorkRequestStatusEnum
func GetWorkRequestStatusEnumValues() []WorkRequestStatusEnum {
	values := make([]WorkRequestStatusEnum, 0)
	for _, v := range mappingWorkRequestStatus {
		values = append(values, v)
	}
	return values
}
