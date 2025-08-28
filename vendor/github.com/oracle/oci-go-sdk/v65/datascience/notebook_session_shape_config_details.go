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

// NotebookSessionShapeConfigDetails Details for the notebook session shape configuration.
type NotebookSessionShapeConfigDetails struct {

	// The total number of OCPUs available to the notebook session instance.
	Ocpus *float32 `mandatory:"false" json:"ocpus"`

	// The total amount of memory available to the notebook session instance, in gigabytes.
	MemoryInGBs *float32 `mandatory:"false" json:"memoryInGBs"`

	// The baseline OCPU utilization for a subcore burstable VM instance. If this attribute is left bank, it will default to `BASELINE_1_1`.
	// The following values are supported:
	//   BASELINE_1_8 - baseline usage is 1/8 of an OCPU.
	//   BASELINE_1_2 - baseline usage is 1/2 of an OCPU.
	//   BASELINE_1_1 - baseline usage is an entire OCPU. This represents a non-burstable instance.
	CpuBaseline NotebookSessionShapeConfigDetailsCpuBaselineEnum `mandatory:"false" json:"cpuBaseline,omitempty"`
}

func (m NotebookSessionShapeConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NotebookSessionShapeConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingNotebookSessionShapeConfigDetailsCpuBaselineEnum(string(m.CpuBaseline)); !ok && m.CpuBaseline != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CpuBaseline: %s. Supported values are: %s.", m.CpuBaseline, strings.Join(GetNotebookSessionShapeConfigDetailsCpuBaselineEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// NotebookSessionShapeConfigDetailsCpuBaselineEnum Enum with underlying type: string
type NotebookSessionShapeConfigDetailsCpuBaselineEnum string

// Set of constants representing the allowable values for NotebookSessionShapeConfigDetailsCpuBaselineEnum
const (
	NotebookSessionShapeConfigDetailsCpuBaseline8 NotebookSessionShapeConfigDetailsCpuBaselineEnum = "BASELINE_1_8"
	NotebookSessionShapeConfigDetailsCpuBaseline2 NotebookSessionShapeConfigDetailsCpuBaselineEnum = "BASELINE_1_2"
	NotebookSessionShapeConfigDetailsCpuBaseline1 NotebookSessionShapeConfigDetailsCpuBaselineEnum = "BASELINE_1_1"
)

var mappingNotebookSessionShapeConfigDetailsCpuBaselineEnum = map[string]NotebookSessionShapeConfigDetailsCpuBaselineEnum{
	"BASELINE_1_8": NotebookSessionShapeConfigDetailsCpuBaseline8,
	"BASELINE_1_2": NotebookSessionShapeConfigDetailsCpuBaseline2,
	"BASELINE_1_1": NotebookSessionShapeConfigDetailsCpuBaseline1,
}

var mappingNotebookSessionShapeConfigDetailsCpuBaselineEnumLowerCase = map[string]NotebookSessionShapeConfigDetailsCpuBaselineEnum{
	"baseline_1_8": NotebookSessionShapeConfigDetailsCpuBaseline8,
	"baseline_1_2": NotebookSessionShapeConfigDetailsCpuBaseline2,
	"baseline_1_1": NotebookSessionShapeConfigDetailsCpuBaseline1,
}

// GetNotebookSessionShapeConfigDetailsCpuBaselineEnumValues Enumerates the set of values for NotebookSessionShapeConfigDetailsCpuBaselineEnum
func GetNotebookSessionShapeConfigDetailsCpuBaselineEnumValues() []NotebookSessionShapeConfigDetailsCpuBaselineEnum {
	values := make([]NotebookSessionShapeConfigDetailsCpuBaselineEnum, 0)
	for _, v := range mappingNotebookSessionShapeConfigDetailsCpuBaselineEnum {
		values = append(values, v)
	}
	return values
}

// GetNotebookSessionShapeConfigDetailsCpuBaselineEnumStringValues Enumerates the set of values in String for NotebookSessionShapeConfigDetailsCpuBaselineEnum
func GetNotebookSessionShapeConfigDetailsCpuBaselineEnumStringValues() []string {
	return []string{
		"BASELINE_1_8",
		"BASELINE_1_2",
		"BASELINE_1_1",
	}
}

// GetMappingNotebookSessionShapeConfigDetailsCpuBaselineEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNotebookSessionShapeConfigDetailsCpuBaselineEnum(val string) (NotebookSessionShapeConfigDetailsCpuBaselineEnum, bool) {
	enum, ok := mappingNotebookSessionShapeConfigDetailsCpuBaselineEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
