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

// UnifiedAgentMultilineGrokParser multiline grok parser
type UnifiedAgentMultilineGrokParser struct {

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

	GrokNameKey *string `mandatory:"false" json:"grokNameKey"`

	GrokFailureKey *string `mandatory:"false" json:"grokFailureKey"`

	MultiLineStartRegexp *string `mandatory:"false" json:"multiLineStartRegexp"`

	Patterns []GrokPattern `mandatory:"false" json:"patterns"`
}

//GetFieldTimeKey returns FieldTimeKey
func (m UnifiedAgentMultilineGrokParser) GetFieldTimeKey() *string {
	return m.FieldTimeKey
}

//GetTypes returns Types
func (m UnifiedAgentMultilineGrokParser) GetTypes() map[string]string {
	return m.Types
}

//GetNullValuePattern returns NullValuePattern
func (m UnifiedAgentMultilineGrokParser) GetNullValuePattern() *string {
	return m.NullValuePattern
}

//GetIsNullEmptyString returns IsNullEmptyString
func (m UnifiedAgentMultilineGrokParser) GetIsNullEmptyString() *bool {
	return m.IsNullEmptyString
}

//GetIsEstimateCurrentEvent returns IsEstimateCurrentEvent
func (m UnifiedAgentMultilineGrokParser) GetIsEstimateCurrentEvent() *bool {
	return m.IsEstimateCurrentEvent
}

//GetIsKeepTimeKey returns IsKeepTimeKey
func (m UnifiedAgentMultilineGrokParser) GetIsKeepTimeKey() *bool {
	return m.IsKeepTimeKey
}

//GetTimeoutInMilliseconds returns TimeoutInMilliseconds
func (m UnifiedAgentMultilineGrokParser) GetTimeoutInMilliseconds() *int {
	return m.TimeoutInMilliseconds
}

func (m UnifiedAgentMultilineGrokParser) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UnifiedAgentMultilineGrokParser) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUnifiedAgentMultilineGrokParser UnifiedAgentMultilineGrokParser
	s := struct {
		DiscriminatorParam string `json:"parserType"`
		MarshalTypeUnifiedAgentMultilineGrokParser
	}{
		"MULTILINE_GROK",
		(MarshalTypeUnifiedAgentMultilineGrokParser)(m),
	}

	return json.Marshal(&s)
}
