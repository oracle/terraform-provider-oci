// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SoftwarePackageSummary Provides summary information for a software package.
type SoftwarePackageSummary struct {

	// Package name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Unique identifier for the package. Note that this is not an OCID.
	Name *string `mandatory:"true" json:"name"`

	// Type of the package.
	Type *string `mandatory:"true" json:"type"`

	// Version of the package.
	Version *string `mandatory:"true" json:"version"`

	// The architecture for which this software was built.
	Architecture SoftwarePackageArchitectureEnum `mandatory:"false" json:"architecture,omitempty"`

	// Checksum of the package.
	Checksum *string `mandatory:"false" json:"checksum"`

	// Type of the checksum.
	ChecksumType *string `mandatory:"false" json:"checksumType"`

	// Indicates whether this package is the latest version.
	IsLatest *bool `mandatory:"false" json:"isLatest"`

	// List of software sources that provide the software package. This property is deprecated and it will be removed in a future API release.
	SoftwareSources []SoftwareSourceDetails `mandatory:"false" json:"softwareSources"`

	// The OS families the package belongs to.
	OsFamilies []OsFamilyEnum `mandatory:"false" json:"osFamilies"`
}

func (m SoftwarePackageSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SoftwarePackageSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSoftwarePackageArchitectureEnum(string(m.Architecture)); !ok && m.Architecture != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Architecture: %s. Supported values are: %s.", m.Architecture, strings.Join(GetSoftwarePackageArchitectureEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
