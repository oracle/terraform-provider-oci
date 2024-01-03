// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// API Gateway API
//
// API for the API Gateway service. Use this API to manage gateways, deployments, and related items.
// For more information, see
// Overview of API Gateway (https://docs.cloud.oracle.com/iaas/Content/APIGateway/Concepts/apigatewayoverview.htm).
//

package apigateway

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DynamicAuthenticationPolicy Policy on how to authenticate requests when multiple authentication options are configured for a deployment. For an incoming request, the value of selector specified under selectionSource will be matched against the keys specified for each authentication server. The authentication server whose key matches the value of selector will be used for authentication.
type DynamicAuthenticationPolicy struct {
	SelectionSource SelectionSourcePolicy `mandatory:"true" json:"selectionSource"`

	// List of authentication servers to choose from during dynamic authentication.
	AuthenticationServers []AuthenticationServerPolicy `mandatory:"true" json:"authenticationServers"`
}

func (m DynamicAuthenticationPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DynamicAuthenticationPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *DynamicAuthenticationPolicy) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		SelectionSource       selectionsourcepolicy        `json:"selectionSource"`
		AuthenticationServers []AuthenticationServerPolicy `json:"authenticationServers"`
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

	m.AuthenticationServers = make([]AuthenticationServerPolicy, len(model.AuthenticationServers))
	copy(m.AuthenticationServers, model.AuthenticationServers)
	return
}
