// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MigrationBandwidth Defines the bandwidth used by migration
type MigrationBandwidth struct {

	// Defines how the migration bandwidth is assigned
	AssignmentMethod MigrationBandwidthAssignmentMethodEnum `mandatory:"false" json:"assignmentMethod,omitempty"`
}

func (m MigrationBandwidth) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MigrationBandwidth) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingMigrationBandwidthAssignmentMethodEnum(string(m.AssignmentMethod)); !ok && m.AssignmentMethod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AssignmentMethod: %s. Supported values are: %s.", m.AssignmentMethod, strings.Join(GetMigrationBandwidthAssignmentMethodEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MigrationBandwidthAssignmentMethodEnum Enum with underlying type: string
type MigrationBandwidthAssignmentMethodEnum string

// Set of constants representing the allowable values for MigrationBandwidthAssignmentMethodEnum
const (
	MigrationBandwidthAssignmentMethodAuto              MigrationBandwidthAssignmentMethodEnum = "AUTO"
	MigrationBandwidthAssignmentMethodCustom            MigrationBandwidthAssignmentMethodEnum = "CUSTOM"
	MigrationBandwidthAssignmentMethodHypervisorDefault MigrationBandwidthAssignmentMethodEnum = "HYPERVISOR_DEFAULT"
)

var mappingMigrationBandwidthAssignmentMethodEnum = map[string]MigrationBandwidthAssignmentMethodEnum{
	"AUTO":               MigrationBandwidthAssignmentMethodAuto,
	"CUSTOM":             MigrationBandwidthAssignmentMethodCustom,
	"HYPERVISOR_DEFAULT": MigrationBandwidthAssignmentMethodHypervisorDefault,
}

var mappingMigrationBandwidthAssignmentMethodEnumLowerCase = map[string]MigrationBandwidthAssignmentMethodEnum{
	"auto":               MigrationBandwidthAssignmentMethodAuto,
	"custom":             MigrationBandwidthAssignmentMethodCustom,
	"hypervisor_default": MigrationBandwidthAssignmentMethodHypervisorDefault,
}

// GetMigrationBandwidthAssignmentMethodEnumValues Enumerates the set of values for MigrationBandwidthAssignmentMethodEnum
func GetMigrationBandwidthAssignmentMethodEnumValues() []MigrationBandwidthAssignmentMethodEnum {
	values := make([]MigrationBandwidthAssignmentMethodEnum, 0)
	for _, v := range mappingMigrationBandwidthAssignmentMethodEnum {
		values = append(values, v)
	}
	return values
}

// GetMigrationBandwidthAssignmentMethodEnumStringValues Enumerates the set of values in String for MigrationBandwidthAssignmentMethodEnum
func GetMigrationBandwidthAssignmentMethodEnumStringValues() []string {
	return []string{
		"AUTO",
		"CUSTOM",
		"HYPERVISOR_DEFAULT",
	}
}

// GetMappingMigrationBandwidthAssignmentMethodEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMigrationBandwidthAssignmentMethodEnum(val string) (MigrationBandwidthAssignmentMethodEnum, bool) {
	enum, ok := mappingMigrationBandwidthAssignmentMethodEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
