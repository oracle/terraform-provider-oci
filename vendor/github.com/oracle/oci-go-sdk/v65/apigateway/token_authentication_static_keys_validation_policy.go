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

// TokenAuthenticationStaticKeysValidationPolicy A set of static public keys that will be used to verify the JWT signature.
type TokenAuthenticationStaticKeysValidationPolicy struct {
	AdditionalValidationPolicy *AdditionalValidationPolicy `mandatory:"false" json:"additionalValidationPolicy"`

	// The set of static public keys.
	Keys []StaticPublicKey `mandatory:"false" json:"keys"`
}

// GetAdditionalValidationPolicy returns AdditionalValidationPolicy
func (m TokenAuthenticationStaticKeysValidationPolicy) GetAdditionalValidationPolicy() *AdditionalValidationPolicy {
	return m.AdditionalValidationPolicy
}

func (m TokenAuthenticationStaticKeysValidationPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TokenAuthenticationStaticKeysValidationPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m TokenAuthenticationStaticKeysValidationPolicy) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTokenAuthenticationStaticKeysValidationPolicy TokenAuthenticationStaticKeysValidationPolicy
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeTokenAuthenticationStaticKeysValidationPolicy
	}{
		"STATIC_KEYS",
		(MarshalTypeTokenAuthenticationStaticKeysValidationPolicy)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *TokenAuthenticationStaticKeysValidationPolicy) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		AdditionalValidationPolicy *AdditionalValidationPolicy `json:"additionalValidationPolicy"`
		Keys                       []staticpublickey           `json:"keys"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.AdditionalValidationPolicy = model.AdditionalValidationPolicy

	m.Keys = make([]StaticPublicKey, len(model.Keys))
	for i, n := range model.Keys {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Keys[i] = nn.(StaticPublicKey)
		} else {
			m.Keys[i] = nil
		}
	}
	return
}
