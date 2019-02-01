// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Announcements Service API
//
// A description of the AnnouncementsService API
//

package announcementsservice

import (
	"github.com/oracle/oci-go-sdk/common"
)

// AffectedResource Descrption of a resource affected by the announcement
type AffectedResource struct {

	// The OCID of the resource
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// User-friendly name of the resource
	ResourceName *string `mandatory:"true" json:"resourceName"`

	// Region where this resource belongs to
	Region *string `mandatory:"true" json:"region"`
}

func (m AffectedResource) String() string {
	return common.PointerString(m)
}
