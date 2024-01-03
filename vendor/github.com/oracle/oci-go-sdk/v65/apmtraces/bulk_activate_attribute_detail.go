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

// BulkActivateAttributeDetail Object that contains the details about a single attribute in the bulk request to be activated.  The attributeNameSpace and
// unit are optional parameters, and the attributeNameSpace will default to TRACES if it is not passed in.
type BulkActivateAttributeDetail struct {

	// Name of the attribute to be activated.
	AttributeName *string `mandatory:"true" json:"attributeName"`

	// Type of the attribute to be activated.
	AttributeType BulkActivateAttributeDetailAttributeTypeEnum `mandatory:"true" json:"attributeType"`

	// Unit of the attribute to be updated.
	Unit BulkActivateAttributeDetailUnitEnum `mandatory:"false" json:"unit,omitempty"`

	// Namespace of the attribute to be activated.  The attributeNameSpace will default to TRACES if it is
	// not passed in.
	AttributeNameSpace BulkActivateAttributeDetailAttributeNameSpaceEnum `mandatory:"false" json:"attributeNameSpace,omitempty"`
}

func (m BulkActivateAttributeDetail) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BulkActivateAttributeDetail) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBulkActivateAttributeDetailAttributeTypeEnum(string(m.AttributeType)); !ok && m.AttributeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AttributeType: %s. Supported values are: %s.", m.AttributeType, strings.Join(GetBulkActivateAttributeDetailAttributeTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingBulkActivateAttributeDetailUnitEnum(string(m.Unit)); !ok && m.Unit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Unit: %s. Supported values are: %s.", m.Unit, strings.Join(GetBulkActivateAttributeDetailUnitEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBulkActivateAttributeDetailAttributeNameSpaceEnum(string(m.AttributeNameSpace)); !ok && m.AttributeNameSpace != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AttributeNameSpace: %s. Supported values are: %s.", m.AttributeNameSpace, strings.Join(GetBulkActivateAttributeDetailAttributeNameSpaceEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BulkActivateAttributeDetailAttributeTypeEnum Enum with underlying type: string
type BulkActivateAttributeDetailAttributeTypeEnum string

// Set of constants representing the allowable values for BulkActivateAttributeDetailAttributeTypeEnum
const (
	BulkActivateAttributeDetailAttributeTypeNumeric BulkActivateAttributeDetailAttributeTypeEnum = "NUMERIC"
	BulkActivateAttributeDetailAttributeTypeString  BulkActivateAttributeDetailAttributeTypeEnum = "STRING"
)

var mappingBulkActivateAttributeDetailAttributeTypeEnum = map[string]BulkActivateAttributeDetailAttributeTypeEnum{
	"NUMERIC": BulkActivateAttributeDetailAttributeTypeNumeric,
	"STRING":  BulkActivateAttributeDetailAttributeTypeString,
}

var mappingBulkActivateAttributeDetailAttributeTypeEnumLowerCase = map[string]BulkActivateAttributeDetailAttributeTypeEnum{
	"numeric": BulkActivateAttributeDetailAttributeTypeNumeric,
	"string":  BulkActivateAttributeDetailAttributeTypeString,
}

// GetBulkActivateAttributeDetailAttributeTypeEnumValues Enumerates the set of values for BulkActivateAttributeDetailAttributeTypeEnum
func GetBulkActivateAttributeDetailAttributeTypeEnumValues() []BulkActivateAttributeDetailAttributeTypeEnum {
	values := make([]BulkActivateAttributeDetailAttributeTypeEnum, 0)
	for _, v := range mappingBulkActivateAttributeDetailAttributeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBulkActivateAttributeDetailAttributeTypeEnumStringValues Enumerates the set of values in String for BulkActivateAttributeDetailAttributeTypeEnum
func GetBulkActivateAttributeDetailAttributeTypeEnumStringValues() []string {
	return []string{
		"NUMERIC",
		"STRING",
	}
}

// GetMappingBulkActivateAttributeDetailAttributeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBulkActivateAttributeDetailAttributeTypeEnum(val string) (BulkActivateAttributeDetailAttributeTypeEnum, bool) {
	enum, ok := mappingBulkActivateAttributeDetailAttributeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BulkActivateAttributeDetailUnitEnum Enum with underlying type: string
type BulkActivateAttributeDetailUnitEnum string

// Set of constants representing the allowable values for BulkActivateAttributeDetailUnitEnum
const (
	BulkActivateAttributeDetailUnitNone        BulkActivateAttributeDetailUnitEnum = "NONE"
	BulkActivateAttributeDetailUnitEpochTimeMs BulkActivateAttributeDetailUnitEnum = "EPOCH_TIME_MS"
	BulkActivateAttributeDetailUnitBytes       BulkActivateAttributeDetailUnitEnum = "BYTES"
	BulkActivateAttributeDetailUnitCount       BulkActivateAttributeDetailUnitEnum = "COUNT"
	BulkActivateAttributeDetailUnitDurationMs  BulkActivateAttributeDetailUnitEnum = "DURATION_MS"
	BulkActivateAttributeDetailUnitTraceStatus BulkActivateAttributeDetailUnitEnum = "TRACE_STATUS"
	BulkActivateAttributeDetailUnitPercentage  BulkActivateAttributeDetailUnitEnum = "PERCENTAGE"
)

var mappingBulkActivateAttributeDetailUnitEnum = map[string]BulkActivateAttributeDetailUnitEnum{
	"NONE":          BulkActivateAttributeDetailUnitNone,
	"EPOCH_TIME_MS": BulkActivateAttributeDetailUnitEpochTimeMs,
	"BYTES":         BulkActivateAttributeDetailUnitBytes,
	"COUNT":         BulkActivateAttributeDetailUnitCount,
	"DURATION_MS":   BulkActivateAttributeDetailUnitDurationMs,
	"TRACE_STATUS":  BulkActivateAttributeDetailUnitTraceStatus,
	"PERCENTAGE":    BulkActivateAttributeDetailUnitPercentage,
}

var mappingBulkActivateAttributeDetailUnitEnumLowerCase = map[string]BulkActivateAttributeDetailUnitEnum{
	"none":          BulkActivateAttributeDetailUnitNone,
	"epoch_time_ms": BulkActivateAttributeDetailUnitEpochTimeMs,
	"bytes":         BulkActivateAttributeDetailUnitBytes,
	"count":         BulkActivateAttributeDetailUnitCount,
	"duration_ms":   BulkActivateAttributeDetailUnitDurationMs,
	"trace_status":  BulkActivateAttributeDetailUnitTraceStatus,
	"percentage":    BulkActivateAttributeDetailUnitPercentage,
}

// GetBulkActivateAttributeDetailUnitEnumValues Enumerates the set of values for BulkActivateAttributeDetailUnitEnum
func GetBulkActivateAttributeDetailUnitEnumValues() []BulkActivateAttributeDetailUnitEnum {
	values := make([]BulkActivateAttributeDetailUnitEnum, 0)
	for _, v := range mappingBulkActivateAttributeDetailUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetBulkActivateAttributeDetailUnitEnumStringValues Enumerates the set of values in String for BulkActivateAttributeDetailUnitEnum
func GetBulkActivateAttributeDetailUnitEnumStringValues() []string {
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

// GetMappingBulkActivateAttributeDetailUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBulkActivateAttributeDetailUnitEnum(val string) (BulkActivateAttributeDetailUnitEnum, bool) {
	enum, ok := mappingBulkActivateAttributeDetailUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BulkActivateAttributeDetailAttributeNameSpaceEnum Enum with underlying type: string
type BulkActivateAttributeDetailAttributeNameSpaceEnum string

// Set of constants representing the allowable values for BulkActivateAttributeDetailAttributeNameSpaceEnum
const (
	BulkActivateAttributeDetailAttributeNameSpaceTraces    BulkActivateAttributeDetailAttributeNameSpaceEnum = "TRACES"
	BulkActivateAttributeDetailAttributeNameSpaceSynthetic BulkActivateAttributeDetailAttributeNameSpaceEnum = "SYNTHETIC"
)

var mappingBulkActivateAttributeDetailAttributeNameSpaceEnum = map[string]BulkActivateAttributeDetailAttributeNameSpaceEnum{
	"TRACES":    BulkActivateAttributeDetailAttributeNameSpaceTraces,
	"SYNTHETIC": BulkActivateAttributeDetailAttributeNameSpaceSynthetic,
}

var mappingBulkActivateAttributeDetailAttributeNameSpaceEnumLowerCase = map[string]BulkActivateAttributeDetailAttributeNameSpaceEnum{
	"traces":    BulkActivateAttributeDetailAttributeNameSpaceTraces,
	"synthetic": BulkActivateAttributeDetailAttributeNameSpaceSynthetic,
}

// GetBulkActivateAttributeDetailAttributeNameSpaceEnumValues Enumerates the set of values for BulkActivateAttributeDetailAttributeNameSpaceEnum
func GetBulkActivateAttributeDetailAttributeNameSpaceEnumValues() []BulkActivateAttributeDetailAttributeNameSpaceEnum {
	values := make([]BulkActivateAttributeDetailAttributeNameSpaceEnum, 0)
	for _, v := range mappingBulkActivateAttributeDetailAttributeNameSpaceEnum {
		values = append(values, v)
	}
	return values
}

// GetBulkActivateAttributeDetailAttributeNameSpaceEnumStringValues Enumerates the set of values in String for BulkActivateAttributeDetailAttributeNameSpaceEnum
func GetBulkActivateAttributeDetailAttributeNameSpaceEnumStringValues() []string {
	return []string{
		"TRACES",
		"SYNTHETIC",
	}
}

// GetMappingBulkActivateAttributeDetailAttributeNameSpaceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBulkActivateAttributeDetailAttributeNameSpaceEnum(val string) (BulkActivateAttributeDetailAttributeNameSpaceEnum, bool) {
	enum, ok := mappingBulkActivateAttributeDetailAttributeNameSpaceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
