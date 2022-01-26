// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/v56/common"
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

// PatchSummaryPatchTypeEnum Enum with underlying type: string
type PatchSummaryPatchTypeEnum string

// Set of constants representing the allowable values for PatchSummaryPatchTypeEnum
const (
	PatchSummaryPatchTypePublish   PatchSummaryPatchTypeEnum = "PUBLISH"
	PatchSummaryPatchTypeRefresh   PatchSummaryPatchTypeEnum = "REFRESH"
	PatchSummaryPatchTypeUnpublish PatchSummaryPatchTypeEnum = "UNPUBLISH"
)

var mappingPatchSummaryPatchType = map[string]PatchSummaryPatchTypeEnum{
	"PUBLISH":   PatchSummaryPatchTypePublish,
	"REFRESH":   PatchSummaryPatchTypeRefresh,
	"UNPUBLISH": PatchSummaryPatchTypeUnpublish,
}

// GetPatchSummaryPatchTypeEnumValues Enumerates the set of values for PatchSummaryPatchTypeEnum
func GetPatchSummaryPatchTypeEnumValues() []PatchSummaryPatchTypeEnum {
	values := make([]PatchSummaryPatchTypeEnum, 0)
	for _, v := range mappingPatchSummaryPatchType {
		values = append(values, v)
	}
	return values
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

var mappingPatchSummaryPatchStatus = map[string]PatchSummaryPatchStatusEnum{
	"QUEUED":      PatchSummaryPatchStatusQueued,
	"SUCCESSFUL":  PatchSummaryPatchStatusSuccessful,
	"FAILED":      PatchSummaryPatchStatusFailed,
	"IN_PROGRESS": PatchSummaryPatchStatusInProgress,
}

// GetPatchSummaryPatchStatusEnumValues Enumerates the set of values for PatchSummaryPatchStatusEnum
func GetPatchSummaryPatchStatusEnumValues() []PatchSummaryPatchStatusEnum {
	values := make([]PatchSummaryPatchStatusEnum, 0)
	for _, v := range mappingPatchSummaryPatchStatus {
		values = append(values, v)
	}
	return values
}
