// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Integration API
//
// Oracle Integration API.
//

package integration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// IntegrationInstanceSummary Summary of the Integration Instance.
type IntegrationInstanceSummary struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// Integration Instance Identifier, can be renamed.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Standard or Enterprise type
	IntegrationInstanceType IntegrationInstanceSummaryIntegrationInstanceTypeEnum `mandatory:"true" json:"integrationInstanceType"`

	// Bring your own license.
	IsByol *bool `mandatory:"true" json:"isByol"`

	// The Integration Instance URL.
	InstanceUrl *string `mandatory:"true" json:"instanceUrl"`

	// The number of configured message packs (if any)
	MessagePacks *int `mandatory:"true" json:"messagePacks"`

	// The time the the Integration Instance was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the IntegrationInstance was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the Integration Instance.
	LifecycleState IntegrationInstanceSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// An message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	StateMessage *string `mandatory:"false" json:"stateMessage"`

	// The file server is enabled or not.
	IsFileServerEnabled *bool `mandatory:"false" json:"isFileServerEnabled"`

	// Visual Builder is enabled or not.
	IsVisualBuilderEnabled *bool `mandatory:"false" json:"isVisualBuilderEnabled"`

	CustomEndpoint *CustomEndpointDetails `mandatory:"false" json:"customEndpoint"`

	// A list of alternate custom endpoints used for the integration instance URL.
	AlternateCustomEndpoints []CustomEndpointDetails `mandatory:"false" json:"alternateCustomEndpoints"`

	// The entitlement used for billing purposes.
	ConsumptionModel IntegrationInstanceSummaryConsumptionModelEnum `mandatory:"false" json:"consumptionModel,omitempty"`

	NetworkEndpointDetails NetworkEndpointDetails `mandatory:"false" json:"networkEndpointDetails"`
}

func (m IntegrationInstanceSummary) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *IntegrationInstanceSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TimeCreated              *common.SDKTime                                       `json:"timeCreated"`
		TimeUpdated              *common.SDKTime                                       `json:"timeUpdated"`
		LifecycleState           IntegrationInstanceSummaryLifecycleStateEnum          `json:"lifecycleState"`
		StateMessage             *string                                               `json:"stateMessage"`
		IsFileServerEnabled      *bool                                                 `json:"isFileServerEnabled"`
		IsVisualBuilderEnabled   *bool                                                 `json:"isVisualBuilderEnabled"`
		CustomEndpoint           *CustomEndpointDetails                                `json:"customEndpoint"`
		AlternateCustomEndpoints []CustomEndpointDetails                               `json:"alternateCustomEndpoints"`
		ConsumptionModel         IntegrationInstanceSummaryConsumptionModelEnum        `json:"consumptionModel"`
		NetworkEndpointDetails   networkendpointdetails                                `json:"networkEndpointDetails"`
		Id                       *string                                               `json:"id"`
		DisplayName              *string                                               `json:"displayName"`
		CompartmentId            *string                                               `json:"compartmentId"`
		IntegrationInstanceType  IntegrationInstanceSummaryIntegrationInstanceTypeEnum `json:"integrationInstanceType"`
		IsByol                   *bool                                                 `json:"isByol"`
		InstanceUrl              *string                                               `json:"instanceUrl"`
		MessagePacks             *int                                                  `json:"messagePacks"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

	m.StateMessage = model.StateMessage

	m.IsFileServerEnabled = model.IsFileServerEnabled

	m.IsVisualBuilderEnabled = model.IsVisualBuilderEnabled

	m.CustomEndpoint = model.CustomEndpoint

	m.AlternateCustomEndpoints = make([]CustomEndpointDetails, len(model.AlternateCustomEndpoints))
	for i, n := range model.AlternateCustomEndpoints {
		m.AlternateCustomEndpoints[i] = n
	}

	m.ConsumptionModel = model.ConsumptionModel

	nn, e = model.NetworkEndpointDetails.UnmarshalPolymorphicJSON(model.NetworkEndpointDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.NetworkEndpointDetails = nn.(NetworkEndpointDetails)
	} else {
		m.NetworkEndpointDetails = nil
	}

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.IntegrationInstanceType = model.IntegrationInstanceType

	m.IsByol = model.IsByol

	m.InstanceUrl = model.InstanceUrl

	m.MessagePacks = model.MessagePacks

	return
}

// IntegrationInstanceSummaryIntegrationInstanceTypeEnum Enum with underlying type: string
type IntegrationInstanceSummaryIntegrationInstanceTypeEnum string

// Set of constants representing the allowable values for IntegrationInstanceSummaryIntegrationInstanceTypeEnum
const (
	IntegrationInstanceSummaryIntegrationInstanceTypeStandard   IntegrationInstanceSummaryIntegrationInstanceTypeEnum = "STANDARD"
	IntegrationInstanceSummaryIntegrationInstanceTypeEnterprise IntegrationInstanceSummaryIntegrationInstanceTypeEnum = "ENTERPRISE"
)

var mappingIntegrationInstanceSummaryIntegrationInstanceType = map[string]IntegrationInstanceSummaryIntegrationInstanceTypeEnum{
	"STANDARD":   IntegrationInstanceSummaryIntegrationInstanceTypeStandard,
	"ENTERPRISE": IntegrationInstanceSummaryIntegrationInstanceTypeEnterprise,
}

// GetIntegrationInstanceSummaryIntegrationInstanceTypeEnumValues Enumerates the set of values for IntegrationInstanceSummaryIntegrationInstanceTypeEnum
func GetIntegrationInstanceSummaryIntegrationInstanceTypeEnumValues() []IntegrationInstanceSummaryIntegrationInstanceTypeEnum {
	values := make([]IntegrationInstanceSummaryIntegrationInstanceTypeEnum, 0)
	for _, v := range mappingIntegrationInstanceSummaryIntegrationInstanceType {
		values = append(values, v)
	}
	return values
}

// IntegrationInstanceSummaryLifecycleStateEnum Enum with underlying type: string
type IntegrationInstanceSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for IntegrationInstanceSummaryLifecycleStateEnum
const (
	IntegrationInstanceSummaryLifecycleStateCreating IntegrationInstanceSummaryLifecycleStateEnum = "CREATING"
	IntegrationInstanceSummaryLifecycleStateUpdating IntegrationInstanceSummaryLifecycleStateEnum = "UPDATING"
	IntegrationInstanceSummaryLifecycleStateActive   IntegrationInstanceSummaryLifecycleStateEnum = "ACTIVE"
	IntegrationInstanceSummaryLifecycleStateInactive IntegrationInstanceSummaryLifecycleStateEnum = "INACTIVE"
	IntegrationInstanceSummaryLifecycleStateDeleting IntegrationInstanceSummaryLifecycleStateEnum = "DELETING"
	IntegrationInstanceSummaryLifecycleStateDeleted  IntegrationInstanceSummaryLifecycleStateEnum = "DELETED"
	IntegrationInstanceSummaryLifecycleStateFailed   IntegrationInstanceSummaryLifecycleStateEnum = "FAILED"
)

var mappingIntegrationInstanceSummaryLifecycleState = map[string]IntegrationInstanceSummaryLifecycleStateEnum{
	"CREATING": IntegrationInstanceSummaryLifecycleStateCreating,
	"UPDATING": IntegrationInstanceSummaryLifecycleStateUpdating,
	"ACTIVE":   IntegrationInstanceSummaryLifecycleStateActive,
	"INACTIVE": IntegrationInstanceSummaryLifecycleStateInactive,
	"DELETING": IntegrationInstanceSummaryLifecycleStateDeleting,
	"DELETED":  IntegrationInstanceSummaryLifecycleStateDeleted,
	"FAILED":   IntegrationInstanceSummaryLifecycleStateFailed,
}

// GetIntegrationInstanceSummaryLifecycleStateEnumValues Enumerates the set of values for IntegrationInstanceSummaryLifecycleStateEnum
func GetIntegrationInstanceSummaryLifecycleStateEnumValues() []IntegrationInstanceSummaryLifecycleStateEnum {
	values := make([]IntegrationInstanceSummaryLifecycleStateEnum, 0)
	for _, v := range mappingIntegrationInstanceSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}

// IntegrationInstanceSummaryConsumptionModelEnum Enum with underlying type: string
type IntegrationInstanceSummaryConsumptionModelEnum string

// Set of constants representing the allowable values for IntegrationInstanceSummaryConsumptionModelEnum
const (
	IntegrationInstanceSummaryConsumptionModelUcm      IntegrationInstanceSummaryConsumptionModelEnum = "UCM"
	IntegrationInstanceSummaryConsumptionModelGov      IntegrationInstanceSummaryConsumptionModelEnum = "GOV"
	IntegrationInstanceSummaryConsumptionModelOic4saas IntegrationInstanceSummaryConsumptionModelEnum = "OIC4SAAS"
)

var mappingIntegrationInstanceSummaryConsumptionModel = map[string]IntegrationInstanceSummaryConsumptionModelEnum{
	"UCM":      IntegrationInstanceSummaryConsumptionModelUcm,
	"GOV":      IntegrationInstanceSummaryConsumptionModelGov,
	"OIC4SAAS": IntegrationInstanceSummaryConsumptionModelOic4saas,
}

// GetIntegrationInstanceSummaryConsumptionModelEnumValues Enumerates the set of values for IntegrationInstanceSummaryConsumptionModelEnum
func GetIntegrationInstanceSummaryConsumptionModelEnumValues() []IntegrationInstanceSummaryConsumptionModelEnum {
	values := make([]IntegrationInstanceSummaryConsumptionModelEnum, 0)
	for _, v := range mappingIntegrationInstanceSummaryConsumptionModel {
		values = append(values, v)
	}
	return values
}
