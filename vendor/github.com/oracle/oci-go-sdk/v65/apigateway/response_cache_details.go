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

// ResponseCacheDetails Base Gateway response cache.
type ResponseCacheDetails interface {
}

type responsecachedetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *responsecachedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerresponsecachedetails responsecachedetails
	s := struct {
		Model Unmarshalerresponsecachedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *responsecachedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "EXTERNAL_RESP_CACHE":
		mm := ExternalRespCache{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NONE":
		mm := NoCache{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ResponseCacheDetails: %s.", m.Type)
		return *m, nil
	}
}

func (m responsecachedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m responsecachedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ResponseCacheDetailsTypeEnum Enum with underlying type: string
type ResponseCacheDetailsTypeEnum string

// Set of constants representing the allowable values for ResponseCacheDetailsTypeEnum
const (
	ResponseCacheDetailsTypeExternalRespCache ResponseCacheDetailsTypeEnum = "EXTERNAL_RESP_CACHE"
	ResponseCacheDetailsTypeNone              ResponseCacheDetailsTypeEnum = "NONE"
)

var mappingResponseCacheDetailsTypeEnum = map[string]ResponseCacheDetailsTypeEnum{
	"EXTERNAL_RESP_CACHE": ResponseCacheDetailsTypeExternalRespCache,
	"NONE":                ResponseCacheDetailsTypeNone,
}

var mappingResponseCacheDetailsTypeEnumLowerCase = map[string]ResponseCacheDetailsTypeEnum{
	"external_resp_cache": ResponseCacheDetailsTypeExternalRespCache,
	"none":                ResponseCacheDetailsTypeNone,
}

// GetResponseCacheDetailsTypeEnumValues Enumerates the set of values for ResponseCacheDetailsTypeEnum
func GetResponseCacheDetailsTypeEnumValues() []ResponseCacheDetailsTypeEnum {
	values := make([]ResponseCacheDetailsTypeEnum, 0)
	for _, v := range mappingResponseCacheDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetResponseCacheDetailsTypeEnumStringValues Enumerates the set of values in String for ResponseCacheDetailsTypeEnum
func GetResponseCacheDetailsTypeEnumStringValues() []string {
	return []string{
		"EXTERNAL_RESP_CACHE",
		"NONE",
	}
}

// GetMappingResponseCacheDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResponseCacheDetailsTypeEnum(val string) (ResponseCacheDetailsTypeEnum, bool) {
	enum, ok := mappingResponseCacheDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
