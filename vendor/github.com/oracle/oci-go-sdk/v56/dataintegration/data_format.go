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

var mappingDataFormatType = map[string]DataFormatTypeEnum{
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
	for _, v := range mappingDataFormatType {
		values = append(values, v)
	}
	return values
}
