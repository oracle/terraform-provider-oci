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

// MacsManagedExternalHostInsightSummary Summary of a MACS-managed external host insight resource.
type MacsManagedExternalHostInsightSummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the host insight resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The host name. The host name is unique amongst the hosts managed by the same management agent.
	HostName *string `mandatory:"true" json:"hostName"`

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
	PlatformType MacsManagedExternalHostInsightSummaryPlatformTypeEnum `mandatory:"false" json:"platformType,omitempty"`

	// Indicates the status of a host insight in Ops Insights
	Status ResourceStatusEnum `mandatory:"false" json:"status,omitempty"`

	// The current state of the host.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

// GetId returns Id
func (m MacsManagedExternalHostInsightSummary) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m MacsManagedExternalHostInsightSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetHostName returns HostName
func (m MacsManagedExternalHostInsightSummary) GetHostName() *string {
	return m.HostName
}

// GetHostDisplayName returns HostDisplayName
func (m MacsManagedExternalHostInsightSummary) GetHostDisplayName() *string {
	return m.HostDisplayName
}

// GetHostType returns HostType
func (m MacsManagedExternalHostInsightSummary) GetHostType() *string {
	return m.HostType
}

// GetProcessorCount returns ProcessorCount
func (m MacsManagedExternalHostInsightSummary) GetProcessorCount() *int {
	return m.ProcessorCount
}

// GetFreeformTags returns FreeformTags
func (m MacsManagedExternalHostInsightSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m MacsManagedExternalHostInsightSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m MacsManagedExternalHostInsightSummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetOpsiPrivateEndpointId returns OpsiPrivateEndpointId
func (m MacsManagedExternalHostInsightSummary) GetOpsiPrivateEndpointId() *string {
	return m.OpsiPrivateEndpointId
}

// GetStatus returns Status
func (m MacsManagedExternalHostInsightSummary) GetStatus() ResourceStatusEnum {
	return m.Status
}

// GetTimeCreated returns TimeCreated
func (m MacsManagedExternalHostInsightSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m MacsManagedExternalHostInsightSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m MacsManagedExternalHostInsightSummary) GetLifecycleState() LifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m MacsManagedExternalHostInsightSummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

func (m MacsManagedExternalHostInsightSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MacsManagedExternalHostInsightSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMacsManagedExternalHostInsightSummaryPlatformTypeEnum(string(m.PlatformType)); !ok && m.PlatformType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlatformType: %s. Supported values are: %s.", m.PlatformType, strings.Join(GetMacsManagedExternalHostInsightSummaryPlatformTypeEnumStringValues(), ",")))
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
func (m MacsManagedExternalHostInsightSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeMacsManagedExternalHostInsightSummary MacsManagedExternalHostInsightSummary
	s := struct {
		DiscriminatorParam string `json:"entitySource"`
		MarshalTypeMacsManagedExternalHostInsightSummary
	}{
		"MACS_MANAGED_EXTERNAL_HOST",
		(MarshalTypeMacsManagedExternalHostInsightSummary)(m),
	}

	return json.Marshal(&s)
}

// MacsManagedExternalHostInsightSummaryPlatformTypeEnum Enum with underlying type: string
type MacsManagedExternalHostInsightSummaryPlatformTypeEnum string

// Set of constants representing the allowable values for MacsManagedExternalHostInsightSummaryPlatformTypeEnum
const (
	MacsManagedExternalHostInsightSummaryPlatformTypeLinux   MacsManagedExternalHostInsightSummaryPlatformTypeEnum = "LINUX"
	MacsManagedExternalHostInsightSummaryPlatformTypeSolaris MacsManagedExternalHostInsightSummaryPlatformTypeEnum = "SOLARIS"
	MacsManagedExternalHostInsightSummaryPlatformTypeSunos   MacsManagedExternalHostInsightSummaryPlatformTypeEnum = "SUNOS"
	MacsManagedExternalHostInsightSummaryPlatformTypeZlinux  MacsManagedExternalHostInsightSummaryPlatformTypeEnum = "ZLINUX"
	MacsManagedExternalHostInsightSummaryPlatformTypeWindows MacsManagedExternalHostInsightSummaryPlatformTypeEnum = "WINDOWS"
	MacsManagedExternalHostInsightSummaryPlatformTypeAix     MacsManagedExternalHostInsightSummaryPlatformTypeEnum = "AIX"
	MacsManagedExternalHostInsightSummaryPlatformTypeHpUx    MacsManagedExternalHostInsightSummaryPlatformTypeEnum = "HP_UX"
)

var mappingMacsManagedExternalHostInsightSummaryPlatformTypeEnum = map[string]MacsManagedExternalHostInsightSummaryPlatformTypeEnum{
	"LINUX":   MacsManagedExternalHostInsightSummaryPlatformTypeLinux,
	"SOLARIS": MacsManagedExternalHostInsightSummaryPlatformTypeSolaris,
	"SUNOS":   MacsManagedExternalHostInsightSummaryPlatformTypeSunos,
	"ZLINUX":  MacsManagedExternalHostInsightSummaryPlatformTypeZlinux,
	"WINDOWS": MacsManagedExternalHostInsightSummaryPlatformTypeWindows,
	"AIX":     MacsManagedExternalHostInsightSummaryPlatformTypeAix,
	"HP_UX":   MacsManagedExternalHostInsightSummaryPlatformTypeHpUx,
}

var mappingMacsManagedExternalHostInsightSummaryPlatformTypeEnumLowerCase = map[string]MacsManagedExternalHostInsightSummaryPlatformTypeEnum{
	"linux":   MacsManagedExternalHostInsightSummaryPlatformTypeLinux,
	"solaris": MacsManagedExternalHostInsightSummaryPlatformTypeSolaris,
	"sunos":   MacsManagedExternalHostInsightSummaryPlatformTypeSunos,
	"zlinux":  MacsManagedExternalHostInsightSummaryPlatformTypeZlinux,
	"windows": MacsManagedExternalHostInsightSummaryPlatformTypeWindows,
	"aix":     MacsManagedExternalHostInsightSummaryPlatformTypeAix,
	"hp_ux":   MacsManagedExternalHostInsightSummaryPlatformTypeHpUx,
}

// GetMacsManagedExternalHostInsightSummaryPlatformTypeEnumValues Enumerates the set of values for MacsManagedExternalHostInsightSummaryPlatformTypeEnum
func GetMacsManagedExternalHostInsightSummaryPlatformTypeEnumValues() []MacsManagedExternalHostInsightSummaryPlatformTypeEnum {
	values := make([]MacsManagedExternalHostInsightSummaryPlatformTypeEnum, 0)
	for _, v := range mappingMacsManagedExternalHostInsightSummaryPlatformTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMacsManagedExternalHostInsightSummaryPlatformTypeEnumStringValues Enumerates the set of values in String for MacsManagedExternalHostInsightSummaryPlatformTypeEnum
func GetMacsManagedExternalHostInsightSummaryPlatformTypeEnumStringValues() []string {
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

// GetMappingMacsManagedExternalHostInsightSummaryPlatformTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMacsManagedExternalHostInsightSummaryPlatformTypeEnum(val string) (MacsManagedExternalHostInsightSummaryPlatformTypeEnum, bool) {
	enum, ok := mappingMacsManagedExternalHostInsightSummaryPlatformTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
