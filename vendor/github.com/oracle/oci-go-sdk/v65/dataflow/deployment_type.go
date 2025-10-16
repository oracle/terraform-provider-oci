// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"strings"
)

// DeploymentTypeEnum Enum with underlying type: string
type DeploymentTypeEnum string

// Set of constants representing the allowable values for DeploymentTypeEnum
const (
	DeploymentTypeDeploy  DeploymentTypeEnum = "DEPLOY"
	DeploymentTypeExecute DeploymentTypeEnum = "EXECUTE"
)

var mappingDeploymentTypeEnum = map[string]DeploymentTypeEnum{
	"DEPLOY":  DeploymentTypeDeploy,
	"EXECUTE": DeploymentTypeExecute,
}

var mappingDeploymentTypeEnumLowerCase = map[string]DeploymentTypeEnum{
	"deploy":  DeploymentTypeDeploy,
	"execute": DeploymentTypeExecute,
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
		"DEPLOY",
		"EXECUTE",
	}
}

// GetMappingDeploymentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDeploymentTypeEnum(val string) (DeploymentTypeEnum, bool) {
	enum, ok := mappingDeploymentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
