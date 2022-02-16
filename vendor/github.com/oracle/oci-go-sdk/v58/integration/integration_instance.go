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
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// IntegrationInstance Description of Integration Instance.
type IntegrationInstance struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// Integration Instance Identifier, can be renamed.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Standard or Enterprise type
	IntegrationInstanceType IntegrationInstanceIntegrationInstanceTypeEnum `mandatory:"true" json:"integrationInstanceType"`

	// Bring your own license.
	IsByol *bool `mandatory:"true" json:"isByol"`

	// The Integration Instance URL.
	InstanceUrl *string `mandatory:"true" json:"instanceUrl"`

	// The number of configured message packs (if any)
	MessagePacks *int `mandatory:"true" json:"messagePacks"`

	// The time the the IntegrationInstance was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the IntegrationInstance was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the integration instance.
	LifecycleState IntegrationInstanceLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// An message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	StateMessage *string `mandatory:"false" json:"stateMessage"`

	// Simple key-value pair that is applied without any predefined name,
	// type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to
	// namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The file server is enabled or not.
	IsFileServerEnabled *bool `mandatory:"false" json:"isFileServerEnabled"`

	// VisualBuilder is enabled or not.
	IsVisualBuilderEnabled *bool `mandatory:"false" json:"isVisualBuilderEnabled"`

	CustomEndpoint *CustomEndpointDetails `mandatory:"false" json:"customEndpoint"`

	// A list of alternate custom endpoints used for the integration instance URL.
	AlternateCustomEndpoints []CustomEndpointDetails `mandatory:"false" json:"alternateCustomEndpoints"`

	// The entitlement used for billing purposes.
	ConsumptionModel IntegrationInstanceConsumptionModelEnum `mandatory:"false" json:"consumptionModel,omitempty"`

	NetworkEndpointDetails NetworkEndpointDetails `mandatory:"false" json:"networkEndpointDetails"`
}

func (m IntegrationInstance) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IntegrationInstance) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingIntegrationInstanceIntegrationInstanceTypeEnum(string(m.IntegrationInstanceType)); !ok && m.IntegrationInstanceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IntegrationInstanceType: %s. Supported values are: %s.", m.IntegrationInstanceType, strings.Join(GetIntegrationInstanceIntegrationInstanceTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingIntegrationInstanceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetIntegrationInstanceLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingIntegrationInstanceConsumptionModelEnum(string(m.ConsumptionModel)); !ok && m.ConsumptionModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConsumptionModel: %s. Supported values are: %s.", m.ConsumptionModel, strings.Join(GetIntegrationInstanceConsumptionModelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *IntegrationInstance) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TimeCreated              *common.SDKTime                                `json:"timeCreated"`
		TimeUpdated              *common.SDKTime                                `json:"timeUpdated"`
		LifecycleState           IntegrationInstanceLifecycleStateEnum          `json:"lifecycleState"`
		StateMessage             *string                                        `json:"stateMessage"`
		FreeformTags             map[string]string                              `json:"freeformTags"`
		DefinedTags              map[string]map[string]interface{}              `json:"definedTags"`
		IsFileServerEnabled      *bool                                          `json:"isFileServerEnabled"`
		IsVisualBuilderEnabled   *bool                                          `json:"isVisualBuilderEnabled"`
		CustomEndpoint           *CustomEndpointDetails                         `json:"customEndpoint"`
		AlternateCustomEndpoints []CustomEndpointDetails                        `json:"alternateCustomEndpoints"`
		ConsumptionModel         IntegrationInstanceConsumptionModelEnum        `json:"consumptionModel"`
		NetworkEndpointDetails   networkendpointdetails                         `json:"networkEndpointDetails"`
		Id                       *string                                        `json:"id"`
		DisplayName              *string                                        `json:"displayName"`
		CompartmentId            *string                                        `json:"compartmentId"`
		IntegrationInstanceType  IntegrationInstanceIntegrationInstanceTypeEnum `json:"integrationInstanceType"`
		IsByol                   *bool                                          `json:"isByol"`
		InstanceUrl              *string                                        `json:"instanceUrl"`
		MessagePacks             *int                                           `json:"messagePacks"`
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

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

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

// IntegrationInstanceIntegrationInstanceTypeEnum Enum with underlying type: string
type IntegrationInstanceIntegrationInstanceTypeEnum string

// Set of constants representing the allowable values for IntegrationInstanceIntegrationInstanceTypeEnum
const (
	IntegrationInstanceIntegrationInstanceTypeStandard   IntegrationInstanceIntegrationInstanceTypeEnum = "STANDARD"
	IntegrationInstanceIntegrationInstanceTypeEnterprise IntegrationInstanceIntegrationInstanceTypeEnum = "ENTERPRISE"
)

var mappingIntegrationInstanceIntegrationInstanceTypeEnum = map[string]IntegrationInstanceIntegrationInstanceTypeEnum{
	"STANDARD":   IntegrationInstanceIntegrationInstanceTypeStandard,
	"ENTERPRISE": IntegrationInstanceIntegrationInstanceTypeEnterprise,
}

// GetIntegrationInstanceIntegrationInstanceTypeEnumValues Enumerates the set of values for IntegrationInstanceIntegrationInstanceTypeEnum
func GetIntegrationInstanceIntegrationInstanceTypeEnumValues() []IntegrationInstanceIntegrationInstanceTypeEnum {
	values := make([]IntegrationInstanceIntegrationInstanceTypeEnum, 0)
	for _, v := range mappingIntegrationInstanceIntegrationInstanceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetIntegrationInstanceIntegrationInstanceTypeEnumStringValues Enumerates the set of values in String for IntegrationInstanceIntegrationInstanceTypeEnum
func GetIntegrationInstanceIntegrationInstanceTypeEnumStringValues() []string {
	return []string{
		"STANDARD",
		"ENTERPRISE",
	}
}

// GetMappingIntegrationInstanceIntegrationInstanceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIntegrationInstanceIntegrationInstanceTypeEnum(val string) (IntegrationInstanceIntegrationInstanceTypeEnum, bool) {
	mappingIntegrationInstanceIntegrationInstanceTypeEnumIgnoreCase := make(map[string]IntegrationInstanceIntegrationInstanceTypeEnum)
	for k, v := range mappingIntegrationInstanceIntegrationInstanceTypeEnum {
		mappingIntegrationInstanceIntegrationInstanceTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingIntegrationInstanceIntegrationInstanceTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// IntegrationInstanceLifecycleStateEnum Enum with underlying type: string
type IntegrationInstanceLifecycleStateEnum string

// Set of constants representing the allowable values for IntegrationInstanceLifecycleStateEnum
const (
	IntegrationInstanceLifecycleStateCreating IntegrationInstanceLifecycleStateEnum = "CREATING"
	IntegrationInstanceLifecycleStateUpdating IntegrationInstanceLifecycleStateEnum = "UPDATING"
	IntegrationInstanceLifecycleStateActive   IntegrationInstanceLifecycleStateEnum = "ACTIVE"
	IntegrationInstanceLifecycleStateInactive IntegrationInstanceLifecycleStateEnum = "INACTIVE"
	IntegrationInstanceLifecycleStateDeleting IntegrationInstanceLifecycleStateEnum = "DELETING"
	IntegrationInstanceLifecycleStateDeleted  IntegrationInstanceLifecycleStateEnum = "DELETED"
	IntegrationInstanceLifecycleStateFailed   IntegrationInstanceLifecycleStateEnum = "FAILED"
)

var mappingIntegrationInstanceLifecycleStateEnum = map[string]IntegrationInstanceLifecycleStateEnum{
	"CREATING": IntegrationInstanceLifecycleStateCreating,
	"UPDATING": IntegrationInstanceLifecycleStateUpdating,
	"ACTIVE":   IntegrationInstanceLifecycleStateActive,
	"INACTIVE": IntegrationInstanceLifecycleStateInactive,
	"DELETING": IntegrationInstanceLifecycleStateDeleting,
	"DELETED":  IntegrationInstanceLifecycleStateDeleted,
	"FAILED":   IntegrationInstanceLifecycleStateFailed,
}

// GetIntegrationInstanceLifecycleStateEnumValues Enumerates the set of values for IntegrationInstanceLifecycleStateEnum
func GetIntegrationInstanceLifecycleStateEnumValues() []IntegrationInstanceLifecycleStateEnum {
	values := make([]IntegrationInstanceLifecycleStateEnum, 0)
	for _, v := range mappingIntegrationInstanceLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetIntegrationInstanceLifecycleStateEnumStringValues Enumerates the set of values in String for IntegrationInstanceLifecycleStateEnum
func GetIntegrationInstanceLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingIntegrationInstanceLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIntegrationInstanceLifecycleStateEnum(val string) (IntegrationInstanceLifecycleStateEnum, bool) {
	mappingIntegrationInstanceLifecycleStateEnumIgnoreCase := make(map[string]IntegrationInstanceLifecycleStateEnum)
	for k, v := range mappingIntegrationInstanceLifecycleStateEnum {
		mappingIntegrationInstanceLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingIntegrationInstanceLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// IntegrationInstanceConsumptionModelEnum Enum with underlying type: string
type IntegrationInstanceConsumptionModelEnum string

// Set of constants representing the allowable values for IntegrationInstanceConsumptionModelEnum
const (
	IntegrationInstanceConsumptionModelUcm      IntegrationInstanceConsumptionModelEnum = "UCM"
	IntegrationInstanceConsumptionModelGov      IntegrationInstanceConsumptionModelEnum = "GOV"
	IntegrationInstanceConsumptionModelOic4saas IntegrationInstanceConsumptionModelEnum = "OIC4SAAS"
)

var mappingIntegrationInstanceConsumptionModelEnum = map[string]IntegrationInstanceConsumptionModelEnum{
	"UCM":      IntegrationInstanceConsumptionModelUcm,
	"GOV":      IntegrationInstanceConsumptionModelGov,
	"OIC4SAAS": IntegrationInstanceConsumptionModelOic4saas,
}

// GetIntegrationInstanceConsumptionModelEnumValues Enumerates the set of values for IntegrationInstanceConsumptionModelEnum
func GetIntegrationInstanceConsumptionModelEnumValues() []IntegrationInstanceConsumptionModelEnum {
	values := make([]IntegrationInstanceConsumptionModelEnum, 0)
	for _, v := range mappingIntegrationInstanceConsumptionModelEnum {
		values = append(values, v)
	}
	return values
}

// GetIntegrationInstanceConsumptionModelEnumStringValues Enumerates the set of values in String for IntegrationInstanceConsumptionModelEnum
func GetIntegrationInstanceConsumptionModelEnumStringValues() []string {
	return []string{
		"UCM",
		"GOV",
		"OIC4SAAS",
	}
}

// GetMappingIntegrationInstanceConsumptionModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIntegrationInstanceConsumptionModelEnum(val string) (IntegrationInstanceConsumptionModelEnum, bool) {
	mappingIntegrationInstanceConsumptionModelEnumIgnoreCase := make(map[string]IntegrationInstanceConsumptionModelEnum)
	for k, v := range mappingIntegrationInstanceConsumptionModelEnum {
		mappingIntegrationInstanceConsumptionModelEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingIntegrationInstanceConsumptionModelEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
