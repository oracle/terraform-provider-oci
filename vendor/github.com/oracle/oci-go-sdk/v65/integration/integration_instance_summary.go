// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// IntegrationInstanceSummary Summary of the Integration Instance.
type IntegrationInstanceSummary struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// Integration Instance Identifier, can be renamed.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Standard or Enterprise type,
	// Oracle Integration Generation 2 uses ENTERPRISE and STANDARD,
	// Oracle Integration 3 uses ENTERPRISEX and STANDARDX
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

	// Additional details of lifecycleState or substates
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// An message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	StateMessage *string `mandatory:"false" json:"stateMessage"`

	// The Integration Instance Design Time URL
	InstanceDesignTimeUrl *string `mandatory:"false" json:"instanceDesignTimeUrl"`

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

	// Simple key-value pair that is applied without any predefined name,
	// type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to
	// namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// Shape
	Shape IntegrationInstanceSummaryShapeEnum `mandatory:"false" json:"shape,omitempty"`

	PrivateEndpointOutboundConnection OutboundConnection `mandatory:"false" json:"privateEndpointOutboundConnection"`

	// Is Disaster Recovery enabled for the integrationInstance
	IsDisasterRecoveryEnabled *bool `mandatory:"false" json:"isDisasterRecoveryEnabled"`

	// Data retention period set for given integration instance
	DataRetentionPeriod IntegrationInstanceSummaryDataRetentionPeriodEnum `mandatory:"false" json:"dataRetentionPeriod,omitempty"`
}

func (m IntegrationInstanceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IntegrationInstanceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingIntegrationInstanceSummaryIntegrationInstanceTypeEnum(string(m.IntegrationInstanceType)); !ok && m.IntegrationInstanceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IntegrationInstanceType: %s. Supported values are: %s.", m.IntegrationInstanceType, strings.Join(GetIntegrationInstanceSummaryIntegrationInstanceTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingIntegrationInstanceSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetIntegrationInstanceSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingIntegrationInstanceSummaryConsumptionModelEnum(string(m.ConsumptionModel)); !ok && m.ConsumptionModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConsumptionModel: %s. Supported values are: %s.", m.ConsumptionModel, strings.Join(GetIntegrationInstanceSummaryConsumptionModelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingIntegrationInstanceSummaryShapeEnum(string(m.Shape)); !ok && m.Shape != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Shape: %s. Supported values are: %s.", m.Shape, strings.Join(GetIntegrationInstanceSummaryShapeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingIntegrationInstanceSummaryDataRetentionPeriodEnum(string(m.DataRetentionPeriod)); !ok && m.DataRetentionPeriod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataRetentionPeriod: %s. Supported values are: %s.", m.DataRetentionPeriod, strings.Join(GetIntegrationInstanceSummaryDataRetentionPeriodEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *IntegrationInstanceSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TimeCreated                       *common.SDKTime                                       `json:"timeCreated"`
		TimeUpdated                       *common.SDKTime                                       `json:"timeUpdated"`
		LifecycleState                    IntegrationInstanceSummaryLifecycleStateEnum          `json:"lifecycleState"`
		LifecycleDetails                  *string                                               `json:"lifecycleDetails"`
		StateMessage                      *string                                               `json:"stateMessage"`
		InstanceDesignTimeUrl             *string                                               `json:"instanceDesignTimeUrl"`
		IsFileServerEnabled               *bool                                                 `json:"isFileServerEnabled"`
		IsVisualBuilderEnabled            *bool                                                 `json:"isVisualBuilderEnabled"`
		CustomEndpoint                    *CustomEndpointDetails                                `json:"customEndpoint"`
		AlternateCustomEndpoints          []CustomEndpointDetails                               `json:"alternateCustomEndpoints"`
		ConsumptionModel                  IntegrationInstanceSummaryConsumptionModelEnum        `json:"consumptionModel"`
		NetworkEndpointDetails            networkendpointdetails                                `json:"networkEndpointDetails"`
		FreeformTags                      map[string]string                                     `json:"freeformTags"`
		DefinedTags                       map[string]map[string]interface{}                     `json:"definedTags"`
		SystemTags                        map[string]map[string]interface{}                     `json:"systemTags"`
		Shape                             IntegrationInstanceSummaryShapeEnum                   `json:"shape"`
		PrivateEndpointOutboundConnection outboundconnection                                    `json:"privateEndpointOutboundConnection"`
		IsDisasterRecoveryEnabled         *bool                                                 `json:"isDisasterRecoveryEnabled"`
		DataRetentionPeriod               IntegrationInstanceSummaryDataRetentionPeriodEnum     `json:"dataRetentionPeriod"`
		Id                                *string                                               `json:"id"`
		DisplayName                       *string                                               `json:"displayName"`
		CompartmentId                     *string                                               `json:"compartmentId"`
		IntegrationInstanceType           IntegrationInstanceSummaryIntegrationInstanceTypeEnum `json:"integrationInstanceType"`
		IsByol                            *bool                                                 `json:"isByol"`
		InstanceUrl                       *string                                               `json:"instanceUrl"`
		MessagePacks                      *int                                                  `json:"messagePacks"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

	m.LifecycleDetails = model.LifecycleDetails

	m.StateMessage = model.StateMessage

	m.InstanceDesignTimeUrl = model.InstanceDesignTimeUrl

	m.IsFileServerEnabled = model.IsFileServerEnabled

	m.IsVisualBuilderEnabled = model.IsVisualBuilderEnabled

	m.CustomEndpoint = model.CustomEndpoint

	m.AlternateCustomEndpoints = make([]CustomEndpointDetails, len(model.AlternateCustomEndpoints))
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

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Shape = model.Shape

	nn, e = model.PrivateEndpointOutboundConnection.UnmarshalPolymorphicJSON(model.PrivateEndpointOutboundConnection.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.PrivateEndpointOutboundConnection = nn.(OutboundConnection)
	} else {
		m.PrivateEndpointOutboundConnection = nil
	}

	m.IsDisasterRecoveryEnabled = model.IsDisasterRecoveryEnabled

	m.DataRetentionPeriod = model.DataRetentionPeriod

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
	IntegrationInstanceSummaryIntegrationInstanceTypeStandard    IntegrationInstanceSummaryIntegrationInstanceTypeEnum = "STANDARD"
	IntegrationInstanceSummaryIntegrationInstanceTypeEnterprise  IntegrationInstanceSummaryIntegrationInstanceTypeEnum = "ENTERPRISE"
	IntegrationInstanceSummaryIntegrationInstanceTypeStandardx   IntegrationInstanceSummaryIntegrationInstanceTypeEnum = "STANDARDX"
	IntegrationInstanceSummaryIntegrationInstanceTypeEnterprisex IntegrationInstanceSummaryIntegrationInstanceTypeEnum = "ENTERPRISEX"
	IntegrationInstanceSummaryIntegrationInstanceTypeHealthcare  IntegrationInstanceSummaryIntegrationInstanceTypeEnum = "HEALTHCARE"
)

var mappingIntegrationInstanceSummaryIntegrationInstanceTypeEnum = map[string]IntegrationInstanceSummaryIntegrationInstanceTypeEnum{
	"STANDARD":    IntegrationInstanceSummaryIntegrationInstanceTypeStandard,
	"ENTERPRISE":  IntegrationInstanceSummaryIntegrationInstanceTypeEnterprise,
	"STANDARDX":   IntegrationInstanceSummaryIntegrationInstanceTypeStandardx,
	"ENTERPRISEX": IntegrationInstanceSummaryIntegrationInstanceTypeEnterprisex,
	"HEALTHCARE":  IntegrationInstanceSummaryIntegrationInstanceTypeHealthcare,
}

var mappingIntegrationInstanceSummaryIntegrationInstanceTypeEnumLowerCase = map[string]IntegrationInstanceSummaryIntegrationInstanceTypeEnum{
	"standard":    IntegrationInstanceSummaryIntegrationInstanceTypeStandard,
	"enterprise":  IntegrationInstanceSummaryIntegrationInstanceTypeEnterprise,
	"standardx":   IntegrationInstanceSummaryIntegrationInstanceTypeStandardx,
	"enterprisex": IntegrationInstanceSummaryIntegrationInstanceTypeEnterprisex,
	"healthcare":  IntegrationInstanceSummaryIntegrationInstanceTypeHealthcare,
}

// GetIntegrationInstanceSummaryIntegrationInstanceTypeEnumValues Enumerates the set of values for IntegrationInstanceSummaryIntegrationInstanceTypeEnum
func GetIntegrationInstanceSummaryIntegrationInstanceTypeEnumValues() []IntegrationInstanceSummaryIntegrationInstanceTypeEnum {
	values := make([]IntegrationInstanceSummaryIntegrationInstanceTypeEnum, 0)
	for _, v := range mappingIntegrationInstanceSummaryIntegrationInstanceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetIntegrationInstanceSummaryIntegrationInstanceTypeEnumStringValues Enumerates the set of values in String for IntegrationInstanceSummaryIntegrationInstanceTypeEnum
func GetIntegrationInstanceSummaryIntegrationInstanceTypeEnumStringValues() []string {
	return []string{
		"STANDARD",
		"ENTERPRISE",
		"STANDARDX",
		"ENTERPRISEX",
		"HEALTHCARE",
	}
}

// GetMappingIntegrationInstanceSummaryIntegrationInstanceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIntegrationInstanceSummaryIntegrationInstanceTypeEnum(val string) (IntegrationInstanceSummaryIntegrationInstanceTypeEnum, bool) {
	enum, ok := mappingIntegrationInstanceSummaryIntegrationInstanceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
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

var mappingIntegrationInstanceSummaryLifecycleStateEnum = map[string]IntegrationInstanceSummaryLifecycleStateEnum{
	"CREATING": IntegrationInstanceSummaryLifecycleStateCreating,
	"UPDATING": IntegrationInstanceSummaryLifecycleStateUpdating,
	"ACTIVE":   IntegrationInstanceSummaryLifecycleStateActive,
	"INACTIVE": IntegrationInstanceSummaryLifecycleStateInactive,
	"DELETING": IntegrationInstanceSummaryLifecycleStateDeleting,
	"DELETED":  IntegrationInstanceSummaryLifecycleStateDeleted,
	"FAILED":   IntegrationInstanceSummaryLifecycleStateFailed,
}

var mappingIntegrationInstanceSummaryLifecycleStateEnumLowerCase = map[string]IntegrationInstanceSummaryLifecycleStateEnum{
	"creating": IntegrationInstanceSummaryLifecycleStateCreating,
	"updating": IntegrationInstanceSummaryLifecycleStateUpdating,
	"active":   IntegrationInstanceSummaryLifecycleStateActive,
	"inactive": IntegrationInstanceSummaryLifecycleStateInactive,
	"deleting": IntegrationInstanceSummaryLifecycleStateDeleting,
	"deleted":  IntegrationInstanceSummaryLifecycleStateDeleted,
	"failed":   IntegrationInstanceSummaryLifecycleStateFailed,
}

// GetIntegrationInstanceSummaryLifecycleStateEnumValues Enumerates the set of values for IntegrationInstanceSummaryLifecycleStateEnum
func GetIntegrationInstanceSummaryLifecycleStateEnumValues() []IntegrationInstanceSummaryLifecycleStateEnum {
	values := make([]IntegrationInstanceSummaryLifecycleStateEnum, 0)
	for _, v := range mappingIntegrationInstanceSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetIntegrationInstanceSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for IntegrationInstanceSummaryLifecycleStateEnum
func GetIntegrationInstanceSummaryLifecycleStateEnumStringValues() []string {
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

// GetMappingIntegrationInstanceSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIntegrationInstanceSummaryLifecycleStateEnum(val string) (IntegrationInstanceSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingIntegrationInstanceSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// IntegrationInstanceSummaryConsumptionModelEnum Enum with underlying type: string
type IntegrationInstanceSummaryConsumptionModelEnum string

// Set of constants representing the allowable values for IntegrationInstanceSummaryConsumptionModelEnum
const (
	IntegrationInstanceSummaryConsumptionModelUcm      IntegrationInstanceSummaryConsumptionModelEnum = "UCM"
	IntegrationInstanceSummaryConsumptionModelGov      IntegrationInstanceSummaryConsumptionModelEnum = "GOV"
	IntegrationInstanceSummaryConsumptionModelOic4saas IntegrationInstanceSummaryConsumptionModelEnum = "OIC4SAAS"
)

var mappingIntegrationInstanceSummaryConsumptionModelEnum = map[string]IntegrationInstanceSummaryConsumptionModelEnum{
	"UCM":      IntegrationInstanceSummaryConsumptionModelUcm,
	"GOV":      IntegrationInstanceSummaryConsumptionModelGov,
	"OIC4SAAS": IntegrationInstanceSummaryConsumptionModelOic4saas,
}

var mappingIntegrationInstanceSummaryConsumptionModelEnumLowerCase = map[string]IntegrationInstanceSummaryConsumptionModelEnum{
	"ucm":      IntegrationInstanceSummaryConsumptionModelUcm,
	"gov":      IntegrationInstanceSummaryConsumptionModelGov,
	"oic4saas": IntegrationInstanceSummaryConsumptionModelOic4saas,
}

// GetIntegrationInstanceSummaryConsumptionModelEnumValues Enumerates the set of values for IntegrationInstanceSummaryConsumptionModelEnum
func GetIntegrationInstanceSummaryConsumptionModelEnumValues() []IntegrationInstanceSummaryConsumptionModelEnum {
	values := make([]IntegrationInstanceSummaryConsumptionModelEnum, 0)
	for _, v := range mappingIntegrationInstanceSummaryConsumptionModelEnum {
		values = append(values, v)
	}
	return values
}

// GetIntegrationInstanceSummaryConsumptionModelEnumStringValues Enumerates the set of values in String for IntegrationInstanceSummaryConsumptionModelEnum
func GetIntegrationInstanceSummaryConsumptionModelEnumStringValues() []string {
	return []string{
		"UCM",
		"GOV",
		"OIC4SAAS",
	}
}

// GetMappingIntegrationInstanceSummaryConsumptionModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIntegrationInstanceSummaryConsumptionModelEnum(val string) (IntegrationInstanceSummaryConsumptionModelEnum, bool) {
	enum, ok := mappingIntegrationInstanceSummaryConsumptionModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// IntegrationInstanceSummaryShapeEnum Enum with underlying type: string
type IntegrationInstanceSummaryShapeEnum string

// Set of constants representing the allowable values for IntegrationInstanceSummaryShapeEnum
const (
	IntegrationInstanceSummaryShapeDevelopment IntegrationInstanceSummaryShapeEnum = "DEVELOPMENT"
	IntegrationInstanceSummaryShapeProduction  IntegrationInstanceSummaryShapeEnum = "PRODUCTION"
)

var mappingIntegrationInstanceSummaryShapeEnum = map[string]IntegrationInstanceSummaryShapeEnum{
	"DEVELOPMENT": IntegrationInstanceSummaryShapeDevelopment,
	"PRODUCTION":  IntegrationInstanceSummaryShapeProduction,
}

var mappingIntegrationInstanceSummaryShapeEnumLowerCase = map[string]IntegrationInstanceSummaryShapeEnum{
	"development": IntegrationInstanceSummaryShapeDevelopment,
	"production":  IntegrationInstanceSummaryShapeProduction,
}

// GetIntegrationInstanceSummaryShapeEnumValues Enumerates the set of values for IntegrationInstanceSummaryShapeEnum
func GetIntegrationInstanceSummaryShapeEnumValues() []IntegrationInstanceSummaryShapeEnum {
	values := make([]IntegrationInstanceSummaryShapeEnum, 0)
	for _, v := range mappingIntegrationInstanceSummaryShapeEnum {
		values = append(values, v)
	}
	return values
}

// GetIntegrationInstanceSummaryShapeEnumStringValues Enumerates the set of values in String for IntegrationInstanceSummaryShapeEnum
func GetIntegrationInstanceSummaryShapeEnumStringValues() []string {
	return []string{
		"DEVELOPMENT",
		"PRODUCTION",
	}
}

// GetMappingIntegrationInstanceSummaryShapeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIntegrationInstanceSummaryShapeEnum(val string) (IntegrationInstanceSummaryShapeEnum, bool) {
	enum, ok := mappingIntegrationInstanceSummaryShapeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// IntegrationInstanceSummaryDataRetentionPeriodEnum Enum with underlying type: string
type IntegrationInstanceSummaryDataRetentionPeriodEnum string

// Set of constants representing the allowable values for IntegrationInstanceSummaryDataRetentionPeriodEnum
const (
	IntegrationInstanceSummaryDataRetentionPeriod1 IntegrationInstanceSummaryDataRetentionPeriodEnum = "MONTHS_1"
	IntegrationInstanceSummaryDataRetentionPeriod3 IntegrationInstanceSummaryDataRetentionPeriodEnum = "MONTHS_3"
	IntegrationInstanceSummaryDataRetentionPeriod6 IntegrationInstanceSummaryDataRetentionPeriodEnum = "MONTHS_6"
)

var mappingIntegrationInstanceSummaryDataRetentionPeriodEnum = map[string]IntegrationInstanceSummaryDataRetentionPeriodEnum{
	"MONTHS_1": IntegrationInstanceSummaryDataRetentionPeriod1,
	"MONTHS_3": IntegrationInstanceSummaryDataRetentionPeriod3,
	"MONTHS_6": IntegrationInstanceSummaryDataRetentionPeriod6,
}

var mappingIntegrationInstanceSummaryDataRetentionPeriodEnumLowerCase = map[string]IntegrationInstanceSummaryDataRetentionPeriodEnum{
	"months_1": IntegrationInstanceSummaryDataRetentionPeriod1,
	"months_3": IntegrationInstanceSummaryDataRetentionPeriod3,
	"months_6": IntegrationInstanceSummaryDataRetentionPeriod6,
}

// GetIntegrationInstanceSummaryDataRetentionPeriodEnumValues Enumerates the set of values for IntegrationInstanceSummaryDataRetentionPeriodEnum
func GetIntegrationInstanceSummaryDataRetentionPeriodEnumValues() []IntegrationInstanceSummaryDataRetentionPeriodEnum {
	values := make([]IntegrationInstanceSummaryDataRetentionPeriodEnum, 0)
	for _, v := range mappingIntegrationInstanceSummaryDataRetentionPeriodEnum {
		values = append(values, v)
	}
	return values
}

// GetIntegrationInstanceSummaryDataRetentionPeriodEnumStringValues Enumerates the set of values in String for IntegrationInstanceSummaryDataRetentionPeriodEnum
func GetIntegrationInstanceSummaryDataRetentionPeriodEnumStringValues() []string {
	return []string{
		"MONTHS_1",
		"MONTHS_3",
		"MONTHS_6",
	}
}

// GetMappingIntegrationInstanceSummaryDataRetentionPeriodEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIntegrationInstanceSummaryDataRetentionPeriodEnum(val string) (IntegrationInstanceSummaryDataRetentionPeriodEnum, bool) {
	enum, ok := mappingIntegrationInstanceSummaryDataRetentionPeriodEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
