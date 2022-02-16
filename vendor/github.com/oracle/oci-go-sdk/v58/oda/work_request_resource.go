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
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WorkRequestResource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingWorkRequestResourceResourceActionEnum(string(m.ResourceAction)); !ok && m.ResourceAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceAction: %s. Supported values are: %s.", m.ResourceAction, strings.Join(GetWorkRequestResourceResourceActionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingWorkRequestResourceStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetWorkRequestResourceStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingWorkRequestResourceResourceActionEnum = map[string]WorkRequestResourceResourceActionEnum{
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
	for _, v := range mappingWorkRequestResourceResourceActionEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestResourceResourceActionEnumStringValues Enumerates the set of values in String for WorkRequestResourceResourceActionEnum
func GetWorkRequestResourceResourceActionEnumStringValues() []string {
	return []string{
		"CREATE",
		"DELETE",
		"PURGE",
		"RECOVER",
		"STOP",
		"START",
		"CHANGE_COMPARTMENT",
		"CREATE_ASSOCIATION",
		"DELETE_ASSOCIATION",
		"UPDATE_ENTITLEMENTS_FOR_CACCT",
	}
}

// GetMappingWorkRequestResourceResourceActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestResourceResourceActionEnum(val string) (WorkRequestResourceResourceActionEnum, bool) {
	mappingWorkRequestResourceResourceActionEnumIgnoreCase := make(map[string]WorkRequestResourceResourceActionEnum)
	for k, v := range mappingWorkRequestResourceResourceActionEnum {
		mappingWorkRequestResourceResourceActionEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingWorkRequestResourceResourceActionEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
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

var mappingWorkRequestResourceStatusEnum = map[string]WorkRequestResourceStatusEnum{
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
	for _, v := range mappingWorkRequestResourceStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestResourceStatusEnumStringValues Enumerates the set of values in String for WorkRequestResourceStatusEnum
func GetWorkRequestResourceStatusEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingWorkRequestResourceStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestResourceStatusEnum(val string) (WorkRequestResourceStatusEnum, bool) {
	mappingWorkRequestResourceStatusEnumIgnoreCase := make(map[string]WorkRequestResourceStatusEnum)
	for k, v := range mappingWorkRequestResourceStatusEnum {
		mappingWorkRequestResourceStatusEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingWorkRequestResourceStatusEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
