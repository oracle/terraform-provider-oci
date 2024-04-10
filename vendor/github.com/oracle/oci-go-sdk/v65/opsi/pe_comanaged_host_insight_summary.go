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

// PeComanagedHostInsightSummary Summary of a Private Endpoint host insight resource.
type PeComanagedHostInsightSummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the host insight resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The host name. The host name is unique amongst the hosts managed by the same management agent.
	HostName *string `mandatory:"true" json:"hostName"`

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

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM Cluster or DB System ID, depending on which configuration the resource belongs to.
	ParentId *string `mandatory:"false" json:"parentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Infrastructure.
	RootId *string `mandatory:"false" json:"rootId"`

	// Platform type.
	// Supported platformType(s) for MACS-managed external host insight: [LINUX, SOLARIS, WINDOWS].
	// Supported platformType(s) for MACS-managed cloud host insight: [LINUX].
	// Supported platformType(s) for EM-managed external host insight: [LINUX, SOLARIS, SUNOS, ZLINUX, WINDOWS, AIX, HP-UX].
	PlatformType PeComanagedHostInsightSummaryPlatformTypeEnum `mandatory:"false" json:"platformType,omitempty"`

	// Indicates the status of a host insight in Ops Insights
	Status ResourceStatusEnum `mandatory:"false" json:"status,omitempty"`

	// The current state of the host.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

// GetId returns Id
func (m PeComanagedHostInsightSummary) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m PeComanagedHostInsightSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetHostName returns HostName
func (m PeComanagedHostInsightSummary) GetHostName() *string {
	return m.HostName
}

// GetHostDisplayName returns HostDisplayName
func (m PeComanagedHostInsightSummary) GetHostDisplayName() *string {
	return m.HostDisplayName
}

// GetHostType returns HostType
func (m PeComanagedHostInsightSummary) GetHostType() *string {
	return m.HostType
}

// GetProcessorCount returns ProcessorCount
func (m PeComanagedHostInsightSummary) GetProcessorCount() *int {
	return m.ProcessorCount
}

// GetFreeformTags returns FreeformTags
func (m PeComanagedHostInsightSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m PeComanagedHostInsightSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m PeComanagedHostInsightSummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetOpsiPrivateEndpointId returns OpsiPrivateEndpointId
func (m PeComanagedHostInsightSummary) GetOpsiPrivateEndpointId() *string {
	return m.OpsiPrivateEndpointId
}

// GetStatus returns Status
func (m PeComanagedHostInsightSummary) GetStatus() ResourceStatusEnum {
	return m.Status
}

// GetTimeCreated returns TimeCreated
func (m PeComanagedHostInsightSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m PeComanagedHostInsightSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m PeComanagedHostInsightSummary) GetLifecycleState() LifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m PeComanagedHostInsightSummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

func (m PeComanagedHostInsightSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PeComanagedHostInsightSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPeComanagedHostInsightSummaryPlatformTypeEnum(string(m.PlatformType)); !ok && m.PlatformType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlatformType: %s. Supported values are: %s.", m.PlatformType, strings.Join(GetPeComanagedHostInsightSummaryPlatformTypeEnumStringValues(), ",")))
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
func (m PeComanagedHostInsightSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePeComanagedHostInsightSummary PeComanagedHostInsightSummary
	s := struct {
		DiscriminatorParam string `json:"entitySource"`
		MarshalTypePeComanagedHostInsightSummary
	}{
		"PE_COMANAGED_HOST",
		(MarshalTypePeComanagedHostInsightSummary)(m),
	}

	return json.Marshal(&s)
}

// PeComanagedHostInsightSummaryPlatformTypeEnum Enum with underlying type: string
type PeComanagedHostInsightSummaryPlatformTypeEnum string

// Set of constants representing the allowable values for PeComanagedHostInsightSummaryPlatformTypeEnum
const (
	PeComanagedHostInsightSummaryPlatformTypeLinux   PeComanagedHostInsightSummaryPlatformTypeEnum = "LINUX"
	PeComanagedHostInsightSummaryPlatformTypeSolaris PeComanagedHostInsightSummaryPlatformTypeEnum = "SOLARIS"
	PeComanagedHostInsightSummaryPlatformTypeSunos   PeComanagedHostInsightSummaryPlatformTypeEnum = "SUNOS"
	PeComanagedHostInsightSummaryPlatformTypeZlinux  PeComanagedHostInsightSummaryPlatformTypeEnum = "ZLINUX"
	PeComanagedHostInsightSummaryPlatformTypeWindows PeComanagedHostInsightSummaryPlatformTypeEnum = "WINDOWS"
	PeComanagedHostInsightSummaryPlatformTypeAix     PeComanagedHostInsightSummaryPlatformTypeEnum = "AIX"
	PeComanagedHostInsightSummaryPlatformTypeHpUx    PeComanagedHostInsightSummaryPlatformTypeEnum = "HP_UX"
)

var mappingPeComanagedHostInsightSummaryPlatformTypeEnum = map[string]PeComanagedHostInsightSummaryPlatformTypeEnum{
	"LINUX":   PeComanagedHostInsightSummaryPlatformTypeLinux,
	"SOLARIS": PeComanagedHostInsightSummaryPlatformTypeSolaris,
	"SUNOS":   PeComanagedHostInsightSummaryPlatformTypeSunos,
	"ZLINUX":  PeComanagedHostInsightSummaryPlatformTypeZlinux,
	"WINDOWS": PeComanagedHostInsightSummaryPlatformTypeWindows,
	"AIX":     PeComanagedHostInsightSummaryPlatformTypeAix,
	"HP_UX":   PeComanagedHostInsightSummaryPlatformTypeHpUx,
}

var mappingPeComanagedHostInsightSummaryPlatformTypeEnumLowerCase = map[string]PeComanagedHostInsightSummaryPlatformTypeEnum{
	"linux":   PeComanagedHostInsightSummaryPlatformTypeLinux,
	"solaris": PeComanagedHostInsightSummaryPlatformTypeSolaris,
	"sunos":   PeComanagedHostInsightSummaryPlatformTypeSunos,
	"zlinux":  PeComanagedHostInsightSummaryPlatformTypeZlinux,
	"windows": PeComanagedHostInsightSummaryPlatformTypeWindows,
	"aix":     PeComanagedHostInsightSummaryPlatformTypeAix,
	"hp_ux":   PeComanagedHostInsightSummaryPlatformTypeHpUx,
}

// GetPeComanagedHostInsightSummaryPlatformTypeEnumValues Enumerates the set of values for PeComanagedHostInsightSummaryPlatformTypeEnum
func GetPeComanagedHostInsightSummaryPlatformTypeEnumValues() []PeComanagedHostInsightSummaryPlatformTypeEnum {
	values := make([]PeComanagedHostInsightSummaryPlatformTypeEnum, 0)
	for _, v := range mappingPeComanagedHostInsightSummaryPlatformTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPeComanagedHostInsightSummaryPlatformTypeEnumStringValues Enumerates the set of values in String for PeComanagedHostInsightSummaryPlatformTypeEnum
func GetPeComanagedHostInsightSummaryPlatformTypeEnumStringValues() []string {
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

// GetMappingPeComanagedHostInsightSummaryPlatformTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPeComanagedHostInsightSummaryPlatformTypeEnum(val string) (PeComanagedHostInsightSummaryPlatformTypeEnum, bool) {
	enum, ok := mappingPeComanagedHostInsightSummaryPlatformTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
