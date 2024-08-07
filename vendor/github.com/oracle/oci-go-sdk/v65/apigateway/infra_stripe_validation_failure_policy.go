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

// InfraStripeValidationFailurePolicy Policy to specify idcs infra stripe configuration.
type InfraStripeValidationFailurePolicy struct {
	ClientDetails ClientAppDetails `mandatory:"true" json:"clientDetails"`

	SourceUriDetails SourceUriDetails `mandatory:"true" json:"sourceUriDetails"`

	// List of scopes.
	Scopes []string `mandatory:"true" json:"scopes"`

	// The path to be used as fallback after OAuth2.
	FallbackRedirectPath *string `mandatory:"false" json:"fallbackRedirectPath"`

	// The path to be used as logout.
	LogoutPath *string `mandatory:"false" json:"logoutPath"`

	// The duration for which token should be cached before it is fetched again.
	MaxCacheDurationInHours *int `mandatory:"false" json:"maxCacheDurationInHours"`

	// Response Type.
	ResponseType InfraStripeValidationFailurePolicyResponseTypeEnum `mandatory:"true" json:"responseType"`
}

func (m InfraStripeValidationFailurePolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InfraStripeValidationFailurePolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInfraStripeValidationFailurePolicyResponseTypeEnum(string(m.ResponseType)); !ok && m.ResponseType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResponseType: %s. Supported values are: %s.", m.ResponseType, strings.Join(GetInfraStripeValidationFailurePolicyResponseTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m InfraStripeValidationFailurePolicy) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInfraStripeValidationFailurePolicy InfraStripeValidationFailurePolicy
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeInfraStripeValidationFailurePolicy
	}{
		"INFRA_STRIPE",
		(MarshalTypeInfraStripeValidationFailurePolicy)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *InfraStripeValidationFailurePolicy) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		FallbackRedirectPath    *string                                            `json:"fallbackRedirectPath"`
		LogoutPath              *string                                            `json:"logoutPath"`
		MaxCacheDurationInHours *int                                               `json:"maxCacheDurationInHours"`
		ClientDetails           clientappdetails                                   `json:"clientDetails"`
		SourceUriDetails        sourceuridetails                                   `json:"sourceUriDetails"`
		Scopes                  []string                                           `json:"scopes"`
		ResponseType            InfraStripeValidationFailurePolicyResponseTypeEnum `json:"responseType"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.FallbackRedirectPath = model.FallbackRedirectPath

	m.LogoutPath = model.LogoutPath

	m.MaxCacheDurationInHours = model.MaxCacheDurationInHours

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

// InfraStripeValidationFailurePolicyResponseTypeEnum Enum with underlying type: string
type InfraStripeValidationFailurePolicyResponseTypeEnum string

// Set of constants representing the allowable values for InfraStripeValidationFailurePolicyResponseTypeEnum
const (
	InfraStripeValidationFailurePolicyResponseTypeCode InfraStripeValidationFailurePolicyResponseTypeEnum = "CODE"
)

var mappingInfraStripeValidationFailurePolicyResponseTypeEnum = map[string]InfraStripeValidationFailurePolicyResponseTypeEnum{
	"CODE": InfraStripeValidationFailurePolicyResponseTypeCode,
}

var mappingInfraStripeValidationFailurePolicyResponseTypeEnumLowerCase = map[string]InfraStripeValidationFailurePolicyResponseTypeEnum{
	"code": InfraStripeValidationFailurePolicyResponseTypeCode,
}

// GetInfraStripeValidationFailurePolicyResponseTypeEnumValues Enumerates the set of values for InfraStripeValidationFailurePolicyResponseTypeEnum
func GetInfraStripeValidationFailurePolicyResponseTypeEnumValues() []InfraStripeValidationFailurePolicyResponseTypeEnum {
	values := make([]InfraStripeValidationFailurePolicyResponseTypeEnum, 0)
	for _, v := range mappingInfraStripeValidationFailurePolicyResponseTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetInfraStripeValidationFailurePolicyResponseTypeEnumStringValues Enumerates the set of values in String for InfraStripeValidationFailurePolicyResponseTypeEnum
func GetInfraStripeValidationFailurePolicyResponseTypeEnumStringValues() []string {
	return []string{
		"CODE",
	}
}

// GetMappingInfraStripeValidationFailurePolicyResponseTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInfraStripeValidationFailurePolicyResponseTypeEnum(val string) (InfraStripeValidationFailurePolicyResponseTypeEnum, bool) {
	enum, ok := mappingInfraStripeValidationFailurePolicyResponseTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
