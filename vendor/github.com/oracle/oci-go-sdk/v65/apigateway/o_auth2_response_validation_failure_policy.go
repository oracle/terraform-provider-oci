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

// OAuth2ResponseValidationFailurePolicy Policy to specify OAuth2 flow configuration.
type OAuth2ResponseValidationFailurePolicy struct {
	ClientDetails ClientAppDetails `mandatory:"true" json:"clientDetails"`

	SourceUriDetails SourceUriDetails `mandatory:"true" json:"sourceUriDetails"`

	// List of scopes.
	Scopes []string `mandatory:"true" json:"scopes"`

	// The duration for which the OAuth2 success token should be cached before it is
	// fetched again.
	MaxExpiryDurationInHours *int `mandatory:"false" json:"maxExpiryDurationInHours"`

	// Defines whether or not to use cookies for session maintenance.
	UseCookiesForSession *bool `mandatory:"false" json:"useCookiesForSession"`

	// Defines whether or not to use cookies for OAuth2 intermediate steps.
	UseCookiesForIntermediateSteps *bool `mandatory:"false" json:"useCookiesForIntermediateSteps"`

	// Defines whether or not to support PKCE.
	UsePkce *bool `mandatory:"false" json:"usePkce"`

	// The path to be used as fallback after OAuth2.
	FallbackRedirectPath *string `mandatory:"false" json:"fallbackRedirectPath"`

	// The path to be used as logout.
	LogoutPath *string `mandatory:"false" json:"logoutPath"`

	// Response Type.
	ResponseType OAuth2ResponseValidationFailurePolicyResponseTypeEnum `mandatory:"true" json:"responseType"`
}

func (m OAuth2ResponseValidationFailurePolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OAuth2ResponseValidationFailurePolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOAuth2ResponseValidationFailurePolicyResponseTypeEnum(string(m.ResponseType)); !ok && m.ResponseType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResponseType: %s. Supported values are: %s.", m.ResponseType, strings.Join(GetOAuth2ResponseValidationFailurePolicyResponseTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OAuth2ResponseValidationFailurePolicy) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOAuth2ResponseValidationFailurePolicy OAuth2ResponseValidationFailurePolicy
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeOAuth2ResponseValidationFailurePolicy
	}{
		"OAUTH2",
		(MarshalTypeOAuth2ResponseValidationFailurePolicy)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *OAuth2ResponseValidationFailurePolicy) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		MaxExpiryDurationInHours       *int                                                  `json:"maxExpiryDurationInHours"`
		UseCookiesForSession           *bool                                                 `json:"useCookiesForSession"`
		UseCookiesForIntermediateSteps *bool                                                 `json:"useCookiesForIntermediateSteps"`
		UsePkce                        *bool                                                 `json:"usePkce"`
		FallbackRedirectPath           *string                                               `json:"fallbackRedirectPath"`
		LogoutPath                     *string                                               `json:"logoutPath"`
		ClientDetails                  clientappdetails                                      `json:"clientDetails"`
		SourceUriDetails               sourceuridetails                                      `json:"sourceUriDetails"`
		Scopes                         []string                                              `json:"scopes"`
		ResponseType                   OAuth2ResponseValidationFailurePolicyResponseTypeEnum `json:"responseType"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.MaxExpiryDurationInHours = model.MaxExpiryDurationInHours

	m.UseCookiesForSession = model.UseCookiesForSession

	m.UseCookiesForIntermediateSteps = model.UseCookiesForIntermediateSteps

	m.UsePkce = model.UsePkce

	m.FallbackRedirectPath = model.FallbackRedirectPath

	m.LogoutPath = model.LogoutPath

	nn, e = model.ClientDetails.UnmarshalPolymorphicJSON(model.ClientDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ClientDetails = nn.(ClientAppDetails)
	} else {
		m.ClientDetails = nil
	}

	nn, e = model.SourceUriDetails.UnmarshalPolymorphicJSON(model.SourceUriDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.SourceUriDetails = nn.(SourceUriDetails)
	} else {
		m.SourceUriDetails = nil
	}

	m.Scopes = make([]string, len(model.Scopes))
	copy(m.Scopes, model.Scopes)
	m.ResponseType = model.ResponseType

	return
}

// OAuth2ResponseValidationFailurePolicyResponseTypeEnum Enum with underlying type: string
type OAuth2ResponseValidationFailurePolicyResponseTypeEnum string

// Set of constants representing the allowable values for OAuth2ResponseValidationFailurePolicyResponseTypeEnum
const (
	OAuth2ResponseValidationFailurePolicyResponseTypeCode OAuth2ResponseValidationFailurePolicyResponseTypeEnum = "CODE"
)

var mappingOAuth2ResponseValidationFailurePolicyResponseTypeEnum = map[string]OAuth2ResponseValidationFailurePolicyResponseTypeEnum{
	"CODE": OAuth2ResponseValidationFailurePolicyResponseTypeCode,
}

var mappingOAuth2ResponseValidationFailurePolicyResponseTypeEnumLowerCase = map[string]OAuth2ResponseValidationFailurePolicyResponseTypeEnum{
	"code": OAuth2ResponseValidationFailurePolicyResponseTypeCode,
}

// GetOAuth2ResponseValidationFailurePolicyResponseTypeEnumValues Enumerates the set of values for OAuth2ResponseValidationFailurePolicyResponseTypeEnum
func GetOAuth2ResponseValidationFailurePolicyResponseTypeEnumValues() []OAuth2ResponseValidationFailurePolicyResponseTypeEnum {
	values := make([]OAuth2ResponseValidationFailurePolicyResponseTypeEnum, 0)
	for _, v := range mappingOAuth2ResponseValidationFailurePolicyResponseTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOAuth2ResponseValidationFailurePolicyResponseTypeEnumStringValues Enumerates the set of values in String for OAuth2ResponseValidationFailurePolicyResponseTypeEnum
func GetOAuth2ResponseValidationFailurePolicyResponseTypeEnumStringValues() []string {
	return []string{
		"CODE",
	}
}

// GetMappingOAuth2ResponseValidationFailurePolicyResponseTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOAuth2ResponseValidationFailurePolicyResponseTypeEnum(val string) (OAuth2ResponseValidationFailurePolicyResponseTypeEnum, bool) {
	enum, ok := mappingOAuth2ResponseValidationFailurePolicyResponseTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
