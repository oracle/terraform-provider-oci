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

// PatchObjectMetadata A summary type containing information about the object including its key, name and when/who created/updated it
type PatchObjectMetadata struct {

	// The key of the object.
	Key *string `mandatory:"false" json:"key"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	Name *string `mandatory:"false" json:"name"`

	// The fully qualified path of the published object which would include its project and folder.
	NamePath *string `mandatory:"false" json:"namePath"`

	// The type of the object in patch.
	Type PatchObjectMetadataTypeEnum `mandatory:"false" json:"type,omitempty"`

	// The object version.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be edited by the user.
	Identifier *string `mandatory:"false" json:"identifier"`

	// The patch action, if object was created, updated or deleted.
	Action PatchObjectMetadataActionEnum `mandatory:"false" json:"action,omitempty"`
}

func (m PatchObjectMetadata) String() string {
	return common.PointerString(m)
}

// PatchObjectMetadataTypeEnum Enum with underlying type: string
type PatchObjectMetadataTypeEnum string

// Set of constants representing the allowable values for PatchObjectMetadataTypeEnum
const (
	PatchObjectMetadataTypeIntegrationTask PatchObjectMetadataTypeEnum = "INTEGRATION_TASK"
	PatchObjectMetadataTypeDataLoaderTask  PatchObjectMetadataTypeEnum = "DATA_LOADER_TASK"
)

var mappingPatchObjectMetadataType = map[string]PatchObjectMetadataTypeEnum{
	"INTEGRATION_TASK": PatchObjectMetadataTypeIntegrationTask,
	"DATA_LOADER_TASK": PatchObjectMetadataTypeDataLoaderTask,
}

// GetPatchObjectMetadataTypeEnumValues Enumerates the set of values for PatchObjectMetadataTypeEnum
func GetPatchObjectMetadataTypeEnumValues() []PatchObjectMetadataTypeEnum {
	values := make([]PatchObjectMetadataTypeEnum, 0)
	for _, v := range mappingPatchObjectMetadataType {
		values = append(values, v)
	}
	return values
}

// PatchObjectMetadataActionEnum Enum with underlying type: string
type PatchObjectMetadataActionEnum string

// Set of constants representing the allowable values for PatchObjectMetadataActionEnum
const (
	PatchObjectMetadataActionCreated PatchObjectMetadataActionEnum = "CREATED"
	PatchObjectMetadataActionDeleted PatchObjectMetadataActionEnum = "DELETED"
	PatchObjectMetadataActionUpdated PatchObjectMetadataActionEnum = "UPDATED"
)

var mappingPatchObjectMetadataAction = map[string]PatchObjectMetadataActionEnum{
	"CREATED": PatchObjectMetadataActionCreated,
	"DELETED": PatchObjectMetadataActionDeleted,
	"UPDATED": PatchObjectMetadataActionUpdated,
}

// GetPatchObjectMetadataActionEnumValues Enumerates the set of values for PatchObjectMetadataActionEnum
func GetPatchObjectMetadataActionEnumValues() []PatchObjectMetadataActionEnum {
	values := make([]PatchObjectMetadataActionEnum, 0)
	for _, v := range mappingPatchObjectMetadataAction {
		values = append(values, v)
	}
	return values
}
