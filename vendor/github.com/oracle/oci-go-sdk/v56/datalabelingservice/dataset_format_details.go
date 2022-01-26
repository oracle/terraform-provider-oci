// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Labeling Management API
//
// Use Data Labeling Management API to create, list, edit & delete datasets.
//

package datalabelingservice

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// DatasetFormatDetails Specifies how to process the data. Supported formats include DOCUMENT, IMAGE and TEXT.
type DatasetFormatDetails interface {
}

type datasetformatdetails struct {
	JsonData   []byte
	FormatType string `json:"formatType"`
}

// UnmarshalJSON unmarshals json
func (m *datasetformatdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatasetformatdetails datasetformatdetails
	s := struct {
		Model Unmarshalerdatasetformatdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.FormatType = s.Model.FormatType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *datasetformatdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.FormatType {
	case "IMAGE":
		mm := ImageDatasetFormatDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DOCUMENT":
		mm := DocumentDatasetFormatDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TEXT":
		mm := TextDatasetFormatDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m datasetformatdetails) String() string {
	return common.PointerString(m)
}

// DatasetFormatDetailsFormatTypeEnum Enum with underlying type: string
type DatasetFormatDetailsFormatTypeEnum string

// Set of constants representing the allowable values for DatasetFormatDetailsFormatTypeEnum
const (
	DatasetFormatDetailsFormatTypeDocument DatasetFormatDetailsFormatTypeEnum = "DOCUMENT"
	DatasetFormatDetailsFormatTypeImage    DatasetFormatDetailsFormatTypeEnum = "IMAGE"
	DatasetFormatDetailsFormatTypeText     DatasetFormatDetailsFormatTypeEnum = "TEXT"
)

var mappingDatasetFormatDetailsFormatType = map[string]DatasetFormatDetailsFormatTypeEnum{
	"DOCUMENT": DatasetFormatDetailsFormatTypeDocument,
	"IMAGE":    DatasetFormatDetailsFormatTypeImage,
	"TEXT":     DatasetFormatDetailsFormatTypeText,
}

// GetDatasetFormatDetailsFormatTypeEnumValues Enumerates the set of values for DatasetFormatDetailsFormatTypeEnum
func GetDatasetFormatDetailsFormatTypeEnumValues() []DatasetFormatDetailsFormatTypeEnum {
	values := make([]DatasetFormatDetailsFormatTypeEnum, 0)
	for _, v := range mappingDatasetFormatDetailsFormatType {
		values = append(values, v)
	}
	return values
}
