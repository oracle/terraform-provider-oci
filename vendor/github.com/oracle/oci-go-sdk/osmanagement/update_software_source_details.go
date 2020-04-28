// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management API
//
// API for the OS Management service. Use these API operations for working
// with Managed instances and Managed instance groups.
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
