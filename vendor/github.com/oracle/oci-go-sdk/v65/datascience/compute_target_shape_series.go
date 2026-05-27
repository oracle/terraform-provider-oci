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

// ComputeTargetShapeSeriesEnum Enum with underlying type: string
type ComputeTargetShapeSeriesEnum string

// Set of constants representing the allowable values for ComputeTargetShapeSeriesEnum
const (
	ComputeTargetShapeSeriesAmdRome      ComputeTargetShapeSeriesEnum = "AMD_ROME"
	ComputeTargetShapeSeriesIntelSkylake ComputeTargetShapeSeriesEnum = "INTEL_SKYLAKE"
	ComputeTargetShapeSeriesNvidiaGpu    ComputeTargetShapeSeriesEnum = "NVIDIA_GPU"
	ComputeTargetShapeSeriesGeneric      ComputeTargetShapeSeriesEnum = "GENERIC"
	ComputeTargetShapeSeriesLegacy       ComputeTargetShapeSeriesEnum = "LEGACY"
	ComputeTargetShapeSeriesArm          ComputeTargetShapeSeriesEnum = "ARM"
)

var mappingComputeTargetShapeSeriesEnum = map[string]ComputeTargetShapeSeriesEnum{
	"AMD_ROME":      ComputeTargetShapeSeriesAmdRome,
	"INTEL_SKYLAKE": ComputeTargetShapeSeriesIntelSkylake,
	"NVIDIA_GPU":    ComputeTargetShapeSeriesNvidiaGpu,
	"GENERIC":       ComputeTargetShapeSeriesGeneric,
	"LEGACY":        ComputeTargetShapeSeriesLegacy,
	"ARM":           ComputeTargetShapeSeriesArm,
}

var mappingComputeTargetShapeSeriesEnumLowerCase = map[string]ComputeTargetShapeSeriesEnum{
	"amd_rome":      ComputeTargetShapeSeriesAmdRome,
	"intel_skylake": ComputeTargetShapeSeriesIntelSkylake,
	"nvidia_gpu":    ComputeTargetShapeSeriesNvidiaGpu,
	"generic":       ComputeTargetShapeSeriesGeneric,
	"legacy":        ComputeTargetShapeSeriesLegacy,
	"arm":           ComputeTargetShapeSeriesArm,
}

// GetComputeTargetShapeSeriesEnumValues Enumerates the set of values for ComputeTargetShapeSeriesEnum
func GetComputeTargetShapeSeriesEnumValues() []ComputeTargetShapeSeriesEnum {
	values := make([]ComputeTargetShapeSeriesEnum, 0)
	for _, v := range mappingComputeTargetShapeSeriesEnum {
		values = append(values, v)
	}
	return values
}

// GetComputeTargetShapeSeriesEnumStringValues Enumerates the set of values in String for ComputeTargetShapeSeriesEnum
func GetComputeTargetShapeSeriesEnumStringValues() []string {
	return []string{
		"AMD_ROME",
		"INTEL_SKYLAKE",
		"NVIDIA_GPU",
		"GENERIC",
		"LEGACY",
		"ARM",
	}
}

// GetMappingComputeTargetShapeSeriesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingComputeTargetShapeSeriesEnum(val string) (ComputeTargetShapeSeriesEnum, bool) {
	enum, ok := mappingComputeTargetShapeSeriesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
