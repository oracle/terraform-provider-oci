// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// ModelDeploymentEnvironmentConfigurationTypeEnum Enum with underlying type: string
type ModelDeploymentEnvironmentConfigurationTypeEnum string

// Set of constants representing the allowable values for ModelDeploymentEnvironmentConfigurationTypeEnum
const (
	ModelDeploymentEnvironmentConfigurationTypeDefault       ModelDeploymentEnvironmentConfigurationTypeEnum = "DEFAULT"
	ModelDeploymentEnvironmentConfigurationTypeOcirContainer ModelDeploymentEnvironmentConfigurationTypeEnum = "OCIR_CONTAINER"
)

var mappingModelDeploymentEnvironmentConfigurationTypeEnum = map[string]ModelDeploymentEnvironmentConfigurationTypeEnum{
	"DEFAULT":        ModelDeploymentEnvironmentConfigurationTypeDefault,
	"OCIR_CONTAINER": ModelDeploymentEnvironmentConfigurationTypeOcirContainer,
}

var mappingModelDeploymentEnvironmentConfigurationTypeEnumLowerCase = map[string]ModelDeploymentEnvironmentConfigurationTypeEnum{
	"default":        ModelDeploymentEnvironmentConfigurationTypeDefault,
	"ocir_container": ModelDeploymentEnvironmentConfigurationTypeOcirContainer,
}

// GetModelDeploymentEnvironmentConfigurationTypeEnumValues Enumerates the set of values for ModelDeploymentEnvironmentConfigurationTypeEnum
func GetModelDeploymentEnvironmentConfigurationTypeEnumValues() []ModelDeploymentEnvironmentConfigurationTypeEnum {
	values := make([]ModelDeploymentEnvironmentConfigurationTypeEnum, 0)
	for _, v := range mappingModelDeploymentEnvironmentConfigurationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetModelDeploymentEnvironmentConfigurationTypeEnumStringValues Enumerates the set of values in String for ModelDeploymentEnvironmentConfigurationTypeEnum
func GetModelDeploymentEnvironmentConfigurationTypeEnumStringValues() []string {
	return []string{
		"DEFAULT",
		"OCIR_CONTAINER",
	}
}

// GetMappingModelDeploymentEnvironmentConfigurationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModelDeploymentEnvironmentConfigurationTypeEnum(val string) (ModelDeploymentEnvironmentConfigurationTypeEnum, bool) {
	enum, ok := mappingModelDeploymentEnvironmentConfigurationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
