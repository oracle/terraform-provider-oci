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

// AuthenticationServerPolicy Policy for the details regarding each authentication server under dynamic authentication. We specify the value of selectors for which this authentication server must be selected for a request under keys. We specify the configuration details of authentication server under authenticationServerDetail.
type AuthenticationServerPolicy struct {
	Key DynamicSelectionKey `mandatory:"true" json:"key"`

	AuthenticationServerDetail AuthenticationPolicy `mandatory:"true" json:"authenticationServerDetail"`
}

func (m AuthenticationServerPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AuthenticationServerPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *AuthenticationServerPolicy) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key                        dynamicselectionkey  `json:"key"`
		AuthenticationServerDetail authenticationpolicy `json:"authenticationServerDetail"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.Key.UnmarshalPolymorphicJSON(model.Key.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Key = nn.(DynamicSelectionKey)
	} else {
		m.Key = nil
	}

	nn, e = model.AuthenticationServerDetail.UnmarshalPolymorphicJSON(model.AuthenticationServerDetail.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.AuthenticationServerDetail = nn.(AuthenticationPolicy)
	} else {
		m.AuthenticationServerDetail = nil
	}

	return
}
