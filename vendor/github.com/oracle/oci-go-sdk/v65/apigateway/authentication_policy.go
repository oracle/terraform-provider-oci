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

// AuthenticationPolicy Information on how to authenticate incoming requests.
type AuthenticationPolicy interface {

	// Whether an unauthenticated user may access the API. Must be "true" to enable ANONYMOUS
	// route authorization.
	GetIsAnonymousAccessAllowed() *bool
}

type authenticationpolicy struct {
	JsonData                 []byte
	IsAnonymousAccessAllowed *bool  `mandatory:"false" json:"isAnonymousAccessAllowed"`
	Type                     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *authenticationpolicy) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerauthenticationpolicy authenticationpolicy
	s := struct {
		Model Unmarshalerauthenticationpolicy
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.IsAnonymousAccessAllowed = s.Model.IsAnonymousAccessAllowed
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *authenticationpolicy) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "TOKEN_AUTHENTICATION":
		mm := TokenAuthenticationPolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "JWT_AUTHENTICATION":
		mm := JwtAuthenticationPolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CUSTOM_AUTHENTICATION":
		mm := CustomAuthenticationPolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for AuthenticationPolicy: %s.", m.Type)
		return *m, nil
	}
}

// GetIsAnonymousAccessAllowed returns IsAnonymousAccessAllowed
func (m authenticationpolicy) GetIsAnonymousAccessAllowed() *bool {
	return m.IsAnonymousAccessAllowed
}

func (m authenticationpolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m authenticationpolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AuthenticationPolicyTypeEnum Enum with underlying type: string
type AuthenticationPolicyTypeEnum string

// Set of constants representing the allowable values for AuthenticationPolicyTypeEnum
const (
	AuthenticationPolicyTypeCustomAuthentication AuthenticationPolicyTypeEnum = "CUSTOM_AUTHENTICATION"
	AuthenticationPolicyTypeJwtAuthentication    AuthenticationPolicyTypeEnum = "JWT_AUTHENTICATION"
	AuthenticationPolicyTypeTokenAuthentication  AuthenticationPolicyTypeEnum = "TOKEN_AUTHENTICATION"
)

var mappingAuthenticationPolicyTypeEnum = map[string]AuthenticationPolicyTypeEnum{
	"CUSTOM_AUTHENTICATION": AuthenticationPolicyTypeCustomAuthentication,
	"JWT_AUTHENTICATION":    AuthenticationPolicyTypeJwtAuthentication,
	"TOKEN_AUTHENTICATION":  AuthenticationPolicyTypeTokenAuthentication,
}

var mappingAuthenticationPolicyTypeEnumLowerCase = map[string]AuthenticationPolicyTypeEnum{
	"custom_authentication": AuthenticationPolicyTypeCustomAuthentication,
	"jwt_authentication":    AuthenticationPolicyTypeJwtAuthentication,
	"token_authentication":  AuthenticationPolicyTypeTokenAuthentication,
}

// GetAuthenticationPolicyTypeEnumValues Enumerates the set of values for AuthenticationPolicyTypeEnum
func GetAuthenticationPolicyTypeEnumValues() []AuthenticationPolicyTypeEnum {
	values := make([]AuthenticationPolicyTypeEnum, 0)
	for _, v := range mappingAuthenticationPolicyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAuthenticationPolicyTypeEnumStringValues Enumerates the set of values in String for AuthenticationPolicyTypeEnum
func GetAuthenticationPolicyTypeEnumStringValues() []string {
	return []string{
		"CUSTOM_AUTHENTICATION",
		"JWT_AUTHENTICATION",
		"TOKEN_AUTHENTICATION",
	}
}

// GetMappingAuthenticationPolicyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAuthenticationPolicyTypeEnum(val string) (AuthenticationPolicyTypeEnum, bool) {
	enum, ok := mappingAuthenticationPolicyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
