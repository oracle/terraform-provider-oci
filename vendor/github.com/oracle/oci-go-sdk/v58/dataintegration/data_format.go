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

// DataFormat The data format object.
type DataFormat struct {
	FormatAttribute AbstractFormatAttribute `mandatory:"false" json:"formatAttribute"`

	// type
	Type DataFormatTypeEnum `mandatory:"false" json:"type,omitempty"`

	CompressionConfig *Compression `mandatory:"false" json:"compressionConfig"`
}

func (m DataFormat) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataFormat) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDataFormatTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetDataFormatTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *DataFormat) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		FormatAttribute   abstractformatattribute `json:"formatAttribute"`
		Type              DataFormatTypeEnum      `json:"type"`
		CompressionConfig *Compression            `json:"compressionConfig"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.FormatAttribute.UnmarshalPolymorphicJSON(model.FormatAttribute.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.FormatAttribute = nn.(AbstractFormatAttribute)
	} else {
		m.FormatAttribute = nil
	}

	m.Type = model.Type

	m.CompressionConfig = model.CompressionConfig

	return
}

// DataFormatTypeEnum Enum with underlying type: string
type DataFormatTypeEnum string

// Set of constants representing the allowable values for DataFormatTypeEnum
const (
	DataFormatTypeXml     DataFormatTypeEnum = "XML"
	DataFormatTypeJson    DataFormatTypeEnum = "JSON"
	DataFormatTypeCsv     DataFormatTypeEnum = "CSV"
	DataFormatTypeOrc     DataFormatTypeEnum = "ORC"
	DataFormatTypeParquet DataFormatTypeEnum = "PARQUET"
	DataFormatTypeAvro    DataFormatTypeEnum = "AVRO"
)

var mappingDataFormatTypeEnum = map[string]DataFormatTypeEnum{
	"XML":     DataFormatTypeXml,
	"JSON":    DataFormatTypeJson,
	"CSV":     DataFormatTypeCsv,
	"ORC":     DataFormatTypeOrc,
	"PARQUET": DataFormatTypeParquet,
	"AVRO":    DataFormatTypeAvro,
}

// GetDataFormatTypeEnumValues Enumerates the set of values for DataFormatTypeEnum
func GetDataFormatTypeEnumValues() []DataFormatTypeEnum {
	values := make([]DataFormatTypeEnum, 0)
	for _, v := range mappingDataFormatTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDataFormatTypeEnumStringValues Enumerates the set of values in String for DataFormatTypeEnum
func GetDataFormatTypeEnumStringValues() []string {
	return []string{
		"XML",
		"JSON",
		"CSV",
		"ORC",
		"PARQUET",
		"AVRO",
	}
}

// GetMappingDataFormatTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataFormatTypeEnum(val string) (DataFormatTypeEnum, bool) {
	mappingDataFormatTypeEnumIgnoreCase := make(map[string]DataFormatTypeEnum)
	for k, v := range mappingDataFormatTypeEnum {
		mappingDataFormatTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingDataFormatTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
