// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// IamWorkRequestResource A IAM work request resource entry
type IamWorkRequestResource struct {

	// The way in which this resource is affected by the work tracked in the work request.
	// A resource being created, updated, or deleted will remain in the IN_PROGRESS state until
	// work is complete for that resource at which point it will transition to CREATED, UPDATED,
	// or DELETED, respectively.
	ActionType IamWorkRequestResourceActionTypeEnum `mandatory:"true" json:"actionType"`

	// The resource type the work request is affects.
	EntityType *string `mandatory:"true" json:"entityType"`

	// An OCID of the resource that the work request affects.
	Identifier *string `mandatory:"true" json:"identifier"`

	// The URI path that the user can do a GET on to access the resource metadata.
	EntityUri *string `mandatory:"false" json:"entityUri"`
}

func (m IamWorkRequestResource) String() string {
	return common.PointerString(m)
}

// IamWorkRequestResourceActionTypeEnum Enum with underlying type: string
type IamWorkRequestResourceActionTypeEnum string

// Set of constants representing the allowable values for IamWorkRequestResourceActionTypeEnum
const (
	IamWorkRequestResourceActionTypeCreated    IamWorkRequestResourceActionTypeEnum = "CREATED"
	IamWorkRequestResourceActionTypeUpdated    IamWorkRequestResourceActionTypeEnum = "UPDATED"
	IamWorkRequestResourceActionTypeDeleted    IamWorkRequestResourceActionTypeEnum = "DELETED"
	IamWorkRequestResourceActionTypeRelated    IamWorkRequestResourceActionTypeEnum = "RELATED"
	IamWorkRequestResourceActionTypeInProgress IamWorkRequestResourceActionTypeEnum = "IN_PROGRESS"
)

var mappingIamWorkRequestResourceActionType = map[string]IamWorkRequestResourceActionTypeEnum{
	"CREATED":     IamWorkRequestResourceActionTypeCreated,
	"UPDATED":     IamWorkRequestResourceActionTypeUpdated,
	"DELETED":     IamWorkRequestResourceActionTypeDeleted,
	"RELATED":     IamWorkRequestResourceActionTypeRelated,
	"IN_PROGRESS": IamWorkRequestResourceActionTypeInProgress,
}

// GetIamWorkRequestResourceActionTypeEnumValues Enumerates the set of values for IamWorkRequestResourceActionTypeEnum
func GetIamWorkRequestResourceActionTypeEnumValues() []IamWorkRequestResourceActionTypeEnum {
	values := make([]IamWorkRequestResourceActionTypeEnum, 0)
	for _, v := range mappingIamWorkRequestResourceActionType {
		values = append(values, v)
	}
	return values
}
