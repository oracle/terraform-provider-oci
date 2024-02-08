// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.cloud.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.cloud.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ResponderRuleSummary Summary of the ResponderRule.
type ResponderRuleSummary struct {

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
	SupportedModes []ResponderRuleSummarySupportedModesEnum `mandatory:"false" json:"supportedModes,omitempty"`

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

func (m ResponderRuleSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ResponderRuleSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingResponderTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetResponderTypeEnumStringValues(), ",")))
	}

	for _, val := range m.SupportedModes {
		if _, ok := GetMappingResponderRuleSummarySupportedModesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SupportedModes: %s. Supported values are: %s.", val, strings.Join(GetResponderRuleSummarySupportedModesEnumStringValues(), ",")))
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

// ResponderRuleSummarySupportedModesEnum Enum with underlying type: string
type ResponderRuleSummarySupportedModesEnum string

// Set of constants representing the allowable values for ResponderRuleSummarySupportedModesEnum
const (
	ResponderRuleSummarySupportedModesAutoaction ResponderRuleSummarySupportedModesEnum = "AUTOACTION"
	ResponderRuleSummarySupportedModesUseraction ResponderRuleSummarySupportedModesEnum = "USERACTION"
)

var mappingResponderRuleSummarySupportedModesEnum = map[string]ResponderRuleSummarySupportedModesEnum{
	"AUTOACTION": ResponderRuleSummarySupportedModesAutoaction,
	"USERACTION": ResponderRuleSummarySupportedModesUseraction,
}

var mappingResponderRuleSummarySupportedModesEnumLowerCase = map[string]ResponderRuleSummarySupportedModesEnum{
	"autoaction": ResponderRuleSummarySupportedModesAutoaction,
	"useraction": ResponderRuleSummarySupportedModesUseraction,
}

// GetResponderRuleSummarySupportedModesEnumValues Enumerates the set of values for ResponderRuleSummarySupportedModesEnum
func GetResponderRuleSummarySupportedModesEnumValues() []ResponderRuleSummarySupportedModesEnum {
	values := make([]ResponderRuleSummarySupportedModesEnum, 0)
	for _, v := range mappingResponderRuleSummarySupportedModesEnum {
		values = append(values, v)
	}
	return values
}

// GetResponderRuleSummarySupportedModesEnumStringValues Enumerates the set of values in String for ResponderRuleSummarySupportedModesEnum
func GetResponderRuleSummarySupportedModesEnumStringValues() []string {
	return []string{
		"AUTOACTION",
		"USERACTION",
	}
}

// GetMappingResponderRuleSummarySupportedModesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResponderRuleSummarySupportedModesEnum(val string) (ResponderRuleSummarySupportedModesEnum, bool) {
	enum, ok := mappingResponderRuleSummarySupportedModesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
