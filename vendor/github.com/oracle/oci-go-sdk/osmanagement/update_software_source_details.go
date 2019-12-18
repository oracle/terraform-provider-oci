// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// OSMS
//
// OS Management as a Service API definition
//

package osmanagement

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateSoftwareSourceDetails Information for updating a software source on the management system
type UpdateSoftwareSourceDetails struct {

	// User friendly name for the software source
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Information specified by the user about the software source
	Description *string `mandatory:"false" json:"description"`

	// Name of the person maintaining this software source
	MaintainerName *string `mandatory:"false" json:"maintainerName"`

	// Email address of the person maintaining this software source
	MaintainerEmail *string `mandatory:"false" json:"maintainerEmail"`

	// Phone number of the person maintaining this software source
	MaintainerPhone *string `mandatory:"false" json:"maintainerPhone"`

	// The yum repository checksum type used by this software source
	ChecksumType ChecksumTypesEnum `mandatory:"false" json:"checksumType,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateSoftwareSourceDetails) String() string {
	return common.PointerString(m)
}
