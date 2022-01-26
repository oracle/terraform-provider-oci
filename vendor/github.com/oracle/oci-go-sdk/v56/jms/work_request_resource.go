// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// WorkRequestResource A resource that is created or operated on by an asynchronous operation that is tracked by a work request.
type WorkRequestResource struct {

	// The resource type affected by the work request.
	EntityType *string `mandatory:"true" json:"entityType"`

	// The way in which this resource was affected by the operation that spawned the work request.
	// A resource being created, updated, or deleted will remain in the IN_PROGRESS state until
	// work is complete for that resource at which point it will transition to CREATED, UPDATED,
	// or DELETED, respectively.
	ActionType ActionTypeEnum `mandatory:"true" json:"actionType"`

	// An OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) or other unique identifier of the resource affected by the work request.
	Identifier *string `mandatory:"true" json:"identifier"`

	// The URI path that the user can perform a GET operation to access the resource metadata.
	EntityUri *string `mandatory:"false" json:"entityUri"`
}

func (m WorkRequestResource) String() string {
	return common.PointerString(m)
}
