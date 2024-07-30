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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TargetEntry Details to specify a target to add or remove from a Exadata Fleet Update Collection.
type TargetEntry struct {

	// Resource entity type
	EntityType TargetEntryEntityTypeEnum `mandatory:"true" json:"entityType"`

	// Resource identifier OCID
	Identifier *string `mandatory:"true" json:"identifier"`
}

func (m TargetEntry) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TargetEntry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTargetEntryEntityTypeEnum(string(m.EntityType)); !ok && m.EntityType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EntityType: %s. Supported values are: %s.", m.EntityType, strings.Join(GetTargetEntryEntityTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TargetEntryEntityTypeEnum Enum with underlying type: string
type TargetEntryEntityTypeEnum string

// Set of constants representing the allowable values for TargetEntryEntityTypeEnum
const (
	TargetEntryEntityTypeDatabase       TargetEntryEntityTypeEnum = "DATABASE"
	TargetEntryEntityTypeVmcluster      TargetEntryEntityTypeEnum = "VMCLUSTER"
	TargetEntryEntityTypeCloudvmcluster TargetEntryEntityTypeEnum = "CLOUDVMCLUSTER"
)

var mappingTargetEntryEntityTypeEnum = map[string]TargetEntryEntityTypeEnum{
	"DATABASE":       TargetEntryEntityTypeDatabase,
	"VMCLUSTER":      TargetEntryEntityTypeVmcluster,
	"CLOUDVMCLUSTER": TargetEntryEntityTypeCloudvmcluster,
}

var mappingTargetEntryEntityTypeEnumLowerCase = map[string]TargetEntryEntityTypeEnum{
	"database":       TargetEntryEntityTypeDatabase,
	"vmcluster":      TargetEntryEntityTypeVmcluster,
	"cloudvmcluster": TargetEntryEntityTypeCloudvmcluster,
}

// GetTargetEntryEntityTypeEnumValues Enumerates the set of values for TargetEntryEntityTypeEnum
func GetTargetEntryEntityTypeEnumValues() []TargetEntryEntityTypeEnum {
	values := make([]TargetEntryEntityTypeEnum, 0)
	for _, v := range mappingTargetEntryEntityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTargetEntryEntityTypeEnumStringValues Enumerates the set of values in String for TargetEntryEntityTypeEnum
func GetTargetEntryEntityTypeEnumStringValues() []string {
	return []string{
		"DATABASE",
		"VMCLUSTER",
		"CLOUDVMCLUSTER",
	}
}

// GetMappingTargetEntryEntityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTargetEntryEntityTypeEnum(val string) (TargetEntryEntityTypeEnum, bool) {
	enum, ok := mappingTargetEntryEntityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
