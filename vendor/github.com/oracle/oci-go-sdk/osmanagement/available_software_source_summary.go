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

// AvailableSoftwareSourceSummary A software source which can be added to a managed instance. Once a software source is added, packages from that software source can be installed on that managed instance.
type AvailableSoftwareSourceSummary struct {

	// unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// OCID for the Compartment
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// User friendly name for the software source
	DisplayName *string `mandatory:"true" json:"displayName"`

	// OCID for the parent software source, if there is one
	ParentId *string `mandatory:"false" json:"parentId"`

	// Display name of the parent software source, if there is one
	ParentName *string `mandatory:"false" json:"parentName"`
}

func (m AvailableSoftwareSourceSummary) String() string {
	return common.PointerString(m)
}
