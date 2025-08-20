// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ComputePerformanceSummary Parameters detailing the compute performance for a specified DB system shape.
type ComputePerformanceSummary struct {

	// The amount of memory allocated for the VMDB System.
	MemoryInGBs *float64 `mandatory:"true" json:"memoryInGBs"`

	// The number of CPU cores available.
	CpuCoreCount *int `mandatory:"false" json:"cpuCoreCount"`

	// The network bandwidth of the VMDB system in gbps.
	NetworkBandwidthInGbps *float32 `mandatory:"false" json:"networkBandwidthInGbps"`

	// IOPS for the VMDB System.
	NetworkIops *float32 `mandatory:"false" json:"networkIops"`

	// Network throughput for the VMDB System.
	NetworkThroughputInMbps *float32 `mandatory:"false" json:"networkThroughputInMbps"`

	// The compute model for Base Database Service. This is required if using the `computeCount` parameter. If using `cpuCoreCount` then it is an error to specify `computeModel` to a non-null value. The ECPU compute model is the recommended model, and the OCPU compute model is legacy.
	ComputeModel ComputePerformanceSummaryComputeModelEnum `mandatory:"false" json:"computeModel,omitempty"`

	// The number of compute servers for the DB system.
	ComputeCount *int `mandatory:"false" json:"computeCount"`
}

func (m ComputePerformanceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComputePerformanceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingComputePerformanceSummaryComputeModelEnum(string(m.ComputeModel)); !ok && m.ComputeModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ComputeModel: %s. Supported values are: %s.", m.ComputeModel, strings.Join(GetComputePerformanceSummaryComputeModelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ComputePerformanceSummaryComputeModelEnum Enum with underlying type: string
type ComputePerformanceSummaryComputeModelEnum string

// Set of constants representing the allowable values for ComputePerformanceSummaryComputeModelEnum
const (
	ComputePerformanceSummaryComputeModelEcpu ComputePerformanceSummaryComputeModelEnum = "ECPU"
	ComputePerformanceSummaryComputeModelOcpu ComputePerformanceSummaryComputeModelEnum = "OCPU"
)

var mappingComputePerformanceSummaryComputeModelEnum = map[string]ComputePerformanceSummaryComputeModelEnum{
	"ECPU": ComputePerformanceSummaryComputeModelEcpu,
	"OCPU": ComputePerformanceSummaryComputeModelOcpu,
}

var mappingComputePerformanceSummaryComputeModelEnumLowerCase = map[string]ComputePerformanceSummaryComputeModelEnum{
	"ecpu": ComputePerformanceSummaryComputeModelEcpu,
	"ocpu": ComputePerformanceSummaryComputeModelOcpu,
}

// GetComputePerformanceSummaryComputeModelEnumValues Enumerates the set of values for ComputePerformanceSummaryComputeModelEnum
func GetComputePerformanceSummaryComputeModelEnumValues() []ComputePerformanceSummaryComputeModelEnum {
	values := make([]ComputePerformanceSummaryComputeModelEnum, 0)
	for _, v := range mappingComputePerformanceSummaryComputeModelEnum {
		values = append(values, v)
	}
	return values
}

// GetComputePerformanceSummaryComputeModelEnumStringValues Enumerates the set of values in String for ComputePerformanceSummaryComputeModelEnum
func GetComputePerformanceSummaryComputeModelEnumStringValues() []string {
	return []string{
		"ECPU",
		"OCPU",
	}
}

// GetMappingComputePerformanceSummaryComputeModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingComputePerformanceSummaryComputeModelEnum(val string) (ComputePerformanceSummaryComputeModelEnum, bool) {
	enum, ok := mappingComputePerformanceSummaryComputeModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
