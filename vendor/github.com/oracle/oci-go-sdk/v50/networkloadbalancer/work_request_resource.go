// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// NetworkLoadBalancer API
//
// A description of the network load balancer API
//

package networkloadbalancer

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// WorkRequestResource A resource that a work request creates or on which the work request operates.
type WorkRequestResource struct {

	// The resource type that the work request affects.
	EntityType *string `mandatory:"true" json:"entityType"`

	// The way in which this resource is affected by the work tracked in the work request.
	// A resource being created, updated, or deleted remains in the IN_PROGRESS state until
	// work is complete for that resource, at which point the resource transitions to CREATED, UPDATED,
	// or DELETED, respectively.
	ActionType ActionTypeEnum `mandatory:"true" json:"actionType"`

	// The identifier of the resource the work request affects.
	Identifier *string `mandatory:"true" json:"identifier"`

	// The URI path on which the user can perform a GET request to access the resource metadata.
	EntityUri *string `mandatory:"false" json:"entityUri"`
}

func (m WorkRequestResource) String() string {
	return common.PointerString(m)
}
