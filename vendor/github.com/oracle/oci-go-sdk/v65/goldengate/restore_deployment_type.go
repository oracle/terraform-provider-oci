// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// RestoreDeploymentTypeEnum Enum with underlying type: string
type RestoreDeploymentTypeEnum string

// Set of constants representing the allowable values for RestoreDeploymentTypeEnum
const (
	RestoreDeploymentTypeDefault RestoreDeploymentTypeEnum = "DEFAULT"
)

var mappingRestoreDeploymentTypeEnum = map[string]RestoreDeploymentTypeEnum{
	"DEFAULT": RestoreDeploymentTypeDefault,
}

var mappingRestoreDeploymentTypeEnumLowerCase = map[string]RestoreDeploymentTypeEnum{
	"default": RestoreDeploymentTypeDefault,
}

// GetRestoreDeploymentTypeEnumValues Enumerates the set of values for RestoreDeploymentTypeEnum
func GetRestoreDeploymentTypeEnumValues() []RestoreDeploymentTypeEnum {
	values := make([]RestoreDeploymentTypeEnum, 0)
	for _, v := range mappingRestoreDeploymentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRestoreDeploymentTypeEnumStringValues Enumerates the set of values in String for RestoreDeploymentTypeEnum
func GetRestoreDeploymentTypeEnumStringValues() []string {
	return []string{
		"DEFAULT",
	}
}

// GetMappingRestoreDeploymentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRestoreDeploymentTypeEnum(val string) (RestoreDeploymentTypeEnum, bool) {
	enum, ok := mappingRestoreDeploymentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
