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

// AttributePinResponse Response of an individual attribute item in the bulk pin operation.
type AttributePinResponse struct {

	// Attribute that was pinned by this bulk operation.
	AttributeName *string `mandatory:"true" json:"attributeName"`

	// Notes added to this attribute.
	Notes *string `mandatory:"true" json:"notes"`

	// Type of operation - pin.
	OperationType AttributePinResponseOperationTypeEnum `mandatory:"true" json:"operationType"`

	// Namespace of the attribute whose properties were updated.  The attributeNameSpace will default to TRACES if it is
	// not passed in.
	AttributeNameSpace AttributePinResponseAttributeNameSpaceEnum `mandatory:"true" json:"attributeNameSpace"`

	// Status of the attribute after this operation.  The attribute can have one of the following statuses after the pin operation.  The attribute
	// can have either a success status or an error status.  The status of the attribute must be correlated with the operation status property in the bulk operation metadata
	// object.  The bulk operation will be successful only when all attributes in the bulk request are processed successfully and they get a success status back.
	// The following are successful status values of individual attribute items in a bulk attribute pin operation.
	// ATTRIBUTE_PINNED - The attribute is marked pinned and associated notes have been added.
	// ATTRIBUTE_ALREADY_PINNED - The caller is trying to pin an attribute that has already been pinned.
	// DUPLICATE_ATTRIBUTE - The attribute is a duplicate of an attribute that was present in this bulk request.  Note that we deduplicate the attribute collection, process only unique attributes,
	// and call out duplicates.  A duplicate attribute in a bulk request will not prevent the processing of further attributes in the bulk operation.
	// The following values are error statuses and the bulk processing is stopped when the first error is encountered.
	// PIN_NOT_ALLOWED - The caller has asked to pin an active attribute which is not allowed.
	// INVALID_ATTRIBUTE - The attribute is invalid.
	// ATTRIBUTE_NOT_PROCESSED - The attribute was not processed, as there was another attribute in this bulk request collection that resulted in a processing error.
	AttributeStatus AttributePinResponseAttributeStatusEnum `mandatory:"true" json:"attributeStatus"`

	// Time when the attribute was activated or deactivated.  Note that ingest might not have picked up the changes even if this time has elapsed.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`
}

func (m AttributePinResponse) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AttributePinResponse) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAttributePinResponseOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetAttributePinResponseOperationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAttributePinResponseAttributeNameSpaceEnum(string(m.AttributeNameSpace)); !ok && m.AttributeNameSpace != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AttributeNameSpace: %s. Supported values are: %s.", m.AttributeNameSpace, strings.Join(GetAttributePinResponseAttributeNameSpaceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAttributePinResponseAttributeStatusEnum(string(m.AttributeStatus)); !ok && m.AttributeStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AttributeStatus: %s. Supported values are: %s.", m.AttributeStatus, strings.Join(GetAttributePinResponseAttributeStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AttributePinResponseOperationTypeEnum Enum with underlying type: string
type AttributePinResponseOperationTypeEnum string

// Set of constants representing the allowable values for AttributePinResponseOperationTypeEnum
const (
	AttributePinResponseOperationTypePin AttributePinResponseOperationTypeEnum = "PIN"
)

var mappingAttributePinResponseOperationTypeEnum = map[string]AttributePinResponseOperationTypeEnum{
	"PIN": AttributePinResponseOperationTypePin,
}

var mappingAttributePinResponseOperationTypeEnumLowerCase = map[string]AttributePinResponseOperationTypeEnum{
	"pin": AttributePinResponseOperationTypePin,
}

// GetAttributePinResponseOperationTypeEnumValues Enumerates the set of values for AttributePinResponseOperationTypeEnum
func GetAttributePinResponseOperationTypeEnumValues() []AttributePinResponseOperationTypeEnum {
	values := make([]AttributePinResponseOperationTypeEnum, 0)
	for _, v := range mappingAttributePinResponseOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAttributePinResponseOperationTypeEnumStringValues Enumerates the set of values in String for AttributePinResponseOperationTypeEnum
func GetAttributePinResponseOperationTypeEnumStringValues() []string {
	return []string{
		"PIN",
	}
}

// GetMappingAttributePinResponseOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAttributePinResponseOperationTypeEnum(val string) (AttributePinResponseOperationTypeEnum, bool) {
	enum, ok := mappingAttributePinResponseOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AttributePinResponseAttributeNameSpaceEnum Enum with underlying type: string
type AttributePinResponseAttributeNameSpaceEnum string

// Set of constants representing the allowable values for AttributePinResponseAttributeNameSpaceEnum
const (
	AttributePinResponseAttributeNameSpaceTraces    AttributePinResponseAttributeNameSpaceEnum = "TRACES"
	AttributePinResponseAttributeNameSpaceSynthetic AttributePinResponseAttributeNameSpaceEnum = "SYNTHETIC"
)

var mappingAttributePinResponseAttributeNameSpaceEnum = map[string]AttributePinResponseAttributeNameSpaceEnum{
	"TRACES":    AttributePinResponseAttributeNameSpaceTraces,
	"SYNTHETIC": AttributePinResponseAttributeNameSpaceSynthetic,
}

var mappingAttributePinResponseAttributeNameSpaceEnumLowerCase = map[string]AttributePinResponseAttributeNameSpaceEnum{
	"traces":    AttributePinResponseAttributeNameSpaceTraces,
	"synthetic": AttributePinResponseAttributeNameSpaceSynthetic,
}

// GetAttributePinResponseAttributeNameSpaceEnumValues Enumerates the set of values for AttributePinResponseAttributeNameSpaceEnum
func GetAttributePinResponseAttributeNameSpaceEnumValues() []AttributePinResponseAttributeNameSpaceEnum {
	values := make([]AttributePinResponseAttributeNameSpaceEnum, 0)
	for _, v := range mappingAttributePinResponseAttributeNameSpaceEnum {
		values = append(values, v)
	}
	return values
}

// GetAttributePinResponseAttributeNameSpaceEnumStringValues Enumerates the set of values in String for AttributePinResponseAttributeNameSpaceEnum
func GetAttributePinResponseAttributeNameSpaceEnumStringValues() []string {
	return []string{
		"TRACES",
		"SYNTHETIC",
	}
}

// GetMappingAttributePinResponseAttributeNameSpaceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAttributePinResponseAttributeNameSpaceEnum(val string) (AttributePinResponseAttributeNameSpaceEnum, bool) {
	enum, ok := mappingAttributePinResponseAttributeNameSpaceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AttributePinResponseAttributeStatusEnum Enum with underlying type: string
type AttributePinResponseAttributeStatusEnum string

// Set of constants representing the allowable values for AttributePinResponseAttributeStatusEnum
const (
	AttributePinResponseAttributeStatusAttributeAlreadyPinned AttributePinResponseAttributeStatusEnum = "ATTRIBUTE_ALREADY_PINNED"
	AttributePinResponseAttributeStatusAttributePinned        AttributePinResponseAttributeStatusEnum = "ATTRIBUTE_PINNED"
	AttributePinResponseAttributeStatusPinNotAllowed          AttributePinResponseAttributeStatusEnum = "PIN_NOT_ALLOWED"
	AttributePinResponseAttributeStatusDuplicateAttribute     AttributePinResponseAttributeStatusEnum = "DUPLICATE_ATTRIBUTE"
	AttributePinResponseAttributeStatusInvalidAttribute       AttributePinResponseAttributeStatusEnum = "INVALID_ATTRIBUTE"
	AttributePinResponseAttributeStatusAttributeNotProcessed  AttributePinResponseAttributeStatusEnum = "ATTRIBUTE_NOT_PROCESSED"
)

var mappingAttributePinResponseAttributeStatusEnum = map[string]AttributePinResponseAttributeStatusEnum{
	"ATTRIBUTE_ALREADY_PINNED": AttributePinResponseAttributeStatusAttributeAlreadyPinned,
	"ATTRIBUTE_PINNED":         AttributePinResponseAttributeStatusAttributePinned,
	"PIN_NOT_ALLOWED":          AttributePinResponseAttributeStatusPinNotAllowed,
	"DUPLICATE_ATTRIBUTE":      AttributePinResponseAttributeStatusDuplicateAttribute,
	"INVALID_ATTRIBUTE":        AttributePinResponseAttributeStatusInvalidAttribute,
	"ATTRIBUTE_NOT_PROCESSED":  AttributePinResponseAttributeStatusAttributeNotProcessed,
}

var mappingAttributePinResponseAttributeStatusEnumLowerCase = map[string]AttributePinResponseAttributeStatusEnum{
	"attribute_already_pinned": AttributePinResponseAttributeStatusAttributeAlreadyPinned,
	"attribute_pinned":         AttributePinResponseAttributeStatusAttributePinned,
	"pin_not_allowed":          AttributePinResponseAttributeStatusPinNotAllowed,
	"duplicate_attribute":      AttributePinResponseAttributeStatusDuplicateAttribute,
	"invalid_attribute":        AttributePinResponseAttributeStatusInvalidAttribute,
	"attribute_not_processed":  AttributePinResponseAttributeStatusAttributeNotProcessed,
}

// GetAttributePinResponseAttributeStatusEnumValues Enumerates the set of values for AttributePinResponseAttributeStatusEnum
func GetAttributePinResponseAttributeStatusEnumValues() []AttributePinResponseAttributeStatusEnum {
	values := make([]AttributePinResponseAttributeStatusEnum, 0)
	for _, v := range mappingAttributePinResponseAttributeStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetAttributePinResponseAttributeStatusEnumStringValues Enumerates the set of values in String for AttributePinResponseAttributeStatusEnum
func GetAttributePinResponseAttributeStatusEnumStringValues() []string {
	return []string{
		"ATTRIBUTE_ALREADY_PINNED",
		"ATTRIBUTE_PINNED",
		"PIN_NOT_ALLOWED",
		"DUPLICATE_ATTRIBUTE",
		"INVALID_ATTRIBUTE",
		"ATTRIBUTE_NOT_PROCESSED",
	}
}

// GetMappingAttributePinResponseAttributeStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAttributePinResponseAttributeStatusEnum(val string) (AttributePinResponseAttributeStatusEnum, bool) {
	enum, ok := mappingAttributePinResponseAttributeStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
