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

// WorkRequest The description of work request, including its status.
type WorkRequest struct {

	// The identifier of the work request.
	Id *string `mandatory:"true" json:"id"`

	// The identifier of the compartment that contains the work request.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The identifier of the Digital Assistant instance to which this work request pertains.
	OdaInstanceId *string `mandatory:"true" json:"odaInstanceId"`

	// The type of the operation that's associated with the work request.
	RequestAction WorkRequestRequestActionEnum `mandatory:"true" json:"requestAction"`

	// The status of current work request.
	Status WorkRequestStatusEnum `mandatory:"true" json:"status"`

	// The resources that this work request affects.
	Resources []WorkRequestResource `mandatory:"true" json:"resources"`

	// Percentage of the request completed.
	PercentComplete *float32 `mandatory:"true" json:"percentComplete"`

	// The date and time that the request was created, as described in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeAccepted *common.SDKTime `mandatory:"true" json:"timeAccepted"`

	// A short message that provides more detail about the current status.
	// For example, if a work request fails, then this may include information
	// about why it failed.
	StatusMessage *string `mandatory:"false" json:"statusMessage"`

	// The date and time that the request was started, as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339), CKQ
	// section 14.29.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time that the object finished, as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339). CKQ
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`
}

func (m WorkRequest) String() string {
	return common.PointerString(m)
}

// WorkRequestRequestActionEnum Enum with underlying type: string
type WorkRequestRequestActionEnum string

// Set of constants representing the allowable values for WorkRequestRequestActionEnum
const (
	WorkRequestRequestActionCreateOdaInstance            WorkRequestRequestActionEnum = "CREATE_ODA_INSTANCE"
	WorkRequestRequestActionUpgradeOdaInstance           WorkRequestRequestActionEnum = "UPGRADE_ODA_INSTANCE"
	WorkRequestRequestActionDeleteOdaInstance            WorkRequestRequestActionEnum = "DELETE_ODA_INSTANCE"
	WorkRequestRequestActionPurgeOdaInstance             WorkRequestRequestActionEnum = "PURGE_ODA_INSTANCE"
	WorkRequestRequestActionRecoverOdaInstance           WorkRequestRequestActionEnum = "RECOVER_ODA_INSTANCE"
	WorkRequestRequestActionStopOdaInstance              WorkRequestRequestActionEnum = "STOP_ODA_INSTANCE"
	WorkRequestRequestActionStartOdaInstance             WorkRequestRequestActionEnum = "START_ODA_INSTANCE"
	WorkRequestRequestActionChangeOdaInstanceCompartment WorkRequestRequestActionEnum = "CHANGE_ODA_INSTANCE_COMPARTMENT"
	WorkRequestRequestActionCreateAssociation            WorkRequestRequestActionEnum = "CREATE_ASSOCIATION"
	WorkRequestRequestActionDeleteAssociation            WorkRequestRequestActionEnum = "DELETE_ASSOCIATION"
	WorkRequestRequestActionUpdateEntitlementsForCacct   WorkRequestRequestActionEnum = "UPDATE_ENTITLEMENTS_FOR_CACCT"
	WorkRequestRequestActionLookupOdaInstancesForCacct   WorkRequestRequestActionEnum = "LOOKUP_ODA_INSTANCES_FOR_CACCT"
)

var mappingWorkRequestRequestAction = map[string]WorkRequestRequestActionEnum{
	"CREATE_ODA_INSTANCE":             WorkRequestRequestActionCreateOdaInstance,
	"UPGRADE_ODA_INSTANCE":            WorkRequestRequestActionUpgradeOdaInstance,
	"DELETE_ODA_INSTANCE":             WorkRequestRequestActionDeleteOdaInstance,
	"PURGE_ODA_INSTANCE":              WorkRequestRequestActionPurgeOdaInstance,
	"RECOVER_ODA_INSTANCE":            WorkRequestRequestActionRecoverOdaInstance,
	"STOP_ODA_INSTANCE":               WorkRequestRequestActionStopOdaInstance,
	"START_ODA_INSTANCE":              WorkRequestRequestActionStartOdaInstance,
	"CHANGE_ODA_INSTANCE_COMPARTMENT": WorkRequestRequestActionChangeOdaInstanceCompartment,
	"CREATE_ASSOCIATION":              WorkRequestRequestActionCreateAssociation,
	"DELETE_ASSOCIATION":              WorkRequestRequestActionDeleteAssociation,
	"UPDATE_ENTITLEMENTS_FOR_CACCT":   WorkRequestRequestActionUpdateEntitlementsForCacct,
	"LOOKUP_ODA_INSTANCES_FOR_CACCT":  WorkRequestRequestActionLookupOdaInstancesForCacct,
}

// GetWorkRequestRequestActionEnumValues Enumerates the set of values for WorkRequestRequestActionEnum
func GetWorkRequestRequestActionEnumValues() []WorkRequestRequestActionEnum {
	values := make([]WorkRequestRequestActionEnum, 0)
	for _, v := range mappingWorkRequestRequestAction {
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
	WorkRequestStatusSucceeded  WorkRequestStatusEnum = "SUCCEEDED"
	WorkRequestStatusFailed     WorkRequestStatusEnum = "FAILED"
	WorkRequestStatusCanceling  WorkRequestStatusEnum = "CANCELING"
	WorkRequestStatusCanceled   WorkRequestStatusEnum = "CANCELED"
)

var mappingWorkRequestStatus = map[string]WorkRequestStatusEnum{
	"ACCEPTED":    WorkRequestStatusAccepted,
	"IN_PROGRESS": WorkRequestStatusInProgress,
	"SUCCEEDED":   WorkRequestStatusSucceeded,
	"FAILED":      WorkRequestStatusFailed,
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
