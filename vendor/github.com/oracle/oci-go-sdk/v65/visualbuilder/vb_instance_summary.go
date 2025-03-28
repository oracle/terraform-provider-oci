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

// VbInstanceSummary Summary of the Vb Instance.
type VbInstanceSummary struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// Vb Instance Identifier, can be renamed.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current state of the Vb Instance.
	LifecycleState VbInstanceSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The Vb Instance URL.
	InstanceUrl *string `mandatory:"true" json:"instanceUrl"`

	// The number of Nodes
	NodeCount *int `mandatory:"true" json:"nodeCount"`

	// The time the the Vb Instance was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the VbInstance was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// An message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	StateMessage *string `mandatory:"false" json:"stateMessage"`

	// Visual Builder is enabled or not.
	IsVisualBuilderEnabled *bool `mandatory:"false" json:"isVisualBuilderEnabled"`

	CustomEndpoint *CustomEndpointDetails `mandatory:"false" json:"customEndpoint"`

	// A list of alternate custom endpoints used for the vb instance URL.
	AlternateCustomEndpoints []CustomEndpointDetails `mandatory:"false" json:"alternateCustomEndpoints"`

	// The entitlement used for billing purposes.
	ConsumptionModel VbInstanceSummaryConsumptionModelEnum `mandatory:"false" json:"consumptionModel,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	NetworkEndpointDetails NetworkEndpointDetails `mandatory:"false" json:"networkEndpointDetails"`
}

func (m VbInstanceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VbInstanceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVbInstanceSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetVbInstanceSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingVbInstanceSummaryConsumptionModelEnum(string(m.ConsumptionModel)); !ok && m.ConsumptionModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConsumptionModel: %s. Supported values are: %s.", m.ConsumptionModel, strings.Join(GetVbInstanceSummaryConsumptionModelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *VbInstanceSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TimeCreated              *common.SDKTime                       `json:"timeCreated"`
		TimeUpdated              *common.SDKTime                       `json:"timeUpdated"`
		StateMessage             *string                               `json:"stateMessage"`
		IsVisualBuilderEnabled   *bool                                 `json:"isVisualBuilderEnabled"`
		CustomEndpoint           *CustomEndpointDetails                `json:"customEndpoint"`
		AlternateCustomEndpoints []CustomEndpointDetails               `json:"alternateCustomEndpoints"`
		ConsumptionModel         VbInstanceSummaryConsumptionModelEnum `json:"consumptionModel"`
		FreeformTags             map[string]string                     `json:"freeformTags"`
		DefinedTags              map[string]map[string]interface{}     `json:"definedTags"`
		SystemTags               map[string]map[string]interface{}     `json:"systemTags"`
		NetworkEndpointDetails   networkendpointdetails                `json:"networkEndpointDetails"`
		Id                       *string                               `json:"id"`
		DisplayName              *string                               `json:"displayName"`
		CompartmentId            *string                               `json:"compartmentId"`
		LifecycleState           VbInstanceSummaryLifecycleStateEnum   `json:"lifecycleState"`
		InstanceUrl              *string                               `json:"instanceUrl"`
		NodeCount                *int                                  `json:"nodeCount"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.StateMessage = model.StateMessage

	m.IsVisualBuilderEnabled = model.IsVisualBuilderEnabled

	m.CustomEndpoint = model.CustomEndpoint

	m.AlternateCustomEndpoints = make([]CustomEndpointDetails, len(model.AlternateCustomEndpoints))
	copy(m.AlternateCustomEndpoints, model.AlternateCustomEndpoints)
	m.ConsumptionModel = model.ConsumptionModel

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

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

	m.LifecycleState = model.LifecycleState

	m.InstanceUrl = model.InstanceUrl

	m.NodeCount = model.NodeCount

	return
}

// VbInstanceSummaryLifecycleStateEnum Enum with underlying type: string
type VbInstanceSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for VbInstanceSummaryLifecycleStateEnum
const (
	VbInstanceSummaryLifecycleStateCreating VbInstanceSummaryLifecycleStateEnum = "CREATING"
	VbInstanceSummaryLifecycleStateUpdating VbInstanceSummaryLifecycleStateEnum = "UPDATING"
	VbInstanceSummaryLifecycleStateActive   VbInstanceSummaryLifecycleStateEnum = "ACTIVE"
	VbInstanceSummaryLifecycleStateInactive VbInstanceSummaryLifecycleStateEnum = "INACTIVE"
	VbInstanceSummaryLifecycleStateDeleting VbInstanceSummaryLifecycleStateEnum = "DELETING"
	VbInstanceSummaryLifecycleStateDeleted  VbInstanceSummaryLifecycleStateEnum = "DELETED"
	VbInstanceSummaryLifecycleStateFailed   VbInstanceSummaryLifecycleStateEnum = "FAILED"
)

var mappingVbInstanceSummaryLifecycleStateEnum = map[string]VbInstanceSummaryLifecycleStateEnum{
	"CREATING": VbInstanceSummaryLifecycleStateCreating,
	"UPDATING": VbInstanceSummaryLifecycleStateUpdating,
	"ACTIVE":   VbInstanceSummaryLifecycleStateActive,
	"INACTIVE": VbInstanceSummaryLifecycleStateInactive,
	"DELETING": VbInstanceSummaryLifecycleStateDeleting,
	"DELETED":  VbInstanceSummaryLifecycleStateDeleted,
	"FAILED":   VbInstanceSummaryLifecycleStateFailed,
}

var mappingVbInstanceSummaryLifecycleStateEnumLowerCase = map[string]VbInstanceSummaryLifecycleStateEnum{
	"creating": VbInstanceSummaryLifecycleStateCreating,
	"updating": VbInstanceSummaryLifecycleStateUpdating,
	"active":   VbInstanceSummaryLifecycleStateActive,
	"inactive": VbInstanceSummaryLifecycleStateInactive,
	"deleting": VbInstanceSummaryLifecycleStateDeleting,
	"deleted":  VbInstanceSummaryLifecycleStateDeleted,
	"failed":   VbInstanceSummaryLifecycleStateFailed,
}

// GetVbInstanceSummaryLifecycleStateEnumValues Enumerates the set of values for VbInstanceSummaryLifecycleStateEnum
func GetVbInstanceSummaryLifecycleStateEnumValues() []VbInstanceSummaryLifecycleStateEnum {
	values := make([]VbInstanceSummaryLifecycleStateEnum, 0)
	for _, v := range mappingVbInstanceSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetVbInstanceSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for VbInstanceSummaryLifecycleStateEnum
func GetVbInstanceSummaryLifecycleStateEnumStringValues() []string {
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

// GetMappingVbInstanceSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVbInstanceSummaryLifecycleStateEnum(val string) (VbInstanceSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingVbInstanceSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// VbInstanceSummaryConsumptionModelEnum Enum with underlying type: string
type VbInstanceSummaryConsumptionModelEnum string

// Set of constants representing the allowable values for VbInstanceSummaryConsumptionModelEnum
const (
	VbInstanceSummaryConsumptionModelUcm     VbInstanceSummaryConsumptionModelEnum = "UCM"
	VbInstanceSummaryConsumptionModelGov     VbInstanceSummaryConsumptionModelEnum = "GOV"
	VbInstanceSummaryConsumptionModelVb4saas VbInstanceSummaryConsumptionModelEnum = "VB4SAAS"
)

var mappingVbInstanceSummaryConsumptionModelEnum = map[string]VbInstanceSummaryConsumptionModelEnum{
	"UCM":     VbInstanceSummaryConsumptionModelUcm,
	"GOV":     VbInstanceSummaryConsumptionModelGov,
	"VB4SAAS": VbInstanceSummaryConsumptionModelVb4saas,
}

var mappingVbInstanceSummaryConsumptionModelEnumLowerCase = map[string]VbInstanceSummaryConsumptionModelEnum{
	"ucm":     VbInstanceSummaryConsumptionModelUcm,
	"gov":     VbInstanceSummaryConsumptionModelGov,
	"vb4saas": VbInstanceSummaryConsumptionModelVb4saas,
}

// GetVbInstanceSummaryConsumptionModelEnumValues Enumerates the set of values for VbInstanceSummaryConsumptionModelEnum
func GetVbInstanceSummaryConsumptionModelEnumValues() []VbInstanceSummaryConsumptionModelEnum {
	values := make([]VbInstanceSummaryConsumptionModelEnum, 0)
	for _, v := range mappingVbInstanceSummaryConsumptionModelEnum {
		values = append(values, v)
	}
	return values
}

// GetVbInstanceSummaryConsumptionModelEnumStringValues Enumerates the set of values in String for VbInstanceSummaryConsumptionModelEnum
func GetVbInstanceSummaryConsumptionModelEnumStringValues() []string {
	return []string{
		"UCM",
		"GOV",
		"VB4SAAS",
	}
}

// GetMappingVbInstanceSummaryConsumptionModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVbInstanceSummaryConsumptionModelEnum(val string) (VbInstanceSummaryConsumptionModelEnum, bool) {
	enum, ok := mappingVbInstanceSummaryConsumptionModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
