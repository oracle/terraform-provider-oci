// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the DCMS APIs to perform Metadata/Data operations.
//

package dataconnectivity

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v62/common"
	"strings"
)

// DataFormat The data format object.
type DataFormat struct {

	// type
	Type DataFormatTypeEnum `mandatory:"true" json:"type"`

	FormatAttribute AbstractFormatAttribute `mandatory:"false" json:"formatAttribute"`

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
		CompressionConfig *Compression            `json:"compressionConfig"`
		Type              DataFormatTypeEnum      `json:"type"`
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

	m.CompressionConfig = model.CompressionConfig

	m.Type = model.Type

	return
}

// DataFormatTypeEnum Enum with underlying type: string
type DataFormatTypeEnum string

// Set of constants representing the allowable values for DataFormatTypeEnum
const (
	DataFormatTypeJson    DataFormatTypeEnum = "JSON"
	DataFormatTypeCsv     DataFormatTypeEnum = "CSV"
	DataFormatTypeParquet DataFormatTypeEnum = "PARQUET"
	DataFormatTypeAvro    DataFormatTypeEnum = "AVRO"
)

var mappingDataFormatTypeEnum = map[string]DataFormatTypeEnum{
	"JSON":    DataFormatTypeJson,
	"CSV":     DataFormatTypeCsv,
	"PARQUET": DataFormatTypeParquet,
	"AVRO":    DataFormatTypeAvro,
}

var mappingDataFormatTypeEnumLowerCase = map[string]DataFormatTypeEnum{
	"json":    DataFormatTypeJson,
	"csv":     DataFormatTypeCsv,
	"parquet": DataFormatTypeParquet,
	"avro":    DataFormatTypeAvro,
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
		"JSON",
		"CSV",
		"PARQUET",
		"AVRO",
	}
}

// GetMappingDataFormatTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataFormatTypeEnum(val string) (DataFormatTypeEnum, bool) {
	enum, ok := mappingDataFormatTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
