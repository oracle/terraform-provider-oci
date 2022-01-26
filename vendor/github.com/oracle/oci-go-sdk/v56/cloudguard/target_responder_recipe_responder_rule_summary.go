// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// TargetResponderRecipeResponderRuleSummary Summary of ResponderRule.
type TargetResponderRecipeResponderRuleSummary struct {

	// Identifier for ResponderRule.
	Id *string `mandatory:"true" json:"id"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// ResponderRule Display Name
	DisplayName *string `mandatory:"false" json:"displayName"`

	// ResponderRule Description
	Description *string `mandatory:"false" json:"description"`

	// Type of Responder
	Type ResponderTypeEnum `mandatory:"false" json:"type,omitempty"`

	// List of Policy
	Policies []string `mandatory:"false" json:"policies"`

	// Supported Execution Modes
	SupportedModes []TargetResponderRecipeResponderRuleSummarySupportedModesEnum `mandatory:"false" json:"supportedModes,omitempty"`

	Details *ResponderRuleDetails `mandatory:"false" json:"details"`

	// The date and time the target responder recipe rule was created. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the target responder recipe rule was updated. Format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the ResponderRule.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m TargetResponderRecipeResponderRuleSummary) String() string {
	return common.PointerString(m)
}

// TargetResponderRecipeResponderRuleSummarySupportedModesEnum Enum with underlying type: string
type TargetResponderRecipeResponderRuleSummarySupportedModesEnum string

// Set of constants representing the allowable values for TargetResponderRecipeResponderRuleSummarySupportedModesEnum
const (
	TargetResponderRecipeResponderRuleSummarySupportedModesAutoaction TargetResponderRecipeResponderRuleSummarySupportedModesEnum = "AUTOACTION"
	TargetResponderRecipeResponderRuleSummarySupportedModesUseraction TargetResponderRecipeResponderRuleSummarySupportedModesEnum = "USERACTION"
)

var mappingTargetResponderRecipeResponderRuleSummarySupportedModes = map[string]TargetResponderRecipeResponderRuleSummarySupportedModesEnum{
	"AUTOACTION": TargetResponderRecipeResponderRuleSummarySupportedModesAutoaction,
	"USERACTION": TargetResponderRecipeResponderRuleSummarySupportedModesUseraction,
}

// GetTargetResponderRecipeResponderRuleSummarySupportedModesEnumValues Enumerates the set of values for TargetResponderRecipeResponderRuleSummarySupportedModesEnum
func GetTargetResponderRecipeResponderRuleSummarySupportedModesEnumValues() []TargetResponderRecipeResponderRuleSummarySupportedModesEnum {
	values := make([]TargetResponderRecipeResponderRuleSummarySupportedModesEnum, 0)
	for _, v := range mappingTargetResponderRecipeResponderRuleSummarySupportedModes {
		values = append(values, v)
	}
	return values
}
