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

// ManagementAgentSummary The summary of the Management Agent inventory including the associated plugins.
type ManagementAgentSummary struct {

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

	// true if the agent can be upgraded automatically; false if it must be upgraded manually. This flag is derived from the tenancy level auto upgrade preference.
	IsAgentAutoUpgradable *bool `mandatory:"false" json:"isAgentAutoUpgradable"`

	// The time the Management Agent was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the Management Agent was last updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Management Agent host machine name
	Host *string `mandatory:"false" json:"host"`

	// Host resource ocid
	HostId *string `mandatory:"false" json:"hostId"`

	// list of managementAgentPlugins associated with the agent
	PluginList []ManagementAgentPluginDetails `mandatory:"false" json:"pluginList"`

	// The time the Management Agent has last recorded its heartbeat. An RFC3339 formatted datetime string
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

	// list of dataSources summaries associated with the agent
	DataSourceSummaryList []DataSourceSummaryItem `mandatory:"false" json:"dataSourceSummaryList"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ManagementAgentSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagementAgentSummary) ValidateEnumValue() (bool, error) {
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
func (m *ManagementAgentSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		InstallKeyId            *string                           `json:"installKeyId"`
		DisplayName             *string                           `json:"displayName"`
		PlatformType            PlatformTypesEnum                 `json:"platformType"`
		PlatformName            *string                           `json:"platformName"`
		PlatformVersion         *string                           `json:"platformVersion"`
		ResourceArtifactVersion *string                           `json:"resourceArtifactVersion"`
		IsAgentAutoUpgradable   *bool                             `json:"isAgentAutoUpgradable"`
		TimeCreated             *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated             *common.SDKTime                   `json:"timeUpdated"`
		Host                    *string                           `json:"host"`
		HostId                  *string                           `json:"hostId"`
		PluginList              []ManagementAgentPluginDetails    `json:"pluginList"`
		TimeLastHeartbeat       *common.SDKTime                   `json:"timeLastHeartbeat"`
		AvailabilityStatus      AvailabilityStatusEnum            `json:"availabilityStatus"`
		LifecycleState          LifecycleStatesEnum               `json:"lifecycleState"`
		LifecycleDetails        *string                           `json:"lifecycleDetails"`
		IsCustomerDeployed      *bool                             `json:"isCustomerDeployed"`
		InstallType             InstallTypesEnum                  `json:"installType"`
		DataSourceSummaryList   []datasourcesummaryitem           `json:"dataSourceSummaryList"`
		FreeformTags            map[string]string                 `json:"freeformTags"`
		DefinedTags             map[string]map[string]interface{} `json:"definedTags"`
		Id                      *string                           `json:"id"`
		Version                 *string                           `json:"version"`
		CompartmentId           *string                           `json:"compartmentId"`
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

	m.IsAgentAutoUpgradable = model.IsAgentAutoUpgradable

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.Host = model.Host

	m.HostId = model.HostId

	m.PluginList = make([]ManagementAgentPluginDetails, len(model.PluginList))
	copy(m.PluginList, model.PluginList)
	m.TimeLastHeartbeat = model.TimeLastHeartbeat

	m.AvailabilityStatus = model.AvailabilityStatus

	m.LifecycleState = model.LifecycleState

	m.LifecycleDetails = model.LifecycleDetails

	m.IsCustomerDeployed = model.IsCustomerDeployed

	m.InstallType = model.InstallType

	m.DataSourceSummaryList = make([]DataSourceSummaryItem, len(model.DataSourceSummaryList))
	for i, n := range model.DataSourceSummaryList {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.DataSourceSummaryList[i] = nn.(DataSourceSummaryItem)
		} else {
			m.DataSourceSummaryList[i] = nil
		}
	}
	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.Id = model.Id

	m.Version = model.Version

	m.CompartmentId = model.CompartmentId

	return
}
