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

// DatasetSourceDetails This allows the customer to specify the source of the dataset.
type DatasetSourceDetails interface {
}

type datasetsourcedetails struct {
	JsonData   []byte
	SourceType string `json:"sourceType"`
}

// UnmarshalJSON unmarshals json
func (m *datasetsourcedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatasetsourcedetails datasetsourcedetails
	s := struct {
		Model Unmarshalerdatasetsourcedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SourceType = s.Model.SourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *datasetsourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SourceType {
	case "OBJECT_STORAGE":
		mm := ObjectStorageSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m datasetsourcedetails) String() string {
	return common.PointerString(m)
}

// DatasetSourceDetailsSourceTypeEnum Enum with underlying type: string
type DatasetSourceDetailsSourceTypeEnum string

// Set of constants representing the allowable values for DatasetSourceDetailsSourceTypeEnum
const (
	DatasetSourceDetailsSourceTypeObjectStorage DatasetSourceDetailsSourceTypeEnum = "OBJECT_STORAGE"
)

var mappingDatasetSourceDetailsSourceType = map[string]DatasetSourceDetailsSourceTypeEnum{
	"OBJECT_STORAGE": DatasetSourceDetailsSourceTypeObjectStorage,
}

// GetDatasetSourceDetailsSourceTypeEnumValues Enumerates the set of values for DatasetSourceDetailsSourceTypeEnum
func GetDatasetSourceDetailsSourceTypeEnumValues() []DatasetSourceDetailsSourceTypeEnum {
	values := make([]DatasetSourceDetailsSourceTypeEnum, 0)
	for _, v := range mappingDatasetSourceDetailsSourceType {
		values = append(values, v)
	}
	return values
}
