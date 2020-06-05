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

// WorkRequestSummary A description of the work request's status.
type WorkRequestSummary struct {

	// The identifier of the work request.
	Id *string `mandatory:"true" json:"id"`

	// The identifier of the compartment that contains the work request.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The identifier of the Digital Assistant instance to which this work request pertains.
	OdaInstanceId *string `mandatory:"true" json:"odaInstanceId"`

	// The type of the operation that's associated with the work request.
	RequestAction WorkRequestSummaryRequestActionEnum `mandatory:"true" json:"requestAction"`

	// The status of current work request.
	Status WorkRequestSummaryStatusEnum `mandatory:"true" json:"status"`

	// The resources that this work request affects.
	Resources []WorkRequestResource `mandatory:"true" json:"resources"`
}

func (m WorkRequestSummary) String() string {
	return common.PointerString(m)
}

// WorkRequestSummaryRequestActionEnum Enum with underlying type: string
type WorkRequestSummaryRequestActionEnum string

// Set of constants representing the allowable values for WorkRequestSummaryRequestActionEnum
const (
	WorkRequestSummaryRequestActionCreateOdaInstance            WorkRequestSummaryRequestActionEnum = "CREATE_ODA_INSTANCE"
	WorkRequestSummaryRequestActionUpgradeOdaInstance           WorkRequestSummaryRequestActionEnum = "UPGRADE_ODA_INSTANCE"
	WorkRequestSummaryRequestActionDeleteOdaInstance            WorkRequestSummaryRequestActionEnum = "DELETE_ODA_INSTANCE"
	WorkRequestSummaryRequestActionPurgeOdaInstance             WorkRequestSummaryRequestActionEnum = "PURGE_ODA_INSTANCE"
	WorkRequestSummaryRequestActionRecoverOdaInstance           WorkRequestSummaryRequestActionEnum = "RECOVER_ODA_INSTANCE"
	WorkRequestSummaryRequestActionStopOdaInstance              WorkRequestSummaryRequestActionEnum = "STOP_ODA_INSTANCE"
	WorkRequestSummaryRequestActionStartOdaInstance             WorkRequestSummaryRequestActionEnum = "START_ODA_INSTANCE"
	WorkRequestSummaryRequestActionChangeOdaInstanceCompartment WorkRequestSummaryRequestActionEnum = "CHANGE_ODA_INSTANCE_COMPARTMENT"
	WorkRequestSummaryRequestActionCreateAssociation            WorkRequestSummaryRequestActionEnum = "CREATE_ASSOCIATION"
	WorkRequestSummaryRequestActionDeleteAssociation            WorkRequestSummaryRequestActionEnum = "DELETE_ASSOCIATION"
	WorkRequestSummaryRequestActionUpdateEntitlementsForCacct   WorkRequestSummaryRequestActionEnum = "UPDATE_ENTITLEMENTS_FOR_CACCT"
	WorkRequestSummaryRequestActionLookupOdaInstancesForCacct   WorkRequestSummaryRequestActionEnum = "LOOKUP_ODA_INSTANCES_FOR_CACCT"
)

var mappingWorkRequestSummaryRequestAction = map[string]WorkRequestSummaryRequestActionEnum{
	"CREATE_ODA_INSTANCE":             WorkRequestSummaryRequestActionCreateOdaInstance,
	"UPGRADE_ODA_INSTANCE":            WorkRequestSummaryRequestActionUpgradeOdaInstance,
	"DELETE_ODA_INSTANCE":             WorkRequestSummaryRequestActionDeleteOdaInstance,
	"PURGE_ODA_INSTANCE":              WorkRequestSummaryRequestActionPurgeOdaInstance,
	"RECOVER_ODA_INSTANCE":            WorkRequestSummaryRequestActionRecoverOdaInstance,
	"STOP_ODA_INSTANCE":               WorkRequestSummaryRequestActionStopOdaInstance,
	"START_ODA_INSTANCE":              WorkRequestSummaryRequestActionStartOdaInstance,
	"CHANGE_ODA_INSTANCE_COMPARTMENT": WorkRequestSummaryRequestActionChangeOdaInstanceCompartment,
	"CREATE_ASSOCIATION":              WorkRequestSummaryRequestActionCreateAssociation,
	"DELETE_ASSOCIATION":              WorkRequestSummaryRequestActionDeleteAssociation,
	"UPDATE_ENTITLEMENTS_FOR_CACCT":   WorkRequestSummaryRequestActionUpdateEntitlementsForCacct,
	"LOOKUP_ODA_INSTANCES_FOR_CACCT":  WorkRequestSummaryRequestActionLookupOdaInstancesForCacct,
}

// GetWorkRequestSummaryRequestActionEnumValues Enumerates the set of values for WorkRequestSummaryRequestActionEnum
func GetWorkRequestSummaryRequestActionEnumValues() []WorkRequestSummaryRequestActionEnum {
	values := make([]WorkRequestSummaryRequestActionEnum, 0)
	for _, v := range mappingWorkRequestSummaryRequestAction {
		values = append(values, v)
	}
	return values
}

// WorkRequestSummaryStatusEnum Enum with underlying type: string
type WorkRequestSummaryStatusEnum string

// Set of constants representing the allowable values for WorkRequestSummaryStatusEnum
const (
	WorkRequestSummaryStatusAccepted   WorkRequestSummaryStatusEnum = "ACCEPTED"
	WorkRequestSummaryStatusInProgress WorkRequestSummaryStatusEnum = "IN_PROGRESS"
	WorkRequestSummaryStatusSucceeded  WorkRequestSummaryStatusEnum = "SUCCEEDED"
	WorkRequestSummaryStatusFailed     WorkRequestSummaryStatusEnum = "FAILED"
	WorkRequestSummaryStatusCanceling  WorkRequestSummaryStatusEnum = "CANCELING"
	WorkRequestSummaryStatusCanceled   WorkRequestSummaryStatusEnum = "CANCELED"
)

var mappingWorkRequestSummaryStatus = map[string]WorkRequestSummaryStatusEnum{
	"ACCEPTED":    WorkRequestSummaryStatusAccepted,
	"IN_PROGRESS": WorkRequestSummaryStatusInProgress,
	"SUCCEEDED":   WorkRequestSummaryStatusSucceeded,
	"FAILED":      WorkRequestSummaryStatusFailed,
	"CANCELING":   WorkRequestSummaryStatusCanceling,
	"CANCELED":    WorkRequestSummaryStatusCanceled,
}

// GetWorkRequestSummaryStatusEnumValues Enumerates the set of values for WorkRequestSummaryStatusEnum
func GetWorkRequestSummaryStatusEnumValues() []WorkRequestSummaryStatusEnum {
	values := make([]WorkRequestSummaryStatusEnum, 0)
	for _, v := range mappingWorkRequestSummaryStatus {
		values = append(values, v)
	}
	return values
}
