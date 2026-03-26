// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// SoftwareSourceRepoSummary Provides summary information for available repos to add directly to compartments. A software source contains a collection of packages. For more information, see Managing Software Sources (https://docs.oracle.com/iaas/osmh/doc/software-sources.htm).
type SoftwareSourceRepoSummary struct {

	// User-friendly name for the software source.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The repository ID for the software source.
	RepoId *string `mandatory:"true" json:"repoId"`

	// Type of software source.
	SoftwareSourceType SoftwareSourceTypeEnum `mandatory:"true" json:"softwareSourceType"`

	// The OS family of the software source.
	OsFamily OsFamilyEnum `mandatory:"true" json:"osFamily"`

	// The architecture type supported by the software source.
	ArchType ArchTypeEnum `mandatory:"true" json:"archType"`

	// Description of the software source. For custom software sources, this is user-specified.
	Description *string `mandatory:"false" json:"description"`
}

func (m SoftwareSourceRepoSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SoftwareSourceRepoSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSoftwareSourceTypeEnum(string(m.SoftwareSourceType)); !ok && m.SoftwareSourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SoftwareSourceType: %s. Supported values are: %s.", m.SoftwareSourceType, strings.Join(GetSoftwareSourceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOsFamilyEnum(string(m.OsFamily)); !ok && m.OsFamily != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsFamily: %s. Supported values are: %s.", m.OsFamily, strings.Join(GetOsFamilyEnumStringValues(), ",")))
	}
	if _, ok := GetMappingArchTypeEnum(string(m.ArchType)); !ok && m.ArchType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ArchType: %s. Supported values are: %s.", m.ArchType, strings.Join(GetArchTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
