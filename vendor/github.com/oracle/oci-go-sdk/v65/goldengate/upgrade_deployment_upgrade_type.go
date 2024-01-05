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

// UpgradeDeploymentUpgradeTypeEnum Enum with underlying type: string
type UpgradeDeploymentUpgradeTypeEnum string

// Set of constants representing the allowable values for UpgradeDeploymentUpgradeTypeEnum
const (
	UpgradeDeploymentUpgradeTypeDefault UpgradeDeploymentUpgradeTypeEnum = "DEFAULT"
)

var mappingUpgradeDeploymentUpgradeTypeEnum = map[string]UpgradeDeploymentUpgradeTypeEnum{
	"DEFAULT": UpgradeDeploymentUpgradeTypeDefault,
}

var mappingUpgradeDeploymentUpgradeTypeEnumLowerCase = map[string]UpgradeDeploymentUpgradeTypeEnum{
	"default": UpgradeDeploymentUpgradeTypeDefault,
}

// GetUpgradeDeploymentUpgradeTypeEnumValues Enumerates the set of values for UpgradeDeploymentUpgradeTypeEnum
func GetUpgradeDeploymentUpgradeTypeEnumValues() []UpgradeDeploymentUpgradeTypeEnum {
	values := make([]UpgradeDeploymentUpgradeTypeEnum, 0)
	for _, v := range mappingUpgradeDeploymentUpgradeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpgradeDeploymentUpgradeTypeEnumStringValues Enumerates the set of values in String for UpgradeDeploymentUpgradeTypeEnum
func GetUpgradeDeploymentUpgradeTypeEnumStringValues() []string {
	return []string{
		"DEFAULT",
	}
}

// GetMappingUpgradeDeploymentUpgradeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpgradeDeploymentUpgradeTypeEnum(val string) (UpgradeDeploymentUpgradeTypeEnum, bool) {
	enum, ok := mappingUpgradeDeploymentUpgradeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
