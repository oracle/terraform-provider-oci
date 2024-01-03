// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Agent API
//
// API for the Oracle Cloud Agent software running on compute instances. Oracle Cloud Agent
// is a lightweight process that monitors and manages compute instances.
//

package computeinstanceagent

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InstanceAgentPlugin The agent plugin
type InstanceAgentPlugin struct {

	// The plugin name
	Name *string `mandatory:"true" json:"name"`

	// The plugin status Specified the plugin state on the instance * `RUNNING` - The plugin is in running state * `STOPPED` - The plugin is in stopped state * `NOT_SUPPORTED` - The plugin is not supported on this platform * `INVALID` - The plugin state is not recognizable by the service
	Status InstanceAgentPluginStatusEnum `mandatory:"true" json:"status"`

	// The last update time of the plugin in UTC
	TimeLastUpdatedUtc *common.SDKTime `mandatory:"true" json:"timeLastUpdatedUtc"`

	// The optional message from the agent plugin
	Message *string `mandatory:"false" json:"message"`
}

func (m InstanceAgentPlugin) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstanceAgentPlugin) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInstanceAgentPluginStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetInstanceAgentPluginStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InstanceAgentPluginStatusEnum Enum with underlying type: string
type InstanceAgentPluginStatusEnum string

// Set of constants representing the allowable values for InstanceAgentPluginStatusEnum
const (
	InstanceAgentPluginStatusRunning      InstanceAgentPluginStatusEnum = "RUNNING"
	InstanceAgentPluginStatusStopped      InstanceAgentPluginStatusEnum = "STOPPED"
	InstanceAgentPluginStatusNotSupported InstanceAgentPluginStatusEnum = "NOT_SUPPORTED"
	InstanceAgentPluginStatusInvalid      InstanceAgentPluginStatusEnum = "INVALID"
)

var mappingInstanceAgentPluginStatusEnum = map[string]InstanceAgentPluginStatusEnum{
	"RUNNING":       InstanceAgentPluginStatusRunning,
	"STOPPED":       InstanceAgentPluginStatusStopped,
	"NOT_SUPPORTED": InstanceAgentPluginStatusNotSupported,
	"INVALID":       InstanceAgentPluginStatusInvalid,
}

var mappingInstanceAgentPluginStatusEnumLowerCase = map[string]InstanceAgentPluginStatusEnum{
	"running":       InstanceAgentPluginStatusRunning,
	"stopped":       InstanceAgentPluginStatusStopped,
	"not_supported": InstanceAgentPluginStatusNotSupported,
	"invalid":       InstanceAgentPluginStatusInvalid,
}

// GetInstanceAgentPluginStatusEnumValues Enumerates the set of values for InstanceAgentPluginStatusEnum
func GetInstanceAgentPluginStatusEnumValues() []InstanceAgentPluginStatusEnum {
	values := make([]InstanceAgentPluginStatusEnum, 0)
	for _, v := range mappingInstanceAgentPluginStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetInstanceAgentPluginStatusEnumStringValues Enumerates the set of values in String for InstanceAgentPluginStatusEnum
func GetInstanceAgentPluginStatusEnumStringValues() []string {
	return []string{
		"RUNNING",
		"STOPPED",
		"NOT_SUPPORTED",
		"INVALID",
	}
}

// GetMappingInstanceAgentPluginStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInstanceAgentPluginStatusEnum(val string) (InstanceAgentPluginStatusEnum, bool) {
	enum, ok := mappingInstanceAgentPluginStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
