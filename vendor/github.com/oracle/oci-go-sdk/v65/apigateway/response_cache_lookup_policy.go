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

// ResponseCacheLookupPolicy Base policy for Response Cache lookup.
type ResponseCacheLookupPolicy interface {

	// Whether this policy is currently enabled.
	GetIsEnabled() *bool

	// Set true to allow caching responses where the request has an Authorization header. Ensure you have configured your
	// cache key additions to get the level of isolation across authenticated requests that you require.
	// When false, any request with an Authorization header will not be stored in the Response Cache.
	// If using the CustomAuthenticationPolicy then the tokenHeader/tokenQueryParam are also subject to this check.
	GetIsPrivateCachingEnabled() *bool
}

type responsecachelookuppolicy struct {
	JsonData                []byte
	IsEnabled               *bool  `mandatory:"false" json:"isEnabled"`
	IsPrivateCachingEnabled *bool  `mandatory:"false" json:"isPrivateCachingEnabled"`
	Type                    string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *responsecachelookuppolicy) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerresponsecachelookuppolicy responsecachelookuppolicy
	s := struct {
		Model Unmarshalerresponsecachelookuppolicy
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.IsEnabled = s.Model.IsEnabled
	m.IsPrivateCachingEnabled = s.Model.IsPrivateCachingEnabled
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *responsecachelookuppolicy) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "SIMPLE_LOOKUP_POLICY":
		mm := SimpleLookupPolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ResponseCacheLookupPolicy: %s.", m.Type)
		return *m, nil
	}
}

// GetIsEnabled returns IsEnabled
func (m responsecachelookuppolicy) GetIsEnabled() *bool {
	return m.IsEnabled
}

// GetIsPrivateCachingEnabled returns IsPrivateCachingEnabled
func (m responsecachelookuppolicy) GetIsPrivateCachingEnabled() *bool {
	return m.IsPrivateCachingEnabled
}

func (m responsecachelookuppolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m responsecachelookuppolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ResponseCacheLookupPolicyTypeEnum Enum with underlying type: string
type ResponseCacheLookupPolicyTypeEnum string

// Set of constants representing the allowable values for ResponseCacheLookupPolicyTypeEnum
const (
	ResponseCacheLookupPolicyTypeSimpleLookupPolicy ResponseCacheLookupPolicyTypeEnum = "SIMPLE_LOOKUP_POLICY"
)

var mappingResponseCacheLookupPolicyTypeEnum = map[string]ResponseCacheLookupPolicyTypeEnum{
	"SIMPLE_LOOKUP_POLICY": ResponseCacheLookupPolicyTypeSimpleLookupPolicy,
}

var mappingResponseCacheLookupPolicyTypeEnumLowerCase = map[string]ResponseCacheLookupPolicyTypeEnum{
	"simple_lookup_policy": ResponseCacheLookupPolicyTypeSimpleLookupPolicy,
}

// GetResponseCacheLookupPolicyTypeEnumValues Enumerates the set of values for ResponseCacheLookupPolicyTypeEnum
func GetResponseCacheLookupPolicyTypeEnumValues() []ResponseCacheLookupPolicyTypeEnum {
	values := make([]ResponseCacheLookupPolicyTypeEnum, 0)
	for _, v := range mappingResponseCacheLookupPolicyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetResponseCacheLookupPolicyTypeEnumStringValues Enumerates the set of values in String for ResponseCacheLookupPolicyTypeEnum
func GetResponseCacheLookupPolicyTypeEnumStringValues() []string {
	return []string{
		"SIMPLE_LOOKUP_POLICY",
	}
}

// GetMappingResponseCacheLookupPolicyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResponseCacheLookupPolicyTypeEnum(val string) (ResponseCacheLookupPolicyTypeEnum, bool) {
	enum, ok := mappingResponseCacheLookupPolicyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
