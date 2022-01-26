// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Agent API
//
// API for the Oracle Cloud Agent software running on compute instances. Oracle Cloud Agent
// is a lightweight process that monitors and manages compute instances.
//

package computeinstanceagent

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// InstanceAgentCommandSourceDetails The source of the command.
type InstanceAgentCommandSourceDetails interface {
}

type instanceagentcommandsourcedetails struct {
	JsonData   []byte
	SourceType string `json:"sourceType"`
}

// UnmarshalJSON unmarshals json
func (m *instanceagentcommandsourcedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerinstanceagentcommandsourcedetails instanceagentcommandsourcedetails
	s := struct {
		Model Unmarshalerinstanceagentcommandsourcedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SourceType = s.Model.SourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *instanceagentcommandsourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SourceType {
	case "OBJECT_STORAGE_TUPLE":
		mm := InstanceAgentCommandSourceViaObjectStorageTupleDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OBJECT_STORAGE_URI":
		mm := InstanceAgentCommandSourceViaObjectStorageUriDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TEXT":
		mm := InstanceAgentCommandSourceViaTextDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m instanceagentcommandsourcedetails) String() string {
	return common.PointerString(m)
}

// InstanceAgentCommandSourceDetailsSourceTypeEnum Enum with underlying type: string
type InstanceAgentCommandSourceDetailsSourceTypeEnum string

// Set of constants representing the allowable values for InstanceAgentCommandSourceDetailsSourceTypeEnum
const (
	InstanceAgentCommandSourceDetailsSourceTypeText               InstanceAgentCommandSourceDetailsSourceTypeEnum = "TEXT"
	InstanceAgentCommandSourceDetailsSourceTypeObjectStorageUri   InstanceAgentCommandSourceDetailsSourceTypeEnum = "OBJECT_STORAGE_URI"
	InstanceAgentCommandSourceDetailsSourceTypeObjectStorageTuple InstanceAgentCommandSourceDetailsSourceTypeEnum = "OBJECT_STORAGE_TUPLE"
)

var mappingInstanceAgentCommandSourceDetailsSourceType = map[string]InstanceAgentCommandSourceDetailsSourceTypeEnum{
	"TEXT":                 InstanceAgentCommandSourceDetailsSourceTypeText,
	"OBJECT_STORAGE_URI":   InstanceAgentCommandSourceDetailsSourceTypeObjectStorageUri,
	"OBJECT_STORAGE_TUPLE": InstanceAgentCommandSourceDetailsSourceTypeObjectStorageTuple,
}

// GetInstanceAgentCommandSourceDetailsSourceTypeEnumValues Enumerates the set of values for InstanceAgentCommandSourceDetailsSourceTypeEnum
func GetInstanceAgentCommandSourceDetailsSourceTypeEnumValues() []InstanceAgentCommandSourceDetailsSourceTypeEnum {
	values := make([]InstanceAgentCommandSourceDetailsSourceTypeEnum, 0)
	for _, v := range mappingInstanceAgentCommandSourceDetailsSourceType {
		values = append(values, v)
	}
	return values
}
