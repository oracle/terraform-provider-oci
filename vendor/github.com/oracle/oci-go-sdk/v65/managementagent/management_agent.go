// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// Use the Management Agent API to manage your infrastructure's management agents, including their plugins and install keys.
// For more information, see Management Agent (https://docs.cloud.oracle.com/iaas/management-agents/index.html).
//

package managementagent

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

	// Version of the deployment artifact instantiated by this Management Agent.
	// The format for Standalone resourceMode is YYMMDD.HHMM, and the format for other modes
	// (whose artifacts are based upon Standalone but can advance independently)
	// is YYMMDD.HHMM.VVVVVVVVVVVV.
	// VVVVVVVVVVVV is always a numeric value between 000000000000 and 999999999999
	ResourceArtifactVersion *string `mandatory:"false" json:"resourceArtifactVersion"`

	// Management Agent host machine name
	Host *string `mandatory:"false" json:"host"`

	// Host resource ocid
	HostId *string `mandatory:"false" json:"hostId"`

	// Path where Management Agent is installed
	InstallPath *string `mandatory:"false" json:"installPath"`

	// list of managementAgentPlugins associated with the agent
	PluginList []ManagementAgentPluginDetails `mandatory:"false" json:"pluginList"`

	// true if the agent can be upgraded automatically; false if it must be upgraded manually. This flag is derived from the tenancy level auto upgrade preference.
	IsAgentAutoUpgradable *bool `mandatory:"false" json:"isAgentAutoUpgradable"`

	// The time the Management Agent was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the Management Agent was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The time the Management Agent has last recorded its health status in telemetry. This value will be null if the agent has not recorded its health status in last 7 days. An RFC3339 formatted datetime string
	TimeLastHeartbeat *common.SDKTime `mandatory:"false" json:"timeLastHeartbeat"`

	// The current availability status of managementAgent
	AvailabilityStatus AvailabilityStatusEnum `mandatory:"false" json:"availabilityStatus,omitempty"`

	// The current state of managementAgent
	LifecycleState LifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// true, if the agent image is manually downloaded and installed. false, if the agent is deployed as a plugin in Oracle Cloud Agent.
	IsCustomerDeployed *bool `mandatory:"false" json:"isCustomerDeployed"`

	// The install type, either AGENT or GATEWAY
	InstallType InstallTypesEnum `mandatory:"false" json:"installType,omitempty"`

	// Additional properties for this Management Agent
	ManagementAgentProperties []ManagementAgentProperty `mandatory:"false" json:"managementAgentProperties"`

	// list of dataSources associated with the agent
	DataSourceList []DataSource `mandatory:"false" json:"dataSourceList"`

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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagementAgent) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPlatformTypesEnum(string(m.PlatformType)); !ok && m.PlatformType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlatformType: %s. Supported values are: %s.", m.PlatformType, strings.Join(GetPlatformTypesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAvailabilityStatusEnum(string(m.AvailabilityStatus)); !ok && m.AvailabilityStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AvailabilityStatus: %s. Supported values are: %s.", m.AvailabilityStatus, strings.Join(GetAvailabilityStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStatesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingInstallTypesEnum(string(m.InstallType)); !ok && m.InstallType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InstallType: %s. Supported values are: %s.", m.InstallType, strings.Join(GetInstallTypesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ManagementAgent) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		InstallKeyId              *string                           `json:"installKeyId"`
		DisplayName               *string                           `json:"displayName"`
		PlatformType              PlatformTypesEnum                 `json:"platformType"`
		PlatformName              *string                           `json:"platformName"`
		PlatformVersion           *string                           `json:"platformVersion"`
		ResourceArtifactVersion   *string                           `json:"resourceArtifactVersion"`
		Host                      *string                           `json:"host"`
		HostId                    *string                           `json:"hostId"`
		InstallPath               *string                           `json:"installPath"`
		PluginList                []ManagementAgentPluginDetails    `json:"pluginList"`
		IsAgentAutoUpgradable     *bool                             `json:"isAgentAutoUpgradable"`
		TimeCreated               *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated               *common.SDKTime                   `json:"timeUpdated"`
		TimeLastHeartbeat         *common.SDKTime                   `json:"timeLastHeartbeat"`
		AvailabilityStatus        AvailabilityStatusEnum            `json:"availabilityStatus"`
		LifecycleState            LifecycleStatesEnum               `json:"lifecycleState"`
		LifecycleDetails          *string                           `json:"lifecycleDetails"`
		IsCustomerDeployed        *bool                             `json:"isCustomerDeployed"`
		InstallType               InstallTypesEnum                  `json:"installType"`
		ManagementAgentProperties []ManagementAgentProperty         `json:"managementAgentProperties"`
		DataSourceList            []datasource                      `json:"dataSourceList"`
		FreeformTags              map[string]string                 `json:"freeformTags"`
		DefinedTags               map[string]map[string]interface{} `json:"definedTags"`
		Id                        *string                           `json:"id"`
		Version                   *string                           `json:"version"`
		CompartmentId             *string                           `json:"compartmentId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.InstallKeyId = model.InstallKeyId

	m.DisplayName = model.DisplayName

	m.PlatformType = model.PlatformType

	m.PlatformName = model.PlatformName

	m.PlatformVersion = model.PlatformVersion

	m.ResourceArtifactVersion = model.ResourceArtifactVersion

	m.Host = model.Host

	m.HostId = model.HostId

	m.InstallPath = model.InstallPath

	m.PluginList = make([]ManagementAgentPluginDetails, len(model.PluginList))
	copy(m.PluginList, model.PluginList)
	m.IsAgentAutoUpgradable = model.IsAgentAutoUpgradable

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.TimeLastHeartbeat = model.TimeLastHeartbeat

	m.AvailabilityStatus = model.AvailabilityStatus

	m.LifecycleState = model.LifecycleState

	m.LifecycleDetails = model.LifecycleDetails

	m.IsCustomerDeployed = model.IsCustomerDeployed

	m.InstallType = model.InstallType

	m.ManagementAgentProperties = make([]ManagementAgentProperty, len(model.ManagementAgentProperties))
	copy(m.ManagementAgentProperties, model.ManagementAgentProperties)
	m.DataSourceList = make([]DataSource, len(model.DataSourceList))
	for i, n := range model.DataSourceList {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.DataSourceList[i] = nn.(DataSource)
		} else {
			m.DataSourceList[i] = nil
		}
	}
	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.Id = model.Id

	m.Version = model.Version

	m.CompartmentId = model.CompartmentId

	return
}
