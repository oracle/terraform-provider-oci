// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// GiResourceIdFilter Related resource Ids to include in the discovery.
type GiResourceIdFilter struct {

	// Related resource Ids to include in the discovery.
	// All must match the specified entityType.
	Identifiers []string `mandatory:"true" json:"identifiers"`

	// Type of resource to match in the discovery.
	EntityType GiResourceIdFilterEntityTypeEnum `mandatory:"true" json:"entityType"`

	// INCLUDE or EXCLUDE the filter results in the discovery for GI targets.
	// Supported for 'FSUCOLLECTION' RESOURCE_ID filter only.
	Mode GiFleetDiscoveryFilterModeEnum `mandatory:"false" json:"mode,omitempty"`

	// Type of join for each element in this filter.
	Operator FleetDiscoveryOperatorsEnum `mandatory:"false" json:"operator,omitempty"`
}

// GetMode returns Mode
func (m GiResourceIdFilter) GetMode() GiFleetDiscoveryFilterModeEnum {
	return m.Mode
}

func (m GiResourceIdFilter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GiResourceIdFilter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGiResourceIdFilterEntityTypeEnum(string(m.EntityType)); !ok && m.EntityType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EntityType: %s. Supported values are: %s.", m.EntityType, strings.Join(GetGiResourceIdFilterEntityTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingGiFleetDiscoveryFilterModeEnum(string(m.Mode)); !ok && m.Mode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Mode: %s. Supported values are: %s.", m.Mode, strings.Join(GetGiFleetDiscoveryFilterModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingFleetDiscoveryOperatorsEnum(string(m.Operator)); !ok && m.Operator != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Operator: %s. Supported values are: %s.", m.Operator, strings.Join(GetFleetDiscoveryOperatorsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m GiResourceIdFilter) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGiResourceIdFilter GiResourceIdFilter
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeGiResourceIdFilter
	}{
		"RESOURCE_ID",
		(MarshalTypeGiResourceIdFilter)(m),
	}

	return json.Marshal(&s)
}

// GiResourceIdFilterEntityTypeEnum Enum with underlying type: string
type GiResourceIdFilterEntityTypeEnum string

// Set of constants representing the allowable values for GiResourceIdFilterEntityTypeEnum
const (
	GiResourceIdFilterEntityTypeDatabasesoftwareimage      GiResourceIdFilterEntityTypeEnum = "DATABASESOFTWAREIMAGE"
	GiResourceIdFilterEntityTypeExadatainfrastructure      GiResourceIdFilterEntityTypeEnum = "EXADATAINFRASTRUCTURE"
	GiResourceIdFilterEntityTypeCloudexadatainfrastructure GiResourceIdFilterEntityTypeEnum = "CLOUDEXADATAINFRASTRUCTURE"
	GiResourceIdFilterEntityTypeVmcluster                  GiResourceIdFilterEntityTypeEnum = "VMCLUSTER"
	GiResourceIdFilterEntityTypeCloudvmcluster             GiResourceIdFilterEntityTypeEnum = "CLOUDVMCLUSTER"
	GiResourceIdFilterEntityTypeFsucollection              GiResourceIdFilterEntityTypeEnum = "FSUCOLLECTION"
)

var mappingGiResourceIdFilterEntityTypeEnum = map[string]GiResourceIdFilterEntityTypeEnum{
	"DATABASESOFTWAREIMAGE":      GiResourceIdFilterEntityTypeDatabasesoftwareimage,
	"EXADATAINFRASTRUCTURE":      GiResourceIdFilterEntityTypeExadatainfrastructure,
	"CLOUDEXADATAINFRASTRUCTURE": GiResourceIdFilterEntityTypeCloudexadatainfrastructure,
	"VMCLUSTER":                  GiResourceIdFilterEntityTypeVmcluster,
	"CLOUDVMCLUSTER":             GiResourceIdFilterEntityTypeCloudvmcluster,
	"FSUCOLLECTION":              GiResourceIdFilterEntityTypeFsucollection,
}

var mappingGiResourceIdFilterEntityTypeEnumLowerCase = map[string]GiResourceIdFilterEntityTypeEnum{
	"databasesoftwareimage":      GiResourceIdFilterEntityTypeDatabasesoftwareimage,
	"exadatainfrastructure":      GiResourceIdFilterEntityTypeExadatainfrastructure,
	"cloudexadatainfrastructure": GiResourceIdFilterEntityTypeCloudexadatainfrastructure,
	"vmcluster":                  GiResourceIdFilterEntityTypeVmcluster,
	"cloudvmcluster":             GiResourceIdFilterEntityTypeCloudvmcluster,
	"fsucollection":              GiResourceIdFilterEntityTypeFsucollection,
}

// GetGiResourceIdFilterEntityTypeEnumValues Enumerates the set of values for GiResourceIdFilterEntityTypeEnum
func GetGiResourceIdFilterEntityTypeEnumValues() []GiResourceIdFilterEntityTypeEnum {
	values := make([]GiResourceIdFilterEntityTypeEnum, 0)
	for _, v := range mappingGiResourceIdFilterEntityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGiResourceIdFilterEntityTypeEnumStringValues Enumerates the set of values in String for GiResourceIdFilterEntityTypeEnum
func GetGiResourceIdFilterEntityTypeEnumStringValues() []string {
	return []string{
		"DATABASESOFTWAREIMAGE",
		"EXADATAINFRASTRUCTURE",
		"CLOUDEXADATAINFRASTRUCTURE",
		"VMCLUSTER",
		"CLOUDVMCLUSTER",
		"FSUCOLLECTION",
	}
}

// GetMappingGiResourceIdFilterEntityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGiResourceIdFilterEntityTypeEnum(val string) (GiResourceIdFilterEntityTypeEnum, bool) {
	enum, ok := mappingGiResourceIdFilterEntityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
