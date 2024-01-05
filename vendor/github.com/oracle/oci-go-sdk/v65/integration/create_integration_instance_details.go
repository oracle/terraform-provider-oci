// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Integration API
//
// Oracle Integration API.
//

package integration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateIntegrationInstanceDetails The information about new IntegrationInstance.
// Some properties may not be applicable to specific integration types,
// see Differences in Instance Management (https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/application-integration&id=INTOO-GUID-931B5E33-4FE6-4997-93E5-8748516F46AA__GUID-176E43D5-4116-4828-8120-B929DF2A6B5E)
// for details.
type CreateIntegrationInstanceDetails struct {

	// Integration Instance Identifier.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Standard or Enterprise type,
	// Oracle Integration Generation 2 uses ENTERPRISE and STANDARD,
	// Oracle Integration 3 uses ENTERPRISEX and STANDARDX
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

	// Shape
	Shape CreateIntegrationInstanceDetailsShapeEnum `mandatory:"false" json:"shape,omitempty"`

	// The OCID of the identity domain, that will be used to determine the
	// corresponding Idcs Stripe and create an Idcs application within the stripe.
	// This parameter is mutually exclusive with parameter: idcsAt, i.e only one of
	// two parameters should be specified.
	DomainId *string `mandatory:"false" json:"domainId"`
}

func (m CreateIntegrationInstanceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateIntegrationInstanceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateIntegrationInstanceDetailsIntegrationInstanceTypeEnum(string(m.IntegrationInstanceType)); !ok && m.IntegrationInstanceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IntegrationInstanceType: %s. Supported values are: %s.", m.IntegrationInstanceType, strings.Join(GetCreateIntegrationInstanceDetailsIntegrationInstanceTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingCreateIntegrationInstanceDetailsConsumptionModelEnum(string(m.ConsumptionModel)); !ok && m.ConsumptionModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConsumptionModel: %s. Supported values are: %s.", m.ConsumptionModel, strings.Join(GetCreateIntegrationInstanceDetailsConsumptionModelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateIntegrationInstanceDetailsShapeEnum(string(m.Shape)); !ok && m.Shape != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Shape: %s. Supported values are: %s.", m.Shape, strings.Join(GetCreateIntegrationInstanceDetailsShapeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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
		Shape                    CreateIntegrationInstanceDetailsShapeEnum                   `json:"shape"`
		DomainId                 *string                                                     `json:"domainId"`
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
	copy(m.AlternateCustomEndpoints, model.AlternateCustomEndpoints)
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

	m.Shape = model.Shape

	m.DomainId = model.DomainId

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
	CreateIntegrationInstanceDetailsIntegrationInstanceTypeStandard    CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnum = "STANDARD"
	CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnterprise  CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnum = "ENTERPRISE"
	CreateIntegrationInstanceDetailsIntegrationInstanceTypeStandardx   CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnum = "STANDARDX"
	CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnterprisex CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnum = "ENTERPRISEX"
)

var mappingCreateIntegrationInstanceDetailsIntegrationInstanceTypeEnum = map[string]CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnum{
	"STANDARD":    CreateIntegrationInstanceDetailsIntegrationInstanceTypeStandard,
	"ENTERPRISE":  CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnterprise,
	"STANDARDX":   CreateIntegrationInstanceDetailsIntegrationInstanceTypeStandardx,
	"ENTERPRISEX": CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnterprisex,
}

var mappingCreateIntegrationInstanceDetailsIntegrationInstanceTypeEnumLowerCase = map[string]CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnum{
	"standard":    CreateIntegrationInstanceDetailsIntegrationInstanceTypeStandard,
	"enterprise":  CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnterprise,
	"standardx":   CreateIntegrationInstanceDetailsIntegrationInstanceTypeStandardx,
	"enterprisex": CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnterprisex,
}

// GetCreateIntegrationInstanceDetailsIntegrationInstanceTypeEnumValues Enumerates the set of values for CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnum
func GetCreateIntegrationInstanceDetailsIntegrationInstanceTypeEnumValues() []CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnum {
	values := make([]CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnum, 0)
	for _, v := range mappingCreateIntegrationInstanceDetailsIntegrationInstanceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateIntegrationInstanceDetailsIntegrationInstanceTypeEnumStringValues Enumerates the set of values in String for CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnum
func GetCreateIntegrationInstanceDetailsIntegrationInstanceTypeEnumStringValues() []string {
	return []string{
		"STANDARD",
		"ENTERPRISE",
		"STANDARDX",
		"ENTERPRISEX",
	}
}

// GetMappingCreateIntegrationInstanceDetailsIntegrationInstanceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateIntegrationInstanceDetailsIntegrationInstanceTypeEnum(val string) (CreateIntegrationInstanceDetailsIntegrationInstanceTypeEnum, bool) {
	enum, ok := mappingCreateIntegrationInstanceDetailsIntegrationInstanceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateIntegrationInstanceDetailsConsumptionModelEnum Enum with underlying type: string
type CreateIntegrationInstanceDetailsConsumptionModelEnum string

// Set of constants representing the allowable values for CreateIntegrationInstanceDetailsConsumptionModelEnum
const (
	CreateIntegrationInstanceDetailsConsumptionModelUcm      CreateIntegrationInstanceDetailsConsumptionModelEnum = "UCM"
	CreateIntegrationInstanceDetailsConsumptionModelGov      CreateIntegrationInstanceDetailsConsumptionModelEnum = "GOV"
	CreateIntegrationInstanceDetailsConsumptionModelOic4saas CreateIntegrationInstanceDetailsConsumptionModelEnum = "OIC4SAAS"
)

var mappingCreateIntegrationInstanceDetailsConsumptionModelEnum = map[string]CreateIntegrationInstanceDetailsConsumptionModelEnum{
	"UCM":      CreateIntegrationInstanceDetailsConsumptionModelUcm,
	"GOV":      CreateIntegrationInstanceDetailsConsumptionModelGov,
	"OIC4SAAS": CreateIntegrationInstanceDetailsConsumptionModelOic4saas,
}

var mappingCreateIntegrationInstanceDetailsConsumptionModelEnumLowerCase = map[string]CreateIntegrationInstanceDetailsConsumptionModelEnum{
	"ucm":      CreateIntegrationInstanceDetailsConsumptionModelUcm,
	"gov":      CreateIntegrationInstanceDetailsConsumptionModelGov,
	"oic4saas": CreateIntegrationInstanceDetailsConsumptionModelOic4saas,
}

// GetCreateIntegrationInstanceDetailsConsumptionModelEnumValues Enumerates the set of values for CreateIntegrationInstanceDetailsConsumptionModelEnum
func GetCreateIntegrationInstanceDetailsConsumptionModelEnumValues() []CreateIntegrationInstanceDetailsConsumptionModelEnum {
	values := make([]CreateIntegrationInstanceDetailsConsumptionModelEnum, 0)
	for _, v := range mappingCreateIntegrationInstanceDetailsConsumptionModelEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateIntegrationInstanceDetailsConsumptionModelEnumStringValues Enumerates the set of values in String for CreateIntegrationInstanceDetailsConsumptionModelEnum
func GetCreateIntegrationInstanceDetailsConsumptionModelEnumStringValues() []string {
	return []string{
		"UCM",
		"GOV",
		"OIC4SAAS",
	}
}

// GetMappingCreateIntegrationInstanceDetailsConsumptionModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateIntegrationInstanceDetailsConsumptionModelEnum(val string) (CreateIntegrationInstanceDetailsConsumptionModelEnum, bool) {
	enum, ok := mappingCreateIntegrationInstanceDetailsConsumptionModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateIntegrationInstanceDetailsShapeEnum Enum with underlying type: string
type CreateIntegrationInstanceDetailsShapeEnum string

// Set of constants representing the allowable values for CreateIntegrationInstanceDetailsShapeEnum
const (
	CreateIntegrationInstanceDetailsShapeDevelopment CreateIntegrationInstanceDetailsShapeEnum = "DEVELOPMENT"
	CreateIntegrationInstanceDetailsShapeProduction  CreateIntegrationInstanceDetailsShapeEnum = "PRODUCTION"
)

var mappingCreateIntegrationInstanceDetailsShapeEnum = map[string]CreateIntegrationInstanceDetailsShapeEnum{
	"DEVELOPMENT": CreateIntegrationInstanceDetailsShapeDevelopment,
	"PRODUCTION":  CreateIntegrationInstanceDetailsShapeProduction,
}

var mappingCreateIntegrationInstanceDetailsShapeEnumLowerCase = map[string]CreateIntegrationInstanceDetailsShapeEnum{
	"development": CreateIntegrationInstanceDetailsShapeDevelopment,
	"production":  CreateIntegrationInstanceDetailsShapeProduction,
}

// GetCreateIntegrationInstanceDetailsShapeEnumValues Enumerates the set of values for CreateIntegrationInstanceDetailsShapeEnum
func GetCreateIntegrationInstanceDetailsShapeEnumValues() []CreateIntegrationInstanceDetailsShapeEnum {
	values := make([]CreateIntegrationInstanceDetailsShapeEnum, 0)
	for _, v := range mappingCreateIntegrationInstanceDetailsShapeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateIntegrationInstanceDetailsShapeEnumStringValues Enumerates the set of values in String for CreateIntegrationInstanceDetailsShapeEnum
func GetCreateIntegrationInstanceDetailsShapeEnumStringValues() []string {
	return []string{
		"DEVELOPMENT",
		"PRODUCTION",
	}
}

// GetMappingCreateIntegrationInstanceDetailsShapeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateIntegrationInstanceDetailsShapeEnum(val string) (CreateIntegrationInstanceDetailsShapeEnum, bool) {
	enum, ok := mappingCreateIntegrationInstanceDetailsShapeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
