// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// AttributeUnpinResponse Response of an individual attribute item in the bulk un-pin operation.
type AttributeUnpinResponse struct {

	// Attribute that was un-pinned by this bulk operation.
	AttributeName *string `mandatory:"true" json:"attributeName"`

	// Type of operation - un-pin.
	OperationType AttributeUnpinResponseOperationTypeEnum `mandatory:"true" json:"operationType"`

	// Status of the attribute after this operation.  The attribute can have one of the following statuses after the un-pin operation.  The attribute
	// can have either a success status or an error status.  The status of the attribute must be correlated with the operation status property in the bulk operation metadata
	// object.  The bulk operation will be successful only when all attributes in the bulk request are processed successful and they get a success status back.
	// The following are successful status values of individual attribute items in a bulk attribute un-pin operation.
	// ATTRIBUTE_UNPINNED - The attribute is marked un-pinned and associated notes have been cleared.
	// DUPLICATE_ATTRIBUTE - The attriubute is a duplicate of an attribute that was present in this bulk request.  Note that we de-duplicate the attribute collection, process only unique attributes,
	// and call out duplicates.  A duplicate attribute in a bulk request will not prevent the processing of further attributes in the bulk operation.
	// The following values are error statuses and the bulk processing is stopped when the first error is encountered.
	// INVALID_ATTRIBUTE - The attribute is invalid.  The size of the attribute is more than a 1000 chars.
	// ATTRIBUTE_NOT_FOUND - The attribute was not found in the set of pinned attributes.
	// ATTRIBUTE_NOT_PROCESSED - The attribute was not processed, as there was another attribute in this bulk request collection that resulted in a processing error.
	AttributeStatus AttributeUnpinResponseAttributeStatusEnum `mandatory:"true" json:"attributeStatus"`

	// Time when the attribute was activated or de-activated.  Note that the span processor might not have picked up the changes even if this time has elapsed.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`
}

func (m AttributeUnpinResponse) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AttributeUnpinResponse) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAttributeUnpinResponseOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetAttributeUnpinResponseOperationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAttributeUnpinResponseAttributeStatusEnum(string(m.AttributeStatus)); !ok && m.AttributeStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AttributeStatus: %s. Supported values are: %s.", m.AttributeStatus, strings.Join(GetAttributeUnpinResponseAttributeStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AttributeUnpinResponseOperationTypeEnum Enum with underlying type: string
type AttributeUnpinResponseOperationTypeEnum string

// Set of constants representing the allowable values for AttributeUnpinResponseOperationTypeEnum
const (
	AttributeUnpinResponseOperationTypeUnpin AttributeUnpinResponseOperationTypeEnum = "UNPIN"
)

var mappingAttributeUnpinResponseOperationTypeEnum = map[string]AttributeUnpinResponseOperationTypeEnum{
	"UNPIN": AttributeUnpinResponseOperationTypeUnpin,
}

var mappingAttributeUnpinResponseOperationTypeEnumLowerCase = map[string]AttributeUnpinResponseOperationTypeEnum{
	"unpin": AttributeUnpinResponseOperationTypeUnpin,
}

// GetAttributeUnpinResponseOperationTypeEnumValues Enumerates the set of values for AttributeUnpinResponseOperationTypeEnum
func GetAttributeUnpinResponseOperationTypeEnumValues() []AttributeUnpinResponseOperationTypeEnum {
	values := make([]AttributeUnpinResponseOperationTypeEnum, 0)
	for _, v := range mappingAttributeUnpinResponseOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAttributeUnpinResponseOperationTypeEnumStringValues Enumerates the set of values in String for AttributeUnpinResponseOperationTypeEnum
func GetAttributeUnpinResponseOperationTypeEnumStringValues() []string {
	return []string{
		"UNPIN",
	}
}

// GetMappingAttributeUnpinResponseOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAttributeUnpinResponseOperationTypeEnum(val string) (AttributeUnpinResponseOperationTypeEnum, bool) {
	enum, ok := mappingAttributeUnpinResponseOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AttributeUnpinResponseAttributeStatusEnum Enum with underlying type: string
type AttributeUnpinResponseAttributeStatusEnum string

// Set of constants representing the allowable values for AttributeUnpinResponseAttributeStatusEnum
const (
	AttributeUnpinResponseAttributeStatusAttributeUnpinned     AttributeUnpinResponseAttributeStatusEnum = "ATTRIBUTE_UNPINNED"
	AttributeUnpinResponseAttributeStatusDuplicateAttribute    AttributeUnpinResponseAttributeStatusEnum = "DUPLICATE_ATTRIBUTE"
	AttributeUnpinResponseAttributeStatusInvalidAttribute      AttributeUnpinResponseAttributeStatusEnum = "INVALID_ATTRIBUTE"
	AttributeUnpinResponseAttributeStatusAttributeNotFound     AttributeUnpinResponseAttributeStatusEnum = "ATTRIBUTE_NOT_FOUND"
	AttributeUnpinResponseAttributeStatusAttributeNotProcessed AttributeUnpinResponseAttributeStatusEnum = "ATTRIBUTE_NOT_PROCESSED"
)

var mappingAttributeUnpinResponseAttributeStatusEnum = map[string]AttributeUnpinResponseAttributeStatusEnum{
	"ATTRIBUTE_UNPINNED":      AttributeUnpinResponseAttributeStatusAttributeUnpinned,
	"DUPLICATE_ATTRIBUTE":     AttributeUnpinResponseAttributeStatusDuplicateAttribute,
	"INVALID_ATTRIBUTE":       AttributeUnpinResponseAttributeStatusInvalidAttribute,
	"ATTRIBUTE_NOT_FOUND":     AttributeUnpinResponseAttributeStatusAttributeNotFound,
	"ATTRIBUTE_NOT_PROCESSED": AttributeUnpinResponseAttributeStatusAttributeNotProcessed,
}

var mappingAttributeUnpinResponseAttributeStatusEnumLowerCase = map[string]AttributeUnpinResponseAttributeStatusEnum{
	"attribute_unpinned":      AttributeUnpinResponseAttributeStatusAttributeUnpinned,
	"duplicate_attribute":     AttributeUnpinResponseAttributeStatusDuplicateAttribute,
	"invalid_attribute":       AttributeUnpinResponseAttributeStatusInvalidAttribute,
	"attribute_not_found":     AttributeUnpinResponseAttributeStatusAttributeNotFound,
	"attribute_not_processed": AttributeUnpinResponseAttributeStatusAttributeNotProcessed,
}

// GetAttributeUnpinResponseAttributeStatusEnumValues Enumerates the set of values for AttributeUnpinResponseAttributeStatusEnum
func GetAttributeUnpinResponseAttributeStatusEnumValues() []AttributeUnpinResponseAttributeStatusEnum {
	values := make([]AttributeUnpinResponseAttributeStatusEnum, 0)
	for _, v := range mappingAttributeUnpinResponseAttributeStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetAttributeUnpinResponseAttributeStatusEnumStringValues Enumerates the set of values in String for AttributeUnpinResponseAttributeStatusEnum
func GetAttributeUnpinResponseAttributeStatusEnumStringValues() []string {
	return []string{
		"ATTRIBUTE_UNPINNED",
		"DUPLICATE_ATTRIBUTE",
		"INVALID_ATTRIBUTE",
		"ATTRIBUTE_NOT_FOUND",
		"ATTRIBUTE_NOT_PROCESSED",
	}
}

// GetMappingAttributeUnpinResponseAttributeStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAttributeUnpinResponseAttributeStatusEnum(val string) (AttributeUnpinResponseAttributeStatusEnum, bool) {
	enum, ok := mappingAttributeUnpinResponseAttributeStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
