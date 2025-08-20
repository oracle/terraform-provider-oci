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

// GuestOsResourceIdFilter The OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of associated resources to include when discovering Exadata VM Cluster targets for a 'GUEST_OS' collection.
type GuestOsResourceIdFilter struct {

	// The OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of associated resources to include when discovering Exadata VM Cluster targets for a 'GUEST_OS' collection.
	// Specified resources must match the specified 'entityType'.
	// FsuCollection of type 'GI' or 'GUEST_OS' can be specified.
	Identifiers []string `mandatory:"true" json:"identifiers"`

	// Type of associated resource.
	EntityType GuestOsResourceIdFilterEntityTypeEnum `mandatory:"true" json:"entityType"`

	// INCLUDE or EXCLUDE the filter results when discovering Exadata VM Cluster targets for a 'GUEST_OS' collection.
	// Supported only for RESOURCE_ID filter.
	Mode GuestOsFleetDiscoveryFilterModeEnum `mandatory:"false" json:"mode,omitempty"`

	// Type of join for each element in this filter.
	Operator FleetDiscoveryOperatorsEnum `mandatory:"false" json:"operator,omitempty"`
}

// GetMode returns Mode
func (m GuestOsResourceIdFilter) GetMode() GuestOsFleetDiscoveryFilterModeEnum {
	return m.Mode
}

func (m GuestOsResourceIdFilter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GuestOsResourceIdFilter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGuestOsResourceIdFilterEntityTypeEnum(string(m.EntityType)); !ok && m.EntityType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EntityType: %s. Supported values are: %s.", m.EntityType, strings.Join(GetGuestOsResourceIdFilterEntityTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingGuestOsFleetDiscoveryFilterModeEnum(string(m.Mode)); !ok && m.Mode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Mode: %s. Supported values are: %s.", m.Mode, strings.Join(GetGuestOsFleetDiscoveryFilterModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingFleetDiscoveryOperatorsEnum(string(m.Operator)); !ok && m.Operator != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Operator: %s. Supported values are: %s.", m.Operator, strings.Join(GetFleetDiscoveryOperatorsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m GuestOsResourceIdFilter) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGuestOsResourceIdFilter GuestOsResourceIdFilter
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeGuestOsResourceIdFilter
	}{
		"RESOURCE_ID",
		(MarshalTypeGuestOsResourceIdFilter)(m),
	}

	return json.Marshal(&s)
}

// GuestOsResourceIdFilterEntityTypeEnum Enum with underlying type: string
type GuestOsResourceIdFilterEntityTypeEnum string

// Set of constants representing the allowable values for GuestOsResourceIdFilterEntityTypeEnum
const (
	GuestOsResourceIdFilterEntityTypeExadatainfrastructure      GuestOsResourceIdFilterEntityTypeEnum = "EXADATAINFRASTRUCTURE"
	GuestOsResourceIdFilterEntityTypeCloudexadatainfrastructure GuestOsResourceIdFilterEntityTypeEnum = "CLOUDEXADATAINFRASTRUCTURE"
	GuestOsResourceIdFilterEntityTypeVmcluster                  GuestOsResourceIdFilterEntityTypeEnum = "VMCLUSTER"
	GuestOsResourceIdFilterEntityTypeCloudvmcluster             GuestOsResourceIdFilterEntityTypeEnum = "CLOUDVMCLUSTER"
	GuestOsResourceIdFilterEntityTypeFsucollection              GuestOsResourceIdFilterEntityTypeEnum = "FSUCOLLECTION"
)

var mappingGuestOsResourceIdFilterEntityTypeEnum = map[string]GuestOsResourceIdFilterEntityTypeEnum{
	"EXADATAINFRASTRUCTURE":      GuestOsResourceIdFilterEntityTypeExadatainfrastructure,
	"CLOUDEXADATAINFRASTRUCTURE": GuestOsResourceIdFilterEntityTypeCloudexadatainfrastructure,
	"VMCLUSTER":                  GuestOsResourceIdFilterEntityTypeVmcluster,
	"CLOUDVMCLUSTER":             GuestOsResourceIdFilterEntityTypeCloudvmcluster,
	"FSUCOLLECTION":              GuestOsResourceIdFilterEntityTypeFsucollection,
}

var mappingGuestOsResourceIdFilterEntityTypeEnumLowerCase = map[string]GuestOsResourceIdFilterEntityTypeEnum{
	"exadatainfrastructure":      GuestOsResourceIdFilterEntityTypeExadatainfrastructure,
	"cloudexadatainfrastructure": GuestOsResourceIdFilterEntityTypeCloudexadatainfrastructure,
	"vmcluster":                  GuestOsResourceIdFilterEntityTypeVmcluster,
	"cloudvmcluster":             GuestOsResourceIdFilterEntityTypeCloudvmcluster,
	"fsucollection":              GuestOsResourceIdFilterEntityTypeFsucollection,
}

// GetGuestOsResourceIdFilterEntityTypeEnumValues Enumerates the set of values for GuestOsResourceIdFilterEntityTypeEnum
func GetGuestOsResourceIdFilterEntityTypeEnumValues() []GuestOsResourceIdFilterEntityTypeEnum {
	values := make([]GuestOsResourceIdFilterEntityTypeEnum, 0)
	for _, v := range mappingGuestOsResourceIdFilterEntityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGuestOsResourceIdFilterEntityTypeEnumStringValues Enumerates the set of values in String for GuestOsResourceIdFilterEntityTypeEnum
func GetGuestOsResourceIdFilterEntityTypeEnumStringValues() []string {
	return []string{
		"EXADATAINFRASTRUCTURE",
		"CLOUDEXADATAINFRASTRUCTURE",
		"VMCLUSTER",
		"CLOUDVMCLUSTER",
		"FSUCOLLECTION",
	}
}

// GetMappingGuestOsResourceIdFilterEntityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGuestOsResourceIdFilterEntityTypeEnum(val string) (GuestOsResourceIdFilterEntityTypeEnum, bool) {
	enum, ok := mappingGuestOsResourceIdFilterEntityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
