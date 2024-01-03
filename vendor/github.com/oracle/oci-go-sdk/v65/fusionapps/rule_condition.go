// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fusion Applications Environment Management API
//
// Use the Fusion Applications Environment Management API to manage the environments where your Fusion Applications run. For more information, see the Fusion Applications Environment Management documentation (https://docs.cloud.oracle.com/iaas/Content/fusion-applications/home.htm).
//

package fusionapps

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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
	case "SOURCE_VCN_IP_ADDRESS":
		mm := SourceVcnIpAddressCondition{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for RuleCondition: %s.", m.AttributeName)
		return *m, nil
	}
}

func (m rulecondition) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m rulecondition) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RuleConditionAttributeNameEnum Enum with underlying type: string
type RuleConditionAttributeNameEnum string

// Set of constants representing the allowable values for RuleConditionAttributeNameEnum
const (
	RuleConditionAttributeNameIpAddress    RuleConditionAttributeNameEnum = "SOURCE_IP_ADDRESS"
	RuleConditionAttributeNameVcnId        RuleConditionAttributeNameEnum = "SOURCE_VCN_ID"
	RuleConditionAttributeNameVcnIpAddress RuleConditionAttributeNameEnum = "SOURCE_VCN_IP_ADDRESS"
)

var mappingRuleConditionAttributeNameEnum = map[string]RuleConditionAttributeNameEnum{
	"SOURCE_IP_ADDRESS":     RuleConditionAttributeNameIpAddress,
	"SOURCE_VCN_ID":         RuleConditionAttributeNameVcnId,
	"SOURCE_VCN_IP_ADDRESS": RuleConditionAttributeNameVcnIpAddress,
}

var mappingRuleConditionAttributeNameEnumLowerCase = map[string]RuleConditionAttributeNameEnum{
	"source_ip_address":     RuleConditionAttributeNameIpAddress,
	"source_vcn_id":         RuleConditionAttributeNameVcnId,
	"source_vcn_ip_address": RuleConditionAttributeNameVcnIpAddress,
}

// GetRuleConditionAttributeNameEnumValues Enumerates the set of values for RuleConditionAttributeNameEnum
func GetRuleConditionAttributeNameEnumValues() []RuleConditionAttributeNameEnum {
	values := make([]RuleConditionAttributeNameEnum, 0)
	for _, v := range mappingRuleConditionAttributeNameEnum {
		values = append(values, v)
	}
	return values
}

// GetRuleConditionAttributeNameEnumStringValues Enumerates the set of values in String for RuleConditionAttributeNameEnum
func GetRuleConditionAttributeNameEnumStringValues() []string {
	return []string{
		"SOURCE_IP_ADDRESS",
		"SOURCE_VCN_ID",
		"SOURCE_VCN_IP_ADDRESS",
	}
}

// GetMappingRuleConditionAttributeNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRuleConditionAttributeNameEnum(val string) (RuleConditionAttributeNameEnum, bool) {
	enum, ok := mappingRuleConditionAttributeNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
