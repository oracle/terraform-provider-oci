// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// ReadinessCheckTargetEntry Details to specify a target to include in the Exadata Fleet Update Readiness Check run.
type ReadinessCheckTargetEntry struct {

	// Resource entity type
	EntityType ReadinessCheckTargetEntryEntityTypeEnum `mandatory:"true" json:"entityType"`

	// Resource identifier OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
	EntityId *string `mandatory:"true" json:"entityId"`
}

func (m ReadinessCheckTargetEntry) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ReadinessCheckTargetEntry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingReadinessCheckTargetEntryEntityTypeEnum(string(m.EntityType)); !ok && m.EntityType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EntityType: %s. Supported values are: %s.", m.EntityType, strings.Join(GetReadinessCheckTargetEntryEntityTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ReadinessCheckTargetEntryEntityTypeEnum Enum with underlying type: string
type ReadinessCheckTargetEntryEntityTypeEnum string

// Set of constants representing the allowable values for ReadinessCheckTargetEntryEntityTypeEnum
const (
	ReadinessCheckTargetEntryEntityTypeDatabase ReadinessCheckTargetEntryEntityTypeEnum = "DATABASE"
)

var mappingReadinessCheckTargetEntryEntityTypeEnum = map[string]ReadinessCheckTargetEntryEntityTypeEnum{
	"DATABASE": ReadinessCheckTargetEntryEntityTypeDatabase,
}

var mappingReadinessCheckTargetEntryEntityTypeEnumLowerCase = map[string]ReadinessCheckTargetEntryEntityTypeEnum{
	"database": ReadinessCheckTargetEntryEntityTypeDatabase,
}

// GetReadinessCheckTargetEntryEntityTypeEnumValues Enumerates the set of values for ReadinessCheckTargetEntryEntityTypeEnum
func GetReadinessCheckTargetEntryEntityTypeEnumValues() []ReadinessCheckTargetEntryEntityTypeEnum {
	values := make([]ReadinessCheckTargetEntryEntityTypeEnum, 0)
	for _, v := range mappingReadinessCheckTargetEntryEntityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetReadinessCheckTargetEntryEntityTypeEnumStringValues Enumerates the set of values in String for ReadinessCheckTargetEntryEntityTypeEnum
func GetReadinessCheckTargetEntryEntityTypeEnumStringValues() []string {
	return []string{
		"DATABASE",
	}
}

// GetMappingReadinessCheckTargetEntryEntityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReadinessCheckTargetEntryEntityTypeEnum(val string) (ReadinessCheckTargetEntryEntityTypeEnum, bool) {
	enum, ok := mappingReadinessCheckTargetEntryEntityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
