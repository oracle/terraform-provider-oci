// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SubsetRuleEntry The details of the subset rule
type SubsetRuleEntry interface {
}

type subsetruleentry struct {
	JsonData []byte
	RuleType string `json:"ruleType"`
}

// UnmarshalJSON unmarshals json
func (m *subsetruleentry) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersubsetruleentry subsetruleentry
	s := struct {
		Model Unmarshalersubsetruleentry
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.RuleType = s.Model.RuleType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *subsetruleentry) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.RuleType {
	case "PARTITION":
		mm := PartitionSubsetRuleEntry{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PERCENT":
		mm := PercentSubsetRuleEntry{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CONDITION":
		mm := ConditionSubsetRuleEntry{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for SubsetRuleEntry: %s.", m.RuleType)
		return *m, nil
	}
}

func (m subsetruleentry) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m subsetruleentry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SubsetRuleEntryRuleTypeEnum Enum with underlying type: string
type SubsetRuleEntryRuleTypeEnum string

// Set of constants representing the allowable values for SubsetRuleEntryRuleTypeEnum
const (
	SubsetRuleEntryRuleTypePercent   SubsetRuleEntryRuleTypeEnum = "PERCENT"
	SubsetRuleEntryRuleTypeCondition SubsetRuleEntryRuleTypeEnum = "CONDITION"
	SubsetRuleEntryRuleTypePartition SubsetRuleEntryRuleTypeEnum = "PARTITION"
)

var mappingSubsetRuleEntryRuleTypeEnum = map[string]SubsetRuleEntryRuleTypeEnum{
	"PERCENT":   SubsetRuleEntryRuleTypePercent,
	"CONDITION": SubsetRuleEntryRuleTypeCondition,
	"PARTITION": SubsetRuleEntryRuleTypePartition,
}

var mappingSubsetRuleEntryRuleTypeEnumLowerCase = map[string]SubsetRuleEntryRuleTypeEnum{
	"percent":   SubsetRuleEntryRuleTypePercent,
	"condition": SubsetRuleEntryRuleTypeCondition,
	"partition": SubsetRuleEntryRuleTypePartition,
}

// GetSubsetRuleEntryRuleTypeEnumValues Enumerates the set of values for SubsetRuleEntryRuleTypeEnum
func GetSubsetRuleEntryRuleTypeEnumValues() []SubsetRuleEntryRuleTypeEnum {
	values := make([]SubsetRuleEntryRuleTypeEnum, 0)
	for _, v := range mappingSubsetRuleEntryRuleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSubsetRuleEntryRuleTypeEnumStringValues Enumerates the set of values in String for SubsetRuleEntryRuleTypeEnum
func GetSubsetRuleEntryRuleTypeEnumStringValues() []string {
	return []string{
		"PERCENT",
		"CONDITION",
		"PARTITION",
	}
}

// GetMappingSubsetRuleEntryRuleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSubsetRuleEntryRuleTypeEnum(val string) (SubsetRuleEntryRuleTypeEnum, bool) {
	enum, ok := mappingSubsetRuleEntryRuleTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
