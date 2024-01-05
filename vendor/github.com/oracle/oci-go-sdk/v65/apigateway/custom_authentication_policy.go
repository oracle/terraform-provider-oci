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

// CustomAuthenticationPolicy Use a function to validate a custom header or query parameter sent with the request authentication.
// A valid policy must specify either tokenHeader or tokenQueryParam.
type CustomAuthenticationPolicy struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Oracle Functions function resource.
	FunctionId *string `mandatory:"true" json:"functionId"`

	// Whether an unauthenticated user may access the API. Must be "true" to enable ANONYMOUS
	// route authorization.
	IsAnonymousAccessAllowed *bool `mandatory:"false" json:"isAnonymousAccessAllowed"`

	// The name of the header containing the authentication token.
	TokenHeader *string `mandatory:"false" json:"tokenHeader"`

	// The name of the query parameter containing the authentication token.
	TokenQueryParam *string `mandatory:"false" json:"tokenQueryParam"`

	// A map where key is a user defined string and value is a context expressions whose values will be sent to the custom auth function. Values should contain an expression.
	// Example: `{"foo": "request.header[abc]"}`
	Parameters map[string]string `mandatory:"false" json:"parameters"`

	// A list of keys from "parameters" attribute value whose values will be added to the cache key.
	CacheKey []string `mandatory:"false" json:"cacheKey"`

	ValidationFailurePolicy ValidationFailurePolicy `mandatory:"false" json:"validationFailurePolicy"`
}

// GetIsAnonymousAccessAllowed returns IsAnonymousAccessAllowed
func (m CustomAuthenticationPolicy) GetIsAnonymousAccessAllowed() *bool {
	return m.IsAnonymousAccessAllowed
}

func (m CustomAuthenticationPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CustomAuthenticationPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CustomAuthenticationPolicy) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCustomAuthenticationPolicy CustomAuthenticationPolicy
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCustomAuthenticationPolicy
	}{
		"CUSTOM_AUTHENTICATION",
		(MarshalTypeCustomAuthenticationPolicy)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CustomAuthenticationPolicy) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		IsAnonymousAccessAllowed *bool                   `json:"isAnonymousAccessAllowed"`
		TokenHeader              *string                 `json:"tokenHeader"`
		TokenQueryParam          *string                 `json:"tokenQueryParam"`
		Parameters               map[string]string       `json:"parameters"`
		CacheKey                 []string                `json:"cacheKey"`
		ValidationFailurePolicy  validationfailurepolicy `json:"validationFailurePolicy"`
		FunctionId               *string                 `json:"functionId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.IsAnonymousAccessAllowed = model.IsAnonymousAccessAllowed

	m.TokenHeader = model.TokenHeader

	m.TokenQueryParam = model.TokenQueryParam

	m.Parameters = model.Parameters

	m.CacheKey = make([]string, len(model.CacheKey))
	copy(m.CacheKey, model.CacheKey)
	nn, e = model.ValidationFailurePolicy.UnmarshalPolymorphicJSON(model.ValidationFailurePolicy.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ValidationFailurePolicy = nn.(ValidationFailurePolicy)
	} else {
		m.ValidationFailurePolicy = nil
	}

	m.FunctionId = model.FunctionId

	return
}
