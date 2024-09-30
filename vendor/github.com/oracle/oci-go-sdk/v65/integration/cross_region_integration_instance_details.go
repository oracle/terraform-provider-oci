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

// CrossRegionIntegrationInstanceDetails Details of integration instance created in cross region for disaster recovery.
type CrossRegionIntegrationInstanceDetails struct {

	// Role of the integration instance in the region
	Role CrossRegionIntegrationInstanceDetailsRoleEnum `mandatory:"false" json:"role,omitempty"`

	// Cross region integration instance identifier
	Id *string `mandatory:"false" json:"id"`

	// Cross region where integration instance is created
	Region *string `mandatory:"false" json:"region"`

	// Time when cross region integration instance role was changed
	TimeRoleChanged *common.SDKTime `mandatory:"false" json:"timeRoleChanged"`
}

func (m CrossRegionIntegrationInstanceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CrossRegionIntegrationInstanceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCrossRegionIntegrationInstanceDetailsRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetCrossRegionIntegrationInstanceDetailsRoleEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CrossRegionIntegrationInstanceDetailsRoleEnum Enum with underlying type: string
type CrossRegionIntegrationInstanceDetailsRoleEnum string

// Set of constants representing the allowable values for CrossRegionIntegrationInstanceDetailsRoleEnum
const (
	CrossRegionIntegrationInstanceDetailsRolePrimary   CrossRegionIntegrationInstanceDetailsRoleEnum = "PRIMARY"
	CrossRegionIntegrationInstanceDetailsRoleSecondary CrossRegionIntegrationInstanceDetailsRoleEnum = "SECONDARY"
	CrossRegionIntegrationInstanceDetailsRoleUnknown   CrossRegionIntegrationInstanceDetailsRoleEnum = "UNKNOWN"
)

var mappingCrossRegionIntegrationInstanceDetailsRoleEnum = map[string]CrossRegionIntegrationInstanceDetailsRoleEnum{
	"PRIMARY":   CrossRegionIntegrationInstanceDetailsRolePrimary,
	"SECONDARY": CrossRegionIntegrationInstanceDetailsRoleSecondary,
	"UNKNOWN":   CrossRegionIntegrationInstanceDetailsRoleUnknown,
}

var mappingCrossRegionIntegrationInstanceDetailsRoleEnumLowerCase = map[string]CrossRegionIntegrationInstanceDetailsRoleEnum{
	"primary":   CrossRegionIntegrationInstanceDetailsRolePrimary,
	"secondary": CrossRegionIntegrationInstanceDetailsRoleSecondary,
	"unknown":   CrossRegionIntegrationInstanceDetailsRoleUnknown,
}

// GetCrossRegionIntegrationInstanceDetailsRoleEnumValues Enumerates the set of values for CrossRegionIntegrationInstanceDetailsRoleEnum
func GetCrossRegionIntegrationInstanceDetailsRoleEnumValues() []CrossRegionIntegrationInstanceDetailsRoleEnum {
	values := make([]CrossRegionIntegrationInstanceDetailsRoleEnum, 0)
	for _, v := range mappingCrossRegionIntegrationInstanceDetailsRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetCrossRegionIntegrationInstanceDetailsRoleEnumStringValues Enumerates the set of values in String for CrossRegionIntegrationInstanceDetailsRoleEnum
func GetCrossRegionIntegrationInstanceDetailsRoleEnumStringValues() []string {
	return []string{
		"PRIMARY",
		"SECONDARY",
		"UNKNOWN",
	}
}

// GetMappingCrossRegionIntegrationInstanceDetailsRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCrossRegionIntegrationInstanceDetailsRoleEnum(val string) (CrossRegionIntegrationInstanceDetailsRoleEnum, bool) {
	enum, ok := mappingCrossRegionIntegrationInstanceDetailsRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
