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

// Rule An object that represents an action to apply to a listener.
type Rule interface {
}

type rule struct {
	JsonData []byte
	Action   string `json:"action"`
}

// UnmarshalJSON unmarshals json
func (m *rule) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerrule rule
	s := struct {
		Model Unmarshalerrule
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Action = s.Model.Action

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *rule) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Action {
	case "ALLOW":
		mm := AllowRule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for Rule: %s.", m.Action)
		return *m, nil
	}
}

func (m rule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m rule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RuleActionEnum Enum with underlying type: string
type RuleActionEnum string

// Set of constants representing the allowable values for RuleActionEnum
const (
	RuleActionAllow RuleActionEnum = "ALLOW"
)

var mappingRuleActionEnum = map[string]RuleActionEnum{
	"ALLOW": RuleActionAllow,
}

var mappingRuleActionEnumLowerCase = map[string]RuleActionEnum{
	"allow": RuleActionAllow,
}

// GetRuleActionEnumValues Enumerates the set of values for RuleActionEnum
func GetRuleActionEnumValues() []RuleActionEnum {
	values := make([]RuleActionEnum, 0)
	for _, v := range mappingRuleActionEnum {
		values = append(values, v)
	}
	return values
}

// GetRuleActionEnumStringValues Enumerates the set of values in String for RuleActionEnum
func GetRuleActionEnumStringValues() []string {
	return []string{
		"ALLOW",
	}
}

// GetMappingRuleActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRuleActionEnum(val string) (RuleActionEnum, bool) {
	enum, ok := mappingRuleActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
