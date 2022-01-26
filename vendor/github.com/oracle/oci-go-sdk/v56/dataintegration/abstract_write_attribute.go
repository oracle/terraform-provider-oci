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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// AbstractWriteAttribute The abstract write attribute.
type AbstractWriteAttribute interface {
}

type abstractwriteattribute struct {
	JsonData  []byte
	ModelType string `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *abstractwriteattribute) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerabstractwriteattribute abstractwriteattribute
	s := struct {
		Model Unmarshalerabstractwriteattribute
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *abstractwriteattribute) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "ORACLEADWCWRITEATTRIBUTE":
		mm := OracleAdwcWriteAttribute{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_ATP_WRITE_ATTRIBUTE":
		mm := OracleAtpWriteAttributes{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLEWRITEATTRIBUTE":
		mm := OracleWriteAttribute{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_WRITE_ATTRIBUTE":
		mm := OracleWriteAttributes{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLEATPWRITEATTRIBUTE":
		mm := OracleAtpWriteAttribute{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OBJECTSTORAGEWRITEATTRIBUTE":
		mm := ObjectStorageWriteAttribute{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_ADWC_WRITE_ATTRIBUTE":
		mm := OracleAdwcWriteAttributes{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OBJECT_STORAGE_WRITE_ATTRIBUTE":
		mm := ObjectStorageWriteAttributes{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m abstractwriteattribute) String() string {
	return common.PointerString(m)
}

// AbstractWriteAttributeModelTypeEnum Enum with underlying type: string
type AbstractWriteAttributeModelTypeEnum string

// Set of constants representing the allowable values for AbstractWriteAttributeModelTypeEnum
const (
	AbstractWriteAttributeModelTypeOraclewriteattribute        AbstractWriteAttributeModelTypeEnum = "ORACLEWRITEATTRIBUTE"
	AbstractWriteAttributeModelTypeOracleatpwriteattribute     AbstractWriteAttributeModelTypeEnum = "ORACLEATPWRITEATTRIBUTE"
	AbstractWriteAttributeModelTypeOracleadwcwriteattribute    AbstractWriteAttributeModelTypeEnum = "ORACLEADWCWRITEATTRIBUTE"
	AbstractWriteAttributeModelTypeObjectstoragewriteattribute AbstractWriteAttributeModelTypeEnum = "OBJECTSTORAGEWRITEATTRIBUTE"
	AbstractWriteAttributeModelTypeOracleWriteAttribute        AbstractWriteAttributeModelTypeEnum = "ORACLE_WRITE_ATTRIBUTE"
	AbstractWriteAttributeModelTypeOracleAtpWriteAttribute     AbstractWriteAttributeModelTypeEnum = "ORACLE_ATP_WRITE_ATTRIBUTE"
	AbstractWriteAttributeModelTypeOracleAdwcWriteAttribute    AbstractWriteAttributeModelTypeEnum = "ORACLE_ADWC_WRITE_ATTRIBUTE"
	AbstractWriteAttributeModelTypeObjectStorageWriteAttribute AbstractWriteAttributeModelTypeEnum = "OBJECT_STORAGE_WRITE_ATTRIBUTE"
)

var mappingAbstractWriteAttributeModelType = map[string]AbstractWriteAttributeModelTypeEnum{
	"ORACLEWRITEATTRIBUTE":           AbstractWriteAttributeModelTypeOraclewriteattribute,
	"ORACLEATPWRITEATTRIBUTE":        AbstractWriteAttributeModelTypeOracleatpwriteattribute,
	"ORACLEADWCWRITEATTRIBUTE":       AbstractWriteAttributeModelTypeOracleadwcwriteattribute,
	"OBJECTSTORAGEWRITEATTRIBUTE":    AbstractWriteAttributeModelTypeObjectstoragewriteattribute,
	"ORACLE_WRITE_ATTRIBUTE":         AbstractWriteAttributeModelTypeOracleWriteAttribute,
	"ORACLE_ATP_WRITE_ATTRIBUTE":     AbstractWriteAttributeModelTypeOracleAtpWriteAttribute,
	"ORACLE_ADWC_WRITE_ATTRIBUTE":    AbstractWriteAttributeModelTypeOracleAdwcWriteAttribute,
	"OBJECT_STORAGE_WRITE_ATTRIBUTE": AbstractWriteAttributeModelTypeObjectStorageWriteAttribute,
}

// GetAbstractWriteAttributeModelTypeEnumValues Enumerates the set of values for AbstractWriteAttributeModelTypeEnum
func GetAbstractWriteAttributeModelTypeEnumValues() []AbstractWriteAttributeModelTypeEnum {
	values := make([]AbstractWriteAttributeModelTypeEnum, 0)
	for _, v := range mappingAbstractWriteAttributeModelType {
		values = append(values, v)
	}
	return values
}
