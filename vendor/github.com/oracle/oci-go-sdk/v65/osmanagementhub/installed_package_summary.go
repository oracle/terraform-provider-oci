// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for the operating system environments in your private data centers through a single management console. For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InstalledPackageSummary A software package installed on a managed instance.
type InstalledPackageSummary struct {

	// Package name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Unique identifier for the package.
	Name *string `mandatory:"true" json:"name"`

	// Type of the package.
	Type *string `mandatory:"true" json:"type"`

	// Version of the installed package.
	Version *string `mandatory:"true" json:"version"`

	// The date and time the package was installed, as described in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeInstalled *common.SDKTime `mandatory:"true" json:"timeInstalled"`

	// list of software sources that provide the software package.
	SoftwareSources []SoftwareSourceDetails `mandatory:"false" json:"softwareSources"`

	// The date and time the package was issued by a providing erratum (if available), as described in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeIssued *common.SDKTime `mandatory:"false" json:"timeIssued"`

	// The architecture for which this package was built.
	Architecture ArchTypeEnum `mandatory:"false" json:"architecture,omitempty"`
}

// GetDisplayName returns DisplayName
func (m InstalledPackageSummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetName returns Name
func (m InstalledPackageSummary) GetName() *string {
	return m.Name
}

// GetType returns Type
func (m InstalledPackageSummary) GetType() *string {
	return m.Type
}

// GetVersion returns Version
func (m InstalledPackageSummary) GetVersion() *string {
	return m.Version
}

// GetArchitecture returns Architecture
func (m InstalledPackageSummary) GetArchitecture() ArchTypeEnum {
	return m.Architecture
}

// GetSoftwareSources returns SoftwareSources
func (m InstalledPackageSummary) GetSoftwareSources() []SoftwareSourceDetails {
	return m.SoftwareSources
}

func (m InstalledPackageSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstalledPackageSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingArchTypeEnum(string(m.Architecture)); !ok && m.Architecture != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Architecture: %s. Supported values are: %s.", m.Architecture, strings.Join(GetArchTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m InstalledPackageSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInstalledPackageSummary InstalledPackageSummary
	s := struct {
		DiscriminatorParam string `json:"packageClassification"`
		MarshalTypeInstalledPackageSummary
	}{
		"INSTALLED",
		(MarshalTypeInstalledPackageSummary)(m),
	}

	return json.Marshal(&s)
}
