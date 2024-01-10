// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fusion Applications Environment Management API
//
// Use the Fusion Applications Environment Management API to manage the environments where your Fusion Applications run. For more information, see the Fusion Applications Environment Management documentation (https://docs.cloud.oracle.com/iaas/Content/fusion-applications/home.htm).
//

package fusionapps

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EnvironmentRole Describes the role of the FA Environment.
type EnvironmentRole struct {

	// The current role of the environment
	CurrentRole EnvironmentRoleCurrentRoleEnum `mandatory:"false" json:"currentRole,omitempty"`

	// Region the standby environment is in
	StandbyEnvironmentRegion *string `mandatory:"false" json:"standbyEnvironmentRegion"`

	// Fusion Environment ID of the standby environment
	StandbyEnvironmentId *string `mandatory:"false" json:"standbyEnvironmentId"`
}

func (m EnvironmentRole) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EnvironmentRole) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingEnvironmentRoleCurrentRoleEnum(string(m.CurrentRole)); !ok && m.CurrentRole != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CurrentRole: %s. Supported values are: %s.", m.CurrentRole, strings.Join(GetEnvironmentRoleCurrentRoleEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EnvironmentRoleCurrentRoleEnum Enum with underlying type: string
type EnvironmentRoleCurrentRoleEnum string

// Set of constants representing the allowable values for EnvironmentRoleCurrentRoleEnum
const (
	EnvironmentRoleCurrentRolePrimary EnvironmentRoleCurrentRoleEnum = "PRIMARY"
	EnvironmentRoleCurrentRoleStandby EnvironmentRoleCurrentRoleEnum = "STANDBY"
)

var mappingEnvironmentRoleCurrentRoleEnum = map[string]EnvironmentRoleCurrentRoleEnum{
	"PRIMARY": EnvironmentRoleCurrentRolePrimary,
	"STANDBY": EnvironmentRoleCurrentRoleStandby,
}

var mappingEnvironmentRoleCurrentRoleEnumLowerCase = map[string]EnvironmentRoleCurrentRoleEnum{
	"primary": EnvironmentRoleCurrentRolePrimary,
	"standby": EnvironmentRoleCurrentRoleStandby,
}

// GetEnvironmentRoleCurrentRoleEnumValues Enumerates the set of values for EnvironmentRoleCurrentRoleEnum
func GetEnvironmentRoleCurrentRoleEnumValues() []EnvironmentRoleCurrentRoleEnum {
	values := make([]EnvironmentRoleCurrentRoleEnum, 0)
	for _, v := range mappingEnvironmentRoleCurrentRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetEnvironmentRoleCurrentRoleEnumStringValues Enumerates the set of values in String for EnvironmentRoleCurrentRoleEnum
func GetEnvironmentRoleCurrentRoleEnumStringValues() []string {
	return []string{
		"PRIMARY",
		"STANDBY",
	}
}

// GetMappingEnvironmentRoleCurrentRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEnvironmentRoleCurrentRoleEnum(val string) (EnvironmentRoleCurrentRoleEnum, bool) {
	enum, ok := mappingEnvironmentRoleCurrentRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
