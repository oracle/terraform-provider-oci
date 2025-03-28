// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Visual Builder API
//
// Oracle Visual Builder enables developers to quickly build web and mobile applications. With a visual development environment that makes it easy to connect to Oracle data and third-party REST services, developers can build modern, consumer-grade applications in a fraction of the time it would take in other tools.
// The Visual Builder Instance Management API allows users to create and manage a Visual Builder instance.
//

package visualbuilder

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateVbInstanceDetails The information about new VbInstance.
type CreateVbInstanceDetails struct {

	// Vb Instance Identifier.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The number of Nodes
	NodeCount *int `mandatory:"true" json:"nodeCount"`

	// Simple key-value pair that is applied without any predefined name,
	// type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to
	// namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Encrypted IDCS Open ID token. This is required for pre-UCPIS cloud accounts, but not UCPIS, hence not a required parameter
	IdcsOpenId *string `mandatory:"false" json:"idcsOpenId"`

	// Visual Builder is enabled or not.
	IsVisualBuilderEnabled *bool `mandatory:"false" json:"isVisualBuilderEnabled"`

	CustomEndpoint *CreateCustomEndpointDetails `mandatory:"false" json:"customEndpoint"`

	// A list of alternate custom endpoints to be used for the vb instance URL
	// (contact Oracle for alternateCustomEndpoints availability for a specific instance).
	AlternateCustomEndpoints []CreateCustomEndpointDetails `mandatory:"false" json:"alternateCustomEndpoints"`

	// Optional parameter specifying which entitlement to use for billing purposes. Only required if the account possesses more than one entitlement.
	ConsumptionModel CreateVbInstanceDetailsConsumptionModelEnum `mandatory:"false" json:"consumptionModel,omitempty"`

	NetworkEndpointDetails NetworkEndpointDetails `mandatory:"false" json:"networkEndpointDetails"`
}

func (m CreateVbInstanceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateVbInstanceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateVbInstanceDetailsConsumptionModelEnum(string(m.ConsumptionModel)); !ok && m.ConsumptionModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConsumptionModel: %s. Supported values are: %s.", m.ConsumptionModel, strings.Join(GetCreateVbInstanceDetailsConsumptionModelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateVbInstanceDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		FreeformTags             map[string]string                           `json:"freeformTags"`
		DefinedTags              map[string]map[string]interface{}           `json:"definedTags"`
		IdcsOpenId               *string                                     `json:"idcsOpenId"`
		IsVisualBuilderEnabled   *bool                                       `json:"isVisualBuilderEnabled"`
		CustomEndpoint           *CreateCustomEndpointDetails                `json:"customEndpoint"`
		AlternateCustomEndpoints []CreateCustomEndpointDetails               `json:"alternateCustomEndpoints"`
		ConsumptionModel         CreateVbInstanceDetailsConsumptionModelEnum `json:"consumptionModel"`
		NetworkEndpointDetails   networkendpointdetails                      `json:"networkEndpointDetails"`
		DisplayName              *string                                     `json:"displayName"`
		CompartmentId            *string                                     `json:"compartmentId"`
		NodeCount                *int                                        `json:"nodeCount"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.IdcsOpenId = model.IdcsOpenId

	m.IsVisualBuilderEnabled = model.IsVisualBuilderEnabled

	m.CustomEndpoint = model.CustomEndpoint

	m.AlternateCustomEndpoints = make([]CreateCustomEndpointDetails, len(model.AlternateCustomEndpoints))
	copy(m.AlternateCustomEndpoints, model.AlternateCustomEndpoints)
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

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.NodeCount = model.NodeCount

	return
}

// CreateVbInstanceDetailsConsumptionModelEnum Enum with underlying type: string
type CreateVbInstanceDetailsConsumptionModelEnum string

// Set of constants representing the allowable values for CreateVbInstanceDetailsConsumptionModelEnum
const (
	CreateVbInstanceDetailsConsumptionModelUcm     CreateVbInstanceDetailsConsumptionModelEnum = "UCM"
	CreateVbInstanceDetailsConsumptionModelGov     CreateVbInstanceDetailsConsumptionModelEnum = "GOV"
	CreateVbInstanceDetailsConsumptionModelVb4saas CreateVbInstanceDetailsConsumptionModelEnum = "VB4SAAS"
)

var mappingCreateVbInstanceDetailsConsumptionModelEnum = map[string]CreateVbInstanceDetailsConsumptionModelEnum{
	"UCM":     CreateVbInstanceDetailsConsumptionModelUcm,
	"GOV":     CreateVbInstanceDetailsConsumptionModelGov,
	"VB4SAAS": CreateVbInstanceDetailsConsumptionModelVb4saas,
}

var mappingCreateVbInstanceDetailsConsumptionModelEnumLowerCase = map[string]CreateVbInstanceDetailsConsumptionModelEnum{
	"ucm":     CreateVbInstanceDetailsConsumptionModelUcm,
	"gov":     CreateVbInstanceDetailsConsumptionModelGov,
	"vb4saas": CreateVbInstanceDetailsConsumptionModelVb4saas,
}

// GetCreateVbInstanceDetailsConsumptionModelEnumValues Enumerates the set of values for CreateVbInstanceDetailsConsumptionModelEnum
func GetCreateVbInstanceDetailsConsumptionModelEnumValues() []CreateVbInstanceDetailsConsumptionModelEnum {
	values := make([]CreateVbInstanceDetailsConsumptionModelEnum, 0)
	for _, v := range mappingCreateVbInstanceDetailsConsumptionModelEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateVbInstanceDetailsConsumptionModelEnumStringValues Enumerates the set of values in String for CreateVbInstanceDetailsConsumptionModelEnum
func GetCreateVbInstanceDetailsConsumptionModelEnumStringValues() []string {
	return []string{
		"UCM",
		"GOV",
		"VB4SAAS",
	}
}

// GetMappingCreateVbInstanceDetailsConsumptionModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateVbInstanceDetailsConsumptionModelEnum(val string) (CreateVbInstanceDetailsConsumptionModelEnum, bool) {
	enum, ok := mappingCreateVbInstanceDetailsConsumptionModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
