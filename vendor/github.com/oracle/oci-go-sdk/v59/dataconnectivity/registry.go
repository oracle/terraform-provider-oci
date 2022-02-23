// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the DCMS APIs to perform Metadata/Data operations.
//

package dataconnectivity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v59/common"
	"strings"
)

// Registry A registry is an organizational construct to keep multiple data Connectivity Management solutions and their resources (data assets, data flows, tasks, and so on) separate from each other, helping you to stay organized. For example, you could have separate registries for development, testing, and production.
type Registry struct {

	// Data Connectivity Management Registry display name, registries can be renamed
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Registry description
	Description *string `mandatory:"false" json:"description"`

	// Name of the user who updated the DCMS Registry.
	UpdatedBy *string `mandatory:"false" json:"updatedBy"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The time the Data Connectivity Management Registry was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the Data Connectivity Management Registry was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Lifecycle states for registries in Data Connectivity Management Service
	// CREATING - The resource is being created and may not be usable until the entire metadata is defined
	// UPDATING - The resource is being updated and may not be usable until all changes are commited
	// DELETING - The resource is being deleted and might require deep cleanup of children.
	// ACTIVE   - The resource is valid and available for access
	// INACTIVE - The resource might be incomplete in its definition or might have been made unavailable for
	//          administrative reasons
	// DELETED  - The resource has been deleted and isn't available
	// FAILED   - The resource is in a failed state due to validation or other errors
	LifecycleState RegistryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	StateMessage *string `mandatory:"false" json:"stateMessage"`
}

func (m Registry) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Registry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := mappingRegistryLifecycleStateEnum[string(m.LifecycleState)]; !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetRegistryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RegistryLifecycleStateEnum Enum with underlying type: string
type RegistryLifecycleStateEnum string

// Set of constants representing the allowable values for RegistryLifecycleStateEnum
const (
	RegistryLifecycleStateCreating RegistryLifecycleStateEnum = "CREATING"
	RegistryLifecycleStateActive   RegistryLifecycleStateEnum = "ACTIVE"
	RegistryLifecycleStateInactive RegistryLifecycleStateEnum = "INACTIVE"
	RegistryLifecycleStateUpdating RegistryLifecycleStateEnum = "UPDATING"
	RegistryLifecycleStateDeleting RegistryLifecycleStateEnum = "DELETING"
	RegistryLifecycleStateDeleted  RegistryLifecycleStateEnum = "DELETED"
	RegistryLifecycleStateFailed   RegistryLifecycleStateEnum = "FAILED"
)

var mappingRegistryLifecycleStateEnum = map[string]RegistryLifecycleStateEnum{
	"CREATING": RegistryLifecycleStateCreating,
	"ACTIVE":   RegistryLifecycleStateActive,
	"INACTIVE": RegistryLifecycleStateInactive,
	"UPDATING": RegistryLifecycleStateUpdating,
	"DELETING": RegistryLifecycleStateDeleting,
	"DELETED":  RegistryLifecycleStateDeleted,
	"FAILED":   RegistryLifecycleStateFailed,
}

// GetRegistryLifecycleStateEnumValues Enumerates the set of values for RegistryLifecycleStateEnum
func GetRegistryLifecycleStateEnumValues() []RegistryLifecycleStateEnum {
	values := make([]RegistryLifecycleStateEnum, 0)
	for _, v := range mappingRegistryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetRegistryLifecycleStateEnumStringValues Enumerates the set of values in String for RegistryLifecycleStateEnum
func GetRegistryLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}
