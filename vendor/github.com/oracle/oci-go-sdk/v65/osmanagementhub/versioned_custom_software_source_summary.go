// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// VersionedCustomSoftwareSourceSummary An immutable custom software source that is assigned a version and contains a custom collection of packages.
type VersionedCustomSoftwareSourceSummary struct {

	// The OCID for the software source.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the tenancy containing the software source.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// User friendly name for the software source.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The Repo ID for the software source.
	RepoId *string `mandatory:"true" json:"repoId"`

	// URL for the repository.
	Url *string `mandatory:"true" json:"url"`

	// The date and time the software source was created, as described in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time of when the software source was updated as described in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// List of vendor software sources.
	VendorSoftwareSources []Id `mandatory:"true" json:"vendorSoftwareSources"`

	// The version to assign to this custom software source.
	SoftwareSourceVersion *string `mandatory:"true" json:"softwareSourceVersion"`

	// Information specified by the user about the software source.
	Description *string `mandatory:"false" json:"description"`

	// Number of packages.
	PackageCount *int64 `mandatory:"false" json:"packageCount"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The current state of the software source.
	LifecycleState SoftwareSourceLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Possible availabilities of a software source.
	Availability AvailabilityEnum `mandatory:"true" json:"availability"`

	// The OS family the software source belongs to.
	OsFamily OsFamilyEnum `mandatory:"true" json:"osFamily"`

	// The architecture type supported by the software source.
	ArchType ArchTypeEnum `mandatory:"true" json:"archType"`
}

// GetId returns Id
func (m VersionedCustomSoftwareSourceSummary) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m VersionedCustomSoftwareSourceSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m VersionedCustomSoftwareSourceSummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetRepoId returns RepoId
func (m VersionedCustomSoftwareSourceSummary) GetRepoId() *string {
	return m.RepoId
}

// GetUrl returns Url
func (m VersionedCustomSoftwareSourceSummary) GetUrl() *string {
	return m.Url
}

// GetTimeCreated returns TimeCreated
func (m VersionedCustomSoftwareSourceSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m VersionedCustomSoftwareSourceSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetDescription returns Description
func (m VersionedCustomSoftwareSourceSummary) GetDescription() *string {
	return m.Description
}

// GetAvailability returns Availability
func (m VersionedCustomSoftwareSourceSummary) GetAvailability() AvailabilityEnum {
	return m.Availability
}

// GetOsFamily returns OsFamily
func (m VersionedCustomSoftwareSourceSummary) GetOsFamily() OsFamilyEnum {
	return m.OsFamily
}

// GetArchType returns ArchType
func (m VersionedCustomSoftwareSourceSummary) GetArchType() ArchTypeEnum {
	return m.ArchType
}

// GetPackageCount returns PackageCount
func (m VersionedCustomSoftwareSourceSummary) GetPackageCount() *int64 {
	return m.PackageCount
}

// GetLifecycleState returns LifecycleState
func (m VersionedCustomSoftwareSourceSummary) GetLifecycleState() SoftwareSourceLifecycleStateEnum {
	return m.LifecycleState
}

// GetFreeformTags returns FreeformTags
func (m VersionedCustomSoftwareSourceSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m VersionedCustomSoftwareSourceSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m VersionedCustomSoftwareSourceSummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m VersionedCustomSoftwareSourceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VersionedCustomSoftwareSourceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSoftwareSourceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSoftwareSourceLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAvailabilityEnum(string(m.Availability)); !ok && m.Availability != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Availability: %s. Supported values are: %s.", m.Availability, strings.Join(GetAvailabilityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOsFamilyEnum(string(m.OsFamily)); !ok && m.OsFamily != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsFamily: %s. Supported values are: %s.", m.OsFamily, strings.Join(GetOsFamilyEnumStringValues(), ",")))
	}
	if _, ok := GetMappingArchTypeEnum(string(m.ArchType)); !ok && m.ArchType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ArchType: %s. Supported values are: %s.", m.ArchType, strings.Join(GetArchTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m VersionedCustomSoftwareSourceSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeVersionedCustomSoftwareSourceSummary VersionedCustomSoftwareSourceSummary
	s := struct {
		DiscriminatorParam string `json:"softwareSourceType"`
		MarshalTypeVersionedCustomSoftwareSourceSummary
	}{
		"VERSIONED",
		(MarshalTypeVersionedCustomSoftwareSourceSummary)(m),
	}

	return json.Marshal(&s)
}
