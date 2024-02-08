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

// FastLaunchJobConfigSummary The shape config to launch a fast launch capable job instance
type FastLaunchJobConfigSummary struct {

	// The name of the fast launch job config
	Name *string `mandatory:"true" json:"name"`

	// The name of the fast launch job shape.
	ShapeName *string `mandatory:"true" json:"shapeName"`

	// The number of cores associated with this fast launch job shape.
	CoreCount *int `mandatory:"true" json:"coreCount"`

	// The number of cores associated with this fast launch job shape.
	MemoryInGBs *int `mandatory:"true" json:"memoryInGBs"`

	// The family that the compute shape belongs to.
	ShapeSeries FastLaunchJobConfigSummaryShapeSeriesEnum `mandatory:"true" json:"shapeSeries"`

	// The managed egress support
	ManagedEgressSupport FastLaunchJobConfigSummaryManagedEgressSupportEnum `mandatory:"true" json:"managedEgressSupport"`
}

func (m FastLaunchJobConfigSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FastLaunchJobConfigSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFastLaunchJobConfigSummaryShapeSeriesEnum(string(m.ShapeSeries)); !ok && m.ShapeSeries != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ShapeSeries: %s. Supported values are: %s.", m.ShapeSeries, strings.Join(GetFastLaunchJobConfigSummaryShapeSeriesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingFastLaunchJobConfigSummaryManagedEgressSupportEnum(string(m.ManagedEgressSupport)); !ok && m.ManagedEgressSupport != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ManagedEgressSupport: %s. Supported values are: %s.", m.ManagedEgressSupport, strings.Join(GetFastLaunchJobConfigSummaryManagedEgressSupportEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FastLaunchJobConfigSummaryShapeSeriesEnum Enum with underlying type: string
type FastLaunchJobConfigSummaryShapeSeriesEnum string

// Set of constants representing the allowable values for FastLaunchJobConfigSummaryShapeSeriesEnum
const (
	FastLaunchJobConfigSummaryShapeSeriesAmdRome      FastLaunchJobConfigSummaryShapeSeriesEnum = "AMD_ROME"
	FastLaunchJobConfigSummaryShapeSeriesIntelSkylake FastLaunchJobConfigSummaryShapeSeriesEnum = "INTEL_SKYLAKE"
	FastLaunchJobConfigSummaryShapeSeriesNvidiaGpu    FastLaunchJobConfigSummaryShapeSeriesEnum = "NVIDIA_GPU"
	FastLaunchJobConfigSummaryShapeSeriesLegacy       FastLaunchJobConfigSummaryShapeSeriesEnum = "LEGACY"
	FastLaunchJobConfigSummaryShapeSeriesArm          FastLaunchJobConfigSummaryShapeSeriesEnum = "ARM"
)

var mappingFastLaunchJobConfigSummaryShapeSeriesEnum = map[string]FastLaunchJobConfigSummaryShapeSeriesEnum{
	"AMD_ROME":      FastLaunchJobConfigSummaryShapeSeriesAmdRome,
	"INTEL_SKYLAKE": FastLaunchJobConfigSummaryShapeSeriesIntelSkylake,
	"NVIDIA_GPU":    FastLaunchJobConfigSummaryShapeSeriesNvidiaGpu,
	"LEGACY":        FastLaunchJobConfigSummaryShapeSeriesLegacy,
	"ARM":           FastLaunchJobConfigSummaryShapeSeriesArm,
}

var mappingFastLaunchJobConfigSummaryShapeSeriesEnumLowerCase = map[string]FastLaunchJobConfigSummaryShapeSeriesEnum{
	"amd_rome":      FastLaunchJobConfigSummaryShapeSeriesAmdRome,
	"intel_skylake": FastLaunchJobConfigSummaryShapeSeriesIntelSkylake,
	"nvidia_gpu":    FastLaunchJobConfigSummaryShapeSeriesNvidiaGpu,
	"legacy":        FastLaunchJobConfigSummaryShapeSeriesLegacy,
	"arm":           FastLaunchJobConfigSummaryShapeSeriesArm,
}

// GetFastLaunchJobConfigSummaryShapeSeriesEnumValues Enumerates the set of values for FastLaunchJobConfigSummaryShapeSeriesEnum
func GetFastLaunchJobConfigSummaryShapeSeriesEnumValues() []FastLaunchJobConfigSummaryShapeSeriesEnum {
	values := make([]FastLaunchJobConfigSummaryShapeSeriesEnum, 0)
	for _, v := range mappingFastLaunchJobConfigSummaryShapeSeriesEnum {
		values = append(values, v)
	}
	return values
}

// GetFastLaunchJobConfigSummaryShapeSeriesEnumStringValues Enumerates the set of values in String for FastLaunchJobConfigSummaryShapeSeriesEnum
func GetFastLaunchJobConfigSummaryShapeSeriesEnumStringValues() []string {
	return []string{
		"AMD_ROME",
		"INTEL_SKYLAKE",
		"NVIDIA_GPU",
		"LEGACY",
		"ARM",
	}
}

// GetMappingFastLaunchJobConfigSummaryShapeSeriesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFastLaunchJobConfigSummaryShapeSeriesEnum(val string) (FastLaunchJobConfigSummaryShapeSeriesEnum, bool) {
	enum, ok := mappingFastLaunchJobConfigSummaryShapeSeriesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// FastLaunchJobConfigSummaryManagedEgressSupportEnum Enum with underlying type: string
type FastLaunchJobConfigSummaryManagedEgressSupportEnum string

// Set of constants representing the allowable values for FastLaunchJobConfigSummaryManagedEgressSupportEnum
const (
	FastLaunchJobConfigSummaryManagedEgressSupportRequired    FastLaunchJobConfigSummaryManagedEgressSupportEnum = "REQUIRED"
	FastLaunchJobConfigSummaryManagedEgressSupportSupported   FastLaunchJobConfigSummaryManagedEgressSupportEnum = "SUPPORTED"
	FastLaunchJobConfigSummaryManagedEgressSupportUnsupported FastLaunchJobConfigSummaryManagedEgressSupportEnum = "UNSUPPORTED"
)

var mappingFastLaunchJobConfigSummaryManagedEgressSupportEnum = map[string]FastLaunchJobConfigSummaryManagedEgressSupportEnum{
	"REQUIRED":    FastLaunchJobConfigSummaryManagedEgressSupportRequired,
	"SUPPORTED":   FastLaunchJobConfigSummaryManagedEgressSupportSupported,
	"UNSUPPORTED": FastLaunchJobConfigSummaryManagedEgressSupportUnsupported,
}

var mappingFastLaunchJobConfigSummaryManagedEgressSupportEnumLowerCase = map[string]FastLaunchJobConfigSummaryManagedEgressSupportEnum{
	"required":    FastLaunchJobConfigSummaryManagedEgressSupportRequired,
	"supported":   FastLaunchJobConfigSummaryManagedEgressSupportSupported,
	"unsupported": FastLaunchJobConfigSummaryManagedEgressSupportUnsupported,
}

// GetFastLaunchJobConfigSummaryManagedEgressSupportEnumValues Enumerates the set of values for FastLaunchJobConfigSummaryManagedEgressSupportEnum
func GetFastLaunchJobConfigSummaryManagedEgressSupportEnumValues() []FastLaunchJobConfigSummaryManagedEgressSupportEnum {
	values := make([]FastLaunchJobConfigSummaryManagedEgressSupportEnum, 0)
	for _, v := range mappingFastLaunchJobConfigSummaryManagedEgressSupportEnum {
		values = append(values, v)
	}
	return values
}

// GetFastLaunchJobConfigSummaryManagedEgressSupportEnumStringValues Enumerates the set of values in String for FastLaunchJobConfigSummaryManagedEgressSupportEnum
func GetFastLaunchJobConfigSummaryManagedEgressSupportEnumStringValues() []string {
	return []string{
		"REQUIRED",
		"SUPPORTED",
		"UNSUPPORTED",
	}
}

// GetMappingFastLaunchJobConfigSummaryManagedEgressSupportEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFastLaunchJobConfigSummaryManagedEgressSupportEnum(val string) (FastLaunchJobConfigSummaryManagedEgressSupportEnum, bool) {
	enum, ok := mappingFastLaunchJobConfigSummaryManagedEgressSupportEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
