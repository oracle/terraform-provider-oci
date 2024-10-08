// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Secure Desktops API
//
// Create and manage cloud-hosted desktops which can be accessed from a web browser or installed client.
//

package desktops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DesktopPoolShapeConfig The shape configuration used for each desktop compute instance in the desktop pool.
type DesktopPoolShapeConfig struct {

	// The total number of OCPUs available for each desktop compute instance in the desktop pool.
	Ocpus *int64 `mandatory:"false" json:"ocpus"`

	// The total amount of memory available in gigabytes for each desktop compute instance in the desktop pool.
	MemoryInGBs *int64 `mandatory:"false" json:"memoryInGBs"`

	// The baseline OCPU utilization for a subcore burstable VM instance used for each desktop compute instance in
	// the desktop pool.
	// Leave this attribute blank for a non-burstable instance, or explicitly specify non-burstable with
	// `BASELINE_1_1`.
	//
	// The following values are supported:
	// - `BASELINE_1_8` - baseline usage is 1/8 of an OCPU.
	// - `BASELINE_1_2` - baseline usage is 1/2 of an OCPU.
	// - `BASELINE_1_1` - baseline usage is the entire OCPU. This represents a non-burstable instance.
	BaselineOcpuUtilization DesktopPoolShapeConfigBaselineOcpuUtilizationEnum `mandatory:"false" json:"baselineOcpuUtilization,omitempty"`
}

func (m DesktopPoolShapeConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DesktopPoolShapeConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDesktopPoolShapeConfigBaselineOcpuUtilizationEnum(string(m.BaselineOcpuUtilization)); !ok && m.BaselineOcpuUtilization != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BaselineOcpuUtilization: %s. Supported values are: %s.", m.BaselineOcpuUtilization, strings.Join(GetDesktopPoolShapeConfigBaselineOcpuUtilizationEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DesktopPoolShapeConfigBaselineOcpuUtilizationEnum Enum with underlying type: string
type DesktopPoolShapeConfigBaselineOcpuUtilizationEnum string

// Set of constants representing the allowable values for DesktopPoolShapeConfigBaselineOcpuUtilizationEnum
const (
	DesktopPoolShapeConfigBaselineOcpuUtilization8 DesktopPoolShapeConfigBaselineOcpuUtilizationEnum = "BASELINE_1_8"
	DesktopPoolShapeConfigBaselineOcpuUtilization2 DesktopPoolShapeConfigBaselineOcpuUtilizationEnum = "BASELINE_1_2"
	DesktopPoolShapeConfigBaselineOcpuUtilization1 DesktopPoolShapeConfigBaselineOcpuUtilizationEnum = "BASELINE_1_1"
)

var mappingDesktopPoolShapeConfigBaselineOcpuUtilizationEnum = map[string]DesktopPoolShapeConfigBaselineOcpuUtilizationEnum{
	"BASELINE_1_8": DesktopPoolShapeConfigBaselineOcpuUtilization8,
	"BASELINE_1_2": DesktopPoolShapeConfigBaselineOcpuUtilization2,
	"BASELINE_1_1": DesktopPoolShapeConfigBaselineOcpuUtilization1,
}

var mappingDesktopPoolShapeConfigBaselineOcpuUtilizationEnumLowerCase = map[string]DesktopPoolShapeConfigBaselineOcpuUtilizationEnum{
	"baseline_1_8": DesktopPoolShapeConfigBaselineOcpuUtilization8,
	"baseline_1_2": DesktopPoolShapeConfigBaselineOcpuUtilization2,
	"baseline_1_1": DesktopPoolShapeConfigBaselineOcpuUtilization1,
}

// GetDesktopPoolShapeConfigBaselineOcpuUtilizationEnumValues Enumerates the set of values for DesktopPoolShapeConfigBaselineOcpuUtilizationEnum
func GetDesktopPoolShapeConfigBaselineOcpuUtilizationEnumValues() []DesktopPoolShapeConfigBaselineOcpuUtilizationEnum {
	values := make([]DesktopPoolShapeConfigBaselineOcpuUtilizationEnum, 0)
	for _, v := range mappingDesktopPoolShapeConfigBaselineOcpuUtilizationEnum {
		values = append(values, v)
	}
	return values
}

// GetDesktopPoolShapeConfigBaselineOcpuUtilizationEnumStringValues Enumerates the set of values in String for DesktopPoolShapeConfigBaselineOcpuUtilizationEnum
func GetDesktopPoolShapeConfigBaselineOcpuUtilizationEnumStringValues() []string {
	return []string{
		"BASELINE_1_8",
		"BASELINE_1_2",
		"BASELINE_1_1",
	}
}

// GetMappingDesktopPoolShapeConfigBaselineOcpuUtilizationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDesktopPoolShapeConfigBaselineOcpuUtilizationEnum(val string) (DesktopPoolShapeConfigBaselineOcpuUtilizationEnum, bool) {
	enum, ok := mappingDesktopPoolShapeConfigBaselineOcpuUtilizationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
