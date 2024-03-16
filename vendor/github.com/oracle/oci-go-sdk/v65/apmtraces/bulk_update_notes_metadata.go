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

// BulkUpdateNotesMetadata Metadata about the bulk update notes operation.  The bulk update notes operation is atomic and binary.  If the processing of any of the attributes
// in the bulk update notes request results in a processing or validation error, then none of the notes in the update notes request are updated.
type BulkUpdateNotesMetadata struct {

	// Operation status of the bulk update notes operation.
	// SUCCESS - The bulk updates notes request has succeeded and all the attributes in the bulk update notes request have been updated with the provided notes.
	// The following are error statuses for the bulk update notes operation.
	// EMPTY_ATTRIBUTE_LIST - The bulk update notes request object was empty and did not contain any attributes for which notes had to be updated.
	// INVALID_BULK_REQUEST - The bulk request contains invalid attribute(s), or attribute(s) that resulted in a validation error, or an attribute that resulted
	// in a processing error.
	OperationStatus BulkUpdateNotesMetadataOperationStatusEnum `mandatory:"true" json:"operationStatus"`

	// Type of operation.
	OperationType BulkUpdateNotesMetadataOperationTypeEnum `mandatory:"true" json:"operationType"`

	// Total number attributes (both string and numeric) in TRACES namespace for which notes were updated.
	AttributesNotesUpdated *int `mandatory:"true" json:"attributesNotesUpdated"`

	// Total number attributes (both string and numeric) in SYNTHETIC namespace that were pinned.
	SyntheticAttributesPinned *int `mandatory:"false" json:"syntheticAttributesPinned"`
}

func (m BulkUpdateNotesMetadata) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BulkUpdateNotesMetadata) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBulkUpdateNotesMetadataOperationStatusEnum(string(m.OperationStatus)); !ok && m.OperationStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationStatus: %s. Supported values are: %s.", m.OperationStatus, strings.Join(GetBulkUpdateNotesMetadataOperationStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBulkUpdateNotesMetadataOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetBulkUpdateNotesMetadataOperationTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BulkUpdateNotesMetadataOperationStatusEnum Enum with underlying type: string
type BulkUpdateNotesMetadataOperationStatusEnum string

// Set of constants representing the allowable values for BulkUpdateNotesMetadataOperationStatusEnum
const (
	BulkUpdateNotesMetadataOperationStatusSuccess            BulkUpdateNotesMetadataOperationStatusEnum = "SUCCESS"
	BulkUpdateNotesMetadataOperationStatusEmptyAttributeList BulkUpdateNotesMetadataOperationStatusEnum = "EMPTY_ATTRIBUTE_LIST"
	BulkUpdateNotesMetadataOperationStatusInvalidBulkRequest BulkUpdateNotesMetadataOperationStatusEnum = "INVALID_BULK_REQUEST"
)

var mappingBulkUpdateNotesMetadataOperationStatusEnum = map[string]BulkUpdateNotesMetadataOperationStatusEnum{
	"SUCCESS":              BulkUpdateNotesMetadataOperationStatusSuccess,
	"EMPTY_ATTRIBUTE_LIST": BulkUpdateNotesMetadataOperationStatusEmptyAttributeList,
	"INVALID_BULK_REQUEST": BulkUpdateNotesMetadataOperationStatusInvalidBulkRequest,
}

var mappingBulkUpdateNotesMetadataOperationStatusEnumLowerCase = map[string]BulkUpdateNotesMetadataOperationStatusEnum{
	"success":              BulkUpdateNotesMetadataOperationStatusSuccess,
	"empty_attribute_list": BulkUpdateNotesMetadataOperationStatusEmptyAttributeList,
	"invalid_bulk_request": BulkUpdateNotesMetadataOperationStatusInvalidBulkRequest,
}

// GetBulkUpdateNotesMetadataOperationStatusEnumValues Enumerates the set of values for BulkUpdateNotesMetadataOperationStatusEnum
func GetBulkUpdateNotesMetadataOperationStatusEnumValues() []BulkUpdateNotesMetadataOperationStatusEnum {
	values := make([]BulkUpdateNotesMetadataOperationStatusEnum, 0)
	for _, v := range mappingBulkUpdateNotesMetadataOperationStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetBulkUpdateNotesMetadataOperationStatusEnumStringValues Enumerates the set of values in String for BulkUpdateNotesMetadataOperationStatusEnum
func GetBulkUpdateNotesMetadataOperationStatusEnumStringValues() []string {
	return []string{
		"SUCCESS",
		"EMPTY_ATTRIBUTE_LIST",
		"INVALID_BULK_REQUEST",
	}
}

// GetMappingBulkUpdateNotesMetadataOperationStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBulkUpdateNotesMetadataOperationStatusEnum(val string) (BulkUpdateNotesMetadataOperationStatusEnum, bool) {
	enum, ok := mappingBulkUpdateNotesMetadataOperationStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BulkUpdateNotesMetadataOperationTypeEnum Enum with underlying type: string
type BulkUpdateNotesMetadataOperationTypeEnum string

// Set of constants representing the allowable values for BulkUpdateNotesMetadataOperationTypeEnum
const (
	BulkUpdateNotesMetadataOperationTypeUpdateNotes BulkUpdateNotesMetadataOperationTypeEnum = "UPDATE_NOTES"
)

var mappingBulkUpdateNotesMetadataOperationTypeEnum = map[string]BulkUpdateNotesMetadataOperationTypeEnum{
	"UPDATE_NOTES": BulkUpdateNotesMetadataOperationTypeUpdateNotes,
}

var mappingBulkUpdateNotesMetadataOperationTypeEnumLowerCase = map[string]BulkUpdateNotesMetadataOperationTypeEnum{
	"update_notes": BulkUpdateNotesMetadataOperationTypeUpdateNotes,
}

// GetBulkUpdateNotesMetadataOperationTypeEnumValues Enumerates the set of values for BulkUpdateNotesMetadataOperationTypeEnum
func GetBulkUpdateNotesMetadataOperationTypeEnumValues() []BulkUpdateNotesMetadataOperationTypeEnum {
	values := make([]BulkUpdateNotesMetadataOperationTypeEnum, 0)
	for _, v := range mappingBulkUpdateNotesMetadataOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBulkUpdateNotesMetadataOperationTypeEnumStringValues Enumerates the set of values in String for BulkUpdateNotesMetadataOperationTypeEnum
func GetBulkUpdateNotesMetadataOperationTypeEnumStringValues() []string {
	return []string{
		"UPDATE_NOTES",
	}
}

// GetMappingBulkUpdateNotesMetadataOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBulkUpdateNotesMetadataOperationTypeEnum(val string) (BulkUpdateNotesMetadataOperationTypeEnum, bool) {
	enum, ok := mappingBulkUpdateNotesMetadataOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
