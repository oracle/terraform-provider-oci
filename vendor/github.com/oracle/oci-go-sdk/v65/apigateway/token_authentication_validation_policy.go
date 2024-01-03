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

// TokenAuthenticationValidationPolicy Authentication Policies for the Token Authentication types.
type TokenAuthenticationValidationPolicy interface {
	GetAdditionalValidationPolicy() *AdditionalValidationPolicy
}

type tokenauthenticationvalidationpolicy struct {
	JsonData                   []byte
	AdditionalValidationPolicy *AdditionalValidationPolicy `mandatory:"false" json:"additionalValidationPolicy"`
	Type                       string                      `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *tokenauthenticationvalidationpolicy) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertokenauthenticationvalidationpolicy tokenauthenticationvalidationpolicy
	s := struct {
		Model Unmarshalertokenauthenticationvalidationpolicy
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.AdditionalValidationPolicy = s.Model.AdditionalValidationPolicy
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *tokenauthenticationvalidationpolicy) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "REMOTE_JWKS":
		mm := TokenAuthenticationRemoteJwksValidationPolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "REMOTE_DISCOVERY":
		mm := TokenAuthenticationRemoteDiscoveryValidationPolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "STATIC_KEYS":
		mm := TokenAuthenticationStaticKeysValidationPolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for TokenAuthenticationValidationPolicy: %s.", m.Type)
		return *m, nil
	}
}

// GetAdditionalValidationPolicy returns AdditionalValidationPolicy
func (m tokenauthenticationvalidationpolicy) GetAdditionalValidationPolicy() *AdditionalValidationPolicy {
	return m.AdditionalValidationPolicy
}

func (m tokenauthenticationvalidationpolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m tokenauthenticationvalidationpolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TokenAuthenticationValidationPolicyTypeEnum Enum with underlying type: string
type TokenAuthenticationValidationPolicyTypeEnum string

// Set of constants representing the allowable values for TokenAuthenticationValidationPolicyTypeEnum
const (
	TokenAuthenticationValidationPolicyTypeStaticKeys      TokenAuthenticationValidationPolicyTypeEnum = "STATIC_KEYS"
	TokenAuthenticationValidationPolicyTypeRemoteJwks      TokenAuthenticationValidationPolicyTypeEnum = "REMOTE_JWKS"
	TokenAuthenticationValidationPolicyTypeRemoteDiscovery TokenAuthenticationValidationPolicyTypeEnum = "REMOTE_DISCOVERY"
)

var mappingTokenAuthenticationValidationPolicyTypeEnum = map[string]TokenAuthenticationValidationPolicyTypeEnum{
	"STATIC_KEYS":      TokenAuthenticationValidationPolicyTypeStaticKeys,
	"REMOTE_JWKS":      TokenAuthenticationValidationPolicyTypeRemoteJwks,
	"REMOTE_DISCOVERY": TokenAuthenticationValidationPolicyTypeRemoteDiscovery,
}

var mappingTokenAuthenticationValidationPolicyTypeEnumLowerCase = map[string]TokenAuthenticationValidationPolicyTypeEnum{
	"static_keys":      TokenAuthenticationValidationPolicyTypeStaticKeys,
	"remote_jwks":      TokenAuthenticationValidationPolicyTypeRemoteJwks,
	"remote_discovery": TokenAuthenticationValidationPolicyTypeRemoteDiscovery,
}

// GetTokenAuthenticationValidationPolicyTypeEnumValues Enumerates the set of values for TokenAuthenticationValidationPolicyTypeEnum
func GetTokenAuthenticationValidationPolicyTypeEnumValues() []TokenAuthenticationValidationPolicyTypeEnum {
	values := make([]TokenAuthenticationValidationPolicyTypeEnum, 0)
	for _, v := range mappingTokenAuthenticationValidationPolicyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTokenAuthenticationValidationPolicyTypeEnumStringValues Enumerates the set of values in String for TokenAuthenticationValidationPolicyTypeEnum
func GetTokenAuthenticationValidationPolicyTypeEnumStringValues() []string {
	return []string{
		"STATIC_KEYS",
		"REMOTE_JWKS",
		"REMOTE_DISCOVERY",
	}
}

// GetMappingTokenAuthenticationValidationPolicyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTokenAuthenticationValidationPolicyTypeEnum(val string) (TokenAuthenticationValidationPolicyTypeEnum, bool) {
	enum, ok := mappingTokenAuthenticationValidationPolicyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
