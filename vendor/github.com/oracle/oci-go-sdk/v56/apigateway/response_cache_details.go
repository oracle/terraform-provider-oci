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
	"github.com/oracle/oci-go-sdk/v56/common"
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
		return *m, nil
	}
}

func (m responsecachedetails) String() string {
	return common.PointerString(m)
}

// ResponseCacheDetailsTypeEnum Enum with underlying type: string
type ResponseCacheDetailsTypeEnum string

// Set of constants representing the allowable values for ResponseCacheDetailsTypeEnum
const (
	ResponseCacheDetailsTypeExternalRespCache ResponseCacheDetailsTypeEnum = "EXTERNAL_RESP_CACHE"
	ResponseCacheDetailsTypeNone              ResponseCacheDetailsTypeEnum = "NONE"
)

var mappingResponseCacheDetailsType = map[string]ResponseCacheDetailsTypeEnum{
	"EXTERNAL_RESP_CACHE": ResponseCacheDetailsTypeExternalRespCache,
	"NONE":                ResponseCacheDetailsTypeNone,
}

// GetResponseCacheDetailsTypeEnumValues Enumerates the set of values for ResponseCacheDetailsTypeEnum
func GetResponseCacheDetailsTypeEnumValues() []ResponseCacheDetailsTypeEnum {
	values := make([]ResponseCacheDetailsTypeEnum, 0)
	for _, v := range mappingResponseCacheDetailsType {
		values = append(values, v)
	}
	return values
}
