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

// UpdatablePackageSummary A software package available for install on a managed instance.
type UpdatablePackageSummary struct {

	// Package name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Unique identifier for the package.
	Name *string `mandatory:"true" json:"name"`

	// Type of the package.
	Type *string `mandatory:"true" json:"type"`

	// Version of the installed package.
	Version *string `mandatory:"true" json:"version"`

	// list of software sources that provide the software package.
	SoftwareSources []SoftwareSourceDetails `mandatory:"false" json:"softwareSources"`

	// The version of this upgradable package already installed on the instance.
	InstalledVersion *string `mandatory:"false" json:"installedVersion"`

	// List of errata containing this update.
	Errata []string `mandatory:"false" json:"errata"`

	// List of CVEs applicable to this erratum.
	RelatedCves []string `mandatory:"false" json:"relatedCves"`

	// The architecture for which this package was built.
	Architecture ArchTypeEnum `mandatory:"false" json:"architecture,omitempty"`

	// The classification of this update.
	UpdateType ClassificationTypesEnum `mandatory:"true" json:"updateType"`
}

// GetDisplayName returns DisplayName
func (m UpdatablePackageSummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetName returns Name
func (m UpdatablePackageSummary) GetName() *string {
	return m.Name
}

// GetType returns Type
func (m UpdatablePackageSummary) GetType() *string {
	return m.Type
}

// GetVersion returns Version
func (m UpdatablePackageSummary) GetVersion() *string {
	return m.Version
}

// GetArchitecture returns Architecture
func (m UpdatablePackageSummary) GetArchitecture() ArchTypeEnum {
	return m.Architecture
}

// GetSoftwareSources returns SoftwareSources
func (m UpdatablePackageSummary) GetSoftwareSources() []SoftwareSourceDetails {
	return m.SoftwareSources
}

func (m UpdatablePackageSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdatablePackageSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingArchTypeEnum(string(m.Architecture)); !ok && m.Architecture != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Architecture: %s. Supported values are: %s.", m.Architecture, strings.Join(GetArchTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingClassificationTypesEnum(string(m.UpdateType)); !ok && m.UpdateType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateType: %s. Supported values are: %s.", m.UpdateType, strings.Join(GetClassificationTypesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdatablePackageSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdatablePackageSummary UpdatablePackageSummary
	s := struct {
		DiscriminatorParam string `json:"packageClassification"`
		MarshalTypeUpdatablePackageSummary
	}{
		"UPDATABLE",
		(MarshalTypeUpdatablePackageSummary)(m),
	}

	return json.Marshal(&s)
}
