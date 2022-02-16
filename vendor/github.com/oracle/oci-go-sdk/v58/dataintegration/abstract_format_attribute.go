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

// AbstractFormatAttribute The abstract format attribute.
type AbstractFormatAttribute interface {

	// Defines whether a file pattern is supported.
	GetIsFilePattern() *bool
}

type abstractformatattribute struct {
	JsonData      []byte
	IsFilePattern *bool  `mandatory:"false" json:"isFilePattern"`
	ModelType     string `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *abstractformatattribute) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerabstractformatattribute abstractformatattribute
	s := struct {
		Model Unmarshalerabstractformatattribute
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.IsFilePattern = s.Model.IsFilePattern
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *abstractformatattribute) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "AVRO_FORMAT":
		mm := AvroFormatAttribute{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "JSON_FORMAT":
		mm := JsonFormatAttribute{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CSV_FORMAT":
		mm := CsvFormatAttribute{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PARQUET_FORMAT":
		mm := ParquetFormatAttribute{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetIsFilePattern returns IsFilePattern
func (m abstractformatattribute) GetIsFilePattern() *bool {
	return m.IsFilePattern
}

func (m abstractformatattribute) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m abstractformatattribute) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AbstractFormatAttributeModelTypeEnum Enum with underlying type: string
type AbstractFormatAttributeModelTypeEnum string

// Set of constants representing the allowable values for AbstractFormatAttributeModelTypeEnum
const (
	AbstractFormatAttributeModelTypeJsonFormat AbstractFormatAttributeModelTypeEnum = "JSON_FORMAT"
	AbstractFormatAttributeModelTypeCsvFormat  AbstractFormatAttributeModelTypeEnum = "CSV_FORMAT"
	AbstractFormatAttributeModelTypeAvroFormat AbstractFormatAttributeModelTypeEnum = "AVRO_FORMAT"
)

var mappingAbstractFormatAttributeModelTypeEnum = map[string]AbstractFormatAttributeModelTypeEnum{
	"JSON_FORMAT": AbstractFormatAttributeModelTypeJsonFormat,
	"CSV_FORMAT":  AbstractFormatAttributeModelTypeCsvFormat,
	"AVRO_FORMAT": AbstractFormatAttributeModelTypeAvroFormat,
}

// GetAbstractFormatAttributeModelTypeEnumValues Enumerates the set of values for AbstractFormatAttributeModelTypeEnum
func GetAbstractFormatAttributeModelTypeEnumValues() []AbstractFormatAttributeModelTypeEnum {
	values := make([]AbstractFormatAttributeModelTypeEnum, 0)
	for _, v := range mappingAbstractFormatAttributeModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAbstractFormatAttributeModelTypeEnumStringValues Enumerates the set of values in String for AbstractFormatAttributeModelTypeEnum
func GetAbstractFormatAttributeModelTypeEnumStringValues() []string {
	return []string{
		"JSON_FORMAT",
		"CSV_FORMAT",
		"AVRO_FORMAT",
	}
}

// GetMappingAbstractFormatAttributeModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAbstractFormatAttributeModelTypeEnum(val string) (AbstractFormatAttributeModelTypeEnum, bool) {
	mappingAbstractFormatAttributeModelTypeEnumIgnoreCase := make(map[string]AbstractFormatAttributeModelTypeEnum)
	for k, v := range mappingAbstractFormatAttributeModelTypeEnum {
		mappingAbstractFormatAttributeModelTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingAbstractFormatAttributeModelTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
