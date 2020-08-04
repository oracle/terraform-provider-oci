// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"github.com/oracle/oci-go-sdk/common"
)

// WorkRequestResource A resource related to a Data Flow work request.
type WorkRequestResource struct {

	// The way in which this resource is affected by the work tracked in the work request.
	ActionType WorkRequestResourceActionTypeEnum `mandatory:"true" json:"actionType"`

	// The id of the releated resource. See resourceType to identity the specific type of resource.
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// The type of resource.  See resourceId for the id of the specific resource.
	ResourceType *string `mandatory:"true" json:"resourceType"`

	// The id of a work request resource object.
	Id *int64 `mandatory:"false" json:"id"`

	// The URI path that the user can use to get access to the resource metadata
	ResourceUri *string `mandatory:"false" json:"resourceUri"`

	// The OCID of a work request.
	WorkRequestid *string `mandatory:"false" json:"workRequestid"`
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
	WorkRequestResourceActionTypeInprogress WorkRequestResourceActionTypeEnum = "INPROGRESS"
	WorkRequestResourceActionTypeRelated    WorkRequestResourceActionTypeEnum = "RELATED"
)

var mappingWorkRequestResourceActionType = map[string]WorkRequestResourceActionTypeEnum{
	"CREATED":    WorkRequestResourceActionTypeCreated,
	"UPDATED":    WorkRequestResourceActionTypeUpdated,
	"DELETED":    WorkRequestResourceActionTypeDeleted,
	"INPROGRESS": WorkRequestResourceActionTypeInprogress,
	"RELATED":    WorkRequestResourceActionTypeRelated,
}

// GetWorkRequestResourceActionTypeEnumValues Enumerates the set of values for WorkRequestResourceActionTypeEnum
func GetWorkRequestResourceActionTypeEnumValues() []WorkRequestResourceActionTypeEnum {
	values := make([]WorkRequestResourceActionTypeEnum, 0)
	for _, v := range mappingWorkRequestResourceActionType {
		values = append(values, v)
	}
	return values
}
