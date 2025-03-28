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

// IntegrationInstance Description of Integration Instance.
type IntegrationInstance struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// Integration Instance Identifier, can be renamed.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Standard or Enterprise type,
	// Oracle Integration Generation 2 uses ENTERPRISE and STANDARD,
	// Oracle Integration 3 uses ENTERPRISEX and STANDARDX
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

	// Additional details of lifecycleState or substates
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

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

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The Integration Instance Design Time URL
	InstanceDesignTimeUrl *string `mandatory:"false" json:"instanceDesignTimeUrl"`

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

	IdcsInfo *IdcsInfoDetails `mandatory:"false" json:"idcsInfo"`

	// A list of associated attachments to other services
	Attachments []AttachmentDetails `mandatory:"false" json:"attachments"`

	// Shape
	Shape IntegrationInstanceShapeEnum `mandatory:"false" json:"shape,omitempty"`

	PrivateEndpointOutboundConnection OutboundConnection `mandatory:"false" json:"privateEndpointOutboundConnection"`

	// Is Disaster Recovery enabled for the integrationInstance
	IsDisasterRecoveryEnabled *bool `mandatory:"false" json:"isDisasterRecoveryEnabled"`

	DisasterRecoveryDetails *DisasterRecoveryDetails `mandatory:"false" json:"disasterRecoveryDetails"`

	// Data retention period set for given integration instance
	DataRetentionPeriod IntegrationInstanceDataRetentionPeriodEnum `mandatory:"false" json:"dataRetentionPeriod,omitempty"`
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
	if _, ok := GetMappingIntegrationInstanceShapeEnum(string(m.Shape)); !ok && m.Shape != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Shape: %s. Supported values are: %s.", m.Shape, strings.Join(GetIntegrationInstanceShapeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingIntegrationInstanceDataRetentionPeriodEnum(string(m.DataRetentionPeriod)); !ok && m.DataRetentionPeriod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataRetentionPeriod: %s. Supported values are: %s.", m.DataRetentionPeriod, strings.Join(GetIntegrationInstanceDataRetentionPeriodEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *IntegrationInstance) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TimeCreated                       *common.SDKTime                                `json:"timeCreated"`
		TimeUpdated                       *common.SDKTime                                `json:"timeUpdated"`
		LifecycleState                    IntegrationInstanceLifecycleStateEnum          `json:"lifecycleState"`
		LifecycleDetails                  *string                                        `json:"lifecycleDetails"`
		StateMessage                      *string                                        `json:"stateMessage"`
		FreeformTags                      map[string]string                              `json:"freeformTags"`
		DefinedTags                       map[string]map[string]interface{}              `json:"definedTags"`
		SystemTags                        map[string]map[string]interface{}              `json:"systemTags"`
		InstanceDesignTimeUrl             *string                                        `json:"instanceDesignTimeUrl"`
		IsFileServerEnabled               *bool                                          `json:"isFileServerEnabled"`
		IsVisualBuilderEnabled            *bool                                          `json:"isVisualBuilderEnabled"`
		CustomEndpoint                    *CustomEndpointDetails                         `json:"customEndpoint"`
		AlternateCustomEndpoints          []CustomEndpointDetails                        `json:"alternateCustomEndpoints"`
		ConsumptionModel                  IntegrationInstanceConsumptionModelEnum        `json:"consumptionModel"`
		NetworkEndpointDetails            networkendpointdetails                         `json:"networkEndpointDetails"`
		IdcsInfo                          *IdcsInfoDetails                               `json:"idcsInfo"`
		Attachments                       []AttachmentDetails                            `json:"attachments"`
		Shape                             IntegrationInstanceShapeEnum                   `json:"shape"`
		PrivateEndpointOutboundConnection outboundconnection                             `json:"privateEndpointOutboundConnection"`
		IsDisasterRecoveryEnabled         *bool                                          `json:"isDisasterRecoveryEnabled"`
		DisasterRecoveryDetails           *DisasterRecoveryDetails                       `json:"disasterRecoveryDetails"`
		DataRetentionPeriod               IntegrationInstanceDataRetentionPeriodEnum     `json:"dataRetentionPeriod"`
		Id                                *string                                        `json:"id"`
		DisplayName                       *string                                        `json:"displayName"`
		CompartmentId                     *string                                        `json:"compartmentId"`
		IntegrationInstanceType           IntegrationInstanceIntegrationInstanceTypeEnum `json:"integrationInstanceType"`
		IsByol                            *bool                                          `json:"isByol"`
		InstanceUrl                       *string                                        `json:"instanceUrl"`
		MessagePacks                      *int                                           `json:"messagePacks"`
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

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

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

	m.IdcsInfo = model.IdcsInfo

	m.Attachments = make([]AttachmentDetails, len(model.Attachments))
	copy(m.Attachments, model.Attachments)
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

	m.DisasterRecoveryDetails = model.DisasterRecoveryDetails

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

// IntegrationInstanceIntegrationInstanceTypeEnum Enum with underlying type: string
type IntegrationInstanceIntegrationInstanceTypeEnum string

// Set of constants representing the allowable values for IntegrationInstanceIntegrationInstanceTypeEnum
const (
	IntegrationInstanceIntegrationInstanceTypeStandard    IntegrationInstanceIntegrationInstanceTypeEnum = "STANDARD"
	IntegrationInstanceIntegrationInstanceTypeEnterprise  IntegrationInstanceIntegrationInstanceTypeEnum = "ENTERPRISE"
	IntegrationInstanceIntegrationInstanceTypeStandardx   IntegrationInstanceIntegrationInstanceTypeEnum = "STANDARDX"
	IntegrationInstanceIntegrationInstanceTypeEnterprisex IntegrationInstanceIntegrationInstanceTypeEnum = "ENTERPRISEX"
	IntegrationInstanceIntegrationInstanceTypeHealthcare  IntegrationInstanceIntegrationInstanceTypeEnum = "HEALTHCARE"
)

var mappingIntegrationInstanceIntegrationInstanceTypeEnum = map[string]IntegrationInstanceIntegrationInstanceTypeEnum{
	"STANDARD":    IntegrationInstanceIntegrationInstanceTypeStandard,
	"ENTERPRISE":  IntegrationInstanceIntegrationInstanceTypeEnterprise,
	"STANDARDX":   IntegrationInstanceIntegrationInstanceTypeStandardx,
	"ENTERPRISEX": IntegrationInstanceIntegrationInstanceTypeEnterprisex,
	"HEALTHCARE":  IntegrationInstanceIntegrationInstanceTypeHealthcare,
}

var mappingIntegrationInstanceIntegrationInstanceTypeEnumLowerCase = map[string]IntegrationInstanceIntegrationInstanceTypeEnum{
	"standard":    IntegrationInstanceIntegrationInstanceTypeStandard,
	"enterprise":  IntegrationInstanceIntegrationInstanceTypeEnterprise,
	"standardx":   IntegrationInstanceIntegrationInstanceTypeStandardx,
	"enterprisex": IntegrationInstanceIntegrationInstanceTypeEnterprisex,
	"healthcare":  IntegrationInstanceIntegrationInstanceTypeHealthcare,
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
		"STANDARDX",
		"ENTERPRISEX",
		"HEALTHCARE",
	}
}

// GetMappingIntegrationInstanceIntegrationInstanceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIntegrationInstanceIntegrationInstanceTypeEnum(val string) (IntegrationInstanceIntegrationInstanceTypeEnum, bool) {
	enum, ok := mappingIntegrationInstanceIntegrationInstanceTypeEnumLowerCase[strings.ToLower(val)]
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

var mappingIntegrationInstanceLifecycleStateEnumLowerCase = map[string]IntegrationInstanceLifecycleStateEnum{
	"creating": IntegrationInstanceLifecycleStateCreating,
	"updating": IntegrationInstanceLifecycleStateUpdating,
	"active":   IntegrationInstanceLifecycleStateActive,
	"inactive": IntegrationInstanceLifecycleStateInactive,
	"deleting": IntegrationInstanceLifecycleStateDeleting,
	"deleted":  IntegrationInstanceLifecycleStateDeleted,
	"failed":   IntegrationInstanceLifecycleStateFailed,
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
	enum, ok := mappingIntegrationInstanceLifecycleStateEnumLowerCase[strings.ToLower(val)]
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

var mappingIntegrationInstanceConsumptionModelEnumLowerCase = map[string]IntegrationInstanceConsumptionModelEnum{
	"ucm":      IntegrationInstanceConsumptionModelUcm,
	"gov":      IntegrationInstanceConsumptionModelGov,
	"oic4saas": IntegrationInstanceConsumptionModelOic4saas,
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
	enum, ok := mappingIntegrationInstanceConsumptionModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// IntegrationInstanceShapeEnum Enum with underlying type: string
type IntegrationInstanceShapeEnum string

// Set of constants representing the allowable values for IntegrationInstanceShapeEnum
const (
	IntegrationInstanceShapeDevelopment IntegrationInstanceShapeEnum = "DEVELOPMENT"
	IntegrationInstanceShapeProduction  IntegrationInstanceShapeEnum = "PRODUCTION"
)

var mappingIntegrationInstanceShapeEnum = map[string]IntegrationInstanceShapeEnum{
	"DEVELOPMENT": IntegrationInstanceShapeDevelopment,
	"PRODUCTION":  IntegrationInstanceShapeProduction,
}

var mappingIntegrationInstanceShapeEnumLowerCase = map[string]IntegrationInstanceShapeEnum{
	"development": IntegrationInstanceShapeDevelopment,
	"production":  IntegrationInstanceShapeProduction,
}

// GetIntegrationInstanceShapeEnumValues Enumerates the set of values for IntegrationInstanceShapeEnum
func GetIntegrationInstanceShapeEnumValues() []IntegrationInstanceShapeEnum {
	values := make([]IntegrationInstanceShapeEnum, 0)
	for _, v := range mappingIntegrationInstanceShapeEnum {
		values = append(values, v)
	}
	return values
}

// GetIntegrationInstanceShapeEnumStringValues Enumerates the set of values in String for IntegrationInstanceShapeEnum
func GetIntegrationInstanceShapeEnumStringValues() []string {
	return []string{
		"DEVELOPMENT",
		"PRODUCTION",
	}
}

// GetMappingIntegrationInstanceShapeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIntegrationInstanceShapeEnum(val string) (IntegrationInstanceShapeEnum, bool) {
	enum, ok := mappingIntegrationInstanceShapeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// IntegrationInstanceDataRetentionPeriodEnum Enum with underlying type: string
type IntegrationInstanceDataRetentionPeriodEnum string

// Set of constants representing the allowable values for IntegrationInstanceDataRetentionPeriodEnum
const (
	IntegrationInstanceDataRetentionPeriod1 IntegrationInstanceDataRetentionPeriodEnum = "MONTHS_1"
	IntegrationInstanceDataRetentionPeriod3 IntegrationInstanceDataRetentionPeriodEnum = "MONTHS_3"
	IntegrationInstanceDataRetentionPeriod6 IntegrationInstanceDataRetentionPeriodEnum = "MONTHS_6"
)

var mappingIntegrationInstanceDataRetentionPeriodEnum = map[string]IntegrationInstanceDataRetentionPeriodEnum{
	"MONTHS_1": IntegrationInstanceDataRetentionPeriod1,
	"MONTHS_3": IntegrationInstanceDataRetentionPeriod3,
	"MONTHS_6": IntegrationInstanceDataRetentionPeriod6,
}

var mappingIntegrationInstanceDataRetentionPeriodEnumLowerCase = map[string]IntegrationInstanceDataRetentionPeriodEnum{
	"months_1": IntegrationInstanceDataRetentionPeriod1,
	"months_3": IntegrationInstanceDataRetentionPeriod3,
	"months_6": IntegrationInstanceDataRetentionPeriod6,
}

// GetIntegrationInstanceDataRetentionPeriodEnumValues Enumerates the set of values for IntegrationInstanceDataRetentionPeriodEnum
func GetIntegrationInstanceDataRetentionPeriodEnumValues() []IntegrationInstanceDataRetentionPeriodEnum {
	values := make([]IntegrationInstanceDataRetentionPeriodEnum, 0)
	for _, v := range mappingIntegrationInstanceDataRetentionPeriodEnum {
		values = append(values, v)
	}
	return values
}

// GetIntegrationInstanceDataRetentionPeriodEnumStringValues Enumerates the set of values in String for IntegrationInstanceDataRetentionPeriodEnum
func GetIntegrationInstanceDataRetentionPeriodEnumStringValues() []string {
	return []string{
		"MONTHS_1",
		"MONTHS_3",
		"MONTHS_6",
	}
}

// GetMappingIntegrationInstanceDataRetentionPeriodEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIntegrationInstanceDataRetentionPeriodEnum(val string) (IntegrationInstanceDataRetentionPeriodEnum, bool) {
	enum, ok := mappingIntegrationInstanceDataRetentionPeriodEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
