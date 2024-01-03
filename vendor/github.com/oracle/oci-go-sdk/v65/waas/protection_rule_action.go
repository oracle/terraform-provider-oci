// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ProtectionRuleAction A protection rule key and the associated action to apply to that rule.
type ProtectionRuleAction struct {

	// The unique key of the protection rule.
	Key *string `mandatory:"true" json:"key"`

	// The action to apply to the protection rule. If unspecified, defaults to `OFF`.
	Action ProtectionRuleActionActionEnum `mandatory:"true" json:"action"`

	// The types of requests excluded from the protection rule action. If the requests matches the criteria in the `exclusions`, the protection rule action will not be executed.
	Exclusions []ProtectionRuleExclusion `mandatory:"false" json:"exclusions"`
}

func (m ProtectionRuleAction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ProtectionRuleAction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingProtectionRuleActionActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetProtectionRuleActionActionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ProtectionRuleActionActionEnum Enum with underlying type: string
type ProtectionRuleActionActionEnum string

// Set of constants representing the allowable values for ProtectionRuleActionActionEnum
const (
	ProtectionRuleActionActionOff    ProtectionRuleActionActionEnum = "OFF"
	ProtectionRuleActionActionDetect ProtectionRuleActionActionEnum = "DETECT"
	ProtectionRuleActionActionBlock  ProtectionRuleActionActionEnum = "BLOCK"
)

var mappingProtectionRuleActionActionEnum = map[string]ProtectionRuleActionActionEnum{
	"OFF":    ProtectionRuleActionActionOff,
	"DETECT": ProtectionRuleActionActionDetect,
	"BLOCK":  ProtectionRuleActionActionBlock,
}

var mappingProtectionRuleActionActionEnumLowerCase = map[string]ProtectionRuleActionActionEnum{
	"off":    ProtectionRuleActionActionOff,
	"detect": ProtectionRuleActionActionDetect,
	"block":  ProtectionRuleActionActionBlock,
}

// GetProtectionRuleActionActionEnumValues Enumerates the set of values for ProtectionRuleActionActionEnum
func GetProtectionRuleActionActionEnumValues() []ProtectionRuleActionActionEnum {
	values := make([]ProtectionRuleActionActionEnum, 0)
	for _, v := range mappingProtectionRuleActionActionEnum {
		values = append(values, v)
	}
	return values
}

// GetProtectionRuleActionActionEnumStringValues Enumerates the set of values in String for ProtectionRuleActionActionEnum
func GetProtectionRuleActionActionEnumStringValues() []string {
	return []string{
		"OFF",
		"DETECT",
		"BLOCK",
	}
}

// GetMappingProtectionRuleActionActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProtectionRuleActionActionEnum(val string) (ProtectionRuleActionActionEnum, bool) {
	enum, ok := mappingProtectionRuleActionActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
