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

// UnifiedAgentLoggingSource Logging source object.
type UnifiedAgentLoggingSource interface {

	// Unique name for the source.
	GetName() *string
}

type unifiedagentloggingsource struct {
	JsonData   []byte
	Name       *string `mandatory:"true" json:"name"`
	SourceType string  `json:"sourceType"`
}

// UnmarshalJSON unmarshals json
func (m *unifiedagentloggingsource) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerunifiedagentloggingsource unifiedagentloggingsource
	s := struct {
		Model Unmarshalerunifiedagentloggingsource
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.SourceType = s.Model.SourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *unifiedagentloggingsource) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SourceType {
	case "WINDOWS_EVENT_LOG":
		mm := UnifiedAgentWindowsEventSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LOG_TAIL":
		mm := UnifiedAgentTailLogSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for UnifiedAgentLoggingSource: %s.", m.SourceType)
		return *m, nil
	}
}

// GetName returns Name
func (m unifiedagentloggingsource) GetName() *string {
	return m.Name
}

func (m unifiedagentloggingsource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m unifiedagentloggingsource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnifiedAgentLoggingSourceSourceTypeEnum Enum with underlying type: string
type UnifiedAgentLoggingSourceSourceTypeEnum string

// Set of constants representing the allowable values for UnifiedAgentLoggingSourceSourceTypeEnum
const (
	UnifiedAgentLoggingSourceSourceTypeLogTail         UnifiedAgentLoggingSourceSourceTypeEnum = "LOG_TAIL"
	UnifiedAgentLoggingSourceSourceTypeWindowsEventLog UnifiedAgentLoggingSourceSourceTypeEnum = "WINDOWS_EVENT_LOG"
)

var mappingUnifiedAgentLoggingSourceSourceTypeEnum = map[string]UnifiedAgentLoggingSourceSourceTypeEnum{
	"LOG_TAIL":          UnifiedAgentLoggingSourceSourceTypeLogTail,
	"WINDOWS_EVENT_LOG": UnifiedAgentLoggingSourceSourceTypeWindowsEventLog,
}

var mappingUnifiedAgentLoggingSourceSourceTypeEnumLowerCase = map[string]UnifiedAgentLoggingSourceSourceTypeEnum{
	"log_tail":          UnifiedAgentLoggingSourceSourceTypeLogTail,
	"windows_event_log": UnifiedAgentLoggingSourceSourceTypeWindowsEventLog,
}

// GetUnifiedAgentLoggingSourceSourceTypeEnumValues Enumerates the set of values for UnifiedAgentLoggingSourceSourceTypeEnum
func GetUnifiedAgentLoggingSourceSourceTypeEnumValues() []UnifiedAgentLoggingSourceSourceTypeEnum {
	values := make([]UnifiedAgentLoggingSourceSourceTypeEnum, 0)
	for _, v := range mappingUnifiedAgentLoggingSourceSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUnifiedAgentLoggingSourceSourceTypeEnumStringValues Enumerates the set of values in String for UnifiedAgentLoggingSourceSourceTypeEnum
func GetUnifiedAgentLoggingSourceSourceTypeEnumStringValues() []string {
	return []string{
		"LOG_TAIL",
		"WINDOWS_EVENT_LOG",
	}
}

// GetMappingUnifiedAgentLoggingSourceSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUnifiedAgentLoggingSourceSourceTypeEnum(val string) (UnifiedAgentLoggingSourceSourceTypeEnum, bool) {
	enum, ok := mappingUnifiedAgentLoggingSourceSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
