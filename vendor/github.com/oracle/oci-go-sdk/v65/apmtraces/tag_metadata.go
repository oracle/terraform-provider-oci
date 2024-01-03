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

// TagMetadata Definition of the tag metadata.
type TagMetadata struct {

	// Type associated with the tag key.
	TagType TagMetadataTagTypeEnum `mandatory:"false" json:"tagType,omitempty"`

	// Unit associated with the tag key.
	TagUnit TagMetadataTagUnitEnum `mandatory:"false" json:"tagUnit,omitempty"`
}

func (m TagMetadata) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TagMetadata) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingTagMetadataTagTypeEnum(string(m.TagType)); !ok && m.TagType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TagType: %s. Supported values are: %s.", m.TagType, strings.Join(GetTagMetadataTagTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTagMetadataTagUnitEnum(string(m.TagUnit)); !ok && m.TagUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TagUnit: %s. Supported values are: %s.", m.TagUnit, strings.Join(GetTagMetadataTagUnitEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TagMetadataTagTypeEnum Enum with underlying type: string
type TagMetadataTagTypeEnum string

// Set of constants representing the allowable values for TagMetadataTagTypeEnum
const (
	TagMetadataTagTypeString  TagMetadataTagTypeEnum = "STRING"
	TagMetadataTagTypeNumeric TagMetadataTagTypeEnum = "NUMERIC"
)

var mappingTagMetadataTagTypeEnum = map[string]TagMetadataTagTypeEnum{
	"STRING":  TagMetadataTagTypeString,
	"NUMERIC": TagMetadataTagTypeNumeric,
}

var mappingTagMetadataTagTypeEnumLowerCase = map[string]TagMetadataTagTypeEnum{
	"string":  TagMetadataTagTypeString,
	"numeric": TagMetadataTagTypeNumeric,
}

// GetTagMetadataTagTypeEnumValues Enumerates the set of values for TagMetadataTagTypeEnum
func GetTagMetadataTagTypeEnumValues() []TagMetadataTagTypeEnum {
	values := make([]TagMetadataTagTypeEnum, 0)
	for _, v := range mappingTagMetadataTagTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTagMetadataTagTypeEnumStringValues Enumerates the set of values in String for TagMetadataTagTypeEnum
func GetTagMetadataTagTypeEnumStringValues() []string {
	return []string{
		"STRING",
		"NUMERIC",
	}
}

// GetMappingTagMetadataTagTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTagMetadataTagTypeEnum(val string) (TagMetadataTagTypeEnum, bool) {
	enum, ok := mappingTagMetadataTagTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TagMetadataTagUnitEnum Enum with underlying type: string
type TagMetadataTagUnitEnum string

// Set of constants representing the allowable values for TagMetadataTagUnitEnum
const (
	TagMetadataTagUnitEpochTimeMs TagMetadataTagUnitEnum = "EPOCH_TIME_MS"
	TagMetadataTagUnitBytes       TagMetadataTagUnitEnum = "BYTES"
	TagMetadataTagUnitCount       TagMetadataTagUnitEnum = "COUNT"
	TagMetadataTagUnitDurationMs  TagMetadataTagUnitEnum = "DURATION_MS"
	TagMetadataTagUnitTraceStatus TagMetadataTagUnitEnum = "TRACE_STATUS"
	TagMetadataTagUnitPercentage  TagMetadataTagUnitEnum = "PERCENTAGE"
	TagMetadataTagUnitNone        TagMetadataTagUnitEnum = "NONE"
)

var mappingTagMetadataTagUnitEnum = map[string]TagMetadataTagUnitEnum{
	"EPOCH_TIME_MS": TagMetadataTagUnitEpochTimeMs,
	"BYTES":         TagMetadataTagUnitBytes,
	"COUNT":         TagMetadataTagUnitCount,
	"DURATION_MS":   TagMetadataTagUnitDurationMs,
	"TRACE_STATUS":  TagMetadataTagUnitTraceStatus,
	"PERCENTAGE":    TagMetadataTagUnitPercentage,
	"NONE":          TagMetadataTagUnitNone,
}

var mappingTagMetadataTagUnitEnumLowerCase = map[string]TagMetadataTagUnitEnum{
	"epoch_time_ms": TagMetadataTagUnitEpochTimeMs,
	"bytes":         TagMetadataTagUnitBytes,
	"count":         TagMetadataTagUnitCount,
	"duration_ms":   TagMetadataTagUnitDurationMs,
	"trace_status":  TagMetadataTagUnitTraceStatus,
	"percentage":    TagMetadataTagUnitPercentage,
	"none":          TagMetadataTagUnitNone,
}

// GetTagMetadataTagUnitEnumValues Enumerates the set of values for TagMetadataTagUnitEnum
func GetTagMetadataTagUnitEnumValues() []TagMetadataTagUnitEnum {
	values := make([]TagMetadataTagUnitEnum, 0)
	for _, v := range mappingTagMetadataTagUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetTagMetadataTagUnitEnumStringValues Enumerates the set of values in String for TagMetadataTagUnitEnum
func GetTagMetadataTagUnitEnumStringValues() []string {
	return []string{
		"EPOCH_TIME_MS",
		"BYTES",
		"COUNT",
		"DURATION_MS",
		"TRACE_STATUS",
		"PERCENTAGE",
		"NONE",
	}
}

// GetMappingTagMetadataTagUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTagMetadataTagUnitEnum(val string) (TagMetadataTagUnitEnum, bool) {
	enum, ok := mappingTagMetadataTagUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
