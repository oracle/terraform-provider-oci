// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateTagDefaultDetails The representation of CreateTagDefaultDetails
type CreateTagDefaultDetails struct {

	// The OCID of the Compartment. The Tag Default will apply to any resource contained in this Compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the Tag Definition. The Tag Default will always assign a default value for this Tag Definition.
	TagDefinitionId *string `mandatory:"true" json:"tagDefinitionId"`

	// The default value for the Tag Definition. This will be applied to all resources created in the Compartment.
	Value *string `mandatory:"true" json:"value"`
}

func (m CreateTagDefaultDetails) String() string {
	return common.PointerString(m)
}
