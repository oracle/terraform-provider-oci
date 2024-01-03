// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, move and delete
// log groups, log objects, log saved searches, and agent configurations.
// For more information, see Logging Overview (https://docs.cloud.oracle.com/iaas/Content/Logging/Concepts/loggingoverview.htm).
//

package logging

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UnifiedAgentSyslogParser Syslog Parser.
type UnifiedAgentSyslogParser struct {

	// Specifies the time field for the event time. If the event doesn't have this field, the current time is used.
	FieldTimeKey *string `mandatory:"false" json:"fieldTimeKey"`

	// Specify types for converting a field into another type.
	// For example,
	//   With this configuration:
	//       <parse>
	//         @type csv
	//         keys time,host,req_id,user
	//         time_key time
	//       </parse>
	//   This incoming event:
	//     "2013/02/28 12:00:00,192.168.0.1,111,-"
	//   is parsed as:
	//     1362020400 (2013/02/28/ 12:00:00)
	//     record:
	//     {
	//       "host"   : "192.168.0.1",
	//       "req_id" : "111",
	//       "user"   : "-"
	//     }
	Types map[string]string `mandatory:"false" json:"types"`

	// Specify the null value pattern.
	NullValuePattern *string `mandatory:"false" json:"nullValuePattern"`

	// If true, an empty string field is replaced with a null value.
	IsNullEmptyString *bool `mandatory:"false" json:"isNullEmptyString"`

	// If true, use Fluent::EventTime.now(current time) as a timestamp when the time_key is specified.
	IsEstimateCurrentEvent *bool `mandatory:"false" json:"isEstimateCurrentEvent"`

	// If true, keep the time field in the record.
	IsKeepTimeKey *bool `mandatory:"false" json:"isKeepTimeKey"`

	// Specify the timeout for parse processing. This is mainly for detecting an incorrect regexp pattern.
	TimeoutInMilliseconds *int `mandatory:"false" json:"timeoutInMilliseconds"`

	// Time format.
	TimeFormat *string `mandatory:"false" json:"timeFormat"`

	// RFC 5424 time format.
	Rfc5424TimeFormat *string `mandatory:"false" json:"rfc5424TimeFormat"`

	// Specifies with priority or not. Corresponds to the Fluentd with_priority parameter.
	IsWithPriority *bool `mandatory:"false" json:"isWithPriority"`

	// Specifies whether or not to support colonless ident. Corresponds to the Fluentd support_colonless_ident parameter.
	IsSupportColonlessIdent *bool `mandatory:"false" json:"isSupportColonlessIdent"`

	// Syslog message format.
	MessageFormat UnifiedAgentSyslogParserMessageFormatEnum `mandatory:"false" json:"messageFormat,omitempty"`

	// Syslog parser type.
	SyslogParserType UnifiedAgentSyslogParserSyslogParserTypeEnum `mandatory:"false" json:"syslogParserType,omitempty"`
}

// GetFieldTimeKey returns FieldTimeKey
func (m UnifiedAgentSyslogParser) GetFieldTimeKey() *string {
	return m.FieldTimeKey
}

// GetTypes returns Types
func (m UnifiedAgentSyslogParser) GetTypes() map[string]string {
	return m.Types
}

// GetNullValuePattern returns NullValuePattern
func (m UnifiedAgentSyslogParser) GetNullValuePattern() *string {
	return m.NullValuePattern
}

// GetIsNullEmptyString returns IsNullEmptyString
func (m UnifiedAgentSyslogParser) GetIsNullEmptyString() *bool {
	return m.IsNullEmptyString
}

// GetIsEstimateCurrentEvent returns IsEstimateCurrentEvent
func (m UnifiedAgentSyslogParser) GetIsEstimateCurrentEvent() *bool {
	return m.IsEstimateCurrentEvent
}

// GetIsKeepTimeKey returns IsKeepTimeKey
func (m UnifiedAgentSyslogParser) GetIsKeepTimeKey() *bool {
	return m.IsKeepTimeKey
}

// GetTimeoutInMilliseconds returns TimeoutInMilliseconds
func (m UnifiedAgentSyslogParser) GetTimeoutInMilliseconds() *int {
	return m.TimeoutInMilliseconds
}

func (m UnifiedAgentSyslogParser) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UnifiedAgentSyslogParser) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUnifiedAgentSyslogParserMessageFormatEnum(string(m.MessageFormat)); !ok && m.MessageFormat != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MessageFormat: %s. Supported values are: %s.", m.MessageFormat, strings.Join(GetUnifiedAgentSyslogParserMessageFormatEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUnifiedAgentSyslogParserSyslogParserTypeEnum(string(m.SyslogParserType)); !ok && m.SyslogParserType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SyslogParserType: %s. Supported values are: %s.", m.SyslogParserType, strings.Join(GetUnifiedAgentSyslogParserSyslogParserTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingUnifiedAgentSyslogParserMessageFormatEnum = map[string]UnifiedAgentSyslogParserMessageFormatEnum{
	"RFC3164": UnifiedAgentSyslogParserMessageFormatRfc3164,
	"RFC5424": UnifiedAgentSyslogParserMessageFormatRfc5424,
	"AUTO":    UnifiedAgentSyslogParserMessageFormatAuto,
}

var mappingUnifiedAgentSyslogParserMessageFormatEnumLowerCase = map[string]UnifiedAgentSyslogParserMessageFormatEnum{
	"rfc3164": UnifiedAgentSyslogParserMessageFormatRfc3164,
	"rfc5424": UnifiedAgentSyslogParserMessageFormatRfc5424,
	"auto":    UnifiedAgentSyslogParserMessageFormatAuto,
}

// GetUnifiedAgentSyslogParserMessageFormatEnumValues Enumerates the set of values for UnifiedAgentSyslogParserMessageFormatEnum
func GetUnifiedAgentSyslogParserMessageFormatEnumValues() []UnifiedAgentSyslogParserMessageFormatEnum {
	values := make([]UnifiedAgentSyslogParserMessageFormatEnum, 0)
	for _, v := range mappingUnifiedAgentSyslogParserMessageFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetUnifiedAgentSyslogParserMessageFormatEnumStringValues Enumerates the set of values in String for UnifiedAgentSyslogParserMessageFormatEnum
func GetUnifiedAgentSyslogParserMessageFormatEnumStringValues() []string {
	return []string{
		"RFC3164",
		"RFC5424",
		"AUTO",
	}
}

// GetMappingUnifiedAgentSyslogParserMessageFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUnifiedAgentSyslogParserMessageFormatEnum(val string) (UnifiedAgentSyslogParserMessageFormatEnum, bool) {
	enum, ok := mappingUnifiedAgentSyslogParserMessageFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UnifiedAgentSyslogParserSyslogParserTypeEnum Enum with underlying type: string
type UnifiedAgentSyslogParserSyslogParserTypeEnum string

// Set of constants representing the allowable values for UnifiedAgentSyslogParserSyslogParserTypeEnum
const (
	UnifiedAgentSyslogParserSyslogParserTypeString UnifiedAgentSyslogParserSyslogParserTypeEnum = "STRING"
	UnifiedAgentSyslogParserSyslogParserTypeRegexp UnifiedAgentSyslogParserSyslogParserTypeEnum = "REGEXP"
)

var mappingUnifiedAgentSyslogParserSyslogParserTypeEnum = map[string]UnifiedAgentSyslogParserSyslogParserTypeEnum{
	"STRING": UnifiedAgentSyslogParserSyslogParserTypeString,
	"REGEXP": UnifiedAgentSyslogParserSyslogParserTypeRegexp,
}

var mappingUnifiedAgentSyslogParserSyslogParserTypeEnumLowerCase = map[string]UnifiedAgentSyslogParserSyslogParserTypeEnum{
	"string": UnifiedAgentSyslogParserSyslogParserTypeString,
	"regexp": UnifiedAgentSyslogParserSyslogParserTypeRegexp,
}

// GetUnifiedAgentSyslogParserSyslogParserTypeEnumValues Enumerates the set of values for UnifiedAgentSyslogParserSyslogParserTypeEnum
func GetUnifiedAgentSyslogParserSyslogParserTypeEnumValues() []UnifiedAgentSyslogParserSyslogParserTypeEnum {
	values := make([]UnifiedAgentSyslogParserSyslogParserTypeEnum, 0)
	for _, v := range mappingUnifiedAgentSyslogParserSyslogParserTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUnifiedAgentSyslogParserSyslogParserTypeEnumStringValues Enumerates the set of values in String for UnifiedAgentSyslogParserSyslogParserTypeEnum
func GetUnifiedAgentSyslogParserSyslogParserTypeEnumStringValues() []string {
	return []string{
		"STRING",
		"REGEXP",
	}
}

// GetMappingUnifiedAgentSyslogParserSyslogParserTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUnifiedAgentSyslogParserSyslogParserTypeEnum(val string) (UnifiedAgentSyslogParserSyslogParserTypeEnum, bool) {
	enum, ok := mappingUnifiedAgentSyslogParserSyslogParserTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
