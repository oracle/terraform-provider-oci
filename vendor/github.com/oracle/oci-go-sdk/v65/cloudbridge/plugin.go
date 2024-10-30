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

// Plugin Description of plugin
type Plugin struct {

	// Plugin identifier, which can be renamed.
	Name *string `mandatory:"true" json:"name"`

	// Agent identifier.
	AgentId *string `mandatory:"true" json:"agentId"`

	// Plugin version.
	PluginVersion *string `mandatory:"true" json:"pluginVersion"`

	// The time when the Agent was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the plugin.
	LifecycleState PluginLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// State to which the customer wants the plugin to move to.
	DesiredState PluginDesiredStateEnum `mandatory:"false" json:"desiredState,omitempty"`

	// The time when the Agent was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, it can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace/scope. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The system tags associated with this resource, if any. The system tags are set by Oracle cloud infrastructure services. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m Plugin) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Plugin) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPluginLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPluginLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingPluginDesiredStateEnum(string(m.DesiredState)); !ok && m.DesiredState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DesiredState: %s. Supported values are: %s.", m.DesiredState, strings.Join(GetPluginDesiredStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PluginDesiredStateEnum Enum with underlying type: string
type PluginDesiredStateEnum string

// Set of constants representing the allowable values for PluginDesiredStateEnum
const (
	PluginDesiredStateEnabled  PluginDesiredStateEnum = "ENABLED"
	PluginDesiredStateDisabled PluginDesiredStateEnum = "DISABLED"
)

var mappingPluginDesiredStateEnum = map[string]PluginDesiredStateEnum{
	"ENABLED":  PluginDesiredStateEnabled,
	"DISABLED": PluginDesiredStateDisabled,
}

var mappingPluginDesiredStateEnumLowerCase = map[string]PluginDesiredStateEnum{
	"enabled":  PluginDesiredStateEnabled,
	"disabled": PluginDesiredStateDisabled,
}

// GetPluginDesiredStateEnumValues Enumerates the set of values for PluginDesiredStateEnum
func GetPluginDesiredStateEnumValues() []PluginDesiredStateEnum {
	values := make([]PluginDesiredStateEnum, 0)
	for _, v := range mappingPluginDesiredStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPluginDesiredStateEnumStringValues Enumerates the set of values in String for PluginDesiredStateEnum
func GetPluginDesiredStateEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingPluginDesiredStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPluginDesiredStateEnum(val string) (PluginDesiredStateEnum, bool) {
	enum, ok := mappingPluginDesiredStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// PluginLifecycleStateEnum Enum with underlying type: string
type PluginLifecycleStateEnum string

// Set of constants representing the allowable values for PluginLifecycleStateEnum
const (
	PluginLifecycleStateUpdating       PluginLifecycleStateEnum = "UPDATING"
	PluginLifecycleStateActive         PluginLifecycleStateEnum = "ACTIVE"
	PluginLifecycleStateInactive       PluginLifecycleStateEnum = "INACTIVE"
	PluginLifecycleStateNeedsAttention PluginLifecycleStateEnum = "NEEDS_ATTENTION"
	PluginLifecycleStateDeleted        PluginLifecycleStateEnum = "DELETED"
)

var mappingPluginLifecycleStateEnum = map[string]PluginLifecycleStateEnum{
	"UPDATING":        PluginLifecycleStateUpdating,
	"ACTIVE":          PluginLifecycleStateActive,
	"INACTIVE":        PluginLifecycleStateInactive,
	"NEEDS_ATTENTION": PluginLifecycleStateNeedsAttention,
	"DELETED":         PluginLifecycleStateDeleted,
}

var mappingPluginLifecycleStateEnumLowerCase = map[string]PluginLifecycleStateEnum{
	"updating":        PluginLifecycleStateUpdating,
	"active":          PluginLifecycleStateActive,
	"inactive":        PluginLifecycleStateInactive,
	"needs_attention": PluginLifecycleStateNeedsAttention,
	"deleted":         PluginLifecycleStateDeleted,
}

// GetPluginLifecycleStateEnumValues Enumerates the set of values for PluginLifecycleStateEnum
func GetPluginLifecycleStateEnumValues() []PluginLifecycleStateEnum {
	values := make([]PluginLifecycleStateEnum, 0)
	for _, v := range mappingPluginLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPluginLifecycleStateEnumStringValues Enumerates the set of values in String for PluginLifecycleStateEnum
func GetPluginLifecycleStateEnumStringValues() []string {
	return []string{
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"NEEDS_ATTENTION",
		"DELETED",
	}
}

// GetMappingPluginLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPluginLifecycleStateEnum(val string) (PluginLifecycleStateEnum, bool) {
	enum, ok := mappingPluginLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
