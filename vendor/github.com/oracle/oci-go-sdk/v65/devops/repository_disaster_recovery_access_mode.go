// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"strings"
)

// RepositoryDisasterRecoveryAccessModeEnum Enum with underlying type: string
type RepositoryDisasterRecoveryAccessModeEnum string

// Set of constants representing the allowable values for RepositoryDisasterRecoveryAccessModeEnum
const (
	RepositoryDisasterRecoveryAccessModeActive  RepositoryDisasterRecoveryAccessModeEnum = "ACTIVE"
	RepositoryDisasterRecoveryAccessModeStandby RepositoryDisasterRecoveryAccessModeEnum = "STANDBY"
)

var mappingRepositoryDisasterRecoveryAccessModeEnum = map[string]RepositoryDisasterRecoveryAccessModeEnum{
	"ACTIVE":  RepositoryDisasterRecoveryAccessModeActive,
	"STANDBY": RepositoryDisasterRecoveryAccessModeStandby,
}

var mappingRepositoryDisasterRecoveryAccessModeEnumLowerCase = map[string]RepositoryDisasterRecoveryAccessModeEnum{
	"active":  RepositoryDisasterRecoveryAccessModeActive,
	"standby": RepositoryDisasterRecoveryAccessModeStandby,
}

// GetRepositoryDisasterRecoveryAccessModeEnumValues Enumerates the set of values for RepositoryDisasterRecoveryAccessModeEnum
func GetRepositoryDisasterRecoveryAccessModeEnumValues() []RepositoryDisasterRecoveryAccessModeEnum {
	values := make([]RepositoryDisasterRecoveryAccessModeEnum, 0)
	for _, v := range mappingRepositoryDisasterRecoveryAccessModeEnum {
		values = append(values, v)
	}
	return values
}

// GetRepositoryDisasterRecoveryAccessModeEnumStringValues Enumerates the set of values in String for RepositoryDisasterRecoveryAccessModeEnum
func GetRepositoryDisasterRecoveryAccessModeEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"STANDBY",
	}
}

// GetMappingRepositoryDisasterRecoveryAccessModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRepositoryDisasterRecoveryAccessModeEnum(val string) (RepositoryDisasterRecoveryAccessModeEnum, bool) {
	enum, ok := mappingRepositoryDisasterRecoveryAccessModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
