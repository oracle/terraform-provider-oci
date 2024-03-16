// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Trace Explorer API
//
// Use the Application Performance Monitoring Trace Explorer API to query traces and associated spans in Trace Explorer. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmtraces

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AttributeUpdateNotesResponse Response of an individual attribute item in the bulk update notes operation.
type AttributeUpdateNotesResponse struct {

	// Attribute for which notes were added to or edited in this bulk operation.
	AttributeName *string `mandatory:"true" json:"attributeName"`

	// Notes added to or edited for this attribute.
	Notes *string `mandatory:"true" json:"notes"`

	// Type of operation - UPDATE_NOTES.
	OperationType AttributeUpdateNotesResponseOperationTypeEnum `mandatory:"true" json:"operationType"`

	// Namespace of the attribute whose notes were updated.  The attributeNameSpace will default to TRACES if it is
	// not passed in.
	AttributeNameSpace AttributeUpdateNotesResponseAttributeNameSpaceEnum `mandatory:"true" json:"attributeNameSpace"`

	// Status of the attribute after this operation.  The attribute can have one of the following statuses after the update notes operation.  The attribute
	// can have either a success status or an error status.  The status of the attribute must be correlated with the operation status property in the bulk operation metadata
	// object.  The bulk operation will be successful only when all attributes in the bulk request are processed successfully and they get a success status back.
	// The following are successful status values of individual attribute items in a bulk update notes operation.
	// ATTRIBUTE_NOTES_UPDATED - The attribute's notes have been updated with the given notes.
	// DUPLICATE_ATTRIBUTE - The attribute is a duplicate of an attribute that was present in this bulk request.  Note that we deduplicate the attribute collection, process only unique attributes,
	// and call out duplicates.  A duplicate attribute in a bulk request will not prevent the processing of further attributes in the bulk operation.
	// The following values are error statuses and the bulk processing is stopped when the first error is encountered.
	// INVALID_ATTRIBUTE - The attribute is invalid.  The length of the notes is more than a 1000 characters.
	// ATTRIBUTE_NOT_PROCESSED - The attribute was not processed, as there was another attribute in this bulk request collection that resulted in a processing error.
	// ATTRIBUTE_DOES_NOT_EXIST - Attribute was neither active nor pinned inactive.
	// NOTES_TOO_LONG - Attribute notes were too long (more than 1000 chars).
	AttributeStatus AttributeUpdateNotesResponseAttributeStatusEnum `mandatory:"true" json:"attributeStatus"`

	// Time when the attribute's notes were updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`
}

func (m AttributeUpdateNotesResponse) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AttributeUpdateNotesResponse) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAttributeUpdateNotesResponseOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetAttributeUpdateNotesResponseOperationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAttributeUpdateNotesResponseAttributeNameSpaceEnum(string(m.AttributeNameSpace)); !ok && m.AttributeNameSpace != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AttributeNameSpace: %s. Supported values are: %s.", m.AttributeNameSpace, strings.Join(GetAttributeUpdateNotesResponseAttributeNameSpaceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAttributeUpdateNotesResponseAttributeStatusEnum(string(m.AttributeStatus)); !ok && m.AttributeStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AttributeStatus: %s. Supported values are: %s.", m.AttributeStatus, strings.Join(GetAttributeUpdateNotesResponseAttributeStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AttributeUpdateNotesResponseOperationTypeEnum Enum with underlying type: string
type AttributeUpdateNotesResponseOperationTypeEnum string

// Set of constants representing the allowable values for AttributeUpdateNotesResponseOperationTypeEnum
const (
	AttributeUpdateNotesResponseOperationTypeUpdateNotes AttributeUpdateNotesResponseOperationTypeEnum = "UPDATE_NOTES"
)

var mappingAttributeUpdateNotesResponseOperationTypeEnum = map[string]AttributeUpdateNotesResponseOperationTypeEnum{
	"UPDATE_NOTES": AttributeUpdateNotesResponseOperationTypeUpdateNotes,
}

var mappingAttributeUpdateNotesResponseOperationTypeEnumLowerCase = map[string]AttributeUpdateNotesResponseOperationTypeEnum{
	"update_notes": AttributeUpdateNotesResponseOperationTypeUpdateNotes,
}

// GetAttributeUpdateNotesResponseOperationTypeEnumValues Enumerates the set of values for AttributeUpdateNotesResponseOperationTypeEnum
func GetAttributeUpdateNotesResponseOperationTypeEnumValues() []AttributeUpdateNotesResponseOperationTypeEnum {
	values := make([]AttributeUpdateNotesResponseOperationTypeEnum, 0)
	for _, v := range mappingAttributeUpdateNotesResponseOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAttributeUpdateNotesResponseOperationTypeEnumStringValues Enumerates the set of values in String for AttributeUpdateNotesResponseOperationTypeEnum
func GetAttributeUpdateNotesResponseOperationTypeEnumStringValues() []string {
	return []string{
		"UPDATE_NOTES",
	}
}

// GetMappingAttributeUpdateNotesResponseOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAttributeUpdateNotesResponseOperationTypeEnum(val string) (AttributeUpdateNotesResponseOperationTypeEnum, bool) {
	enum, ok := mappingAttributeUpdateNotesResponseOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AttributeUpdateNotesResponseAttributeNameSpaceEnum Enum with underlying type: string
type AttributeUpdateNotesResponseAttributeNameSpaceEnum string

// Set of constants representing the allowable values for AttributeUpdateNotesResponseAttributeNameSpaceEnum
const (
	AttributeUpdateNotesResponseAttributeNameSpaceTraces    AttributeUpdateNotesResponseAttributeNameSpaceEnum = "TRACES"
	AttributeUpdateNotesResponseAttributeNameSpaceSynthetic AttributeUpdateNotesResponseAttributeNameSpaceEnum = "SYNTHETIC"
)

var mappingAttributeUpdateNotesResponseAttributeNameSpaceEnum = map[string]AttributeUpdateNotesResponseAttributeNameSpaceEnum{
	"TRACES":    AttributeUpdateNotesResponseAttributeNameSpaceTraces,
	"SYNTHETIC": AttributeUpdateNotesResponseAttributeNameSpaceSynthetic,
}

var mappingAttributeUpdateNotesResponseAttributeNameSpaceEnumLowerCase = map[string]AttributeUpdateNotesResponseAttributeNameSpaceEnum{
	"traces":    AttributeUpdateNotesResponseAttributeNameSpaceTraces,
	"synthetic": AttributeUpdateNotesResponseAttributeNameSpaceSynthetic,
}

// GetAttributeUpdateNotesResponseAttributeNameSpaceEnumValues Enumerates the set of values for AttributeUpdateNotesResponseAttributeNameSpaceEnum
func GetAttributeUpdateNotesResponseAttributeNameSpaceEnumValues() []AttributeUpdateNotesResponseAttributeNameSpaceEnum {
	values := make([]AttributeUpdateNotesResponseAttributeNameSpaceEnum, 0)
	for _, v := range mappingAttributeUpdateNotesResponseAttributeNameSpaceEnum {
		values = append(values, v)
	}
	return values
}

// GetAttributeUpdateNotesResponseAttributeNameSpaceEnumStringValues Enumerates the set of values in String for AttributeUpdateNotesResponseAttributeNameSpaceEnum
func GetAttributeUpdateNotesResponseAttributeNameSpaceEnumStringValues() []string {
	return []string{
		"TRACES",
		"SYNTHETIC",
	}
}

// GetMappingAttributeUpdateNotesResponseAttributeNameSpaceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAttributeUpdateNotesResponseAttributeNameSpaceEnum(val string) (AttributeUpdateNotesResponseAttributeNameSpaceEnum, bool) {
	enum, ok := mappingAttributeUpdateNotesResponseAttributeNameSpaceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AttributeUpdateNotesResponseAttributeStatusEnum Enum with underlying type: string
type AttributeUpdateNotesResponseAttributeStatusEnum string

// Set of constants representing the allowable values for AttributeUpdateNotesResponseAttributeStatusEnum
const (
	AttributeUpdateNotesResponseAttributeStatusAttributeNotesUpdated AttributeUpdateNotesResponseAttributeStatusEnum = "ATTRIBUTE_NOTES_UPDATED"
	AttributeUpdateNotesResponseAttributeStatusDuplicateAttribute    AttributeUpdateNotesResponseAttributeStatusEnum = "DUPLICATE_ATTRIBUTE"
	AttributeUpdateNotesResponseAttributeStatusInvalidAttribute      AttributeUpdateNotesResponseAttributeStatusEnum = "INVALID_ATTRIBUTE"
	AttributeUpdateNotesResponseAttributeStatusAttributeNotProcessed AttributeUpdateNotesResponseAttributeStatusEnum = "ATTRIBUTE_NOT_PROCESSED"
	AttributeUpdateNotesResponseAttributeStatusAttributeDoesNotExist AttributeUpdateNotesResponseAttributeStatusEnum = "ATTRIBUTE_DOES_NOT_EXIST"
	AttributeUpdateNotesResponseAttributeStatusNotesTooLong          AttributeUpdateNotesResponseAttributeStatusEnum = "NOTES_TOO_LONG"
)

var mappingAttributeUpdateNotesResponseAttributeStatusEnum = map[string]AttributeUpdateNotesResponseAttributeStatusEnum{
	"ATTRIBUTE_NOTES_UPDATED":  AttributeUpdateNotesResponseAttributeStatusAttributeNotesUpdated,
	"DUPLICATE_ATTRIBUTE":      AttributeUpdateNotesResponseAttributeStatusDuplicateAttribute,
	"INVALID_ATTRIBUTE":        AttributeUpdateNotesResponseAttributeStatusInvalidAttribute,
	"ATTRIBUTE_NOT_PROCESSED":  AttributeUpdateNotesResponseAttributeStatusAttributeNotProcessed,
	"ATTRIBUTE_DOES_NOT_EXIST": AttributeUpdateNotesResponseAttributeStatusAttributeDoesNotExist,
	"NOTES_TOO_LONG":           AttributeUpdateNotesResponseAttributeStatusNotesTooLong,
}

var mappingAttributeUpdateNotesResponseAttributeStatusEnumLowerCase = map[string]AttributeUpdateNotesResponseAttributeStatusEnum{
	"attribute_notes_updated":  AttributeUpdateNotesResponseAttributeStatusAttributeNotesUpdated,
	"duplicate_attribute":      AttributeUpdateNotesResponseAttributeStatusDuplicateAttribute,
	"invalid_attribute":        AttributeUpdateNotesResponseAttributeStatusInvalidAttribute,
	"attribute_not_processed":  AttributeUpdateNotesResponseAttributeStatusAttributeNotProcessed,
	"attribute_does_not_exist": AttributeUpdateNotesResponseAttributeStatusAttributeDoesNotExist,
	"notes_too_long":           AttributeUpdateNotesResponseAttributeStatusNotesTooLong,
}

// GetAttributeUpdateNotesResponseAttributeStatusEnumValues Enumerates the set of values for AttributeUpdateNotesResponseAttributeStatusEnum
func GetAttributeUpdateNotesResponseAttributeStatusEnumValues() []AttributeUpdateNotesResponseAttributeStatusEnum {
	values := make([]AttributeUpdateNotesResponseAttributeStatusEnum, 0)
	for _, v := range mappingAttributeUpdateNotesResponseAttributeStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetAttributeUpdateNotesResponseAttributeStatusEnumStringValues Enumerates the set of values in String for AttributeUpdateNotesResponseAttributeStatusEnum
func GetAttributeUpdateNotesResponseAttributeStatusEnumStringValues() []string {
	return []string{
		"ATTRIBUTE_NOTES_UPDATED",
		"DUPLICATE_ATTRIBUTE",
		"INVALID_ATTRIBUTE",
		"ATTRIBUTE_NOT_PROCESSED",
		"ATTRIBUTE_DOES_NOT_EXIST",
		"NOTES_TOO_LONG",
	}
}

// GetMappingAttributeUpdateNotesResponseAttributeStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAttributeUpdateNotesResponseAttributeStatusEnum(val string) (AttributeUpdateNotesResponseAttributeStatusEnum, bool) {
	enum, ok := mappingAttributeUpdateNotesResponseAttributeStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
