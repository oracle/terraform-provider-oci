// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// API Gateway API
//
// API for the API Gateway service. Use this API to manage gateways, deployments, and related items.
// For more information, see
// Overview of API Gateway (https://docs.oracle.com/iaas/Content/APIGateway/Concepts/apigatewayoverview.htm).
//

package apigateway

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DynamicRoutingBackend Send the request to the backend dynamically selected based on the incoming request's context.
type DynamicRoutingBackend struct {
	SelectionSource SelectionSourcePolicy `mandatory:"true" json:"selectionSource"`

	// List of backends to chose from for Dynamic Routing.
	RoutingBackends []DynamicRoutingTypeRoutingBackend `mandatory:"true" json:"routingBackends"`
}

func (m DynamicRoutingBackend) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DynamicRoutingBackend) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DynamicRoutingBackend) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDynamicRoutingBackend DynamicRoutingBackend
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeDynamicRoutingBackend
	}{
		"DYNAMIC_ROUTING_BACKEND",
		(MarshalTypeDynamicRoutingBackend)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *DynamicRoutingBackend) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		SelectionSource selectionsourcepolicy              `json:"selectionSource"`
		RoutingBackends []DynamicRoutingTypeRoutingBackend `json:"routingBackends"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.SelectionSource.UnmarshalPolymorphicJSON(model.SelectionSource.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.SelectionSource = nn.(SelectionSourcePolicy)
	} else {
		m.SelectionSource = nil
	}

	m.RoutingBackends = make([]DynamicRoutingTypeRoutingBackend, len(model.RoutingBackends))
	copy(m.RoutingBackends, model.RoutingBackends)
	return
}
