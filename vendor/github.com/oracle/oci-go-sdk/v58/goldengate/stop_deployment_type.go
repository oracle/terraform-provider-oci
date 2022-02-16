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

// StopDeploymentTypeEnum Enum with underlying type: string
type StopDeploymentTypeEnum string

// Set of constants representing the allowable values for StopDeploymentTypeEnum
const (
	StopDeploymentTypeDefault StopDeploymentTypeEnum = "DEFAULT"
)

var mappingStopDeploymentTypeEnum = map[string]StopDeploymentTypeEnum{
	"DEFAULT": StopDeploymentTypeDefault,
}

// GetStopDeploymentTypeEnumValues Enumerates the set of values for StopDeploymentTypeEnum
func GetStopDeploymentTypeEnumValues() []StopDeploymentTypeEnum {
	values := make([]StopDeploymentTypeEnum, 0)
	for _, v := range mappingStopDeploymentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetStopDeploymentTypeEnumStringValues Enumerates the set of values in String for StopDeploymentTypeEnum
func GetStopDeploymentTypeEnumStringValues() []string {
	return []string{
		"DEFAULT",
	}
}

// GetMappingStopDeploymentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStopDeploymentTypeEnum(val string) (StopDeploymentTypeEnum, bool) {
	mappingStopDeploymentTypeEnumIgnoreCase := make(map[string]StopDeploymentTypeEnum)
	for k, v := range mappingStopDeploymentTypeEnum {
		mappingStopDeploymentTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingStopDeploymentTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
