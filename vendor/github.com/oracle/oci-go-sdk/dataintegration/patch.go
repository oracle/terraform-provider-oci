// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Patch The patch object contains the audit summary information and the definition of the patch.
type Patch struct {

	// The key of the object.
	Key *string `mandatory:"false" json:"key"`

	// The type of the object.
	ModelType *string `mandatory:"false" json:"modelType"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	Name *string `mandatory:"false" json:"name"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be edited by the user.
	Identifier *string `mandatory:"false" json:"identifier"`

	// The date and time the patch was applied, in the timestamp format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimePatched *common.SDKTime `mandatory:"false" json:"timePatched"`

	// The errors encountered while applying the patch, if any.
	ErrorMessages map[string]string `mandatory:"false" json:"errorMessages"`

	// The application version of the patch.
	ApplicationVersion *int `mandatory:"false" json:"applicationVersion"`

	// The type of the patch applied or being applied on the application.
	PatchType PatchPatchTypeEnum `mandatory:"false" json:"patchType,omitempty"`

	// Status of the patch applied or being applied on the application
	PatchStatus PatchPatchStatusEnum `mandatory:"false" json:"patchStatus,omitempty"`

	// List of dependent objects in this patch.
	DependentObjectMetadata []PatchObjectMetadata `mandatory:"false" json:"dependentObjectMetadata"`

	// List of objects that are published / unpublished in this patch.
	PatchObjectMetadata []PatchObjectMetadata `mandatory:"false" json:"patchObjectMetadata"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	Metadata *ObjectMetadata `mandatory:"false" json:"metadata"`

	// A map, if provided key is replaced with generated key, this structure provides mapping between user provided key and generated key
	KeyMap map[string]string `mandatory:"false" json:"keyMap"`
}

func (m Patch) String() string {
	return common.PointerString(m)
}

// PatchPatchTypeEnum Enum with underlying type: string
type PatchPatchTypeEnum string

// Set of constants representing the allowable values for PatchPatchTypeEnum
const (
	PatchPatchTypePublish   PatchPatchTypeEnum = "PUBLISH"
	PatchPatchTypeUnpublish PatchPatchTypeEnum = "UNPUBLISH"
)

var mappingPatchPatchType = map[string]PatchPatchTypeEnum{
	"PUBLISH":   PatchPatchTypePublish,
	"UNPUBLISH": PatchPatchTypeUnpublish,
}

// GetPatchPatchTypeEnumValues Enumerates the set of values for PatchPatchTypeEnum
func GetPatchPatchTypeEnumValues() []PatchPatchTypeEnum {
	values := make([]PatchPatchTypeEnum, 0)
	for _, v := range mappingPatchPatchType {
		values = append(values, v)
	}
	return values
}

// PatchPatchStatusEnum Enum with underlying type: string
type PatchPatchStatusEnum string

// Set of constants representing the allowable values for PatchPatchStatusEnum
const (
	PatchPatchStatusQueued     PatchPatchStatusEnum = "QUEUED"
	PatchPatchStatusSuccessful PatchPatchStatusEnum = "SUCCESSFUL"
	PatchPatchStatusFailed     PatchPatchStatusEnum = "FAILED"
	PatchPatchStatusInProgress PatchPatchStatusEnum = "IN_PROGRESS"
)

var mappingPatchPatchStatus = map[string]PatchPatchStatusEnum{
	"QUEUED":      PatchPatchStatusQueued,
	"SUCCESSFUL":  PatchPatchStatusSuccessful,
	"FAILED":      PatchPatchStatusFailed,
	"IN_PROGRESS": PatchPatchStatusInProgress,
}

// GetPatchPatchStatusEnumValues Enumerates the set of values for PatchPatchStatusEnum
func GetPatchPatchStatusEnumValues() []PatchPatchStatusEnum {
	values := make([]PatchPatchStatusEnum, 0)
	for _, v := range mappingPatchPatchStatus {
		values = append(values, v)
	}
	return values
}
