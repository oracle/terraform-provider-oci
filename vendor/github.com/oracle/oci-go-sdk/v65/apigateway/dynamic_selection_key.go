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

// DynamicSelectionKey Base policy for defining how to match the context variable in an incoming request with selection keys when dynamically routing and dynamically authenticating requests.
type DynamicSelectionKey interface {

	// Name assigned to the branch.
	GetName() *string

	// Specifies whether to use the route or authentication server associated with this selection key as the default. The default is used if the value of a context variable in an incoming request does not match any of the other selection key values when dynamically routing and dynamically authenticating requests.
	GetIsDefault() *bool
}

type dynamicselectionkey struct {
	JsonData  []byte
	IsDefault *bool   `mandatory:"false" json:"isDefault"`
	Name      *string `mandatory:"true" json:"name"`
	Type      string  `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *dynamicselectionkey) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdynamicselectionkey dynamicselectionkey
	s := struct {
		Model Unmarshalerdynamicselectionkey
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.IsDefault = s.Model.IsDefault
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *dynamicselectionkey) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "WILDCARD":
		mm := WildcardSelectionKey{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ANY_OF":
		mm := AnyOfSelectionKey{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DynamicSelectionKey: %s.", m.Type)
		return *m, nil
	}
}

// GetIsDefault returns IsDefault
func (m dynamicselectionkey) GetIsDefault() *bool {
	return m.IsDefault
}

// GetName returns Name
func (m dynamicselectionkey) GetName() *string {
	return m.Name
}

func (m dynamicselectionkey) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m dynamicselectionkey) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DynamicSelectionKeyTypeEnum Enum with underlying type: string
type DynamicSelectionKeyTypeEnum string

// Set of constants representing the allowable values for DynamicSelectionKeyTypeEnum
const (
	DynamicSelectionKeyTypeAnyOf    DynamicSelectionKeyTypeEnum = "ANY_OF"
	DynamicSelectionKeyTypeWildcard DynamicSelectionKeyTypeEnum = "WILDCARD"
)

var mappingDynamicSelectionKeyTypeEnum = map[string]DynamicSelectionKeyTypeEnum{
	"ANY_OF":   DynamicSelectionKeyTypeAnyOf,
	"WILDCARD": DynamicSelectionKeyTypeWildcard,
}

var mappingDynamicSelectionKeyTypeEnumLowerCase = map[string]DynamicSelectionKeyTypeEnum{
	"any_of":   DynamicSelectionKeyTypeAnyOf,
	"wildcard": DynamicSelectionKeyTypeWildcard,
}

// GetDynamicSelectionKeyTypeEnumValues Enumerates the set of values for DynamicSelectionKeyTypeEnum
func GetDynamicSelectionKeyTypeEnumValues() []DynamicSelectionKeyTypeEnum {
	values := make([]DynamicSelectionKeyTypeEnum, 0)
	for _, v := range mappingDynamicSelectionKeyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDynamicSelectionKeyTypeEnumStringValues Enumerates the set of values in String for DynamicSelectionKeyTypeEnum
func GetDynamicSelectionKeyTypeEnumStringValues() []string {
	return []string{
		"ANY_OF",
		"WILDCARD",
	}
}

// GetMappingDynamicSelectionKeyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDynamicSelectionKeyTypeEnum(val string) (DynamicSelectionKeyTypeEnum, bool) {
	enum, ok := mappingDynamicSelectionKeyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
