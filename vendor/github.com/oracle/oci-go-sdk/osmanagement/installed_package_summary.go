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

// InstalledPackageSummary A software package installed on a managed instance
type InstalledPackageSummary struct {

	// Package name
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Unique identifier for the package. NOTE - This is not an OCID
	Name *string `mandatory:"true" json:"name"`

	// Type of the package
	Type *string `mandatory:"true" json:"type"`

	// Version of the installed package
	Version *string `mandatory:"true" json:"version"`

	// The architecture for which this package was built
	Architecture *string `mandatory:"false" json:"architecture"`

	// Install time of the package
	InstallTime *string `mandatory:"false" json:"installTime"`

	// list of software sources that provide the software package
	SoftwareSources []SoftwareSourceId `mandatory:"false" json:"softwareSources"`
}

func (m InstalledPackageSummary) String() string {
	return common.PointerString(m)
}
