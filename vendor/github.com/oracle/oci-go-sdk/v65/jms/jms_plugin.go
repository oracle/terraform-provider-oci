// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Fleets API
//
// The APIs for the Fleet Management (https://docs.oracle.com/en-us/iaas/jms/doc/fleet-management.html) feature of Java Management Service to monitor and manage the usage of Java in your enterprise. Use these APIs to manage fleets, configure managed instances to report to fleets, and gain insights into the Java workloads running on these instances by carrying out basic and advanced features.
//

package jms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// JmsPlugin Information about a JmsPlugin that has been registered.
type JmsPlugin struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) to identify this JmsPlugin.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Management Agent (OMA) or the Oracle Cloud Agent (OCA) instance where the JMS plugin is deployed.
	AgentId *string `mandatory:"true" json:"agentId"`

	// The agent type.
	AgentType AgentTypeEnum `mandatory:"true" json:"agentType"`

	// The lifecycle state.
	LifecycleState JmsPluginLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The availability status.
	AvailabilityStatus JmsPluginAvailabilityStatusEnum `mandatory:"true" json:"availabilityStatus"`

	// The date and time the plugin was registered.
	TimeRegistered *common.SDKTime `mandatory:"true" json:"timeRegistered"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the fleet.
	FleetId *string `mandatory:"false" json:"fleetId"`

	// The OMA/OCA agent's compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The hostname of the agent.
	Hostname *string `mandatory:"false" json:"hostname"`

	// The operating system family for the plugin.
	OsFamily OsFamilyEnum `mandatory:"false" json:"osFamily,omitempty"`

	// The architecture of the operating system of the plugin.
	OsArchitecture *string `mandatory:"false" json:"osArchitecture"`

	// The distribution of the operating system of the plugin.
	OsDistribution *string `mandatory:"false" json:"osDistribution"`

	// The version of the plugin.
	PluginVersion *string `mandatory:"false" json:"pluginVersion"`

	// The date and time the resource was _last_ reported to JMS.
	// This is potentially _after_ the specified time period provided by the filters.
	// For example, a resource can be last reported to JMS before the start of a specified time period,
	// if it is also reported during the time period.
	TimeLastSeen *common.SDKTime `mandatory:"false" json:"timeLastSeen"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`. (See Understanding Free-form Tags (https://docs.cloud.oracle.com/Content/Tagging/Tasks/managingtagsandtagnamespaces.htm)).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`. (See Managing Tags and Tag Namespaces (https://docs.cloud.oracle.com/Content/Tagging/Concepts/understandingfreeformtags.htm).)
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// System tags can be viewed by users, but can only be created by the system.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m JmsPlugin) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JmsPlugin) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAgentTypeEnum(string(m.AgentType)); !ok && m.AgentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AgentType: %s. Supported values are: %s.", m.AgentType, strings.Join(GetAgentTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingJmsPluginLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetJmsPluginLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingJmsPluginAvailabilityStatusEnum(string(m.AvailabilityStatus)); !ok && m.AvailabilityStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AvailabilityStatus: %s. Supported values are: %s.", m.AvailabilityStatus, strings.Join(GetJmsPluginAvailabilityStatusEnumStringValues(), ",")))
	}

	if _, ok := GetMappingOsFamilyEnum(string(m.OsFamily)); !ok && m.OsFamily != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsFamily: %s. Supported values are: %s.", m.OsFamily, strings.Join(GetOsFamilyEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
