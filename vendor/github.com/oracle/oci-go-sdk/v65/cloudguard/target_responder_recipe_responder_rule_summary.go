// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TargetResponderRecipeResponderRuleSummary Summary information for a target responder recipe responder rule.
type TargetResponderRecipeResponderRuleSummary struct {

	// Unique identifier for the responder rule
	Id *string `mandatory:"true" json:"id"`

	// Compartment OCID
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Responder rule display name
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Responder rule description
	Description *string `mandatory:"false" json:"description"`

	// Type of responder
	Type ResponderTypeEnum `mandatory:"false" json:"type,omitempty"`

	// List of policies
	Policies []string `mandatory:"false" json:"policies"`

	// Supported execution modes for the responder rule
	SupportedModes []TargetResponderRecipeResponderRuleSummarySupportedModesEnum `mandatory:"false" json:"supportedModes,omitempty"`

	Details *ResponderRuleDetails `mandatory:"false" json:"details"`

	// The date and time the target responder recipe rule was created. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the target responder recipe rule was last updated. Format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current lifecycle state of the responder rule
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m TargetResponderRecipeResponderRuleSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TargetResponderRecipeResponderRuleSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingResponderTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetResponderTypeEnumStringValues(), ",")))
	}
	for _, val := range m.SupportedModes {
		if _, ok := GetMappingTargetResponderRecipeResponderRuleSummarySupportedModesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SupportedModes: %s. Supported values are: %s.", val, strings.Join(GetTargetResponderRecipeResponderRuleSummarySupportedModesEnumStringValues(), ",")))
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

// TargetResponderRecipeResponderRuleSummarySupportedModesEnum Enum with underlying type: string
type TargetResponderRecipeResponderRuleSummarySupportedModesEnum string

// Set of constants representing the allowable values for TargetResponderRecipeResponderRuleSummarySupportedModesEnum
const (
	TargetResponderRecipeResponderRuleSummarySupportedModesAutoaction TargetResponderRecipeResponderRuleSummarySupportedModesEnum = "AUTOACTION"
	TargetResponderRecipeResponderRuleSummarySupportedModesUseraction TargetResponderRecipeResponderRuleSummarySupportedModesEnum = "USERACTION"
)

var mappingTargetResponderRecipeResponderRuleSummarySupportedModesEnum = map[string]TargetResponderRecipeResponderRuleSummarySupportedModesEnum{
	"AUTOACTION": TargetResponderRecipeResponderRuleSummarySupportedModesAutoaction,
	"USERACTION": TargetResponderRecipeResponderRuleSummarySupportedModesUseraction,
}

var mappingTargetResponderRecipeResponderRuleSummarySupportedModesEnumLowerCase = map[string]TargetResponderRecipeResponderRuleSummarySupportedModesEnum{
	"autoaction": TargetResponderRecipeResponderRuleSummarySupportedModesAutoaction,
	"useraction": TargetResponderRecipeResponderRuleSummarySupportedModesUseraction,
}

// GetTargetResponderRecipeResponderRuleSummarySupportedModesEnumValues Enumerates the set of values for TargetResponderRecipeResponderRuleSummarySupportedModesEnum
func GetTargetResponderRecipeResponderRuleSummarySupportedModesEnumValues() []TargetResponderRecipeResponderRuleSummarySupportedModesEnum {
	values := make([]TargetResponderRecipeResponderRuleSummarySupportedModesEnum, 0)
	for _, v := range mappingTargetResponderRecipeResponderRuleSummarySupportedModesEnum {
		values = append(values, v)
	}
	return values
}

// GetTargetResponderRecipeResponderRuleSummarySupportedModesEnumStringValues Enumerates the set of values in String for TargetResponderRecipeResponderRuleSummarySupportedModesEnum
func GetTargetResponderRecipeResponderRuleSummarySupportedModesEnumStringValues() []string {
	return []string{
		"AUTOACTION",
		"USERACTION",
	}
}

// GetMappingTargetResponderRecipeResponderRuleSummarySupportedModesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTargetResponderRecipeResponderRuleSummarySupportedModesEnum(val string) (TargetResponderRecipeResponderRuleSummarySupportedModesEnum, bool) {
	enum, ok := mappingTargetResponderRecipeResponderRuleSummarySupportedModesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
