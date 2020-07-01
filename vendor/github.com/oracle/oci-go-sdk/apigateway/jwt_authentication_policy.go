// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/common"
)

// JwtAuthenticationPolicy Validate a JWT token present in the header or query parameter. A valid
// policy must specify either tokenHeader or tokenQueryParam.
type JwtAuthenticationPolicy struct {

	// A list of parties that could have issued the token.
	Issuers []string `mandatory:"true" json:"issuers"`

	// The list of intended recipients for the token.
	Audiences []string `mandatory:"true" json:"audiences"`

	PublicKeys PublicKeySet `mandatory:"true" json:"publicKeys"`

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

	// A list of claims which should be validated to consider the token valid.
	VerifyClaims []JsonWebTokenClaim `mandatory:"false" json:"verifyClaims"`

	// The maximum expected time difference between the system clocks
	// of the token issuer and the API Gateway.
	MaxClockSkewInSeconds *float32 `mandatory:"false" json:"maxClockSkewInSeconds"`
}

//GetIsAnonymousAccessAllowed returns IsAnonymousAccessAllowed
func (m JwtAuthenticationPolicy) GetIsAnonymousAccessAllowed() *bool {
	return m.IsAnonymousAccessAllowed
}

func (m JwtAuthenticationPolicy) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m JwtAuthenticationPolicy) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeJwtAuthenticationPolicy JwtAuthenticationPolicy
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeJwtAuthenticationPolicy
	}{
		"JWT_AUTHENTICATION",
		(MarshalTypeJwtAuthenticationPolicy)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *JwtAuthenticationPolicy) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		IsAnonymousAccessAllowed *bool               `json:"isAnonymousAccessAllowed"`
		TokenHeader              *string             `json:"tokenHeader"`
		TokenQueryParam          *string             `json:"tokenQueryParam"`
		TokenAuthScheme          *string             `json:"tokenAuthScheme"`
		VerifyClaims             []JsonWebTokenClaim `json:"verifyClaims"`
		MaxClockSkewInSeconds    *float32            `json:"maxClockSkewInSeconds"`
		Issuers                  []string            `json:"issuers"`
		Audiences                []string            `json:"audiences"`
		PublicKeys               publickeyset        `json:"publicKeys"`
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

	m.VerifyClaims = make([]JsonWebTokenClaim, len(model.VerifyClaims))
	for i, n := range model.VerifyClaims {
		m.VerifyClaims[i] = n
	}

	m.MaxClockSkewInSeconds = model.MaxClockSkewInSeconds

	m.Issuers = make([]string, len(model.Issuers))
	for i, n := range model.Issuers {
		m.Issuers[i] = n
	}

	m.Audiences = make([]string, len(model.Audiences))
	for i, n := range model.Audiences {
		m.Audiences[i] = n
	}

	nn, e = model.PublicKeys.UnmarshalPolymorphicJSON(model.PublicKeys.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.PublicKeys = nn.(PublicKeySet)
	} else {
		m.PublicKeys = nil
	}

	return
}
