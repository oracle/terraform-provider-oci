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

// SoftwarePackage The details for a software package
type SoftwarePackage struct {

	// Package name
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Unique identifier for the package. NOTE - This is not an OCID
	Name *string `mandatory:"true" json:"name"`

	// Type of the package
	Type *string `mandatory:"true" json:"type"`

	// Version of the package
	Version *string `mandatory:"true" json:"version"`

	// the architecture for which this software was built
	Architecture *string `mandatory:"false" json:"architecture"`

	// date of the last update to the package
	LastModifiedDate *string `mandatory:"false" json:"lastModifiedDate"`

	// checksum of the package
	Checksum *string `mandatory:"false" json:"checksum"`

	// type of the checksum
	ChecksumType *string `mandatory:"false" json:"checksumType"`

	// description of the package
	Description *string `mandatory:"false" json:"description"`

	// size of the package in bytes
	SizeInBytes *int64 `mandatory:"false" json:"sizeInBytes"`

	// list of dependencies for the software package
	Dependencies []SoftwarePackageDependency `mandatory:"false" json:"dependencies"`

	// list of files for the software package
	Files []SoftwarePackageFile `mandatory:"false" json:"files"`

	// list of software sources that provide the software package
	SoftwareSources []SoftwareSourceId `mandatory:"false" json:"softwareSources"`
}

func (m SoftwarePackage) String() string {
	return common.PointerString(m)
}
