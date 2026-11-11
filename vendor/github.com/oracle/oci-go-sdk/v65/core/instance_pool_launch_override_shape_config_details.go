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

// InstancePoolLaunchOverrideShapeConfigDetails The shape configuration requested for the override option.
type InstancePoolLaunchOverrideShapeConfigDetails struct {

	// The total number of OCPUs available to the instance.
	Ocpus *float32 `mandatory:"false" json:"ocpus"`

	// The total number of VCPUs available to the instance. This can be used instead of OCPUs,
	// in which case the actual number of OCPUs will be calculated based on this value
	// and the actual hardware. This must be a multiple of 2.
	Vcpus *int `mandatory:"false" json:"vcpus"`

	// The total amount of memory available to the instance, in gigabytes.
	MemoryInGBs *float32 `mandatory:"false" json:"memoryInGBs"`

	// The baseline OCPU utilization for a subcore burstable VM instance. Leave this attribute blank for a
	// non-burstable instance, or explicitly specify non-burstable with `BASELINE_1_1`.
	// The following values are supported:
	// - `BASELINE_1_8` - baseline usage is 1/8 of an OCPU.
	// - `BASELINE_1_2` - baseline usage is 1/2 of an OCPU.
	// - `BASELINE_1_1` - baseline usage is an entire OCPU. This represents a non-burstable instance.
	BaselineOcpuUtilization InstancePoolLaunchOverrideShapeConfigDetailsBaselineOcpuUtilizationEnum `mandatory:"false" json:"baselineOcpuUtilization,omitempty"`

	// The number of NVMe drives to be used for storage. A single drive has 6.8 TB available.
	Nvmes *int `mandatory:"false" json:"nvmes"`

	// This field is reserved for internal use.
	ResourceManagement InstancePoolLaunchOverrideShapeConfigDetailsResourceManagementEnum `mandatory:"false" json:"resourceManagement,omitempty"`

	// The NVMe-backed local storage capacity, in GB, for flexible dense (DenseLV) VM shapes. If the selected shape
	// is DenseLV, the value must be greater than 0. For all other shapes, the value must be null (if specified);
	// any non-null value for a non-DenseLV shape results in an error.
	LocalVolumeSizeInGBs *int `mandatory:"false" json:"localVolumeSizeInGBs"`
}

func (m InstancePoolLaunchOverrideShapeConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstancePoolLaunchOverrideShapeConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingInstancePoolLaunchOverrideShapeConfigDetailsBaselineOcpuUtilizationEnum(string(m.BaselineOcpuUtilization)); !ok && m.BaselineOcpuUtilization != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BaselineOcpuUtilization: %s. Supported values are: %s.", m.BaselineOcpuUtilization, strings.Join(GetInstancePoolLaunchOverrideShapeConfigDetailsBaselineOcpuUtilizationEnumStringValues(), ",")))
	}
	if _, ok := GetMappingInstancePoolLaunchOverrideShapeConfigDetailsResourceManagementEnum(string(m.ResourceManagement)); !ok && m.ResourceManagement != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceManagement: %s. Supported values are: %s.", m.ResourceManagement, strings.Join(GetInstancePoolLaunchOverrideShapeConfigDetailsResourceManagementEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InstancePoolLaunchOverrideShapeConfigDetailsBaselineOcpuUtilizationEnum Enum with underlying type: string
type InstancePoolLaunchOverrideShapeConfigDetailsBaselineOcpuUtilizationEnum string

// Set of constants representing the allowable values for InstancePoolLaunchOverrideShapeConfigDetailsBaselineOcpuUtilizationEnum
const (
	InstancePoolLaunchOverrideShapeConfigDetailsBaselineOcpuUtilization8 InstancePoolLaunchOverrideShapeConfigDetailsBaselineOcpuUtilizationEnum = "BASELINE_1_8"
	InstancePoolLaunchOverrideShapeConfigDetailsBaselineOcpuUtilization2 InstancePoolLaunchOverrideShapeConfigDetailsBaselineOcpuUtilizationEnum = "BASELINE_1_2"
	InstancePoolLaunchOverrideShapeConfigDetailsBaselineOcpuUtilization1 InstancePoolLaunchOverrideShapeConfigDetailsBaselineOcpuUtilizationEnum = "BASELINE_1_1"
)

var mappingInstancePoolLaunchOverrideShapeConfigDetailsBaselineOcpuUtilizationEnum = map[string]InstancePoolLaunchOverrideShapeConfigDetailsBaselineOcpuUtilizationEnum{
	"BASELINE_1_8": InstancePoolLaunchOverrideShapeConfigDetailsBaselineOcpuUtilization8,
	"BASELINE_1_2": InstancePoolLaunchOverrideShapeConfigDetailsBaselineOcpuUtilization2,
	"BASELINE_1_1": InstancePoolLaunchOverrideShapeConfigDetailsBaselineOcpuUtilization1,
}

var mappingInstancePoolLaunchOverrideShapeConfigDetailsBaselineOcpuUtilizationEnumLowerCase = map[string]InstancePoolLaunchOverrideShapeConfigDetailsBaselineOcpuUtilizationEnum{
	"baseline_1_8": InstancePoolLaunchOverrideShapeConfigDetailsBaselineOcpuUtilization8,
	"baseline_1_2": InstancePoolLaunchOverrideShapeConfigDetailsBaselineOcpuUtilization2,
	"baseline_1_1": InstancePoolLaunchOverrideShapeConfigDetailsBaselineOcpuUtilization1,
}

// GetInstancePoolLaunchOverrideShapeConfigDetailsBaselineOcpuUtilizationEnumValues Enumerates the set of values for InstancePoolLaunchOverrideShapeConfigDetailsBaselineOcpuUtilizationEnum
func GetInstancePoolLaunchOverrideShapeConfigDetailsBaselineOcpuUtilizationEnumValues() []InstancePoolLaunchOverrideShapeConfigDetailsBaselineOcpuUtilizationEnum {
	values := make([]InstancePoolLaunchOverrideShapeConfigDetailsBaselineOcpuUtilizationEnum, 0)
	for _, v := range mappingInstancePoolLaunchOverrideShapeConfigDetailsBaselineOcpuUtilizationEnum {
		values = append(values, v)
	}
	return values
}

// GetInstancePoolLaunchOverrideShapeConfigDetailsBaselineOcpuUtilizationEnumStringValues Enumerates the set of values in String for InstancePoolLaunchOverrideShapeConfigDetailsBaselineOcpuUtilizationEnum
func GetInstancePoolLaunchOverrideShapeConfigDetailsBaselineOcpuUtilizationEnumStringValues() []string {
	return []string{
		"BASELINE_1_8",
		"BASELINE_1_2",
		"BASELINE_1_1",
	}
}

// GetMappingInstancePoolLaunchOverrideShapeConfigDetailsBaselineOcpuUtilizationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInstancePoolLaunchOverrideShapeConfigDetailsBaselineOcpuUtilizationEnum(val string) (InstancePoolLaunchOverrideShapeConfigDetailsBaselineOcpuUtilizationEnum, bool) {
	enum, ok := mappingInstancePoolLaunchOverrideShapeConfigDetailsBaselineOcpuUtilizationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// InstancePoolLaunchOverrideShapeConfigDetailsResourceManagementEnum Enum with underlying type: string
type InstancePoolLaunchOverrideShapeConfigDetailsResourceManagementEnum string

// Set of constants representing the allowable values for InstancePoolLaunchOverrideShapeConfigDetailsResourceManagementEnum
const (
	InstancePoolLaunchOverrideShapeConfigDetailsResourceManagementDynamic InstancePoolLaunchOverrideShapeConfigDetailsResourceManagementEnum = "DYNAMIC"
	InstancePoolLaunchOverrideShapeConfigDetailsResourceManagementStatic  InstancePoolLaunchOverrideShapeConfigDetailsResourceManagementEnum = "STATIC"
)

var mappingInstancePoolLaunchOverrideShapeConfigDetailsResourceManagementEnum = map[string]InstancePoolLaunchOverrideShapeConfigDetailsResourceManagementEnum{
	"DYNAMIC": InstancePoolLaunchOverrideShapeConfigDetailsResourceManagementDynamic,
	"STATIC":  InstancePoolLaunchOverrideShapeConfigDetailsResourceManagementStatic,
}

var mappingInstancePoolLaunchOverrideShapeConfigDetailsResourceManagementEnumLowerCase = map[string]InstancePoolLaunchOverrideShapeConfigDetailsResourceManagementEnum{
	"dynamic": InstancePoolLaunchOverrideShapeConfigDetailsResourceManagementDynamic,
	"static":  InstancePoolLaunchOverrideShapeConfigDetailsResourceManagementStatic,
}

// GetInstancePoolLaunchOverrideShapeConfigDetailsResourceManagementEnumValues Enumerates the set of values for InstancePoolLaunchOverrideShapeConfigDetailsResourceManagementEnum
func GetInstancePoolLaunchOverrideShapeConfigDetailsResourceManagementEnumValues() []InstancePoolLaunchOverrideShapeConfigDetailsResourceManagementEnum {
	values := make([]InstancePoolLaunchOverrideShapeConfigDetailsResourceManagementEnum, 0)
	for _, v := range mappingInstancePoolLaunchOverrideShapeConfigDetailsResourceManagementEnum {
		values = append(values, v)
	}
	return values
}

// GetInstancePoolLaunchOverrideShapeConfigDetailsResourceManagementEnumStringValues Enumerates the set of values in String for InstancePoolLaunchOverrideShapeConfigDetailsResourceManagementEnum
func GetInstancePoolLaunchOverrideShapeConfigDetailsResourceManagementEnumStringValues() []string {
	return []string{
		"DYNAMIC",
		"STATIC",
	}
}

// GetMappingInstancePoolLaunchOverrideShapeConfigDetailsResourceManagementEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInstancePoolLaunchOverrideShapeConfigDetailsResourceManagementEnum(val string) (InstancePoolLaunchOverrideShapeConfigDetailsResourceManagementEnum, bool) {
	enum, ok := mappingInstancePoolLaunchOverrideShapeConfigDetailsResourceManagementEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
