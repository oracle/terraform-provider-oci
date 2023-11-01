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

// StationProfile Definition of a registration profile of type STATION.
type StationProfile struct {

	// The OCID of the profile that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the tenancy containing the registration profile.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The description of the registration profile.
	Description *string `mandatory:"false" json:"description"`

	// The OCID of the management station.
	ManagementStationId *string `mandatory:"false" json:"managementStationId"`

	// The time the the registration profile was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

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

	// The software source vendor name.
	VendorName VendorNameEnum `mandatory:"true" json:"vendorName"`

	// The operating system family.
	OsFamily OsFamilyEnum `mandatory:"true" json:"osFamily"`

	// The architecture type.
	ArchType ArchTypeEnum `mandatory:"true" json:"archType"`

	// The current state of the registration profile.
	LifecycleState ProfileLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

// GetId returns Id
func (m StationProfile) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m StationProfile) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m StationProfile) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m StationProfile) GetDescription() *string {
	return m.Description
}

// GetManagementStationId returns ManagementStationId
func (m StationProfile) GetManagementStationId() *string {
	return m.ManagementStationId
}

// GetVendorName returns VendorName
func (m StationProfile) GetVendorName() VendorNameEnum {
	return m.VendorName
}

// GetOsFamily returns OsFamily
func (m StationProfile) GetOsFamily() OsFamilyEnum {
	return m.OsFamily
}

// GetArchType returns ArchType
func (m StationProfile) GetArchType() ArchTypeEnum {
	return m.ArchType
}

// GetTimeCreated returns TimeCreated
func (m StationProfile) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetLifecycleState returns LifecycleState
func (m StationProfile) GetLifecycleState() ProfileLifecycleStateEnum {
	return m.LifecycleState
}

// GetFreeformTags returns FreeformTags
func (m StationProfile) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m StationProfile) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m StationProfile) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m StationProfile) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StationProfile) ValidateEnumValue() (bool, error) {
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
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m StationProfile) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeStationProfile StationProfile
	s := struct {
		DiscriminatorParam string `json:"profileType"`
		MarshalTypeStationProfile
	}{
		"STATION",
		(MarshalTypeStationProfile)(m),
	}

	return json.Marshal(&s)
}
