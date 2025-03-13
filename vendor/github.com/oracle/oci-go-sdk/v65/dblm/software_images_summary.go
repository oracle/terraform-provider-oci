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

// SoftwareImagesSummary Summary of the SoftwareImage.
type SoftwareImagesSummary struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// Name of the SoftwareImage resource.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time the the SoftwareImage was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// The time the SoftwareImage was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Platform of the SoftwareImage.
	Platform *string `mandatory:"false" json:"platform"`

	// Name of the current software image version
	ImageVersion *string `mandatory:"false" json:"imageVersion"`

	// base release of the SoftwareImage.
	ImageRelease *string `mandatory:"false" json:"imageRelease"`

	// SoftwareImage description.
	Description *string `mandatory:"false" json:"description"`

	// Number of databases subscribed to this software image.
	SubscribedResourcesCount *int `mandatory:"false" json:"subscribedResourcesCount"`

	// Number of databases subscribed to this software image that are drifted.
	DriftersCount *int `mandatory:"false" json:"driftersCount"`

	ResourcesPatchComplianceSummary *ResourcesPatchComplianceSummary `mandatory:"false" json:"resourcesPatchComplianceSummary"`

	// SoftwareImage Status.
	Status SoftwareImagesSummaryStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Software image version details.
	ImageVersionsDetails []ImageVersionsDetails `mandatory:"false" json:"imageVersionsDetails"`

	ImagePatchRecommendationsSummary *ImagePatchRecommendationsSummary `mandatory:"false" json:"imagePatchRecommendationsSummary"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m SoftwareImagesSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SoftwareImagesSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSoftwareImagesSummaryStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetSoftwareImagesSummaryStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SoftwareImagesSummaryStatusEnum Enum with underlying type: string
type SoftwareImagesSummaryStatusEnum string

// Set of constants representing the allowable values for SoftwareImagesSummaryStatusEnum
const (
	SoftwareImagesSummaryStatusDraft      SoftwareImagesSummaryStatusEnum = "DRAFT"
	SoftwareImagesSummaryStatusDeleted    SoftwareImagesSummaryStatusEnum = "DELETED"
	SoftwareImagesSummaryStatusScheduled  SoftwareImagesSummaryStatusEnum = "SCHEDULED"
	SoftwareImagesSummaryStatusProduction SoftwareImagesSummaryStatusEnum = "PRODUCTION"
)

var mappingSoftwareImagesSummaryStatusEnum = map[string]SoftwareImagesSummaryStatusEnum{
	"DRAFT":      SoftwareImagesSummaryStatusDraft,
	"DELETED":    SoftwareImagesSummaryStatusDeleted,
	"SCHEDULED":  SoftwareImagesSummaryStatusScheduled,
	"PRODUCTION": SoftwareImagesSummaryStatusProduction,
}

var mappingSoftwareImagesSummaryStatusEnumLowerCase = map[string]SoftwareImagesSummaryStatusEnum{
	"draft":      SoftwareImagesSummaryStatusDraft,
	"deleted":    SoftwareImagesSummaryStatusDeleted,
	"scheduled":  SoftwareImagesSummaryStatusScheduled,
	"production": SoftwareImagesSummaryStatusProduction,
}

// GetSoftwareImagesSummaryStatusEnumValues Enumerates the set of values for SoftwareImagesSummaryStatusEnum
func GetSoftwareImagesSummaryStatusEnumValues() []SoftwareImagesSummaryStatusEnum {
	values := make([]SoftwareImagesSummaryStatusEnum, 0)
	for _, v := range mappingSoftwareImagesSummaryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetSoftwareImagesSummaryStatusEnumStringValues Enumerates the set of values in String for SoftwareImagesSummaryStatusEnum
func GetSoftwareImagesSummaryStatusEnumStringValues() []string {
	return []string{
		"DRAFT",
		"DELETED",
		"SCHEDULED",
		"PRODUCTION",
	}
}

// GetMappingSoftwareImagesSummaryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSoftwareImagesSummaryStatusEnum(val string) (SoftwareImagesSummaryStatusEnum, bool) {
	enum, ok := mappingSoftwareImagesSummaryStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
