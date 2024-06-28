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

// MacsManagedCloudHostInsightSummary Summary of a MACS-managed cloud host insight resource.
type MacsManagedCloudHostInsightSummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the host insight resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The host name. The host name is unique amongst the hosts managed by the same management agent.
	HostName *string `mandatory:"true" json:"hostName"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Compute Instance
	ComputeId *string `mandatory:"true" json:"computeId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Management Agent
	ManagementAgentId *string `mandatory:"true" json:"managementAgentId"`

	// The user-friendly name for the host. The name does not have to be unique.
	HostDisplayName *string `mandatory:"false" json:"hostDisplayName"`

	// Ops Insights internal representation of the host type. Possible value is EXTERNAL-HOST.
	HostType *string `mandatory:"false" json:"hostType"`

	// Processor count. This is the OCPU count for Autonomous Database and CPU core count for other database types.
	ProcessorCount *int `mandatory:"false" json:"processorCount"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the OPSI private endpoint
	OpsiPrivateEndpointId *string `mandatory:"false" json:"opsiPrivateEndpointId"`

	// The time the the host insight was first enabled. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the host insight was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Platform type.
	// Supported platformType(s) for MACS-managed external host insight: [LINUX, SOLARIS, WINDOWS].
	// Supported platformType(s) for MACS-managed cloud host insight: [LINUX].
	// Supported platformType(s) for EM-managed external host insight: [LINUX, SOLARIS, SUNOS, ZLINUX, WINDOWS, AIX, HP-UX].
	PlatformType MacsManagedCloudHostInsightSummaryPlatformTypeEnum `mandatory:"false" json:"platformType,omitempty"`

	// Indicates the status of a host insight in Ops Insights
	Status ResourceStatusEnum `mandatory:"false" json:"status,omitempty"`

	// The current state of the host.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

// GetId returns Id
func (m MacsManagedCloudHostInsightSummary) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m MacsManagedCloudHostInsightSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetHostName returns HostName
func (m MacsManagedCloudHostInsightSummary) GetHostName() *string {
	return m.HostName
}

// GetHostDisplayName returns HostDisplayName
func (m MacsManagedCloudHostInsightSummary) GetHostDisplayName() *string {
	return m.HostDisplayName
}

// GetHostType returns HostType
func (m MacsManagedCloudHostInsightSummary) GetHostType() *string {
	return m.HostType
}

// GetProcessorCount returns ProcessorCount
func (m MacsManagedCloudHostInsightSummary) GetProcessorCount() *int {
	return m.ProcessorCount
}

// GetFreeformTags returns FreeformTags
func (m MacsManagedCloudHostInsightSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m MacsManagedCloudHostInsightSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m MacsManagedCloudHostInsightSummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetOpsiPrivateEndpointId returns OpsiPrivateEndpointId
func (m MacsManagedCloudHostInsightSummary) GetOpsiPrivateEndpointId() *string {
	return m.OpsiPrivateEndpointId
}

// GetStatus returns Status
func (m MacsManagedCloudHostInsightSummary) GetStatus() ResourceStatusEnum {
	return m.Status
}

// GetTimeCreated returns TimeCreated
func (m MacsManagedCloudHostInsightSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m MacsManagedCloudHostInsightSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m MacsManagedCloudHostInsightSummary) GetLifecycleState() LifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m MacsManagedCloudHostInsightSummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

func (m MacsManagedCloudHostInsightSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MacsManagedCloudHostInsightSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMacsManagedCloudHostInsightSummaryPlatformTypeEnum(string(m.PlatformType)); !ok && m.PlatformType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlatformType: %s. Supported values are: %s.", m.PlatformType, strings.Join(GetMacsManagedCloudHostInsightSummaryPlatformTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingResourceStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetResourceStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m MacsManagedCloudHostInsightSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeMacsManagedCloudHostInsightSummary MacsManagedCloudHostInsightSummary
	s := struct {
		DiscriminatorParam string `json:"entitySource"`
		MarshalTypeMacsManagedCloudHostInsightSummary
	}{
		"MACS_MANAGED_CLOUD_HOST",
		(MarshalTypeMacsManagedCloudHostInsightSummary)(m),
	}

	return json.Marshal(&s)
}

// MacsManagedCloudHostInsightSummaryPlatformTypeEnum Enum with underlying type: string
type MacsManagedCloudHostInsightSummaryPlatformTypeEnum string

// Set of constants representing the allowable values for MacsManagedCloudHostInsightSummaryPlatformTypeEnum
const (
	MacsManagedCloudHostInsightSummaryPlatformTypeLinux   MacsManagedCloudHostInsightSummaryPlatformTypeEnum = "LINUX"
	MacsManagedCloudHostInsightSummaryPlatformTypeSolaris MacsManagedCloudHostInsightSummaryPlatformTypeEnum = "SOLARIS"
	MacsManagedCloudHostInsightSummaryPlatformTypeSunos   MacsManagedCloudHostInsightSummaryPlatformTypeEnum = "SUNOS"
	MacsManagedCloudHostInsightSummaryPlatformTypeZlinux  MacsManagedCloudHostInsightSummaryPlatformTypeEnum = "ZLINUX"
	MacsManagedCloudHostInsightSummaryPlatformTypeWindows MacsManagedCloudHostInsightSummaryPlatformTypeEnum = "WINDOWS"
	MacsManagedCloudHostInsightSummaryPlatformTypeAix     MacsManagedCloudHostInsightSummaryPlatformTypeEnum = "AIX"
	MacsManagedCloudHostInsightSummaryPlatformTypeHpUx    MacsManagedCloudHostInsightSummaryPlatformTypeEnum = "HP_UX"
)

var mappingMacsManagedCloudHostInsightSummaryPlatformTypeEnum = map[string]MacsManagedCloudHostInsightSummaryPlatformTypeEnum{
	"LINUX":   MacsManagedCloudHostInsightSummaryPlatformTypeLinux,
	"SOLARIS": MacsManagedCloudHostInsightSummaryPlatformTypeSolaris,
	"SUNOS":   MacsManagedCloudHostInsightSummaryPlatformTypeSunos,
	"ZLINUX":  MacsManagedCloudHostInsightSummaryPlatformTypeZlinux,
	"WINDOWS": MacsManagedCloudHostInsightSummaryPlatformTypeWindows,
	"AIX":     MacsManagedCloudHostInsightSummaryPlatformTypeAix,
	"HP_UX":   MacsManagedCloudHostInsightSummaryPlatformTypeHpUx,
}

var mappingMacsManagedCloudHostInsightSummaryPlatformTypeEnumLowerCase = map[string]MacsManagedCloudHostInsightSummaryPlatformTypeEnum{
	"linux":   MacsManagedCloudHostInsightSummaryPlatformTypeLinux,
	"solaris": MacsManagedCloudHostInsightSummaryPlatformTypeSolaris,
	"sunos":   MacsManagedCloudHostInsightSummaryPlatformTypeSunos,
	"zlinux":  MacsManagedCloudHostInsightSummaryPlatformTypeZlinux,
	"windows": MacsManagedCloudHostInsightSummaryPlatformTypeWindows,
	"aix":     MacsManagedCloudHostInsightSummaryPlatformTypeAix,
	"hp_ux":   MacsManagedCloudHostInsightSummaryPlatformTypeHpUx,
}

// GetMacsManagedCloudHostInsightSummaryPlatformTypeEnumValues Enumerates the set of values for MacsManagedCloudHostInsightSummaryPlatformTypeEnum
func GetMacsManagedCloudHostInsightSummaryPlatformTypeEnumValues() []MacsManagedCloudHostInsightSummaryPlatformTypeEnum {
	values := make([]MacsManagedCloudHostInsightSummaryPlatformTypeEnum, 0)
	for _, v := range mappingMacsManagedCloudHostInsightSummaryPlatformTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMacsManagedCloudHostInsightSummaryPlatformTypeEnumStringValues Enumerates the set of values in String for MacsManagedCloudHostInsightSummaryPlatformTypeEnum
func GetMacsManagedCloudHostInsightSummaryPlatformTypeEnumStringValues() []string {
	return []string{
		"LINUX",
		"SOLARIS",
		"SUNOS",
		"ZLINUX",
		"WINDOWS",
		"AIX",
		"HP_UX",
	}
}

// GetMappingMacsManagedCloudHostInsightSummaryPlatformTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMacsManagedCloudHostInsightSummaryPlatformTypeEnum(val string) (MacsManagedCloudHostInsightSummaryPlatformTypeEnum, bool) {
	enum, ok := mappingMacsManagedCloudHostInsightSummaryPlatformTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
