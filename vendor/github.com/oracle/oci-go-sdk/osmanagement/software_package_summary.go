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

// SoftwarePackageSummary Summary information for a software package
type SoftwarePackageSummary struct {

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

	// checksum of the package
	Checksum *string `mandatory:"false" json:"checksum"`

	// type of the checksum
	ChecksumType *string `mandatory:"false" json:"checksumType"`
}

func (m SoftwarePackageSummary) String() string {
	return common.PointerString(m)
}
