// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// PartitionConfig The information about partition configuration.
type PartitionConfig interface {
}

type partitionconfig struct {
	JsonData  []byte
	ModelType string `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *partitionconfig) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerpartitionconfig partitionconfig
	s := struct {
		Model Unmarshalerpartitionconfig
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *partitionconfig) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "KEYRANGEPARTITIONCONFIG":
		mm := KeyRangePartitionConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m partitionconfig) String() string {
	return common.PointerString(m)
}

// PartitionConfigModelTypeEnum Enum with underlying type: string
type PartitionConfigModelTypeEnum string

// Set of constants representing the allowable values for PartitionConfigModelTypeEnum
const (
	PartitionConfigModelTypeKeyrangepartitionconfig PartitionConfigModelTypeEnum = "KEYRANGEPARTITIONCONFIG"
)

var mappingPartitionConfigModelType = map[string]PartitionConfigModelTypeEnum{
	"KEYRANGEPARTITIONCONFIG": PartitionConfigModelTypeKeyrangepartitionconfig,
}

// GetPartitionConfigModelTypeEnumValues Enumerates the set of values for PartitionConfigModelTypeEnum
func GetPartitionConfigModelTypeEnumValues() []PartitionConfigModelTypeEnum {
	values := make([]PartitionConfigModelTypeEnum, 0)
	for _, v := range mappingPartitionConfigModelType {
		values = append(values, v)
	}
	return values
}
