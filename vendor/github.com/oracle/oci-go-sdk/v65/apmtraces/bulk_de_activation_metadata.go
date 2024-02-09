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

// BulkDeActivationMetadata Metadata about the bulk deactivation operation.  The bulk deactivation operation is atomic and binary.  If the processing of any of the attributes
// in the bulk deactivation request results in a processing or validation error, then none of the attributes in the request are deactivated.
// The bulk deactivation request succeeds only when all the attributes in the bulk deactivation request are processed and they get a successful
// attributeStatus back.
type BulkDeActivationMetadata struct {

	// Operation status of the bulk deactivation operation.  The bulk deactivation operation could have either a success or an error status as defined below.  Note that
	// if a bulk operation has not succeeded, we do not deactivate any tags in the request set.
	// SUCCESS - The bulk deactivation operation has succeeded and all the attributes in the bulk deactivation request have been deactivated by this operation or deactivated earlier.
	// The following are error statuses for the bulk deactivation operation.  Note that none of the attributes (string or numeric) in the bulk request have been deactivated by this bulk
	// deactivation operation if any of the below statuses are returned.
	// EMPTY_ATTRIBUTE_LIST - The bulk deactivation request object was empty and did not contain any attributes to be deactivated.
	// NUMERIC_ATTRIBUTE_LIMIT_EXCEEDED - The number of numeric attributes in the bulk request exceeded the maximum limit (100) of numeric attributes that could be present in the APM Domain.
	// STRING_ATTRIBUTE_LIMIT_EXCEEDED - The number of string attributes in the bulk request exceeded the maximum limit (700) of string attributes that could be present in the APM Domain.
	// INVALID_BULK_REQUEST - The bulk request contains invalid attribute(s), or attribute(s) that resulted in a validation error, or an attribute that resulted
	// in a processing error.
	OperationStatus BulkDeActivationMetadataOperationStatusEnum `mandatory:"true" json:"operationStatus"`

	// Type of operation.
	OperationType BulkDeActivationMetadataOperationTypeEnum `mandatory:"true" json:"operationType"`

	// Total number attributes (both string and numeric) in TRACES namespace that were deactivated.
	AttributesDeActivated *int `mandatory:"false" json:"attributesDeActivated"`

	// Total number attributes (both string and numeric) in SYNTHETIC namespace that were deactivated.
	SyntheticAttributesDeActivated *int `mandatory:"false" json:"syntheticAttributesDeActivated"`

	// Total number of free slots available for activation of additional string attributes in TRACES namespace in the APM Domain.
	AvailableStringAttributes *int `mandatory:"false" json:"availableStringAttributes"`

	// Total number of free slots available for activation of additional numeric attributes in TRACES namespace in the APM Domain.
	AvailableNumericAttributes *int `mandatory:"false" json:"availableNumericAttributes"`

	// Total number of free slots available for activation of additional string attributes in SYNTHETIC namespace in the APM Domain.
	AvailableSyntheticStringAttributes *int `mandatory:"false" json:"availableSyntheticStringAttributes"`

	// Total number of free slots available for activation of additional numeric attributes in SYNTHETIC namespace in the APM Domain.
	AvailableSyntheticNumericAttributes *int `mandatory:"false" json:"availableSyntheticNumericAttributes"`
}

func (m BulkDeActivationMetadata) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BulkDeActivationMetadata) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBulkDeActivationMetadataOperationStatusEnum(string(m.OperationStatus)); !ok && m.OperationStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationStatus: %s. Supported values are: %s.", m.OperationStatus, strings.Join(GetBulkDeActivationMetadataOperationStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBulkDeActivationMetadataOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetBulkDeActivationMetadataOperationTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BulkDeActivationMetadataOperationStatusEnum Enum with underlying type: string
type BulkDeActivationMetadataOperationStatusEnum string

// Set of constants representing the allowable values for BulkDeActivationMetadataOperationStatusEnum
const (
	BulkDeActivationMetadataOperationStatusSuccess                       BulkDeActivationMetadataOperationStatusEnum = "SUCCESS"
	BulkDeActivationMetadataOperationStatusEmptyAttributeList            BulkDeActivationMetadataOperationStatusEnum = "EMPTY_ATTRIBUTE_LIST"
	BulkDeActivationMetadataOperationStatusNumericAttributeLimitExceeded BulkDeActivationMetadataOperationStatusEnum = "NUMERIC_ATTRIBUTE_LIMIT_EXCEEDED"
	BulkDeActivationMetadataOperationStatusStringAttributeLimitExceeded  BulkDeActivationMetadataOperationStatusEnum = "STRING_ATTRIBUTE_LIMIT_EXCEEDED"
	BulkDeActivationMetadataOperationStatusInvalidBulkRequest            BulkDeActivationMetadataOperationStatusEnum = "INVALID_BULK_REQUEST"
)

var mappingBulkDeActivationMetadataOperationStatusEnum = map[string]BulkDeActivationMetadataOperationStatusEnum{
	"SUCCESS":                          BulkDeActivationMetadataOperationStatusSuccess,
	"EMPTY_ATTRIBUTE_LIST":             BulkDeActivationMetadataOperationStatusEmptyAttributeList,
	"NUMERIC_ATTRIBUTE_LIMIT_EXCEEDED": BulkDeActivationMetadataOperationStatusNumericAttributeLimitExceeded,
	"STRING_ATTRIBUTE_LIMIT_EXCEEDED":  BulkDeActivationMetadataOperationStatusStringAttributeLimitExceeded,
	"INVALID_BULK_REQUEST":             BulkDeActivationMetadataOperationStatusInvalidBulkRequest,
}

var mappingBulkDeActivationMetadataOperationStatusEnumLowerCase = map[string]BulkDeActivationMetadataOperationStatusEnum{
	"success":                          BulkDeActivationMetadataOperationStatusSuccess,
	"empty_attribute_list":             BulkDeActivationMetadataOperationStatusEmptyAttributeList,
	"numeric_attribute_limit_exceeded": BulkDeActivationMetadataOperationStatusNumericAttributeLimitExceeded,
	"string_attribute_limit_exceeded":  BulkDeActivationMetadataOperationStatusStringAttributeLimitExceeded,
	"invalid_bulk_request":             BulkDeActivationMetadataOperationStatusInvalidBulkRequest,
}

// GetBulkDeActivationMetadataOperationStatusEnumValues Enumerates the set of values for BulkDeActivationMetadataOperationStatusEnum
func GetBulkDeActivationMetadataOperationStatusEnumValues() []BulkDeActivationMetadataOperationStatusEnum {
	values := make([]BulkDeActivationMetadataOperationStatusEnum, 0)
	for _, v := range mappingBulkDeActivationMetadataOperationStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetBulkDeActivationMetadataOperationStatusEnumStringValues Enumerates the set of values in String for BulkDeActivationMetadataOperationStatusEnum
func GetBulkDeActivationMetadataOperationStatusEnumStringValues() []string {
	return []string{
		"SUCCESS",
		"EMPTY_ATTRIBUTE_LIST",
		"NUMERIC_ATTRIBUTE_LIMIT_EXCEEDED",
		"STRING_ATTRIBUTE_LIMIT_EXCEEDED",
		"INVALID_BULK_REQUEST",
	}
}

// GetMappingBulkDeActivationMetadataOperationStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBulkDeActivationMetadataOperationStatusEnum(val string) (BulkDeActivationMetadataOperationStatusEnum, bool) {
	enum, ok := mappingBulkDeActivationMetadataOperationStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BulkDeActivationMetadataOperationTypeEnum Enum with underlying type: string
type BulkDeActivationMetadataOperationTypeEnum string

// Set of constants representing the allowable values for BulkDeActivationMetadataOperationTypeEnum
const (
	BulkDeActivationMetadataOperationTypeDeactivate BulkDeActivationMetadataOperationTypeEnum = "DEACTIVATE"
)

var mappingBulkDeActivationMetadataOperationTypeEnum = map[string]BulkDeActivationMetadataOperationTypeEnum{
	"DEACTIVATE": BulkDeActivationMetadataOperationTypeDeactivate,
}

var mappingBulkDeActivationMetadataOperationTypeEnumLowerCase = map[string]BulkDeActivationMetadataOperationTypeEnum{
	"deactivate": BulkDeActivationMetadataOperationTypeDeactivate,
}

// GetBulkDeActivationMetadataOperationTypeEnumValues Enumerates the set of values for BulkDeActivationMetadataOperationTypeEnum
func GetBulkDeActivationMetadataOperationTypeEnumValues() []BulkDeActivationMetadataOperationTypeEnum {
	values := make([]BulkDeActivationMetadataOperationTypeEnum, 0)
	for _, v := range mappingBulkDeActivationMetadataOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBulkDeActivationMetadataOperationTypeEnumStringValues Enumerates the set of values in String for BulkDeActivationMetadataOperationTypeEnum
func GetBulkDeActivationMetadataOperationTypeEnumStringValues() []string {
	return []string{
		"DEACTIVATE",
	}
}

// GetMappingBulkDeActivationMetadataOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBulkDeActivationMetadataOperationTypeEnum(val string) (BulkDeActivationMetadataOperationTypeEnum, bool) {
	enum, ok := mappingBulkDeActivationMetadataOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
