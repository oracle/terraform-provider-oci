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
}

func (m VbInstanceSummary) String() string {
	return common.PointerString(m)
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

var mappingVbInstanceSummaryLifecycleState = map[string]VbInstanceSummaryLifecycleStateEnum{
	"CREATING": VbInstanceSummaryLifecycleStateCreating,
	"UPDATING": VbInstanceSummaryLifecycleStateUpdating,
	"ACTIVE":   VbInstanceSummaryLifecycleStateActive,
	"INACTIVE": VbInstanceSummaryLifecycleStateInactive,
	"DELETING": VbInstanceSummaryLifecycleStateDeleting,
	"DELETED":  VbInstanceSummaryLifecycleStateDeleted,
	"FAILED":   VbInstanceSummaryLifecycleStateFailed,
}

// GetVbInstanceSummaryLifecycleStateEnumValues Enumerates the set of values for VbInstanceSummaryLifecycleStateEnum
func GetVbInstanceSummaryLifecycleStateEnumValues() []VbInstanceSummaryLifecycleStateEnum {
	values := make([]VbInstanceSummaryLifecycleStateEnum, 0)
	for _, v := range mappingVbInstanceSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}

// VbInstanceSummaryConsumptionModelEnum Enum with underlying type: string
type VbInstanceSummaryConsumptionModelEnum string

// Set of constants representing the allowable values for VbInstanceSummaryConsumptionModelEnum
const (
	VbInstanceSummaryConsumptionModelUcm     VbInstanceSummaryConsumptionModelEnum = "UCM"
	VbInstanceSummaryConsumptionModelGov     VbInstanceSummaryConsumptionModelEnum = "GOV"
	VbInstanceSummaryConsumptionModelVb4saas VbInstanceSummaryConsumptionModelEnum = "VB4SAAS"
)

var mappingVbInstanceSummaryConsumptionModel = map[string]VbInstanceSummaryConsumptionModelEnum{
	"UCM":     VbInstanceSummaryConsumptionModelUcm,
	"GOV":     VbInstanceSummaryConsumptionModelGov,
	"VB4SAAS": VbInstanceSummaryConsumptionModelVb4saas,
}

// GetVbInstanceSummaryConsumptionModelEnumValues Enumerates the set of values for VbInstanceSummaryConsumptionModelEnum
func GetVbInstanceSummaryConsumptionModelEnumValues() []VbInstanceSummaryConsumptionModelEnum {
	values := make([]VbInstanceSummaryConsumptionModelEnum, 0)
	for _, v := range mappingVbInstanceSummaryConsumptionModel {
		values = append(values, v)
	}
	return values
}
