// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, policies, and identity domains.
//

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TaggingWorkRequest The asynchronous API request does not take effect immediately. This request spawns an asynchronous
// workflow to fulfill the request. WorkRequest objects provide visibility for in-progress workflows.
type TaggingWorkRequest struct {

	// The OCID of the work request.
	Id *string `mandatory:"true" json:"id"`

	// An enum-like description of the type of work the work request is doing.
	OperationType TaggingWorkRequestOperationTypeEnum `mandatory:"true" json:"operationType"`

	// The current status of the work request.
	Status TaggingWorkRequestStatusEnum `mandatory:"true" json:"status"`

	// The OCID of the compartment that contains the work request.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The resources this work request affects.
	Resources []WorkRequestResource `mandatory:"false" json:"resources"`

	// Date and time the work was accepted, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeAccepted *common.SDKTime `mandatory:"false" json:"timeAccepted"`

	// Date and time the work started, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// Date and time the work completed, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// How much progress the operation has made.
	PercentComplete *float32 `mandatory:"false" json:"percentComplete"`
}

func (m TaggingWorkRequest) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TaggingWorkRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTaggingWorkRequestOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetTaggingWorkRequestOperationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTaggingWorkRequestStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetTaggingWorkRequestStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TaggingWorkRequestOperationTypeEnum Enum with underlying type: string
type TaggingWorkRequestOperationTypeEnum string

// Set of constants representing the allowable values for TaggingWorkRequestOperationTypeEnum
const (
	TaggingWorkRequestOperationTypeDeleteTagDefinition        TaggingWorkRequestOperationTypeEnum = "DELETE_TAG_DEFINITION"
	TaggingWorkRequestOperationTypeDeleteNonEmptyTagNamespace TaggingWorkRequestOperationTypeEnum = "DELETE_NON_EMPTY_TAG_NAMESPACE"
	TaggingWorkRequestOperationTypeBulkDeleteTagDefinition    TaggingWorkRequestOperationTypeEnum = "BULK_DELETE_TAG_DEFINITION"
	TaggingWorkRequestOperationTypeBulkEditOfTags             TaggingWorkRequestOperationTypeEnum = "BULK_EDIT_OF_TAGS"
	TaggingWorkRequestOperationTypeImportStandardTags         TaggingWorkRequestOperationTypeEnum = "IMPORT_STANDARD_TAGS"
)

var mappingTaggingWorkRequestOperationTypeEnum = map[string]TaggingWorkRequestOperationTypeEnum{
	"DELETE_TAG_DEFINITION":          TaggingWorkRequestOperationTypeDeleteTagDefinition,
	"DELETE_NON_EMPTY_TAG_NAMESPACE": TaggingWorkRequestOperationTypeDeleteNonEmptyTagNamespace,
	"BULK_DELETE_TAG_DEFINITION":     TaggingWorkRequestOperationTypeBulkDeleteTagDefinition,
	"BULK_EDIT_OF_TAGS":              TaggingWorkRequestOperationTypeBulkEditOfTags,
	"IMPORT_STANDARD_TAGS":           TaggingWorkRequestOperationTypeImportStandardTags,
}

var mappingTaggingWorkRequestOperationTypeEnumLowerCase = map[string]TaggingWorkRequestOperationTypeEnum{
	"delete_tag_definition":          TaggingWorkRequestOperationTypeDeleteTagDefinition,
	"delete_non_empty_tag_namespace": TaggingWorkRequestOperationTypeDeleteNonEmptyTagNamespace,
	"bulk_delete_tag_definition":     TaggingWorkRequestOperationTypeBulkDeleteTagDefinition,
	"bulk_edit_of_tags":              TaggingWorkRequestOperationTypeBulkEditOfTags,
	"import_standard_tags":           TaggingWorkRequestOperationTypeImportStandardTags,
}

// GetTaggingWorkRequestOperationTypeEnumValues Enumerates the set of values for TaggingWorkRequestOperationTypeEnum
func GetTaggingWorkRequestOperationTypeEnumValues() []TaggingWorkRequestOperationTypeEnum {
	values := make([]TaggingWorkRequestOperationTypeEnum, 0)
	for _, v := range mappingTaggingWorkRequestOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTaggingWorkRequestOperationTypeEnumStringValues Enumerates the set of values in String for TaggingWorkRequestOperationTypeEnum
func GetTaggingWorkRequestOperationTypeEnumStringValues() []string {
	return []string{
		"DELETE_TAG_DEFINITION",
		"DELETE_NON_EMPTY_TAG_NAMESPACE",
		"BULK_DELETE_TAG_DEFINITION",
		"BULK_EDIT_OF_TAGS",
		"IMPORT_STANDARD_TAGS",
	}
}

// GetMappingTaggingWorkRequestOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTaggingWorkRequestOperationTypeEnum(val string) (TaggingWorkRequestOperationTypeEnum, bool) {
	enum, ok := mappingTaggingWorkRequestOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TaggingWorkRequestStatusEnum Enum with underlying type: string
type TaggingWorkRequestStatusEnum string

// Set of constants representing the allowable values for TaggingWorkRequestStatusEnum
const (
	TaggingWorkRequestStatusAccepted           TaggingWorkRequestStatusEnum = "ACCEPTED"
	TaggingWorkRequestStatusInProgress         TaggingWorkRequestStatusEnum = "IN_PROGRESS"
	TaggingWorkRequestStatusFailed             TaggingWorkRequestStatusEnum = "FAILED"
	TaggingWorkRequestStatusSucceeded          TaggingWorkRequestStatusEnum = "SUCCEEDED"
	TaggingWorkRequestStatusPartiallySucceeded TaggingWorkRequestStatusEnum = "PARTIALLY_SUCCEEDED"
	TaggingWorkRequestStatusCanceling          TaggingWorkRequestStatusEnum = "CANCELING"
	TaggingWorkRequestStatusCanceled           TaggingWorkRequestStatusEnum = "CANCELED"
)

var mappingTaggingWorkRequestStatusEnum = map[string]TaggingWorkRequestStatusEnum{
	"ACCEPTED":            TaggingWorkRequestStatusAccepted,
	"IN_PROGRESS":         TaggingWorkRequestStatusInProgress,
	"FAILED":              TaggingWorkRequestStatusFailed,
	"SUCCEEDED":           TaggingWorkRequestStatusSucceeded,
	"PARTIALLY_SUCCEEDED": TaggingWorkRequestStatusPartiallySucceeded,
	"CANCELING":           TaggingWorkRequestStatusCanceling,
	"CANCELED":            TaggingWorkRequestStatusCanceled,
}

var mappingTaggingWorkRequestStatusEnumLowerCase = map[string]TaggingWorkRequestStatusEnum{
	"accepted":            TaggingWorkRequestStatusAccepted,
	"in_progress":         TaggingWorkRequestStatusInProgress,
	"failed":              TaggingWorkRequestStatusFailed,
	"succeeded":           TaggingWorkRequestStatusSucceeded,
	"partially_succeeded": TaggingWorkRequestStatusPartiallySucceeded,
	"canceling":           TaggingWorkRequestStatusCanceling,
	"canceled":            TaggingWorkRequestStatusCanceled,
}

// GetTaggingWorkRequestStatusEnumValues Enumerates the set of values for TaggingWorkRequestStatusEnum
func GetTaggingWorkRequestStatusEnumValues() []TaggingWorkRequestStatusEnum {
	values := make([]TaggingWorkRequestStatusEnum, 0)
	for _, v := range mappingTaggingWorkRequestStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetTaggingWorkRequestStatusEnumStringValues Enumerates the set of values in String for TaggingWorkRequestStatusEnum
func GetTaggingWorkRequestStatusEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
		"PARTIALLY_SUCCEEDED",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingTaggingWorkRequestStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTaggingWorkRequestStatusEnum(val string) (TaggingWorkRequestStatusEnum, bool) {
	enum, ok := mappingTaggingWorkRequestStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
