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

// ClientAppDetails Client App Credential details.
type ClientAppDetails interface {
}

type clientappdetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *clientappdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerclientappdetails clientappdetails
	s := struct {
		Model Unmarshalerclientappdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *clientappdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "CUSTOM":
		mm := CustomClientAppDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VALIDATION_BLOCK":
		mm := ValidationBlockClientAppDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ClientAppDetails: %s.", m.Type)
		return *m, nil
	}
}

func (m clientappdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m clientappdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ClientAppDetailsTypeEnum Enum with underlying type: string
type ClientAppDetailsTypeEnum string

// Set of constants representing the allowable values for ClientAppDetailsTypeEnum
const (
	ClientAppDetailsTypeValidationBlock ClientAppDetailsTypeEnum = "VALIDATION_BLOCK"
	ClientAppDetailsTypeCustom          ClientAppDetailsTypeEnum = "CUSTOM"
)

var mappingClientAppDetailsTypeEnum = map[string]ClientAppDetailsTypeEnum{
	"VALIDATION_BLOCK": ClientAppDetailsTypeValidationBlock,
	"CUSTOM":           ClientAppDetailsTypeCustom,
}

var mappingClientAppDetailsTypeEnumLowerCase = map[string]ClientAppDetailsTypeEnum{
	"validation_block": ClientAppDetailsTypeValidationBlock,
	"custom":           ClientAppDetailsTypeCustom,
}

// GetClientAppDetailsTypeEnumValues Enumerates the set of values for ClientAppDetailsTypeEnum
func GetClientAppDetailsTypeEnumValues() []ClientAppDetailsTypeEnum {
	values := make([]ClientAppDetailsTypeEnum, 0)
	for _, v := range mappingClientAppDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetClientAppDetailsTypeEnumStringValues Enumerates the set of values in String for ClientAppDetailsTypeEnum
func GetClientAppDetailsTypeEnumStringValues() []string {
	return []string{
		"VALIDATION_BLOCK",
		"CUSTOM",
	}
}

// GetMappingClientAppDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingClientAppDetailsTypeEnum(val string) (ClientAppDetailsTypeEnum, bool) {
	enum, ok := mappingClientAppDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
