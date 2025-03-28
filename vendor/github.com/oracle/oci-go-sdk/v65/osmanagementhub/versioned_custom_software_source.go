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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VersionedCustomSoftwareSource The object that defines a versioned custom software source.
type VersionedCustomSoftwareSource struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the software source.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// User-friendly name for the software source.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The date and time the software source was created (in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) format).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The repository ID for the software source.
	RepoId *string `mandatory:"true" json:"repoId"`

	// URL for the repository. For vendor software sources, this is the URL to the regional yum server. For custom software sources, this is 'custom/<repoId>'.
	Url *string `mandatory:"true" json:"url"`

	// List of vendor software sources that are used for the basis of the versioned custom software source.
	VendorSoftwareSources []Id `mandatory:"true" json:"vendorSoftwareSources"`

	// The version to assign to this custom software source.
	SoftwareSourceVersion *string `mandatory:"true" json:"softwareSourceVersion"`

	// User-specified description for the software source.
	Description *string `mandatory:"false" json:"description"`

	// Number of packages the software source contains.
	PackageCount *int64 `mandatory:"false" json:"packageCount"`

	// URI of the GPG key for this software source.
	GpgKeyUrl *string `mandatory:"false" json:"gpgKeyUrl"`

	// ID of the GPG key for this software source.
	GpgKeyId *string `mandatory:"false" json:"gpgKeyId"`

	// Fingerprint of the GPG key for this software source.
	GpgKeyFingerprint *string `mandatory:"false" json:"gpgKeyFingerprint"`

	// The size of the software source in bytes (B).
	Size *float64 `mandatory:"false" json:"size"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	CustomSoftwareSourceFilter *CustomSoftwareSourceFilter `mandatory:"false" json:"customSoftwareSourceFilter"`

	// Indicates whether the service should automatically resolve package dependencies when including specific packages in the software source.
	IsAutoResolveDependencies *bool `mandatory:"false" json:"isAutoResolveDependencies"`

	// Indicates whether the service should create the software source from a list of packages provided by the user.
	IsCreatedFromPackageList *bool `mandatory:"false" json:"isCreatedFromPackageList"`

	// Indicates whether the software source will include only the latest versions of content from vendor software sources, while accounting for other constraints set in the custom or versioned custom software source (such as a package list or filters).
	// * For a module filter that does not specify a stream, this will include all available streams, and within each stream only the latest version of packages.
	// * For a module filter that does specify a stream, this will include only the latest version of packages for the specified stream.
	// * For a package filter that does not specify a version, this will include only the latest available version of the package.
	// * For a package filter that does specify a version, this will include only the specified version of the package (the isLatestContentOnly attribute is ignored).
	// * For a package list, this will include only the specified version of packages and modules in the list (the isLatestContentOnly attribute is ignored).
	IsLatestContentOnly *bool `mandatory:"false" json:"isLatestContentOnly"`

	// The packages in the software source.
	Packages []string `mandatory:"false" json:"packages"`

	// The date and time the metadata for this software source was last updated (in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) format).
	TimeMetadataUpdated *common.SDKTime `mandatory:"false" json:"timeMetadataUpdated"`

	// Availability of the software source (for non-OCI environments).
	Availability AvailabilityEnum `mandatory:"true" json:"availability"`

	// Availability of the software source (for OCI environments).
	AvailabilityAtOci AvailabilityEnum `mandatory:"true" json:"availabilityAtOci"`

	// The OS family of the software source.
	OsFamily OsFamilyEnum `mandatory:"true" json:"osFamily"`

	// The architecture type supported by the software source.
	ArchType ArchTypeEnum `mandatory:"true" json:"archType"`

	// The current state of the software source.
	LifecycleState SoftwareSourceLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The yum repository checksum type used by this software source.
	ChecksumType ChecksumTypeEnum `mandatory:"false" json:"checksumType,omitempty"`

	// Identifies how the versioned custom software source was created.
	SoftwareSourceSubType SoftwareSourceSubTypeEnum `mandatory:"false" json:"softwareSourceSubType,omitempty"`
}

// GetId returns Id
func (m VersionedCustomSoftwareSource) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m VersionedCustomSoftwareSource) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m VersionedCustomSoftwareSource) GetDisplayName() *string {
	return m.DisplayName
}

// GetTimeCreated returns TimeCreated
func (m VersionedCustomSoftwareSource) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetDescription returns Description
func (m VersionedCustomSoftwareSource) GetDescription() *string {
	return m.Description
}

// GetAvailability returns Availability
func (m VersionedCustomSoftwareSource) GetAvailability() AvailabilityEnum {
	return m.Availability
}

// GetAvailabilityAtOci returns AvailabilityAtOci
func (m VersionedCustomSoftwareSource) GetAvailabilityAtOci() AvailabilityEnum {
	return m.AvailabilityAtOci
}

// GetRepoId returns RepoId
func (m VersionedCustomSoftwareSource) GetRepoId() *string {
	return m.RepoId
}

// GetOsFamily returns OsFamily
func (m VersionedCustomSoftwareSource) GetOsFamily() OsFamilyEnum {
	return m.OsFamily
}

// GetArchType returns ArchType
func (m VersionedCustomSoftwareSource) GetArchType() ArchTypeEnum {
	return m.ArchType
}

// GetLifecycleState returns LifecycleState
func (m VersionedCustomSoftwareSource) GetLifecycleState() SoftwareSourceLifecycleStateEnum {
	return m.LifecycleState
}

// GetPackageCount returns PackageCount
func (m VersionedCustomSoftwareSource) GetPackageCount() *int64 {
	return m.PackageCount
}

// GetUrl returns Url
func (m VersionedCustomSoftwareSource) GetUrl() *string {
	return m.Url
}

// GetChecksumType returns ChecksumType
func (m VersionedCustomSoftwareSource) GetChecksumType() ChecksumTypeEnum {
	return m.ChecksumType
}

// GetGpgKeyUrl returns GpgKeyUrl
func (m VersionedCustomSoftwareSource) GetGpgKeyUrl() *string {
	return m.GpgKeyUrl
}

// GetGpgKeyId returns GpgKeyId
func (m VersionedCustomSoftwareSource) GetGpgKeyId() *string {
	return m.GpgKeyId
}

// GetGpgKeyFingerprint returns GpgKeyFingerprint
func (m VersionedCustomSoftwareSource) GetGpgKeyFingerprint() *string {
	return m.GpgKeyFingerprint
}

// GetSize returns Size
func (m VersionedCustomSoftwareSource) GetSize() *float64 {
	return m.Size
}

// GetFreeformTags returns FreeformTags
func (m VersionedCustomSoftwareSource) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m VersionedCustomSoftwareSource) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m VersionedCustomSoftwareSource) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m VersionedCustomSoftwareSource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VersionedCustomSoftwareSource) ValidateEnumValue() (bool, error) {
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
	if _, ok := GetMappingChecksumTypeEnum(string(m.ChecksumType)); !ok && m.ChecksumType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ChecksumType: %s. Supported values are: %s.", m.ChecksumType, strings.Join(GetChecksumTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSoftwareSourceSubTypeEnum(string(m.SoftwareSourceSubType)); !ok && m.SoftwareSourceSubType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SoftwareSourceSubType: %s. Supported values are: %s.", m.SoftwareSourceSubType, strings.Join(GetSoftwareSourceSubTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m VersionedCustomSoftwareSource) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeVersionedCustomSoftwareSource VersionedCustomSoftwareSource
	s := struct {
		DiscriminatorParam string `json:"softwareSourceType"`
		MarshalTypeVersionedCustomSoftwareSource
	}{
		"VERSIONED",
		(MarshalTypeVersionedCustomSoftwareSource)(m),
	}

	return json.Marshal(&s)
}
