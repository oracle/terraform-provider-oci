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

// SelectionSourcePolicy The type of selector to use when dynamically routing and dynamically authenticating requests.
type SelectionSourcePolicy interface {
}

type selectionsourcepolicy struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *selectionsourcepolicy) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerselectionsourcepolicy selectionsourcepolicy
	s := struct {
		Model Unmarshalerselectionsourcepolicy
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *selectionsourcepolicy) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "SINGLE":
		mm := SingleSelectionSourcePolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for SelectionSourcePolicy: %s.", m.Type)
		return *m, nil
	}
}

func (m selectionsourcepolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m selectionsourcepolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SelectionSourcePolicyTypeEnum Enum with underlying type: string
type SelectionSourcePolicyTypeEnum string

// Set of constants representing the allowable values for SelectionSourcePolicyTypeEnum
const (
	SelectionSourcePolicyTypeSingle SelectionSourcePolicyTypeEnum = "SINGLE"
)

var mappingSelectionSourcePolicyTypeEnum = map[string]SelectionSourcePolicyTypeEnum{
	"SINGLE": SelectionSourcePolicyTypeSingle,
}

var mappingSelectionSourcePolicyTypeEnumLowerCase = map[string]SelectionSourcePolicyTypeEnum{
	"single": SelectionSourcePolicyTypeSingle,
}

// GetSelectionSourcePolicyTypeEnumValues Enumerates the set of values for SelectionSourcePolicyTypeEnum
func GetSelectionSourcePolicyTypeEnumValues() []SelectionSourcePolicyTypeEnum {
	values := make([]SelectionSourcePolicyTypeEnum, 0)
	for _, v := range mappingSelectionSourcePolicyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSelectionSourcePolicyTypeEnumStringValues Enumerates the set of values in String for SelectionSourcePolicyTypeEnum
func GetSelectionSourcePolicyTypeEnumStringValues() []string {
	return []string{
		"SINGLE",
	}
}

// GetMappingSelectionSourcePolicyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSelectionSourcePolicyTypeEnum(val string) (SelectionSourcePolicyTypeEnum, bool) {
	enum, ok := mappingSelectionSourcePolicyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
