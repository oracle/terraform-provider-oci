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

// CreateIntegrationInstanceDetails The information about new IntegrationInstance.
type CreateIntegrationInstanceDetails struct {

	// Integration Instance Identifier.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Standard or Enterprise type
	IntegrationInstanceType CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnum `mandatory:"true" json:"integrationInstanceType"`

	// Bring your own license.
	IsByol *bool `mandatory:"true" json:"isByol"`

	// The number of configured message packs
	MessagePacks *int `mandatory:"true" json:"messagePacks"`

	// Simple key-value pair that is applied without any predefined name,
	// type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to
	// namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// IDCS Authentication token. This is required for all realms with IDCS. Its optional as its not required for non IDCS realms.
	IdcsAt *string `mandatory:"false" json:"idcsAt"`

	// Visual Builder is enabled or not.
	IsVisualBuilderEnabled *bool `mandatory:"false" json:"isVisualBuilderEnabled"`

	CustomEndpoint *CreateCustomEndpointDetails `mandatory:"false" json:"customEndpoint"`

	// A list of alternate custom endpoints to be used for the integration instance URL
	// (contact Oracle for alternateCustomEndpoints availability for a specific instance).
	AlternateCustomEndpoints []CreateCustomEndpointDetails `mandatory:"false" json:"alternateCustomEndpoints"`

	// Optional parameter specifying which entitlement to use for billing purposes. Only required if the account possesses more than one entitlement.
	ConsumptionModel CreateIntegrationInstanceDetailsConsumptionModelEnum `mandatory:"false" json:"consumptionModel,omitempty"`

	// The file server is enabled or not.
	IsFileServerEnabled *bool `mandatory:"false" json:"isFileServerEnabled"`

	NetworkEndpointDetails NetworkEndpointDetails `mandatory:"false" json:"networkEndpointDetails"`
}

func (m CreateIntegrationInstanceDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *CreateIntegrationInstanceDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		FreeformTags             map[string]string                                           `json:"freeformTags"`
		DefinedTags              map[string]map[string]interface{}                           `json:"definedTags"`
		IdcsAt                   *string                                                     `json:"idcsAt"`
		IsVisualBuilderEnabled   *bool                                                       `json:"isVisualBuilderEnabled"`
		CustomEndpoint           *CreateCustomEndpointDetails                                `json:"customEndpoint"`
		AlternateCustomEndpoints []CreateCustomEndpointDetails                               `json:"alternateCustomEndpoints"`
		ConsumptionModel         CreateIntegrationInstanceDetailsConsumptionModelEnum        `json:"consumptionModel"`
		IsFileServerEnabled      *bool                                                       `json:"isFileServerEnabled"`
		NetworkEndpointDetails   networkendpointdetails                                      `json:"networkEndpointDetails"`
		DisplayName              *string                                                     `json:"displayName"`
		CompartmentId            *string                                                     `json:"compartmentId"`
		IntegrationInstanceType  CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnum `json:"integrationInstanceType"`
		IsByol                   *bool                                                       `json:"isByol"`
		MessagePacks             *int                                                        `json:"messagePacks"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.IdcsAt = model.IdcsAt

	m.IsVisualBuilderEnabled = model.IsVisualBuilderEnabled

	m.CustomEndpoint = model.CustomEndpoint

	m.AlternateCustomEndpoints = make([]CreateCustomEndpointDetails, len(model.AlternateCustomEndpoints))
	for i, n := range model.AlternateCustomEndpoints {
		m.AlternateCustomEndpoints[i] = n
	}

	m.ConsumptionModel = model.ConsumptionModel

	m.IsFileServerEnabled = model.IsFileServerEnabled

	nn, e = model.NetworkEndpointDetails.UnmarshalPolymorphicJSON(model.NetworkEndpointDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.NetworkEndpointDetails = nn.(NetworkEndpointDetails)
	} else {
		m.NetworkEndpointDetails = nil
	}

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.IntegrationInstanceType = model.IntegrationInstanceType

	m.IsByol = model.IsByol

	m.MessagePacks = model.MessagePacks

	return
}

// CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnum Enum with underlying type: string
type CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnum string

// Set of constants representing the allowable values for CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnum
const (
	CreateIntegrationInstanceDetailsIntegrationInstanceTypeStandard   CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnum = "STANDARD"
	CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnterprise CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnum = "ENTERPRISE"
)

var mappingCreateIntegrationInstanceDetailsIntegrationInstanceType = map[string]CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnum{
	"STANDARD":   CreateIntegrationInstanceDetailsIntegrationInstanceTypeStandard,
	"ENTERPRISE": CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnterprise,
}

// GetCreateIntegrationInstanceDetailsIntegrationInstanceTypeEnumValues Enumerates the set of values for CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnum
func GetCreateIntegrationInstanceDetailsIntegrationInstanceTypeEnumValues() []CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnum {
	values := make([]CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnum, 0)
	for _, v := range mappingCreateIntegrationInstanceDetailsIntegrationInstanceType {
		values = append(values, v)
	}
	return values
}

// CreateIntegrationInstanceDetailsConsumptionModelEnum Enum with underlying type: string
type CreateIntegrationInstanceDetailsConsumptionModelEnum string

// Set of constants representing the allowable values for CreateIntegrationInstanceDetailsConsumptionModelEnum
const (
	CreateIntegrationInstanceDetailsConsumptionModelUcm      CreateIntegrationInstanceDetailsConsumptionModelEnum = "UCM"
	CreateIntegrationInstanceDetailsConsumptionModelGov      CreateIntegrationInstanceDetailsConsumptionModelEnum = "GOV"
	CreateIntegrationInstanceDetailsConsumptionModelOic4saas CreateIntegrationInstanceDetailsConsumptionModelEnum = "OIC4SAAS"
)

var mappingCreateIntegrationInstanceDetailsConsumptionModel = map[string]CreateIntegrationInstanceDetailsConsumptionModelEnum{
	"UCM":      CreateIntegrationInstanceDetailsConsumptionModelUcm,
	"GOV":      CreateIntegrationInstanceDetailsConsumptionModelGov,
	"OIC4SAAS": CreateIntegrationInstanceDetailsConsumptionModelOic4saas,
}

// GetCreateIntegrationInstanceDetailsConsumptionModelEnumValues Enumerates the set of values for CreateIntegrationInstanceDetailsConsumptionModelEnum
func GetCreateIntegrationInstanceDetailsConsumptionModelEnumValues() []CreateIntegrationInstanceDetailsConsumptionModelEnum {
	values := make([]CreateIntegrationInstanceDetailsConsumptionModelEnum, 0)
	for _, v := range mappingCreateIntegrationInstanceDetailsConsumptionModel {
		values = append(values, v)
	}
	return values
}
