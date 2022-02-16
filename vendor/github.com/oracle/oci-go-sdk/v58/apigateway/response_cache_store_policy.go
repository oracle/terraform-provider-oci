// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ResponseCacheStorePolicy Base policy for how a response from a backend is cached in the Response Cache.
type ResponseCacheStorePolicy interface {
}

type responsecachestorepolicy struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *responsecachestorepolicy) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerresponsecachestorepolicy responsecachestorepolicy
	s := struct {
		Model Unmarshalerresponsecachestorepolicy
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *responsecachestorepolicy) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "FIXED_TTL_STORE_POLICY":
		mm := FixedTtlResponseCacheStorePolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m responsecachestorepolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m responsecachestorepolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ResponseCacheStorePolicyTypeEnum Enum with underlying type: string
type ResponseCacheStorePolicyTypeEnum string

// Set of constants representing the allowable values for ResponseCacheStorePolicyTypeEnum
const (
	ResponseCacheStorePolicyTypeFixedTtlStorePolicy ResponseCacheStorePolicyTypeEnum = "FIXED_TTL_STORE_POLICY"
)

var mappingResponseCacheStorePolicyTypeEnum = map[string]ResponseCacheStorePolicyTypeEnum{
	"FIXED_TTL_STORE_POLICY": ResponseCacheStorePolicyTypeFixedTtlStorePolicy,
}

// GetResponseCacheStorePolicyTypeEnumValues Enumerates the set of values for ResponseCacheStorePolicyTypeEnum
func GetResponseCacheStorePolicyTypeEnumValues() []ResponseCacheStorePolicyTypeEnum {
	values := make([]ResponseCacheStorePolicyTypeEnum, 0)
	for _, v := range mappingResponseCacheStorePolicyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetResponseCacheStorePolicyTypeEnumStringValues Enumerates the set of values in String for ResponseCacheStorePolicyTypeEnum
func GetResponseCacheStorePolicyTypeEnumStringValues() []string {
	return []string{
		"FIXED_TTL_STORE_POLICY",
	}
}

// GetMappingResponseCacheStorePolicyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResponseCacheStorePolicyTypeEnum(val string) (ResponseCacheStorePolicyTypeEnum, bool) {
	mappingResponseCacheStorePolicyTypeEnumIgnoreCase := make(map[string]ResponseCacheStorePolicyTypeEnum)
	for k, v := range mappingResponseCacheStorePolicyTypeEnum {
		mappingResponseCacheStorePolicyTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingResponseCacheStorePolicyTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
