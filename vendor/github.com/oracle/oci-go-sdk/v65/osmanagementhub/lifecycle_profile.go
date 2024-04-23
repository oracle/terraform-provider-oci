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

// LifecycleProfile Provides the information for a lifecycle environment registration profile.
type LifecycleProfile struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the registration profile.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the registration profile.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name for the profile.
	DisplayName *string `mandatory:"true" json:"displayName"`

	LifecycleStage *LifecycleStageDetails `mandatory:"true" json:"lifecycleStage"`

	// The description of the registration profile.
	Description *string `mandatory:"false" json:"description"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management station to associate with an instance once registered. Associating with a management station applies only to non-OCI instances.
	ManagementStationId *string `mandatory:"false" json:"managementStationId"`

	// The time the registration profile was created (in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) format).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Indicates if the profile is set as the default. There is exactly one default profile for a specified architecture, OS family, registration type, and vendor. When registering an instance with the corresonding characteristics, the default profile is used, unless another profile is specified.
	IsDefaultProfile *bool `mandatory:"false" json:"isDefaultProfile"`

	// Indicates if the profile was created by the service. OS Management Hub provides a limited set of standardized profiles that can be used to register Autonomous Linux or Windows instances.
	IsServiceProvidedProfile *bool `mandatory:"false" json:"isServiceProvidedProfile"`

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

	LifecycleEnvironment *LifecycleEnvironmentDetails `mandatory:"false" json:"lifecycleEnvironment"`

	// The vendor of the operating system for the instance.
	VendorName VendorNameEnum `mandatory:"true" json:"vendorName"`

	// The operating system family.
	OsFamily OsFamilyEnum `mandatory:"true" json:"osFamily"`

	// The architecture type.
	ArchType ArchTypeEnum `mandatory:"true" json:"archType"`

	// The current state of the registration profile.
	LifecycleState ProfileLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The type of instance to register.
	RegistrationType ProfileRegistrationTypeEnum `mandatory:"false" json:"registrationType,omitempty"`
}

// GetId returns Id
func (m LifecycleProfile) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m LifecycleProfile) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m LifecycleProfile) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m LifecycleProfile) GetDescription() *string {
	return m.Description
}

// GetManagementStationId returns ManagementStationId
func (m LifecycleProfile) GetManagementStationId() *string {
	return m.ManagementStationId
}

// GetVendorName returns VendorName
func (m LifecycleProfile) GetVendorName() VendorNameEnum {
	return m.VendorName
}

// GetOsFamily returns OsFamily
func (m LifecycleProfile) GetOsFamily() OsFamilyEnum {
	return m.OsFamily
}

// GetArchType returns ArchType
func (m LifecycleProfile) GetArchType() ArchTypeEnum {
	return m.ArchType
}

// GetTimeCreated returns TimeCreated
func (m LifecycleProfile) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetLifecycleState returns LifecycleState
func (m LifecycleProfile) GetLifecycleState() ProfileLifecycleStateEnum {
	return m.LifecycleState
}

// GetRegistrationType returns RegistrationType
func (m LifecycleProfile) GetRegistrationType() ProfileRegistrationTypeEnum {
	return m.RegistrationType
}

// GetIsDefaultProfile returns IsDefaultProfile
func (m LifecycleProfile) GetIsDefaultProfile() *bool {
	return m.IsDefaultProfile
}

// GetIsServiceProvidedProfile returns IsServiceProvidedProfile
func (m LifecycleProfile) GetIsServiceProvidedProfile() *bool {
	return m.IsServiceProvidedProfile
}

// GetFreeformTags returns FreeformTags
func (m LifecycleProfile) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m LifecycleProfile) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m LifecycleProfile) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m LifecycleProfile) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LifecycleProfile) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingVendorNameEnum(string(m.VendorName)); !ok && m.VendorName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VendorName: %s. Supported values are: %s.", m.VendorName, strings.Join(GetVendorNameEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOsFamilyEnum(string(m.OsFamily)); !ok && m.OsFamily != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsFamily: %s. Supported values are: %s.", m.OsFamily, strings.Join(GetOsFamilyEnumStringValues(), ",")))
	}
	if _, ok := GetMappingArchTypeEnum(string(m.ArchType)); !ok && m.ArchType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ArchType: %s. Supported values are: %s.", m.ArchType, strings.Join(GetArchTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingProfileLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetProfileLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingProfileRegistrationTypeEnum(string(m.RegistrationType)); !ok && m.RegistrationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RegistrationType: %s. Supported values are: %s.", m.RegistrationType, strings.Join(GetProfileRegistrationTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m LifecycleProfile) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeLifecycleProfile LifecycleProfile
	s := struct {
		DiscriminatorParam string `json:"profileType"`
		MarshalTypeLifecycleProfile
	}{
		"LIFECYCLE",
		(MarshalTypeLifecycleProfile)(m),
	}

	return json.Marshal(&s)
}
