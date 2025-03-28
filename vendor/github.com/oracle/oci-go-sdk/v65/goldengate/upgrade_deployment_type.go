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

// UpgradeDeploymentTypeEnum Enum with underlying type: string
type UpgradeDeploymentTypeEnum string

// Set of constants representing the allowable values for UpgradeDeploymentTypeEnum
const (
	UpgradeDeploymentTypeCurrentRelease  UpgradeDeploymentTypeEnum = "CURRENT_RELEASE"
	UpgradeDeploymentTypeSpecificRelease UpgradeDeploymentTypeEnum = "SPECIFIC_RELEASE"
)

var mappingUpgradeDeploymentTypeEnum = map[string]UpgradeDeploymentTypeEnum{
	"CURRENT_RELEASE":  UpgradeDeploymentTypeCurrentRelease,
	"SPECIFIC_RELEASE": UpgradeDeploymentTypeSpecificRelease,
}

var mappingUpgradeDeploymentTypeEnumLowerCase = map[string]UpgradeDeploymentTypeEnum{
	"current_release":  UpgradeDeploymentTypeCurrentRelease,
	"specific_release": UpgradeDeploymentTypeSpecificRelease,
}

// GetUpgradeDeploymentTypeEnumValues Enumerates the set of values for UpgradeDeploymentTypeEnum
func GetUpgradeDeploymentTypeEnumValues() []UpgradeDeploymentTypeEnum {
	values := make([]UpgradeDeploymentTypeEnum, 0)
	for _, v := range mappingUpgradeDeploymentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpgradeDeploymentTypeEnumStringValues Enumerates the set of values in String for UpgradeDeploymentTypeEnum
func GetUpgradeDeploymentTypeEnumStringValues() []string {
	return []string{
		"CURRENT_RELEASE",
		"SPECIFIC_RELEASE",
	}
}

// GetMappingUpgradeDeploymentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpgradeDeploymentTypeEnum(val string) (UpgradeDeploymentTypeEnum, bool) {
	enum, ok := mappingUpgradeDeploymentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
