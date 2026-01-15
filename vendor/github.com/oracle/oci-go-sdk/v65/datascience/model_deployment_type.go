// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// ModelDeploymentTypeEnum Enum with underlying type: string
type ModelDeploymentTypeEnum string

// Set of constants representing the allowable values for ModelDeploymentTypeEnum
const (
	ModelDeploymentTypeSingleModel ModelDeploymentTypeEnum = "SINGLE_MODEL"
	ModelDeploymentTypeModelGroup  ModelDeploymentTypeEnum = "MODEL_GROUP"
)

var mappingModelDeploymentTypeEnum = map[string]ModelDeploymentTypeEnum{
	"SINGLE_MODEL": ModelDeploymentTypeSingleModel,
	"MODEL_GROUP":  ModelDeploymentTypeModelGroup,
}

var mappingModelDeploymentTypeEnumLowerCase = map[string]ModelDeploymentTypeEnum{
	"single_model": ModelDeploymentTypeSingleModel,
	"model_group":  ModelDeploymentTypeModelGroup,
}

// GetModelDeploymentTypeEnumValues Enumerates the set of values for ModelDeploymentTypeEnum
func GetModelDeploymentTypeEnumValues() []ModelDeploymentTypeEnum {
	values := make([]ModelDeploymentTypeEnum, 0)
	for _, v := range mappingModelDeploymentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetModelDeploymentTypeEnumStringValues Enumerates the set of values in String for ModelDeploymentTypeEnum
func GetModelDeploymentTypeEnumStringValues() []string {
	return []string{
		"SINGLE_MODEL",
		"MODEL_GROUP",
	}
}

// GetMappingModelDeploymentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModelDeploymentTypeEnum(val string) (ModelDeploymentTypeEnum, bool) {
	enum, ok := mappingModelDeploymentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
