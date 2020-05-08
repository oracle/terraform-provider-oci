// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud VMware Solution API
//
// Use this API to manage the Oracle Cloud VMware Solution.
//

package ocvp

import (
	"github.com/oracle/oci-go-sdk/common"
)

// WorkRequestResource A resource that is created or operated on by an asynchronous operation that is
// tracked by a work request.
type WorkRequestResource struct {

	// The resource type the work request affects.
	EntityType *string `mandatory:"true" json:"entityType"`

	// The way in which this resource was affected by the operation that spawned the
	// work request.
	ActionType ActionTypesEnum `mandatory:"true" json:"actionType"`

	// An OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) or other unique identifier
	// for the resource.
	Identifier *string `mandatory:"true" json:"identifier"`

	// The URI path that you can use for a GET request to access the resource metadata.
	EntityUri *string `mandatory:"false" json:"entityUri"`
}

func (m WorkRequestResource) String() string {
	return common.PointerString(m)
}
