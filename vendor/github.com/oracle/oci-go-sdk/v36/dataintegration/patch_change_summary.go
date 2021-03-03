// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/v36/common"
)

// PatchChangeSummary This is the patch report summary information.
type PatchChangeSummary struct {

	// The key of the object.
	Key *string `mandatory:"false" json:"key"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// The fully qualified path of the published object, which would include its project and folder.
	NamePath *string `mandatory:"false" json:"namePath"`

	// The type of the object in patch.
	Type PatchChangeSummaryTypeEnum `mandatory:"false" json:"type,omitempty"`

	// The object version.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	// The patch action indicating if object was created, updated, or deleted.
	Action PatchChangeSummaryActionEnum `mandatory:"false" json:"action,omitempty"`
}

func (m PatchChangeSummary) String() string {
	return common.PointerString(m)
}

// PatchChangeSummaryTypeEnum Enum with underlying type: string
type PatchChangeSummaryTypeEnum string

// Set of constants representing the allowable values for PatchChangeSummaryTypeEnum
const (
	PatchChangeSummaryTypeIntegrationTask PatchChangeSummaryTypeEnum = "INTEGRATION_TASK"
	PatchChangeSummaryTypeDataLoaderTask  PatchChangeSummaryTypeEnum = "DATA_LOADER_TASK"
	PatchChangeSummaryTypePipelineTask    PatchChangeSummaryTypeEnum = "PIPELINE_TASK"
)

var mappingPatchChangeSummaryType = map[string]PatchChangeSummaryTypeEnum{
	"INTEGRATION_TASK": PatchChangeSummaryTypeIntegrationTask,
	"DATA_LOADER_TASK": PatchChangeSummaryTypeDataLoaderTask,
	"PIPELINE_TASK":    PatchChangeSummaryTypePipelineTask,
}

// GetPatchChangeSummaryTypeEnumValues Enumerates the set of values for PatchChangeSummaryTypeEnum
func GetPatchChangeSummaryTypeEnumValues() []PatchChangeSummaryTypeEnum {
	values := make([]PatchChangeSummaryTypeEnum, 0)
	for _, v := range mappingPatchChangeSummaryType {
		values = append(values, v)
	}
	return values
}

// PatchChangeSummaryActionEnum Enum with underlying type: string
type PatchChangeSummaryActionEnum string

// Set of constants representing the allowable values for PatchChangeSummaryActionEnum
const (
	PatchChangeSummaryActionCreated PatchChangeSummaryActionEnum = "CREATED"
	PatchChangeSummaryActionDeleted PatchChangeSummaryActionEnum = "DELETED"
	PatchChangeSummaryActionUpdated PatchChangeSummaryActionEnum = "UPDATED"
)

var mappingPatchChangeSummaryAction = map[string]PatchChangeSummaryActionEnum{
	"CREATED": PatchChangeSummaryActionCreated,
	"DELETED": PatchChangeSummaryActionDeleted,
	"UPDATED": PatchChangeSummaryActionUpdated,
}

// GetPatchChangeSummaryActionEnumValues Enumerates the set of values for PatchChangeSummaryActionEnum
func GetPatchChangeSummaryActionEnumValues() []PatchChangeSummaryActionEnum {
	values := make([]PatchChangeSummaryActionEnum, 0)
	for _, v := range mappingPatchChangeSummaryAction {
		values = append(values, v)
	}
	return values
}
