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

// CreateStationProfileDetails Provides the information used to create the management station profile.
type CreateStationProfileDetails struct {

	// A user-friendly name. Does not have to be unique and you can change the name later. Avoid entering
	// confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the registration profile.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// User-specified description of the registration profile.
	Description *string `mandatory:"false" json:"description"`

	// description: The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management station to associate
	// with an instance once registered. This is required when creating a profile for non-OCI instances.
	ManagementStationId *string `mandatory:"false" json:"managementStationId"`

	// Indicates if the profile is set as the default. There is exactly one default profile for a specified architecture, OS family, registration type, and vendor. When registering an instance with the corresonding characteristics, the default profile is used, unless another profile is specified.
	IsDefaultProfile *bool `mandatory:"false" json:"isDefaultProfile"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The type of instance to register.
	RegistrationType ProfileRegistrationTypeEnum `mandatory:"false" json:"registrationType,omitempty"`

	// The vendor of the operating system for the instance.
	VendorName VendorNameEnum `mandatory:"false" json:"vendorName,omitempty"`

	// The operating system family.
	OsFamily OsFamilyEnum `mandatory:"false" json:"osFamily,omitempty"`

	// The architecture type.
	ArchType ArchTypeEnum `mandatory:"false" json:"archType,omitempty"`
}

// GetDisplayName returns DisplayName
func (m CreateStationProfileDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m CreateStationProfileDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDescription returns Description
func (m CreateStationProfileDetails) GetDescription() *string {
	return m.Description
}

// GetManagementStationId returns ManagementStationId
func (m CreateStationProfileDetails) GetManagementStationId() *string {
	return m.ManagementStationId
}

// GetRegistrationType returns RegistrationType
func (m CreateStationProfileDetails) GetRegistrationType() ProfileRegistrationTypeEnum {
	return m.RegistrationType
}

// GetIsDefaultProfile returns IsDefaultProfile
func (m CreateStationProfileDetails) GetIsDefaultProfile() *bool {
	return m.IsDefaultProfile
}

// GetFreeformTags returns FreeformTags
func (m CreateStationProfileDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateStationProfileDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateStationProfileDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateStationProfileDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingProfileRegistrationTypeEnum(string(m.RegistrationType)); !ok && m.RegistrationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RegistrationType: %s. Supported values are: %s.", m.RegistrationType, strings.Join(GetProfileRegistrationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingVendorNameEnum(string(m.VendorName)); !ok && m.VendorName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VendorName: %s. Supported values are: %s.", m.VendorName, strings.Join(GetVendorNameEnumStringValues(), ",")))
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
func (m CreateStationProfileDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateStationProfileDetails CreateStationProfileDetails
	s := struct {
		DiscriminatorParam string `json:"profileType"`
		MarshalTypeCreateStationProfileDetails
	}{
		"STATION",
		(MarshalTypeCreateStationProfileDetails)(m),
	}

	return json.Marshal(&s)
}
