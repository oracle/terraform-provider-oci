// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// WriteOperationConfig The information about the write operation.
type WriteOperationConfig struct {

	// The object key.
	Key *string `mandatory:"false" json:"key"`

	// The object's model version.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// An array of operations.
	Operations []PushDownOperation `mandatory:"false" json:"operations"`

	DataFormat *DataFormat `mandatory:"false" json:"dataFormat"`

	PartitionConfig PartitionConfig `mandatory:"false" json:"partitionConfig"`

	WriteAttribute AbstractWriteAttribute `mandatory:"false" json:"writeAttribute"`

	MergeKey UniqueKey `mandatory:"false" json:"mergeKey"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// The mode for the write operation.
	WriteMode WriteOperationConfigWriteModeEnum `mandatory:"false" json:"writeMode,omitempty"`
}

func (m WriteOperationConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WriteOperationConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingWriteOperationConfigWriteModeEnum(string(m.WriteMode)); !ok && m.WriteMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for WriteMode: %s. Supported values are: %s.", m.WriteMode, strings.Join(GetWriteOperationConfigWriteModeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m WriteOperationConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeWriteOperationConfig WriteOperationConfig
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeWriteOperationConfig
	}{
		"WRITE_OPERATION_CONFIG",
		(MarshalTypeWriteOperationConfig)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *WriteOperationConfig) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key             *string                           `json:"key"`
		ModelVersion    *string                           `json:"modelVersion"`
		ParentRef       *ParentReference                  `json:"parentRef"`
		Operations      []pushdownoperation               `json:"operations"`
		DataFormat      *DataFormat                       `json:"dataFormat"`
		PartitionConfig partitionconfig                   `json:"partitionConfig"`
		WriteAttribute  abstractwriteattribute            `json:"writeAttribute"`
		WriteMode       WriteOperationConfigWriteModeEnum `json:"writeMode"`
		MergeKey        uniquekey                         `json:"mergeKey"`
		ObjectStatus    *int                              `json:"objectStatus"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Key = model.Key

	m.ModelVersion = model.ModelVersion

	m.ParentRef = model.ParentRef

	m.Operations = make([]PushDownOperation, len(model.Operations))
	for i, n := range model.Operations {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Operations[i] = nn.(PushDownOperation)
		} else {
			m.Operations[i] = nil
		}
	}

	m.DataFormat = model.DataFormat

	nn, e = model.PartitionConfig.UnmarshalPolymorphicJSON(model.PartitionConfig.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.PartitionConfig = nn.(PartitionConfig)
	} else {
		m.PartitionConfig = nil
	}

	nn, e = model.WriteAttribute.UnmarshalPolymorphicJSON(model.WriteAttribute.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.WriteAttribute = nn.(AbstractWriteAttribute)
	} else {
		m.WriteAttribute = nil
	}

	m.WriteMode = model.WriteMode

	nn, e = model.MergeKey.UnmarshalPolymorphicJSON(model.MergeKey.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.MergeKey = nn.(UniqueKey)
	} else {
		m.MergeKey = nil
	}

	m.ObjectStatus = model.ObjectStatus

	return
}

// WriteOperationConfigWriteModeEnum Enum with underlying type: string
type WriteOperationConfigWriteModeEnum string

// Set of constants representing the allowable values for WriteOperationConfigWriteModeEnum
const (
	WriteOperationConfigWriteModeOverwrite WriteOperationConfigWriteModeEnum = "OVERWRITE"
	WriteOperationConfigWriteModeAppend    WriteOperationConfigWriteModeEnum = "APPEND"
	WriteOperationConfigWriteModeMerge     WriteOperationConfigWriteModeEnum = "MERGE"
	WriteOperationConfigWriteModeIgnore    WriteOperationConfigWriteModeEnum = "IGNORE"
)

var mappingWriteOperationConfigWriteModeEnum = map[string]WriteOperationConfigWriteModeEnum{
	"OVERWRITE": WriteOperationConfigWriteModeOverwrite,
	"APPEND":    WriteOperationConfigWriteModeAppend,
	"MERGE":     WriteOperationConfigWriteModeMerge,
	"IGNORE":    WriteOperationConfigWriteModeIgnore,
}

// GetWriteOperationConfigWriteModeEnumValues Enumerates the set of values for WriteOperationConfigWriteModeEnum
func GetWriteOperationConfigWriteModeEnumValues() []WriteOperationConfigWriteModeEnum {
	values := make([]WriteOperationConfigWriteModeEnum, 0)
	for _, v := range mappingWriteOperationConfigWriteModeEnum {
		values = append(values, v)
	}
	return values
}

// GetWriteOperationConfigWriteModeEnumStringValues Enumerates the set of values in String for WriteOperationConfigWriteModeEnum
func GetWriteOperationConfigWriteModeEnumStringValues() []string {
	return []string{
		"OVERWRITE",
		"APPEND",
		"MERGE",
		"IGNORE",
	}
}

// GetMappingWriteOperationConfigWriteModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWriteOperationConfigWriteModeEnum(val string) (WriteOperationConfigWriteModeEnum, bool) {
	mappingWriteOperationConfigWriteModeEnumIgnoreCase := make(map[string]WriteOperationConfigWriteModeEnum)
	for k, v := range mappingWriteOperationConfigWriteModeEnum {
		mappingWriteOperationConfigWriteModeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingWriteOperationConfigWriteModeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
