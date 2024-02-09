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

// ResponderRecipeResponderRule A ResponderRecipeRule resource contains a specific instance of a
// single responder rule.
// A ResponderRecipeRule resource:
// * Is effectively a copy of a ResponderRule resource in which users can
// make certain changes if it’s Oracle-managed, and other changes if it’s user-managed.
// * Can also be created by cloning an existing ResponderRecipe resource, either
// user-managed or Oracle-managed; cloning the ResponderRecipe resource also clones
// its associated ResponderRule resources as ResponderRecipeRule resources.
// * Is visible on the Cloud Guard Responder Recipes, Responder Details page.
// * Is effectively located in a specific OCI compartment, through the ResponderRecipe
// resource to which it belongs.
// * Can be modified by users, programmatically or through the UI.
// * Changes that can be made here apply globally, to all resources in OCI compartments
// mapped to a target that attaches the associated responder recipe (in a
// TargetResponderRecipe resource), but are overridden by changes made in the
// corresponding TargetResponderRecipe resource (which is effectively a copy of the
// ResponderRecipe resource).
// type: object
type ResponderRecipeResponderRule struct {

	// Unique identifier for the responder rule
	ResponderRuleId *string `mandatory:"true" json:"responderRuleId"`

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
	SupportedModes []ResponderRecipeResponderRuleSupportedModesEnum `mandatory:"false" json:"supportedModes,omitempty"`

	Details *ResponderRuleDetails `mandatory:"false" json:"details"`

	// The date and time the responder recipe rule was created. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the responder recipe rule was last updated. Format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current lifecycle state of the responder rule
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`
}

func (m ResponderRecipeResponderRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ResponderRecipeResponderRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingResponderTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetResponderTypeEnumStringValues(), ",")))
	}
	for _, val := range m.SupportedModes {
		if _, ok := GetMappingResponderRecipeResponderRuleSupportedModesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SupportedModes: %s. Supported values are: %s.", val, strings.Join(GetResponderRecipeResponderRuleSupportedModesEnumStringValues(), ",")))
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

// ResponderRecipeResponderRuleSupportedModesEnum Enum with underlying type: string
type ResponderRecipeResponderRuleSupportedModesEnum string

// Set of constants representing the allowable values for ResponderRecipeResponderRuleSupportedModesEnum
const (
	ResponderRecipeResponderRuleSupportedModesAutoaction ResponderRecipeResponderRuleSupportedModesEnum = "AUTOACTION"
	ResponderRecipeResponderRuleSupportedModesUseraction ResponderRecipeResponderRuleSupportedModesEnum = "USERACTION"
)

var mappingResponderRecipeResponderRuleSupportedModesEnum = map[string]ResponderRecipeResponderRuleSupportedModesEnum{
	"AUTOACTION": ResponderRecipeResponderRuleSupportedModesAutoaction,
	"USERACTION": ResponderRecipeResponderRuleSupportedModesUseraction,
}

var mappingResponderRecipeResponderRuleSupportedModesEnumLowerCase = map[string]ResponderRecipeResponderRuleSupportedModesEnum{
	"autoaction": ResponderRecipeResponderRuleSupportedModesAutoaction,
	"useraction": ResponderRecipeResponderRuleSupportedModesUseraction,
}

// GetResponderRecipeResponderRuleSupportedModesEnumValues Enumerates the set of values for ResponderRecipeResponderRuleSupportedModesEnum
func GetResponderRecipeResponderRuleSupportedModesEnumValues() []ResponderRecipeResponderRuleSupportedModesEnum {
	values := make([]ResponderRecipeResponderRuleSupportedModesEnum, 0)
	for _, v := range mappingResponderRecipeResponderRuleSupportedModesEnum {
		values = append(values, v)
	}
	return values
}

// GetResponderRecipeResponderRuleSupportedModesEnumStringValues Enumerates the set of values in String for ResponderRecipeResponderRuleSupportedModesEnum
func GetResponderRecipeResponderRuleSupportedModesEnumStringValues() []string {
	return []string{
		"AUTOACTION",
		"USERACTION",
	}
}

// GetMappingResponderRecipeResponderRuleSupportedModesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResponderRecipeResponderRuleSupportedModesEnum(val string) (ResponderRecipeResponderRuleSupportedModesEnum, bool) {
	enum, ok := mappingResponderRecipeResponderRuleSupportedModesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
