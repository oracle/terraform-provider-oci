// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.oracle.com/iaas/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MacsManagedCloudHostInsight MACS-managed OCI Compute host insight resource.
type MacsManagedCloudHostInsight struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the host insight resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The host name. The host name is unique amongst the hosts managed by the same management agent.
	HostName *string `mandatory:"true" json:"hostName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// The time the the host insight was first enabled. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Compute Instance
	ComputeId *string `mandatory:"true" json:"computeId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Management Agent
	ManagementAgentId *string `mandatory:"true" json:"managementAgentId"`

	// The user-friendly name for the host. The name does not have to be unique.
	HostDisplayName *string `mandatory:"false" json:"hostDisplayName"`

	// Ops Insights internal representation of the host type. Possible value is EXTERNAL-HOST.
	HostType *string `mandatory:"false" json:"hostType"`

	// Processor count. This is the OCPU count for Autonomous Database and CPU core count for other database types.
	ProcessorCount *int `mandatory:"false" json:"processorCount"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The time the host insight was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Platform name.
	PlatformName *string `mandatory:"false" json:"platformName"`

	// Platform version.
	PlatformVersion *string `mandatory:"false" json:"platformVersion"`

	// Platform type.
	// Supported platformType(s) for MACS-managed external host insight: [LINUX, SOLARIS, WINDOWS].
	// Supported platformType(s) for MACS-managed cloud host insight: [LINUX].
	// Supported platformType(s) for EM-managed external host insight: [LINUX, SOLARIS, SUNOS, ZLINUX, WINDOWS, AIX, HP-UX].
	PlatformType MacsManagedCloudHostInsightPlatformTypeEnum `mandatory:"false" json:"platformType,omitempty"`

	// Indicates the status of a host insight in Operations Insights
	Status ResourceStatusEnum `mandatory:"true" json:"status"`

	// The current state of the host.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

// GetId returns Id
func (m MacsManagedCloudHostInsight) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m MacsManagedCloudHostInsight) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetHostName returns HostName
func (m MacsManagedCloudHostInsight) GetHostName() *string {
	return m.HostName
}

// GetHostDisplayName returns HostDisplayName
func (m MacsManagedCloudHostInsight) GetHostDisplayName() *string {
	return m.HostDisplayName
}

// GetHostType returns HostType
func (m MacsManagedCloudHostInsight) GetHostType() *string {
	return m.HostType
}

// GetProcessorCount returns ProcessorCount
func (m MacsManagedCloudHostInsight) GetProcessorCount() *int {
	return m.ProcessorCount
}

// GetFreeformTags returns FreeformTags
func (m MacsManagedCloudHostInsight) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m MacsManagedCloudHostInsight) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m MacsManagedCloudHostInsight) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetStatus returns Status
func (m MacsManagedCloudHostInsight) GetStatus() ResourceStatusEnum {
	return m.Status
}

// GetTimeCreated returns TimeCreated
func (m MacsManagedCloudHostInsight) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m MacsManagedCloudHostInsight) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m MacsManagedCloudHostInsight) GetLifecycleState() LifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m MacsManagedCloudHostInsight) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

func (m MacsManagedCloudHostInsight) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MacsManagedCloudHostInsight) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMacsManagedCloudHostInsightPlatformTypeEnum(string(m.PlatformType)); !ok && m.PlatformType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlatformType: %s. Supported values are: %s.", m.PlatformType, strings.Join(GetMacsManagedCloudHostInsightPlatformTypeEnumStringValues(), ",")))
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
func (m MacsManagedCloudHostInsight) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeMacsManagedCloudHostInsight MacsManagedCloudHostInsight
	s := struct {
		DiscriminatorParam string `json:"entitySource"`
		MarshalTypeMacsManagedCloudHostInsight
	}{
		"MACS_MANAGED_CLOUD_HOST",
		(MarshalTypeMacsManagedCloudHostInsight)(m),
	}

	return json.Marshal(&s)
}

// MacsManagedCloudHostInsightPlatformTypeEnum Enum with underlying type: string
type MacsManagedCloudHostInsightPlatformTypeEnum string

// Set of constants representing the allowable values for MacsManagedCloudHostInsightPlatformTypeEnum
const (
	MacsManagedCloudHostInsightPlatformTypeLinux   MacsManagedCloudHostInsightPlatformTypeEnum = "LINUX"
	MacsManagedCloudHostInsightPlatformTypeSolaris MacsManagedCloudHostInsightPlatformTypeEnum = "SOLARIS"
	MacsManagedCloudHostInsightPlatformTypeSunos   MacsManagedCloudHostInsightPlatformTypeEnum = "SUNOS"
	MacsManagedCloudHostInsightPlatformTypeZlinux  MacsManagedCloudHostInsightPlatformTypeEnum = "ZLINUX"
	MacsManagedCloudHostInsightPlatformTypeWindows MacsManagedCloudHostInsightPlatformTypeEnum = "WINDOWS"
	MacsManagedCloudHostInsightPlatformTypeAix     MacsManagedCloudHostInsightPlatformTypeEnum = "AIX"
	MacsManagedCloudHostInsightPlatformTypeHpUx    MacsManagedCloudHostInsightPlatformTypeEnum = "HP_UX"
)

var mappingMacsManagedCloudHostInsightPlatformTypeEnum = map[string]MacsManagedCloudHostInsightPlatformTypeEnum{
	"LINUX":   MacsManagedCloudHostInsightPlatformTypeLinux,
	"SOLARIS": MacsManagedCloudHostInsightPlatformTypeSolaris,
	"SUNOS":   MacsManagedCloudHostInsightPlatformTypeSunos,
	"ZLINUX":  MacsManagedCloudHostInsightPlatformTypeZlinux,
	"WINDOWS": MacsManagedCloudHostInsightPlatformTypeWindows,
	"AIX":     MacsManagedCloudHostInsightPlatformTypeAix,
	"HP_UX":   MacsManagedCloudHostInsightPlatformTypeHpUx,
}

var mappingMacsManagedCloudHostInsightPlatformTypeEnumLowerCase = map[string]MacsManagedCloudHostInsightPlatformTypeEnum{
	"linux":   MacsManagedCloudHostInsightPlatformTypeLinux,
	"solaris": MacsManagedCloudHostInsightPlatformTypeSolaris,
	"sunos":   MacsManagedCloudHostInsightPlatformTypeSunos,
	"zlinux":  MacsManagedCloudHostInsightPlatformTypeZlinux,
	"windows": MacsManagedCloudHostInsightPlatformTypeWindows,
	"aix":     MacsManagedCloudHostInsightPlatformTypeAix,
	"hp_ux":   MacsManagedCloudHostInsightPlatformTypeHpUx,
}

// GetMacsManagedCloudHostInsightPlatformTypeEnumValues Enumerates the set of values for MacsManagedCloudHostInsightPlatformTypeEnum
func GetMacsManagedCloudHostInsightPlatformTypeEnumValues() []MacsManagedCloudHostInsightPlatformTypeEnum {
	values := make([]MacsManagedCloudHostInsightPlatformTypeEnum, 0)
	for _, v := range mappingMacsManagedCloudHostInsightPlatformTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMacsManagedCloudHostInsightPlatformTypeEnumStringValues Enumerates the set of values in String for MacsManagedCloudHostInsightPlatformTypeEnum
func GetMacsManagedCloudHostInsightPlatformTypeEnumStringValues() []string {
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

// GetMappingMacsManagedCloudHostInsightPlatformTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMacsManagedCloudHostInsightPlatformTypeEnum(val string) (MacsManagedCloudHostInsightPlatformTypeEnum, bool) {
	enum, ok := mappingMacsManagedCloudHostInsightPlatformTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
