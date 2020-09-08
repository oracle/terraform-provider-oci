// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// loggingManagementControlplane API
//
// loggingManagementControlplane API specification
//

package logging

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// UnifiedAgentSyslogParser Syslog Parser
type UnifiedAgentSyslogParser struct {

	// Specify time field for event time. If the event doesn't have this field, current time is used.
	FieldTimeKey *string `mandatory:"false" json:"fieldTimeKey"`

	// Specify types for converting field into other type.
	Types map[string]string `mandatory:"false" json:"types"`

	// Specify null value pattern
	NullValuePattern *string `mandatory:"false" json:"nullValuePattern"`

	// If true, empty string field is replaced with nil
	IsNullEmptyString *bool `mandatory:"false" json:"isNullEmptyString"`

	// If true, use Fluent::EventTime.now(current time) as a timestamp when time_key is specified
	IsEstimateCurrentEvent *bool `mandatory:"false" json:"isEstimateCurrentEvent"`

	// If true, keep time field in the record.
	IsKeepTimeKey *bool `mandatory:"false" json:"isKeepTimeKey"`

	// Specify timeout for parse processing. This is mainly for detecting wrong regexp pattern.
	TimeoutInMilliseconds *int `mandatory:"false" json:"timeoutInMilliseconds"`

	TimeFormat *string `mandatory:"false" json:"timeFormat"`

	Rfc5424TimeFormat *string `mandatory:"false" json:"rfc5424TimeFormat"`

	IsWithPriority *bool `mandatory:"false" json:"isWithPriority"`

	IsSupportColonlessIdent *bool `mandatory:"false" json:"isSupportColonlessIdent"`

	MessageFormat UnifiedAgentSyslogParserMessageFormatEnum `mandatory:"false" json:"messageFormat,omitempty"`

	SyslogParserType UnifiedAgentSyslogParserSyslogParserTypeEnum `mandatory:"false" json:"syslogParserType,omitempty"`
}

//GetFieldTimeKey returns FieldTimeKey
func (m UnifiedAgentSyslogParser) GetFieldTimeKey() *string {
	return m.FieldTimeKey
}

//GetTypes returns Types
func (m UnifiedAgentSyslogParser) GetTypes() map[string]string {
	return m.Types
}

//GetNullValuePattern returns NullValuePattern
func (m UnifiedAgentSyslogParser) GetNullValuePattern() *string {
	return m.NullValuePattern
}

//GetIsNullEmptyString returns IsNullEmptyString
func (m UnifiedAgentSyslogParser) GetIsNullEmptyString() *bool {
	return m.IsNullEmptyString
}

//GetIsEstimateCurrentEvent returns IsEstimateCurrentEvent
func (m UnifiedAgentSyslogParser) GetIsEstimateCurrentEvent() *bool {
	return m.IsEstimateCurrentEvent
}

//GetIsKeepTimeKey returns IsKeepTimeKey
func (m UnifiedAgentSyslogParser) GetIsKeepTimeKey() *bool {
	return m.IsKeepTimeKey
}

//GetTimeoutInMilliseconds returns TimeoutInMilliseconds
func (m UnifiedAgentSyslogParser) GetTimeoutInMilliseconds() *int {
	return m.TimeoutInMilliseconds
}

func (m UnifiedAgentSyslogParser) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UnifiedAgentSyslogParser) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUnifiedAgentSyslogParser UnifiedAgentSyslogParser
	s := struct {
		DiscriminatorParam string `json:"parserType"`
		MarshalTypeUnifiedAgentSyslogParser
	}{
		"SYSLOG",
		(MarshalTypeUnifiedAgentSyslogParser)(m),
	}

	return json.Marshal(&s)
}

// UnifiedAgentSyslogParserMessageFormatEnum Enum with underlying type: string
type UnifiedAgentSyslogParserMessageFormatEnum string

// Set of constants representing the allowable values for UnifiedAgentSyslogParserMessageFormatEnum
const (
	UnifiedAgentSyslogParserMessageFormatRfc3164 UnifiedAgentSyslogParserMessageFormatEnum = "RFC3164"
	UnifiedAgentSyslogParserMessageFormatRfc5424 UnifiedAgentSyslogParserMessageFormatEnum = "RFC5424"
	UnifiedAgentSyslogParserMessageFormatAuto    UnifiedAgentSyslogParserMessageFormatEnum = "AUTO"
)

var mappingUnifiedAgentSyslogParserMessageFormat = map[string]UnifiedAgentSyslogParserMessageFormatEnum{
	"RFC3164": UnifiedAgentSyslogParserMessageFormatRfc3164,
	"RFC5424": UnifiedAgentSyslogParserMessageFormatRfc5424,
	"AUTO":    UnifiedAgentSyslogParserMessageFormatAuto,
}

// GetUnifiedAgentSyslogParserMessageFormatEnumValues Enumerates the set of values for UnifiedAgentSyslogParserMessageFormatEnum
func GetUnifiedAgentSyslogParserMessageFormatEnumValues() []UnifiedAgentSyslogParserMessageFormatEnum {
	values := make([]UnifiedAgentSyslogParserMessageFormatEnum, 0)
	for _, v := range mappingUnifiedAgentSyslogParserMessageFormat {
		values = append(values, v)
	}
	return values
}

// UnifiedAgentSyslogParserSyslogParserTypeEnum Enum with underlying type: string
type UnifiedAgentSyslogParserSyslogParserTypeEnum string

// Set of constants representing the allowable values for UnifiedAgentSyslogParserSyslogParserTypeEnum
const (
	UnifiedAgentSyslogParserSyslogParserTypeString UnifiedAgentSyslogParserSyslogParserTypeEnum = "STRING"
	UnifiedAgentSyslogParserSyslogParserTypeRegexp UnifiedAgentSyslogParserSyslogParserTypeEnum = "REGEXP"
)

var mappingUnifiedAgentSyslogParserSyslogParserType = map[string]UnifiedAgentSyslogParserSyslogParserTypeEnum{
	"STRING": UnifiedAgentSyslogParserSyslogParserTypeString,
	"REGEXP": UnifiedAgentSyslogParserSyslogParserTypeRegexp,
}

// GetUnifiedAgentSyslogParserSyslogParserTypeEnumValues Enumerates the set of values for UnifiedAgentSyslogParserSyslogParserTypeEnum
func GetUnifiedAgentSyslogParserSyslogParserTypeEnumValues() []UnifiedAgentSyslogParserSyslogParserTypeEnum {
	values := make([]UnifiedAgentSyslogParserSyslogParserTypeEnum, 0)
	for _, v := range mappingUnifiedAgentSyslogParserSyslogParserType {
		values = append(values, v)
	}
	return values
}
