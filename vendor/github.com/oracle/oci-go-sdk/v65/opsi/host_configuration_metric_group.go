// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// HostConfigurationMetricGroup Base Metric Group for Host configuration metrics
type HostConfigurationMetricGroup interface {

	// Collection timestamp
	// Example: `"2020-05-06T00:00:00.000Z"`
	GetTimeCollected() *common.SDKTime
}

type hostconfigurationmetricgroup struct {
	JsonData      []byte
	TimeCollected *common.SDKTime `mandatory:"true" json:"timeCollected"`
	MetricName    string          `json:"metricName"`
}

// UnmarshalJSON unmarshals json
func (m *hostconfigurationmetricgroup) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerhostconfigurationmetricgroup hostconfigurationmetricgroup
	s := struct {
		Model Unmarshalerhostconfigurationmetricgroup
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.TimeCollected = s.Model.TimeCollected
	m.MetricName = s.Model.MetricName

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *hostconfigurationmetricgroup) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.MetricName {
	case "HOST_RESOURCE_ALLOCATION":
		mm := HostResourceAllocation{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HOST_PRODUCT":
		mm := HostProduct{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HOST_FILESYSTEM_CONFIGURATION":
		mm := HostFilesystemConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HOST_NETWORK_CONFIGURATION":
		mm := HostNetworkConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HOST_ENTITIES":
		mm := HostEntities{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HOST_MEMORY_CONFIGURATION":
		mm := HostMemoryConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HOST_CPU_HARDWARE_CONFIGURATION":
		mm := HostCpuHardwareConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HOST_GPU_CONFIGURATION":
		mm := HostGpuConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HOST_HARDWARE_CONFIGURATION":
		mm := HostHardwareConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for HostConfigurationMetricGroup: %s.", m.MetricName)
		return *m, nil
	}
}

// GetTimeCollected returns TimeCollected
func (m hostconfigurationmetricgroup) GetTimeCollected() *common.SDKTime {
	return m.TimeCollected
}

func (m hostconfigurationmetricgroup) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m hostconfigurationmetricgroup) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// HostConfigurationMetricGroupMetricNameEnum Enum with underlying type: string
type HostConfigurationMetricGroupMetricNameEnum string

// Set of constants representing the allowable values for HostConfigurationMetricGroupMetricNameEnum
const (
	HostConfigurationMetricGroupMetricNameProduct                  HostConfigurationMetricGroupMetricNameEnum = "HOST_PRODUCT"
	HostConfigurationMetricGroupMetricNameResourceAllocation       HostConfigurationMetricGroupMetricNameEnum = "HOST_RESOURCE_ALLOCATION"
	HostConfigurationMetricGroupMetricNameMemoryConfiguration      HostConfigurationMetricGroupMetricNameEnum = "HOST_MEMORY_CONFIGURATION"
	HostConfigurationMetricGroupMetricNameHardwareConfiguration    HostConfigurationMetricGroupMetricNameEnum = "HOST_HARDWARE_CONFIGURATION"
	HostConfigurationMetricGroupMetricNameCpuHardwareConfiguration HostConfigurationMetricGroupMetricNameEnum = "HOST_CPU_HARDWARE_CONFIGURATION"
	HostConfigurationMetricGroupMetricNameNetworkConfiguration     HostConfigurationMetricGroupMetricNameEnum = "HOST_NETWORK_CONFIGURATION"
	HostConfigurationMetricGroupMetricNameEntites                  HostConfigurationMetricGroupMetricNameEnum = "HOST_ENTITES"
	HostConfigurationMetricGroupMetricNameFilesystemConfiguration  HostConfigurationMetricGroupMetricNameEnum = "HOST_FILESYSTEM_CONFIGURATION"
	HostConfigurationMetricGroupMetricNameGpuConfiguration         HostConfigurationMetricGroupMetricNameEnum = "HOST_GPU_CONFIGURATION"
)

var mappingHostConfigurationMetricGroupMetricNameEnum = map[string]HostConfigurationMetricGroupMetricNameEnum{
	"HOST_PRODUCT":                    HostConfigurationMetricGroupMetricNameProduct,
	"HOST_RESOURCE_ALLOCATION":        HostConfigurationMetricGroupMetricNameResourceAllocation,
	"HOST_MEMORY_CONFIGURATION":       HostConfigurationMetricGroupMetricNameMemoryConfiguration,
	"HOST_HARDWARE_CONFIGURATION":     HostConfigurationMetricGroupMetricNameHardwareConfiguration,
	"HOST_CPU_HARDWARE_CONFIGURATION": HostConfigurationMetricGroupMetricNameCpuHardwareConfiguration,
	"HOST_NETWORK_CONFIGURATION":      HostConfigurationMetricGroupMetricNameNetworkConfiguration,
	"HOST_ENTITES":                    HostConfigurationMetricGroupMetricNameEntites,
	"HOST_FILESYSTEM_CONFIGURATION":   HostConfigurationMetricGroupMetricNameFilesystemConfiguration,
	"HOST_GPU_CONFIGURATION":          HostConfigurationMetricGroupMetricNameGpuConfiguration,
}

var mappingHostConfigurationMetricGroupMetricNameEnumLowerCase = map[string]HostConfigurationMetricGroupMetricNameEnum{
	"host_product":                    HostConfigurationMetricGroupMetricNameProduct,
	"host_resource_allocation":        HostConfigurationMetricGroupMetricNameResourceAllocation,
	"host_memory_configuration":       HostConfigurationMetricGroupMetricNameMemoryConfiguration,
	"host_hardware_configuration":     HostConfigurationMetricGroupMetricNameHardwareConfiguration,
	"host_cpu_hardware_configuration": HostConfigurationMetricGroupMetricNameCpuHardwareConfiguration,
	"host_network_configuration":      HostConfigurationMetricGroupMetricNameNetworkConfiguration,
	"host_entites":                    HostConfigurationMetricGroupMetricNameEntites,
	"host_filesystem_configuration":   HostConfigurationMetricGroupMetricNameFilesystemConfiguration,
	"host_gpu_configuration":          HostConfigurationMetricGroupMetricNameGpuConfiguration,
}

// GetHostConfigurationMetricGroupMetricNameEnumValues Enumerates the set of values for HostConfigurationMetricGroupMetricNameEnum
func GetHostConfigurationMetricGroupMetricNameEnumValues() []HostConfigurationMetricGroupMetricNameEnum {
	values := make([]HostConfigurationMetricGroupMetricNameEnum, 0)
	for _, v := range mappingHostConfigurationMetricGroupMetricNameEnum {
		values = append(values, v)
	}
	return values
}

// GetHostConfigurationMetricGroupMetricNameEnumStringValues Enumerates the set of values in String for HostConfigurationMetricGroupMetricNameEnum
func GetHostConfigurationMetricGroupMetricNameEnumStringValues() []string {
	return []string{
		"HOST_PRODUCT",
		"HOST_RESOURCE_ALLOCATION",
		"HOST_MEMORY_CONFIGURATION",
		"HOST_HARDWARE_CONFIGURATION",
		"HOST_CPU_HARDWARE_CONFIGURATION",
		"HOST_NETWORK_CONFIGURATION",
		"HOST_ENTITES",
		"HOST_FILESYSTEM_CONFIGURATION",
		"HOST_GPU_CONFIGURATION",
	}
}

// GetMappingHostConfigurationMetricGroupMetricNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHostConfigurationMetricGroupMetricNameEnum(val string) (HostConfigurationMetricGroupMetricNameEnum, bool) {
	enum, ok := mappingHostConfigurationMetricGroupMetricNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
