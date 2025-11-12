// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Prompt Prompt.
type Prompt struct {

	// acceptance condition.
	AcceptanceCondition PromptAcceptanceConditionEnum `mandatory:"true" json:"acceptanceCondition"`

	// List of choice.
	ChoiceGroup []Choice `mandatory:"false" json:"choiceGroup"`
}

func (m Prompt) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Prompt) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPromptAcceptanceConditionEnum(string(m.AcceptanceCondition)); !ok && m.AcceptanceCondition != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AcceptanceCondition: %s. Supported values are: %s.", m.AcceptanceCondition, strings.Join(GetPromptAcceptanceConditionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PromptAcceptanceConditionEnum Enum with underlying type: string
type PromptAcceptanceConditionEnum string

// Set of constants representing the allowable values for PromptAcceptanceConditionEnum
const (
	PromptAcceptanceConditionAny PromptAcceptanceConditionEnum = "ANY"
	PromptAcceptanceConditionAll PromptAcceptanceConditionEnum = "ALL"
)

var mappingPromptAcceptanceConditionEnum = map[string]PromptAcceptanceConditionEnum{
	"ANY": PromptAcceptanceConditionAny,
	"ALL": PromptAcceptanceConditionAll,
}

var mappingPromptAcceptanceConditionEnumLowerCase = map[string]PromptAcceptanceConditionEnum{
	"any": PromptAcceptanceConditionAny,
	"all": PromptAcceptanceConditionAll,
}

// GetPromptAcceptanceConditionEnumValues Enumerates the set of values for PromptAcceptanceConditionEnum
func GetPromptAcceptanceConditionEnumValues() []PromptAcceptanceConditionEnum {
	values := make([]PromptAcceptanceConditionEnum, 0)
	for _, v := range mappingPromptAcceptanceConditionEnum {
		values = append(values, v)
	}
	return values
}

// GetPromptAcceptanceConditionEnumStringValues Enumerates the set of values in String for PromptAcceptanceConditionEnum
func GetPromptAcceptanceConditionEnumStringValues() []string {
	return []string{
		"ANY",
		"ALL",
	}
}

// GetMappingPromptAcceptanceConditionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPromptAcceptanceConditionEnum(val string) (PromptAcceptanceConditionEnum, bool) {
	enum, ok := mappingPromptAcceptanceConditionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
