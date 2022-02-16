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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchChangeSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPatchChangeSummaryTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetPatchChangeSummaryTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPatchChangeSummaryActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetPatchChangeSummaryActionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PatchChangeSummaryTypeEnum Enum with underlying type: string
type PatchChangeSummaryTypeEnum string

// Set of constants representing the allowable values for PatchChangeSummaryTypeEnum
const (
	PatchChangeSummaryTypeIntegrationTask PatchChangeSummaryTypeEnum = "INTEGRATION_TASK"
	PatchChangeSummaryTypeDataLoaderTask  PatchChangeSummaryTypeEnum = "DATA_LOADER_TASK"
	PatchChangeSummaryTypePipelineTask    PatchChangeSummaryTypeEnum = "PIPELINE_TASK"
	PatchChangeSummaryTypeSqlTask         PatchChangeSummaryTypeEnum = "SQL_TASK"
	PatchChangeSummaryTypeOciDataflowTask PatchChangeSummaryTypeEnum = "OCI_DATAFLOW_TASK"
	PatchChangeSummaryTypeRestTask        PatchChangeSummaryTypeEnum = "REST_TASK"
)

var mappingPatchChangeSummaryTypeEnum = map[string]PatchChangeSummaryTypeEnum{
	"INTEGRATION_TASK":  PatchChangeSummaryTypeIntegrationTask,
	"DATA_LOADER_TASK":  PatchChangeSummaryTypeDataLoaderTask,
	"PIPELINE_TASK":     PatchChangeSummaryTypePipelineTask,
	"SQL_TASK":          PatchChangeSummaryTypeSqlTask,
	"OCI_DATAFLOW_TASK": PatchChangeSummaryTypeOciDataflowTask,
	"REST_TASK":         PatchChangeSummaryTypeRestTask,
}

// GetPatchChangeSummaryTypeEnumValues Enumerates the set of values for PatchChangeSummaryTypeEnum
func GetPatchChangeSummaryTypeEnumValues() []PatchChangeSummaryTypeEnum {
	values := make([]PatchChangeSummaryTypeEnum, 0)
	for _, v := range mappingPatchChangeSummaryTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchChangeSummaryTypeEnumStringValues Enumerates the set of values in String for PatchChangeSummaryTypeEnum
func GetPatchChangeSummaryTypeEnumStringValues() []string {
	return []string{
		"INTEGRATION_TASK",
		"DATA_LOADER_TASK",
		"PIPELINE_TASK",
		"SQL_TASK",
		"OCI_DATAFLOW_TASK",
		"REST_TASK",
	}
}

// GetMappingPatchChangeSummaryTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchChangeSummaryTypeEnum(val string) (PatchChangeSummaryTypeEnum, bool) {
	mappingPatchChangeSummaryTypeEnumIgnoreCase := make(map[string]PatchChangeSummaryTypeEnum)
	for k, v := range mappingPatchChangeSummaryTypeEnum {
		mappingPatchChangeSummaryTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingPatchChangeSummaryTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// PatchChangeSummaryActionEnum Enum with underlying type: string
type PatchChangeSummaryActionEnum string

// Set of constants representing the allowable values for PatchChangeSummaryActionEnum
const (
	PatchChangeSummaryActionCreated PatchChangeSummaryActionEnum = "CREATED"
	PatchChangeSummaryActionDeleted PatchChangeSummaryActionEnum = "DELETED"
	PatchChangeSummaryActionUpdated PatchChangeSummaryActionEnum = "UPDATED"
)

var mappingPatchChangeSummaryActionEnum = map[string]PatchChangeSummaryActionEnum{
	"CREATED": PatchChangeSummaryActionCreated,
	"DELETED": PatchChangeSummaryActionDeleted,
	"UPDATED": PatchChangeSummaryActionUpdated,
}

// GetPatchChangeSummaryActionEnumValues Enumerates the set of values for PatchChangeSummaryActionEnum
func GetPatchChangeSummaryActionEnumValues() []PatchChangeSummaryActionEnum {
	values := make([]PatchChangeSummaryActionEnum, 0)
	for _, v := range mappingPatchChangeSummaryActionEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchChangeSummaryActionEnumStringValues Enumerates the set of values in String for PatchChangeSummaryActionEnum
func GetPatchChangeSummaryActionEnumStringValues() []string {
	return []string{
		"CREATED",
		"DELETED",
		"UPDATED",
	}
}

// GetMappingPatchChangeSummaryActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchChangeSummaryActionEnum(val string) (PatchChangeSummaryActionEnum, bool) {
	mappingPatchChangeSummaryActionEnumIgnoreCase := make(map[string]PatchChangeSummaryActionEnum)
	for k, v := range mappingPatchChangeSummaryActionEnum {
		mappingPatchChangeSummaryActionEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingPatchChangeSummaryActionEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
