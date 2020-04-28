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

// AvailableUpdateSummary An update available for a managed instance
type AvailableUpdateSummary struct {

	// Package name
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Unique identifier for the package available for update. NOTE - This is not an OCID
	Name *string `mandatory:"true" json:"name"`

	// Type of the package
	Type *string `mandatory:"true" json:"type"`

	// Version of the installed package
	InstalledVersion *string `mandatory:"true" json:"installedVersion"`

	// Version of the package available for update
	AvailableVersion *string `mandatory:"true" json:"availableVersion"`

	// The purpose of this update.
	UpdateType UpdateTypesEnum `mandatory:"false" json:"updateType,omitempty"`

	// The architecture for which this package was built
	Architecture *string `mandatory:"false" json:"architecture"`

	// List of errata containing this update
	Errata []Id `mandatory:"false" json:"errata"`

	// List of CVEs applicable to this erratum
	RelatedCves []string `mandatory:"false" json:"relatedCves"`

	// list of software sources that provide the software package
	SoftwareSources []SoftwareSourceId `mandatory:"false" json:"softwareSources"`
}

func (m AvailableUpdateSummary) String() string {
	return common.PointerString(m)
}
