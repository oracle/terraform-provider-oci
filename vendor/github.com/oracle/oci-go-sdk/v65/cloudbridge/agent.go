// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Bridge API
//
// API for Oracle Cloud Bridge service.
//

package cloudbridge

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Agent Description of Agent.
type Agent struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// Agent identifier, can be renamed.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Type of the Agent.
	AgentType AgentAgentTypeEnum `mandatory:"true" json:"agentType"`

	// Agent identifier.
	AgentVersion *string `mandatory:"true" json:"agentVersion"`

	// OS version.
	OsVersion *string `mandatory:"true" json:"osVersion"`

	// The time when the Agent was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Environment identifier.
	EnvironmentId *string `mandatory:"true" json:"environmentId"`

	// The current state of the Agent.
	LifecycleState AgentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace/scope. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// The time when the Agent was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The time when the last heartbeat of the Agent was noted. An RFC3339 formatted datetime string.
	TimeLastSyncReceived *common.SDKTime `mandatory:"false" json:"timeLastSyncReceived"`

	// The current heartbeat status of the Agent based on its timeLastSyncReceived value.
	HeartBeatStatus AgentHeartBeatStatusEnum `mandatory:"false" json:"heartBeatStatus,omitempty"`

	// Resource principal public key.
	AgentPubKey *string `mandatory:"false" json:"agentPubKey"`

	// The time since epoch for when the public key will expire. An RFC3339 formatted datetime string.
	TimeExpireAgentKeyInMs *common.SDKTime `mandatory:"false" json:"timeExpireAgentKeyInMs"`

	// A message describing the current state of the Agent in more detail. For example, it can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// List of plugins associated with the agent.
	PluginList []PluginSummary `mandatory:"false" json:"pluginList"`

	// The system tags associated with this resource, if any. The system tags are set by Oracle cloud infrastructure services. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m Agent) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Agent) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAgentAgentTypeEnum(string(m.AgentType)); !ok && m.AgentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AgentType: %s. Supported values are: %s.", m.AgentType, strings.Join(GetAgentAgentTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAgentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAgentLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingAgentHeartBeatStatusEnum(string(m.HeartBeatStatus)); !ok && m.HeartBeatStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for HeartBeatStatus: %s. Supported values are: %s.", m.HeartBeatStatus, strings.Join(GetAgentHeartBeatStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AgentAgentTypeEnum Enum with underlying type: string
type AgentAgentTypeEnum string

// Set of constants representing the allowable values for AgentAgentTypeEnum
const (
	AgentAgentTypeAppliance AgentAgentTypeEnum = "APPLIANCE"
)

var mappingAgentAgentTypeEnum = map[string]AgentAgentTypeEnum{
	"APPLIANCE": AgentAgentTypeAppliance,
}

var mappingAgentAgentTypeEnumLowerCase = map[string]AgentAgentTypeEnum{
	"appliance": AgentAgentTypeAppliance,
}

// GetAgentAgentTypeEnumValues Enumerates the set of values for AgentAgentTypeEnum
func GetAgentAgentTypeEnumValues() []AgentAgentTypeEnum {
	values := make([]AgentAgentTypeEnum, 0)
	for _, v := range mappingAgentAgentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAgentAgentTypeEnumStringValues Enumerates the set of values in String for AgentAgentTypeEnum
func GetAgentAgentTypeEnumStringValues() []string {
	return []string{
		"APPLIANCE",
	}
}

// GetMappingAgentAgentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAgentAgentTypeEnum(val string) (AgentAgentTypeEnum, bool) {
	enum, ok := mappingAgentAgentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AgentHeartBeatStatusEnum Enum with underlying type: string
type AgentHeartBeatStatusEnum string

// Set of constants representing the allowable values for AgentHeartBeatStatusEnum
const (
	AgentHeartBeatStatusHealthy   AgentHeartBeatStatusEnum = "HEALTHY"
	AgentHeartBeatStatusUnhealthy AgentHeartBeatStatusEnum = "UNHEALTHY"
	AgentHeartBeatStatusFailed    AgentHeartBeatStatusEnum = "FAILED"
	AgentHeartBeatStatusInactive  AgentHeartBeatStatusEnum = "INACTIVE"
)

var mappingAgentHeartBeatStatusEnum = map[string]AgentHeartBeatStatusEnum{
	"HEALTHY":   AgentHeartBeatStatusHealthy,
	"UNHEALTHY": AgentHeartBeatStatusUnhealthy,
	"FAILED":    AgentHeartBeatStatusFailed,
	"INACTIVE":  AgentHeartBeatStatusInactive,
}

var mappingAgentHeartBeatStatusEnumLowerCase = map[string]AgentHeartBeatStatusEnum{
	"healthy":   AgentHeartBeatStatusHealthy,
	"unhealthy": AgentHeartBeatStatusUnhealthy,
	"failed":    AgentHeartBeatStatusFailed,
	"inactive":  AgentHeartBeatStatusInactive,
}

// GetAgentHeartBeatStatusEnumValues Enumerates the set of values for AgentHeartBeatStatusEnum
func GetAgentHeartBeatStatusEnumValues() []AgentHeartBeatStatusEnum {
	values := make([]AgentHeartBeatStatusEnum, 0)
	for _, v := range mappingAgentHeartBeatStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetAgentHeartBeatStatusEnumStringValues Enumerates the set of values in String for AgentHeartBeatStatusEnum
func GetAgentHeartBeatStatusEnumStringValues() []string {
	return []string{
		"HEALTHY",
		"UNHEALTHY",
		"FAILED",
		"INACTIVE",
	}
}

// GetMappingAgentHeartBeatStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAgentHeartBeatStatusEnum(val string) (AgentHeartBeatStatusEnum, bool) {
	enum, ok := mappingAgentHeartBeatStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AgentLifecycleStateEnum Enum with underlying type: string
type AgentLifecycleStateEnum string

// Set of constants representing the allowable values for AgentLifecycleStateEnum
const (
	AgentLifecycleStateCreating AgentLifecycleStateEnum = "CREATING"
	AgentLifecycleStateActive   AgentLifecycleStateEnum = "ACTIVE"
	AgentLifecycleStateInactive AgentLifecycleStateEnum = "INACTIVE"
	AgentLifecycleStateDeleted  AgentLifecycleStateEnum = "DELETED"
	AgentLifecycleStateFailed   AgentLifecycleStateEnum = "FAILED"
)

var mappingAgentLifecycleStateEnum = map[string]AgentLifecycleStateEnum{
	"CREATING": AgentLifecycleStateCreating,
	"ACTIVE":   AgentLifecycleStateActive,
	"INACTIVE": AgentLifecycleStateInactive,
	"DELETED":  AgentLifecycleStateDeleted,
	"FAILED":   AgentLifecycleStateFailed,
}

var mappingAgentLifecycleStateEnumLowerCase = map[string]AgentLifecycleStateEnum{
	"creating": AgentLifecycleStateCreating,
	"active":   AgentLifecycleStateActive,
	"inactive": AgentLifecycleStateInactive,
	"deleted":  AgentLifecycleStateDeleted,
	"failed":   AgentLifecycleStateFailed,
}

// GetAgentLifecycleStateEnumValues Enumerates the set of values for AgentLifecycleStateEnum
func GetAgentLifecycleStateEnumValues() []AgentLifecycleStateEnum {
	values := make([]AgentLifecycleStateEnum, 0)
	for _, v := range mappingAgentLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAgentLifecycleStateEnumStringValues Enumerates the set of values in String for AgentLifecycleStateEnum
func GetAgentLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"DELETED",
		"FAILED",
	}
}

// GetMappingAgentLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAgentLifecycleStateEnum(val string) (AgentLifecycleStateEnum, bool) {
	enum, ok := mappingAgentLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
