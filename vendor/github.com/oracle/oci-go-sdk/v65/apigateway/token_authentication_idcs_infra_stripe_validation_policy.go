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

// TokenAuthenticationIdcsInfraStripeValidationPolicy This contains the idcs infra stripe details and stripe manager function details to fetch the
// jwks for individual tenant at runtime.
type TokenAuthenticationIdcsInfraStripeValidationPolicy struct {
	ClientDetails ClientAppDetails `mandatory:"true" json:"clientDetails"`

	SourceUriDetails SourceUriDetails `mandatory:"true" json:"sourceUriDetails"`

	IdcsConfiguration *IdcsConfiguration `mandatory:"true" json:"idcsConfiguration"`

	AdditionalValidationPolicy *AdditionalValidationPolicy `mandatory:"false" json:"additionalValidationPolicy"`

	// Defines whether or not to uphold SSL verification.
	IsSslVerifyDisabled *bool `mandatory:"false" json:"isSslVerifyDisabled"`

	// The duration for which the introspect URL response should be cached before it is
	// fetched again.
	MaxCacheDurationInHours *int `mandatory:"false" json:"maxCacheDurationInHours"`
}

// GetAdditionalValidationPolicy returns AdditionalValidationPolicy
func (m TokenAuthenticationIdcsInfraStripeValidationPolicy) GetAdditionalValidationPolicy() *AdditionalValidationPolicy {
	return m.AdditionalValidationPolicy
}

func (m TokenAuthenticationIdcsInfraStripeValidationPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TokenAuthenticationIdcsInfraStripeValidationPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m TokenAuthenticationIdcsInfraStripeValidationPolicy) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTokenAuthenticationIdcsInfraStripeValidationPolicy TokenAuthenticationIdcsInfraStripeValidationPolicy
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeTokenAuthenticationIdcsInfraStripeValidationPolicy
	}{
		"IDCS_INFRA_STRIPE",
		(MarshalTypeTokenAuthenticationIdcsInfraStripeValidationPolicy)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *TokenAuthenticationIdcsInfraStripeValidationPolicy) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		AdditionalValidationPolicy *AdditionalValidationPolicy `json:"additionalValidationPolicy"`
		IsSslVerifyDisabled        *bool                       `json:"isSslVerifyDisabled"`
		MaxCacheDurationInHours    *int                        `json:"maxCacheDurationInHours"`
		ClientDetails              clientappdetails            `json:"clientDetails"`
		SourceUriDetails           sourceuridetails            `json:"sourceUriDetails"`
		IdcsConfiguration          *IdcsConfiguration          `json:"idcsConfiguration"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.AdditionalValidationPolicy = model.AdditionalValidationPolicy

	m.IsSslVerifyDisabled = model.IsSslVerifyDisabled

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

	m.IdcsConfiguration = model.IdcsConfiguration

	return
}
