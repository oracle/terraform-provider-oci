// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// JobShapeConfigDetails Details for the job run shape configuration. Specify only when a flex shape is selected.
type JobShapeConfigDetails struct {

	// The total number of OCPUs available to the job run instance.
	Ocpus *float32 `mandatory:"false" json:"ocpus"`

	// The total amount of memory available to the job run instance, in gigabytes.
	MemoryInGBs *float32 `mandatory:"false" json:"memoryInGBs"`

	// The baseline OCPU utilization for a subcore burstable VM instance. If this attribute is left blank, it will default to `BASELINE_1_1`.
	// The following values are supported:
	//   BASELINE_1_8 - baseline usage is 1/8 of an OCPU.
	//   BASELINE_1_2 - baseline usage is 1/2 of an OCPU.
	//   BASELINE_1_1 - baseline usage is an entire OCPU. This represents a non-burstable instance.
	CpuBaseline JobShapeConfigDetailsCpuBaselineEnum `mandatory:"false" json:"cpuBaseline,omitempty"`
}

func (m JobShapeConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JobShapeConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingJobShapeConfigDetailsCpuBaselineEnum(string(m.CpuBaseline)); !ok && m.CpuBaseline != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CpuBaseline: %s. Supported values are: %s.", m.CpuBaseline, strings.Join(GetJobShapeConfigDetailsCpuBaselineEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// JobShapeConfigDetailsCpuBaselineEnum Enum with underlying type: string
type JobShapeConfigDetailsCpuBaselineEnum string

// Set of constants representing the allowable values for JobShapeConfigDetailsCpuBaselineEnum
const (
	JobShapeConfigDetailsCpuBaseline8 JobShapeConfigDetailsCpuBaselineEnum = "BASELINE_1_8"
	JobShapeConfigDetailsCpuBaseline2 JobShapeConfigDetailsCpuBaselineEnum = "BASELINE_1_2"
	JobShapeConfigDetailsCpuBaseline1 JobShapeConfigDetailsCpuBaselineEnum = "BASELINE_1_1"
)

var mappingJobShapeConfigDetailsCpuBaselineEnum = map[string]JobShapeConfigDetailsCpuBaselineEnum{
	"BASELINE_1_8": JobShapeConfigDetailsCpuBaseline8,
	"BASELINE_1_2": JobShapeConfigDetailsCpuBaseline2,
	"BASELINE_1_1": JobShapeConfigDetailsCpuBaseline1,
}

var mappingJobShapeConfigDetailsCpuBaselineEnumLowerCase = map[string]JobShapeConfigDetailsCpuBaselineEnum{
	"baseline_1_8": JobShapeConfigDetailsCpuBaseline8,
	"baseline_1_2": JobShapeConfigDetailsCpuBaseline2,
	"baseline_1_1": JobShapeConfigDetailsCpuBaseline1,
}

// GetJobShapeConfigDetailsCpuBaselineEnumValues Enumerates the set of values for JobShapeConfigDetailsCpuBaselineEnum
func GetJobShapeConfigDetailsCpuBaselineEnumValues() []JobShapeConfigDetailsCpuBaselineEnum {
	values := make([]JobShapeConfigDetailsCpuBaselineEnum, 0)
	for _, v := range mappingJobShapeConfigDetailsCpuBaselineEnum {
		values = append(values, v)
	}
	return values
}

// GetJobShapeConfigDetailsCpuBaselineEnumStringValues Enumerates the set of values in String for JobShapeConfigDetailsCpuBaselineEnum
func GetJobShapeConfigDetailsCpuBaselineEnumStringValues() []string {
	return []string{
		"BASELINE_1_8",
		"BASELINE_1_2",
		"BASELINE_1_1",
	}
}

// GetMappingJobShapeConfigDetailsCpuBaselineEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJobShapeConfigDetailsCpuBaselineEnum(val string) (JobShapeConfigDetailsCpuBaselineEnum, bool) {
	enum, ok := mappingJobShapeConfigDetailsCpuBaselineEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
