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

// UnifiedAgentLoggingFilter Logging filter object.
type UnifiedAgentLoggingFilter interface {

	// Unique name for the filter.
	GetName() *string
}

type unifiedagentloggingfilter struct {
	JsonData   []byte
	Name       *string `mandatory:"true" json:"name"`
	FilterType string  `json:"filterType"`
}

// UnmarshalJSON unmarshals json
func (m *unifiedagentloggingfilter) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerunifiedagentloggingfilter unifiedagentloggingfilter
	s := struct {
		Model Unmarshalerunifiedagentloggingfilter
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.FilterType = s.Model.FilterType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *unifiedagentloggingfilter) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.FilterType {
	case "CUSTOM_FILTER":
		mm := UnifiedAgentCustomFilter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PARSER_FILTER":
		mm := UnifiedAgentParserFilter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GREP_FILTER":
		mm := UnifiedAgentLoggingGrepFilter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RECORD_TRANSFORMER_FILTER":
		mm := UnifiedAgentLoggingRecordTransformerFilter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for UnifiedAgentLoggingFilter: %s.", m.FilterType)
		return *m, nil
	}
}

// GetName returns Name
func (m unifiedagentloggingfilter) GetName() *string {
	return m.Name
}

func (m unifiedagentloggingfilter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m unifiedagentloggingfilter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnifiedAgentLoggingFilterFilterTypeEnum Enum with underlying type: string
type UnifiedAgentLoggingFilterFilterTypeEnum string

// Set of constants representing the allowable values for UnifiedAgentLoggingFilterFilterTypeEnum
const (
	UnifiedAgentLoggingFilterFilterTypeParserFilter            UnifiedAgentLoggingFilterFilterTypeEnum = "PARSER_FILTER"
	UnifiedAgentLoggingFilterFilterTypeGrepFilter              UnifiedAgentLoggingFilterFilterTypeEnum = "GREP_FILTER"
	UnifiedAgentLoggingFilterFilterTypeRecordTransformerFilter UnifiedAgentLoggingFilterFilterTypeEnum = "RECORD_TRANSFORMER_FILTER"
	UnifiedAgentLoggingFilterFilterTypeCustomFilter            UnifiedAgentLoggingFilterFilterTypeEnum = "CUSTOM_FILTER"
)

var mappingUnifiedAgentLoggingFilterFilterTypeEnum = map[string]UnifiedAgentLoggingFilterFilterTypeEnum{
	"PARSER_FILTER":             UnifiedAgentLoggingFilterFilterTypeParserFilter,
	"GREP_FILTER":               UnifiedAgentLoggingFilterFilterTypeGrepFilter,
	"RECORD_TRANSFORMER_FILTER": UnifiedAgentLoggingFilterFilterTypeRecordTransformerFilter,
	"CUSTOM_FILTER":             UnifiedAgentLoggingFilterFilterTypeCustomFilter,
}

var mappingUnifiedAgentLoggingFilterFilterTypeEnumLowerCase = map[string]UnifiedAgentLoggingFilterFilterTypeEnum{
	"parser_filter":             UnifiedAgentLoggingFilterFilterTypeParserFilter,
	"grep_filter":               UnifiedAgentLoggingFilterFilterTypeGrepFilter,
	"record_transformer_filter": UnifiedAgentLoggingFilterFilterTypeRecordTransformerFilter,
	"custom_filter":             UnifiedAgentLoggingFilterFilterTypeCustomFilter,
}

// GetUnifiedAgentLoggingFilterFilterTypeEnumValues Enumerates the set of values for UnifiedAgentLoggingFilterFilterTypeEnum
func GetUnifiedAgentLoggingFilterFilterTypeEnumValues() []UnifiedAgentLoggingFilterFilterTypeEnum {
	values := make([]UnifiedAgentLoggingFilterFilterTypeEnum, 0)
	for _, v := range mappingUnifiedAgentLoggingFilterFilterTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUnifiedAgentLoggingFilterFilterTypeEnumStringValues Enumerates the set of values in String for UnifiedAgentLoggingFilterFilterTypeEnum
func GetUnifiedAgentLoggingFilterFilterTypeEnumStringValues() []string {
	return []string{
		"PARSER_FILTER",
		"GREP_FILTER",
		"RECORD_TRANSFORMER_FILTER",
		"CUSTOM_FILTER",
	}
}

// GetMappingUnifiedAgentLoggingFilterFilterTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUnifiedAgentLoggingFilterFilterTypeEnum(val string) (UnifiedAgentLoggingFilterFilterTypeEnum, bool) {
	enum, ok := mappingUnifiedAgentLoggingFilterFilterTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
