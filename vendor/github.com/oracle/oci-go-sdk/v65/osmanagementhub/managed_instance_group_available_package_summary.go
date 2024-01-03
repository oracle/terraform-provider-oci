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

// ManagedInstanceGroupAvailablePackageSummary Summary information pertaining to an available package for a managed instance group.
type ManagedInstanceGroupAvailablePackageSummary struct {

	// Package name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Unique identifier for the package. NOTE - This is not an OCID.
	Name *string `mandatory:"true" json:"name"`

	// Type of the package.
	Type *string `mandatory:"true" json:"type"`

	// Version of the installed package.
	Version *string `mandatory:"true" json:"version"`

	// The architecture for which this package was built.
	Architecture ArchTypeEnum `mandatory:"false" json:"architecture,omitempty"`

	// List of software sources that provide the software package.
	SoftwareSources []SoftwareSourceDetails `mandatory:"false" json:"softwareSources"`

	// Flag to return only latest package versions.
	IsLatest *bool `mandatory:"false" json:"isLatest"`
}

func (m ManagedInstanceGroupAvailablePackageSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagedInstanceGroupAvailablePackageSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingArchTypeEnum(string(m.Architecture)); !ok && m.Architecture != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Architecture: %s. Supported values are: %s.", m.Architecture, strings.Join(GetArchTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
