// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// HostConfigurationSummary Summary of a host configuration for a resource.
type HostConfigurationSummary interface {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the host insight resource.
	GetHostInsightId() *string

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	GetCompartmentId() *string

	// The host name. The host name is unique amongst the hosts managed by the same management agent.
	GetHostName() *string

	// Platform type.
	// Supported platformType(s) for MACS-managed external host insight: [LINUX].
	// Supported platformType(s) for EM-managed external host insight: [LINUX, SOLARIS, SUNOS].
	GetPlatformType() HostConfigurationSummaryPlatformTypeEnum

	// Platform version.
	GetPlatformVersion() *string

	// Platform vendor.
	GetPlatformVendor() *string

	// Total CPU on this host.
	GetTotalCpus() *int

	// Total amount of usable physical memory in gibabytes
	GetTotalMemoryInGBs() *float64

	// CPU architechure
	GetCpuArchitecture() *string

	// Size of cache memory in megabytes.
	GetCpuCacheInMBs() *float64

	// Name of the CPU vendor.
	GetCpuVendor() *string

	// Clock frequency of the processor in megahertz.
	GetCpuFrequencyInMhz() *float64

	// Model name of processor.
	GetCpuImplementation() *string

	// Number of cores per socket.
	GetCoresPerSocket() *int

	// Number of total sockets.
	GetTotalSockets() *int

	// Number of threads per socket.
	GetThreadsPerSocket() *int

	// Indicates if hyper-threading is enabled or not
	GetIsHyperThreadingEnabled() *bool

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string
}

type hostconfigurationsummary struct {
	JsonData                []byte
	HostInsightId           *string                                  `mandatory:"true" json:"hostInsightId"`
	CompartmentId           *string                                  `mandatory:"true" json:"compartmentId"`
	HostName                *string                                  `mandatory:"true" json:"hostName"`
	PlatformType            HostConfigurationSummaryPlatformTypeEnum `mandatory:"true" json:"platformType"`
	PlatformVersion         *string                                  `mandatory:"true" json:"platformVersion"`
	PlatformVendor          *string                                  `mandatory:"true" json:"platformVendor"`
	TotalCpus               *int                                     `mandatory:"true" json:"totalCpus"`
	TotalMemoryInGBs        *float64                                 `mandatory:"true" json:"totalMemoryInGBs"`
	CpuArchitecture         *string                                  `mandatory:"true" json:"cpuArchitecture"`
	CpuCacheInMBs           *float64                                 `mandatory:"true" json:"cpuCacheInMBs"`
	CpuVendor               *string                                  `mandatory:"true" json:"cpuVendor"`
	CpuFrequencyInMhz       *float64                                 `mandatory:"true" json:"cpuFrequencyInMhz"`
	CpuImplementation       *string                                  `mandatory:"true" json:"cpuImplementation"`
	CoresPerSocket          *int                                     `mandatory:"true" json:"coresPerSocket"`
	TotalSockets            *int                                     `mandatory:"true" json:"totalSockets"`
	ThreadsPerSocket        *int                                     `mandatory:"true" json:"threadsPerSocket"`
	IsHyperThreadingEnabled *bool                                    `mandatory:"true" json:"isHyperThreadingEnabled"`
	DefinedTags             map[string]map[string]interface{}        `mandatory:"true" json:"definedTags"`
	FreeformTags            map[string]string                        `mandatory:"true" json:"freeformTags"`
	EntitySource            string                                   `json:"entitySource"`
}

// UnmarshalJSON unmarshals json
func (m *hostconfigurationsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerhostconfigurationsummary hostconfigurationsummary
	s := struct {
		Model Unmarshalerhostconfigurationsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.HostInsightId = s.Model.HostInsightId
	m.CompartmentId = s.Model.CompartmentId
	m.HostName = s.Model.HostName
	m.PlatformType = s.Model.PlatformType
	m.PlatformVersion = s.Model.PlatformVersion
	m.PlatformVendor = s.Model.PlatformVendor
	m.TotalCpus = s.Model.TotalCpus
	m.TotalMemoryInGBs = s.Model.TotalMemoryInGBs
	m.CpuArchitecture = s.Model.CpuArchitecture
	m.CpuCacheInMBs = s.Model.CpuCacheInMBs
	m.CpuVendor = s.Model.CpuVendor
	m.CpuFrequencyInMhz = s.Model.CpuFrequencyInMhz
	m.CpuImplementation = s.Model.CpuImplementation
	m.CoresPerSocket = s.Model.CoresPerSocket
	m.TotalSockets = s.Model.TotalSockets
	m.ThreadsPerSocket = s.Model.ThreadsPerSocket
	m.IsHyperThreadingEnabled = s.Model.IsHyperThreadingEnabled
	m.DefinedTags = s.Model.DefinedTags
	m.FreeformTags = s.Model.FreeformTags
	m.EntitySource = s.Model.EntitySource

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *hostconfigurationsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.EntitySource {
	case "MACS_MANAGED_EXTERNAL_HOST":
		mm := MacsManagedExternalHostConfigurationSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "EM_MANAGED_EXTERNAL_HOST":
		mm := EmManagedExternalHostConfigurationSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetHostInsightId returns HostInsightId
func (m hostconfigurationsummary) GetHostInsightId() *string {
	return m.HostInsightId
}

//GetCompartmentId returns CompartmentId
func (m hostconfigurationsummary) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetHostName returns HostName
func (m hostconfigurationsummary) GetHostName() *string {
	return m.HostName
}

//GetPlatformType returns PlatformType
func (m hostconfigurationsummary) GetPlatformType() HostConfigurationSummaryPlatformTypeEnum {
	return m.PlatformType
}

//GetPlatformVersion returns PlatformVersion
func (m hostconfigurationsummary) GetPlatformVersion() *string {
	return m.PlatformVersion
}

//GetPlatformVendor returns PlatformVendor
func (m hostconfigurationsummary) GetPlatformVendor() *string {
	return m.PlatformVendor
}

//GetTotalCpus returns TotalCpus
func (m hostconfigurationsummary) GetTotalCpus() *int {
	return m.TotalCpus
}

//GetTotalMemoryInGBs returns TotalMemoryInGBs
func (m hostconfigurationsummary) GetTotalMemoryInGBs() *float64 {
	return m.TotalMemoryInGBs
}

//GetCpuArchitecture returns CpuArchitecture
func (m hostconfigurationsummary) GetCpuArchitecture() *string {
	return m.CpuArchitecture
}

//GetCpuCacheInMBs returns CpuCacheInMBs
func (m hostconfigurationsummary) GetCpuCacheInMBs() *float64 {
	return m.CpuCacheInMBs
}

//GetCpuVendor returns CpuVendor
func (m hostconfigurationsummary) GetCpuVendor() *string {
	return m.CpuVendor
}

//GetCpuFrequencyInMhz returns CpuFrequencyInMhz
func (m hostconfigurationsummary) GetCpuFrequencyInMhz() *float64 {
	return m.CpuFrequencyInMhz
}

//GetCpuImplementation returns CpuImplementation
func (m hostconfigurationsummary) GetCpuImplementation() *string {
	return m.CpuImplementation
}

//GetCoresPerSocket returns CoresPerSocket
func (m hostconfigurationsummary) GetCoresPerSocket() *int {
	return m.CoresPerSocket
}

//GetTotalSockets returns TotalSockets
func (m hostconfigurationsummary) GetTotalSockets() *int {
	return m.TotalSockets
}

//GetThreadsPerSocket returns ThreadsPerSocket
func (m hostconfigurationsummary) GetThreadsPerSocket() *int {
	return m.ThreadsPerSocket
}

//GetIsHyperThreadingEnabled returns IsHyperThreadingEnabled
func (m hostconfigurationsummary) GetIsHyperThreadingEnabled() *bool {
	return m.IsHyperThreadingEnabled
}

//GetDefinedTags returns DefinedTags
func (m hostconfigurationsummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetFreeformTags returns FreeformTags
func (m hostconfigurationsummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

func (m hostconfigurationsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m hostconfigurationsummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingHostConfigurationSummaryPlatformTypeEnum(string(m.PlatformType)); !ok && m.PlatformType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlatformType: %s. Supported values are: %s.", m.PlatformType, strings.Join(GetHostConfigurationSummaryPlatformTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// HostConfigurationSummaryPlatformTypeEnum Enum with underlying type: string
type HostConfigurationSummaryPlatformTypeEnum string

// Set of constants representing the allowable values for HostConfigurationSummaryPlatformTypeEnum
const (
	HostConfigurationSummaryPlatformTypeLinux   HostConfigurationSummaryPlatformTypeEnum = "LINUX"
	HostConfigurationSummaryPlatformTypeSolaris HostConfigurationSummaryPlatformTypeEnum = "SOLARIS"
	HostConfigurationSummaryPlatformTypeSunos   HostConfigurationSummaryPlatformTypeEnum = "SUNOS"
)

var mappingHostConfigurationSummaryPlatformTypeEnum = map[string]HostConfigurationSummaryPlatformTypeEnum{
	"LINUX":   HostConfigurationSummaryPlatformTypeLinux,
	"SOLARIS": HostConfigurationSummaryPlatformTypeSolaris,
	"SUNOS":   HostConfigurationSummaryPlatformTypeSunos,
}

// GetHostConfigurationSummaryPlatformTypeEnumValues Enumerates the set of values for HostConfigurationSummaryPlatformTypeEnum
func GetHostConfigurationSummaryPlatformTypeEnumValues() []HostConfigurationSummaryPlatformTypeEnum {
	values := make([]HostConfigurationSummaryPlatformTypeEnum, 0)
	for _, v := range mappingHostConfigurationSummaryPlatformTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetHostConfigurationSummaryPlatformTypeEnumStringValues Enumerates the set of values in String for HostConfigurationSummaryPlatformTypeEnum
func GetHostConfigurationSummaryPlatformTypeEnumStringValues() []string {
	return []string{
		"LINUX",
		"SOLARIS",
		"SUNOS",
	}
}

// GetMappingHostConfigurationSummaryPlatformTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHostConfigurationSummaryPlatformTypeEnum(val string) (HostConfigurationSummaryPlatformTypeEnum, bool) {
	mappingHostConfigurationSummaryPlatformTypeEnumIgnoreCase := make(map[string]HostConfigurationSummaryPlatformTypeEnum)
	for k, v := range mappingHostConfigurationSummaryPlatformTypeEnum {
		mappingHostConfigurationSummaryPlatformTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingHostConfigurationSummaryPlatformTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
