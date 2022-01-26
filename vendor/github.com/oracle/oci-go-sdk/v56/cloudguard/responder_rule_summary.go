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

// ResponderRuleSummarySupportedModesEnum Enum with underlying type: string
type ResponderRuleSummarySupportedModesEnum string

// Set of constants representing the allowable values for ResponderRuleSummarySupportedModesEnum
const (
	ResponderRuleSummarySupportedModesAutoaction ResponderRuleSummarySupportedModesEnum = "AUTOACTION"
	ResponderRuleSummarySupportedModesUseraction ResponderRuleSummarySupportedModesEnum = "USERACTION"
)

var mappingResponderRuleSummarySupportedModes = map[string]ResponderRuleSummarySupportedModesEnum{
	"AUTOACTION": ResponderRuleSummarySupportedModesAutoaction,
	"USERACTION": ResponderRuleSummarySupportedModesUseraction,
}

// GetResponderRuleSummarySupportedModesEnumValues Enumerates the set of values for ResponderRuleSummarySupportedModesEnum
func GetResponderRuleSummarySupportedModesEnumValues() []ResponderRuleSummarySupportedModesEnum {
	values := make([]ResponderRuleSummarySupportedModesEnum, 0)
	for _, v := range mappingResponderRuleSummarySupportedModes {
		values = append(values, v)
	}
	return values
}
