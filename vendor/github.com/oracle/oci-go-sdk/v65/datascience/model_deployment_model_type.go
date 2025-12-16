// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"strings"
)

// ModelDeploymentModelTypeEnum Enum with underlying type: string
type ModelDeploymentModelTypeEnum string

// Set of constants representing the allowable values for ModelDeploymentModelTypeEnum
const (
	ModelDeploymentModelTypeManagedModel ModelDeploymentModelTypeEnum = "MANAGED_MODEL"
)

var mappingModelDeploymentModelTypeEnum = map[string]ModelDeploymentModelTypeEnum{
	"MANAGED_MODEL": ModelDeploymentModelTypeManagedModel,
}

var mappingModelDeploymentModelTypeEnumLowerCase = map[string]ModelDeploymentModelTypeEnum{
	"managed_model": ModelDeploymentModelTypeManagedModel,
}

// GetModelDeploymentModelTypeEnumValues Enumerates the set of values for ModelDeploymentModelTypeEnum
func GetModelDeploymentModelTypeEnumValues() []ModelDeploymentModelTypeEnum {
	values := make([]ModelDeploymentModelTypeEnum, 0)
	for _, v := range mappingModelDeploymentModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetModelDeploymentModelTypeEnumStringValues Enumerates the set of values in String for ModelDeploymentModelTypeEnum
func GetModelDeploymentModelTypeEnumStringValues() []string {
	return []string{
		"MANAGED_MODEL",
	}
}

// GetMappingModelDeploymentModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModelDeploymentModelTypeEnum(val string) (ModelDeploymentModelTypeEnum, bool) {
	enum, ok := mappingModelDeploymentModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
