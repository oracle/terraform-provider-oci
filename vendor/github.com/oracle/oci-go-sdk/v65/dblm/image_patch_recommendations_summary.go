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

// ImagePatchRecommendationsSummary Recommended patches for current image version.
type ImagePatchRecommendationsSummary struct {

	// Total number of recommended patches.
	Total *int `mandatory:"false" json:"total"`

	// Patch recommendation status for SoftwareImage.
	Status ImagePatchRecommendationsSummaryStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Image version without patch recommendations.
	UpToDateImageVersion *string `mandatory:"false" json:"upToDateImageVersion"`

	// List of the patch recommendations for databases
	ImagePatchRecommendationsDetails []ImagePatchRecommendationsDetails `mandatory:"false" json:"imagePatchRecommendationsDetails"`
}

func (m ImagePatchRecommendationsSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ImagePatchRecommendationsSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingImagePatchRecommendationsSummaryStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetImagePatchRecommendationsSummaryStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ImagePatchRecommendationsSummaryStatusEnum Enum with underlying type: string
type ImagePatchRecommendationsSummaryStatusEnum string

// Set of constants representing the allowable values for ImagePatchRecommendationsSummaryStatusEnum
const (
	ImagePatchRecommendationsSummaryStatusGreen  ImagePatchRecommendationsSummaryStatusEnum = "GREEN"
	ImagePatchRecommendationsSummaryStatusYellow ImagePatchRecommendationsSummaryStatusEnum = "YELLOW"
	ImagePatchRecommendationsSummaryStatusRed    ImagePatchRecommendationsSummaryStatusEnum = "RED"
)

var mappingImagePatchRecommendationsSummaryStatusEnum = map[string]ImagePatchRecommendationsSummaryStatusEnum{
	"GREEN":  ImagePatchRecommendationsSummaryStatusGreen,
	"YELLOW": ImagePatchRecommendationsSummaryStatusYellow,
	"RED":    ImagePatchRecommendationsSummaryStatusRed,
}

var mappingImagePatchRecommendationsSummaryStatusEnumLowerCase = map[string]ImagePatchRecommendationsSummaryStatusEnum{
	"green":  ImagePatchRecommendationsSummaryStatusGreen,
	"yellow": ImagePatchRecommendationsSummaryStatusYellow,
	"red":    ImagePatchRecommendationsSummaryStatusRed,
}

// GetImagePatchRecommendationsSummaryStatusEnumValues Enumerates the set of values for ImagePatchRecommendationsSummaryStatusEnum
func GetImagePatchRecommendationsSummaryStatusEnumValues() []ImagePatchRecommendationsSummaryStatusEnum {
	values := make([]ImagePatchRecommendationsSummaryStatusEnum, 0)
	for _, v := range mappingImagePatchRecommendationsSummaryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetImagePatchRecommendationsSummaryStatusEnumStringValues Enumerates the set of values in String for ImagePatchRecommendationsSummaryStatusEnum
func GetImagePatchRecommendationsSummaryStatusEnumStringValues() []string {
	return []string{
		"GREEN",
		"YELLOW",
		"RED",
	}
}

// GetMappingImagePatchRecommendationsSummaryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingImagePatchRecommendationsSummaryStatusEnum(val string) (ImagePatchRecommendationsSummaryStatusEnum, bool) {
	enum, ok := mappingImagePatchRecommendationsSummaryStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
