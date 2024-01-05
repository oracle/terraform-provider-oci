// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PatchSummary The patch summary type contains the audit summary information and the definition of the patch.
type PatchSummary struct {

	// The object key.
	Key *string `mandatory:"false" json:"key"`

	// The object type.
	ModelType *string `mandatory:"false" json:"modelType"`

	// The object's model version.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	// The date and time the patch was applied, in the timestamp format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimePatched *common.SDKTime `mandatory:"false" json:"timePatched"`

	// The errors encountered while applying the patch, if any.
	ErrorMessages map[string]string `mandatory:"false" json:"errorMessages"`

	// The application version of the patch.
	ApplicationVersion *int `mandatory:"false" json:"applicationVersion"`

	// The type of the patch applied or being applied on the application.
	PatchType PatchSummaryPatchTypeEnum `mandatory:"false" json:"patchType,omitempty"`

	// Status of the patch applied or being applied on the application
	PatchStatus PatchSummaryPatchStatusEnum `mandatory:"false" json:"patchStatus,omitempty"`

	// List of dependent objects in this patch.
	DependentObjectMetadata []PatchObjectMetadata `mandatory:"false" json:"dependentObjectMetadata"`

	// List of objects that are published or unpublished in this patch.
	PatchObjectMetadata []PatchObjectMetadata `mandatory:"false" json:"patchObjectMetadata"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	Metadata *ObjectMetadata `mandatory:"false" json:"metadata"`

	// A key map. If provided, key is replaced with generated key. This structure provides mapping between user provided key and generated key.
	KeyMap map[string]string `mandatory:"false" json:"keyMap"`
}

func (m PatchSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPatchSummaryPatchTypeEnum(string(m.PatchType)); !ok && m.PatchType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PatchType: %s. Supported values are: %s.", m.PatchType, strings.Join(GetPatchSummaryPatchTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPatchSummaryPatchStatusEnum(string(m.PatchStatus)); !ok && m.PatchStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PatchStatus: %s. Supported values are: %s.", m.PatchStatus, strings.Join(GetPatchSummaryPatchStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PatchSummaryPatchTypeEnum Enum with underlying type: string
type PatchSummaryPatchTypeEnum string

// Set of constants representing the allowable values for PatchSummaryPatchTypeEnum
const (
	PatchSummaryPatchTypePublish   PatchSummaryPatchTypeEnum = "PUBLISH"
	PatchSummaryPatchTypeRefresh   PatchSummaryPatchTypeEnum = "REFRESH"
	PatchSummaryPatchTypeUnpublish PatchSummaryPatchTypeEnum = "UNPUBLISH"
)

var mappingPatchSummaryPatchTypeEnum = map[string]PatchSummaryPatchTypeEnum{
	"PUBLISH":   PatchSummaryPatchTypePublish,
	"REFRESH":   PatchSummaryPatchTypeRefresh,
	"UNPUBLISH": PatchSummaryPatchTypeUnpublish,
}

var mappingPatchSummaryPatchTypeEnumLowerCase = map[string]PatchSummaryPatchTypeEnum{
	"publish":   PatchSummaryPatchTypePublish,
	"refresh":   PatchSummaryPatchTypeRefresh,
	"unpublish": PatchSummaryPatchTypeUnpublish,
}

// GetPatchSummaryPatchTypeEnumValues Enumerates the set of values for PatchSummaryPatchTypeEnum
func GetPatchSummaryPatchTypeEnumValues() []PatchSummaryPatchTypeEnum {
	values := make([]PatchSummaryPatchTypeEnum, 0)
	for _, v := range mappingPatchSummaryPatchTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchSummaryPatchTypeEnumStringValues Enumerates the set of values in String for PatchSummaryPatchTypeEnum
func GetPatchSummaryPatchTypeEnumStringValues() []string {
	return []string{
		"PUBLISH",
		"REFRESH",
		"UNPUBLISH",
	}
}

// GetMappingPatchSummaryPatchTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchSummaryPatchTypeEnum(val string) (PatchSummaryPatchTypeEnum, bool) {
	enum, ok := mappingPatchSummaryPatchTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// PatchSummaryPatchStatusEnum Enum with underlying type: string
type PatchSummaryPatchStatusEnum string

// Set of constants representing the allowable values for PatchSummaryPatchStatusEnum
const (
	PatchSummaryPatchStatusQueued     PatchSummaryPatchStatusEnum = "QUEUED"
	PatchSummaryPatchStatusSuccessful PatchSummaryPatchStatusEnum = "SUCCESSFUL"
	PatchSummaryPatchStatusFailed     PatchSummaryPatchStatusEnum = "FAILED"
	PatchSummaryPatchStatusInProgress PatchSummaryPatchStatusEnum = "IN_PROGRESS"
)

var mappingPatchSummaryPatchStatusEnum = map[string]PatchSummaryPatchStatusEnum{
	"QUEUED":      PatchSummaryPatchStatusQueued,
	"SUCCESSFUL":  PatchSummaryPatchStatusSuccessful,
	"FAILED":      PatchSummaryPatchStatusFailed,
	"IN_PROGRESS": PatchSummaryPatchStatusInProgress,
}

var mappingPatchSummaryPatchStatusEnumLowerCase = map[string]PatchSummaryPatchStatusEnum{
	"queued":      PatchSummaryPatchStatusQueued,
	"successful":  PatchSummaryPatchStatusSuccessful,
	"failed":      PatchSummaryPatchStatusFailed,
	"in_progress": PatchSummaryPatchStatusInProgress,
}

// GetPatchSummaryPatchStatusEnumValues Enumerates the set of values for PatchSummaryPatchStatusEnum
func GetPatchSummaryPatchStatusEnumValues() []PatchSummaryPatchStatusEnum {
	values := make([]PatchSummaryPatchStatusEnum, 0)
	for _, v := range mappingPatchSummaryPatchStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchSummaryPatchStatusEnumStringValues Enumerates the set of values in String for PatchSummaryPatchStatusEnum
func GetPatchSummaryPatchStatusEnumStringValues() []string {
	return []string{
		"QUEUED",
		"SUCCESSFUL",
		"FAILED",
		"IN_PROGRESS",
	}
}

// GetMappingPatchSummaryPatchStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchSummaryPatchStatusEnum(val string) (PatchSummaryPatchStatusEnum, bool) {
	enum, ok := mappingPatchSummaryPatchStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
