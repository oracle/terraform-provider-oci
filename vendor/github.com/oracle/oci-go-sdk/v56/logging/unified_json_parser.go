// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, and delete log groups, log objects, and agent configurations.
//

package logging

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// UnifiedJsonParser JSON parser.
type UnifiedJsonParser struct {

	// Specify time field for the event time. If the event doesn't have this field, the current time is used.
	FieldTimeKey *string `mandatory:"false" json:"fieldTimeKey"`

	// Specify types for converting a field into another type.
	Types map[string]string `mandatory:"false" json:"types"`

	// Specify the null value pattern.
	NullValuePattern *string `mandatory:"false" json:"nullValuePattern"`

	// If true, an empty string field is replaced with nil.
	IsNullEmptyString *bool `mandatory:"false" json:"isNullEmptyString"`

	// If true, use Fluent::EventTime.now(current time) as a timestamp when time_key is specified.
	IsEstimateCurrentEvent *bool `mandatory:"false" json:"isEstimateCurrentEvent"`

	// If true, keep time field in the record.
	IsKeepTimeKey *bool `mandatory:"false" json:"isKeepTimeKey"`

	// Specify the timeout for parse processing. This is mainly for detecting an incorrect regexp pattern.
	TimeoutInMilliseconds *int `mandatory:"false" json:"timeoutInMilliseconds"`

	TimeFormat *string `mandatory:"false" json:"timeFormat"`

	TimeType UnifiedJsonParserTimeTypeEnum `mandatory:"false" json:"timeType,omitempty"`
}

//GetFieldTimeKey returns FieldTimeKey
func (m UnifiedJsonParser) GetFieldTimeKey() *string {
	return m.FieldTimeKey
}

//GetTypes returns Types
func (m UnifiedJsonParser) GetTypes() map[string]string {
	return m.Types
}

//GetNullValuePattern returns NullValuePattern
func (m UnifiedJsonParser) GetNullValuePattern() *string {
	return m.NullValuePattern
}

//GetIsNullEmptyString returns IsNullEmptyString
func (m UnifiedJsonParser) GetIsNullEmptyString() *bool {
	return m.IsNullEmptyString
}

//GetIsEstimateCurrentEvent returns IsEstimateCurrentEvent
func (m UnifiedJsonParser) GetIsEstimateCurrentEvent() *bool {
	return m.IsEstimateCurrentEvent
}

//GetIsKeepTimeKey returns IsKeepTimeKey
func (m UnifiedJsonParser) GetIsKeepTimeKey() *bool {
	return m.IsKeepTimeKey
}

//GetTimeoutInMilliseconds returns TimeoutInMilliseconds
func (m UnifiedJsonParser) GetTimeoutInMilliseconds() *int {
	return m.TimeoutInMilliseconds
}

func (m UnifiedJsonParser) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UnifiedJsonParser) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUnifiedJsonParser UnifiedJsonParser
	s := struct {
		DiscriminatorParam string `json:"parserType"`
		MarshalTypeUnifiedJsonParser
	}{
		"JSON",
		(MarshalTypeUnifiedJsonParser)(m),
	}

	return json.Marshal(&s)
}

// UnifiedJsonParserTimeTypeEnum Enum with underlying type: string
type UnifiedJsonParserTimeTypeEnum string

// Set of constants representing the allowable values for UnifiedJsonParserTimeTypeEnum
const (
	UnifiedJsonParserTimeTypeFloat    UnifiedJsonParserTimeTypeEnum = "FLOAT"
	UnifiedJsonParserTimeTypeUnixtime UnifiedJsonParserTimeTypeEnum = "UNIXTIME"
	UnifiedJsonParserTimeTypeString   UnifiedJsonParserTimeTypeEnum = "STRING"
)

var mappingUnifiedJsonParserTimeType = map[string]UnifiedJsonParserTimeTypeEnum{
	"FLOAT":    UnifiedJsonParserTimeTypeFloat,
	"UNIXTIME": UnifiedJsonParserTimeTypeUnixtime,
	"STRING":   UnifiedJsonParserTimeTypeString,
}

// GetUnifiedJsonParserTimeTypeEnumValues Enumerates the set of values for UnifiedJsonParserTimeTypeEnum
func GetUnifiedJsonParserTimeTypeEnumValues() []UnifiedJsonParserTimeTypeEnum {
	values := make([]UnifiedJsonParserTimeTypeEnum, 0)
	for _, v := range mappingUnifiedJsonParserTimeType {
		values = append(values, v)
	}
	return values
}
