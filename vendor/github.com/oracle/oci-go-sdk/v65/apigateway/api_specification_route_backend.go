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

// ApiSpecificationRouteBackend The backend to forward requests to.
type ApiSpecificationRouteBackend interface {
}

type apispecificationroutebackend struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *apispecificationroutebackend) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerapispecificationroutebackend apispecificationroutebackend
	s := struct {
		Model Unmarshalerapispecificationroutebackend
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *apispecificationroutebackend) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "OAUTH2_LOGOUT_BACKEND":
		mm := OAuth2LogoutBackend{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HTTP_BACKEND":
		mm := HttpBackend{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_FUNCTIONS_BACKEND":
		mm := OracleFunctionBackend{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "STOCK_RESPONSE_BACKEND":
		mm := StockResponseBackend{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DYNAMIC_ROUTING_BACKEND":
		mm := DynamicRoutingBackend{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ApiSpecificationRouteBackend: %s.", m.Type)
		return *m, nil
	}
}

func (m apispecificationroutebackend) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m apispecificationroutebackend) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ApiSpecificationRouteBackendTypeEnum Enum with underlying type: string
type ApiSpecificationRouteBackendTypeEnum string

// Set of constants representing the allowable values for ApiSpecificationRouteBackendTypeEnum
const (
	ApiSpecificationRouteBackendTypeOracleFunctionsBackend ApiSpecificationRouteBackendTypeEnum = "ORACLE_FUNCTIONS_BACKEND"
	ApiSpecificationRouteBackendTypeHttpBackend            ApiSpecificationRouteBackendTypeEnum = "HTTP_BACKEND"
	ApiSpecificationRouteBackendTypeStockResponseBackend   ApiSpecificationRouteBackendTypeEnum = "STOCK_RESPONSE_BACKEND"
	ApiSpecificationRouteBackendTypeDynamicRoutingBackend  ApiSpecificationRouteBackendTypeEnum = "DYNAMIC_ROUTING_BACKEND"
	ApiSpecificationRouteBackendTypeOauth2LogoutBackend    ApiSpecificationRouteBackendTypeEnum = "OAUTH2_LOGOUT_BACKEND"
)

var mappingApiSpecificationRouteBackendTypeEnum = map[string]ApiSpecificationRouteBackendTypeEnum{
	"ORACLE_FUNCTIONS_BACKEND": ApiSpecificationRouteBackendTypeOracleFunctionsBackend,
	"HTTP_BACKEND":             ApiSpecificationRouteBackendTypeHttpBackend,
	"STOCK_RESPONSE_BACKEND":   ApiSpecificationRouteBackendTypeStockResponseBackend,
	"DYNAMIC_ROUTING_BACKEND":  ApiSpecificationRouteBackendTypeDynamicRoutingBackend,
	"OAUTH2_LOGOUT_BACKEND":    ApiSpecificationRouteBackendTypeOauth2LogoutBackend,
}

var mappingApiSpecificationRouteBackendTypeEnumLowerCase = map[string]ApiSpecificationRouteBackendTypeEnum{
	"oracle_functions_backend": ApiSpecificationRouteBackendTypeOracleFunctionsBackend,
	"http_backend":             ApiSpecificationRouteBackendTypeHttpBackend,
	"stock_response_backend":   ApiSpecificationRouteBackendTypeStockResponseBackend,
	"dynamic_routing_backend":  ApiSpecificationRouteBackendTypeDynamicRoutingBackend,
	"oauth2_logout_backend":    ApiSpecificationRouteBackendTypeOauth2LogoutBackend,
}

// GetApiSpecificationRouteBackendTypeEnumValues Enumerates the set of values for ApiSpecificationRouteBackendTypeEnum
func GetApiSpecificationRouteBackendTypeEnumValues() []ApiSpecificationRouteBackendTypeEnum {
	values := make([]ApiSpecificationRouteBackendTypeEnum, 0)
	for _, v := range mappingApiSpecificationRouteBackendTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetApiSpecificationRouteBackendTypeEnumStringValues Enumerates the set of values in String for ApiSpecificationRouteBackendTypeEnum
func GetApiSpecificationRouteBackendTypeEnumStringValues() []string {
	return []string{
		"ORACLE_FUNCTIONS_BACKEND",
		"HTTP_BACKEND",
		"STOCK_RESPONSE_BACKEND",
		"DYNAMIC_ROUTING_BACKEND",
		"OAUTH2_LOGOUT_BACKEND",
	}
}

// GetMappingApiSpecificationRouteBackendTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingApiSpecificationRouteBackendTypeEnum(val string) (ApiSpecificationRouteBackendTypeEnum, bool) {
	enum, ok := mappingApiSpecificationRouteBackendTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
