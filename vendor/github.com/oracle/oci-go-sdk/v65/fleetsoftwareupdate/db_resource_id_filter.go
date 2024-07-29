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

// DbResourceIdFilter Related resource Ids to include in the discovery.
type DbResourceIdFilter struct {

	// Related resource Ids to include in the discovery.
	// All must match the specified entityType.
	Identifiers []string `mandatory:"true" json:"identifiers"`

	// Type of resource to match in the discovery.
	EntityType DbResourceIdFilterEntityTypeEnum `mandatory:"true" json:"entityType"`

	// INCLUDE or EXCLUDE the filter results in the discovery for DB targets.
	// Supported for 'FSUCOLLECTION' RESOURCE_ID filter only.
	Mode DbFleetDiscoveryFilterModeEnum `mandatory:"false" json:"mode,omitempty"`

	// Type of join for each element in this filter.
	Operator FleetDiscoveryOperatorsEnum `mandatory:"false" json:"operator,omitempty"`
}

// GetMode returns Mode
func (m DbResourceIdFilter) GetMode() DbFleetDiscoveryFilterModeEnum {
	return m.Mode
}

func (m DbResourceIdFilter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbResourceIdFilter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDbResourceIdFilterEntityTypeEnum(string(m.EntityType)); !ok && m.EntityType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EntityType: %s. Supported values are: %s.", m.EntityType, strings.Join(GetDbResourceIdFilterEntityTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDbFleetDiscoveryFilterModeEnum(string(m.Mode)); !ok && m.Mode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Mode: %s. Supported values are: %s.", m.Mode, strings.Join(GetDbFleetDiscoveryFilterModeEnumStringValues(), ",")))
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
func (m DbResourceIdFilter) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDbResourceIdFilter DbResourceIdFilter
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeDbResourceIdFilter
	}{
		"RESOURCE_ID",
		(MarshalTypeDbResourceIdFilter)(m),
	}

	return json.Marshal(&s)
}

// DbResourceIdFilterEntityTypeEnum Enum with underlying type: string
type DbResourceIdFilterEntityTypeEnum string

// Set of constants representing the allowable values for DbResourceIdFilterEntityTypeEnum
const (
	DbResourceIdFilterEntityTypeDatabasesoftwareimage      DbResourceIdFilterEntityTypeEnum = "DATABASESOFTWAREIMAGE"
	DbResourceIdFilterEntityTypeDbhome                     DbResourceIdFilterEntityTypeEnum = "DBHOME"
	DbResourceIdFilterEntityTypeExadatainfrastructure      DbResourceIdFilterEntityTypeEnum = "EXADATAINFRASTRUCTURE"
	DbResourceIdFilterEntityTypeCloudexadatainfrastructure DbResourceIdFilterEntityTypeEnum = "CLOUDEXADATAINFRASTRUCTURE"
	DbResourceIdFilterEntityTypeVmcluster                  DbResourceIdFilterEntityTypeEnum = "VMCLUSTER"
	DbResourceIdFilterEntityTypeCloudvmcluster             DbResourceIdFilterEntityTypeEnum = "CLOUDVMCLUSTER"
	DbResourceIdFilterEntityTypeFsucollection              DbResourceIdFilterEntityTypeEnum = "FSUCOLLECTION"
)

var mappingDbResourceIdFilterEntityTypeEnum = map[string]DbResourceIdFilterEntityTypeEnum{
	"DATABASESOFTWAREIMAGE":      DbResourceIdFilterEntityTypeDatabasesoftwareimage,
	"DBHOME":                     DbResourceIdFilterEntityTypeDbhome,
	"EXADATAINFRASTRUCTURE":      DbResourceIdFilterEntityTypeExadatainfrastructure,
	"CLOUDEXADATAINFRASTRUCTURE": DbResourceIdFilterEntityTypeCloudexadatainfrastructure,
	"VMCLUSTER":                  DbResourceIdFilterEntityTypeVmcluster,
	"CLOUDVMCLUSTER":             DbResourceIdFilterEntityTypeCloudvmcluster,
	"FSUCOLLECTION":              DbResourceIdFilterEntityTypeFsucollection,
}

var mappingDbResourceIdFilterEntityTypeEnumLowerCase = map[string]DbResourceIdFilterEntityTypeEnum{
	"databasesoftwareimage":      DbResourceIdFilterEntityTypeDatabasesoftwareimage,
	"dbhome":                     DbResourceIdFilterEntityTypeDbhome,
	"exadatainfrastructure":      DbResourceIdFilterEntityTypeExadatainfrastructure,
	"cloudexadatainfrastructure": DbResourceIdFilterEntityTypeCloudexadatainfrastructure,
	"vmcluster":                  DbResourceIdFilterEntityTypeVmcluster,
	"cloudvmcluster":             DbResourceIdFilterEntityTypeCloudvmcluster,
	"fsucollection":              DbResourceIdFilterEntityTypeFsucollection,
}

// GetDbResourceIdFilterEntityTypeEnumValues Enumerates the set of values for DbResourceIdFilterEntityTypeEnum
func GetDbResourceIdFilterEntityTypeEnumValues() []DbResourceIdFilterEntityTypeEnum {
	values := make([]DbResourceIdFilterEntityTypeEnum, 0)
	for _, v := range mappingDbResourceIdFilterEntityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDbResourceIdFilterEntityTypeEnumStringValues Enumerates the set of values in String for DbResourceIdFilterEntityTypeEnum
func GetDbResourceIdFilterEntityTypeEnumStringValues() []string {
	return []string{
		"DATABASESOFTWAREIMAGE",
		"DBHOME",
		"EXADATAINFRASTRUCTURE",
		"CLOUDEXADATAINFRASTRUCTURE",
		"VMCLUSTER",
		"CLOUDVMCLUSTER",
		"FSUCOLLECTION",
	}
}

// GetMappingDbResourceIdFilterEntityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbResourceIdFilterEntityTypeEnum(val string) (DbResourceIdFilterEntityTypeEnum, bool) {
	enum, ok := mappingDbResourceIdFilterEntityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
