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

// TargetDetectorRecipe Target Detector recipe
type TargetDetectorRecipe struct {

	// Ocid for detector recipe
	Id *string `mandatory:"true" json:"id"`

	// DisplayName of detector recipe
	DisplayName *string `mandatory:"true" json:"displayName"`

	// compartmentId of detector recipe
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Unique identifier for Detector Recipe of which this is an extension
	DetectorRecipeId *string `mandatory:"true" json:"detectorRecipeId"`

	// Owner of detector recipe
	Owner OwnerTypeEnum `mandatory:"true" json:"owner"`

	// Type of detector
	Detector DetectorEnumEnum `mandatory:"true" json:"detector"`

	// Detector recipe description
	Description *string `mandatory:"false" json:"description"`

	// List of detector rules for the detector type for recipe - user input
	DetectorRules []TargetDetectorRecipeDetectorRule `mandatory:"false" json:"detectorRules"`

	// List of effective detector rules for the detector type for recipe after applying defaults
	EffectiveDetectorRules []TargetDetectorRecipeDetectorRule `mandatory:"false" json:"effectiveDetectorRules"`

	// The date and time the target detector recipe was created. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the target detector recipe was updated. Format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the resource.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m TargetDetectorRecipe) String() string {
	return common.PointerString(m)
}
