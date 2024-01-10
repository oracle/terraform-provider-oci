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

// RollbackDeploymentUpgradeTypeEnum Enum with underlying type: string
type RollbackDeploymentUpgradeTypeEnum string

// Set of constants representing the allowable values for RollbackDeploymentUpgradeTypeEnum
const (
	RollbackDeploymentUpgradeTypeDefault RollbackDeploymentUpgradeTypeEnum = "DEFAULT"
)

var mappingRollbackDeploymentUpgradeTypeEnum = map[string]RollbackDeploymentUpgradeTypeEnum{
	"DEFAULT": RollbackDeploymentUpgradeTypeDefault,
}

var mappingRollbackDeploymentUpgradeTypeEnumLowerCase = map[string]RollbackDeploymentUpgradeTypeEnum{
	"default": RollbackDeploymentUpgradeTypeDefault,
}

// GetRollbackDeploymentUpgradeTypeEnumValues Enumerates the set of values for RollbackDeploymentUpgradeTypeEnum
func GetRollbackDeploymentUpgradeTypeEnumValues() []RollbackDeploymentUpgradeTypeEnum {
	values := make([]RollbackDeploymentUpgradeTypeEnum, 0)
	for _, v := range mappingRollbackDeploymentUpgradeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRollbackDeploymentUpgradeTypeEnumStringValues Enumerates the set of values in String for RollbackDeploymentUpgradeTypeEnum
func GetRollbackDeploymentUpgradeTypeEnumStringValues() []string {
	return []string{
		"DEFAULT",
	}
}

// GetMappingRollbackDeploymentUpgradeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRollbackDeploymentUpgradeTypeEnum(val string) (RollbackDeploymentUpgradeTypeEnum, bool) {
	enum, ok := mappingRollbackDeploymentUpgradeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
