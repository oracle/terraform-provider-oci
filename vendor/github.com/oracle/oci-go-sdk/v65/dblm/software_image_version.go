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

// SoftwareImageVersion SoftwareImageVersion for a softwareImage.
type SoftwareImageVersion struct {

	// Unique identifier that is system generated
	SoftwareImageVersionKey *int64 `mandatory:"true" json:"softwareImageVersionKey"`

	// Unique identifier of SoftwareImage
	SoftwareImageId *string `mandatory:"true" json:"softwareImageId"`

	// Display name of SoftwareImageVersion
	DisplayName *string `mandatory:"true" json:"displayName"`

	// SoftwareImageVersion release.
	ReleaseVersion *string `mandatory:"true" json:"releaseVersion"`

	// Creation status of SoftwareImageVersion
	CreationStatus SoftwareImageVersionCreationStatusEnum `mandatory:"true" json:"creationStatus"`

	// Status of SoftwareImageVersion
	Status SoftwareImageVersionStatusEnum `mandatory:"true" json:"status"`

	// External id/path of SoftwareImageVersion
	ExternalId *string `mandatory:"false" json:"externalId"`

	// The time the the SoftwareImage was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Position of SoftwareImageVersion
	Position *int `mandatory:"false" json:"position"`

	// Hashcode of SoftwareImageVersion.
	Hashcode *string `mandatory:"false" json:"hashcode"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m SoftwareImageVersion) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SoftwareImageVersion) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSoftwareImageVersionCreationStatusEnum(string(m.CreationStatus)); !ok && m.CreationStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CreationStatus: %s. Supported values are: %s.", m.CreationStatus, strings.Join(GetSoftwareImageVersionCreationStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSoftwareImageVersionStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetSoftwareImageVersionStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SoftwareImageVersionCreationStatusEnum Enum with underlying type: string
type SoftwareImageVersionCreationStatusEnum string

// Set of constants representing the allowable values for SoftwareImageVersionCreationStatusEnum
const (
	SoftwareImageVersionCreationStatusReady              SoftwareImageVersionCreationStatusEnum = "READY"
	SoftwareImageVersionCreationStatusErrorInCreation    SoftwareImageVersionCreationStatusEnum = "ERROR_IN_CREATION"
	SoftwareImageVersionCreationStatusCreationInProgress SoftwareImageVersionCreationStatusEnum = "CREATION_IN_PROGRESS"
	SoftwareImageVersionCreationStatusScheduled          SoftwareImageVersionCreationStatusEnum = "SCHEDULED"
)

var mappingSoftwareImageVersionCreationStatusEnum = map[string]SoftwareImageVersionCreationStatusEnum{
	"READY":                SoftwareImageVersionCreationStatusReady,
	"ERROR_IN_CREATION":    SoftwareImageVersionCreationStatusErrorInCreation,
	"CREATION_IN_PROGRESS": SoftwareImageVersionCreationStatusCreationInProgress,
	"SCHEDULED":            SoftwareImageVersionCreationStatusScheduled,
}

var mappingSoftwareImageVersionCreationStatusEnumLowerCase = map[string]SoftwareImageVersionCreationStatusEnum{
	"ready":                SoftwareImageVersionCreationStatusReady,
	"error_in_creation":    SoftwareImageVersionCreationStatusErrorInCreation,
	"creation_in_progress": SoftwareImageVersionCreationStatusCreationInProgress,
	"scheduled":            SoftwareImageVersionCreationStatusScheduled,
}

// GetSoftwareImageVersionCreationStatusEnumValues Enumerates the set of values for SoftwareImageVersionCreationStatusEnum
func GetSoftwareImageVersionCreationStatusEnumValues() []SoftwareImageVersionCreationStatusEnum {
	values := make([]SoftwareImageVersionCreationStatusEnum, 0)
	for _, v := range mappingSoftwareImageVersionCreationStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetSoftwareImageVersionCreationStatusEnumStringValues Enumerates the set of values in String for SoftwareImageVersionCreationStatusEnum
func GetSoftwareImageVersionCreationStatusEnumStringValues() []string {
	return []string{
		"READY",
		"ERROR_IN_CREATION",
		"CREATION_IN_PROGRESS",
		"SCHEDULED",
	}
}

// GetMappingSoftwareImageVersionCreationStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSoftwareImageVersionCreationStatusEnum(val string) (SoftwareImageVersionCreationStatusEnum, bool) {
	enum, ok := mappingSoftwareImageVersionCreationStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SoftwareImageVersionStatusEnum Enum with underlying type: string
type SoftwareImageVersionStatusEnum string

// Set of constants representing the allowable values for SoftwareImageVersionStatusEnum
const (
	SoftwareImageVersionStatusCurrent    SoftwareImageVersionStatusEnum = "CURRENT"
	SoftwareImageVersionStatusActive     SoftwareImageVersionStatusEnum = "ACTIVE"
	SoftwareImageVersionStatusObsolete   SoftwareImageVersionStatusEnum = "OBSOLETE"
	SoftwareImageVersionStatusRestricted SoftwareImageVersionStatusEnum = "RESTRICTED"
	SoftwareImageVersionStatusDraft      SoftwareImageVersionStatusEnum = "DRAFT"
	SoftwareImageVersionStatusDeleted    SoftwareImageVersionStatusEnum = "DELETED"
)

var mappingSoftwareImageVersionStatusEnum = map[string]SoftwareImageVersionStatusEnum{
	"CURRENT":    SoftwareImageVersionStatusCurrent,
	"ACTIVE":     SoftwareImageVersionStatusActive,
	"OBSOLETE":   SoftwareImageVersionStatusObsolete,
	"RESTRICTED": SoftwareImageVersionStatusRestricted,
	"DRAFT":      SoftwareImageVersionStatusDraft,
	"DELETED":    SoftwareImageVersionStatusDeleted,
}

var mappingSoftwareImageVersionStatusEnumLowerCase = map[string]SoftwareImageVersionStatusEnum{
	"current":    SoftwareImageVersionStatusCurrent,
	"active":     SoftwareImageVersionStatusActive,
	"obsolete":   SoftwareImageVersionStatusObsolete,
	"restricted": SoftwareImageVersionStatusRestricted,
	"draft":      SoftwareImageVersionStatusDraft,
	"deleted":    SoftwareImageVersionStatusDeleted,
}

// GetSoftwareImageVersionStatusEnumValues Enumerates the set of values for SoftwareImageVersionStatusEnum
func GetSoftwareImageVersionStatusEnumValues() []SoftwareImageVersionStatusEnum {
	values := make([]SoftwareImageVersionStatusEnum, 0)
	for _, v := range mappingSoftwareImageVersionStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetSoftwareImageVersionStatusEnumStringValues Enumerates the set of values in String for SoftwareImageVersionStatusEnum
func GetSoftwareImageVersionStatusEnumStringValues() []string {
	return []string{
		"CURRENT",
		"ACTIVE",
		"OBSOLETE",
		"RESTRICTED",
		"DRAFT",
		"DELETED",
	}
}

// GetMappingSoftwareImageVersionStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSoftwareImageVersionStatusEnum(val string) (SoftwareImageVersionStatusEnum, bool) {
	enum, ok := mappingSoftwareImageVersionStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
