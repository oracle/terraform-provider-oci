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

// ExadataInsight Exadata insight resource.
type ExadataInsight interface {

	// Exadata insight identifier
	GetId() *string

	// Compartment identifier of the Exadata insight resource
	GetCompartmentId() *string

	// The Exadata system name. If the Exadata systems managed by Enterprise Manager, the name is unique amongst the Exadata systems managed by the same Enterprise Manager.
	GetExadataName() *string

	// Indicates the status of an Exadata insight in Operations Insights
	GetStatus() ResourceStatusEnum

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// The time the the Exadata insight was first enabled. An RFC3339 formatted datetime string
	GetTimeCreated() *common.SDKTime

	// The current state of the Exadata insight.
	GetLifecycleState() ExadataInsightLifecycleStateEnum

	// The user-friendly name for the Exadata system. The name does not have to be unique.
	GetExadataDisplayName() *string

	// Operations Insights internal representation of the the Exadata system type.
	GetExadataType() ExadataTypeEnum

	// Exadata rack type.
	GetExadataRackType() ExadataRackTypeEnum

	// true if virtualization is used in the Exadata system
	GetIsVirtualizedExadata() *bool

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}

	// The time the Exadata insight was updated. An RFC3339 formatted datetime string
	GetTimeUpdated() *common.SDKTime

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	GetLifecycleDetails() *string
}

type exadatainsight struct {
	JsonData             []byte
	Id                   *string                           `mandatory:"true" json:"id"`
	CompartmentId        *string                           `mandatory:"true" json:"compartmentId"`
	ExadataName          *string                           `mandatory:"true" json:"exadataName"`
	Status               ResourceStatusEnum                `mandatory:"true" json:"status"`
	FreeformTags         map[string]string                 `mandatory:"true" json:"freeformTags"`
	DefinedTags          map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`
	TimeCreated          *common.SDKTime                   `mandatory:"true" json:"timeCreated"`
	LifecycleState       ExadataInsightLifecycleStateEnum  `mandatory:"true" json:"lifecycleState"`
	ExadataDisplayName   *string                           `mandatory:"false" json:"exadataDisplayName"`
	ExadataType          ExadataTypeEnum                   `mandatory:"false" json:"exadataType,omitempty"`
	ExadataRackType      ExadataRackTypeEnum               `mandatory:"false" json:"exadataRackType,omitempty"`
	IsVirtualizedExadata *bool                             `mandatory:"false" json:"isVirtualizedExadata"`
	SystemTags           map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	TimeUpdated          *common.SDKTime                   `mandatory:"false" json:"timeUpdated"`
	LifecycleDetails     *string                           `mandatory:"false" json:"lifecycleDetails"`
	EntitySource         string                            `json:"entitySource"`
}

// UnmarshalJSON unmarshals json
func (m *exadatainsight) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerexadatainsight exadatainsight
	s := struct {
		Model Unmarshalerexadatainsight
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.CompartmentId = s.Model.CompartmentId
	m.ExadataName = s.Model.ExadataName
	m.Status = s.Model.Status
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.TimeCreated = s.Model.TimeCreated
	m.LifecycleState = s.Model.LifecycleState
	m.ExadataDisplayName = s.Model.ExadataDisplayName
	m.ExadataType = s.Model.ExadataType
	m.ExadataRackType = s.Model.ExadataRackType
	m.IsVirtualizedExadata = s.Model.IsVirtualizedExadata
	m.SystemTags = s.Model.SystemTags
	m.TimeUpdated = s.Model.TimeUpdated
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.EntitySource = s.Model.EntitySource

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *exadatainsight) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.EntitySource {
	case "EM_MANAGED_EXTERNAL_EXADATA":
		mm := EmManagedExternalExadataInsight{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetId returns Id
func (m exadatainsight) GetId() *string {
	return m.Id
}

//GetCompartmentId returns CompartmentId
func (m exadatainsight) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetExadataName returns ExadataName
func (m exadatainsight) GetExadataName() *string {
	return m.ExadataName
}

//GetStatus returns Status
func (m exadatainsight) GetStatus() ResourceStatusEnum {
	return m.Status
}

//GetFreeformTags returns FreeformTags
func (m exadatainsight) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m exadatainsight) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetTimeCreated returns TimeCreated
func (m exadatainsight) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetLifecycleState returns LifecycleState
func (m exadatainsight) GetLifecycleState() ExadataInsightLifecycleStateEnum {
	return m.LifecycleState
}

//GetExadataDisplayName returns ExadataDisplayName
func (m exadatainsight) GetExadataDisplayName() *string {
	return m.ExadataDisplayName
}

//GetExadataType returns ExadataType
func (m exadatainsight) GetExadataType() ExadataTypeEnum {
	return m.ExadataType
}

//GetExadataRackType returns ExadataRackType
func (m exadatainsight) GetExadataRackType() ExadataRackTypeEnum {
	return m.ExadataRackType
}

//GetIsVirtualizedExadata returns IsVirtualizedExadata
func (m exadatainsight) GetIsVirtualizedExadata() *bool {
	return m.IsVirtualizedExadata
}

//GetSystemTags returns SystemTags
func (m exadatainsight) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

//GetTimeUpdated returns TimeUpdated
func (m exadatainsight) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleDetails returns LifecycleDetails
func (m exadatainsight) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

func (m exadatainsight) String() string {
	return common.PointerString(m)
}
