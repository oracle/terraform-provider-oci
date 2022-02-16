// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ResponderRule Definition of ResponderRule.
type ResponderRule struct {

	// Identifier for ResponderRule.
	Id *string `mandatory:"true" json:"id"`

	// ResponderRule Display Name
	DisplayName *string `mandatory:"true" json:"displayName"`

	// ResponderRule Description
	Description *string `mandatory:"true" json:"description"`

	// Type of Responder
	Type ResponderTypeEnum `mandatory:"true" json:"type"`

	// List of Policy
	Policies []string `mandatory:"false" json:"policies"`

	// Supported Execution Modes
	SupportedModes []ResponderRuleSupportedModesEnum `mandatory:"false" json:"supportedModes,omitempty"`

	Details *ResponderRuleDetails `mandatory:"false" json:"details"`

	// The date and time the responder rule was created. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the responder rule was updated. Format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the ResponderRule.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m ResponderRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ResponderRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingResponderTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetResponderTypeEnumStringValues(), ",")))
	}

	for _, val := range m.SupportedModes {
		if _, ok := GetMappingResponderRuleSupportedModesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SupportedModes: %s. Supported values are: %s.", val, strings.Join(GetResponderRuleSupportedModesEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ResponderRuleSupportedModesEnum Enum with underlying type: string
type ResponderRuleSupportedModesEnum string

// Set of constants representing the allowable values for ResponderRuleSupportedModesEnum
const (
	ResponderRuleSupportedModesAutoaction ResponderRuleSupportedModesEnum = "AUTOACTION"
	ResponderRuleSupportedModesUseraction ResponderRuleSupportedModesEnum = "USERACTION"
)

var mappingResponderRuleSupportedModesEnum = map[string]ResponderRuleSupportedModesEnum{
	"AUTOACTION": ResponderRuleSupportedModesAutoaction,
	"USERACTION": ResponderRuleSupportedModesUseraction,
}

// GetResponderRuleSupportedModesEnumValues Enumerates the set of values for ResponderRuleSupportedModesEnum
func GetResponderRuleSupportedModesEnumValues() []ResponderRuleSupportedModesEnum {
	values := make([]ResponderRuleSupportedModesEnum, 0)
	for _, v := range mappingResponderRuleSupportedModesEnum {
		values = append(values, v)
	}
	return values
}

// GetResponderRuleSupportedModesEnumStringValues Enumerates the set of values in String for ResponderRuleSupportedModesEnum
func GetResponderRuleSupportedModesEnumStringValues() []string {
	return []string{
		"AUTOACTION",
		"USERACTION",
	}
}

// GetMappingResponderRuleSupportedModesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResponderRuleSupportedModesEnum(val string) (ResponderRuleSupportedModesEnum, bool) {
	mappingResponderRuleSupportedModesEnumIgnoreCase := make(map[string]ResponderRuleSupportedModesEnum)
	for k, v := range mappingResponderRuleSupportedModesEnum {
		mappingResponderRuleSupportedModesEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingResponderRuleSupportedModesEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
