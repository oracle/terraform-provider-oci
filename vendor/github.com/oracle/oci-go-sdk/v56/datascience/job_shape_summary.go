// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"github.com/oracle/oci-go-sdk/v56/common"
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

// JobShapeSummaryShapeSeriesEnum Enum with underlying type: string
type JobShapeSummaryShapeSeriesEnum string

// Set of constants representing the allowable values for JobShapeSummaryShapeSeriesEnum
const (
	JobShapeSummaryShapeSeriesAmdRome      JobShapeSummaryShapeSeriesEnum = "AMD_ROME"
	JobShapeSummaryShapeSeriesIntelSkylake JobShapeSummaryShapeSeriesEnum = "INTEL_SKYLAKE"
	JobShapeSummaryShapeSeriesNvidiaGpu    JobShapeSummaryShapeSeriesEnum = "NVIDIA_GPU"
	JobShapeSummaryShapeSeriesLegacy       JobShapeSummaryShapeSeriesEnum = "LEGACY"
)

var mappingJobShapeSummaryShapeSeries = map[string]JobShapeSummaryShapeSeriesEnum{
	"AMD_ROME":      JobShapeSummaryShapeSeriesAmdRome,
	"INTEL_SKYLAKE": JobShapeSummaryShapeSeriesIntelSkylake,
	"NVIDIA_GPU":    JobShapeSummaryShapeSeriesNvidiaGpu,
	"LEGACY":        JobShapeSummaryShapeSeriesLegacy,
}

// GetJobShapeSummaryShapeSeriesEnumValues Enumerates the set of values for JobShapeSummaryShapeSeriesEnum
func GetJobShapeSummaryShapeSeriesEnumValues() []JobShapeSummaryShapeSeriesEnum {
	values := make([]JobShapeSummaryShapeSeriesEnum, 0)
	for _, v := range mappingJobShapeSummaryShapeSeries {
		values = append(values, v)
	}
	return values
}
