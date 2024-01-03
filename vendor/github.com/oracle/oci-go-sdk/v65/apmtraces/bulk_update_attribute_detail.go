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

// BulkUpdateAttributeDetail Object that contains the details about a single attribute in the bulk request for which properties are to be updated.
type BulkUpdateAttributeDetail struct {

	// Name of the attribute for which notes are to be updated.
	AttributeName *string `mandatory:"true" json:"attributeName"`

	// Unit of the attribute to be updated.
	Unit BulkUpdateAttributeDetailUnitEnum `mandatory:"false" json:"unit,omitempty"`

	// Namespace of the attribute for which the properties are to be updated.
	AttributeNameSpace BulkUpdateAttributeDetailAttributeNameSpaceEnum `mandatory:"false" json:"attributeNameSpace,omitempty"`
}

func (m BulkUpdateAttributeDetail) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BulkUpdateAttributeDetail) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingBulkUpdateAttributeDetailUnitEnum(string(m.Unit)); !ok && m.Unit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Unit: %s. Supported values are: %s.", m.Unit, strings.Join(GetBulkUpdateAttributeDetailUnitEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBulkUpdateAttributeDetailAttributeNameSpaceEnum(string(m.AttributeNameSpace)); !ok && m.AttributeNameSpace != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AttributeNameSpace: %s. Supported values are: %s.", m.AttributeNameSpace, strings.Join(GetBulkUpdateAttributeDetailAttributeNameSpaceEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BulkUpdateAttributeDetailUnitEnum Enum with underlying type: string
type BulkUpdateAttributeDetailUnitEnum string

// Set of constants representing the allowable values for BulkUpdateAttributeDetailUnitEnum
const (
	BulkUpdateAttributeDetailUnitNone        BulkUpdateAttributeDetailUnitEnum = "NONE"
	BulkUpdateAttributeDetailUnitEpochTimeMs BulkUpdateAttributeDetailUnitEnum = "EPOCH_TIME_MS"
	BulkUpdateAttributeDetailUnitBytes       BulkUpdateAttributeDetailUnitEnum = "BYTES"
	BulkUpdateAttributeDetailUnitCount       BulkUpdateAttributeDetailUnitEnum = "COUNT"
	BulkUpdateAttributeDetailUnitDurationMs  BulkUpdateAttributeDetailUnitEnum = "DURATION_MS"
	BulkUpdateAttributeDetailUnitTraceStatus BulkUpdateAttributeDetailUnitEnum = "TRACE_STATUS"
	BulkUpdateAttributeDetailUnitPercentage  BulkUpdateAttributeDetailUnitEnum = "PERCENTAGE"
)

var mappingBulkUpdateAttributeDetailUnitEnum = map[string]BulkUpdateAttributeDetailUnitEnum{
	"NONE":          BulkUpdateAttributeDetailUnitNone,
	"EPOCH_TIME_MS": BulkUpdateAttributeDetailUnitEpochTimeMs,
	"BYTES":         BulkUpdateAttributeDetailUnitBytes,
	"COUNT":         BulkUpdateAttributeDetailUnitCount,
	"DURATION_MS":   BulkUpdateAttributeDetailUnitDurationMs,
	"TRACE_STATUS":  BulkUpdateAttributeDetailUnitTraceStatus,
	"PERCENTAGE":    BulkUpdateAttributeDetailUnitPercentage,
}

var mappingBulkUpdateAttributeDetailUnitEnumLowerCase = map[string]BulkUpdateAttributeDetailUnitEnum{
	"none":          BulkUpdateAttributeDetailUnitNone,
	"epoch_time_ms": BulkUpdateAttributeDetailUnitEpochTimeMs,
	"bytes":         BulkUpdateAttributeDetailUnitBytes,
	"count":         BulkUpdateAttributeDetailUnitCount,
	"duration_ms":   BulkUpdateAttributeDetailUnitDurationMs,
	"trace_status":  BulkUpdateAttributeDetailUnitTraceStatus,
	"percentage":    BulkUpdateAttributeDetailUnitPercentage,
}

// GetBulkUpdateAttributeDetailUnitEnumValues Enumerates the set of values for BulkUpdateAttributeDetailUnitEnum
func GetBulkUpdateAttributeDetailUnitEnumValues() []BulkUpdateAttributeDetailUnitEnum {
	values := make([]BulkUpdateAttributeDetailUnitEnum, 0)
	for _, v := range mappingBulkUpdateAttributeDetailUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetBulkUpdateAttributeDetailUnitEnumStringValues Enumerates the set of values in String for BulkUpdateAttributeDetailUnitEnum
func GetBulkUpdateAttributeDetailUnitEnumStringValues() []string {
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

// GetMappingBulkUpdateAttributeDetailUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBulkUpdateAttributeDetailUnitEnum(val string) (BulkUpdateAttributeDetailUnitEnum, bool) {
	enum, ok := mappingBulkUpdateAttributeDetailUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BulkUpdateAttributeDetailAttributeNameSpaceEnum Enum with underlying type: string
type BulkUpdateAttributeDetailAttributeNameSpaceEnum string

// Set of constants representing the allowable values for BulkUpdateAttributeDetailAttributeNameSpaceEnum
const (
	BulkUpdateAttributeDetailAttributeNameSpaceTraces    BulkUpdateAttributeDetailAttributeNameSpaceEnum = "TRACES"
	BulkUpdateAttributeDetailAttributeNameSpaceSynthetic BulkUpdateAttributeDetailAttributeNameSpaceEnum = "SYNTHETIC"
)

var mappingBulkUpdateAttributeDetailAttributeNameSpaceEnum = map[string]BulkUpdateAttributeDetailAttributeNameSpaceEnum{
	"TRACES":    BulkUpdateAttributeDetailAttributeNameSpaceTraces,
	"SYNTHETIC": BulkUpdateAttributeDetailAttributeNameSpaceSynthetic,
}

var mappingBulkUpdateAttributeDetailAttributeNameSpaceEnumLowerCase = map[string]BulkUpdateAttributeDetailAttributeNameSpaceEnum{
	"traces":    BulkUpdateAttributeDetailAttributeNameSpaceTraces,
	"synthetic": BulkUpdateAttributeDetailAttributeNameSpaceSynthetic,
}

// GetBulkUpdateAttributeDetailAttributeNameSpaceEnumValues Enumerates the set of values for BulkUpdateAttributeDetailAttributeNameSpaceEnum
func GetBulkUpdateAttributeDetailAttributeNameSpaceEnumValues() []BulkUpdateAttributeDetailAttributeNameSpaceEnum {
	values := make([]BulkUpdateAttributeDetailAttributeNameSpaceEnum, 0)
	for _, v := range mappingBulkUpdateAttributeDetailAttributeNameSpaceEnum {
		values = append(values, v)
	}
	return values
}

// GetBulkUpdateAttributeDetailAttributeNameSpaceEnumStringValues Enumerates the set of values in String for BulkUpdateAttributeDetailAttributeNameSpaceEnum
func GetBulkUpdateAttributeDetailAttributeNameSpaceEnumStringValues() []string {
	return []string{
		"TRACES",
		"SYNTHETIC",
	}
}

// GetMappingBulkUpdateAttributeDetailAttributeNameSpaceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBulkUpdateAttributeDetailAttributeNameSpaceEnum(val string) (BulkUpdateAttributeDetailAttributeNameSpaceEnum, bool) {
	enum, ok := mappingBulkUpdateAttributeDetailAttributeNameSpaceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
