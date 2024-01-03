// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for the operating system environments in your private data centers through a single management console. For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SoftwarePackage The details for a software package.
type SoftwarePackage struct {

	// Package name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Unique identifier for the package. NOTE - This is not an OCID.
	Name *string `mandatory:"true" json:"name"`

	// Type of the package.
	Type *string `mandatory:"true" json:"type"`

	// Version of the package.
	Version *string `mandatory:"true" json:"version"`

	// The architecture for which this software was built
	Architecture *string `mandatory:"false" json:"architecture"`

	// Date of the last update to the package.
	LastModifiedDate *string `mandatory:"false" json:"lastModifiedDate"`

	// Checksum of the package.
	Checksum *string `mandatory:"false" json:"checksum"`

	// Type of the checksum.
	ChecksumType *string `mandatory:"false" json:"checksumType"`

	// Description of the package.
	Description *string `mandatory:"false" json:"description"`

	// Size of the package in bytes.
	SizeInBytes *int64 `mandatory:"false" json:"sizeInBytes"`

	// List of dependencies for the software package.
	Dependencies []SoftwarePackageDependency `mandatory:"false" json:"dependencies"`

	// List of files for the software package.
	Files []SoftwarePackageFile `mandatory:"false" json:"files"`

	// List of software sources that provide the software package.
	SoftwareSources []SoftwareSourceDetails `mandatory:"false" json:"softwareSources"`

	// Indicates whether this package is the latest version.
	IsLatest *bool `mandatory:"false" json:"isLatest"`
}

func (m SoftwarePackage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SoftwarePackage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
