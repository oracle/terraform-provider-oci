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

// BulkActivationMetadata Metadata about the bulk activation operation.  The bulk activation operation is atomic and binary.  If the processing of any of the attributes in the bulk
// activation request results in a processing or validation error, then none of the attributes in the request are activated.  The bulk activation request succeeds
// only when all the attributes in the bulk activation request are processed and they get a successful attributeStatus back.
type BulkActivationMetadata struct {

	// Operation status of the bulk activation operation.  The bulk  activation operation could have either a success or an error status as defined below.
	// The following is a success status for the bulk activation operation.
	// SUCCESS - The bulk activation operation has succeeded and all the attributes in the bulk activation request have been activated by this operation or activated earlier.
	// The following are error statuses for the bulk activation operation.  Note that none of the attributes (string or numeric) in the bulk request have been activated by this bulk
	// activation operation if any of the below statuses.
	// EMPTY_ATTRIBUTE_LIST - The bulk activation request object was empty and did not contain any attributes to be activated.
	// NUMERIC_ATTRIBUTE_LIMIT_EXCEEDED - The number of numeric attributes in the bulk request exceeded the maximum limit (100) of numeric attributes that could be activated in the APM Domain.
	// STRING_ATTRIBUTE_LIMIT_EXCEEDED - The number of string attributes in the bulk request exceeded the maximum limit (700) of string attributes that could be activated in the APM Domain.
	// INSUFFICIENT_STRING_SPACE - There are not enough free slots available in the APM Domain to activate the string attributes present in the bulk request.
	// INSUFFICIENT_NUMERIC_SPACE - There are not enough free slots available in the APM Domain to activate the numeric attributes present in the bulk request.
	// INVALID_BULK_REQUEST - The bulk request contains invalid attribute(s), or attribute(s) that resulted in a validation error, or an attribute that resulted
	// in a processing error.
	OperationStatus BulkActivationMetadataOperationStatusEnum `mandatory:"true" json:"operationStatus"`

	// Type of operation.
	OperationType BulkActivationMetadataOperationTypeEnum `mandatory:"true" json:"operationType"`

	// Total number of string attributes in TRACES namespace that were activated.
	StringAttributesActivated *int `mandatory:"false" json:"stringAttributesActivated"`

	// Total number of numeric attributes in TRACES namespace that were activated.
	NumericAttributesActivated *int `mandatory:"false" json:"numericAttributesActivated"`

	// Total number of string attributes in SYNTHETIC namespace that were activated.
	SyntheticStringAttributesActivated *int `mandatory:"false" json:"syntheticStringAttributesActivated"`

	// Total number of numeric attributes in SYNTHETIC namespace that were activated.
	SyntheticNumericAttributesActivated *int `mandatory:"false" json:"syntheticNumericAttributesActivated"`

	// Total number of free slots available to activate string attributes in the TRACES namespace in the APM Domain.  Note that if a bulk request has succeeded, this
	// number reflects the total number of free slots available for activation of additional string attributes in the TRACES namespace in the APM Domain.
	AvailableStringAttributes *int `mandatory:"false" json:"availableStringAttributes"`

	// Total number of free slots available to activate numeric attributes in the TRACES namespace in the APM Domain.  Note that if a bulk request has succeeded, this
	// number reflects the total number of free slots available for activation of additional numeric attributes in the TRACES namespace in the APM Domain.
	AvailableNumericAttributes *int `mandatory:"false" json:"availableNumericAttributes"`

	// Total number of free slots available to activate string attributes in the SYNTHETIC namespace in the APM Domain.  Note that if a bulk request has succeeded, this
	// number reflects the total number of free synthetic slots available for activation of additional string attributes in the SYNTHETIC namespace in the APM Domain.
	AvailableSyntheticStringAttributes *int `mandatory:"false" json:"availableSyntheticStringAttributes"`

	// Total number of free slots available to activate numeric attributes in the SYNTHETIC namespace in the APM Domain.  Note that if a bulk request has succeeded, this
	// number reflects the total number of free synthetic slots available for activation of additional numeric attributes in the SYNTHETIC namespace in the APM Domain.
	AvailableSyntheticNumericAttributes *int `mandatory:"false" json:"availableSyntheticNumericAttributes"`
}

func (m BulkActivationMetadata) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BulkActivationMetadata) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBulkActivationMetadataOperationStatusEnum(string(m.OperationStatus)); !ok && m.OperationStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationStatus: %s. Supported values are: %s.", m.OperationStatus, strings.Join(GetBulkActivationMetadataOperationStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBulkActivationMetadataOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetBulkActivationMetadataOperationTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BulkActivationMetadataOperationStatusEnum Enum with underlying type: string
type BulkActivationMetadataOperationStatusEnum string

// Set of constants representing the allowable values for BulkActivationMetadataOperationStatusEnum
const (
	BulkActivationMetadataOperationStatusSuccess                       BulkActivationMetadataOperationStatusEnum = "SUCCESS"
	BulkActivationMetadataOperationStatusEmptyAttributeList            BulkActivationMetadataOperationStatusEnum = "EMPTY_ATTRIBUTE_LIST"
	BulkActivationMetadataOperationStatusNumericAttributeLimitExceeded BulkActivationMetadataOperationStatusEnum = "NUMERIC_ATTRIBUTE_LIMIT_EXCEEDED"
	BulkActivationMetadataOperationStatusStringAttributeLimitExceeded  BulkActivationMetadataOperationStatusEnum = "STRING_ATTRIBUTE_LIMIT_EXCEEDED"
	BulkActivationMetadataOperationStatusInsufficientStringSpace       BulkActivationMetadataOperationStatusEnum = "INSUFFICIENT_STRING_SPACE"
	BulkActivationMetadataOperationStatusInsufficientNumericSpace      BulkActivationMetadataOperationStatusEnum = "INSUFFICIENT_NUMERIC_SPACE"
	BulkActivationMetadataOperationStatusInvalidBulkRequest            BulkActivationMetadataOperationStatusEnum = "INVALID_BULK_REQUEST"
)

var mappingBulkActivationMetadataOperationStatusEnum = map[string]BulkActivationMetadataOperationStatusEnum{
	"SUCCESS":                          BulkActivationMetadataOperationStatusSuccess,
	"EMPTY_ATTRIBUTE_LIST":             BulkActivationMetadataOperationStatusEmptyAttributeList,
	"NUMERIC_ATTRIBUTE_LIMIT_EXCEEDED": BulkActivationMetadataOperationStatusNumericAttributeLimitExceeded,
	"STRING_ATTRIBUTE_LIMIT_EXCEEDED":  BulkActivationMetadataOperationStatusStringAttributeLimitExceeded,
	"INSUFFICIENT_STRING_SPACE":        BulkActivationMetadataOperationStatusInsufficientStringSpace,
	"INSUFFICIENT_NUMERIC_SPACE":       BulkActivationMetadataOperationStatusInsufficientNumericSpace,
	"INVALID_BULK_REQUEST":             BulkActivationMetadataOperationStatusInvalidBulkRequest,
}

var mappingBulkActivationMetadataOperationStatusEnumLowerCase = map[string]BulkActivationMetadataOperationStatusEnum{
	"success":                          BulkActivationMetadataOperationStatusSuccess,
	"empty_attribute_list":             BulkActivationMetadataOperationStatusEmptyAttributeList,
	"numeric_attribute_limit_exceeded": BulkActivationMetadataOperationStatusNumericAttributeLimitExceeded,
	"string_attribute_limit_exceeded":  BulkActivationMetadataOperationStatusStringAttributeLimitExceeded,
	"insufficient_string_space":        BulkActivationMetadataOperationStatusInsufficientStringSpace,
	"insufficient_numeric_space":       BulkActivationMetadataOperationStatusInsufficientNumericSpace,
	"invalid_bulk_request":             BulkActivationMetadataOperationStatusInvalidBulkRequest,
}

// GetBulkActivationMetadataOperationStatusEnumValues Enumerates the set of values for BulkActivationMetadataOperationStatusEnum
func GetBulkActivationMetadataOperationStatusEnumValues() []BulkActivationMetadataOperationStatusEnum {
	values := make([]BulkActivationMetadataOperationStatusEnum, 0)
	for _, v := range mappingBulkActivationMetadataOperationStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetBulkActivationMetadataOperationStatusEnumStringValues Enumerates the set of values in String for BulkActivationMetadataOperationStatusEnum
func GetBulkActivationMetadataOperationStatusEnumStringValues() []string {
	return []string{
		"SUCCESS",
		"EMPTY_ATTRIBUTE_LIST",
		"NUMERIC_ATTRIBUTE_LIMIT_EXCEEDED",
		"STRING_ATTRIBUTE_LIMIT_EXCEEDED",
		"INSUFFICIENT_STRING_SPACE",
		"INSUFFICIENT_NUMERIC_SPACE",
		"INVALID_BULK_REQUEST",
	}
}

// GetMappingBulkActivationMetadataOperationStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBulkActivationMetadataOperationStatusEnum(val string) (BulkActivationMetadataOperationStatusEnum, bool) {
	enum, ok := mappingBulkActivationMetadataOperationStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BulkActivationMetadataOperationTypeEnum Enum with underlying type: string
type BulkActivationMetadataOperationTypeEnum string

// Set of constants representing the allowable values for BulkActivationMetadataOperationTypeEnum
const (
	BulkActivationMetadataOperationTypeActivate BulkActivationMetadataOperationTypeEnum = "ACTIVATE"
)

var mappingBulkActivationMetadataOperationTypeEnum = map[string]BulkActivationMetadataOperationTypeEnum{
	"ACTIVATE": BulkActivationMetadataOperationTypeActivate,
}

var mappingBulkActivationMetadataOperationTypeEnumLowerCase = map[string]BulkActivationMetadataOperationTypeEnum{
	"activate": BulkActivationMetadataOperationTypeActivate,
}

// GetBulkActivationMetadataOperationTypeEnumValues Enumerates the set of values for BulkActivationMetadataOperationTypeEnum
func GetBulkActivationMetadataOperationTypeEnumValues() []BulkActivationMetadataOperationTypeEnum {
	values := make([]BulkActivationMetadataOperationTypeEnum, 0)
	for _, v := range mappingBulkActivationMetadataOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBulkActivationMetadataOperationTypeEnumStringValues Enumerates the set of values in String for BulkActivationMetadataOperationTypeEnum
func GetBulkActivationMetadataOperationTypeEnumStringValues() []string {
	return []string{
		"ACTIVATE",
	}
}

// GetMappingBulkActivationMetadataOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBulkActivationMetadataOperationTypeEnum(val string) (BulkActivationMetadataOperationTypeEnum, bool) {
	enum, ok := mappingBulkActivationMetadataOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
