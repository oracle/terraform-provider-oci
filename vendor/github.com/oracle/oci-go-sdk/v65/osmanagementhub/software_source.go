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

// SoftwareSource A software source contains a collection of packages.
type SoftwareSource interface {

	// OCID for the software source.
	GetId() *string

	// The OCID of the tenancy containing the software source.
	GetCompartmentId() *string

	// User friendly name for the software source.
	GetDisplayName() *string

	// The date and time the software source was created, as described in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	GetTimeCreated() *common.SDKTime

	// Possible availabilities of a software source.
	GetAvailability() AvailabilityEnum

	// The Repo ID for the software source.
	GetRepoId() *string

	// The OS family the software source belongs to.
	GetOsFamily() OsFamilyEnum

	// The architecture type supported by the software source.
	GetArchType() ArchTypeEnum

	// URL for the repository.
	GetUrl() *string

	// Information specified by the user about the software source.
	GetDescription() *string

	// The current state of the software source.
	GetLifecycleState() SoftwareSourceLifecycleStateEnum

	// Number of packages.
	GetPackageCount() *int64

	// The yum repository checksum type used by this software source.
	GetChecksumType() ChecksumTypeEnum

	// URL of the GPG key for this software source.
	GetGpgKeyUrl() *string

	// ID of the GPG key for this software source.
	GetGpgKeyId() *string

	// Fingerprint of the GPG key for this software source.
	GetGpgKeyFingerprint() *string

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

type softwaresource struct {
	JsonData           []byte
	Description        *string                           `mandatory:"false" json:"description"`
	LifecycleState     SoftwareSourceLifecycleStateEnum  `mandatory:"false" json:"lifecycleState,omitempty"`
	PackageCount       *int64                            `mandatory:"false" json:"packageCount"`
	ChecksumType       ChecksumTypeEnum                  `mandatory:"false" json:"checksumType,omitempty"`
	GpgKeyUrl          *string                           `mandatory:"false" json:"gpgKeyUrl"`
	GpgKeyId           *string                           `mandatory:"false" json:"gpgKeyId"`
	GpgKeyFingerprint  *string                           `mandatory:"false" json:"gpgKeyFingerprint"`
	FreeformTags       map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags        map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags         map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	Id                 *string                           `mandatory:"true" json:"id"`
	CompartmentId      *string                           `mandatory:"true" json:"compartmentId"`
	DisplayName        *string                           `mandatory:"true" json:"displayName"`
	TimeCreated        *common.SDKTime                   `mandatory:"true" json:"timeCreated"`
	Availability       AvailabilityEnum                  `mandatory:"true" json:"availability"`
	RepoId             *string                           `mandatory:"true" json:"repoId"`
	OsFamily           OsFamilyEnum                      `mandatory:"true" json:"osFamily"`
	ArchType           ArchTypeEnum                      `mandatory:"true" json:"archType"`
	Url                *string                           `mandatory:"true" json:"url"`
	SoftwareSourceType string                            `json:"softwareSourceType"`
}

// UnmarshalJSON unmarshals json
func (m *softwaresource) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersoftwaresource softwaresource
	s := struct {
		Model Unmarshalersoftwaresource
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.CompartmentId = s.Model.CompartmentId
	m.DisplayName = s.Model.DisplayName
	m.TimeCreated = s.Model.TimeCreated
	m.Availability = s.Model.Availability
	m.RepoId = s.Model.RepoId
	m.OsFamily = s.Model.OsFamily
	m.ArchType = s.Model.ArchType
	m.Url = s.Model.Url
	m.Description = s.Model.Description
	m.LifecycleState = s.Model.LifecycleState
	m.PackageCount = s.Model.PackageCount
	m.ChecksumType = s.Model.ChecksumType
	m.GpgKeyUrl = s.Model.GpgKeyUrl
	m.GpgKeyId = s.Model.GpgKeyId
	m.GpgKeyFingerprint = s.Model.GpgKeyFingerprint
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.SoftwareSourceType = s.Model.SoftwareSourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *softwaresource) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SoftwareSourceType {
	case "VENDOR":
		mm := VendorSoftwareSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CUSTOM":
		mm := CustomSoftwareSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VERSIONED":
		mm := VersionedCustomSoftwareSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for SoftwareSource: %s.", m.SoftwareSourceType)
		return *m, nil
	}
}

// GetDescription returns Description
func (m softwaresource) GetDescription() *string {
	return m.Description
}

// GetLifecycleState returns LifecycleState
func (m softwaresource) GetLifecycleState() SoftwareSourceLifecycleStateEnum {
	return m.LifecycleState
}

// GetPackageCount returns PackageCount
func (m softwaresource) GetPackageCount() *int64 {
	return m.PackageCount
}

// GetChecksumType returns ChecksumType
func (m softwaresource) GetChecksumType() ChecksumTypeEnum {
	return m.ChecksumType
}

// GetGpgKeyUrl returns GpgKeyUrl
func (m softwaresource) GetGpgKeyUrl() *string {
	return m.GpgKeyUrl
}

// GetGpgKeyId returns GpgKeyId
func (m softwaresource) GetGpgKeyId() *string {
	return m.GpgKeyId
}

// GetGpgKeyFingerprint returns GpgKeyFingerprint
func (m softwaresource) GetGpgKeyFingerprint() *string {
	return m.GpgKeyFingerprint
}

// GetFreeformTags returns FreeformTags
func (m softwaresource) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m softwaresource) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m softwaresource) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetId returns Id
func (m softwaresource) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m softwaresource) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m softwaresource) GetDisplayName() *string {
	return m.DisplayName
}

// GetTimeCreated returns TimeCreated
func (m softwaresource) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetAvailability returns Availability
func (m softwaresource) GetAvailability() AvailabilityEnum {
	return m.Availability
}

// GetRepoId returns RepoId
func (m softwaresource) GetRepoId() *string {
	return m.RepoId
}

// GetOsFamily returns OsFamily
func (m softwaresource) GetOsFamily() OsFamilyEnum {
	return m.OsFamily
}

// GetArchType returns ArchType
func (m softwaresource) GetArchType() ArchTypeEnum {
	return m.ArchType
}

// GetUrl returns Url
func (m softwaresource) GetUrl() *string {
	return m.Url
}

func (m softwaresource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m softwaresource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAvailabilityEnum(string(m.Availability)); !ok && m.Availability != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Availability: %s. Supported values are: %s.", m.Availability, strings.Join(GetAvailabilityEnumStringValues(), ",")))
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
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SoftwareSourceLifecycleStateEnum Enum with underlying type: string
type SoftwareSourceLifecycleStateEnum string

// Set of constants representing the allowable values for SoftwareSourceLifecycleStateEnum
const (
	SoftwareSourceLifecycleStateCreating SoftwareSourceLifecycleStateEnum = "CREATING"
	SoftwareSourceLifecycleStateUpdating SoftwareSourceLifecycleStateEnum = "UPDATING"
	SoftwareSourceLifecycleStateActive   SoftwareSourceLifecycleStateEnum = "ACTIVE"
	SoftwareSourceLifecycleStateDeleting SoftwareSourceLifecycleStateEnum = "DELETING"
	SoftwareSourceLifecycleStateDeleted  SoftwareSourceLifecycleStateEnum = "DELETED"
	SoftwareSourceLifecycleStateFailed   SoftwareSourceLifecycleStateEnum = "FAILED"
)

var mappingSoftwareSourceLifecycleStateEnum = map[string]SoftwareSourceLifecycleStateEnum{
	"CREATING": SoftwareSourceLifecycleStateCreating,
	"UPDATING": SoftwareSourceLifecycleStateUpdating,
	"ACTIVE":   SoftwareSourceLifecycleStateActive,
	"DELETING": SoftwareSourceLifecycleStateDeleting,
	"DELETED":  SoftwareSourceLifecycleStateDeleted,
	"FAILED":   SoftwareSourceLifecycleStateFailed,
}

var mappingSoftwareSourceLifecycleStateEnumLowerCase = map[string]SoftwareSourceLifecycleStateEnum{
	"creating": SoftwareSourceLifecycleStateCreating,
	"updating": SoftwareSourceLifecycleStateUpdating,
	"active":   SoftwareSourceLifecycleStateActive,
	"deleting": SoftwareSourceLifecycleStateDeleting,
	"deleted":  SoftwareSourceLifecycleStateDeleted,
	"failed":   SoftwareSourceLifecycleStateFailed,
}

// GetSoftwareSourceLifecycleStateEnumValues Enumerates the set of values for SoftwareSourceLifecycleStateEnum
func GetSoftwareSourceLifecycleStateEnumValues() []SoftwareSourceLifecycleStateEnum {
	values := make([]SoftwareSourceLifecycleStateEnum, 0)
	for _, v := range mappingSoftwareSourceLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSoftwareSourceLifecycleStateEnumStringValues Enumerates the set of values in String for SoftwareSourceLifecycleStateEnum
func GetSoftwareSourceLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingSoftwareSourceLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSoftwareSourceLifecycleStateEnum(val string) (SoftwareSourceLifecycleStateEnum, bool) {
	enum, ok := mappingSoftwareSourceLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
