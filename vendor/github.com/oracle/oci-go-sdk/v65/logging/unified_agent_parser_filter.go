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

// UnifiedAgentParserFilter Logging parser filter object.
// Ref: https://docs.fluentd.org/filter/parser
type UnifiedAgentParserFilter struct {

	// Unique name for the filter.
	Name *string `mandatory:"true" json:"name"`

	Parser UnifiedAgentParser `mandatory:"true" json:"parser"`

	// The field name in the record to parse.
	KeyName *string `mandatory:"true" json:"keyName"`

	// If true, keep the original event time in the parsed result.
	ReserveTime *bool `mandatory:"false" json:"reserveTime"`

	// If true, keep the original key-value pair in the parsed result.
	ReserveData *bool `mandatory:"false" json:"reserveData"`

	// If true, remove the keyName field when parsing is succeeded.
	RemoveKeyNameField *bool `mandatory:"false" json:"removeKeyNameField"`

	// If true, the invalid string is replaced with safe characters and is re-parsed.
	ReplaceInvalidSequence *bool `mandatory:"false" json:"replaceInvalidSequence"`

	// Store the parsed values with the specified key name prefix.
	InjectKeyPrefix *string `mandatory:"false" json:"injectKeyPrefix"`

	// Store the parsed values as a hash value in a field.
	HashValueField *string `mandatory:"false" json:"hashValueField"`

	// If true, emit invalid record to @ERROR label. Invalid cases are: 1) key does not exist; 2) the format
	// does not match; or 3) an unexpected error. You can rescue unexpected format logs in the @ERROR lable.
	// If you want to ignore these errors, set this to false.
	EmitInvalidRecordToError *bool `mandatory:"false" json:"emitInvalidRecordToError"`
}

// GetName returns Name
func (m UnifiedAgentParserFilter) GetName() *string {
	return m.Name
}

func (m UnifiedAgentParserFilter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UnifiedAgentParserFilter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UnifiedAgentParserFilter) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUnifiedAgentParserFilter UnifiedAgentParserFilter
	s := struct {
		DiscriminatorParam string `json:"filterType"`
		MarshalTypeUnifiedAgentParserFilter
	}{
		"PARSER_FILTER",
		(MarshalTypeUnifiedAgentParserFilter)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *UnifiedAgentParserFilter) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ReserveTime              *bool              `json:"reserveTime"`
		ReserveData              *bool              `json:"reserveData"`
		RemoveKeyNameField       *bool              `json:"removeKeyNameField"`
		ReplaceInvalidSequence   *bool              `json:"replaceInvalidSequence"`
		InjectKeyPrefix          *string            `json:"injectKeyPrefix"`
		HashValueField           *string            `json:"hashValueField"`
		EmitInvalidRecordToError *bool              `json:"emitInvalidRecordToError"`
		Name                     *string            `json:"name"`
		Parser                   unifiedagentparser `json:"parser"`
		KeyName                  *string            `json:"keyName"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ReserveTime = model.ReserveTime

	m.ReserveData = model.ReserveData

	m.RemoveKeyNameField = model.RemoveKeyNameField

	m.ReplaceInvalidSequence = model.ReplaceInvalidSequence

	m.InjectKeyPrefix = model.InjectKeyPrefix

	m.HashValueField = model.HashValueField

	m.EmitInvalidRecordToError = model.EmitInvalidRecordToError

	m.Name = model.Name

	nn, e = model.Parser.UnmarshalPolymorphicJSON(model.Parser.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Parser = nn.(UnifiedAgentParser)
	} else {
		m.Parser = nil
	}

	m.KeyName = model.KeyName

	return
}
