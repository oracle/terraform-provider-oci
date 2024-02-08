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

// UnifiedAgentParser Source parser object.
type UnifiedAgentParser interface {

	// Specifies the time field for the event time. If the event doesn't have this field, the current time is used.
	GetFieldTimeKey() *string

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
	GetTypes() map[string]string

	// Specify the null value pattern.
	GetNullValuePattern() *string

	// If true, an empty string field is replaced with a null value.
	GetIsNullEmptyString() *bool

	// If true, use Fluent::EventTime.now(current time) as a timestamp when the time_key is specified.
	GetIsEstimateCurrentEvent() *bool

	// If true, keep the time field in the record.
	GetIsKeepTimeKey() *bool

	// Specify the timeout for parse processing. This is mainly for detecting an incorrect regexp pattern.
	GetTimeoutInMilliseconds() *int
}

type unifiedagentparser struct {
	JsonData               []byte
	FieldTimeKey           *string           `mandatory:"false" json:"fieldTimeKey"`
	Types                  map[string]string `mandatory:"false" json:"types"`
	NullValuePattern       *string           `mandatory:"false" json:"nullValuePattern"`
	IsNullEmptyString      *bool             `mandatory:"false" json:"isNullEmptyString"`
	IsEstimateCurrentEvent *bool             `mandatory:"false" json:"isEstimateCurrentEvent"`
	IsKeepTimeKey          *bool             `mandatory:"false" json:"isKeepTimeKey"`
	TimeoutInMilliseconds  *int              `mandatory:"false" json:"timeoutInMilliseconds"`
	ParserType             string            `json:"parserType"`
}

// UnmarshalJSON unmarshals json
func (m *unifiedagentparser) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerunifiedagentparser unifiedagentparser
	s := struct {
		Model Unmarshalerunifiedagentparser
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.FieldTimeKey = s.Model.FieldTimeKey
	m.Types = s.Model.Types
	m.NullValuePattern = s.Model.NullValuePattern
	m.IsNullEmptyString = s.Model.IsNullEmptyString
	m.IsEstimateCurrentEvent = s.Model.IsEstimateCurrentEvent
	m.IsKeepTimeKey = s.Model.IsKeepTimeKey
	m.TimeoutInMilliseconds = s.Model.TimeoutInMilliseconds
	m.ParserType = s.Model.ParserType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *unifiedagentparser) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ParserType {
	case "MULTILINE_GROK":
		mm := UnifiedAgentMultilineGrokParser{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "JSON":
		mm := UnifiedJsonParser{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GROK":
		mm := UnifiedAgentGrokParser{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NONE":
		mm := UnifiedAgentNoneParser{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SYSLOG":
		mm := UnifiedAgentSyslogParser{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AUDITD":
		mm := UnifiedAgentAuditdParser{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "APACHE2":
		mm := UnifiedAgentApache2Parser{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "REGEXP":
		mm := UnifiedAgentRegexParser{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MULTILINE":
		mm := UnifiedAgentMultilineParser{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TSV":
		mm := UnifiedAgentTsvParser{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CRI":
		mm := UnifiedAgentCriParser{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "APACHE_ERROR":
		mm := UnifiedAgentApacheErrorParser{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MSGPACK":
		mm := UnifiedAgentMsgpackParser{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CSV":
		mm := UnifiedAgentCsvParser{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for UnifiedAgentParser: %s.", m.ParserType)
		return *m, nil
	}
}

// GetFieldTimeKey returns FieldTimeKey
func (m unifiedagentparser) GetFieldTimeKey() *string {
	return m.FieldTimeKey
}

// GetTypes returns Types
func (m unifiedagentparser) GetTypes() map[string]string {
	return m.Types
}

// GetNullValuePattern returns NullValuePattern
func (m unifiedagentparser) GetNullValuePattern() *string {
	return m.NullValuePattern
}

// GetIsNullEmptyString returns IsNullEmptyString
func (m unifiedagentparser) GetIsNullEmptyString() *bool {
	return m.IsNullEmptyString
}

// GetIsEstimateCurrentEvent returns IsEstimateCurrentEvent
func (m unifiedagentparser) GetIsEstimateCurrentEvent() *bool {
	return m.IsEstimateCurrentEvent
}

// GetIsKeepTimeKey returns IsKeepTimeKey
func (m unifiedagentparser) GetIsKeepTimeKey() *bool {
	return m.IsKeepTimeKey
}

// GetTimeoutInMilliseconds returns TimeoutInMilliseconds
func (m unifiedagentparser) GetTimeoutInMilliseconds() *int {
	return m.TimeoutInMilliseconds
}

func (m unifiedagentparser) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m unifiedagentparser) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnifiedAgentParserParserTypeEnum Enum with underlying type: string
type UnifiedAgentParserParserTypeEnum string

// Set of constants representing the allowable values for UnifiedAgentParserParserTypeEnum
const (
	UnifiedAgentParserParserTypeAuditd        UnifiedAgentParserParserTypeEnum = "AUDITD"
	UnifiedAgentParserParserTypeCri           UnifiedAgentParserParserTypeEnum = "CRI"
	UnifiedAgentParserParserTypeJson          UnifiedAgentParserParserTypeEnum = "JSON"
	UnifiedAgentParserParserTypeTsv           UnifiedAgentParserParserTypeEnum = "TSV"
	UnifiedAgentParserParserTypeCsv           UnifiedAgentParserParserTypeEnum = "CSV"
	UnifiedAgentParserParserTypeNone          UnifiedAgentParserParserTypeEnum = "NONE"
	UnifiedAgentParserParserTypeSyslog        UnifiedAgentParserParserTypeEnum = "SYSLOG"
	UnifiedAgentParserParserTypeApache2       UnifiedAgentParserParserTypeEnum = "APACHE2"
	UnifiedAgentParserParserTypeApacheError   UnifiedAgentParserParserTypeEnum = "APACHE_ERROR"
	UnifiedAgentParserParserTypeMsgpack       UnifiedAgentParserParserTypeEnum = "MSGPACK"
	UnifiedAgentParserParserTypeRegexp        UnifiedAgentParserParserTypeEnum = "REGEXP"
	UnifiedAgentParserParserTypeMultiline     UnifiedAgentParserParserTypeEnum = "MULTILINE"
	UnifiedAgentParserParserTypeGrok          UnifiedAgentParserParserTypeEnum = "GROK"
	UnifiedAgentParserParserTypeMultilineGrok UnifiedAgentParserParserTypeEnum = "MULTILINE_GROK"
)

var mappingUnifiedAgentParserParserTypeEnum = map[string]UnifiedAgentParserParserTypeEnum{
	"AUDITD":         UnifiedAgentParserParserTypeAuditd,
	"CRI":            UnifiedAgentParserParserTypeCri,
	"JSON":           UnifiedAgentParserParserTypeJson,
	"TSV":            UnifiedAgentParserParserTypeTsv,
	"CSV":            UnifiedAgentParserParserTypeCsv,
	"NONE":           UnifiedAgentParserParserTypeNone,
	"SYSLOG":         UnifiedAgentParserParserTypeSyslog,
	"APACHE2":        UnifiedAgentParserParserTypeApache2,
	"APACHE_ERROR":   UnifiedAgentParserParserTypeApacheError,
	"MSGPACK":        UnifiedAgentParserParserTypeMsgpack,
	"REGEXP":         UnifiedAgentParserParserTypeRegexp,
	"MULTILINE":      UnifiedAgentParserParserTypeMultiline,
	"GROK":           UnifiedAgentParserParserTypeGrok,
	"MULTILINE_GROK": UnifiedAgentParserParserTypeMultilineGrok,
}

var mappingUnifiedAgentParserParserTypeEnumLowerCase = map[string]UnifiedAgentParserParserTypeEnum{
	"auditd":         UnifiedAgentParserParserTypeAuditd,
	"cri":            UnifiedAgentParserParserTypeCri,
	"json":           UnifiedAgentParserParserTypeJson,
	"tsv":            UnifiedAgentParserParserTypeTsv,
	"csv":            UnifiedAgentParserParserTypeCsv,
	"none":           UnifiedAgentParserParserTypeNone,
	"syslog":         UnifiedAgentParserParserTypeSyslog,
	"apache2":        UnifiedAgentParserParserTypeApache2,
	"apache_error":   UnifiedAgentParserParserTypeApacheError,
	"msgpack":        UnifiedAgentParserParserTypeMsgpack,
	"regexp":         UnifiedAgentParserParserTypeRegexp,
	"multiline":      UnifiedAgentParserParserTypeMultiline,
	"grok":           UnifiedAgentParserParserTypeGrok,
	"multiline_grok": UnifiedAgentParserParserTypeMultilineGrok,
}

// GetUnifiedAgentParserParserTypeEnumValues Enumerates the set of values for UnifiedAgentParserParserTypeEnum
func GetUnifiedAgentParserParserTypeEnumValues() []UnifiedAgentParserParserTypeEnum {
	values := make([]UnifiedAgentParserParserTypeEnum, 0)
	for _, v := range mappingUnifiedAgentParserParserTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUnifiedAgentParserParserTypeEnumStringValues Enumerates the set of values in String for UnifiedAgentParserParserTypeEnum
func GetUnifiedAgentParserParserTypeEnumStringValues() []string {
	return []string{
		"AUDITD",
		"CRI",
		"JSON",
		"TSV",
		"CSV",
		"NONE",
		"SYSLOG",
		"APACHE2",
		"APACHE_ERROR",
		"MSGPACK",
		"REGEXP",
		"MULTILINE",
		"GROK",
		"MULTILINE_GROK",
	}
}

// GetMappingUnifiedAgentParserParserTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUnifiedAgentParserParserTypeEnum(val string) (UnifiedAgentParserParserTypeEnum, bool) {
	enum, ok := mappingUnifiedAgentParserParserTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
