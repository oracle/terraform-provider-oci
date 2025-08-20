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

// ImageVersionsDetails Image Version details.
type ImageVersionsDetails struct {

	// Software image version key.
	ImageVersionKey *int `mandatory:"false" json:"imageVersionKey"`

	// Software image version name.
	Name *string `mandatory:"false" json:"name"`

	// Creation status of SoftwareImageVersion
	CreationStatus ImageVersionsDetailsCreationStatusEnum `mandatory:"false" json:"creationStatus,omitempty"`

	// Software image version recommended patches.
	RecommendedPatches *int `mandatory:"false" json:"recommendedPatches"`

	// Software image version database release.
	DatabaseRelease *string `mandatory:"false" json:"databaseRelease"`

	// Software image version created date.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Software image version status (CURRENT, ACTIVE, DRAFT, RESTRICTED, OBSOLETE, DELETED, UNAVAILABLE).
	Status ImageVersionsDetailsStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Software image version position.
	Position *int `mandatory:"false" json:"position"`

	// Error message if creation fails for this software image version.
	Message *string `mandatory:"false" json:"message"`

	// Existing software image version key for reference
	ReferenceImageVersionKey *int64 `mandatory:"false" json:"referenceImageVersionKey"`

	// List of patchIds used to create this Software Image Version.
	PatchIds []int `mandatory:"false" json:"patchIds"`
}

func (m ImageVersionsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ImageVersionsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingImageVersionsDetailsCreationStatusEnum(string(m.CreationStatus)); !ok && m.CreationStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CreationStatus: %s. Supported values are: %s.", m.CreationStatus, strings.Join(GetImageVersionsDetailsCreationStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingImageVersionsDetailsStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetImageVersionsDetailsStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ImageVersionsDetailsCreationStatusEnum Enum with underlying type: string
type ImageVersionsDetailsCreationStatusEnum string

// Set of constants representing the allowable values for ImageVersionsDetailsCreationStatusEnum
const (
	ImageVersionsDetailsCreationStatusReady              ImageVersionsDetailsCreationStatusEnum = "READY"
	ImageVersionsDetailsCreationStatusErrorInCreation    ImageVersionsDetailsCreationStatusEnum = "ERROR_IN_CREATION"
	ImageVersionsDetailsCreationStatusCreationInProgress ImageVersionsDetailsCreationStatusEnum = "CREATION_IN_PROGRESS"
	ImageVersionsDetailsCreationStatusScheduled          ImageVersionsDetailsCreationStatusEnum = "SCHEDULED"
)

var mappingImageVersionsDetailsCreationStatusEnum = map[string]ImageVersionsDetailsCreationStatusEnum{
	"READY":                ImageVersionsDetailsCreationStatusReady,
	"ERROR_IN_CREATION":    ImageVersionsDetailsCreationStatusErrorInCreation,
	"CREATION_IN_PROGRESS": ImageVersionsDetailsCreationStatusCreationInProgress,
	"SCHEDULED":            ImageVersionsDetailsCreationStatusScheduled,
}

var mappingImageVersionsDetailsCreationStatusEnumLowerCase = map[string]ImageVersionsDetailsCreationStatusEnum{
	"ready":                ImageVersionsDetailsCreationStatusReady,
	"error_in_creation":    ImageVersionsDetailsCreationStatusErrorInCreation,
	"creation_in_progress": ImageVersionsDetailsCreationStatusCreationInProgress,
	"scheduled":            ImageVersionsDetailsCreationStatusScheduled,
}

// GetImageVersionsDetailsCreationStatusEnumValues Enumerates the set of values for ImageVersionsDetailsCreationStatusEnum
func GetImageVersionsDetailsCreationStatusEnumValues() []ImageVersionsDetailsCreationStatusEnum {
	values := make([]ImageVersionsDetailsCreationStatusEnum, 0)
	for _, v := range mappingImageVersionsDetailsCreationStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetImageVersionsDetailsCreationStatusEnumStringValues Enumerates the set of values in String for ImageVersionsDetailsCreationStatusEnum
func GetImageVersionsDetailsCreationStatusEnumStringValues() []string {
	return []string{
		"READY",
		"ERROR_IN_CREATION",
		"CREATION_IN_PROGRESS",
		"SCHEDULED",
	}
}

// GetMappingImageVersionsDetailsCreationStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingImageVersionsDetailsCreationStatusEnum(val string) (ImageVersionsDetailsCreationStatusEnum, bool) {
	enum, ok := mappingImageVersionsDetailsCreationStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ImageVersionsDetailsStatusEnum Enum with underlying type: string
type ImageVersionsDetailsStatusEnum string

// Set of constants representing the allowable values for ImageVersionsDetailsStatusEnum
const (
	ImageVersionsDetailsStatusCurrent     ImageVersionsDetailsStatusEnum = "CURRENT"
	ImageVersionsDetailsStatusActive      ImageVersionsDetailsStatusEnum = "ACTIVE"
	ImageVersionsDetailsStatusObsolete    ImageVersionsDetailsStatusEnum = "OBSOLETE"
	ImageVersionsDetailsStatusRestricted  ImageVersionsDetailsStatusEnum = "RESTRICTED"
	ImageVersionsDetailsStatusDraft       ImageVersionsDetailsStatusEnum = "DRAFT"
	ImageVersionsDetailsStatusDeleted     ImageVersionsDetailsStatusEnum = "DELETED"
	ImageVersionsDetailsStatusUnavailable ImageVersionsDetailsStatusEnum = "UNAVAILABLE"
)

var mappingImageVersionsDetailsStatusEnum = map[string]ImageVersionsDetailsStatusEnum{
	"CURRENT":     ImageVersionsDetailsStatusCurrent,
	"ACTIVE":      ImageVersionsDetailsStatusActive,
	"OBSOLETE":    ImageVersionsDetailsStatusObsolete,
	"RESTRICTED":  ImageVersionsDetailsStatusRestricted,
	"DRAFT":       ImageVersionsDetailsStatusDraft,
	"DELETED":     ImageVersionsDetailsStatusDeleted,
	"UNAVAILABLE": ImageVersionsDetailsStatusUnavailable,
}

var mappingImageVersionsDetailsStatusEnumLowerCase = map[string]ImageVersionsDetailsStatusEnum{
	"current":     ImageVersionsDetailsStatusCurrent,
	"active":      ImageVersionsDetailsStatusActive,
	"obsolete":    ImageVersionsDetailsStatusObsolete,
	"restricted":  ImageVersionsDetailsStatusRestricted,
	"draft":       ImageVersionsDetailsStatusDraft,
	"deleted":     ImageVersionsDetailsStatusDeleted,
	"unavailable": ImageVersionsDetailsStatusUnavailable,
}

// GetImageVersionsDetailsStatusEnumValues Enumerates the set of values for ImageVersionsDetailsStatusEnum
func GetImageVersionsDetailsStatusEnumValues() []ImageVersionsDetailsStatusEnum {
	values := make([]ImageVersionsDetailsStatusEnum, 0)
	for _, v := range mappingImageVersionsDetailsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetImageVersionsDetailsStatusEnumStringValues Enumerates the set of values in String for ImageVersionsDetailsStatusEnum
func GetImageVersionsDetailsStatusEnumStringValues() []string {
	return []string{
		"CURRENT",
		"ACTIVE",
		"OBSOLETE",
		"RESTRICTED",
		"DRAFT",
		"DELETED",
		"UNAVAILABLE",
	}
}

// GetMappingImageVersionsDetailsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingImageVersionsDetailsStatusEnum(val string) (ImageVersionsDetailsStatusEnum, bool) {
	enum, ok := mappingImageVersionsDetailsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
