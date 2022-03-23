// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Digital Assistant Service Instance API
//
// API to create and maintain Oracle Digital Assistant service instances.
//

package oda

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v63/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WorkRequestSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingWorkRequestSummaryRequestActionEnum(string(m.RequestAction)); !ok && m.RequestAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RequestAction: %s. Supported values are: %s.", m.RequestAction, strings.Join(GetWorkRequestSummaryRequestActionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingWorkRequestSummaryStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetWorkRequestSummaryStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingWorkRequestSummaryRequestActionEnum = map[string]WorkRequestSummaryRequestActionEnum{
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

var mappingWorkRequestSummaryRequestActionEnumLowerCase = map[string]WorkRequestSummaryRequestActionEnum{
	"create_oda_instance":             WorkRequestSummaryRequestActionCreateOdaInstance,
	"upgrade_oda_instance":            WorkRequestSummaryRequestActionUpgradeOdaInstance,
	"delete_oda_instance":             WorkRequestSummaryRequestActionDeleteOdaInstance,
	"purge_oda_instance":              WorkRequestSummaryRequestActionPurgeOdaInstance,
	"recover_oda_instance":            WorkRequestSummaryRequestActionRecoverOdaInstance,
	"stop_oda_instance":               WorkRequestSummaryRequestActionStopOdaInstance,
	"start_oda_instance":              WorkRequestSummaryRequestActionStartOdaInstance,
	"change_oda_instance_compartment": WorkRequestSummaryRequestActionChangeOdaInstanceCompartment,
	"create_association":              WorkRequestSummaryRequestActionCreateAssociation,
	"delete_association":              WorkRequestSummaryRequestActionDeleteAssociation,
	"update_entitlements_for_cacct":   WorkRequestSummaryRequestActionUpdateEntitlementsForCacct,
	"lookup_oda_instances_for_cacct":  WorkRequestSummaryRequestActionLookupOdaInstancesForCacct,
}

// GetWorkRequestSummaryRequestActionEnumValues Enumerates the set of values for WorkRequestSummaryRequestActionEnum
func GetWorkRequestSummaryRequestActionEnumValues() []WorkRequestSummaryRequestActionEnum {
	values := make([]WorkRequestSummaryRequestActionEnum, 0)
	for _, v := range mappingWorkRequestSummaryRequestActionEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestSummaryRequestActionEnumStringValues Enumerates the set of values in String for WorkRequestSummaryRequestActionEnum
func GetWorkRequestSummaryRequestActionEnumStringValues() []string {
	return []string{
		"CREATE_ODA_INSTANCE",
		"UPGRADE_ODA_INSTANCE",
		"DELETE_ODA_INSTANCE",
		"PURGE_ODA_INSTANCE",
		"RECOVER_ODA_INSTANCE",
		"STOP_ODA_INSTANCE",
		"START_ODA_INSTANCE",
		"CHANGE_ODA_INSTANCE_COMPARTMENT",
		"CREATE_ASSOCIATION",
		"DELETE_ASSOCIATION",
		"UPDATE_ENTITLEMENTS_FOR_CACCT",
		"LOOKUP_ODA_INSTANCES_FOR_CACCT",
	}
}

// GetMappingWorkRequestSummaryRequestActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestSummaryRequestActionEnum(val string) (WorkRequestSummaryRequestActionEnum, bool) {
	enum, ok := mappingWorkRequestSummaryRequestActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
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

var mappingWorkRequestSummaryStatusEnum = map[string]WorkRequestSummaryStatusEnum{
	"ACCEPTED":    WorkRequestSummaryStatusAccepted,
	"IN_PROGRESS": WorkRequestSummaryStatusInProgress,
	"SUCCEEDED":   WorkRequestSummaryStatusSucceeded,
	"FAILED":      WorkRequestSummaryStatusFailed,
	"CANCELING":   WorkRequestSummaryStatusCanceling,
	"CANCELED":    WorkRequestSummaryStatusCanceled,
}

var mappingWorkRequestSummaryStatusEnumLowerCase = map[string]WorkRequestSummaryStatusEnum{
	"accepted":    WorkRequestSummaryStatusAccepted,
	"in_progress": WorkRequestSummaryStatusInProgress,
	"succeeded":   WorkRequestSummaryStatusSucceeded,
	"failed":      WorkRequestSummaryStatusFailed,
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
		"SUCCEEDED",
		"FAILED",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingWorkRequestSummaryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestSummaryStatusEnum(val string) (WorkRequestSummaryStatusEnum, bool) {
	enum, ok := mappingWorkRequestSummaryStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
