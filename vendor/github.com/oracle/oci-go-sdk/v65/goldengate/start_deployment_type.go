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

// StartDeploymentTypeEnum Enum with underlying type: string
type StartDeploymentTypeEnum string

// Set of constants representing the allowable values for StartDeploymentTypeEnum
const (
	StartDeploymentTypeDefault StartDeploymentTypeEnum = "DEFAULT"
)

var mappingStartDeploymentTypeEnum = map[string]StartDeploymentTypeEnum{
	"DEFAULT": StartDeploymentTypeDefault,
}

var mappingStartDeploymentTypeEnumLowerCase = map[string]StartDeploymentTypeEnum{
	"default": StartDeploymentTypeDefault,
}

// GetStartDeploymentTypeEnumValues Enumerates the set of values for StartDeploymentTypeEnum
func GetStartDeploymentTypeEnumValues() []StartDeploymentTypeEnum {
	values := make([]StartDeploymentTypeEnum, 0)
	for _, v := range mappingStartDeploymentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetStartDeploymentTypeEnumStringValues Enumerates the set of values in String for StartDeploymentTypeEnum
func GetStartDeploymentTypeEnumStringValues() []string {
	return []string{
		"DEFAULT",
	}
}

// GetMappingStartDeploymentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStartDeploymentTypeEnum(val string) (StartDeploymentTypeEnum, bool) {
	enum, ok := mappingStartDeploymentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
