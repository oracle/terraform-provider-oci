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

// AgentSummary Summary of the Agent.
type AgentSummary struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// Agent identifier, which can be renamed.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Type of Agent.
	AgentType AgentAgentTypeEnum `mandatory:"true" json:"agentType"`

	// Agent identifier.
	AgentVersion *string `mandatory:"true" json:"agentVersion"`

	// OS version.
	OsVersion *string `mandatory:"true" json:"osVersion"`

	// The time when the Agent was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time when the Agent was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Environment identifier.
	EnvironmentId *string `mandatory:"true" json:"environmentId"`

	// The current state of the Agent.
	LifecycleState AgentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A message describing the current state in more detail. For example, it can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"true" json:"lifecycleDetails"`

	// The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace/scope. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// The time when the last heartbeat of the Agent was noted. An RFC3339 formatted datetime string.
	TimeLastSyncReceived *common.SDKTime `mandatory:"false" json:"timeLastSyncReceived"`

	// Current heartbeat status of the Agent based on its timeLastSyncReceived value.
	HeartBeatStatus AgentHeartBeatStatusEnum `mandatory:"false" json:"heartBeatStatus,omitempty"`

	// The system tags associated with this resource, if any. The system tags are set by Oracle cloud infrastructure services. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m AgentSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AgentSummary) ValidateEnumValue() (bool, error) {
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
