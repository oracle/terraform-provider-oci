// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Advisor API
//
// Use the Cloud Advisor API to find potential inefficiencies in your tenancy and address them.
// Cloud Advisor can help you save money, improve performance, strengthen system resilience, and improve security.
// For more information, see Cloud Advisor (https://docs.cloud.oracle.com/Content/CloudAdvisor/Concepts/cloudadvisoroverview.htm).
//

package optimizer

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// WorkRequestResource Details about the resource entity.
type WorkRequestResource struct {

	// The resource type the work request affects.
	EntityType *string `mandatory:"true" json:"entityType"`

	// The way in which this resource was affected by the work tracked by the work request.
	// A resource being created, updated, or deleted remains in the `IN_PROGRESS` state until
	// work is complete for that resource. At that point, the resource transitions to the `CREATED`, `UPDATED`,
	// or `DELETED` state.
	ActionType WorkRequestActionTypeEnum `mandatory:"true" json:"actionType"`

	// The resource identifier the work request affects.
	Identifier *string `mandatory:"true" json:"identifier"`

	// The URI path that the user can do a GET on to access the resource metadata
	EntityUri *string `mandatory:"false" json:"entityUri"`

	// Additional information about the resource.
	Metadata map[string]string `mandatory:"false" json:"metadata"`
}

func (m WorkRequestResource) String() string {
	return common.PointerString(m)
}
