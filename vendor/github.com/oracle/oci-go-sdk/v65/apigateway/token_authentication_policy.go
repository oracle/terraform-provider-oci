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

// TokenAuthenticationPolicy Validate a token present in the header or query parameter. A valid
// policy must specify either tokenHeader or tokenQueryParam.
type TokenAuthenticationPolicy struct {
	ValidationPolicy TokenAuthenticationValidationPolicy `mandatory:"true" json:"validationPolicy"`

	// Whether an unauthenticated user may access the API. Must be "true" to enable ANONYMOUS
	// route authorization.
	IsAnonymousAccessAllowed *bool `mandatory:"false" json:"isAnonymousAccessAllowed"`

	// The name of the header containing the authentication token.
	TokenHeader *string `mandatory:"false" json:"tokenHeader"`

	// The name of the query parameter containing the authentication token.
	TokenQueryParam *string `mandatory:"false" json:"tokenQueryParam"`

	// The authentication scheme that is to be used when authenticating
	// the token. This must to be provided if "tokenHeader" is specified.
	TokenAuthScheme *string `mandatory:"false" json:"tokenAuthScheme"`

	// The maximum expected time difference between the system clocks
	// of the token issuer and the API Gateway.
	MaxClockSkewInSeconds *float32 `mandatory:"false" json:"maxClockSkewInSeconds"`

	ValidationFailurePolicy ValidationFailurePolicy `mandatory:"false" json:"validationFailurePolicy"`
}

// GetIsAnonymousAccessAllowed returns IsAnonymousAccessAllowed
func (m TokenAuthenticationPolicy) GetIsAnonymousAccessAllowed() *bool {
	return m.IsAnonymousAccessAllowed
}

func (m TokenAuthenticationPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TokenAuthenticationPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m TokenAuthenticationPolicy) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTokenAuthenticationPolicy TokenAuthenticationPolicy
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeTokenAuthenticationPolicy
	}{
		"TOKEN_AUTHENTICATION",
		(MarshalTypeTokenAuthenticationPolicy)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *TokenAuthenticationPolicy) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		IsAnonymousAccessAllowed *bool                               `json:"isAnonymousAccessAllowed"`
		TokenHeader              *string                             `json:"tokenHeader"`
		TokenQueryParam          *string                             `json:"tokenQueryParam"`
		TokenAuthScheme          *string                             `json:"tokenAuthScheme"`
		MaxClockSkewInSeconds    *float32                            `json:"maxClockSkewInSeconds"`
		ValidationFailurePolicy  validationfailurepolicy             `json:"validationFailurePolicy"`
		ValidationPolicy         tokenauthenticationvalidationpolicy `json:"validationPolicy"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.IsAnonymousAccessAllowed = model.IsAnonymousAccessAllowed

	m.TokenHeader = model.TokenHeader

	m.TokenQueryParam = model.TokenQueryParam

	m.TokenAuthScheme = model.TokenAuthScheme

	m.MaxClockSkewInSeconds = model.MaxClockSkewInSeconds

	nn, e = model.ValidationFailurePolicy.UnmarshalPolymorphicJSON(model.ValidationFailurePolicy.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ValidationFailurePolicy = nn.(ValidationFailurePolicy)
	} else {
		m.ValidationFailurePolicy = nil
	}

	nn, e = model.ValidationPolicy.UnmarshalPolymorphicJSON(model.ValidationPolicy.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ValidationPolicy = nn.(TokenAuthenticationValidationPolicy)
	} else {
		m.ValidationPolicy = nil
	}

	return
}
