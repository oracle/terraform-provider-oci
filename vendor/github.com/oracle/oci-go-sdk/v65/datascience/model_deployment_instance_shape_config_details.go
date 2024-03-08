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

// ModelDeploymentInstanceShapeConfigDetails Details for the model-deployment instance shape configuration.
type ModelDeploymentInstanceShapeConfigDetails struct {

	// A model-deployment instance of type VM.Standard.E3.Flex or VM.Standard.E4.Flex allows the ocpu count to be specified with in the range of 1 to 64 ocpu. VM.Standard3.Flex OCPU range is between 1 to 32 ocpu and for VM.Optimized3.Flex OCPU range is 1 to 18 ocpu.
	Ocpus *float32 `mandatory:"false" json:"ocpus"`

	// A model-deployment instance of type VM.Standard.E3.Flex or VM.Standard.E4.Flex allows the memory to be specified with in the range of 6 to 1024 GB. VM.Standard3.Flex memory range is between 6 to 512 GB and VM.Optimized3.Flex memory range is between 6 to 256 GB.
	MemoryInGBs *float32 `mandatory:"false" json:"memoryInGBs"`

	// The baseline OCPU utilization for a subcore burstable VM instance. If this attribute is left blank, it will default to `BASELINE_1_1`.
	// The following values are supported:
	//   BASELINE_1_8 - baseline usage is 1/8 of an OCPU.
	//   BASELINE_1_2 - baseline usage is 1/2 of an OCPU.
	//   BASELINE_1_1 - baseline usage is an entire OCPU. This represents a non-burstable instance.
	CpuBaseline ModelDeploymentInstanceShapeConfigDetailsCpuBaselineEnum `mandatory:"false" json:"cpuBaseline,omitempty"`
}

func (m ModelDeploymentInstanceShapeConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ModelDeploymentInstanceShapeConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingModelDeploymentInstanceShapeConfigDetailsCpuBaselineEnum(string(m.CpuBaseline)); !ok && m.CpuBaseline != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CpuBaseline: %s. Supported values are: %s.", m.CpuBaseline, strings.Join(GetModelDeploymentInstanceShapeConfigDetailsCpuBaselineEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ModelDeploymentInstanceShapeConfigDetailsCpuBaselineEnum Enum with underlying type: string
type ModelDeploymentInstanceShapeConfigDetailsCpuBaselineEnum string

// Set of constants representing the allowable values for ModelDeploymentInstanceShapeConfigDetailsCpuBaselineEnum
const (
	ModelDeploymentInstanceShapeConfigDetailsCpuBaseline8 ModelDeploymentInstanceShapeConfigDetailsCpuBaselineEnum = "BASELINE_1_8"
	ModelDeploymentInstanceShapeConfigDetailsCpuBaseline2 ModelDeploymentInstanceShapeConfigDetailsCpuBaselineEnum = "BASELINE_1_2"
	ModelDeploymentInstanceShapeConfigDetailsCpuBaseline1 ModelDeploymentInstanceShapeConfigDetailsCpuBaselineEnum = "BASELINE_1_1"
)

var mappingModelDeploymentInstanceShapeConfigDetailsCpuBaselineEnum = map[string]ModelDeploymentInstanceShapeConfigDetailsCpuBaselineEnum{
	"BASELINE_1_8": ModelDeploymentInstanceShapeConfigDetailsCpuBaseline8,
	"BASELINE_1_2": ModelDeploymentInstanceShapeConfigDetailsCpuBaseline2,
	"BASELINE_1_1": ModelDeploymentInstanceShapeConfigDetailsCpuBaseline1,
}

var mappingModelDeploymentInstanceShapeConfigDetailsCpuBaselineEnumLowerCase = map[string]ModelDeploymentInstanceShapeConfigDetailsCpuBaselineEnum{
	"baseline_1_8": ModelDeploymentInstanceShapeConfigDetailsCpuBaseline8,
	"baseline_1_2": ModelDeploymentInstanceShapeConfigDetailsCpuBaseline2,
	"baseline_1_1": ModelDeploymentInstanceShapeConfigDetailsCpuBaseline1,
}

// GetModelDeploymentInstanceShapeConfigDetailsCpuBaselineEnumValues Enumerates the set of values for ModelDeploymentInstanceShapeConfigDetailsCpuBaselineEnum
func GetModelDeploymentInstanceShapeConfigDetailsCpuBaselineEnumValues() []ModelDeploymentInstanceShapeConfigDetailsCpuBaselineEnum {
	values := make([]ModelDeploymentInstanceShapeConfigDetailsCpuBaselineEnum, 0)
	for _, v := range mappingModelDeploymentInstanceShapeConfigDetailsCpuBaselineEnum {
		values = append(values, v)
	}
	return values
}

// GetModelDeploymentInstanceShapeConfigDetailsCpuBaselineEnumStringValues Enumerates the set of values in String for ModelDeploymentInstanceShapeConfigDetailsCpuBaselineEnum
func GetModelDeploymentInstanceShapeConfigDetailsCpuBaselineEnumStringValues() []string {
	return []string{
		"BASELINE_1_8",
		"BASELINE_1_2",
		"BASELINE_1_1",
	}
}

// GetMappingModelDeploymentInstanceShapeConfigDetailsCpuBaselineEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModelDeploymentInstanceShapeConfigDetailsCpuBaselineEnum(val string) (ModelDeploymentInstanceShapeConfigDetailsCpuBaselineEnum, bool) {
	enum, ok := mappingModelDeploymentInstanceShapeConfigDetailsCpuBaselineEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
