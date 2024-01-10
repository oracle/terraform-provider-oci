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

// SnoozeDeploymentUpgradeTypeEnum Enum with underlying type: string
type SnoozeDeploymentUpgradeTypeEnum string

// Set of constants representing the allowable values for SnoozeDeploymentUpgradeTypeEnum
const (
	SnoozeDeploymentUpgradeTypeDefault SnoozeDeploymentUpgradeTypeEnum = "DEFAULT"
)

var mappingSnoozeDeploymentUpgradeTypeEnum = map[string]SnoozeDeploymentUpgradeTypeEnum{
	"DEFAULT": SnoozeDeploymentUpgradeTypeDefault,
}

var mappingSnoozeDeploymentUpgradeTypeEnumLowerCase = map[string]SnoozeDeploymentUpgradeTypeEnum{
	"default": SnoozeDeploymentUpgradeTypeDefault,
}

// GetSnoozeDeploymentUpgradeTypeEnumValues Enumerates the set of values for SnoozeDeploymentUpgradeTypeEnum
func GetSnoozeDeploymentUpgradeTypeEnumValues() []SnoozeDeploymentUpgradeTypeEnum {
	values := make([]SnoozeDeploymentUpgradeTypeEnum, 0)
	for _, v := range mappingSnoozeDeploymentUpgradeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSnoozeDeploymentUpgradeTypeEnumStringValues Enumerates the set of values in String for SnoozeDeploymentUpgradeTypeEnum
func GetSnoozeDeploymentUpgradeTypeEnumStringValues() []string {
	return []string{
		"DEFAULT",
	}
}

// GetMappingSnoozeDeploymentUpgradeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSnoozeDeploymentUpgradeTypeEnum(val string) (SnoozeDeploymentUpgradeTypeEnum, bool) {
	enum, ok := mappingSnoozeDeploymentUpgradeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
