// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PackageSummary Provides summary information for a software package.
type PackageSummary interface {

	// Package name.
	GetDisplayName() *string

	// Unique identifier for the package.
	GetName() *string

	// Type of the package.
	GetType() *string

	// Version of the installed package.
	GetVersion() *string

	// The architecture for which this package was built.
	GetArchitecture() ArchTypeEnum

	// List of software sources that provide the software package.
	GetSoftwareSources() []SoftwareSourceDetails
}

type packagesummary struct {
	JsonData              []byte
	Architecture          ArchTypeEnum            `mandatory:"false" json:"architecture,omitempty"`
	SoftwareSources       []SoftwareSourceDetails `mandatory:"false" json:"softwareSources"`
	DisplayName           *string                 `mandatory:"true" json:"displayName"`
	Name                  *string                 `mandatory:"true" json:"name"`
	Type                  *string                 `mandatory:"true" json:"type"`
	Version               *string                 `mandatory:"true" json:"version"`
	PackageClassification string                  `json:"packageClassification"`
}

// UnmarshalJSON unmarshals json
func (m *packagesummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerpackagesummary packagesummary
	s := struct {
		Model Unmarshalerpackagesummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.Name = s.Model.Name
	m.Type = s.Model.Type
	m.Version = s.Model.Version
	m.Architecture = s.Model.Architecture
	m.SoftwareSources = s.Model.SoftwareSources
	m.PackageClassification = s.Model.PackageClassification

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *packagesummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.PackageClassification {
	case "AVAILABLE":
		mm := AvailablePackageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INSTALLED":
		mm := InstalledPackageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "UPDATABLE":
		mm := UpdatablePackageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for PackageSummary: %s.", m.PackageClassification)
		return *m, nil
	}
}

// GetArchitecture returns Architecture
func (m packagesummary) GetArchitecture() ArchTypeEnum {
	return m.Architecture
}

// GetSoftwareSources returns SoftwareSources
func (m packagesummary) GetSoftwareSources() []SoftwareSourceDetails {
	return m.SoftwareSources
}

// GetDisplayName returns DisplayName
func (m packagesummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetName returns Name
func (m packagesummary) GetName() *string {
	return m.Name
}

// GetType returns Type
func (m packagesummary) GetType() *string {
	return m.Type
}

// GetVersion returns Version
func (m packagesummary) GetVersion() *string {
	return m.Version
}

func (m packagesummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m packagesummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingArchTypeEnum(string(m.Architecture)); !ok && m.Architecture != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Architecture: %s. Supported values are: %s.", m.Architecture, strings.Join(GetArchTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PackageSummaryPackageClassificationEnum Enum with underlying type: string
type PackageSummaryPackageClassificationEnum string

// Set of constants representing the allowable values for PackageSummaryPackageClassificationEnum
const (
	PackageSummaryPackageClassificationInstalled PackageSummaryPackageClassificationEnum = "INSTALLED"
	PackageSummaryPackageClassificationAvailable PackageSummaryPackageClassificationEnum = "AVAILABLE"
	PackageSummaryPackageClassificationUpdatable PackageSummaryPackageClassificationEnum = "UPDATABLE"
)

var mappingPackageSummaryPackageClassificationEnum = map[string]PackageSummaryPackageClassificationEnum{
	"INSTALLED": PackageSummaryPackageClassificationInstalled,
	"AVAILABLE": PackageSummaryPackageClassificationAvailable,
	"UPDATABLE": PackageSummaryPackageClassificationUpdatable,
}

var mappingPackageSummaryPackageClassificationEnumLowerCase = map[string]PackageSummaryPackageClassificationEnum{
	"installed": PackageSummaryPackageClassificationInstalled,
	"available": PackageSummaryPackageClassificationAvailable,
	"updatable": PackageSummaryPackageClassificationUpdatable,
}

// GetPackageSummaryPackageClassificationEnumValues Enumerates the set of values for PackageSummaryPackageClassificationEnum
func GetPackageSummaryPackageClassificationEnumValues() []PackageSummaryPackageClassificationEnum {
	values := make([]PackageSummaryPackageClassificationEnum, 0)
	for _, v := range mappingPackageSummaryPackageClassificationEnum {
		values = append(values, v)
	}
	return values
}

// GetPackageSummaryPackageClassificationEnumStringValues Enumerates the set of values in String for PackageSummaryPackageClassificationEnum
func GetPackageSummaryPackageClassificationEnumStringValues() []string {
	return []string{
		"INSTALLED",
		"AVAILABLE",
		"UPDATABLE",
	}
}

// GetMappingPackageSummaryPackageClassificationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPackageSummaryPackageClassificationEnum(val string) (PackageSummaryPackageClassificationEnum, bool) {
	enum, ok := mappingPackageSummaryPackageClassificationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
