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

// UpdateTargetDetails The information to be updated.
type UpdateTargetDetails struct {

	// DetectorTemplate Identifier
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The current state of the Target.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The details of target detector recipes to be updated.
	TargetDetectorRecipes []UpdateTargetDetectorRecipe `mandatory:"false" json:"targetDetectorRecipes"`

	// The details of target responder recipes to be updated.
	TargetResponderRecipes []UpdateTargetResponderRecipe `mandatory:"false" json:"targetResponderRecipes"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateTargetDetails) String() string {
	return common.PointerString(m)
}
