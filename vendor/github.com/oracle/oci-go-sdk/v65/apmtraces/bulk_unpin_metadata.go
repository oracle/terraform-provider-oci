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

// BulkUnpinMetadata Metadata about the bulk unpin operation.  The bulk unpin operation is atomic and binary.  If the processing of any of the attributes
// in the bulk unpin request results in a processing or validation error, then none of the attributes in the request are unpinned.
type BulkUnpinMetadata struct {

	// Operation status of the bulk unpin operation.
	// SUCCESS - The bulk unpin operation has succeeded and all the attributes in the bulk unpin request have been unpinned by this operation.
	// The following are error statuses for the bulk unpin operation.
	// EMPTY_ATTRIBUTE_LIST - The bulk unpin request object was empty and did not contain any attributes to be unpinned.
	// INVALID_BULK_REQUEST - The bulk request contains invalid attribute(s), or attribute(s) that resulted in a validation error, or an attribute that resulted
	// in a processing error.
	OperationStatus BulkUnpinMetadataOperationStatusEnum `mandatory:"true" json:"operationStatus"`

	// Type of operation.
	OperationType BulkUnpinMetadataOperationTypeEnum `mandatory:"true" json:"operationType"`

	// Total number attributes (both string and numeric) in TRACES namespace that were unpinned.
	AttributesUnpinned *int `mandatory:"true" json:"attributesUnpinned"`

	// Total number attributes (both string and numeric) in SYNTHETIC namespace that were unpinned.
	SyntheticAttributesUnpinned *int `mandatory:"false" json:"syntheticAttributesUnpinned"`
}

func (m BulkUnpinMetadata) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BulkUnpinMetadata) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBulkUnpinMetadataOperationStatusEnum(string(m.OperationStatus)); !ok && m.OperationStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationStatus: %s. Supported values are: %s.", m.OperationStatus, strings.Join(GetBulkUnpinMetadataOperationStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBulkUnpinMetadataOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetBulkUnpinMetadataOperationTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BulkUnpinMetadataOperationStatusEnum Enum with underlying type: string
type BulkUnpinMetadataOperationStatusEnum string

// Set of constants representing the allowable values for BulkUnpinMetadataOperationStatusEnum
const (
	BulkUnpinMetadataOperationStatusSuccess            BulkUnpinMetadataOperationStatusEnum = "SUCCESS"
	BulkUnpinMetadataOperationStatusEmptyAttributeList BulkUnpinMetadataOperationStatusEnum = "EMPTY_ATTRIBUTE_LIST"
	BulkUnpinMetadataOperationStatusInvalidBulkRequest BulkUnpinMetadataOperationStatusEnum = "INVALID_BULK_REQUEST"
)

var mappingBulkUnpinMetadataOperationStatusEnum = map[string]BulkUnpinMetadataOperationStatusEnum{
	"SUCCESS":              BulkUnpinMetadataOperationStatusSuccess,
	"EMPTY_ATTRIBUTE_LIST": BulkUnpinMetadataOperationStatusEmptyAttributeList,
	"INVALID_BULK_REQUEST": BulkUnpinMetadataOperationStatusInvalidBulkRequest,
}

var mappingBulkUnpinMetadataOperationStatusEnumLowerCase = map[string]BulkUnpinMetadataOperationStatusEnum{
	"success":              BulkUnpinMetadataOperationStatusSuccess,
	"empty_attribute_list": BulkUnpinMetadataOperationStatusEmptyAttributeList,
	"invalid_bulk_request": BulkUnpinMetadataOperationStatusInvalidBulkRequest,
}

// GetBulkUnpinMetadataOperationStatusEnumValues Enumerates the set of values for BulkUnpinMetadataOperationStatusEnum
func GetBulkUnpinMetadataOperationStatusEnumValues() []BulkUnpinMetadataOperationStatusEnum {
	values := make([]BulkUnpinMetadataOperationStatusEnum, 0)
	for _, v := range mappingBulkUnpinMetadataOperationStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetBulkUnpinMetadataOperationStatusEnumStringValues Enumerates the set of values in String for BulkUnpinMetadataOperationStatusEnum
func GetBulkUnpinMetadataOperationStatusEnumStringValues() []string {
	return []string{
		"SUCCESS",
		"EMPTY_ATTRIBUTE_LIST",
		"INVALID_BULK_REQUEST",
	}
}

// GetMappingBulkUnpinMetadataOperationStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBulkUnpinMetadataOperationStatusEnum(val string) (BulkUnpinMetadataOperationStatusEnum, bool) {
	enum, ok := mappingBulkUnpinMetadataOperationStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BulkUnpinMetadataOperationTypeEnum Enum with underlying type: string
type BulkUnpinMetadataOperationTypeEnum string

// Set of constants representing the allowable values for BulkUnpinMetadataOperationTypeEnum
const (
	BulkUnpinMetadataOperationTypeUnpin BulkUnpinMetadataOperationTypeEnum = "UNPIN"
)

var mappingBulkUnpinMetadataOperationTypeEnum = map[string]BulkUnpinMetadataOperationTypeEnum{
	"UNPIN": BulkUnpinMetadataOperationTypeUnpin,
}

var mappingBulkUnpinMetadataOperationTypeEnumLowerCase = map[string]BulkUnpinMetadataOperationTypeEnum{
	"unpin": BulkUnpinMetadataOperationTypeUnpin,
}

// GetBulkUnpinMetadataOperationTypeEnumValues Enumerates the set of values for BulkUnpinMetadataOperationTypeEnum
func GetBulkUnpinMetadataOperationTypeEnumValues() []BulkUnpinMetadataOperationTypeEnum {
	values := make([]BulkUnpinMetadataOperationTypeEnum, 0)
	for _, v := range mappingBulkUnpinMetadataOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBulkUnpinMetadataOperationTypeEnumStringValues Enumerates the set of values in String for BulkUnpinMetadataOperationTypeEnum
func GetBulkUnpinMetadataOperationTypeEnumStringValues() []string {
	return []string{
		"UNPIN",
	}
}

// GetMappingBulkUnpinMetadataOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBulkUnpinMetadataOperationTypeEnum(val string) (BulkUnpinMetadataOperationTypeEnum, bool) {
	enum, ok := mappingBulkUnpinMetadataOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
