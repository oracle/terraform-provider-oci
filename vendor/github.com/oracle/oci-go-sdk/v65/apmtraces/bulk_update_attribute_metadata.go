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

// BulkUpdateAttributeMetadata Metadata about the bulk attribute update operation.  The bulk attribute update operation is atomic and binary.  If the processing of any of the attributes
// in the bulk attribute update request results in a processing or validation error, then none of the attributes updated.
type BulkUpdateAttributeMetadata struct {

	// Operation status of the bulk update attribute operation.
	// SUCCESS - The bulk attribute update request has succeeded and all the attributes in the request have been updated.
	// The following are error statuses for the bulk update attributes operation.
	// EMPTY_ATTRIBUTE_LIST - The bulk update attributes request object was empty and did not contain any attributes for which properties had to be updated.
	// INVALID_BULK_REQUEST - The bulk request contains invalid attribute(s), or attribute(s) that resulted in a validation error, or an attribute that resulted
	// in a processing error.
	OperationStatus BulkUpdateAttributeMetadataOperationStatusEnum `mandatory:"true" json:"operationStatus"`

	// Type of operation.
	OperationType BulkUpdateAttributeMetadataOperationTypeEnum `mandatory:"true" json:"operationType"`

	// Total number attributes (both string and numeric) in TRACES namespace for which properties were updated.
	AttributesUpdated *int `mandatory:"true" json:"attributesUpdated"`

	// Total number attributes (both string and numeric) in SYNTHETIC namespace for which properties were updated.
	SyntheticAttributesUpdated *int `mandatory:"false" json:"syntheticAttributesUpdated"`
}

func (m BulkUpdateAttributeMetadata) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BulkUpdateAttributeMetadata) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBulkUpdateAttributeMetadataOperationStatusEnum(string(m.OperationStatus)); !ok && m.OperationStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationStatus: %s. Supported values are: %s.", m.OperationStatus, strings.Join(GetBulkUpdateAttributeMetadataOperationStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBulkUpdateAttributeMetadataOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetBulkUpdateAttributeMetadataOperationTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BulkUpdateAttributeMetadataOperationStatusEnum Enum with underlying type: string
type BulkUpdateAttributeMetadataOperationStatusEnum string

// Set of constants representing the allowable values for BulkUpdateAttributeMetadataOperationStatusEnum
const (
	BulkUpdateAttributeMetadataOperationStatusSuccess            BulkUpdateAttributeMetadataOperationStatusEnum = "SUCCESS"
	BulkUpdateAttributeMetadataOperationStatusEmptyAttributeList BulkUpdateAttributeMetadataOperationStatusEnum = "EMPTY_ATTRIBUTE_LIST"
	BulkUpdateAttributeMetadataOperationStatusInvalidBulkRequest BulkUpdateAttributeMetadataOperationStatusEnum = "INVALID_BULK_REQUEST"
)

var mappingBulkUpdateAttributeMetadataOperationStatusEnum = map[string]BulkUpdateAttributeMetadataOperationStatusEnum{
	"SUCCESS":              BulkUpdateAttributeMetadataOperationStatusSuccess,
	"EMPTY_ATTRIBUTE_LIST": BulkUpdateAttributeMetadataOperationStatusEmptyAttributeList,
	"INVALID_BULK_REQUEST": BulkUpdateAttributeMetadataOperationStatusInvalidBulkRequest,
}

var mappingBulkUpdateAttributeMetadataOperationStatusEnumLowerCase = map[string]BulkUpdateAttributeMetadataOperationStatusEnum{
	"success":              BulkUpdateAttributeMetadataOperationStatusSuccess,
	"empty_attribute_list": BulkUpdateAttributeMetadataOperationStatusEmptyAttributeList,
	"invalid_bulk_request": BulkUpdateAttributeMetadataOperationStatusInvalidBulkRequest,
}

// GetBulkUpdateAttributeMetadataOperationStatusEnumValues Enumerates the set of values for BulkUpdateAttributeMetadataOperationStatusEnum
func GetBulkUpdateAttributeMetadataOperationStatusEnumValues() []BulkUpdateAttributeMetadataOperationStatusEnum {
	values := make([]BulkUpdateAttributeMetadataOperationStatusEnum, 0)
	for _, v := range mappingBulkUpdateAttributeMetadataOperationStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetBulkUpdateAttributeMetadataOperationStatusEnumStringValues Enumerates the set of values in String for BulkUpdateAttributeMetadataOperationStatusEnum
func GetBulkUpdateAttributeMetadataOperationStatusEnumStringValues() []string {
	return []string{
		"SUCCESS",
		"EMPTY_ATTRIBUTE_LIST",
		"INVALID_BULK_REQUEST",
	}
}

// GetMappingBulkUpdateAttributeMetadataOperationStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBulkUpdateAttributeMetadataOperationStatusEnum(val string) (BulkUpdateAttributeMetadataOperationStatusEnum, bool) {
	enum, ok := mappingBulkUpdateAttributeMetadataOperationStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BulkUpdateAttributeMetadataOperationTypeEnum Enum with underlying type: string
type BulkUpdateAttributeMetadataOperationTypeEnum string

// Set of constants representing the allowable values for BulkUpdateAttributeMetadataOperationTypeEnum
const (
	BulkUpdateAttributeMetadataOperationTypeUpdateAttributeProperties BulkUpdateAttributeMetadataOperationTypeEnum = "UPDATE_ATTRIBUTE_PROPERTIES"
)

var mappingBulkUpdateAttributeMetadataOperationTypeEnum = map[string]BulkUpdateAttributeMetadataOperationTypeEnum{
	"UPDATE_ATTRIBUTE_PROPERTIES": BulkUpdateAttributeMetadataOperationTypeUpdateAttributeProperties,
}

var mappingBulkUpdateAttributeMetadataOperationTypeEnumLowerCase = map[string]BulkUpdateAttributeMetadataOperationTypeEnum{
	"update_attribute_properties": BulkUpdateAttributeMetadataOperationTypeUpdateAttributeProperties,
}

// GetBulkUpdateAttributeMetadataOperationTypeEnumValues Enumerates the set of values for BulkUpdateAttributeMetadataOperationTypeEnum
func GetBulkUpdateAttributeMetadataOperationTypeEnumValues() []BulkUpdateAttributeMetadataOperationTypeEnum {
	values := make([]BulkUpdateAttributeMetadataOperationTypeEnum, 0)
	for _, v := range mappingBulkUpdateAttributeMetadataOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBulkUpdateAttributeMetadataOperationTypeEnumStringValues Enumerates the set of values in String for BulkUpdateAttributeMetadataOperationTypeEnum
func GetBulkUpdateAttributeMetadataOperationTypeEnumStringValues() []string {
	return []string{
		"UPDATE_ATTRIBUTE_PROPERTIES",
	}
}

// GetMappingBulkUpdateAttributeMetadataOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBulkUpdateAttributeMetadataOperationTypeEnum(val string) (BulkUpdateAttributeMetadataOperationTypeEnum, bool) {
	enum, ok := mappingBulkUpdateAttributeMetadataOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
