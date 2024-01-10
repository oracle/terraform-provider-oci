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

// SourceUriDetails Auth endpoint details.
type SourceUriDetails interface {
}

type sourceuridetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *sourceuridetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersourceuridetails sourceuridetails
	s := struct {
		Model Unmarshalersourceuridetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *sourceuridetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "DISCOVERY_URI":
		mm := DiscoveryUriSourceUriDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VALIDATION_BLOCK":
		mm := ValidationBlockSourceUriDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for SourceUriDetails: %s.", m.Type)
		return *m, nil
	}
}

func (m sourceuridetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m sourceuridetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SourceUriDetailsTypeEnum Enum with underlying type: string
type SourceUriDetailsTypeEnum string

// Set of constants representing the allowable values for SourceUriDetailsTypeEnum
const (
	SourceUriDetailsTypeDiscoveryUri    SourceUriDetailsTypeEnum = "DISCOVERY_URI"
	SourceUriDetailsTypeValidationBlock SourceUriDetailsTypeEnum = "VALIDATION_BLOCK"
)

var mappingSourceUriDetailsTypeEnum = map[string]SourceUriDetailsTypeEnum{
	"DISCOVERY_URI":    SourceUriDetailsTypeDiscoveryUri,
	"VALIDATION_BLOCK": SourceUriDetailsTypeValidationBlock,
}

var mappingSourceUriDetailsTypeEnumLowerCase = map[string]SourceUriDetailsTypeEnum{
	"discovery_uri":    SourceUriDetailsTypeDiscoveryUri,
	"validation_block": SourceUriDetailsTypeValidationBlock,
}

// GetSourceUriDetailsTypeEnumValues Enumerates the set of values for SourceUriDetailsTypeEnum
func GetSourceUriDetailsTypeEnumValues() []SourceUriDetailsTypeEnum {
	values := make([]SourceUriDetailsTypeEnum, 0)
	for _, v := range mappingSourceUriDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSourceUriDetailsTypeEnumStringValues Enumerates the set of values in String for SourceUriDetailsTypeEnum
func GetSourceUriDetailsTypeEnumStringValues() []string {
	return []string{
		"DISCOVERY_URI",
		"VALIDATION_BLOCK",
	}
}

// GetMappingSourceUriDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSourceUriDetailsTypeEnum(val string) (SourceUriDetailsTypeEnum, bool) {
	enum, ok := mappingSourceUriDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
