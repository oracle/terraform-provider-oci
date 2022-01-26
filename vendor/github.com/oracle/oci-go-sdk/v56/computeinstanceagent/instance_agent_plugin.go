// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Agent API
//
// API for the Oracle Cloud Agent software running on compute instances. Oracle Cloud Agent
// is a lightweight process that monitors and manages compute instances.
//

package computeinstanceagent

import (
	"github.com/oracle/oci-go-sdk/v56/common"
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

// InstanceAgentPluginStatusEnum Enum with underlying type: string
type InstanceAgentPluginStatusEnum string

// Set of constants representing the allowable values for InstanceAgentPluginStatusEnum
const (
	InstanceAgentPluginStatusRunning      InstanceAgentPluginStatusEnum = "RUNNING"
	InstanceAgentPluginStatusStopped      InstanceAgentPluginStatusEnum = "STOPPED"
	InstanceAgentPluginStatusNotSupported InstanceAgentPluginStatusEnum = "NOT_SUPPORTED"
	InstanceAgentPluginStatusInvalid      InstanceAgentPluginStatusEnum = "INVALID"
)

var mappingInstanceAgentPluginStatus = map[string]InstanceAgentPluginStatusEnum{
	"RUNNING":       InstanceAgentPluginStatusRunning,
	"STOPPED":       InstanceAgentPluginStatusStopped,
	"NOT_SUPPORTED": InstanceAgentPluginStatusNotSupported,
	"INVALID":       InstanceAgentPluginStatusInvalid,
}

// GetInstanceAgentPluginStatusEnumValues Enumerates the set of values for InstanceAgentPluginStatusEnum
func GetInstanceAgentPluginStatusEnumValues() []InstanceAgentPluginStatusEnum {
	values := make([]InstanceAgentPluginStatusEnum, 0)
	for _, v := range mappingInstanceAgentPluginStatus {
		values = append(values, v)
	}
	return values
}
