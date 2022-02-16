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

// CancelDeploymentBackupTypeEnum Enum with underlying type: string
type CancelDeploymentBackupTypeEnum string

// Set of constants representing the allowable values for CancelDeploymentBackupTypeEnum
const (
	CancelDeploymentBackupTypeDefault CancelDeploymentBackupTypeEnum = "DEFAULT"
)

var mappingCancelDeploymentBackupTypeEnum = map[string]CancelDeploymentBackupTypeEnum{
	"DEFAULT": CancelDeploymentBackupTypeDefault,
}

// GetCancelDeploymentBackupTypeEnumValues Enumerates the set of values for CancelDeploymentBackupTypeEnum
func GetCancelDeploymentBackupTypeEnumValues() []CancelDeploymentBackupTypeEnum {
	values := make([]CancelDeploymentBackupTypeEnum, 0)
	for _, v := range mappingCancelDeploymentBackupTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCancelDeploymentBackupTypeEnumStringValues Enumerates the set of values in String for CancelDeploymentBackupTypeEnum
func GetCancelDeploymentBackupTypeEnumStringValues() []string {
	return []string{
		"DEFAULT",
	}
}

// GetMappingCancelDeploymentBackupTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCancelDeploymentBackupTypeEnum(val string) (CancelDeploymentBackupTypeEnum, bool) {
	mappingCancelDeploymentBackupTypeEnumIgnoreCase := make(map[string]CancelDeploymentBackupTypeEnum)
	for k, v := range mappingCancelDeploymentBackupTypeEnum {
		mappingCancelDeploymentBackupTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingCancelDeploymentBackupTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
