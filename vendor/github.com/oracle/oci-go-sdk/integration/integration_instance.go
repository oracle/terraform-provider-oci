// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Integration API
//
// Oracle Integration API.
//

package integration

import (
	"github.com/oracle/oci-go-sdk/common"
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
}

func (m IntegrationInstance) String() string {
	return common.PointerString(m)
}

// IntegrationInstanceIntegrationInstanceTypeEnum Enum with underlying type: string
type IntegrationInstanceIntegrationInstanceTypeEnum string

// Set of constants representing the allowable values for IntegrationInstanceIntegrationInstanceTypeEnum
const (
	IntegrationInstanceIntegrationInstanceTypeStandard   IntegrationInstanceIntegrationInstanceTypeEnum = "STANDARD"
	IntegrationInstanceIntegrationInstanceTypeEnterprise IntegrationInstanceIntegrationInstanceTypeEnum = "ENTERPRISE"
)

var mappingIntegrationInstanceIntegrationInstanceType = map[string]IntegrationInstanceIntegrationInstanceTypeEnum{
	"STANDARD":   IntegrationInstanceIntegrationInstanceTypeStandard,
	"ENTERPRISE": IntegrationInstanceIntegrationInstanceTypeEnterprise,
}

// GetIntegrationInstanceIntegrationInstanceTypeEnumValues Enumerates the set of values for IntegrationInstanceIntegrationInstanceTypeEnum
func GetIntegrationInstanceIntegrationInstanceTypeEnumValues() []IntegrationInstanceIntegrationInstanceTypeEnum {
	values := make([]IntegrationInstanceIntegrationInstanceTypeEnum, 0)
	for _, v := range mappingIntegrationInstanceIntegrationInstanceType {
		values = append(values, v)
	}
	return values
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

var mappingIntegrationInstanceLifecycleState = map[string]IntegrationInstanceLifecycleStateEnum{
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
	for _, v := range mappingIntegrationInstanceLifecycleState {
		values = append(values, v)
	}
	return values
}
