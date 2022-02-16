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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m partitionconfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PartitionConfigModelTypeEnum Enum with underlying type: string
type PartitionConfigModelTypeEnum string

// Set of constants representing the allowable values for PartitionConfigModelTypeEnum
const (
	PartitionConfigModelTypeKeyrangepartitionconfig PartitionConfigModelTypeEnum = "KEYRANGEPARTITIONCONFIG"
)

var mappingPartitionConfigModelTypeEnum = map[string]PartitionConfigModelTypeEnum{
	"KEYRANGEPARTITIONCONFIG": PartitionConfigModelTypeKeyrangepartitionconfig,
}

// GetPartitionConfigModelTypeEnumValues Enumerates the set of values for PartitionConfigModelTypeEnum
func GetPartitionConfigModelTypeEnumValues() []PartitionConfigModelTypeEnum {
	values := make([]PartitionConfigModelTypeEnum, 0)
	for _, v := range mappingPartitionConfigModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPartitionConfigModelTypeEnumStringValues Enumerates the set of values in String for PartitionConfigModelTypeEnum
func GetPartitionConfigModelTypeEnumStringValues() []string {
	return []string{
		"KEYRANGEPARTITIONCONFIG",
	}
}

// GetMappingPartitionConfigModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPartitionConfigModelTypeEnum(val string) (PartitionConfigModelTypeEnum, bool) {
	mappingPartitionConfigModelTypeEnumIgnoreCase := make(map[string]PartitionConfigModelTypeEnum)
	for k, v := range mappingPartitionConfigModelTypeEnum {
		mappingPartitionConfigModelTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingPartitionConfigModelTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
