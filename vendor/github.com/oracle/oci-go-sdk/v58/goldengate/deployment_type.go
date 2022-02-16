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

// DeploymentTypeEnum Enum with underlying type: string
type DeploymentTypeEnum string

// Set of constants representing the allowable values for DeploymentTypeEnum
const (
	DeploymentTypeOgg DeploymentTypeEnum = "OGG"
)

var mappingDeploymentTypeEnum = map[string]DeploymentTypeEnum{
	"OGG": DeploymentTypeOgg,
}

// GetDeploymentTypeEnumValues Enumerates the set of values for DeploymentTypeEnum
func GetDeploymentTypeEnumValues() []DeploymentTypeEnum {
	values := make([]DeploymentTypeEnum, 0)
	for _, v := range mappingDeploymentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDeploymentTypeEnumStringValues Enumerates the set of values in String for DeploymentTypeEnum
func GetDeploymentTypeEnumStringValues() []string {
	return []string{
		"OGG",
	}
}

// GetMappingDeploymentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDeploymentTypeEnum(val string) (DeploymentTypeEnum, bool) {
	mappingDeploymentTypeEnumIgnoreCase := make(map[string]DeploymentTypeEnum)
	for k, v := range mappingDeploymentTypeEnum {
		mappingDeploymentTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingDeploymentTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
