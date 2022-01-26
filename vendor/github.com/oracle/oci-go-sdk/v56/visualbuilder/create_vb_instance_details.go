// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Visual Builder API
//
// Oracle Visual Builder enables developers to quickly build web and mobile applications. With a visual development environment that makes it easy to connect to Oracle data and third-party REST services, developers can build modern, consumer-grade applications in a fraction of the time it would take in other tools.
// The Visual Builder Instance Management API allows users to create and manage a Visual Builder instance.
//

package visualbuilder

import (
	"github.com/oracle/oci-go-sdk/v56/common"
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
}

func (m CreateVbInstanceDetails) String() string {
	return common.PointerString(m)
}

// CreateVbInstanceDetailsConsumptionModelEnum Enum with underlying type: string
type CreateVbInstanceDetailsConsumptionModelEnum string

// Set of constants representing the allowable values for CreateVbInstanceDetailsConsumptionModelEnum
const (
	CreateVbInstanceDetailsConsumptionModelUcm     CreateVbInstanceDetailsConsumptionModelEnum = "UCM"
	CreateVbInstanceDetailsConsumptionModelGov     CreateVbInstanceDetailsConsumptionModelEnum = "GOV"
	CreateVbInstanceDetailsConsumptionModelVb4saas CreateVbInstanceDetailsConsumptionModelEnum = "VB4SAAS"
)

var mappingCreateVbInstanceDetailsConsumptionModel = map[string]CreateVbInstanceDetailsConsumptionModelEnum{
	"UCM":     CreateVbInstanceDetailsConsumptionModelUcm,
	"GOV":     CreateVbInstanceDetailsConsumptionModelGov,
	"VB4SAAS": CreateVbInstanceDetailsConsumptionModelVb4saas,
}

// GetCreateVbInstanceDetailsConsumptionModelEnumValues Enumerates the set of values for CreateVbInstanceDetailsConsumptionModelEnum
func GetCreateVbInstanceDetailsConsumptionModelEnumValues() []CreateVbInstanceDetailsConsumptionModelEnum {
	values := make([]CreateVbInstanceDetailsConsumptionModelEnum, 0)
	for _, v := range mappingCreateVbInstanceDetailsConsumptionModel {
		values = append(values, v)
	}
	return values
}
