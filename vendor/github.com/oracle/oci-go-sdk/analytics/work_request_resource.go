// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

import (
	"github.com/oracle/oci-go-sdk/common"
)

// WorkRequestResource The representation of WorkRequestResource
type WorkRequestResource struct {

	// The way in which this resource was affected by this work request.
	ActionResult WorkRequestActionResultEnum `mandatory:"true" json:"actionResult"`

	// The type of the resource the work request is affecting.
	ResourceType WorkRequestResourceTypeEnum `mandatory:"true" json:"resourceType"`

	// The OCID of the resource the work request is affecting.
	Identifier *string `mandatory:"true" json:"identifier"`

	// The URI of the affected resource.
	ResourceUri *string `mandatory:"true" json:"resourceUri"`

	// Additional metadata of the resource.
	Metadata map[string]string `mandatory:"false" json:"metadata"`
}

func (m WorkRequestResource) String() string {
	return common.PointerString(m)
}
