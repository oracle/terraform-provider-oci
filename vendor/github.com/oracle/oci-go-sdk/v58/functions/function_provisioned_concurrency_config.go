// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Functions Service API
//
// API for the Functions service.
//

package functions

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// FunctionProvisionedConcurrencyConfig Define the strategy for provisioned concurrency for the function.
type FunctionProvisionedConcurrencyConfig interface {
}

type functionprovisionedconcurrencyconfig struct {
	JsonData []byte
	Strategy string `json:"strategy"`
}

// UnmarshalJSON unmarshals json
func (m *functionprovisionedconcurrencyconfig) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerfunctionprovisionedconcurrencyconfig functionprovisionedconcurrencyconfig
	s := struct {
		Model Unmarshalerfunctionprovisionedconcurrencyconfig
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Strategy = s.Model.Strategy

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *functionprovisionedconcurrencyconfig) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Strategy {
	case "NONE":
		mm := NoneProvisionedConcurrencyConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CONSTANT":
		mm := ConstantProvisionedConcurrencyConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m functionprovisionedconcurrencyconfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m functionprovisionedconcurrencyconfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FunctionProvisionedConcurrencyConfigStrategyEnum Enum with underlying type: string
type FunctionProvisionedConcurrencyConfigStrategyEnum string

// Set of constants representing the allowable values for FunctionProvisionedConcurrencyConfigStrategyEnum
const (
	FunctionProvisionedConcurrencyConfigStrategyConstant FunctionProvisionedConcurrencyConfigStrategyEnum = "CONSTANT"
	FunctionProvisionedConcurrencyConfigStrategyNone     FunctionProvisionedConcurrencyConfigStrategyEnum = "NONE"
)

var mappingFunctionProvisionedConcurrencyConfigStrategyEnum = map[string]FunctionProvisionedConcurrencyConfigStrategyEnum{
	"CONSTANT": FunctionProvisionedConcurrencyConfigStrategyConstant,
	"NONE":     FunctionProvisionedConcurrencyConfigStrategyNone,
}

// GetFunctionProvisionedConcurrencyConfigStrategyEnumValues Enumerates the set of values for FunctionProvisionedConcurrencyConfigStrategyEnum
func GetFunctionProvisionedConcurrencyConfigStrategyEnumValues() []FunctionProvisionedConcurrencyConfigStrategyEnum {
	values := make([]FunctionProvisionedConcurrencyConfigStrategyEnum, 0)
	for _, v := range mappingFunctionProvisionedConcurrencyConfigStrategyEnum {
		values = append(values, v)
	}
	return values
}

// GetFunctionProvisionedConcurrencyConfigStrategyEnumStringValues Enumerates the set of values in String for FunctionProvisionedConcurrencyConfigStrategyEnum
func GetFunctionProvisionedConcurrencyConfigStrategyEnumStringValues() []string {
	return []string{
		"CONSTANT",
		"NONE",
	}
}
