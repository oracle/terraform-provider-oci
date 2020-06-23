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

// AbstractReadAttribute The abstract read attribute.
type AbstractReadAttribute interface {
}

type abstractreadattribute struct {
	JsonData  []byte
	ModelType string `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *abstractreadattribute) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerabstractreadattribute abstractreadattribute
	s := struct {
		Model Unmarshalerabstractreadattribute
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *abstractreadattribute) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "ORACLEREADATTRIBUTE":
		mm := OracleReadAttribute{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m abstractreadattribute) String() string {
	return common.PointerString(m)
}

// AbstractReadAttributeModelTypeEnum Enum with underlying type: string
type AbstractReadAttributeModelTypeEnum string

// Set of constants representing the allowable values for AbstractReadAttributeModelTypeEnum
const (
	AbstractReadAttributeModelTypeOraclereadattribute AbstractReadAttributeModelTypeEnum = "ORACLEREADATTRIBUTE"
)

var mappingAbstractReadAttributeModelType = map[string]AbstractReadAttributeModelTypeEnum{
	"ORACLEREADATTRIBUTE": AbstractReadAttributeModelTypeOraclereadattribute,
}

// GetAbstractReadAttributeModelTypeEnumValues Enumerates the set of values for AbstractReadAttributeModelTypeEnum
func GetAbstractReadAttributeModelTypeEnumValues() []AbstractReadAttributeModelTypeEnum {
	values := make([]AbstractReadAttributeModelTypeEnum, 0)
	for _, v := range mappingAbstractReadAttributeModelType {
		values = append(values, v)
	}
	return values
}
