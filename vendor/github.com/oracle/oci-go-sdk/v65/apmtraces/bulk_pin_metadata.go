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

// BulkPinMetadata Metadata about the bulk pin operation.  The bulk pin operation is atomic and binary.  If the processing of any of the attributes
// in the bulk pin request results in a processing or validation error, then none of the attributes in the request are pinned.
type BulkPinMetadata struct {

	// Operation status of the bulk pin operation.
	// SUCCESS - The bulk pin operation has succeeded and all the attributes in the bulk pin request have been pinned by this operation or pinned earlier.
	// The following are error statuses for the bulk pin operation.
	// EMPTY_ATTRIBUTE_LIST - The bulk pin request object was empty and did not contain any attributes to be pinned.
	// INVALID_BULK_REQUEST - The bulk request contains invalid attribute(s), or attribute(s) that resulted in a validation error, or an attribute that resulted
	// in a processing error.
	OperationStatus BulkPinMetadataOperationStatusEnum `mandatory:"true" json:"operationStatus"`

	// Type of operation.
	OperationType BulkPinMetadataOperationTypeEnum `mandatory:"true" json:"operationType"`

	// Total number attributes (both string and numeric) in TRACES namespace that were pinned.
	AttributesPinned *int `mandatory:"true" json:"attributesPinned"`

	// Total number attributes (both string and numeric) in SYNTHETIC namespace that were pinned.
	SyntheticAttributesPinned *int `mandatory:"false" json:"syntheticAttributesPinned"`
}

func (m BulkPinMetadata) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BulkPinMetadata) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBulkPinMetadataOperationStatusEnum(string(m.OperationStatus)); !ok && m.OperationStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationStatus: %s. Supported values are: %s.", m.OperationStatus, strings.Join(GetBulkPinMetadataOperationStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBulkPinMetadataOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetBulkPinMetadataOperationTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BulkPinMetadataOperationStatusEnum Enum with underlying type: string
type BulkPinMetadataOperationStatusEnum string

// Set of constants representing the allowable values for BulkPinMetadataOperationStatusEnum
const (
	BulkPinMetadataOperationStatusSuccess            BulkPinMetadataOperationStatusEnum = "SUCCESS"
	BulkPinMetadataOperationStatusEmptyAttributeList BulkPinMetadataOperationStatusEnum = "EMPTY_ATTRIBUTE_LIST"
	BulkPinMetadataOperationStatusInvalidBulkRequest BulkPinMetadataOperationStatusEnum = "INVALID_BULK_REQUEST"
)

var mappingBulkPinMetadataOperationStatusEnum = map[string]BulkPinMetadataOperationStatusEnum{
	"SUCCESS":              BulkPinMetadataOperationStatusSuccess,
	"EMPTY_ATTRIBUTE_LIST": BulkPinMetadataOperationStatusEmptyAttributeList,
	"INVALID_BULK_REQUEST": BulkPinMetadataOperationStatusInvalidBulkRequest,
}

var mappingBulkPinMetadataOperationStatusEnumLowerCase = map[string]BulkPinMetadataOperationStatusEnum{
	"success":              BulkPinMetadataOperationStatusSuccess,
	"empty_attribute_list": BulkPinMetadataOperationStatusEmptyAttributeList,
	"invalid_bulk_request": BulkPinMetadataOperationStatusInvalidBulkRequest,
}

// GetBulkPinMetadataOperationStatusEnumValues Enumerates the set of values for BulkPinMetadataOperationStatusEnum
func GetBulkPinMetadataOperationStatusEnumValues() []BulkPinMetadataOperationStatusEnum {
	values := make([]BulkPinMetadataOperationStatusEnum, 0)
	for _, v := range mappingBulkPinMetadataOperationStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetBulkPinMetadataOperationStatusEnumStringValues Enumerates the set of values in String for BulkPinMetadataOperationStatusEnum
func GetBulkPinMetadataOperationStatusEnumStringValues() []string {
	return []string{
		"SUCCESS",
		"EMPTY_ATTRIBUTE_LIST",
		"INVALID_BULK_REQUEST",
	}
}

// GetMappingBulkPinMetadataOperationStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBulkPinMetadataOperationStatusEnum(val string) (BulkPinMetadataOperationStatusEnum, bool) {
	enum, ok := mappingBulkPinMetadataOperationStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BulkPinMetadataOperationTypeEnum Enum with underlying type: string
type BulkPinMetadataOperationTypeEnum string

// Set of constants representing the allowable values for BulkPinMetadataOperationTypeEnum
const (
	BulkPinMetadataOperationTypePin BulkPinMetadataOperationTypeEnum = "PIN"
)

var mappingBulkPinMetadataOperationTypeEnum = map[string]BulkPinMetadataOperationTypeEnum{
	"PIN": BulkPinMetadataOperationTypePin,
}

var mappingBulkPinMetadataOperationTypeEnumLowerCase = map[string]BulkPinMetadataOperationTypeEnum{
	"pin": BulkPinMetadataOperationTypePin,
}

// GetBulkPinMetadataOperationTypeEnumValues Enumerates the set of values for BulkPinMetadataOperationTypeEnum
func GetBulkPinMetadataOperationTypeEnumValues() []BulkPinMetadataOperationTypeEnum {
	values := make([]BulkPinMetadataOperationTypeEnum, 0)
	for _, v := range mappingBulkPinMetadataOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBulkPinMetadataOperationTypeEnumStringValues Enumerates the set of values in String for BulkPinMetadataOperationTypeEnum
func GetBulkPinMetadataOperationTypeEnumStringValues() []string {
	return []string{
		"PIN",
	}
}

// GetMappingBulkPinMetadataOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBulkPinMetadataOperationTypeEnum(val string) (BulkPinMetadataOperationTypeEnum, bool) {
	enum, ok := mappingBulkPinMetadataOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
