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

// RescheduleDeploymentUpgradeTypeEnum Enum with underlying type: string
type RescheduleDeploymentUpgradeTypeEnum string

// Set of constants representing the allowable values for RescheduleDeploymentUpgradeTypeEnum
const (
	RescheduleDeploymentUpgradeTypeRescheduleToDate RescheduleDeploymentUpgradeTypeEnum = "RESCHEDULE_TO_DATE"
)

var mappingRescheduleDeploymentUpgradeTypeEnum = map[string]RescheduleDeploymentUpgradeTypeEnum{
	"RESCHEDULE_TO_DATE": RescheduleDeploymentUpgradeTypeRescheduleToDate,
}

var mappingRescheduleDeploymentUpgradeTypeEnumLowerCase = map[string]RescheduleDeploymentUpgradeTypeEnum{
	"reschedule_to_date": RescheduleDeploymentUpgradeTypeRescheduleToDate,
}

// GetRescheduleDeploymentUpgradeTypeEnumValues Enumerates the set of values for RescheduleDeploymentUpgradeTypeEnum
func GetRescheduleDeploymentUpgradeTypeEnumValues() []RescheduleDeploymentUpgradeTypeEnum {
	values := make([]RescheduleDeploymentUpgradeTypeEnum, 0)
	for _, v := range mappingRescheduleDeploymentUpgradeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRescheduleDeploymentUpgradeTypeEnumStringValues Enumerates the set of values in String for RescheduleDeploymentUpgradeTypeEnum
func GetRescheduleDeploymentUpgradeTypeEnumStringValues() []string {
	return []string{
		"RESCHEDULE_TO_DATE",
	}
}

// GetMappingRescheduleDeploymentUpgradeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRescheduleDeploymentUpgradeTypeEnum(val string) (RescheduleDeploymentUpgradeTypeEnum, bool) {
	enum, ok := mappingRescheduleDeploymentUpgradeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
