// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// DeploymentUpgradeTypeEnum Enum with underlying type: string
type DeploymentUpgradeTypeEnum string

// Set of constants representing the allowable values for DeploymentUpgradeTypeEnum
const (
	DeploymentUpgradeTypeManual    DeploymentUpgradeTypeEnum = "MANUAL"
	DeploymentUpgradeTypeAutomatic DeploymentUpgradeTypeEnum = "AUTOMATIC"
)

var mappingDeploymentUpgradeTypeEnum = map[string]DeploymentUpgradeTypeEnum{
	"MANUAL":    DeploymentUpgradeTypeManual,
	"AUTOMATIC": DeploymentUpgradeTypeAutomatic,
}

// GetDeploymentUpgradeTypeEnumValues Enumerates the set of values for DeploymentUpgradeTypeEnum
func GetDeploymentUpgradeTypeEnumValues() []DeploymentUpgradeTypeEnum {
	values := make([]DeploymentUpgradeTypeEnum, 0)
	for _, v := range mappingDeploymentUpgradeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDeploymentUpgradeTypeEnumStringValues Enumerates the set of values in String for DeploymentUpgradeTypeEnum
func GetDeploymentUpgradeTypeEnumStringValues() []string {
	return []string{
		"MANUAL",
		"AUTOMATIC",
	}
}

// GetMappingDeploymentUpgradeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDeploymentUpgradeTypeEnum(val string) (DeploymentUpgradeTypeEnum, bool) {
	mappingDeploymentUpgradeTypeEnumIgnoreCase := make(map[string]DeploymentUpgradeTypeEnum)
	for k, v := range mappingDeploymentUpgradeTypeEnum {
		mappingDeploymentUpgradeTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingDeploymentUpgradeTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
