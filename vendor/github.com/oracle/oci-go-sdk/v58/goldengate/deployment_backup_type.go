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

// DeploymentBackupTypeEnum Enum with underlying type: string
type DeploymentBackupTypeEnum string

// Set of constants representing the allowable values for DeploymentBackupTypeEnum
const (
	DeploymentBackupTypeIncremental DeploymentBackupTypeEnum = "INCREMENTAL"
	DeploymentBackupTypeFull        DeploymentBackupTypeEnum = "FULL"
)

var mappingDeploymentBackupTypeEnum = map[string]DeploymentBackupTypeEnum{
	"INCREMENTAL": DeploymentBackupTypeIncremental,
	"FULL":        DeploymentBackupTypeFull,
}

// GetDeploymentBackupTypeEnumValues Enumerates the set of values for DeploymentBackupTypeEnum
func GetDeploymentBackupTypeEnumValues() []DeploymentBackupTypeEnum {
	values := make([]DeploymentBackupTypeEnum, 0)
	for _, v := range mappingDeploymentBackupTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDeploymentBackupTypeEnumStringValues Enumerates the set of values in String for DeploymentBackupTypeEnum
func GetDeploymentBackupTypeEnumStringValues() []string {
	return []string{
		"INCREMENTAL",
		"FULL",
	}
}

// GetMappingDeploymentBackupTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDeploymentBackupTypeEnum(val string) (DeploymentBackupTypeEnum, bool) {
	mappingDeploymentBackupTypeEnumIgnoreCase := make(map[string]DeploymentBackupTypeEnum)
	for k, v := range mappingDeploymentBackupTypeEnum {
		mappingDeploymentBackupTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingDeploymentBackupTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
