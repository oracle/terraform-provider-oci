// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// AttributeMetadata Definition of the attribute metadata.
type AttributeMetadata struct {

	// Type associated with the attribute key.
	AttributeType AttributeMetadataAttributeTypeEnum `mandatory:"false" json:"attributeType,omitempty"`

	// Unit associated with the attribute key.  If unit is not specified, it defaults to NONE.
	AttributeUnit AttributeMetadataAttributeUnitEnum `mandatory:"false" json:"attributeUnit,omitempty"`
}

func (m AttributeMetadata) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AttributeMetadata) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAttributeMetadataAttributeTypeEnum(string(m.AttributeType)); !ok && m.AttributeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AttributeType: %s. Supported values are: %s.", m.AttributeType, strings.Join(GetAttributeMetadataAttributeTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAttributeMetadataAttributeUnitEnum(string(m.AttributeUnit)); !ok && m.AttributeUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AttributeUnit: %s. Supported values are: %s.", m.AttributeUnit, strings.Join(GetAttributeMetadataAttributeUnitEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AttributeMetadataAttributeTypeEnum Enum with underlying type: string
type AttributeMetadataAttributeTypeEnum string

// Set of constants representing the allowable values for AttributeMetadataAttributeTypeEnum
const (
	AttributeMetadataAttributeTypeString  AttributeMetadataAttributeTypeEnum = "STRING"
	AttributeMetadataAttributeTypeNumeric AttributeMetadataAttributeTypeEnum = "NUMERIC"
)

var mappingAttributeMetadataAttributeTypeEnum = map[string]AttributeMetadataAttributeTypeEnum{
	"STRING":  AttributeMetadataAttributeTypeString,
	"NUMERIC": AttributeMetadataAttributeTypeNumeric,
}

var mappingAttributeMetadataAttributeTypeEnumLowerCase = map[string]AttributeMetadataAttributeTypeEnum{
	"string":  AttributeMetadataAttributeTypeString,
	"numeric": AttributeMetadataAttributeTypeNumeric,
}

// GetAttributeMetadataAttributeTypeEnumValues Enumerates the set of values for AttributeMetadataAttributeTypeEnum
func GetAttributeMetadataAttributeTypeEnumValues() []AttributeMetadataAttributeTypeEnum {
	values := make([]AttributeMetadataAttributeTypeEnum, 0)
	for _, v := range mappingAttributeMetadataAttributeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAttributeMetadataAttributeTypeEnumStringValues Enumerates the set of values in String for AttributeMetadataAttributeTypeEnum
func GetAttributeMetadataAttributeTypeEnumStringValues() []string {
	return []string{
		"STRING",
		"NUMERIC",
	}
}

// GetMappingAttributeMetadataAttributeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAttributeMetadataAttributeTypeEnum(val string) (AttributeMetadataAttributeTypeEnum, bool) {
	enum, ok := mappingAttributeMetadataAttributeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AttributeMetadataAttributeUnitEnum Enum with underlying type: string
type AttributeMetadataAttributeUnitEnum string

// Set of constants representing the allowable values for AttributeMetadataAttributeUnitEnum
const (
	AttributeMetadataAttributeUnitEpochTimeMs AttributeMetadataAttributeUnitEnum = "EPOCH_TIME_MS"
	AttributeMetadataAttributeUnitBytes       AttributeMetadataAttributeUnitEnum = "BYTES"
	AttributeMetadataAttributeUnitCount       AttributeMetadataAttributeUnitEnum = "COUNT"
	AttributeMetadataAttributeUnitDurationMs  AttributeMetadataAttributeUnitEnum = "DURATION_MS"
	AttributeMetadataAttributeUnitPercentage  AttributeMetadataAttributeUnitEnum = "PERCENTAGE"
	AttributeMetadataAttributeUnitNone        AttributeMetadataAttributeUnitEnum = "NONE"
)

var mappingAttributeMetadataAttributeUnitEnum = map[string]AttributeMetadataAttributeUnitEnum{
	"EPOCH_TIME_MS": AttributeMetadataAttributeUnitEpochTimeMs,
	"BYTES":         AttributeMetadataAttributeUnitBytes,
	"COUNT":         AttributeMetadataAttributeUnitCount,
	"DURATION_MS":   AttributeMetadataAttributeUnitDurationMs,
	"PERCENTAGE":    AttributeMetadataAttributeUnitPercentage,
	"NONE":          AttributeMetadataAttributeUnitNone,
}

var mappingAttributeMetadataAttributeUnitEnumLowerCase = map[string]AttributeMetadataAttributeUnitEnum{
	"epoch_time_ms": AttributeMetadataAttributeUnitEpochTimeMs,
	"bytes":         AttributeMetadataAttributeUnitBytes,
	"count":         AttributeMetadataAttributeUnitCount,
	"duration_ms":   AttributeMetadataAttributeUnitDurationMs,
	"percentage":    AttributeMetadataAttributeUnitPercentage,
	"none":          AttributeMetadataAttributeUnitNone,
}

// GetAttributeMetadataAttributeUnitEnumValues Enumerates the set of values for AttributeMetadataAttributeUnitEnum
func GetAttributeMetadataAttributeUnitEnumValues() []AttributeMetadataAttributeUnitEnum {
	values := make([]AttributeMetadataAttributeUnitEnum, 0)
	for _, v := range mappingAttributeMetadataAttributeUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetAttributeMetadataAttributeUnitEnumStringValues Enumerates the set of values in String for AttributeMetadataAttributeUnitEnum
func GetAttributeMetadataAttributeUnitEnumStringValues() []string {
	return []string{
		"EPOCH_TIME_MS",
		"BYTES",
		"COUNT",
		"DURATION_MS",
		"PERCENTAGE",
		"NONE",
	}
}

// GetMappingAttributeMetadataAttributeUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAttributeMetadataAttributeUnitEnum(val string) (AttributeMetadataAttributeUnitEnum, bool) {
	enum, ok := mappingAttributeMetadataAttributeUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
