// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Integration API
//
// Oracle Integration API.
//

package integration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DisasterRecoveryDetails Disaster recovery details for the integration instance created in the region.
type DisasterRecoveryDetails struct {

	// Role of the integration instance in the region
	Role DisasterRecoveryDetailsRoleEnum `mandatory:"true" json:"role"`

	// Region specific instance url for the integration instance in the region
	RegionalInstanceUrl *string `mandatory:"true" json:"regionalInstanceUrl"`

	CrossRegionIntegrationInstanceDetails *CrossRegionIntegrationInstanceDetails `mandatory:"true" json:"crossRegionIntegrationInstanceDetails"`
}

func (m DisasterRecoveryDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DisasterRecoveryDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDisasterRecoveryDetailsRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetDisasterRecoveryDetailsRoleEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DisasterRecoveryDetailsRoleEnum Enum with underlying type: string
type DisasterRecoveryDetailsRoleEnum string

// Set of constants representing the allowable values for DisasterRecoveryDetailsRoleEnum
const (
	DisasterRecoveryDetailsRolePrimary   DisasterRecoveryDetailsRoleEnum = "PRIMARY"
	DisasterRecoveryDetailsRoleSecondary DisasterRecoveryDetailsRoleEnum = "SECONDARY"
	DisasterRecoveryDetailsRoleUnknown   DisasterRecoveryDetailsRoleEnum = "UNKNOWN"
)

var mappingDisasterRecoveryDetailsRoleEnum = map[string]DisasterRecoveryDetailsRoleEnum{
	"PRIMARY":   DisasterRecoveryDetailsRolePrimary,
	"SECONDARY": DisasterRecoveryDetailsRoleSecondary,
	"UNKNOWN":   DisasterRecoveryDetailsRoleUnknown,
}

var mappingDisasterRecoveryDetailsRoleEnumLowerCase = map[string]DisasterRecoveryDetailsRoleEnum{
	"primary":   DisasterRecoveryDetailsRolePrimary,
	"secondary": DisasterRecoveryDetailsRoleSecondary,
	"unknown":   DisasterRecoveryDetailsRoleUnknown,
}

// GetDisasterRecoveryDetailsRoleEnumValues Enumerates the set of values for DisasterRecoveryDetailsRoleEnum
func GetDisasterRecoveryDetailsRoleEnumValues() []DisasterRecoveryDetailsRoleEnum {
	values := make([]DisasterRecoveryDetailsRoleEnum, 0)
	for _, v := range mappingDisasterRecoveryDetailsRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetDisasterRecoveryDetailsRoleEnumStringValues Enumerates the set of values in String for DisasterRecoveryDetailsRoleEnum
func GetDisasterRecoveryDetailsRoleEnumStringValues() []string {
	return []string{
		"PRIMARY",
		"SECONDARY",
		"UNKNOWN",
	}
}

// GetMappingDisasterRecoveryDetailsRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDisasterRecoveryDetailsRoleEnum(val string) (DisasterRecoveryDetailsRoleEnum, bool) {
	enum, ok := mappingDisasterRecoveryDetailsRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
