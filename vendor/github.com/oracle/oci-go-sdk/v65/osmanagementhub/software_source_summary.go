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

// SoftwareSourceSummary Provides summary information for a software source. A software source contains a collection of packages. For more information, see Managing Software Sources (https://docs.cloud.oracle.com/iaas/osmh/doc/software-sources.htm).
type SoftwareSourceSummary interface {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source.
	GetId() *string

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the software source.
	GetCompartmentId() *string

	// User-friendly name for the software source.
	GetDisplayName() *string

	// The repository ID for the software source.
	GetRepoId() *string

	// URL for the repository. For vendor software sources, this is the URL to the regional yum server. For custom software sources, this is 'custom/<repoId>'.
	GetUrl() *string

	// The date and time the software source was created (in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) format).
	GetTimeCreated() *common.SDKTime

	// The date and time the software source was updated (in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) format).
	GetTimeUpdated() *common.SDKTime

	// Availability of the software source (for non-OCI environments).
	GetAvailability() AvailabilityEnum

	// Availability of the software source (for OCI environments).
	GetAvailabilityAtOci() AvailabilityEnum

	// The OS family the software source belongs to.
	GetOsFamily() OsFamilyEnum

	// The architecture type supported by the software source.
	GetArchType() ArchTypeEnum

	// Description of the software source. For custom software sources, this is user-specified.
	GetDescription() *string

	// Number of packages the software source contains.
	GetPackageCount() *int64

	// The current state of the software source.
	GetLifecycleState() SoftwareSourceLifecycleStateEnum

	// The size of the software source in gigabytes (GB).
	GetSize() *float64

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}
}

type softwaresourcesummary struct {
	JsonData           []byte
	Description        *string                           `mandatory:"false" json:"description"`
	PackageCount       *int64                            `mandatory:"false" json:"packageCount"`
	LifecycleState     SoftwareSourceLifecycleStateEnum  `mandatory:"false" json:"lifecycleState,omitempty"`
	Size               *float64                          `mandatory:"false" json:"size"`
	FreeformTags       map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags        map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags         map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	Id                 *string                           `mandatory:"true" json:"id"`
	CompartmentId      *string                           `mandatory:"true" json:"compartmentId"`
	DisplayName        *string                           `mandatory:"true" json:"displayName"`
	RepoId             *string                           `mandatory:"true" json:"repoId"`
	Url                *string                           `mandatory:"true" json:"url"`
	TimeCreated        *common.SDKTime                   `mandatory:"true" json:"timeCreated"`
	TimeUpdated        *common.SDKTime                   `mandatory:"true" json:"timeUpdated"`
	Availability       AvailabilityEnum                  `mandatory:"true" json:"availability"`
	AvailabilityAtOci  AvailabilityEnum                  `mandatory:"true" json:"availabilityAtOci"`
	OsFamily           OsFamilyEnum                      `mandatory:"true" json:"osFamily"`
	ArchType           ArchTypeEnum                      `mandatory:"true" json:"archType"`
	SoftwareSourceType string                            `json:"softwareSourceType"`
}

// UnmarshalJSON unmarshals json
func (m *softwaresourcesummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersoftwaresourcesummary softwaresourcesummary
	s := struct {
		Model Unmarshalersoftwaresourcesummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.CompartmentId = s.Model.CompartmentId
	m.DisplayName = s.Model.DisplayName
	m.RepoId = s.Model.RepoId
	m.Url = s.Model.Url
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.Availability = s.Model.Availability
	m.AvailabilityAtOci = s.Model.AvailabilityAtOci
	m.OsFamily = s.Model.OsFamily
	m.ArchType = s.Model.ArchType
	m.Description = s.Model.Description
	m.PackageCount = s.Model.PackageCount
	m.LifecycleState = s.Model.LifecycleState
	m.Size = s.Model.Size
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.SoftwareSourceType = s.Model.SoftwareSourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *softwaresourcesummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SoftwareSourceType {
	case "VENDOR":
		mm := VendorSoftwareSourceSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VERSIONED":
		mm := VersionedCustomSoftwareSourceSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CUSTOM":
		mm := CustomSoftwareSourceSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for SoftwareSourceSummary: %s.", m.SoftwareSourceType)
		return *m, nil
	}
}

// GetDescription returns Description
func (m softwaresourcesummary) GetDescription() *string {
	return m.Description
}

// GetPackageCount returns PackageCount
func (m softwaresourcesummary) GetPackageCount() *int64 {
	return m.PackageCount
}

// GetLifecycleState returns LifecycleState
func (m softwaresourcesummary) GetLifecycleState() SoftwareSourceLifecycleStateEnum {
	return m.LifecycleState
}

// GetSize returns Size
func (m softwaresourcesummary) GetSize() *float64 {
	return m.Size
}

// GetFreeformTags returns FreeformTags
func (m softwaresourcesummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m softwaresourcesummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m softwaresourcesummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetId returns Id
func (m softwaresourcesummary) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m softwaresourcesummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m softwaresourcesummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetRepoId returns RepoId
func (m softwaresourcesummary) GetRepoId() *string {
	return m.RepoId
}

// GetUrl returns Url
func (m softwaresourcesummary) GetUrl() *string {
	return m.Url
}

// GetTimeCreated returns TimeCreated
func (m softwaresourcesummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m softwaresourcesummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetAvailability returns Availability
func (m softwaresourcesummary) GetAvailability() AvailabilityEnum {
	return m.Availability
}

// GetAvailabilityAtOci returns AvailabilityAtOci
func (m softwaresourcesummary) GetAvailabilityAtOci() AvailabilityEnum {
	return m.AvailabilityAtOci
}

// GetOsFamily returns OsFamily
func (m softwaresourcesummary) GetOsFamily() OsFamilyEnum {
	return m.OsFamily
}

// GetArchType returns ArchType
func (m softwaresourcesummary) GetArchType() ArchTypeEnum {
	return m.ArchType
}

func (m softwaresourcesummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m softwaresourcesummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAvailabilityEnum(string(m.Availability)); !ok && m.Availability != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Availability: %s. Supported values are: %s.", m.Availability, strings.Join(GetAvailabilityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAvailabilityEnum(string(m.AvailabilityAtOci)); !ok && m.AvailabilityAtOci != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AvailabilityAtOci: %s. Supported values are: %s.", m.AvailabilityAtOci, strings.Join(GetAvailabilityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOsFamilyEnum(string(m.OsFamily)); !ok && m.OsFamily != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsFamily: %s. Supported values are: %s.", m.OsFamily, strings.Join(GetOsFamilyEnumStringValues(), ",")))
	}
	if _, ok := GetMappingArchTypeEnum(string(m.ArchType)); !ok && m.ArchType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ArchType: %s. Supported values are: %s.", m.ArchType, strings.Join(GetArchTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingSoftwareSourceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSoftwareSourceLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
