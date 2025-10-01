// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Exadata Fleet Update service API
//
// Use the Exadata Fleet Update service to patch large collections of components directly,
// as a single entity, orchestrating the maintenance actions to update all chosen components in the stack in a single cycle.
//

package fleetsoftwareupdate

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// GuestOsFleetDiscoveryFilter Discover Exadata VM Cluster targets for a 'GUEST_OS' collection by querying OCI Search Service (https://docs.oracle.com/iaas/Content/Search/Concepts/queryoverview.htm) using specified filters.
type GuestOsFleetDiscoveryFilter interface {

	// INCLUDE or EXCLUDE the filter results when discovering Exadata VM Cluster targets for a 'GUEST_OS' collection.
	// Supported only for RESOURCE_ID filter.
	GetMode() GuestOsFleetDiscoveryFilterModeEnum
}

type guestosfleetdiscoveryfilter struct {
	JsonData []byte
	Mode     GuestOsFleetDiscoveryFilterModeEnum `mandatory:"false" json:"mode,omitempty"`
	Type     string                              `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *guestosfleetdiscoveryfilter) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerguestosfleetdiscoveryfilter guestosfleetdiscoveryfilter
	s := struct {
		Model Unmarshalerguestosfleetdiscoveryfilter
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Mode = s.Model.Mode
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *guestosfleetdiscoveryfilter) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "VERSION":
		mm := GuestOsVersionFilter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "EXADATA_RELEASE_VERSION":
		mm := GuestOsExadataReleaseVersionFilter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FREEFORM_TAG":
		mm := GuestOsFreeformTagsFilter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DEFINED_TAG":
		mm := GuestOsDefinedTagsFilter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMPARTMENT_ID":
		mm := GuestOsCompartmentIdFilter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RESOURCE_ID":
		mm := GuestOsResourceIdFilter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for GuestOsFleetDiscoveryFilter: %s.", m.Type)
		return *m, nil
	}
}

// GetMode returns Mode
func (m guestosfleetdiscoveryfilter) GetMode() GuestOsFleetDiscoveryFilterModeEnum {
	return m.Mode
}

func (m guestosfleetdiscoveryfilter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m guestosfleetdiscoveryfilter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingGuestOsFleetDiscoveryFilterModeEnum(string(m.Mode)); !ok && m.Mode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Mode: %s. Supported values are: %s.", m.Mode, strings.Join(GetGuestOsFleetDiscoveryFilterModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GuestOsFleetDiscoveryFilterModeEnum Enum with underlying type: string
type GuestOsFleetDiscoveryFilterModeEnum string

// Set of constants representing the allowable values for GuestOsFleetDiscoveryFilterModeEnum
const (
	GuestOsFleetDiscoveryFilterModeInclude GuestOsFleetDiscoveryFilterModeEnum = "INCLUDE"
	GuestOsFleetDiscoveryFilterModeExclude GuestOsFleetDiscoveryFilterModeEnum = "EXCLUDE"
)

var mappingGuestOsFleetDiscoveryFilterModeEnum = map[string]GuestOsFleetDiscoveryFilterModeEnum{
	"INCLUDE": GuestOsFleetDiscoveryFilterModeInclude,
	"EXCLUDE": GuestOsFleetDiscoveryFilterModeExclude,
}

var mappingGuestOsFleetDiscoveryFilterModeEnumLowerCase = map[string]GuestOsFleetDiscoveryFilterModeEnum{
	"include": GuestOsFleetDiscoveryFilterModeInclude,
	"exclude": GuestOsFleetDiscoveryFilterModeExclude,
}

// GetGuestOsFleetDiscoveryFilterModeEnumValues Enumerates the set of values for GuestOsFleetDiscoveryFilterModeEnum
func GetGuestOsFleetDiscoveryFilterModeEnumValues() []GuestOsFleetDiscoveryFilterModeEnum {
	values := make([]GuestOsFleetDiscoveryFilterModeEnum, 0)
	for _, v := range mappingGuestOsFleetDiscoveryFilterModeEnum {
		values = append(values, v)
	}
	return values
}

// GetGuestOsFleetDiscoveryFilterModeEnumStringValues Enumerates the set of values in String for GuestOsFleetDiscoveryFilterModeEnum
func GetGuestOsFleetDiscoveryFilterModeEnumStringValues() []string {
	return []string{
		"INCLUDE",
		"EXCLUDE",
	}
}

// GetMappingGuestOsFleetDiscoveryFilterModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGuestOsFleetDiscoveryFilterModeEnum(val string) (GuestOsFleetDiscoveryFilterModeEnum, bool) {
	enum, ok := mappingGuestOsFleetDiscoveryFilterModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// GuestOsFleetDiscoveryFilterTypeEnum Enum with underlying type: string
type GuestOsFleetDiscoveryFilterTypeEnum string

// Set of constants representing the allowable values for GuestOsFleetDiscoveryFilterTypeEnum
const (
	GuestOsFleetDiscoveryFilterTypeCompartmentId         GuestOsFleetDiscoveryFilterTypeEnum = "COMPARTMENT_ID"
	GuestOsFleetDiscoveryFilterTypeVersion               GuestOsFleetDiscoveryFilterTypeEnum = "VERSION"
	GuestOsFleetDiscoveryFilterTypeExadataReleaseVersion GuestOsFleetDiscoveryFilterTypeEnum = "EXADATA_RELEASE_VERSION"
	GuestOsFleetDiscoveryFilterTypeFreeformTag           GuestOsFleetDiscoveryFilterTypeEnum = "FREEFORM_TAG"
	GuestOsFleetDiscoveryFilterTypeDefinedTag            GuestOsFleetDiscoveryFilterTypeEnum = "DEFINED_TAG"
	GuestOsFleetDiscoveryFilterTypeResourceId            GuestOsFleetDiscoveryFilterTypeEnum = "RESOURCE_ID"
)

var mappingGuestOsFleetDiscoveryFilterTypeEnum = map[string]GuestOsFleetDiscoveryFilterTypeEnum{
	"COMPARTMENT_ID":          GuestOsFleetDiscoveryFilterTypeCompartmentId,
	"VERSION":                 GuestOsFleetDiscoveryFilterTypeVersion,
	"EXADATA_RELEASE_VERSION": GuestOsFleetDiscoveryFilterTypeExadataReleaseVersion,
	"FREEFORM_TAG":            GuestOsFleetDiscoveryFilterTypeFreeformTag,
	"DEFINED_TAG":             GuestOsFleetDiscoveryFilterTypeDefinedTag,
	"RESOURCE_ID":             GuestOsFleetDiscoveryFilterTypeResourceId,
}

var mappingGuestOsFleetDiscoveryFilterTypeEnumLowerCase = map[string]GuestOsFleetDiscoveryFilterTypeEnum{
	"compartment_id":          GuestOsFleetDiscoveryFilterTypeCompartmentId,
	"version":                 GuestOsFleetDiscoveryFilterTypeVersion,
	"exadata_release_version": GuestOsFleetDiscoveryFilterTypeExadataReleaseVersion,
	"freeform_tag":            GuestOsFleetDiscoveryFilterTypeFreeformTag,
	"defined_tag":             GuestOsFleetDiscoveryFilterTypeDefinedTag,
	"resource_id":             GuestOsFleetDiscoveryFilterTypeResourceId,
}

// GetGuestOsFleetDiscoveryFilterTypeEnumValues Enumerates the set of values for GuestOsFleetDiscoveryFilterTypeEnum
func GetGuestOsFleetDiscoveryFilterTypeEnumValues() []GuestOsFleetDiscoveryFilterTypeEnum {
	values := make([]GuestOsFleetDiscoveryFilterTypeEnum, 0)
	for _, v := range mappingGuestOsFleetDiscoveryFilterTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGuestOsFleetDiscoveryFilterTypeEnumStringValues Enumerates the set of values in String for GuestOsFleetDiscoveryFilterTypeEnum
func GetGuestOsFleetDiscoveryFilterTypeEnumStringValues() []string {
	return []string{
		"COMPARTMENT_ID",
		"VERSION",
		"EXADATA_RELEASE_VERSION",
		"FREEFORM_TAG",
		"DEFINED_TAG",
		"RESOURCE_ID",
	}
}

// GetMappingGuestOsFleetDiscoveryFilterTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGuestOsFleetDiscoveryFilterTypeEnum(val string) (GuestOsFleetDiscoveryFilterTypeEnum, bool) {
	enum, ok := mappingGuestOsFleetDiscoveryFilterTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
