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

// ModelDeploymentShapeSeriesEnum Enum with underlying type: string
type ModelDeploymentShapeSeriesEnum string

// Set of constants representing the allowable values for ModelDeploymentShapeSeriesEnum
const (
	ModelDeploymentShapeSeriesAmdRome      ModelDeploymentShapeSeriesEnum = "AMD_ROME"
	ModelDeploymentShapeSeriesIntelSkylake ModelDeploymentShapeSeriesEnum = "INTEL_SKYLAKE"
	ModelDeploymentShapeSeriesNvidiaGpu    ModelDeploymentShapeSeriesEnum = "NVIDIA_GPU"
	ModelDeploymentShapeSeriesGeneric      ModelDeploymentShapeSeriesEnum = "GENERIC"
	ModelDeploymentShapeSeriesLegacy       ModelDeploymentShapeSeriesEnum = "LEGACY"
	ModelDeploymentShapeSeriesArm          ModelDeploymentShapeSeriesEnum = "ARM"
)

var mappingModelDeploymentShapeSeriesEnum = map[string]ModelDeploymentShapeSeriesEnum{
	"AMD_ROME":      ModelDeploymentShapeSeriesAmdRome,
	"INTEL_SKYLAKE": ModelDeploymentShapeSeriesIntelSkylake,
	"NVIDIA_GPU":    ModelDeploymentShapeSeriesNvidiaGpu,
	"GENERIC":       ModelDeploymentShapeSeriesGeneric,
	"LEGACY":        ModelDeploymentShapeSeriesLegacy,
	"ARM":           ModelDeploymentShapeSeriesArm,
}

var mappingModelDeploymentShapeSeriesEnumLowerCase = map[string]ModelDeploymentShapeSeriesEnum{
	"amd_rome":      ModelDeploymentShapeSeriesAmdRome,
	"intel_skylake": ModelDeploymentShapeSeriesIntelSkylake,
	"nvidia_gpu":    ModelDeploymentShapeSeriesNvidiaGpu,
	"generic":       ModelDeploymentShapeSeriesGeneric,
	"legacy":        ModelDeploymentShapeSeriesLegacy,
	"arm":           ModelDeploymentShapeSeriesArm,
}

// GetModelDeploymentShapeSeriesEnumValues Enumerates the set of values for ModelDeploymentShapeSeriesEnum
func GetModelDeploymentShapeSeriesEnumValues() []ModelDeploymentShapeSeriesEnum {
	values := make([]ModelDeploymentShapeSeriesEnum, 0)
	for _, v := range mappingModelDeploymentShapeSeriesEnum {
		values = append(values, v)
	}
	return values
}

// GetModelDeploymentShapeSeriesEnumStringValues Enumerates the set of values in String for ModelDeploymentShapeSeriesEnum
func GetModelDeploymentShapeSeriesEnumStringValues() []string {
	return []string{
		"AMD_ROME",
		"INTEL_SKYLAKE",
		"NVIDIA_GPU",
		"GENERIC",
		"LEGACY",
		"ARM",
	}
}

// GetMappingModelDeploymentShapeSeriesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModelDeploymentShapeSeriesEnum(val string) (ModelDeploymentShapeSeriesEnum, bool) {
	enum, ok := mappingModelDeploymentShapeSeriesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
