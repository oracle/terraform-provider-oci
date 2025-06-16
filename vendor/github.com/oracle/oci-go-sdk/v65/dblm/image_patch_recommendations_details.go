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

// ImagePatchRecommendationsDetails Summary of patch recommendations for image.
type ImagePatchRecommendationsDetails struct {

	// Id for the patch recommendation.
	PatchId *int `mandatory:"false" json:"patchId"`

	// Name for the patch recommendation.
	PatchName *string `mandatory:"false" json:"patchName"`

	// Description of the patch recommendation.
	Description *string `mandatory:"false" json:"description"`

	// Shows if patch is recommended or is an additional patch from an existing database.
	Category ImagePatchRecommendationsDetailsCategoryEnum `mandatory:"false" json:"category,omitempty"`

	// Database with patch.
	DatabasesWithPatchCount *int `mandatory:"false" json:"databasesWithPatchCount"`
}

func (m ImagePatchRecommendationsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ImagePatchRecommendationsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingImagePatchRecommendationsDetailsCategoryEnum(string(m.Category)); !ok && m.Category != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Category: %s. Supported values are: %s.", m.Category, strings.Join(GetImagePatchRecommendationsDetailsCategoryEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ImagePatchRecommendationsDetailsCategoryEnum Enum with underlying type: string
type ImagePatchRecommendationsDetailsCategoryEnum string

// Set of constants representing the allowable values for ImagePatchRecommendationsDetailsCategoryEnum
const (
	ImagePatchRecommendationsDetailsCategoryRecommended     ImagePatchRecommendationsDetailsCategoryEnum = "RECOMMENDED"
	ImagePatchRecommendationsDetailsCategoryAdditionalPatch ImagePatchRecommendationsDetailsCategoryEnum = "ADDITIONAL_PATCH"
)

var mappingImagePatchRecommendationsDetailsCategoryEnum = map[string]ImagePatchRecommendationsDetailsCategoryEnum{
	"RECOMMENDED":      ImagePatchRecommendationsDetailsCategoryRecommended,
	"ADDITIONAL_PATCH": ImagePatchRecommendationsDetailsCategoryAdditionalPatch,
}

var mappingImagePatchRecommendationsDetailsCategoryEnumLowerCase = map[string]ImagePatchRecommendationsDetailsCategoryEnum{
	"recommended":      ImagePatchRecommendationsDetailsCategoryRecommended,
	"additional_patch": ImagePatchRecommendationsDetailsCategoryAdditionalPatch,
}

// GetImagePatchRecommendationsDetailsCategoryEnumValues Enumerates the set of values for ImagePatchRecommendationsDetailsCategoryEnum
func GetImagePatchRecommendationsDetailsCategoryEnumValues() []ImagePatchRecommendationsDetailsCategoryEnum {
	values := make([]ImagePatchRecommendationsDetailsCategoryEnum, 0)
	for _, v := range mappingImagePatchRecommendationsDetailsCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetImagePatchRecommendationsDetailsCategoryEnumStringValues Enumerates the set of values in String for ImagePatchRecommendationsDetailsCategoryEnum
func GetImagePatchRecommendationsDetailsCategoryEnumStringValues() []string {
	return []string{
		"RECOMMENDED",
		"ADDITIONAL_PATCH",
	}
}

// GetMappingImagePatchRecommendationsDetailsCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingImagePatchRecommendationsDetailsCategoryEnum(val string) (ImagePatchRecommendationsDetailsCategoryEnum, bool) {
	enum, ok := mappingImagePatchRecommendationsDetailsCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
