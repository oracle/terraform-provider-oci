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

// IngressGatewayRouteTable This resource represents a customer-managed ingress gateway route table in the Service Mesh.
type IngressGatewayRouteTable struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the ingress gateway.
	IngressGatewayId *string `mandatory:"true" json:"ingressGatewayId"`

	// A user-friendly name. The name must be unique within the same ingress gateway and cannot be changed after creation.
	// Avoid entering confidential information.
	// Example: `My unique resource name`
	Name *string `mandatory:"true" json:"name"`

	// The time when this resource was created in an RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time when this resource was updated in an RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the Resource.
	LifecycleState IngressGatewayRouteTableLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Description of the resource. It can be changed after creation.
	// Avoid entering confidential information.
	// Example: `This is my new resource`
	Description *string `mandatory:"false" json:"description"`

	// The priority of the route table. A lower value means a higher priority. The routes are declared based on the priority.
	Priority *int `mandatory:"false" json:"priority"`

	// The route rules for the ingress gateway.
	RouteRules []IngressGatewayTrafficRouteRule `mandatory:"false" json:"routeRules"`

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

func (m IngressGatewayRouteTable) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IngressGatewayRouteTable) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingIngressGatewayRouteTableLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetIngressGatewayRouteTableLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *IngressGatewayRouteTable) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description      *string                                    `json:"description"`
		Priority         *int                                       `json:"priority"`
		RouteRules       []ingressgatewaytrafficrouterule           `json:"routeRules"`
		LifecycleDetails *string                                    `json:"lifecycleDetails"`
		FreeformTags     map[string]string                          `json:"freeformTags"`
		DefinedTags      map[string]map[string]interface{}          `json:"definedTags"`
		SystemTags       map[string]map[string]interface{}          `json:"systemTags"`
		Id               *string                                    `json:"id"`
		CompartmentId    *string                                    `json:"compartmentId"`
		IngressGatewayId *string                                    `json:"ingressGatewayId"`
		Name             *string                                    `json:"name"`
		TimeCreated      *common.SDKTime                            `json:"timeCreated"`
		TimeUpdated      *common.SDKTime                            `json:"timeUpdated"`
		LifecycleState   IngressGatewayRouteTableLifecycleStateEnum `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.Priority = model.Priority

	m.RouteRules = make([]IngressGatewayTrafficRouteRule, len(model.RouteRules))
	for i, n := range model.RouteRules {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.RouteRules[i] = nn.(IngressGatewayTrafficRouteRule)
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

	m.IngressGatewayId = model.IngressGatewayId

	m.Name = model.Name

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

	return
}

// IngressGatewayRouteTableLifecycleStateEnum Enum with underlying type: string
type IngressGatewayRouteTableLifecycleStateEnum string

// Set of constants representing the allowable values for IngressGatewayRouteTableLifecycleStateEnum
const (
	IngressGatewayRouteTableLifecycleStateCreating IngressGatewayRouteTableLifecycleStateEnum = "CREATING"
	IngressGatewayRouteTableLifecycleStateUpdating IngressGatewayRouteTableLifecycleStateEnum = "UPDATING"
	IngressGatewayRouteTableLifecycleStateActive   IngressGatewayRouteTableLifecycleStateEnum = "ACTIVE"
	IngressGatewayRouteTableLifecycleStateDeleting IngressGatewayRouteTableLifecycleStateEnum = "DELETING"
	IngressGatewayRouteTableLifecycleStateDeleted  IngressGatewayRouteTableLifecycleStateEnum = "DELETED"
	IngressGatewayRouteTableLifecycleStateFailed   IngressGatewayRouteTableLifecycleStateEnum = "FAILED"
)

var mappingIngressGatewayRouteTableLifecycleStateEnum = map[string]IngressGatewayRouteTableLifecycleStateEnum{
	"CREATING": IngressGatewayRouteTableLifecycleStateCreating,
	"UPDATING": IngressGatewayRouteTableLifecycleStateUpdating,
	"ACTIVE":   IngressGatewayRouteTableLifecycleStateActive,
	"DELETING": IngressGatewayRouteTableLifecycleStateDeleting,
	"DELETED":  IngressGatewayRouteTableLifecycleStateDeleted,
	"FAILED":   IngressGatewayRouteTableLifecycleStateFailed,
}

var mappingIngressGatewayRouteTableLifecycleStateEnumLowerCase = map[string]IngressGatewayRouteTableLifecycleStateEnum{
	"creating": IngressGatewayRouteTableLifecycleStateCreating,
	"updating": IngressGatewayRouteTableLifecycleStateUpdating,
	"active":   IngressGatewayRouteTableLifecycleStateActive,
	"deleting": IngressGatewayRouteTableLifecycleStateDeleting,
	"deleted":  IngressGatewayRouteTableLifecycleStateDeleted,
	"failed":   IngressGatewayRouteTableLifecycleStateFailed,
}

// GetIngressGatewayRouteTableLifecycleStateEnumValues Enumerates the set of values for IngressGatewayRouteTableLifecycleStateEnum
func GetIngressGatewayRouteTableLifecycleStateEnumValues() []IngressGatewayRouteTableLifecycleStateEnum {
	values := make([]IngressGatewayRouteTableLifecycleStateEnum, 0)
	for _, v := range mappingIngressGatewayRouteTableLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetIngressGatewayRouteTableLifecycleStateEnumStringValues Enumerates the set of values in String for IngressGatewayRouteTableLifecycleStateEnum
func GetIngressGatewayRouteTableLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingIngressGatewayRouteTableLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIngressGatewayRouteTableLifecycleStateEnum(val string) (IngressGatewayRouteTableLifecycleStateEnum, bool) {
	enum, ok := mappingIngressGatewayRouteTableLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
