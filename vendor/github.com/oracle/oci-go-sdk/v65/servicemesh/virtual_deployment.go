// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Mesh API
//
// Use the Service Mesh API to manage mesh, virtual service, access policy and other mesh related items.
//

package servicemesh

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VirtualDeployment This resource represents a customer-managed virtual service deployment in the Service Mesh.
type VirtualDeployment struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the virtual service in which this virtual deployment is created.
	VirtualServiceId *string `mandatory:"true" json:"virtualServiceId"`

	// A user-friendly name. The name must be unique within the same virtual service and cannot be changed after creation.
	// Avoid entering confidential information.
	// Example: `My unique resource name`
	Name *string `mandatory:"true" json:"name"`

	// The time when this resource was created in an RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time when this resource was updated in an RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the Resource.
	LifecycleState VirtualDeploymentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Description of the resource. It can be changed after creation.
	// Avoid entering confidential information.
	// Example: `This is my new resource`
	Description *string `mandatory:"false" json:"description"`

	ServiceDiscovery ServiceDiscoveryConfiguration `mandatory:"false" json:"serviceDiscovery"`

	// The listeners for the virtual deployment
	Listeners []VirtualDeploymentListener `mandatory:"false" json:"listeners"`

	AccessLogging *AccessLoggingConfiguration `mandatory:"false" json:"accessLogging"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in a Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

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

func (m VirtualDeployment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VirtualDeployment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVirtualDeploymentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetVirtualDeploymentLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *VirtualDeployment) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description      *string                             `json:"description"`
		ServiceDiscovery servicediscoveryconfiguration       `json:"serviceDiscovery"`
		Listeners        []VirtualDeploymentListener         `json:"listeners"`
		AccessLogging    *AccessLoggingConfiguration         `json:"accessLogging"`
		LifecycleDetails *string                             `json:"lifecycleDetails"`
		FreeformTags     map[string]string                   `json:"freeformTags"`
		DefinedTags      map[string]map[string]interface{}   `json:"definedTags"`
		SystemTags       map[string]map[string]interface{}   `json:"systemTags"`
		Id               *string                             `json:"id"`
		CompartmentId    *string                             `json:"compartmentId"`
		VirtualServiceId *string                             `json:"virtualServiceId"`
		Name             *string                             `json:"name"`
		TimeCreated      *common.SDKTime                     `json:"timeCreated"`
		TimeUpdated      *common.SDKTime                     `json:"timeUpdated"`
		LifecycleState   VirtualDeploymentLifecycleStateEnum `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	nn, e = model.ServiceDiscovery.UnmarshalPolymorphicJSON(model.ServiceDiscovery.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ServiceDiscovery = nn.(ServiceDiscoveryConfiguration)
	} else {
		m.ServiceDiscovery = nil
	}

	m.Listeners = make([]VirtualDeploymentListener, len(model.Listeners))
	copy(m.Listeners, model.Listeners)
	m.AccessLogging = model.AccessLogging

	m.LifecycleDetails = model.LifecycleDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.VirtualServiceId = model.VirtualServiceId

	m.Name = model.Name

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

	return
}

// VirtualDeploymentLifecycleStateEnum Enum with underlying type: string
type VirtualDeploymentLifecycleStateEnum string

// Set of constants representing the allowable values for VirtualDeploymentLifecycleStateEnum
const (
	VirtualDeploymentLifecycleStateCreating VirtualDeploymentLifecycleStateEnum = "CREATING"
	VirtualDeploymentLifecycleStateUpdating VirtualDeploymentLifecycleStateEnum = "UPDATING"
	VirtualDeploymentLifecycleStateActive   VirtualDeploymentLifecycleStateEnum = "ACTIVE"
	VirtualDeploymentLifecycleStateDeleting VirtualDeploymentLifecycleStateEnum = "DELETING"
	VirtualDeploymentLifecycleStateDeleted  VirtualDeploymentLifecycleStateEnum = "DELETED"
	VirtualDeploymentLifecycleStateFailed   VirtualDeploymentLifecycleStateEnum = "FAILED"
)

var mappingVirtualDeploymentLifecycleStateEnum = map[string]VirtualDeploymentLifecycleStateEnum{
	"CREATING": VirtualDeploymentLifecycleStateCreating,
	"UPDATING": VirtualDeploymentLifecycleStateUpdating,
	"ACTIVE":   VirtualDeploymentLifecycleStateActive,
	"DELETING": VirtualDeploymentLifecycleStateDeleting,
	"DELETED":  VirtualDeploymentLifecycleStateDeleted,
	"FAILED":   VirtualDeploymentLifecycleStateFailed,
}

var mappingVirtualDeploymentLifecycleStateEnumLowerCase = map[string]VirtualDeploymentLifecycleStateEnum{
	"creating": VirtualDeploymentLifecycleStateCreating,
	"updating": VirtualDeploymentLifecycleStateUpdating,
	"active":   VirtualDeploymentLifecycleStateActive,
	"deleting": VirtualDeploymentLifecycleStateDeleting,
	"deleted":  VirtualDeploymentLifecycleStateDeleted,
	"failed":   VirtualDeploymentLifecycleStateFailed,
}

// GetVirtualDeploymentLifecycleStateEnumValues Enumerates the set of values for VirtualDeploymentLifecycleStateEnum
func GetVirtualDeploymentLifecycleStateEnumValues() []VirtualDeploymentLifecycleStateEnum {
	values := make([]VirtualDeploymentLifecycleStateEnum, 0)
	for _, v := range mappingVirtualDeploymentLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetVirtualDeploymentLifecycleStateEnumStringValues Enumerates the set of values in String for VirtualDeploymentLifecycleStateEnum
func GetVirtualDeploymentLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingVirtualDeploymentLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVirtualDeploymentLifecycleStateEnum(val string) (VirtualDeploymentLifecycleStateEnum, bool) {
	enum, ok := mappingVirtualDeploymentLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
