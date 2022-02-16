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

// EmManagedExternalHostInsightSummary Summary of an EM-managed external host insight resource.
type EmManagedExternalHostInsightSummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the host insight resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The host name. The host name is unique amongst the hosts managed by the same management agent.
	HostName *string `mandatory:"true" json:"hostName"`

	// Enterprise Manager Unique Identifier
	EnterpriseManagerIdentifier *string `mandatory:"true" json:"enterpriseManagerIdentifier"`

	// Enterprise Manager Entity Name
	EnterpriseManagerEntityName *string `mandatory:"true" json:"enterpriseManagerEntityName"`

	// Enterprise Manager Entity Type
	EnterpriseManagerEntityType *string `mandatory:"true" json:"enterpriseManagerEntityType"`

	// Enterprise Manager Entity Unique Identifier
	EnterpriseManagerEntityIdentifier *string `mandatory:"true" json:"enterpriseManagerEntityIdentifier"`

	// OPSI Enterprise Manager Bridge OCID
	EnterpriseManagerBridgeId *string `mandatory:"true" json:"enterpriseManagerBridgeId"`

	// The user-friendly name for the host. The name does not have to be unique.
	HostDisplayName *string `mandatory:"false" json:"hostDisplayName"`

	// Operations Insights internal representation of the host type. Possible value is EXTERNAL-HOST.
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

	// The time the the host insight was first enabled. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the host insight was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Enterprise Manager Entity Display Name
	EnterpriseManagerEntityDisplayName *string `mandatory:"false" json:"enterpriseManagerEntityDisplayName"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata insight.
	ExadataInsightId *string `mandatory:"false" json:"exadataInsightId"`

	// Platform type.
	// Supported platformType(s) for MACS-managed external host insight: [LINUX].
	// Supported platformType(s) for EM-managed external host insight: [LINUX, SOLARIS, SUNOS].
	PlatformType EmManagedExternalHostInsightSummaryPlatformTypeEnum `mandatory:"false" json:"platformType,omitempty"`

	// Indicates the status of a host insight in Operations Insights
	Status ResourceStatusEnum `mandatory:"false" json:"status,omitempty"`

	// The current state of the host.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

//GetId returns Id
func (m EmManagedExternalHostInsightSummary) GetId() *string {
	return m.Id
}

//GetCompartmentId returns CompartmentId
func (m EmManagedExternalHostInsightSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetHostName returns HostName
func (m EmManagedExternalHostInsightSummary) GetHostName() *string {
	return m.HostName
}

//GetHostDisplayName returns HostDisplayName
func (m EmManagedExternalHostInsightSummary) GetHostDisplayName() *string {
	return m.HostDisplayName
}

//GetHostType returns HostType
func (m EmManagedExternalHostInsightSummary) GetHostType() *string {
	return m.HostType
}

//GetProcessorCount returns ProcessorCount
func (m EmManagedExternalHostInsightSummary) GetProcessorCount() *int {
	return m.ProcessorCount
}

//GetFreeformTags returns FreeformTags
func (m EmManagedExternalHostInsightSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m EmManagedExternalHostInsightSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m EmManagedExternalHostInsightSummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

//GetStatus returns Status
func (m EmManagedExternalHostInsightSummary) GetStatus() ResourceStatusEnum {
	return m.Status
}

//GetTimeCreated returns TimeCreated
func (m EmManagedExternalHostInsightSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m EmManagedExternalHostInsightSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleState returns LifecycleState
func (m EmManagedExternalHostInsightSummary) GetLifecycleState() LifecycleStateEnum {
	return m.LifecycleState
}

//GetLifecycleDetails returns LifecycleDetails
func (m EmManagedExternalHostInsightSummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

func (m EmManagedExternalHostInsightSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EmManagedExternalHostInsightSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEmManagedExternalHostInsightSummaryPlatformTypeEnum(string(m.PlatformType)); !ok && m.PlatformType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlatformType: %s. Supported values are: %s.", m.PlatformType, strings.Join(GetEmManagedExternalHostInsightSummaryPlatformTypeEnumStringValues(), ",")))
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
func (m EmManagedExternalHostInsightSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeEmManagedExternalHostInsightSummary EmManagedExternalHostInsightSummary
	s := struct {
		DiscriminatorParam string `json:"entitySource"`
		MarshalTypeEmManagedExternalHostInsightSummary
	}{
		"EM_MANAGED_EXTERNAL_HOST",
		(MarshalTypeEmManagedExternalHostInsightSummary)(m),
	}

	return json.Marshal(&s)
}

// EmManagedExternalHostInsightSummaryPlatformTypeEnum Enum with underlying type: string
type EmManagedExternalHostInsightSummaryPlatformTypeEnum string

// Set of constants representing the allowable values for EmManagedExternalHostInsightSummaryPlatformTypeEnum
const (
	EmManagedExternalHostInsightSummaryPlatformTypeLinux   EmManagedExternalHostInsightSummaryPlatformTypeEnum = "LINUX"
	EmManagedExternalHostInsightSummaryPlatformTypeSolaris EmManagedExternalHostInsightSummaryPlatformTypeEnum = "SOLARIS"
	EmManagedExternalHostInsightSummaryPlatformTypeSunos   EmManagedExternalHostInsightSummaryPlatformTypeEnum = "SUNOS"
)

var mappingEmManagedExternalHostInsightSummaryPlatformTypeEnum = map[string]EmManagedExternalHostInsightSummaryPlatformTypeEnum{
	"LINUX":   EmManagedExternalHostInsightSummaryPlatformTypeLinux,
	"SOLARIS": EmManagedExternalHostInsightSummaryPlatformTypeSolaris,
	"SUNOS":   EmManagedExternalHostInsightSummaryPlatformTypeSunos,
}

// GetEmManagedExternalHostInsightSummaryPlatformTypeEnumValues Enumerates the set of values for EmManagedExternalHostInsightSummaryPlatformTypeEnum
func GetEmManagedExternalHostInsightSummaryPlatformTypeEnumValues() []EmManagedExternalHostInsightSummaryPlatformTypeEnum {
	values := make([]EmManagedExternalHostInsightSummaryPlatformTypeEnum, 0)
	for _, v := range mappingEmManagedExternalHostInsightSummaryPlatformTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetEmManagedExternalHostInsightSummaryPlatformTypeEnumStringValues Enumerates the set of values in String for EmManagedExternalHostInsightSummaryPlatformTypeEnum
func GetEmManagedExternalHostInsightSummaryPlatformTypeEnumStringValues() []string {
	return []string{
		"LINUX",
		"SOLARIS",
		"SUNOS",
	}
}

// GetMappingEmManagedExternalHostInsightSummaryPlatformTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEmManagedExternalHostInsightSummaryPlatformTypeEnum(val string) (EmManagedExternalHostInsightSummaryPlatformTypeEnum, bool) {
	mappingEmManagedExternalHostInsightSummaryPlatformTypeEnumIgnoreCase := make(map[string]EmManagedExternalHostInsightSummaryPlatformTypeEnum)
	for k, v := range mappingEmManagedExternalHostInsightSummaryPlatformTypeEnum {
		mappingEmManagedExternalHostInsightSummaryPlatformTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingEmManagedExternalHostInsightSummaryPlatformTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
