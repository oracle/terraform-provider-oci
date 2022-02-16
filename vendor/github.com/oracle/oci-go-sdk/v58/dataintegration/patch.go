// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// Patch The patch object contains the audit summary information and the definition of the patch.
type Patch struct {

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
	PatchType PatchPatchTypeEnum `mandatory:"false" json:"patchType,omitempty"`

	// Status of the patch applied or being applied on the application
	PatchStatus PatchPatchStatusEnum `mandatory:"false" json:"patchStatus,omitempty"`

	// List of dependent objects in this patch.
	DependentObjectMetadata []PatchObjectMetadata `mandatory:"false" json:"dependentObjectMetadata"`

	// List of objects that are published or unpublished in this patch.
	PatchObjectMetadata []PatchObjectMetadata `mandatory:"false" json:"patchObjectMetadata"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	Metadata *ObjectMetadata `mandatory:"false" json:"metadata"`

	// A key map. If provided, key is replaced with generated key. This structure provides mapping between user provided key and generated key.
	KeyMap map[string]string `mandatory:"false" json:"keyMap"`
}

func (m Patch) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Patch) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPatchPatchTypeEnum(string(m.PatchType)); !ok && m.PatchType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PatchType: %s. Supported values are: %s.", m.PatchType, strings.Join(GetPatchPatchTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPatchPatchStatusEnum(string(m.PatchStatus)); !ok && m.PatchStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PatchStatus: %s. Supported values are: %s.", m.PatchStatus, strings.Join(GetPatchPatchStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PatchPatchTypeEnum Enum with underlying type: string
type PatchPatchTypeEnum string

// Set of constants representing the allowable values for PatchPatchTypeEnum
const (
	PatchPatchTypePublish   PatchPatchTypeEnum = "PUBLISH"
	PatchPatchTypeRefresh   PatchPatchTypeEnum = "REFRESH"
	PatchPatchTypeUnpublish PatchPatchTypeEnum = "UNPUBLISH"
)

var mappingPatchPatchTypeEnum = map[string]PatchPatchTypeEnum{
	"PUBLISH":   PatchPatchTypePublish,
	"REFRESH":   PatchPatchTypeRefresh,
	"UNPUBLISH": PatchPatchTypeUnpublish,
}

// GetPatchPatchTypeEnumValues Enumerates the set of values for PatchPatchTypeEnum
func GetPatchPatchTypeEnumValues() []PatchPatchTypeEnum {
	values := make([]PatchPatchTypeEnum, 0)
	for _, v := range mappingPatchPatchTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchPatchTypeEnumStringValues Enumerates the set of values in String for PatchPatchTypeEnum
func GetPatchPatchTypeEnumStringValues() []string {
	return []string{
		"PUBLISH",
		"REFRESH",
		"UNPUBLISH",
	}
}

// GetMappingPatchPatchTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchPatchTypeEnum(val string) (PatchPatchTypeEnum, bool) {
	mappingPatchPatchTypeEnumIgnoreCase := make(map[string]PatchPatchTypeEnum)
	for k, v := range mappingPatchPatchTypeEnum {
		mappingPatchPatchTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingPatchPatchTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
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

var mappingPatchPatchStatusEnum = map[string]PatchPatchStatusEnum{
	"QUEUED":      PatchPatchStatusQueued,
	"SUCCESSFUL":  PatchPatchStatusSuccessful,
	"FAILED":      PatchPatchStatusFailed,
	"IN_PROGRESS": PatchPatchStatusInProgress,
}

// GetPatchPatchStatusEnumValues Enumerates the set of values for PatchPatchStatusEnum
func GetPatchPatchStatusEnumValues() []PatchPatchStatusEnum {
	values := make([]PatchPatchStatusEnum, 0)
	for _, v := range mappingPatchPatchStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchPatchStatusEnumStringValues Enumerates the set of values in String for PatchPatchStatusEnum
func GetPatchPatchStatusEnumStringValues() []string {
	return []string{
		"QUEUED",
		"SUCCESSFUL",
		"FAILED",
		"IN_PROGRESS",
	}
}

// GetMappingPatchPatchStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchPatchStatusEnum(val string) (PatchPatchStatusEnum, bool) {
	mappingPatchPatchStatusEnumIgnoreCase := make(map[string]PatchPatchStatusEnum)
	for k, v := range mappingPatchPatchStatusEnum {
		mappingPatchPatchStatusEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingPatchPatchStatusEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
