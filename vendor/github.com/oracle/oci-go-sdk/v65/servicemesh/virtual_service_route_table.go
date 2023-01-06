// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// VirtualServiceRouteTable This resource represents a customer-managed service route table in the Service Mesh.
type VirtualServiceRouteTable struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the virtual service in which this virtual service route table is created.
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
	LifecycleState VirtualServiceRouteTableLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Description of the resource. It can be changed after creation.
	// Avoid entering confidential information.
	// Example: `This is my new resource`
	Description *string `mandatory:"false" json:"description"`

	// The priority of the route table. Lower value means higher priority. The routes are declared based on the priority.
	Priority *int `mandatory:"false" json:"priority"`

	// The route rules for the virtual service.
	RouteRules []VirtualServiceTrafficRouteRule `mandatory:"false" json:"routeRules"`

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

func (m VirtualServiceRouteTable) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VirtualServiceRouteTable) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVirtualServiceRouteTableLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetVirtualServiceRouteTableLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *VirtualServiceRouteTable) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description      *string                                    `json:"description"`
		Priority         *int                                       `json:"priority"`
		RouteRules       []virtualservicetrafficrouterule           `json:"routeRules"`
		LifecycleDetails *string                                    `json:"lifecycleDetails"`
		FreeformTags     map[string]string                          `json:"freeformTags"`
		DefinedTags      map[string]map[string]interface{}          `json:"definedTags"`
		SystemTags       map[string]map[string]interface{}          `json:"systemTags"`
		Id               *string                                    `json:"id"`
		CompartmentId    *string                                    `json:"compartmentId"`
		VirtualServiceId *string                                    `json:"virtualServiceId"`
		Name             *string                                    `json:"name"`
		TimeCreated      *common.SDKTime                            `json:"timeCreated"`
		TimeUpdated      *common.SDKTime                            `json:"timeUpdated"`
		LifecycleState   VirtualServiceRouteTableLifecycleStateEnum `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.Priority = model.Priority

	m.RouteRules = make([]VirtualServiceTrafficRouteRule, len(model.RouteRules))
	for i, n := range model.RouteRules {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.RouteRules[i] = nn.(VirtualServiceTrafficRouteRule)
		} else {
			m.RouteRules[i] = nil
		}
	}

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

// VirtualServiceRouteTableLifecycleStateEnum Enum with underlying type: string
type VirtualServiceRouteTableLifecycleStateEnum string

// Set of constants representing the allowable values for VirtualServiceRouteTableLifecycleStateEnum
const (
	VirtualServiceRouteTableLifecycleStateCreating VirtualServiceRouteTableLifecycleStateEnum = "CREATING"
	VirtualServiceRouteTableLifecycleStateUpdating VirtualServiceRouteTableLifecycleStateEnum = "UPDATING"
	VirtualServiceRouteTableLifecycleStateActive   VirtualServiceRouteTableLifecycleStateEnum = "ACTIVE"
	VirtualServiceRouteTableLifecycleStateDeleting VirtualServiceRouteTableLifecycleStateEnum = "DELETING"
	VirtualServiceRouteTableLifecycleStateDeleted  VirtualServiceRouteTableLifecycleStateEnum = "DELETED"
	VirtualServiceRouteTableLifecycleStateFailed   VirtualServiceRouteTableLifecycleStateEnum = "FAILED"
)

var mappingVirtualServiceRouteTableLifecycleStateEnum = map[string]VirtualServiceRouteTableLifecycleStateEnum{
	"CREATING": VirtualServiceRouteTableLifecycleStateCreating,
	"UPDATING": VirtualServiceRouteTableLifecycleStateUpdating,
	"ACTIVE":   VirtualServiceRouteTableLifecycleStateActive,
	"DELETING": VirtualServiceRouteTableLifecycleStateDeleting,
	"DELETED":  VirtualServiceRouteTableLifecycleStateDeleted,
	"FAILED":   VirtualServiceRouteTableLifecycleStateFailed,
}

var mappingVirtualServiceRouteTableLifecycleStateEnumLowerCase = map[string]VirtualServiceRouteTableLifecycleStateEnum{
	"creating": VirtualServiceRouteTableLifecycleStateCreating,
	"updating": VirtualServiceRouteTableLifecycleStateUpdating,
	"active":   VirtualServiceRouteTableLifecycleStateActive,
	"deleting": VirtualServiceRouteTableLifecycleStateDeleting,
	"deleted":  VirtualServiceRouteTableLifecycleStateDeleted,
	"failed":   VirtualServiceRouteTableLifecycleStateFailed,
}

// GetVirtualServiceRouteTableLifecycleStateEnumValues Enumerates the set of values for VirtualServiceRouteTableLifecycleStateEnum
func GetVirtualServiceRouteTableLifecycleStateEnumValues() []VirtualServiceRouteTableLifecycleStateEnum {
	values := make([]VirtualServiceRouteTableLifecycleStateEnum, 0)
	for _, v := range mappingVirtualServiceRouteTableLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetVirtualServiceRouteTableLifecycleStateEnumStringValues Enumerates the set of values in String for VirtualServiceRouteTableLifecycleStateEnum
func GetVirtualServiceRouteTableLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingVirtualServiceRouteTableLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVirtualServiceRouteTableLifecycleStateEnum(val string) (VirtualServiceRouteTableLifecycleStateEnum, bool) {
	enum, ok := mappingVirtualServiceRouteTableLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
