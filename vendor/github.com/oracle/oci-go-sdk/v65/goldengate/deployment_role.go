// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"strings"
)

// DeploymentRoleEnum Enum with underlying type: string
type DeploymentRoleEnum string

// Set of constants representing the allowable values for DeploymentRoleEnum
const (
	DeploymentRolePrimary DeploymentRoleEnum = "PRIMARY"
	DeploymentRoleStandby DeploymentRoleEnum = "STANDBY"
)

var mappingDeploymentRoleEnum = map[string]DeploymentRoleEnum{
	"PRIMARY": DeploymentRolePrimary,
	"STANDBY": DeploymentRoleStandby,
}

var mappingDeploymentRoleEnumLowerCase = map[string]DeploymentRoleEnum{
	"primary": DeploymentRolePrimary,
	"standby": DeploymentRoleStandby,
}

// GetDeploymentRoleEnumValues Enumerates the set of values for DeploymentRoleEnum
func GetDeploymentRoleEnumValues() []DeploymentRoleEnum {
	values := make([]DeploymentRoleEnum, 0)
	for _, v := range mappingDeploymentRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetDeploymentRoleEnumStringValues Enumerates the set of values in String for DeploymentRoleEnum
func GetDeploymentRoleEnumStringValues() []string {
	return []string{
		"PRIMARY",
		"STANDBY",
	}
}

// GetMappingDeploymentRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDeploymentRoleEnum(val string) (DeploymentRoleEnum, bool) {
	enum, ok := mappingDeploymentRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
