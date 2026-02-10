// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CapacityReportInstanceShapeConfig The shape configuration for a shape in a capacity report.
type CapacityReportInstanceShapeConfig struct {

	// The total number of OCPUs available to the instance.
	Ocpus *float32 `mandatory:"false" json:"ocpus"`

	// The total amount of memory available to the instance, in gigabytes.
	MemoryInGBs *float32 `mandatory:"false" json:"memoryInGBs"`

	// The number of NVMe drives to be used for storage.
	Nvmes *int `mandatory:"false" json:"nvmes"`

	// The baseline OCPU utilization for a subcore burstable VM instance. Leave this attribute blank for a
	// non-burstable instance, or explicitly specify non-burstable with `BASELINE_1_1`.
	// The following values are supported:
	// - `BASELINE_1_8` - baseline usage is 1/8 of an OCPU.
	// - `BASELINE_1_2` - baseline usage is 1/2 of an OCPU.
	// - `BASELINE_1_1` - baseline usage is an entire OCPU. This represents a non-burstable instance.
	BaselineOcpuUtilization CapacityReportInstanceShapeConfigBaselineOcpuUtilizationEnum `mandatory:"false" json:"baselineOcpuUtilization,omitempty"`
}

func (m CapacityReportInstanceShapeConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CapacityReportInstanceShapeConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCapacityReportInstanceShapeConfigBaselineOcpuUtilizationEnum(string(m.BaselineOcpuUtilization)); !ok && m.BaselineOcpuUtilization != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BaselineOcpuUtilization: %s. Supported values are: %s.", m.BaselineOcpuUtilization, strings.Join(GetCapacityReportInstanceShapeConfigBaselineOcpuUtilizationEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CapacityReportInstanceShapeConfigBaselineOcpuUtilizationEnum Enum with underlying type: string
type CapacityReportInstanceShapeConfigBaselineOcpuUtilizationEnum string

// Set of constants representing the allowable values for CapacityReportInstanceShapeConfigBaselineOcpuUtilizationEnum
const (
	CapacityReportInstanceShapeConfigBaselineOcpuUtilization8 CapacityReportInstanceShapeConfigBaselineOcpuUtilizationEnum = "BASELINE_1_8"
	CapacityReportInstanceShapeConfigBaselineOcpuUtilization2 CapacityReportInstanceShapeConfigBaselineOcpuUtilizationEnum = "BASELINE_1_2"
	CapacityReportInstanceShapeConfigBaselineOcpuUtilization1 CapacityReportInstanceShapeConfigBaselineOcpuUtilizationEnum = "BASELINE_1_1"
)

var mappingCapacityReportInstanceShapeConfigBaselineOcpuUtilizationEnum = map[string]CapacityReportInstanceShapeConfigBaselineOcpuUtilizationEnum{
	"BASELINE_1_8": CapacityReportInstanceShapeConfigBaselineOcpuUtilization8,
	"BASELINE_1_2": CapacityReportInstanceShapeConfigBaselineOcpuUtilization2,
	"BASELINE_1_1": CapacityReportInstanceShapeConfigBaselineOcpuUtilization1,
}

var mappingCapacityReportInstanceShapeConfigBaselineOcpuUtilizationEnumLowerCase = map[string]CapacityReportInstanceShapeConfigBaselineOcpuUtilizationEnum{
	"baseline_1_8": CapacityReportInstanceShapeConfigBaselineOcpuUtilization8,
	"baseline_1_2": CapacityReportInstanceShapeConfigBaselineOcpuUtilization2,
	"baseline_1_1": CapacityReportInstanceShapeConfigBaselineOcpuUtilization1,
}

// GetCapacityReportInstanceShapeConfigBaselineOcpuUtilizationEnumValues Enumerates the set of values for CapacityReportInstanceShapeConfigBaselineOcpuUtilizationEnum
func GetCapacityReportInstanceShapeConfigBaselineOcpuUtilizationEnumValues() []CapacityReportInstanceShapeConfigBaselineOcpuUtilizationEnum {
	values := make([]CapacityReportInstanceShapeConfigBaselineOcpuUtilizationEnum, 0)
	for _, v := range mappingCapacityReportInstanceShapeConfigBaselineOcpuUtilizationEnum {
		values = append(values, v)
	}
	return values
}

// GetCapacityReportInstanceShapeConfigBaselineOcpuUtilizationEnumStringValues Enumerates the set of values in String for CapacityReportInstanceShapeConfigBaselineOcpuUtilizationEnum
func GetCapacityReportInstanceShapeConfigBaselineOcpuUtilizationEnumStringValues() []string {
	return []string{
		"BASELINE_1_8",
		"BASELINE_1_2",
		"BASELINE_1_1",
	}
}

// GetMappingCapacityReportInstanceShapeConfigBaselineOcpuUtilizationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCapacityReportInstanceShapeConfigBaselineOcpuUtilizationEnum(val string) (CapacityReportInstanceShapeConfigBaselineOcpuUtilizationEnum, bool) {
	enum, ok := mappingCapacityReportInstanceShapeConfigBaselineOcpuUtilizationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
