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
