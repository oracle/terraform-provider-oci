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

// CancelDeploymentUpgradeTypeEnum Enum with underlying type: string
type CancelDeploymentUpgradeTypeEnum string

// Set of constants representing the allowable values for CancelDeploymentUpgradeTypeEnum
const (
	CancelDeploymentUpgradeTypeDefault CancelDeploymentUpgradeTypeEnum = "DEFAULT"
)

var mappingCancelDeploymentUpgradeTypeEnum = map[string]CancelDeploymentUpgradeTypeEnum{
	"DEFAULT": CancelDeploymentUpgradeTypeDefault,
}

var mappingCancelDeploymentUpgradeTypeEnumLowerCase = map[string]CancelDeploymentUpgradeTypeEnum{
	"default": CancelDeploymentUpgradeTypeDefault,
}

// GetCancelDeploymentUpgradeTypeEnumValues Enumerates the set of values for CancelDeploymentUpgradeTypeEnum
func GetCancelDeploymentUpgradeTypeEnumValues() []CancelDeploymentUpgradeTypeEnum {
	values := make([]CancelDeploymentUpgradeTypeEnum, 0)
	for _, v := range mappingCancelDeploymentUpgradeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCancelDeploymentUpgradeTypeEnumStringValues Enumerates the set of values in String for CancelDeploymentUpgradeTypeEnum
func GetCancelDeploymentUpgradeTypeEnumStringValues() []string {
	return []string{
		"DEFAULT",
	}
}

// GetMappingCancelDeploymentUpgradeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCancelDeploymentUpgradeTypeEnum(val string) (CancelDeploymentUpgradeTypeEnum, bool) {
	enum, ok := mappingCancelDeploymentUpgradeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
