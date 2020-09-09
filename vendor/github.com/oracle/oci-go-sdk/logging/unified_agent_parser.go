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

// UnifiedAgentParser source parser object.
type UnifiedAgentParser interface {

	// Specify time field for event time. If the event doesn't have this field, current time is used.
	GetFieldTimeKey() *string

	// Specify types for converting field into other type.
	GetTypes() map[string]string

	// Specify null value pattern
	GetNullValuePattern() *string

	// If true, empty string field is replaced with nil
	GetIsNullEmptyString() *bool

	// If true, use Fluent::EventTime.now(current time) as a timestamp when time_key is specified
	GetIsEstimateCurrentEvent() *bool

	// If true, keep time field in the record.
	GetIsKeepTimeKey() *bool

	// Specify timeout for parse processing. This is mainly for detecting wrong regexp pattern.
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
		return *m, nil
	}
}

//GetFieldTimeKey returns FieldTimeKey
func (m unifiedagentparser) GetFieldTimeKey() *string {
	return m.FieldTimeKey
}

//GetTypes returns Types
func (m unifiedagentparser) GetTypes() map[string]string {
	return m.Types
}

//GetNullValuePattern returns NullValuePattern
func (m unifiedagentparser) GetNullValuePattern() *string {
	return m.NullValuePattern
}

//GetIsNullEmptyString returns IsNullEmptyString
func (m unifiedagentparser) GetIsNullEmptyString() *bool {
	return m.IsNullEmptyString
}

//GetIsEstimateCurrentEvent returns IsEstimateCurrentEvent
func (m unifiedagentparser) GetIsEstimateCurrentEvent() *bool {
	return m.IsEstimateCurrentEvent
}

//GetIsKeepTimeKey returns IsKeepTimeKey
func (m unifiedagentparser) GetIsKeepTimeKey() *bool {
	return m.IsKeepTimeKey
}

//GetTimeoutInMilliseconds returns TimeoutInMilliseconds
func (m unifiedagentparser) GetTimeoutInMilliseconds() *int {
	return m.TimeoutInMilliseconds
}

func (m unifiedagentparser) String() string {
	return common.PointerString(m)
}

// UnifiedAgentParserParserTypeEnum Enum with underlying type: string
type UnifiedAgentParserParserTypeEnum string

// Set of constants representing the allowable values for UnifiedAgentParserParserTypeEnum
const (
	UnifiedAgentParserParserTypeAuditd        UnifiedAgentParserParserTypeEnum = "AUDITD"
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

var mappingUnifiedAgentParserParserType = map[string]UnifiedAgentParserParserTypeEnum{
	"AUDITD":         UnifiedAgentParserParserTypeAuditd,
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

// GetUnifiedAgentParserParserTypeEnumValues Enumerates the set of values for UnifiedAgentParserParserTypeEnum
func GetUnifiedAgentParserParserTypeEnumValues() []UnifiedAgentParserParserTypeEnum {
	values := make([]UnifiedAgentParserParserTypeEnum, 0)
	for _, v := range mappingUnifiedAgentParserParserType {
		values = append(values, v)
	}
	return values
}
