// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// API for Management Agent Cloud Service
//

package managementagent

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ManagementAgent The details of the Management Agent inventory including the associated plugins.
type ManagementAgent struct {

	// agent identifier
	Id *string `mandatory:"true" json:"id"`

	// Management Agent Version
	Version *string `mandatory:"true" json:"version"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// agent install key identifier
	InstallKeyId *string `mandatory:"false" json:"installKeyId"`

	// Management Agent Name
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Platform Type
	PlatformType PlatformTypesEnum `mandatory:"false" json:"platformType,omitempty"`

	// Platform Name
	PlatformName *string `mandatory:"false" json:"platformName"`

	// Platform Version
	PlatformVersion *string `mandatory:"false" json:"platformVersion"`

	// Management Agent host machine name
	Host *string `mandatory:"false" json:"host"`

	// Path where Management Agent is installed
	InstallPath *string `mandatory:"false" json:"installPath"`

	// list of managementAgentPlugins associated with the agent
	PluginList []ManagementAgentPluginDetails `mandatory:"false" json:"pluginList"`

	// true if the agent can be upgraded automatically; false if it must be upgraded manually. true is currently unsupported.
	IsAgentAutoUpgradable *bool `mandatory:"false" json:"isAgentAutoUpgradable"`

	// The time the Management Agent was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the Management Agent was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The time the Management Agent has last recorded its health status in telemetry. This value will be null if the agent has not recorded its health status in last 7 days. An RFC3339 formatted datetime string
	TimeLastHeartbeat *common.SDKTime `mandatory:"false" json:"timeLastHeartbeat"`

	// The current state of managementAgent
	LifecycleState LifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ManagementAgent) String() string {
	return common.PointerString(m)
}
