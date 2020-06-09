// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Digital Assistant Service Instance API
//
// API to create and maintain Oracle Digital Assistant service instances.
//

package oda

import (
	"github.com/oracle/oci-go-sdk/common"
)

// WorkRequestResource A resource created or operated on by a work request.
type WorkRequestResource struct {

	// The action to take against the Digital Assistant instance.
	ResourceAction WorkRequestResourceResourceActionEnum `mandatory:"true" json:"resourceAction"`

	// The resource type that the work request affects.
	ResourceType *string `mandatory:"true" json:"resourceType"`

	// The identifier of the Digital Assistant instance that is the subject of the request.
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// The current state of the work request. The `SUCCEEDED`, `FAILED`, AND `CANCELED` states
	// correspond to the action being performed.
	Status WorkRequestResourceStatusEnum `mandatory:"true" json:"status"`

	// Short message providing more detail for the current status. For example, if an operation fails
	// this may include information about the reason for the failure and a possible resolution.
	StatusMessage *string `mandatory:"false" json:"statusMessage"`

	// The URI path that the user can do a GET on to access the resource metadata.
	ResourceUri *string `mandatory:"false" json:"resourceUri"`
}

func (m WorkRequestResource) String() string {
	return common.PointerString(m)
}

// WorkRequestResourceResourceActionEnum Enum with underlying type: string
type WorkRequestResourceResourceActionEnum string

// Set of constants representing the allowable values for WorkRequestResourceResourceActionEnum
const (
	WorkRequestResourceResourceActionCreate                     WorkRequestResourceResourceActionEnum = "CREATE"
	WorkRequestResourceResourceActionDelete                     WorkRequestResourceResourceActionEnum = "DELETE"
	WorkRequestResourceResourceActionPurge                      WorkRequestResourceResourceActionEnum = "PURGE"
	WorkRequestResourceResourceActionRecover                    WorkRequestResourceResourceActionEnum = "RECOVER"
	WorkRequestResourceResourceActionStop                       WorkRequestResourceResourceActionEnum = "STOP"
	WorkRequestResourceResourceActionStart                      WorkRequestResourceResourceActionEnum = "START"
	WorkRequestResourceResourceActionChangeCompartment          WorkRequestResourceResourceActionEnum = "CHANGE_COMPARTMENT"
	WorkRequestResourceResourceActionCreateAssociation          WorkRequestResourceResourceActionEnum = "CREATE_ASSOCIATION"
	WorkRequestResourceResourceActionDeleteAssociation          WorkRequestResourceResourceActionEnum = "DELETE_ASSOCIATION"
	WorkRequestResourceResourceActionUpdateEntitlementsForCacct WorkRequestResourceResourceActionEnum = "UPDATE_ENTITLEMENTS_FOR_CACCT"
)

var mappingWorkRequestResourceResourceAction = map[string]WorkRequestResourceResourceActionEnum{
	"CREATE":                        WorkRequestResourceResourceActionCreate,
	"DELETE":                        WorkRequestResourceResourceActionDelete,
	"PURGE":                         WorkRequestResourceResourceActionPurge,
	"RECOVER":                       WorkRequestResourceResourceActionRecover,
	"STOP":                          WorkRequestResourceResourceActionStop,
	"START":                         WorkRequestResourceResourceActionStart,
	"CHANGE_COMPARTMENT":            WorkRequestResourceResourceActionChangeCompartment,
	"CREATE_ASSOCIATION":            WorkRequestResourceResourceActionCreateAssociation,
	"DELETE_ASSOCIATION":            WorkRequestResourceResourceActionDeleteAssociation,
	"UPDATE_ENTITLEMENTS_FOR_CACCT": WorkRequestResourceResourceActionUpdateEntitlementsForCacct,
}

// GetWorkRequestResourceResourceActionEnumValues Enumerates the set of values for WorkRequestResourceResourceActionEnum
func GetWorkRequestResourceResourceActionEnumValues() []WorkRequestResourceResourceActionEnum {
	values := make([]WorkRequestResourceResourceActionEnum, 0)
	for _, v := range mappingWorkRequestResourceResourceAction {
		values = append(values, v)
	}
	return values
}

// WorkRequestResourceStatusEnum Enum with underlying type: string
type WorkRequestResourceStatusEnum string

// Set of constants representing the allowable values for WorkRequestResourceStatusEnum
const (
	WorkRequestResourceStatusAccepted   WorkRequestResourceStatusEnum = "ACCEPTED"
	WorkRequestResourceStatusInProgress WorkRequestResourceStatusEnum = "IN_PROGRESS"
	WorkRequestResourceStatusSucceeded  WorkRequestResourceStatusEnum = "SUCCEEDED"
	WorkRequestResourceStatusFailed     WorkRequestResourceStatusEnum = "FAILED"
	WorkRequestResourceStatusCanceling  WorkRequestResourceStatusEnum = "CANCELING"
	WorkRequestResourceStatusCanceled   WorkRequestResourceStatusEnum = "CANCELED"
)

var mappingWorkRequestResourceStatus = map[string]WorkRequestResourceStatusEnum{
	"ACCEPTED":    WorkRequestResourceStatusAccepted,
	"IN_PROGRESS": WorkRequestResourceStatusInProgress,
	"SUCCEEDED":   WorkRequestResourceStatusSucceeded,
	"FAILED":      WorkRequestResourceStatusFailed,
	"CANCELING":   WorkRequestResourceStatusCanceling,
	"CANCELED":    WorkRequestResourceStatusCanceled,
}

// GetWorkRequestResourceStatusEnumValues Enumerates the set of values for WorkRequestResourceStatusEnum
func GetWorkRequestResourceStatusEnumValues() []WorkRequestResourceStatusEnum {
	values := make([]WorkRequestResourceStatusEnum, 0)
	for _, v := range mappingWorkRequestResourceStatus {
		values = append(values, v)
	}
	return values
}
