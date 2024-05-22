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

// NotebookSessionShapeSeriesEnum Enum with underlying type: string
type NotebookSessionShapeSeriesEnum string

// Set of constants representing the allowable values for NotebookSessionShapeSeriesEnum
const (
	NotebookSessionShapeSeriesAmdRome      NotebookSessionShapeSeriesEnum = "AMD_ROME"
	NotebookSessionShapeSeriesIntelSkylake NotebookSessionShapeSeriesEnum = "INTEL_SKYLAKE"
	NotebookSessionShapeSeriesNvidiaGpu    NotebookSessionShapeSeriesEnum = "NVIDIA_GPU"
	NotebookSessionShapeSeriesGeneric      NotebookSessionShapeSeriesEnum = "GENERIC"
	NotebookSessionShapeSeriesLegacy       NotebookSessionShapeSeriesEnum = "LEGACY"
	NotebookSessionShapeSeriesArm          NotebookSessionShapeSeriesEnum = "ARM"
)

var mappingNotebookSessionShapeSeriesEnum = map[string]NotebookSessionShapeSeriesEnum{
	"AMD_ROME":      NotebookSessionShapeSeriesAmdRome,
	"INTEL_SKYLAKE": NotebookSessionShapeSeriesIntelSkylake,
	"NVIDIA_GPU":    NotebookSessionShapeSeriesNvidiaGpu,
	"GENERIC":       NotebookSessionShapeSeriesGeneric,
	"LEGACY":        NotebookSessionShapeSeriesLegacy,
	"ARM":           NotebookSessionShapeSeriesArm,
}

var mappingNotebookSessionShapeSeriesEnumLowerCase = map[string]NotebookSessionShapeSeriesEnum{
	"amd_rome":      NotebookSessionShapeSeriesAmdRome,
	"intel_skylake": NotebookSessionShapeSeriesIntelSkylake,
	"nvidia_gpu":    NotebookSessionShapeSeriesNvidiaGpu,
	"generic":       NotebookSessionShapeSeriesGeneric,
	"legacy":        NotebookSessionShapeSeriesLegacy,
	"arm":           NotebookSessionShapeSeriesArm,
}

// GetNotebookSessionShapeSeriesEnumValues Enumerates the set of values for NotebookSessionShapeSeriesEnum
func GetNotebookSessionShapeSeriesEnumValues() []NotebookSessionShapeSeriesEnum {
	values := make([]NotebookSessionShapeSeriesEnum, 0)
	for _, v := range mappingNotebookSessionShapeSeriesEnum {
		values = append(values, v)
	}
	return values
}

// GetNotebookSessionShapeSeriesEnumStringValues Enumerates the set of values in String for NotebookSessionShapeSeriesEnum
func GetNotebookSessionShapeSeriesEnumStringValues() []string {
	return []string{
		"AMD_ROME",
		"INTEL_SKYLAKE",
		"NVIDIA_GPU",
		"GENERIC",
		"LEGACY",
		"ARM",
	}
}

// GetMappingNotebookSessionShapeSeriesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNotebookSessionShapeSeriesEnum(val string) (NotebookSessionShapeSeriesEnum, bool) {
	enum, ok := mappingNotebookSessionShapeSeriesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
