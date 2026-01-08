// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// SoftwareImage Description of SoftwareImage.
type SoftwareImage struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Name of the SoftwareImage resource
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time the SoftwareImage was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the SoftwareImage.
	LifecycleState SoftwareImageLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// Description of SoftwareImage
	Description *string `mandatory:"false" json:"description"`

	// The deployment type of the image.
	DeploymentType DeploymentTypeEnum `mandatory:"false" json:"deploymentType,omitempty"`

	// The time the SoftwareImage was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// ocid of the user, who last modified this SoftwareImage
	LastModifiedBy *string `mandatory:"false" json:"lastModifiedBy"`

	// ocid of the user, who created the SoftwareImage
	Owner *string `mandatory:"false" json:"owner"`

	// Platform of the SoftwareImage
	Platform *string `mandatory:"false" json:"platform"`

	// Base release version of the image
	ImageVersion *string `mandatory:"false" json:"imageVersion"`

	// Type of Software Image
	ImageType SoftwareImageImageTypeEnum `mandatory:"false" json:"imageType,omitempty"`

	// Name of the first software image version
	VersionName *string `mandatory:"false" json:"versionName"`

	// Release for the software image version
	Release *string `mandatory:"false" json:"release"`

	// One Offs for the software image version
	OneOffs []int `mandatory:"false" json:"oneOffs"`

	MosCredentials *ValidateMosCredentialDetails `mandatory:"false" json:"mosCredentials"`

	// SoftwareImage Status.
	Status SoftwareImageStatusEnum `mandatory:"false" json:"status,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m SoftwareImage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SoftwareImage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSoftwareImageLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSoftwareImageLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDeploymentTypeEnum(string(m.DeploymentType)); !ok && m.DeploymentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeploymentType: %s. Supported values are: %s.", m.DeploymentType, strings.Join(GetDeploymentTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSoftwareImageImageTypeEnum(string(m.ImageType)); !ok && m.ImageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ImageType: %s. Supported values are: %s.", m.ImageType, strings.Join(GetSoftwareImageImageTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSoftwareImageStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetSoftwareImageStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SoftwareImageImageTypeEnum Enum with underlying type: string
type SoftwareImageImageTypeEnum string

// Set of constants representing the allowable values for SoftwareImageImageTypeEnum
const (
	SoftwareImageImageTypeDb SoftwareImageImageTypeEnum = "DB"
)

var mappingSoftwareImageImageTypeEnum = map[string]SoftwareImageImageTypeEnum{
	"DB": SoftwareImageImageTypeDb,
}

var mappingSoftwareImageImageTypeEnumLowerCase = map[string]SoftwareImageImageTypeEnum{
	"db": SoftwareImageImageTypeDb,
}

// GetSoftwareImageImageTypeEnumValues Enumerates the set of values for SoftwareImageImageTypeEnum
func GetSoftwareImageImageTypeEnumValues() []SoftwareImageImageTypeEnum {
	values := make([]SoftwareImageImageTypeEnum, 0)
	for _, v := range mappingSoftwareImageImageTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSoftwareImageImageTypeEnumStringValues Enumerates the set of values in String for SoftwareImageImageTypeEnum
func GetSoftwareImageImageTypeEnumStringValues() []string {
	return []string{
		"DB",
	}
}

// GetMappingSoftwareImageImageTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSoftwareImageImageTypeEnum(val string) (SoftwareImageImageTypeEnum, bool) {
	enum, ok := mappingSoftwareImageImageTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SoftwareImageStatusEnum Enum with underlying type: string
type SoftwareImageStatusEnum string

// Set of constants representing the allowable values for SoftwareImageStatusEnum
const (
	SoftwareImageStatusDraft      SoftwareImageStatusEnum = "DRAFT"
	SoftwareImageStatusDeleted    SoftwareImageStatusEnum = "DELETED"
	SoftwareImageStatusScheduled  SoftwareImageStatusEnum = "SCHEDULED"
	SoftwareImageStatusProduction SoftwareImageStatusEnum = "PRODUCTION"
)

var mappingSoftwareImageStatusEnum = map[string]SoftwareImageStatusEnum{
	"DRAFT":      SoftwareImageStatusDraft,
	"DELETED":    SoftwareImageStatusDeleted,
	"SCHEDULED":  SoftwareImageStatusScheduled,
	"PRODUCTION": SoftwareImageStatusProduction,
}

var mappingSoftwareImageStatusEnumLowerCase = map[string]SoftwareImageStatusEnum{
	"draft":      SoftwareImageStatusDraft,
	"deleted":    SoftwareImageStatusDeleted,
	"scheduled":  SoftwareImageStatusScheduled,
	"production": SoftwareImageStatusProduction,
}

// GetSoftwareImageStatusEnumValues Enumerates the set of values for SoftwareImageStatusEnum
func GetSoftwareImageStatusEnumValues() []SoftwareImageStatusEnum {
	values := make([]SoftwareImageStatusEnum, 0)
	for _, v := range mappingSoftwareImageStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetSoftwareImageStatusEnumStringValues Enumerates the set of values in String for SoftwareImageStatusEnum
func GetSoftwareImageStatusEnumStringValues() []string {
	return []string{
		"DRAFT",
		"DELETED",
		"SCHEDULED",
		"PRODUCTION",
	}
}

// GetMappingSoftwareImageStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSoftwareImageStatusEnum(val string) (SoftwareImageStatusEnum, bool) {
	enum, ok := mappingSoftwareImageStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SoftwareImageLifecycleStateEnum Enum with underlying type: string
type SoftwareImageLifecycleStateEnum string

// Set of constants representing the allowable values for SoftwareImageLifecycleStateEnum
const (
	SoftwareImageLifecycleStateCreating SoftwareImageLifecycleStateEnum = "CREATING"
	SoftwareImageLifecycleStateUpdating SoftwareImageLifecycleStateEnum = "UPDATING"
	SoftwareImageLifecycleStateActive   SoftwareImageLifecycleStateEnum = "ACTIVE"
	SoftwareImageLifecycleStateDeleting SoftwareImageLifecycleStateEnum = "DELETING"
	SoftwareImageLifecycleStateDeleted  SoftwareImageLifecycleStateEnum = "DELETED"
	SoftwareImageLifecycleStateFailed   SoftwareImageLifecycleStateEnum = "FAILED"
)

var mappingSoftwareImageLifecycleStateEnum = map[string]SoftwareImageLifecycleStateEnum{
	"CREATING": SoftwareImageLifecycleStateCreating,
	"UPDATING": SoftwareImageLifecycleStateUpdating,
	"ACTIVE":   SoftwareImageLifecycleStateActive,
	"DELETING": SoftwareImageLifecycleStateDeleting,
	"DELETED":  SoftwareImageLifecycleStateDeleted,
	"FAILED":   SoftwareImageLifecycleStateFailed,
}

var mappingSoftwareImageLifecycleStateEnumLowerCase = map[string]SoftwareImageLifecycleStateEnum{
	"creating": SoftwareImageLifecycleStateCreating,
	"updating": SoftwareImageLifecycleStateUpdating,
	"active":   SoftwareImageLifecycleStateActive,
	"deleting": SoftwareImageLifecycleStateDeleting,
	"deleted":  SoftwareImageLifecycleStateDeleted,
	"failed":   SoftwareImageLifecycleStateFailed,
}

// GetSoftwareImageLifecycleStateEnumValues Enumerates the set of values for SoftwareImageLifecycleStateEnum
func GetSoftwareImageLifecycleStateEnumValues() []SoftwareImageLifecycleStateEnum {
	values := make([]SoftwareImageLifecycleStateEnum, 0)
	for _, v := range mappingSoftwareImageLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSoftwareImageLifecycleStateEnumStringValues Enumerates the set of values in String for SoftwareImageLifecycleStateEnum
func GetSoftwareImageLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingSoftwareImageLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSoftwareImageLifecycleStateEnum(val string) (SoftwareImageLifecycleStateEnum, bool) {
	enum, ok := mappingSoftwareImageLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
