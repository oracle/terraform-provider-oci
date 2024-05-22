// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// JobShapeSummary The compute shape used to launch a job compute instance.
type JobShapeSummary struct {

	// The name of the job shape.
	Name *string `mandatory:"true" json:"name"`

	// The number of cores associated with this job run shape.
	CoreCount *int `mandatory:"true" json:"coreCount"`

	// The number of cores associated with this job shape.
	MemoryInGBs *int `mandatory:"true" json:"memoryInGBs"`

	// The family that the compute shape belongs to.
	ShapeSeries JobShapeSummaryShapeSeriesEnum `mandatory:"true" json:"shapeSeries"`
}

func (m JobShapeSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JobShapeSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingJobShapeSummaryShapeSeriesEnum(string(m.ShapeSeries)); !ok && m.ShapeSeries != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ShapeSeries: %s. Supported values are: %s.", m.ShapeSeries, strings.Join(GetJobShapeSummaryShapeSeriesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// JobShapeSummaryShapeSeriesEnum Enum with underlying type: string
type JobShapeSummaryShapeSeriesEnum string

// Set of constants representing the allowable values for JobShapeSummaryShapeSeriesEnum
const (
	JobShapeSummaryShapeSeriesAmdRome      JobShapeSummaryShapeSeriesEnum = "AMD_ROME"
	JobShapeSummaryShapeSeriesIntelSkylake JobShapeSummaryShapeSeriesEnum = "INTEL_SKYLAKE"
	JobShapeSummaryShapeSeriesNvidiaGpu    JobShapeSummaryShapeSeriesEnum = "NVIDIA_GPU"
	JobShapeSummaryShapeSeriesGeneric      JobShapeSummaryShapeSeriesEnum = "GENERIC"
	JobShapeSummaryShapeSeriesLegacy       JobShapeSummaryShapeSeriesEnum = "LEGACY"
	JobShapeSummaryShapeSeriesArm          JobShapeSummaryShapeSeriesEnum = "ARM"
)

var mappingJobShapeSummaryShapeSeriesEnum = map[string]JobShapeSummaryShapeSeriesEnum{
	"AMD_ROME":      JobShapeSummaryShapeSeriesAmdRome,
	"INTEL_SKYLAKE": JobShapeSummaryShapeSeriesIntelSkylake,
	"NVIDIA_GPU":    JobShapeSummaryShapeSeriesNvidiaGpu,
	"GENERIC":       JobShapeSummaryShapeSeriesGeneric,
	"LEGACY":        JobShapeSummaryShapeSeriesLegacy,
	"ARM":           JobShapeSummaryShapeSeriesArm,
}

var mappingJobShapeSummaryShapeSeriesEnumLowerCase = map[string]JobShapeSummaryShapeSeriesEnum{
	"amd_rome":      JobShapeSummaryShapeSeriesAmdRome,
	"intel_skylake": JobShapeSummaryShapeSeriesIntelSkylake,
	"nvidia_gpu":    JobShapeSummaryShapeSeriesNvidiaGpu,
	"generic":       JobShapeSummaryShapeSeriesGeneric,
	"legacy":        JobShapeSummaryShapeSeriesLegacy,
	"arm":           JobShapeSummaryShapeSeriesArm,
}

// GetJobShapeSummaryShapeSeriesEnumValues Enumerates the set of values for JobShapeSummaryShapeSeriesEnum
func GetJobShapeSummaryShapeSeriesEnumValues() []JobShapeSummaryShapeSeriesEnum {
	values := make([]JobShapeSummaryShapeSeriesEnum, 0)
	for _, v := range mappingJobShapeSummaryShapeSeriesEnum {
		values = append(values, v)
	}
	return values
}

// GetJobShapeSummaryShapeSeriesEnumStringValues Enumerates the set of values in String for JobShapeSummaryShapeSeriesEnum
func GetJobShapeSummaryShapeSeriesEnumStringValues() []string {
	return []string{
		"AMD_ROME",
		"INTEL_SKYLAKE",
		"NVIDIA_GPU",
		"GENERIC",
		"LEGACY",
		"ARM",
	}
}

// GetMappingJobShapeSummaryShapeSeriesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJobShapeSummaryShapeSeriesEnum(val string) (JobShapeSummaryShapeSeriesEnum, bool) {
	enum, ok := mappingJobShapeSummaryShapeSeriesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
