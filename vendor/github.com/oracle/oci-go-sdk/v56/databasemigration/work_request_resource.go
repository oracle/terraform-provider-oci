// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// WorkRequestResource A resource that is created or operated on by an asynchronous operation that is tracked by
// a work request.
type WorkRequestResource struct {

	// The way in which this resource was affected by the operation that spawned the work
	// request.
	ActionType WorkRequestResourceActionTypeEnum `mandatory:"true" json:"actionType"`

	// The resource type the work request affects.
	EntityType *string `mandatory:"true" json:"entityType"`

	// An OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) or other unique identifier for the
	// resource.
	Identifier *string `mandatory:"true" json:"identifier"`

	// The URI path that you can use for a GET request to access the resource metadata.
	EntityUri *string `mandatory:"false" json:"entityUri"`
}

func (m WorkRequestResource) String() string {
	return common.PointerString(m)
}

// WorkRequestResourceActionTypeEnum Enum with underlying type: string
type WorkRequestResourceActionTypeEnum string

// Set of constants representing the allowable values for WorkRequestResourceActionTypeEnum
const (
	WorkRequestResourceActionTypeCreated    WorkRequestResourceActionTypeEnum = "CREATED"
	WorkRequestResourceActionTypeUpdated    WorkRequestResourceActionTypeEnum = "UPDATED"
	WorkRequestResourceActionTypeDeleted    WorkRequestResourceActionTypeEnum = "DELETED"
	WorkRequestResourceActionTypeRelated    WorkRequestResourceActionTypeEnum = "RELATED"
	WorkRequestResourceActionTypeInProgress WorkRequestResourceActionTypeEnum = "IN_PROGRESS"
)

var mappingWorkRequestResourceActionType = map[string]WorkRequestResourceActionTypeEnum{
	"CREATED":     WorkRequestResourceActionTypeCreated,
	"UPDATED":     WorkRequestResourceActionTypeUpdated,
	"DELETED":     WorkRequestResourceActionTypeDeleted,
	"RELATED":     WorkRequestResourceActionTypeRelated,
	"IN_PROGRESS": WorkRequestResourceActionTypeInProgress,
}

// GetWorkRequestResourceActionTypeEnumValues Enumerates the set of values for WorkRequestResourceActionTypeEnum
func GetWorkRequestResourceActionTypeEnumValues() []WorkRequestResourceActionTypeEnum {
	values := make([]WorkRequestResourceActionTypeEnum, 0)
	for _, v := range mappingWorkRequestResourceActionType {
		values = append(values, v)
	}
	return values
}
