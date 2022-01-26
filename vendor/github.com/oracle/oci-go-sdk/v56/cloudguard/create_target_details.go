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

// CreateTargetDetails The information about new Target.
type CreateTargetDetails struct {

	// DetectorTemplate Identifier
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier where the resource is created
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// possible type of targets(compartment/HCMCloud/ERPCloud)
	TargetResourceType TargetResourceTypeEnum `mandatory:"true" json:"targetResourceType"`

	// Resource ID which the target uses to monitor
	TargetResourceId *string `mandatory:"true" json:"targetResourceId"`

	// The target description.
	Description *string `mandatory:"false" json:"description"`

	// List of detector recipes to associate with target
	TargetDetectorRecipes []CreateTargetDetectorRecipeDetails `mandatory:"false" json:"targetDetectorRecipes"`

	// List of responder recipes to associate with target
	TargetResponderRecipes []CreateTargetResponderRecipeDetails `mandatory:"false" json:"targetResponderRecipes"`

	// The current state of the DetectorRule.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateTargetDetails) String() string {
	return common.PointerString(m)
}
