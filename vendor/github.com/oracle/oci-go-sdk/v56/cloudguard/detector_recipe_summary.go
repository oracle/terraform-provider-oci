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

// DetectorRecipeSummary Summary of Detector recipe
type DetectorRecipeSummary struct {

	// Ocid for detector recipe
	Id *string `mandatory:"true" json:"id"`

	// DisplayName of detector recipe
	DisplayName *string `mandatory:"true" json:"displayName"`

	// compartmentId of detector recipe
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Owner of detector recipe
	Owner OwnerTypeEnum `mandatory:"true" json:"owner"`

	// Type of detector
	Detector DetectorEnumEnum `mandatory:"true" json:"detector"`

	// Detector recipe description
	Description *string `mandatory:"false" json:"description"`

	// Recipe Ocid of the Source Recipe to be cloned
	SourceDetectorRecipeId *string `mandatory:"false" json:"sourceDetectorRecipeId"`

	// List of detetor rules for the detector type
	DetectorRules []DetectorRecipeDetectorRule `mandatory:"false" json:"detectorRules"`

	// The date and time the detector recipe was created. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the detector recipe was updated. Format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the resource.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

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

func (m DetectorRecipeSummary) String() string {
	return common.PointerString(m)
}
