// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PeComanagedHostConfigurationSummary Configuration Summary of a PeComanaged host.
type PeComanagedHostConfigurationSummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the host insight resource.
	HostInsightId *string `mandatory:"true" json:"hostInsightId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The host name. The host name is unique amongst the hosts managed by the same management agent.
	HostName *string `mandatory:"true" json:"hostName"`

	// Platform version.
	PlatformVersion *string `mandatory:"true" json:"platformVersion"`

	// Platform vendor.
	PlatformVendor *string `mandatory:"true" json:"platformVendor"`

	// Total CPU on this host.
	TotalCpus *int `mandatory:"true" json:"totalCpus"`

	// Total amount of usable physical memory in gibabytes
	TotalMemoryInGBs *float64 `mandatory:"true" json:"totalMemoryInGBs"`

	// CPU architechure
	CpuArchitecture *string `mandatory:"true" json:"cpuArchitecture"`

	// Size of cache memory in megabytes.
	CpuCacheInMBs *float64 `mandatory:"true" json:"cpuCacheInMBs"`

	// Name of the CPU vendor.
	CpuVendor *string `mandatory:"true" json:"cpuVendor"`

	// Clock frequency of the processor in megahertz.
	CpuFrequencyInMhz *float64 `mandatory:"true" json:"cpuFrequencyInMhz"`

	// Model name of processor.
	CpuImplementation *string `mandatory:"true" json:"cpuImplementation"`

	// Number of cores per socket.
	CoresPerSocket *int `mandatory:"true" json:"coresPerSocket"`

	// Number of total sockets.
	TotalSockets *int `mandatory:"true" json:"totalSockets"`

	// Number of threads per socket.
	ThreadsPerSocket *int `mandatory:"true" json:"threadsPerSocket"`

	// Indicates if hyper-threading is enabled or not
	IsHyperThreadingEnabled *bool `mandatory:"true" json:"isHyperThreadingEnabled"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the OPSI private endpoint
	OpsiPrivateEndpointId *string `mandatory:"true" json:"opsiPrivateEndpointId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.
	ParentId *string `mandatory:"true" json:"parentId"`

	ExadataDetails *ExadataDetails `mandatory:"true" json:"exadataDetails"`

	// Platform type.
	// Supported platformType(s) for MACS-managed external host insight: [LINUX, SOLARIS, WINDOWS].
	// Supported platformType(s) for MACS-managed cloud host insight: [LINUX].
	// Supported platformType(s) for EM-managed external host insight: [LINUX, SOLARIS, SUNOS, ZLINUX, WINDOWS, AIX, HP-UX].
	PlatformType HostConfigurationSummaryPlatformTypeEnum `mandatory:"true" json:"platformType"`
}

// GetHostInsightId returns HostInsightId
func (m PeComanagedHostConfigurationSummary) GetHostInsightId() *string {
	return m.HostInsightId
}

// GetCompartmentId returns CompartmentId
func (m PeComanagedHostConfigurationSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetHostName returns HostName
func (m PeComanagedHostConfigurationSummary) GetHostName() *string {
	return m.HostName
}

// GetPlatformType returns PlatformType
func (m PeComanagedHostConfigurationSummary) GetPlatformType() HostConfigurationSummaryPlatformTypeEnum {
	return m.PlatformType
}

// GetPlatformVersion returns PlatformVersion
func (m PeComanagedHostConfigurationSummary) GetPlatformVersion() *string {
	return m.PlatformVersion
}

// GetPlatformVendor returns PlatformVendor
func (m PeComanagedHostConfigurationSummary) GetPlatformVendor() *string {
	return m.PlatformVendor
}

// GetTotalCpus returns TotalCpus
func (m PeComanagedHostConfigurationSummary) GetTotalCpus() *int {
	return m.TotalCpus
}

// GetTotalMemoryInGBs returns TotalMemoryInGBs
func (m PeComanagedHostConfigurationSummary) GetTotalMemoryInGBs() *float64 {
	return m.TotalMemoryInGBs
}

// GetCpuArchitecture returns CpuArchitecture
func (m PeComanagedHostConfigurationSummary) GetCpuArchitecture() *string {
	return m.CpuArchitecture
}

// GetCpuCacheInMBs returns CpuCacheInMBs
func (m PeComanagedHostConfigurationSummary) GetCpuCacheInMBs() *float64 {
	return m.CpuCacheInMBs
}

// GetCpuVendor returns CpuVendor
func (m PeComanagedHostConfigurationSummary) GetCpuVendor() *string {
	return m.CpuVendor
}

// GetCpuFrequencyInMhz returns CpuFrequencyInMhz
func (m PeComanagedHostConfigurationSummary) GetCpuFrequencyInMhz() *float64 {
	return m.CpuFrequencyInMhz
}

// GetCpuImplementation returns CpuImplementation
func (m PeComanagedHostConfigurationSummary) GetCpuImplementation() *string {
	return m.CpuImplementation
}

// GetCoresPerSocket returns CoresPerSocket
func (m PeComanagedHostConfigurationSummary) GetCoresPerSocket() *int {
	return m.CoresPerSocket
}

// GetTotalSockets returns TotalSockets
func (m PeComanagedHostConfigurationSummary) GetTotalSockets() *int {
	return m.TotalSockets
}

// GetThreadsPerSocket returns ThreadsPerSocket
func (m PeComanagedHostConfigurationSummary) GetThreadsPerSocket() *int {
	return m.ThreadsPerSocket
}

// GetIsHyperThreadingEnabled returns IsHyperThreadingEnabled
func (m PeComanagedHostConfigurationSummary) GetIsHyperThreadingEnabled() *bool {
	return m.IsHyperThreadingEnabled
}

// GetDefinedTags returns DefinedTags
func (m PeComanagedHostConfigurationSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetFreeformTags returns FreeformTags
func (m PeComanagedHostConfigurationSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

func (m PeComanagedHostConfigurationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PeComanagedHostConfigurationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingHostConfigurationSummaryPlatformTypeEnum(string(m.PlatformType)); !ok && m.PlatformType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlatformType: %s. Supported values are: %s.", m.PlatformType, strings.Join(GetHostConfigurationSummaryPlatformTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PeComanagedHostConfigurationSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePeComanagedHostConfigurationSummary PeComanagedHostConfigurationSummary
	s := struct {
		DiscriminatorParam string `json:"entitySource"`
		MarshalTypePeComanagedHostConfigurationSummary
	}{
		"PE_COMANAGED_HOST",
		(MarshalTypePeComanagedHostConfigurationSummary)(m),
	}

	return json.Marshal(&s)
}
