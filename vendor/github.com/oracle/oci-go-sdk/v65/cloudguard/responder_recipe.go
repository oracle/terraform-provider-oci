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

// ResponderRecipe A ResponderRecipe resource contains a specific instance of one of
// the supported detector types (for example, activity, configuration,
// or threat).
// A ResponderRecipe resource:
// * Is effectively a copy of a Responder resource in which users can make
// very limited changes if it’s Oracle-managed, and more changes if it’s user-managed.
// * Can also be created by cloning an existing ResponderRecipe resource, either
// user-managed or Oracle-managed.
// * Is visible on Cloud Guard’s Responder Recipes page.
// * Is located in a specific OCI compartment.
// * Can be modified by users, programmatically or through the UI.
// * Changes that can be made here apply globally, to resources in all OCI compartments
// mapped to a target that attaches the responder recipe, but are overridden by
// any changes made in the corresponding TargetResponderRecipe resource (effectively
// created when the responder recipe is attached to the target).
type ResponderRecipe struct {

	// Unique identifier for the responder recip
	Id *string `mandatory:"true" json:"id"`

	// Compartment OCID
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Responder recipe display name
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Responder recipe description
	Description *string `mandatory:"false" json:"description"`

	// Owner of responder recipe
	Owner OwnerTypeEnum `mandatory:"false" json:"owner,omitempty"`

	// List of responder rules associated with the recipe
	ResponderRules []ResponderRecipeResponderRule `mandatory:"false" json:"responderRules"`

	// List of currently enabled responder rules for the responder type, for recipe after applying defaults
	EffectiveResponderRules []ResponderRecipeResponderRule `mandatory:"false" json:"effectiveResponderRules"`

	// The unique identifier of the source responder recipe
	SourceResponderRecipeId *string `mandatory:"false" json:"sourceResponderRecipeId"`

	// The date and time the responder recipe was created. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the responder recipe was last updated. Format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current lifecycle state of the example
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	// Avoid entering confidential information.
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// System tags can be viewed by users, but can only be created by the system.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m ResponderRecipe) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ResponderRecipe) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOwnerTypeEnum(string(m.Owner)); !ok && m.Owner != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Owner: %s. Supported values are: %s.", m.Owner, strings.Join(GetOwnerTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
