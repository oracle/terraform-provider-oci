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

// CreateDesktopPoolShapeConfigDetails The compute instance shape configuration requested for each desktop in the desktop pool.
type CreateDesktopPoolShapeConfigDetails struct {

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
	BaselineOcpuUtilization CreateDesktopPoolShapeConfigDetailsBaselineOcpuUtilizationEnum `mandatory:"false" json:"baselineOcpuUtilization,omitempty"`
}

func (m CreateDesktopPoolShapeConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDesktopPoolShapeConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateDesktopPoolShapeConfigDetailsBaselineOcpuUtilizationEnum(string(m.BaselineOcpuUtilization)); !ok && m.BaselineOcpuUtilization != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BaselineOcpuUtilization: %s. Supported values are: %s.", m.BaselineOcpuUtilization, strings.Join(GetCreateDesktopPoolShapeConfigDetailsBaselineOcpuUtilizationEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateDesktopPoolShapeConfigDetailsBaselineOcpuUtilizationEnum Enum with underlying type: string
type CreateDesktopPoolShapeConfigDetailsBaselineOcpuUtilizationEnum string

// Set of constants representing the allowable values for CreateDesktopPoolShapeConfigDetailsBaselineOcpuUtilizationEnum
const (
	CreateDesktopPoolShapeConfigDetailsBaselineOcpuUtilization8 CreateDesktopPoolShapeConfigDetailsBaselineOcpuUtilizationEnum = "BASELINE_1_8"
	CreateDesktopPoolShapeConfigDetailsBaselineOcpuUtilization2 CreateDesktopPoolShapeConfigDetailsBaselineOcpuUtilizationEnum = "BASELINE_1_2"
	CreateDesktopPoolShapeConfigDetailsBaselineOcpuUtilization1 CreateDesktopPoolShapeConfigDetailsBaselineOcpuUtilizationEnum = "BASELINE_1_1"
)

var mappingCreateDesktopPoolShapeConfigDetailsBaselineOcpuUtilizationEnum = map[string]CreateDesktopPoolShapeConfigDetailsBaselineOcpuUtilizationEnum{
	"BASELINE_1_8": CreateDesktopPoolShapeConfigDetailsBaselineOcpuUtilization8,
	"BASELINE_1_2": CreateDesktopPoolShapeConfigDetailsBaselineOcpuUtilization2,
	"BASELINE_1_1": CreateDesktopPoolShapeConfigDetailsBaselineOcpuUtilization1,
}

var mappingCreateDesktopPoolShapeConfigDetailsBaselineOcpuUtilizationEnumLowerCase = map[string]CreateDesktopPoolShapeConfigDetailsBaselineOcpuUtilizationEnum{
	"baseline_1_8": CreateDesktopPoolShapeConfigDetailsBaselineOcpuUtilization8,
	"baseline_1_2": CreateDesktopPoolShapeConfigDetailsBaselineOcpuUtilization2,
	"baseline_1_1": CreateDesktopPoolShapeConfigDetailsBaselineOcpuUtilization1,
}

// GetCreateDesktopPoolShapeConfigDetailsBaselineOcpuUtilizationEnumValues Enumerates the set of values for CreateDesktopPoolShapeConfigDetailsBaselineOcpuUtilizationEnum
func GetCreateDesktopPoolShapeConfigDetailsBaselineOcpuUtilizationEnumValues() []CreateDesktopPoolShapeConfigDetailsBaselineOcpuUtilizationEnum {
	values := make([]CreateDesktopPoolShapeConfigDetailsBaselineOcpuUtilizationEnum, 0)
	for _, v := range mappingCreateDesktopPoolShapeConfigDetailsBaselineOcpuUtilizationEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDesktopPoolShapeConfigDetailsBaselineOcpuUtilizationEnumStringValues Enumerates the set of values in String for CreateDesktopPoolShapeConfigDetailsBaselineOcpuUtilizationEnum
func GetCreateDesktopPoolShapeConfigDetailsBaselineOcpuUtilizationEnumStringValues() []string {
	return []string{
		"BASELINE_1_8",
		"BASELINE_1_2",
		"BASELINE_1_1",
	}
}

// GetMappingCreateDesktopPoolShapeConfigDetailsBaselineOcpuUtilizationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDesktopPoolShapeConfigDetailsBaselineOcpuUtilizationEnum(val string) (CreateDesktopPoolShapeConfigDetailsBaselineOcpuUtilizationEnum, bool) {
	enum, ok := mappingCreateDesktopPoolShapeConfigDetailsBaselineOcpuUtilizationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
