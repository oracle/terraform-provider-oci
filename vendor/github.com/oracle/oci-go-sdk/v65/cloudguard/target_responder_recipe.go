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

// TargetResponderRecipe A TargetResponderRecipe resource contains a specific instance of one of the
// supported detector types (for example, activity, configuration, or threat)
// in which some settings can be modified specifically for a single target.
// A TargetResponderRecipe resource:
// * Is effectively a copy of a ResponderRecipe resource in which users can make
// very limited changes if it’s Oracle-managed, and more changes if it’s user-managed.
// * Is visible on the Cloud Guard Targets, Target Details page.
// * Is located in a specific OCI compartment.
// * Can be modified by users, programmatically or through the UI.
// * Changes that can be made here override any settings in the corresponding
// ResponderRecipe, of which the TargetResponderRecipe resource is effectively a copy
// of the ResponderRecipe resource (effectively created when the detector recipe
// is attached to the target).
type TargetResponderRecipe struct {

	// Unique identifier of target responder recipe that can't be changed after creation
	Id *string `mandatory:"true" json:"id"`

	// Unique identifier for the Oracle-managed responder recipe from which this recipe was cloned
	ResponderRecipeId *string `mandatory:"true" json:"responderRecipeId"`

	// Compartment OCID
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Target responder recipe display name
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Target responder description
	Description *string `mandatory:"true" json:"description"`

	// Owner of target responder recipe
	Owner OwnerTypeEnum `mandatory:"true" json:"owner"`

	// The date and time the target responder recipe rule was created. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the target responder recipe rule was last updated. Format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// List of responder rules associated with the recipe - user input
	ResponderRules []TargetResponderRecipeResponderRule `mandatory:"false" json:"responderRules"`

	// List of currently enabled responder rules for the responder type for recipe after applying defaults
	EffectiveResponderRules []TargetResponderRecipeResponderRule `mandatory:"false" json:"effectiveResponderRules"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`
}

func (m TargetResponderRecipe) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TargetResponderRecipe) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOwnerTypeEnum(string(m.Owner)); !ok && m.Owner != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Owner: %s. Supported values are: %s.", m.Owner, strings.Join(GetOwnerTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
