// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// BulkDeActivationMetadata Metadata about the bulk de-activation operation.  The bulk de-activation operation is atomic and binary.  If the processing of any of the attributes
// in the bulk de-activation request results in a processing or validation error, then none of the attributes in the request are de-activated.
// The bulk de-activation request succeeds only when all the attributes in the bulk de-activation request are processed and they get a successful
// attributeStatus back.
type BulkDeActivationMetadata struct {

	// Operation sttus of the bulk de-activation operation.  The bulk de-activation operation could have either a success or an error status as defined below.  Note that
	// if a bulk operation has not succeeded, we do not de-activate any tags in the request set.
	// SUCCESS - The bulk de-activation operation has succeded and all the attributes in the bulk de-activation request have been de-activated by this operation or de-activated earlier.
	// The following are error statuses for the bulk de-activation operation.  Note that none of the attributes (string or numeric) in the bulk request have been de-activated by this bulk
	// de-activation operation if any of the below statuses are returned.
	// EMPTY_ATTRIBUTE_LIST - The bulk de-activation request object was empty and did not contain any attributes to be de-activated.
	// NUMERIC_ATTRIBUTE_LIMIT_EXCEEDED - The number of numeric attributes in the bulk request exceeded the maximum limit (100) of numeric attributes that could be present in the APM Domain.
	// STRING_ATTRIBUTE_LIMIT_EXCEEDED - The number of string attributes in the bulk request exceeded the maximum limit (700) of string attributes that could be present in the APM Domain.
	// INVALID_BULK_REQUEST - The bulk request contains invalid attribute(s), or attribute(s) that resulted in a validation error, or an attribute that resulted
	// in a processing error.
	OperationStatus BulkDeActivationMetadataOperationStatusEnum `mandatory:"true" json:"operationStatus"`

	// Type of operation.
	OperationType BulkDeActivationMetadataOperationTypeEnum `mandatory:"true" json:"operationType"`

	// Total number attributes (both string and numeric) that were de-activated.
	AttributesDeActivated *int `mandatory:"false" json:"attributesDeActivated"`

	// Total number of free slots available for activation of additional string attributes.
	AvailableStringAttributes *int `mandatory:"false" json:"availableStringAttributes"`

	// Total number of free slots available for activation of additional numeric attributes.
	AvailableNumericAttributes *int `mandatory:"false" json:"availableNumericAttributes"`
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
