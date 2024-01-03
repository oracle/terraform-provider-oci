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

// CancelSnoozeDeploymentUpgradeTypeEnum Enum with underlying type: string
type CancelSnoozeDeploymentUpgradeTypeEnum string

// Set of constants representing the allowable values for CancelSnoozeDeploymentUpgradeTypeEnum
const (
	CancelSnoozeDeploymentUpgradeTypeDefault CancelSnoozeDeploymentUpgradeTypeEnum = "DEFAULT"
)

var mappingCancelSnoozeDeploymentUpgradeTypeEnum = map[string]CancelSnoozeDeploymentUpgradeTypeEnum{
	"DEFAULT": CancelSnoozeDeploymentUpgradeTypeDefault,
}

var mappingCancelSnoozeDeploymentUpgradeTypeEnumLowerCase = map[string]CancelSnoozeDeploymentUpgradeTypeEnum{
	"default": CancelSnoozeDeploymentUpgradeTypeDefault,
}

// GetCancelSnoozeDeploymentUpgradeTypeEnumValues Enumerates the set of values for CancelSnoozeDeploymentUpgradeTypeEnum
func GetCancelSnoozeDeploymentUpgradeTypeEnumValues() []CancelSnoozeDeploymentUpgradeTypeEnum {
	values := make([]CancelSnoozeDeploymentUpgradeTypeEnum, 0)
	for _, v := range mappingCancelSnoozeDeploymentUpgradeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCancelSnoozeDeploymentUpgradeTypeEnumStringValues Enumerates the set of values in String for CancelSnoozeDeploymentUpgradeTypeEnum
func GetCancelSnoozeDeploymentUpgradeTypeEnumStringValues() []string {
	return []string{
		"DEFAULT",
	}
}

// GetMappingCancelSnoozeDeploymentUpgradeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCancelSnoozeDeploymentUpgradeTypeEnum(val string) (CancelSnoozeDeploymentUpgradeTypeEnum, bool) {
	enum, ok := mappingCancelSnoozeDeploymentUpgradeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
