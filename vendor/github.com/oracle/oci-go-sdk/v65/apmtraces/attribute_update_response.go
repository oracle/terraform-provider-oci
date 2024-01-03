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

// AttributeUpdateResponse Response of an individual attribute item in the bulk update attribute operation.
type AttributeUpdateResponse struct {

	// Attribute for which properties were updated in this bulk operation.
	AttributeName *string `mandatory:"true" json:"attributeName"`

	// Namespace of the attribute whose properties were updated.  The attributeNameSpace will default to TRACES if it is
	// not passed in.
	AttributeNameSpace AttributeUpdateResponseAttributeNameSpaceEnum `mandatory:"true" json:"attributeNameSpace"`

	// Time when the attribute's properties were updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Type of the attribute.
	AttributeType AttributeUpdateResponseAttributeTypeEnum `mandatory:"false" json:"attributeType,omitempty"`

	// Unit updated for this attribute.
	Unit AttributeUpdateResponseUnitEnum `mandatory:"false" json:"unit,omitempty"`

	// Notes for the attribute.
	Notes *string `mandatory:"false" json:"notes"`

	// Type of operation - UPDATE_ATTRIBUTE_PROPERTIES.
	OperationType AttributeUpdateResponseOperationTypeEnum `mandatory:"false" json:"operationType,omitempty"`

	// Status of the attribute after this operation.  The attribute can have one of the following statuses after the update operation.  The attribute
	// can have either a success status or an error status.  The status of the attribute must be correlated with the operation status property in the bulk operation metadata
	// object.  The bulk operation will be successful only when all attributes in the bulk request are processed successfully and they get a success status back.
	// The following are successful status values of individual attribute items in a bulk update attribute operation.
	// ATTRIBUTE_UPDATED - The attribute's properites have been updated with the given properties.
	// DUPLICATE_ATTRIBUTE - The attribute is a duplicate of an attribute that was present in this bulk request.  Note that we deduplicate the attribute collection, process only unique attributes,
	// and call out duplicates.  A duplicate attribute in a bulk request will not prevent the processing of further attributes in the bulk operation.
	// The following values are error statuses and the bulk processing is stopped when the first error is encountered.
	// INVALID_ATTRIBUTE - The attribute is invalid.
	// ATTRIBUTE_NOT_PROCESSED - The attribute was not processed, as there was another attribute in this bulk request collection that resulted in a processing error.
	// ATTRIBUTE_DOES_NOT_EXIST - Attribute was neither active nor pinned inactive.
	// ATTRIBUTE_UPDATE_NOT_ALLOWED - Attribute update is not allowed as it is an in-built system attribute.
	AttributeStatus AttributeUpdateResponseAttributeStatusEnum `mandatory:"false" json:"attributeStatus,omitempty"`
}

func (m AttributeUpdateResponse) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AttributeUpdateResponse) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAttributeUpdateResponseAttributeNameSpaceEnum(string(m.AttributeNameSpace)); !ok && m.AttributeNameSpace != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AttributeNameSpace: %s. Supported values are: %s.", m.AttributeNameSpace, strings.Join(GetAttributeUpdateResponseAttributeNameSpaceEnumStringValues(), ",")))
	}

	if _, ok := GetMappingAttributeUpdateResponseAttributeTypeEnum(string(m.AttributeType)); !ok && m.AttributeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AttributeType: %s. Supported values are: %s.", m.AttributeType, strings.Join(GetAttributeUpdateResponseAttributeTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAttributeUpdateResponseUnitEnum(string(m.Unit)); !ok && m.Unit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Unit: %s. Supported values are: %s.", m.Unit, strings.Join(GetAttributeUpdateResponseUnitEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAttributeUpdateResponseOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetAttributeUpdateResponseOperationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAttributeUpdateResponseAttributeStatusEnum(string(m.AttributeStatus)); !ok && m.AttributeStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AttributeStatus: %s. Supported values are: %s.", m.AttributeStatus, strings.Join(GetAttributeUpdateResponseAttributeStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AttributeUpdateResponseAttributeTypeEnum Enum with underlying type: string
type AttributeUpdateResponseAttributeTypeEnum string

// Set of constants representing the allowable values for AttributeUpdateResponseAttributeTypeEnum
const (
	AttributeUpdateResponseAttributeTypeNumeric AttributeUpdateResponseAttributeTypeEnum = "NUMERIC"
	AttributeUpdateResponseAttributeTypeString  AttributeUpdateResponseAttributeTypeEnum = "STRING"
)

var mappingAttributeUpdateResponseAttributeTypeEnum = map[string]AttributeUpdateResponseAttributeTypeEnum{
	"NUMERIC": AttributeUpdateResponseAttributeTypeNumeric,
	"STRING":  AttributeUpdateResponseAttributeTypeString,
}

var mappingAttributeUpdateResponseAttributeTypeEnumLowerCase = map[string]AttributeUpdateResponseAttributeTypeEnum{
	"numeric": AttributeUpdateResponseAttributeTypeNumeric,
	"string":  AttributeUpdateResponseAttributeTypeString,
}

// GetAttributeUpdateResponseAttributeTypeEnumValues Enumerates the set of values for AttributeUpdateResponseAttributeTypeEnum
func GetAttributeUpdateResponseAttributeTypeEnumValues() []AttributeUpdateResponseAttributeTypeEnum {
	values := make([]AttributeUpdateResponseAttributeTypeEnum, 0)
	for _, v := range mappingAttributeUpdateResponseAttributeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAttributeUpdateResponseAttributeTypeEnumStringValues Enumerates the set of values in String for AttributeUpdateResponseAttributeTypeEnum
func GetAttributeUpdateResponseAttributeTypeEnumStringValues() []string {
	return []string{
		"NUMERIC",
		"STRING",
	}
}

// GetMappingAttributeUpdateResponseAttributeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAttributeUpdateResponseAttributeTypeEnum(val string) (AttributeUpdateResponseAttributeTypeEnum, bool) {
	enum, ok := mappingAttributeUpdateResponseAttributeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AttributeUpdateResponseUnitEnum Enum with underlying type: string
type AttributeUpdateResponseUnitEnum string

// Set of constants representing the allowable values for AttributeUpdateResponseUnitEnum
const (
	AttributeUpdateResponseUnitNone        AttributeUpdateResponseUnitEnum = "NONE"
	AttributeUpdateResponseUnitEpochTimeMs AttributeUpdateResponseUnitEnum = "EPOCH_TIME_MS"
	AttributeUpdateResponseUnitBytes       AttributeUpdateResponseUnitEnum = "BYTES"
	AttributeUpdateResponseUnitCount       AttributeUpdateResponseUnitEnum = "COUNT"
	AttributeUpdateResponseUnitDurationMs  AttributeUpdateResponseUnitEnum = "DURATION_MS"
	AttributeUpdateResponseUnitTraceStatus AttributeUpdateResponseUnitEnum = "TRACE_STATUS"
	AttributeUpdateResponseUnitPercentage  AttributeUpdateResponseUnitEnum = "PERCENTAGE"
)

var mappingAttributeUpdateResponseUnitEnum = map[string]AttributeUpdateResponseUnitEnum{
	"NONE":          AttributeUpdateResponseUnitNone,
	"EPOCH_TIME_MS": AttributeUpdateResponseUnitEpochTimeMs,
	"BYTES":         AttributeUpdateResponseUnitBytes,
	"COUNT":         AttributeUpdateResponseUnitCount,
	"DURATION_MS":   AttributeUpdateResponseUnitDurationMs,
	"TRACE_STATUS":  AttributeUpdateResponseUnitTraceStatus,
	"PERCENTAGE":    AttributeUpdateResponseUnitPercentage,
}

var mappingAttributeUpdateResponseUnitEnumLowerCase = map[string]AttributeUpdateResponseUnitEnum{
	"none":          AttributeUpdateResponseUnitNone,
	"epoch_time_ms": AttributeUpdateResponseUnitEpochTimeMs,
	"bytes":         AttributeUpdateResponseUnitBytes,
	"count":         AttributeUpdateResponseUnitCount,
	"duration_ms":   AttributeUpdateResponseUnitDurationMs,
	"trace_status":  AttributeUpdateResponseUnitTraceStatus,
	"percentage":    AttributeUpdateResponseUnitPercentage,
}

// GetAttributeUpdateResponseUnitEnumValues Enumerates the set of values for AttributeUpdateResponseUnitEnum
func GetAttributeUpdateResponseUnitEnumValues() []AttributeUpdateResponseUnitEnum {
	values := make([]AttributeUpdateResponseUnitEnum, 0)
	for _, v := range mappingAttributeUpdateResponseUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetAttributeUpdateResponseUnitEnumStringValues Enumerates the set of values in String for AttributeUpdateResponseUnitEnum
func GetAttributeUpdateResponseUnitEnumStringValues() []string {
	return []string{
		"NONE",
		"EPOCH_TIME_MS",
		"BYTES",
		"COUNT",
		"DURATION_MS",
		"TRACE_STATUS",
		"PERCENTAGE",
	}
}

// GetMappingAttributeUpdateResponseUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAttributeUpdateResponseUnitEnum(val string) (AttributeUpdateResponseUnitEnum, bool) {
	enum, ok := mappingAttributeUpdateResponseUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AttributeUpdateResponseOperationTypeEnum Enum with underlying type: string
type AttributeUpdateResponseOperationTypeEnum string

// Set of constants representing the allowable values for AttributeUpdateResponseOperationTypeEnum
const (
	AttributeUpdateResponseOperationTypeUpdateAttributeProperties AttributeUpdateResponseOperationTypeEnum = "UPDATE_ATTRIBUTE_PROPERTIES"
)

var mappingAttributeUpdateResponseOperationTypeEnum = map[string]AttributeUpdateResponseOperationTypeEnum{
	"UPDATE_ATTRIBUTE_PROPERTIES": AttributeUpdateResponseOperationTypeUpdateAttributeProperties,
}

var mappingAttributeUpdateResponseOperationTypeEnumLowerCase = map[string]AttributeUpdateResponseOperationTypeEnum{
	"update_attribute_properties": AttributeUpdateResponseOperationTypeUpdateAttributeProperties,
}

// GetAttributeUpdateResponseOperationTypeEnumValues Enumerates the set of values for AttributeUpdateResponseOperationTypeEnum
func GetAttributeUpdateResponseOperationTypeEnumValues() []AttributeUpdateResponseOperationTypeEnum {
	values := make([]AttributeUpdateResponseOperationTypeEnum, 0)
	for _, v := range mappingAttributeUpdateResponseOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAttributeUpdateResponseOperationTypeEnumStringValues Enumerates the set of values in String for AttributeUpdateResponseOperationTypeEnum
func GetAttributeUpdateResponseOperationTypeEnumStringValues() []string {
	return []string{
		"UPDATE_ATTRIBUTE_PROPERTIES",
	}
}

// GetMappingAttributeUpdateResponseOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAttributeUpdateResponseOperationTypeEnum(val string) (AttributeUpdateResponseOperationTypeEnum, bool) {
	enum, ok := mappingAttributeUpdateResponseOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AttributeUpdateResponseAttributeStatusEnum Enum with underlying type: string
type AttributeUpdateResponseAttributeStatusEnum string

// Set of constants representing the allowable values for AttributeUpdateResponseAttributeStatusEnum
const (
	AttributeUpdateResponseAttributeStatusAttributeUpdated          AttributeUpdateResponseAttributeStatusEnum = "ATTRIBUTE_UPDATED"
	AttributeUpdateResponseAttributeStatusDuplicateAttribute        AttributeUpdateResponseAttributeStatusEnum = "DUPLICATE_ATTRIBUTE"
	AttributeUpdateResponseAttributeStatusInvalidAttribute          AttributeUpdateResponseAttributeStatusEnum = "INVALID_ATTRIBUTE"
	AttributeUpdateResponseAttributeStatusAttributeNotProcessed     AttributeUpdateResponseAttributeStatusEnum = "ATTRIBUTE_NOT_PROCESSED"
	AttributeUpdateResponseAttributeStatusAttributeDoesNotExist     AttributeUpdateResponseAttributeStatusEnum = "ATTRIBUTE_DOES_NOT_EXIST"
	AttributeUpdateResponseAttributeStatusAttributeUpdateNotAllowed AttributeUpdateResponseAttributeStatusEnum = "ATTRIBUTE_UPDATE_NOT_ALLOWED"
)

var mappingAttributeUpdateResponseAttributeStatusEnum = map[string]AttributeUpdateResponseAttributeStatusEnum{
	"ATTRIBUTE_UPDATED":            AttributeUpdateResponseAttributeStatusAttributeUpdated,
	"DUPLICATE_ATTRIBUTE":          AttributeUpdateResponseAttributeStatusDuplicateAttribute,
	"INVALID_ATTRIBUTE":            AttributeUpdateResponseAttributeStatusInvalidAttribute,
	"ATTRIBUTE_NOT_PROCESSED":      AttributeUpdateResponseAttributeStatusAttributeNotProcessed,
	"ATTRIBUTE_DOES_NOT_EXIST":     AttributeUpdateResponseAttributeStatusAttributeDoesNotExist,
	"ATTRIBUTE_UPDATE_NOT_ALLOWED": AttributeUpdateResponseAttributeStatusAttributeUpdateNotAllowed,
}

var mappingAttributeUpdateResponseAttributeStatusEnumLowerCase = map[string]AttributeUpdateResponseAttributeStatusEnum{
	"attribute_updated":            AttributeUpdateResponseAttributeStatusAttributeUpdated,
	"duplicate_attribute":          AttributeUpdateResponseAttributeStatusDuplicateAttribute,
	"invalid_attribute":            AttributeUpdateResponseAttributeStatusInvalidAttribute,
	"attribute_not_processed":      AttributeUpdateResponseAttributeStatusAttributeNotProcessed,
	"attribute_does_not_exist":     AttributeUpdateResponseAttributeStatusAttributeDoesNotExist,
	"attribute_update_not_allowed": AttributeUpdateResponseAttributeStatusAttributeUpdateNotAllowed,
}

// GetAttributeUpdateResponseAttributeStatusEnumValues Enumerates the set of values for AttributeUpdateResponseAttributeStatusEnum
func GetAttributeUpdateResponseAttributeStatusEnumValues() []AttributeUpdateResponseAttributeStatusEnum {
	values := make([]AttributeUpdateResponseAttributeStatusEnum, 0)
	for _, v := range mappingAttributeUpdateResponseAttributeStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetAttributeUpdateResponseAttributeStatusEnumStringValues Enumerates the set of values in String for AttributeUpdateResponseAttributeStatusEnum
func GetAttributeUpdateResponseAttributeStatusEnumStringValues() []string {
	return []string{
		"ATTRIBUTE_UPDATED",
		"DUPLICATE_ATTRIBUTE",
		"INVALID_ATTRIBUTE",
		"ATTRIBUTE_NOT_PROCESSED",
		"ATTRIBUTE_DOES_NOT_EXIST",
		"ATTRIBUTE_UPDATE_NOT_ALLOWED",
	}
}

// GetMappingAttributeUpdateResponseAttributeStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAttributeUpdateResponseAttributeStatusEnum(val string) (AttributeUpdateResponseAttributeStatusEnum, bool) {
	enum, ok := mappingAttributeUpdateResponseAttributeStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AttributeUpdateResponseAttributeNameSpaceEnum Enum with underlying type: string
type AttributeUpdateResponseAttributeNameSpaceEnum string

// Set of constants representing the allowable values for AttributeUpdateResponseAttributeNameSpaceEnum
const (
	AttributeUpdateResponseAttributeNameSpaceTraces    AttributeUpdateResponseAttributeNameSpaceEnum = "TRACES"
	AttributeUpdateResponseAttributeNameSpaceSynthetic AttributeUpdateResponseAttributeNameSpaceEnum = "SYNTHETIC"
)

var mappingAttributeUpdateResponseAttributeNameSpaceEnum = map[string]AttributeUpdateResponseAttributeNameSpaceEnum{
	"TRACES":    AttributeUpdateResponseAttributeNameSpaceTraces,
	"SYNTHETIC": AttributeUpdateResponseAttributeNameSpaceSynthetic,
}

var mappingAttributeUpdateResponseAttributeNameSpaceEnumLowerCase = map[string]AttributeUpdateResponseAttributeNameSpaceEnum{
	"traces":    AttributeUpdateResponseAttributeNameSpaceTraces,
	"synthetic": AttributeUpdateResponseAttributeNameSpaceSynthetic,
}

// GetAttributeUpdateResponseAttributeNameSpaceEnumValues Enumerates the set of values for AttributeUpdateResponseAttributeNameSpaceEnum
func GetAttributeUpdateResponseAttributeNameSpaceEnumValues() []AttributeUpdateResponseAttributeNameSpaceEnum {
	values := make([]AttributeUpdateResponseAttributeNameSpaceEnum, 0)
	for _, v := range mappingAttributeUpdateResponseAttributeNameSpaceEnum {
		values = append(values, v)
	}
	return values
}

// GetAttributeUpdateResponseAttributeNameSpaceEnumStringValues Enumerates the set of values in String for AttributeUpdateResponseAttributeNameSpaceEnum
func GetAttributeUpdateResponseAttributeNameSpaceEnumStringValues() []string {
	return []string{
		"TRACES",
		"SYNTHETIC",
	}
}

// GetMappingAttributeUpdateResponseAttributeNameSpaceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAttributeUpdateResponseAttributeNameSpaceEnum(val string) (AttributeUpdateResponseAttributeNameSpaceEnum, bool) {
	enum, ok := mappingAttributeUpdateResponseAttributeNameSpaceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
