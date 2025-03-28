// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// CreateVirtualServiceRouteTableDetails The information about the new VirtualServiceRouteTable.
type CreateVirtualServiceRouteTableDetails struct {

	// The OCID of the service mesh in which this access policy is created.
	VirtualServiceId *string `mandatory:"true" json:"virtualServiceId"`

	// A user-friendly name. The name must be unique within the same virtual service and cannot be changed after creation.
	// Avoid entering confidential information.
	// Example: `My unique resource name`
	Name *string `mandatory:"true" json:"name"`

	// The route rules for the virtual service.
	RouteRules []VirtualServiceTrafficRouteRuleDetails `mandatory:"true" json:"routeRules"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Description of the resource. It can be changed after creation.
	// Avoid entering confidential information.
	// Example: `This is my new resource`
	Description *string `mandatory:"false" json:"description"`

	// The priority of the route table. Lower value means higher priority. The routes are declared based on the priority.
	Priority *int `mandatory:"false" json:"priority"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateVirtualServiceRouteTableDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateVirtualServiceRouteTableDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateVirtualServiceRouteTableDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description      *string                                 `json:"description"`
		Priority         *int                                    `json:"priority"`
		FreeformTags     map[string]string                       `json:"freeformTags"`
		DefinedTags      map[string]map[string]interface{}       `json:"definedTags"`
		VirtualServiceId *string                                 `json:"virtualServiceId"`
		Name             *string                                 `json:"name"`
		RouteRules       []virtualservicetrafficrouteruledetails `json:"routeRules"`
		CompartmentId    *string                                 `json:"compartmentId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.Priority = model.Priority

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.VirtualServiceId = model.VirtualServiceId

	m.Name = model.Name

	m.RouteRules = make([]VirtualServiceTrafficRouteRuleDetails, len(model.RouteRules))
	for i, n := range model.RouteRules {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.RouteRules[i] = nn.(VirtualServiceTrafficRouteRuleDetails)
		} else {
			m.RouteRules[i] = nil
		}
	}
	m.CompartmentId = model.CompartmentId

	return
}
