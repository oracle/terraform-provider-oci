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

// PipelineShapeConfigDetails Details for the pipeline step run shape configuration. Specify only when a flex shape is selected.
type PipelineShapeConfigDetails struct {

	// The total number of OCPUs available to the pipeline step run instance.
	Ocpus *float32 `mandatory:"false" json:"ocpus"`

	// The total amount of memory available to the pipeline step run instance GBs.
	MemoryInGBs *float32 `mandatory:"false" json:"memoryInGBs"`

	// The baseline OCPU utilization for a subcore burstable VM instance. If this attribute is left blank, it will default to `BASELINE_1_1`.
	// The following values are supported:
	//   BASELINE_1_8 - baseline usage is 1/8 of an OCPU.
	//   BASELINE_1_2 - baseline usage is 1/2 of an OCPU.
	//   BASELINE_1_1 - baseline usage is an entire OCPU. This represents a non-burstable instance.
	CpuBaseline PipelineShapeConfigDetailsCpuBaselineEnum `mandatory:"false" json:"cpuBaseline,omitempty"`

	// The total number of OCPUs available to the pipeline step run instance specified as a parameter. This overrides the ocpus value.
	// The request will fail if the parameters used are null or invalid.
	OcpusParameterized *string `mandatory:"false" json:"ocpusParameterized"`

	// The total amount of memory available to the pipeline step run instance in GBs specified as a parameter. This overrides the memoryInGBs value.
	// The request will fail if the parameters used are null or invalid.
	MemoryInGBsParameterized *string `mandatory:"false" json:"memoryInGBsParameterized"`
}

func (m PipelineShapeConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PipelineShapeConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPipelineShapeConfigDetailsCpuBaselineEnum(string(m.CpuBaseline)); !ok && m.CpuBaseline != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CpuBaseline: %s. Supported values are: %s.", m.CpuBaseline, strings.Join(GetPipelineShapeConfigDetailsCpuBaselineEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PipelineShapeConfigDetailsCpuBaselineEnum Enum with underlying type: string
type PipelineShapeConfigDetailsCpuBaselineEnum string

// Set of constants representing the allowable values for PipelineShapeConfigDetailsCpuBaselineEnum
const (
	PipelineShapeConfigDetailsCpuBaseline8 PipelineShapeConfigDetailsCpuBaselineEnum = "BASELINE_1_8"
	PipelineShapeConfigDetailsCpuBaseline2 PipelineShapeConfigDetailsCpuBaselineEnum = "BASELINE_1_2"
	PipelineShapeConfigDetailsCpuBaseline1 PipelineShapeConfigDetailsCpuBaselineEnum = "BASELINE_1_1"
)

var mappingPipelineShapeConfigDetailsCpuBaselineEnum = map[string]PipelineShapeConfigDetailsCpuBaselineEnum{
	"BASELINE_1_8": PipelineShapeConfigDetailsCpuBaseline8,
	"BASELINE_1_2": PipelineShapeConfigDetailsCpuBaseline2,
	"BASELINE_1_1": PipelineShapeConfigDetailsCpuBaseline1,
}

var mappingPipelineShapeConfigDetailsCpuBaselineEnumLowerCase = map[string]PipelineShapeConfigDetailsCpuBaselineEnum{
	"baseline_1_8": PipelineShapeConfigDetailsCpuBaseline8,
	"baseline_1_2": PipelineShapeConfigDetailsCpuBaseline2,
	"baseline_1_1": PipelineShapeConfigDetailsCpuBaseline1,
}

// GetPipelineShapeConfigDetailsCpuBaselineEnumValues Enumerates the set of values for PipelineShapeConfigDetailsCpuBaselineEnum
func GetPipelineShapeConfigDetailsCpuBaselineEnumValues() []PipelineShapeConfigDetailsCpuBaselineEnum {
	values := make([]PipelineShapeConfigDetailsCpuBaselineEnum, 0)
	for _, v := range mappingPipelineShapeConfigDetailsCpuBaselineEnum {
		values = append(values, v)
	}
	return values
}

// GetPipelineShapeConfigDetailsCpuBaselineEnumStringValues Enumerates the set of values in String for PipelineShapeConfigDetailsCpuBaselineEnum
func GetPipelineShapeConfigDetailsCpuBaselineEnumStringValues() []string {
	return []string{
		"BASELINE_1_8",
		"BASELINE_1_2",
		"BASELINE_1_1",
	}
}

// GetMappingPipelineShapeConfigDetailsCpuBaselineEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPipelineShapeConfigDetailsCpuBaselineEnum(val string) (PipelineShapeConfigDetailsCpuBaselineEnum, bool) {
	enum, ok := mappingPipelineShapeConfigDetailsCpuBaselineEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
