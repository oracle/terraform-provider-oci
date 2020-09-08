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

// UnifiedAgentLoggingSource logging source object.
type UnifiedAgentLoggingSource interface {

	// unique name for the source
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
		return *m, nil
	}
}

//GetName returns Name
func (m unifiedagentloggingsource) GetName() *string {
	return m.Name
}

func (m unifiedagentloggingsource) String() string {
	return common.PointerString(m)
}

// UnifiedAgentLoggingSourceSourceTypeEnum Enum with underlying type: string
type UnifiedAgentLoggingSourceSourceTypeEnum string

// Set of constants representing the allowable values for UnifiedAgentLoggingSourceSourceTypeEnum
const (
	UnifiedAgentLoggingSourceSourceTypeLogTail         UnifiedAgentLoggingSourceSourceTypeEnum = "LOG_TAIL"
	UnifiedAgentLoggingSourceSourceTypeWindowsEventLog UnifiedAgentLoggingSourceSourceTypeEnum = "WINDOWS_EVENT_LOG"
)

var mappingUnifiedAgentLoggingSourceSourceType = map[string]UnifiedAgentLoggingSourceSourceTypeEnum{
	"LOG_TAIL":          UnifiedAgentLoggingSourceSourceTypeLogTail,
	"WINDOWS_EVENT_LOG": UnifiedAgentLoggingSourceSourceTypeWindowsEventLog,
}

// GetUnifiedAgentLoggingSourceSourceTypeEnumValues Enumerates the set of values for UnifiedAgentLoggingSourceSourceTypeEnum
func GetUnifiedAgentLoggingSourceSourceTypeEnumValues() []UnifiedAgentLoggingSourceSourceTypeEnum {
	values := make([]UnifiedAgentLoggingSourceSourceTypeEnum, 0)
	for _, v := range mappingUnifiedAgentLoggingSourceSourceType {
		values = append(values, v)
	}
	return values
}
