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

// Target Description of Target.
type Target struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Compartment Identifier where the resource is created
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// possible type of targets
	TargetResourceType TargetResourceTypeEnum `mandatory:"true" json:"targetResourceType"`

	// Resource ID which the target uses to monitor
	TargetResourceId *string `mandatory:"true" json:"targetResourceId"`

	// Total number of recipes attached to target
	RecipeCount *int `mandatory:"true" json:"recipeCount"`

	// Target Identifier, can be renamed
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The target description.
	Description *string `mandatory:"false" json:"description"`

	// List of detector recipes associated with target
	TargetDetectorRecipes []TargetDetectorRecipe `mandatory:"false" json:"targetDetectorRecipes"`

	// List of responder recipes associated with target
	TargetResponderRecipes []TargetResponderRecipe `mandatory:"false" json:"targetResponderRecipes"`

	// List of inherited compartments
	InheritedByCompartments []string `mandatory:"false" json:"inheritedByCompartments"`

	// The date and time the target was created. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the target was updated. Format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the Target.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecyleDetails *string `mandatory:"false" json:"lifecyleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
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

func (m Target) String() string {
	return common.PointerString(m)
}
