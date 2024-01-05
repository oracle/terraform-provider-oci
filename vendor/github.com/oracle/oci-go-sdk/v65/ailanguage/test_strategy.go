// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Language API
//
// OCI Language Service solutions can help enterprise customers integrate AI into their products immediately using our proven,
// pre-trained and custom models or containers, without a need to set up an house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI and ML operations, which shortens the time to market.
//

package ailanguage

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TestStrategy Possible strategy as testing and validation(optional) dataset.
type TestStrategy interface {
}

type teststrategy struct {
	JsonData     []byte
	StrategyType string `json:"strategyType"`
}

// UnmarshalJSON unmarshals json
func (m *teststrategy) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerteststrategy teststrategy
	s := struct {
		Model Unmarshalerteststrategy
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.StrategyType = s.Model.StrategyType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *teststrategy) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.StrategyType {
	case "TEST_AND_VALIDATION_DATASET":
		mm := TestAndValidationDatasetStrategy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for TestStrategy: %s.", m.StrategyType)
		return *m, nil
	}
}

func (m teststrategy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m teststrategy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TestStrategyStrategyTypeEnum Enum with underlying type: string
type TestStrategyStrategyTypeEnum string

// Set of constants representing the allowable values for TestStrategyStrategyTypeEnum
const (
	TestStrategyStrategyTypeTestAndValidationDataset TestStrategyStrategyTypeEnum = "TEST_AND_VALIDATION_DATASET"
)

var mappingTestStrategyStrategyTypeEnum = map[string]TestStrategyStrategyTypeEnum{
	"TEST_AND_VALIDATION_DATASET": TestStrategyStrategyTypeTestAndValidationDataset,
}

var mappingTestStrategyStrategyTypeEnumLowerCase = map[string]TestStrategyStrategyTypeEnum{
	"test_and_validation_dataset": TestStrategyStrategyTypeTestAndValidationDataset,
}

// GetTestStrategyStrategyTypeEnumValues Enumerates the set of values for TestStrategyStrategyTypeEnum
func GetTestStrategyStrategyTypeEnumValues() []TestStrategyStrategyTypeEnum {
	values := make([]TestStrategyStrategyTypeEnum, 0)
	for _, v := range mappingTestStrategyStrategyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTestStrategyStrategyTypeEnumStringValues Enumerates the set of values in String for TestStrategyStrategyTypeEnum
func GetTestStrategyStrategyTypeEnumStringValues() []string {
	return []string{
		"TEST_AND_VALIDATION_DATASET",
	}
}

// GetMappingTestStrategyStrategyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTestStrategyStrategyTypeEnum(val string) (TestStrategyStrategyTypeEnum, bool) {
	enum, ok := mappingTestStrategyStrategyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
