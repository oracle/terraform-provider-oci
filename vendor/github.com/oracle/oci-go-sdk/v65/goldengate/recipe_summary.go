// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RecipeSummary The list of recipe details to create pipelines.
type RecipeSummary struct {

	// The type of the recipe
	RecipeType RecipeTypeEnum `mandatory:"true" json:"recipeType"`

	// An object's Display Name.
	Name *string `mandatory:"true" json:"name"`

	// An object's Display Name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Array of supported technology types for this recipe.
	SupportedSourceTechnologyTypes []TechnologyTypeEnum `mandatory:"true" json:"supportedSourceTechnologyTypes"`

	// Array of supported technology types for this recipe.
	SupportedTargetTechnologyTypes []TechnologyTypeEnum `mandatory:"true" json:"supportedTargetTechnologyTypes"`

	// Metadata about this specific object.
	Description *string `mandatory:"false" json:"description"`
}

func (m RecipeSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RecipeSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRecipeTypeEnum(string(m.RecipeType)); !ok && m.RecipeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RecipeType: %s. Supported values are: %s.", m.RecipeType, strings.Join(GetRecipeTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
