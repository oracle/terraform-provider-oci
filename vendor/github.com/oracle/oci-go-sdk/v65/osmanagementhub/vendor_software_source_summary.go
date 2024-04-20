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

// VendorSoftwareSourceSummary Provides summary information for a vendor software source.
type VendorSoftwareSourceSummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the software source.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// User-friendly name for the software source.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The repository ID for the software source.
	RepoId *string `mandatory:"true" json:"repoId"`

	// URL for the repository. For vendor software sources, this is the URL to the regional yum server. For custom software sources, this is 'custom/<repoId>'.
	Url *string `mandatory:"true" json:"url"`

	// The date and time the software source was created (in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) format).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the software source was updated (in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) format).
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Description of the software source. For custom software sources, this is user-specified.
	Description *string `mandatory:"false" json:"description"`

	// Number of packages the software source contains.
	PackageCount *int64 `mandatory:"false" json:"packageCount"`

	// The size of the software source in gigabytes (GB).
	Size *float64 `mandatory:"false" json:"size"`

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

	// Indicates whether the software source is required for the Autonomous Linux service.
	IsMandatoryForAutonomousLinux *bool `mandatory:"false" json:"isMandatoryForAutonomousLinux"`

	// The current state of the software source.
	LifecycleState SoftwareSourceLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Availability of the software source (for non-OCI environments).
	Availability AvailabilityEnum `mandatory:"true" json:"availability"`

	// Availability of the software source (for OCI environments).
	AvailabilityAtOci AvailabilityEnum `mandatory:"true" json:"availabilityAtOci"`

	// The OS family the software source belongs to.
	OsFamily OsFamilyEnum `mandatory:"true" json:"osFamily"`

	// The architecture type supported by the software source.
	ArchType ArchTypeEnum `mandatory:"true" json:"archType"`

	// Name of the vendor providing the software source.
	VendorName VendorNameEnum `mandatory:"true" json:"vendorName"`
}

// GetId returns Id
func (m VendorSoftwareSourceSummary) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m VendorSoftwareSourceSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m VendorSoftwareSourceSummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetRepoId returns RepoId
func (m VendorSoftwareSourceSummary) GetRepoId() *string {
	return m.RepoId
}

// GetUrl returns Url
func (m VendorSoftwareSourceSummary) GetUrl() *string {
	return m.Url
}

// GetTimeCreated returns TimeCreated
func (m VendorSoftwareSourceSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m VendorSoftwareSourceSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetDescription returns Description
func (m VendorSoftwareSourceSummary) GetDescription() *string {
	return m.Description
}

// GetAvailability returns Availability
func (m VendorSoftwareSourceSummary) GetAvailability() AvailabilityEnum {
	return m.Availability
}

// GetAvailabilityAtOci returns AvailabilityAtOci
func (m VendorSoftwareSourceSummary) GetAvailabilityAtOci() AvailabilityEnum {
	return m.AvailabilityAtOci
}

// GetOsFamily returns OsFamily
func (m VendorSoftwareSourceSummary) GetOsFamily() OsFamilyEnum {
	return m.OsFamily
}

// GetArchType returns ArchType
func (m VendorSoftwareSourceSummary) GetArchType() ArchTypeEnum {
	return m.ArchType
}

// GetPackageCount returns PackageCount
func (m VendorSoftwareSourceSummary) GetPackageCount() *int64 {
	return m.PackageCount
}

// GetLifecycleState returns LifecycleState
func (m VendorSoftwareSourceSummary) GetLifecycleState() SoftwareSourceLifecycleStateEnum {
	return m.LifecycleState
}

// GetSize returns Size
func (m VendorSoftwareSourceSummary) GetSize() *float64 {
	return m.Size
}

// GetFreeformTags returns FreeformTags
func (m VendorSoftwareSourceSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m VendorSoftwareSourceSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m VendorSoftwareSourceSummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m VendorSoftwareSourceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VendorSoftwareSourceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSoftwareSourceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSoftwareSourceLifecycleStateEnumStringValues(), ",")))
	}
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
	if _, ok := GetMappingVendorNameEnum(string(m.VendorName)); !ok && m.VendorName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VendorName: %s. Supported values are: %s.", m.VendorName, strings.Join(GetVendorNameEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m VendorSoftwareSourceSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeVendorSoftwareSourceSummary VendorSoftwareSourceSummary
	s := struct {
		DiscriminatorParam string `json:"softwareSourceType"`
		MarshalTypeVendorSoftwareSourceSummary
	}{
		"VENDOR",
		(MarshalTypeVendorSoftwareSourceSummary)(m),
	}

	return json.Marshal(&s)
}
