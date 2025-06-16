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

// SoftwareImageSummary Summary of the SoftwareImage.
type SoftwareImageSummary struct {

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

	// The current state of the SoftwareImage.
	LifecycleState SoftwareImageLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

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
	Status SoftwareImageSummaryStatusEnum `mandatory:"false" json:"status,omitempty"`

	// The deployment type of the resource.
	DeploymentType DeploymentTypeEnum `mandatory:"false" json:"deploymentType,omitempty"`

	// Software image version details.
	ImageVersionsDetails []ImageVersionsDetails `mandatory:"false" json:"imageVersionsDetails"`

	ImagePatchRecommendationsSummary *ImagePatchRecommendationsSummary `mandatory:"false" json:"imagePatchRecommendationsSummary"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m SoftwareImageSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SoftwareImageSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSoftwareImageLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSoftwareImageLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSoftwareImageSummaryStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetSoftwareImageSummaryStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDeploymentTypeEnum(string(m.DeploymentType)); !ok && m.DeploymentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeploymentType: %s. Supported values are: %s.", m.DeploymentType, strings.Join(GetDeploymentTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SoftwareImageSummaryStatusEnum Enum with underlying type: string
type SoftwareImageSummaryStatusEnum string

// Set of constants representing the allowable values for SoftwareImageSummaryStatusEnum
const (
	SoftwareImageSummaryStatusDraft      SoftwareImageSummaryStatusEnum = "DRAFT"
	SoftwareImageSummaryStatusDeleted    SoftwareImageSummaryStatusEnum = "DELETED"
	SoftwareImageSummaryStatusScheduled  SoftwareImageSummaryStatusEnum = "SCHEDULED"
	SoftwareImageSummaryStatusProduction SoftwareImageSummaryStatusEnum = "PRODUCTION"
)

var mappingSoftwareImageSummaryStatusEnum = map[string]SoftwareImageSummaryStatusEnum{
	"DRAFT":      SoftwareImageSummaryStatusDraft,
	"DELETED":    SoftwareImageSummaryStatusDeleted,
	"SCHEDULED":  SoftwareImageSummaryStatusScheduled,
	"PRODUCTION": SoftwareImageSummaryStatusProduction,
}

var mappingSoftwareImageSummaryStatusEnumLowerCase = map[string]SoftwareImageSummaryStatusEnum{
	"draft":      SoftwareImageSummaryStatusDraft,
	"deleted":    SoftwareImageSummaryStatusDeleted,
	"scheduled":  SoftwareImageSummaryStatusScheduled,
	"production": SoftwareImageSummaryStatusProduction,
}

// GetSoftwareImageSummaryStatusEnumValues Enumerates the set of values for SoftwareImageSummaryStatusEnum
func GetSoftwareImageSummaryStatusEnumValues() []SoftwareImageSummaryStatusEnum {
	values := make([]SoftwareImageSummaryStatusEnum, 0)
	for _, v := range mappingSoftwareImageSummaryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetSoftwareImageSummaryStatusEnumStringValues Enumerates the set of values in String for SoftwareImageSummaryStatusEnum
func GetSoftwareImageSummaryStatusEnumStringValues() []string {
	return []string{
		"DRAFT",
		"DELETED",
		"SCHEDULED",
		"PRODUCTION",
	}
}

// GetMappingSoftwareImageSummaryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSoftwareImageSummaryStatusEnum(val string) (SoftwareImageSummaryStatusEnum, bool) {
	enum, ok := mappingSoftwareImageSummaryStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
