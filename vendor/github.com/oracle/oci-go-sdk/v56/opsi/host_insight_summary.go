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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// HostInsightSummary Summary of a host insight resource.
type HostInsightSummary interface {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the host insight resource.
	GetId() *string

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	GetCompartmentId() *string

	// The host name. The host name is unique amongst the hosts managed by the same management agent.
	GetHostName() *string

	// The user-friendly name for the host. The name does not have to be unique.
	GetHostDisplayName() *string

	// Operations Insights internal representation of the host type. Possible value is EXTERNAL-HOST.
	GetHostType() *string

	// Processor count. This is the OCPU count for Autonomous Database and CPU core count for other database types.
	GetProcessorCount() *int

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}

	// Indicates the status of a host insight in Operations Insights
	GetStatus() ResourceStatusEnum

	// The time the the host insight was first enabled. An RFC3339 formatted datetime string
	GetTimeCreated() *common.SDKTime

	// The time the host insight was updated. An RFC3339 formatted datetime string
	GetTimeUpdated() *common.SDKTime

	// The current state of the host.
	GetLifecycleState() LifecycleStateEnum

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	GetLifecycleDetails() *string
}

type hostinsightsummary struct {
	JsonData         []byte
	Id               *string                           `mandatory:"true" json:"id"`
	CompartmentId    *string                           `mandatory:"true" json:"compartmentId"`
	HostName         *string                           `mandatory:"true" json:"hostName"`
	HostDisplayName  *string                           `mandatory:"false" json:"hostDisplayName"`
	HostType         *string                           `mandatory:"false" json:"hostType"`
	ProcessorCount   *int                              `mandatory:"false" json:"processorCount"`
	FreeformTags     map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags      map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags       map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	Status           ResourceStatusEnum                `mandatory:"false" json:"status,omitempty"`
	TimeCreated      *common.SDKTime                   `mandatory:"false" json:"timeCreated"`
	TimeUpdated      *common.SDKTime                   `mandatory:"false" json:"timeUpdated"`
	LifecycleState   LifecycleStateEnum                `mandatory:"false" json:"lifecycleState,omitempty"`
	LifecycleDetails *string                           `mandatory:"false" json:"lifecycleDetails"`
	EntitySource     string                            `json:"entitySource"`
}

// UnmarshalJSON unmarshals json
func (m *hostinsightsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerhostinsightsummary hostinsightsummary
	s := struct {
		Model Unmarshalerhostinsightsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.CompartmentId = s.Model.CompartmentId
	m.HostName = s.Model.HostName
	m.HostDisplayName = s.Model.HostDisplayName
	m.HostType = s.Model.HostType
	m.ProcessorCount = s.Model.ProcessorCount
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.Status = s.Model.Status
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.LifecycleState = s.Model.LifecycleState
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.EntitySource = s.Model.EntitySource

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *hostinsightsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.EntitySource {
	case "MACS_MANAGED_EXTERNAL_HOST":
		mm := MacsManagedExternalHostInsightSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "EM_MANAGED_EXTERNAL_HOST":
		mm := EmManagedExternalHostInsightSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetId returns Id
func (m hostinsightsummary) GetId() *string {
	return m.Id
}

//GetCompartmentId returns CompartmentId
func (m hostinsightsummary) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetHostName returns HostName
func (m hostinsightsummary) GetHostName() *string {
	return m.HostName
}

//GetHostDisplayName returns HostDisplayName
func (m hostinsightsummary) GetHostDisplayName() *string {
	return m.HostDisplayName
}

//GetHostType returns HostType
func (m hostinsightsummary) GetHostType() *string {
	return m.HostType
}

//GetProcessorCount returns ProcessorCount
func (m hostinsightsummary) GetProcessorCount() *int {
	return m.ProcessorCount
}

//GetFreeformTags returns FreeformTags
func (m hostinsightsummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m hostinsightsummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m hostinsightsummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

//GetStatus returns Status
func (m hostinsightsummary) GetStatus() ResourceStatusEnum {
	return m.Status
}

//GetTimeCreated returns TimeCreated
func (m hostinsightsummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m hostinsightsummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleState returns LifecycleState
func (m hostinsightsummary) GetLifecycleState() LifecycleStateEnum {
	return m.LifecycleState
}

//GetLifecycleDetails returns LifecycleDetails
func (m hostinsightsummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

func (m hostinsightsummary) String() string {
	return common.PointerString(m)
}
