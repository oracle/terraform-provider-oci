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

// Profile Description of registration profile.
type Profile interface {

	// The OCID of the profile that is immutable on creation.
	GetId() *string

	// The OCID of the tenancy containing the registration profile.
	GetCompartmentId() *string

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	GetDisplayName() *string

	// The software source vendor name.
	GetVendorName() VendorNameEnum

	// The operating system family.
	GetOsFamily() OsFamilyEnum

	// The architecture type.
	GetArchType() ArchTypeEnum

	// The description of the registration profile.
	GetDescription() *string

	// The OCID of the management station.
	GetManagementStationId() *string

	// The time the the registration profile was created. An RFC3339 formatted datetime string.
	GetTimeCreated() *common.SDKTime

	// The current state of the registration profile.
	GetLifecycleState() ProfileLifecycleStateEnum

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

type profile struct {
	JsonData            []byte
	Description         *string                           `mandatory:"false" json:"description"`
	ManagementStationId *string                           `mandatory:"false" json:"managementStationId"`
	TimeCreated         *common.SDKTime                   `mandatory:"false" json:"timeCreated"`
	LifecycleState      ProfileLifecycleStateEnum         `mandatory:"false" json:"lifecycleState,omitempty"`
	FreeformTags        map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags         map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags          map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	Id                  *string                           `mandatory:"true" json:"id"`
	CompartmentId       *string                           `mandatory:"true" json:"compartmentId"`
	DisplayName         *string                           `mandatory:"true" json:"displayName"`
	VendorName          VendorNameEnum                    `mandatory:"true" json:"vendorName"`
	OsFamily            OsFamilyEnum                      `mandatory:"true" json:"osFamily"`
	ArchType            ArchTypeEnum                      `mandatory:"true" json:"archType"`
	ProfileType         string                            `json:"profileType"`
}

// UnmarshalJSON unmarshals json
func (m *profile) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerprofile profile
	s := struct {
		Model Unmarshalerprofile
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.CompartmentId = s.Model.CompartmentId
	m.DisplayName = s.Model.DisplayName
	m.VendorName = s.Model.VendorName
	m.OsFamily = s.Model.OsFamily
	m.ArchType = s.Model.ArchType
	m.Description = s.Model.Description
	m.ManagementStationId = s.Model.ManagementStationId
	m.TimeCreated = s.Model.TimeCreated
	m.LifecycleState = s.Model.LifecycleState
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.ProfileType = s.Model.ProfileType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *profile) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ProfileType {
	case "LIFECYCLE":
		mm := LifecycleProfile{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SOFTWARESOURCE":
		mm := SoftwareSourceProfile{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GROUP":
		mm := GroupProfile{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "STATION":
		mm := StationProfile{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for Profile: %s.", m.ProfileType)
		return *m, nil
	}
}

// GetDescription returns Description
func (m profile) GetDescription() *string {
	return m.Description
}

// GetManagementStationId returns ManagementStationId
func (m profile) GetManagementStationId() *string {
	return m.ManagementStationId
}

// GetTimeCreated returns TimeCreated
func (m profile) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetLifecycleState returns LifecycleState
func (m profile) GetLifecycleState() ProfileLifecycleStateEnum {
	return m.LifecycleState
}

// GetFreeformTags returns FreeformTags
func (m profile) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m profile) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m profile) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetId returns Id
func (m profile) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m profile) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m profile) GetDisplayName() *string {
	return m.DisplayName
}

// GetVendorName returns VendorName
func (m profile) GetVendorName() VendorNameEnum {
	return m.VendorName
}

// GetOsFamily returns OsFamily
func (m profile) GetOsFamily() OsFamilyEnum {
	return m.OsFamily
}

// GetArchType returns ArchType
func (m profile) GetArchType() ArchTypeEnum {
	return m.ArchType
}

func (m profile) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m profile) ValidateEnumValue() (bool, error) {
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

// ProfileLifecycleStateEnum Enum with underlying type: string
type ProfileLifecycleStateEnum string

// Set of constants representing the allowable values for ProfileLifecycleStateEnum
const (
	ProfileLifecycleStateCreating ProfileLifecycleStateEnum = "CREATING"
	ProfileLifecycleStateUpdating ProfileLifecycleStateEnum = "UPDATING"
	ProfileLifecycleStateActive   ProfileLifecycleStateEnum = "ACTIVE"
	ProfileLifecycleStateDeleting ProfileLifecycleStateEnum = "DELETING"
	ProfileLifecycleStateDeleted  ProfileLifecycleStateEnum = "DELETED"
	ProfileLifecycleStateFailed   ProfileLifecycleStateEnum = "FAILED"
)

var mappingProfileLifecycleStateEnum = map[string]ProfileLifecycleStateEnum{
	"CREATING": ProfileLifecycleStateCreating,
	"UPDATING": ProfileLifecycleStateUpdating,
	"ACTIVE":   ProfileLifecycleStateActive,
	"DELETING": ProfileLifecycleStateDeleting,
	"DELETED":  ProfileLifecycleStateDeleted,
	"FAILED":   ProfileLifecycleStateFailed,
}

var mappingProfileLifecycleStateEnumLowerCase = map[string]ProfileLifecycleStateEnum{
	"creating": ProfileLifecycleStateCreating,
	"updating": ProfileLifecycleStateUpdating,
	"active":   ProfileLifecycleStateActive,
	"deleting": ProfileLifecycleStateDeleting,
	"deleted":  ProfileLifecycleStateDeleted,
	"failed":   ProfileLifecycleStateFailed,
}

// GetProfileLifecycleStateEnumValues Enumerates the set of values for ProfileLifecycleStateEnum
func GetProfileLifecycleStateEnumValues() []ProfileLifecycleStateEnum {
	values := make([]ProfileLifecycleStateEnum, 0)
	for _, v := range mappingProfileLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetProfileLifecycleStateEnumStringValues Enumerates the set of values in String for ProfileLifecycleStateEnum
func GetProfileLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingProfileLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProfileLifecycleStateEnum(val string) (ProfileLifecycleStateEnum, bool) {
	enum, ok := mappingProfileLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
