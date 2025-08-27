// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Dblm API
//
// A description of the Dblm API
//

package dblm

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AdditionalPatches Summary of patch recommendations for image.
type AdditionalPatches struct {

	// Id for the patch recommendation.
	PatchId *int `mandatory:"false" json:"patchId"`

	// Name for the patch recommendation.
	PatchName *string `mandatory:"false" json:"patchName"`

	// Description of the patch recommendation.
	Description *string `mandatory:"false" json:"description"`

	// Shows if patch is recommended or is an additional patch from an existing database.
	Category AdditionalPatchesCategoryEnum `mandatory:"false" json:"category,omitempty"`
}

func (m AdditionalPatches) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AdditionalPatches) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAdditionalPatchesCategoryEnum(string(m.Category)); !ok && m.Category != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Category: %s. Supported values are: %s.", m.Category, strings.Join(GetAdditionalPatchesCategoryEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AdditionalPatchesCategoryEnum Enum with underlying type: string
type AdditionalPatchesCategoryEnum string

// Set of constants representing the allowable values for AdditionalPatchesCategoryEnum
const (
	AdditionalPatchesCategoryRecommended     AdditionalPatchesCategoryEnum = "RECOMMENDED"
	AdditionalPatchesCategoryAdditionalPatch AdditionalPatchesCategoryEnum = "ADDITIONAL_PATCH"
)

var mappingAdditionalPatchesCategoryEnum = map[string]AdditionalPatchesCategoryEnum{
	"RECOMMENDED":      AdditionalPatchesCategoryRecommended,
	"ADDITIONAL_PATCH": AdditionalPatchesCategoryAdditionalPatch,
}

var mappingAdditionalPatchesCategoryEnumLowerCase = map[string]AdditionalPatchesCategoryEnum{
	"recommended":      AdditionalPatchesCategoryRecommended,
	"additional_patch": AdditionalPatchesCategoryAdditionalPatch,
}

// GetAdditionalPatchesCategoryEnumValues Enumerates the set of values for AdditionalPatchesCategoryEnum
func GetAdditionalPatchesCategoryEnumValues() []AdditionalPatchesCategoryEnum {
	values := make([]AdditionalPatchesCategoryEnum, 0)
	for _, v := range mappingAdditionalPatchesCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetAdditionalPatchesCategoryEnumStringValues Enumerates the set of values in String for AdditionalPatchesCategoryEnum
func GetAdditionalPatchesCategoryEnumStringValues() []string {
	return []string{
		"RECOMMENDED",
		"ADDITIONAL_PATCH",
	}
}

// GetMappingAdditionalPatchesCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAdditionalPatchesCategoryEnum(val string) (AdditionalPatchesCategoryEnum, bool) {
	enum, ok := mappingAdditionalPatchesCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
