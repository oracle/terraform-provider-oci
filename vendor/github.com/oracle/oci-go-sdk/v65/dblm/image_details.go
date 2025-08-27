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

// ImageDetails Image details containing the subscribed image, its status, version, owner and time of creation.
type ImageDetails struct {

	// Image identifier.
	ImageId *string `mandatory:"false" json:"imageId"`

	// Subscribed image.
	SubscribedImage *string `mandatory:"false" json:"subscribedImage"`

	// Name of the image version marked as current of the image.
	CurrentVersion *string `mandatory:"false" json:"currentVersion"`

	// Image status.
	ImageStatus ImageDetailsImageStatusEnum `mandatory:"false" json:"imageStatus,omitempty"`

	// Release version of the image.
	ImageVersion *string `mandatory:"false" json:"imageVersion"`

	// Owner of the image.
	ImageOwner *string `mandatory:"false" json:"imageOwner"`

	// Name of the person who created the image.
	CreatedBy *string `mandatory:"false" json:"createdBy"`

	// Date when the image was created.
	TimeImageCreation *common.SDKTime `mandatory:"false" json:"timeImageCreation"`

	// An image version name, that is up to date and has no recommendations.
	UpToDateImageVersion *string `mandatory:"false" json:"upToDateImageVersion"`
}

func (m ImageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ImageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingImageDetailsImageStatusEnum(string(m.ImageStatus)); !ok && m.ImageStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ImageStatus: %s. Supported values are: %s.", m.ImageStatus, strings.Join(GetImageDetailsImageStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ImageDetailsImageStatusEnum Enum with underlying type: string
type ImageDetailsImageStatusEnum string

// Set of constants representing the allowable values for ImageDetailsImageStatusEnum
const (
	ImageDetailsImageStatusGreen  ImageDetailsImageStatusEnum = "GREEN"
	ImageDetailsImageStatusYellow ImageDetailsImageStatusEnum = "YELLOW"
	ImageDetailsImageStatusRed    ImageDetailsImageStatusEnum = "RED"
)

var mappingImageDetailsImageStatusEnum = map[string]ImageDetailsImageStatusEnum{
	"GREEN":  ImageDetailsImageStatusGreen,
	"YELLOW": ImageDetailsImageStatusYellow,
	"RED":    ImageDetailsImageStatusRed,
}

var mappingImageDetailsImageStatusEnumLowerCase = map[string]ImageDetailsImageStatusEnum{
	"green":  ImageDetailsImageStatusGreen,
	"yellow": ImageDetailsImageStatusYellow,
	"red":    ImageDetailsImageStatusRed,
}

// GetImageDetailsImageStatusEnumValues Enumerates the set of values for ImageDetailsImageStatusEnum
func GetImageDetailsImageStatusEnumValues() []ImageDetailsImageStatusEnum {
	values := make([]ImageDetailsImageStatusEnum, 0)
	for _, v := range mappingImageDetailsImageStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetImageDetailsImageStatusEnumStringValues Enumerates the set of values in String for ImageDetailsImageStatusEnum
func GetImageDetailsImageStatusEnumStringValues() []string {
	return []string{
		"GREEN",
		"YELLOW",
		"RED",
	}
}

// GetMappingImageDetailsImageStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingImageDetailsImageStatusEnum(val string) (ImageDetailsImageStatusEnum, bool) {
	enum, ok := mappingImageDetailsImageStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
