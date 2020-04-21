// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// RuleCondition A condition to apply to an access control rule.
type RuleCondition interface {
}

type rulecondition struct {
	JsonData      []byte
	AttributeName string `json:"attributeName"`
}

// UnmarshalJSON unmarshals json
func (m *rulecondition) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerrulecondition rulecondition
	s := struct {
		Model Unmarshalerrulecondition
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.AttributeName = s.Model.AttributeName

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *rulecondition) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.AttributeName {
	case "SOURCE_VCN_ID":
		mm := SourceVcnIdCondition{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SOURCE_IP_ADDRESS":
		mm := SourceIpAddressCondition{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PATH":
		mm := PathMatchCondition{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SOURCE_VCN_IP_ADDRESS":
		mm := SourceVcnIpAddressCondition{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m rulecondition) String() string {
	return common.PointerString(m)
}

// RuleConditionAttributeNameEnum Enum with underlying type: string
type RuleConditionAttributeNameEnum string

// Set of constants representing the allowable values for RuleConditionAttributeNameEnum
const (
	RuleConditionAttributeNameSourceIpAddress    RuleConditionAttributeNameEnum = "SOURCE_IP_ADDRESS"
	RuleConditionAttributeNameSourceVcnId        RuleConditionAttributeNameEnum = "SOURCE_VCN_ID"
	RuleConditionAttributeNameSourceVcnIpAddress RuleConditionAttributeNameEnum = "SOURCE_VCN_IP_ADDRESS"
	RuleConditionAttributeNamePath               RuleConditionAttributeNameEnum = "PATH"
)

var mappingRuleConditionAttributeName = map[string]RuleConditionAttributeNameEnum{
	"SOURCE_IP_ADDRESS":     RuleConditionAttributeNameSourceIpAddress,
	"SOURCE_VCN_ID":         RuleConditionAttributeNameSourceVcnId,
	"SOURCE_VCN_IP_ADDRESS": RuleConditionAttributeNameSourceVcnIpAddress,
	"PATH":                  RuleConditionAttributeNamePath,
}

// GetRuleConditionAttributeNameEnumValues Enumerates the set of values for RuleConditionAttributeNameEnum
func GetRuleConditionAttributeNameEnumValues() []RuleConditionAttributeNameEnum {
	values := make([]RuleConditionAttributeNameEnum, 0)
	for _, v := range mappingRuleConditionAttributeName {
		values = append(values, v)
	}
	return values
}
