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

// CreateDetectorRecipeDetails Create of Detector recipe.
type CreateDetectorRecipeDetails struct {

	// DetectorRecipe Display Name
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The id of the source detector recipe.
	SourceDetectorRecipeId *string `mandatory:"true" json:"sourceDetectorRecipeId"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// DetectorRecipe Description
	Description *string `mandatory:"false" json:"description"`

	// Detector Rules to override from source detector recipe
	DetectorRules []UpdateDetectorRecipeDetectorRule `mandatory:"false" json:"detectorRules"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateDetectorRecipeDetails) String() string {
	return common.PointerString(m)
}
