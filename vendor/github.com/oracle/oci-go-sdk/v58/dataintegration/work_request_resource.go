// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// WorkRequestResource The resource that is created or operated on by a work request.
type WorkRequestResource struct {

	// The resource type the work request affects.
	EntityType *string `mandatory:"true" json:"entityType"`

	// The way in which this resource is affected by the work tracked in the work request.
	// A resource being created, updated, or deleted will remain in the `IN_PROGRESS` state until
	// work is complete for that resource at which point it will transition to `CREATED`, `UPDATED`,
	// or `DELETED`, respectively.
	ActionType WorkRequestResourceActionTypeEnum `mandatory:"true" json:"actionType"`

	// The OCID or other unique identifier of the resource the work request affects.
	Identifier *string `mandatory:"true" json:"identifier"`

	// The URI path that is used in a GET request to access the resource metadata.
	EntityUri *string `mandatory:"false" json:"entityUri"`
}

func (m WorkRequestResource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WorkRequestResource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingWorkRequestResourceActionTypeEnum(string(m.ActionType)); !ok && m.ActionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActionType: %s. Supported values are: %s.", m.ActionType, strings.Join(GetWorkRequestResourceActionTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// WorkRequestResourceActionTypeEnum Enum with underlying type: string
type WorkRequestResourceActionTypeEnum string

// Set of constants representing the allowable values for WorkRequestResourceActionTypeEnum
const (
	WorkRequestResourceActionTypeCreated    WorkRequestResourceActionTypeEnum = "CREATED"
	WorkRequestResourceActionTypeUpdated    WorkRequestResourceActionTypeEnum = "UPDATED"
	WorkRequestResourceActionTypeDeleted    WorkRequestResourceActionTypeEnum = "DELETED"
	WorkRequestResourceActionTypeMoved      WorkRequestResourceActionTypeEnum = "MOVED"
	WorkRequestResourceActionTypeInProgress WorkRequestResourceActionTypeEnum = "IN_PROGRESS"
	WorkRequestResourceActionTypeFailed     WorkRequestResourceActionTypeEnum = "FAILED"
	WorkRequestResourceActionTypeStopped    WorkRequestResourceActionTypeEnum = "STOPPED"
	WorkRequestResourceActionTypeStarted    WorkRequestResourceActionTypeEnum = "STARTED"
)

var mappingWorkRequestResourceActionTypeEnum = map[string]WorkRequestResourceActionTypeEnum{
	"CREATED":     WorkRequestResourceActionTypeCreated,
	"UPDATED":     WorkRequestResourceActionTypeUpdated,
	"DELETED":     WorkRequestResourceActionTypeDeleted,
	"MOVED":       WorkRequestResourceActionTypeMoved,
	"IN_PROGRESS": WorkRequestResourceActionTypeInProgress,
	"FAILED":      WorkRequestResourceActionTypeFailed,
	"STOPPED":     WorkRequestResourceActionTypeStopped,
	"STARTED":     WorkRequestResourceActionTypeStarted,
}

// GetWorkRequestResourceActionTypeEnumValues Enumerates the set of values for WorkRequestResourceActionTypeEnum
func GetWorkRequestResourceActionTypeEnumValues() []WorkRequestResourceActionTypeEnum {
	values := make([]WorkRequestResourceActionTypeEnum, 0)
	for _, v := range mappingWorkRequestResourceActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestResourceActionTypeEnumStringValues Enumerates the set of values in String for WorkRequestResourceActionTypeEnum
func GetWorkRequestResourceActionTypeEnumStringValues() []string {
	return []string{
		"CREATED",
		"UPDATED",
		"DELETED",
		"MOVED",
		"IN_PROGRESS",
		"FAILED",
		"STOPPED",
		"STARTED",
	}
}

// GetMappingWorkRequestResourceActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestResourceActionTypeEnum(val string) (WorkRequestResourceActionTypeEnum, bool) {
	mappingWorkRequestResourceActionTypeEnumIgnoreCase := make(map[string]WorkRequestResourceActionTypeEnum)
	for k, v := range mappingWorkRequestResourceActionTypeEnum {
		mappingWorkRequestResourceActionTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingWorkRequestResourceActionTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
