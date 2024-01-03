// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Language API
//
// OCI Language Service solutions can help enterprise customers integrate AI into their products immediately using our proven,
// pre-trained and custom models or containers, without a need to set up an house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI and ML operations, which shortens the time to market.
//

package ailanguage

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatasetDetails Possible data set type
type DatasetDetails interface {
}

type datasetdetails struct {
	JsonData    []byte
	DatasetType string `json:"datasetType"`
}

// UnmarshalJSON unmarshals json
func (m *datasetdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatasetdetails datasetdetails
	s := struct {
		Model Unmarshalerdatasetdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DatasetType = s.Model.DatasetType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *datasetdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DatasetType {
	case "DATA_SCIENCE_LABELING":
		mm := DataScienceLabelingDataset{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OBJECT_STORAGE":
		mm := ObjectStorageDataset{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DatasetDetails: %s.", m.DatasetType)
		return *m, nil
	}
}

func (m datasetdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m datasetdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatasetDetailsDatasetTypeEnum Enum with underlying type: string
type DatasetDetailsDatasetTypeEnum string

// Set of constants representing the allowable values for DatasetDetailsDatasetTypeEnum
const (
	DatasetDetailsDatasetTypeObjectStorage       DatasetDetailsDatasetTypeEnum = "OBJECT_STORAGE"
	DatasetDetailsDatasetTypeDataScienceLabeling DatasetDetailsDatasetTypeEnum = "DATA_SCIENCE_LABELING"
)

var mappingDatasetDetailsDatasetTypeEnum = map[string]DatasetDetailsDatasetTypeEnum{
	"OBJECT_STORAGE":        DatasetDetailsDatasetTypeObjectStorage,
	"DATA_SCIENCE_LABELING": DatasetDetailsDatasetTypeDataScienceLabeling,
}

var mappingDatasetDetailsDatasetTypeEnumLowerCase = map[string]DatasetDetailsDatasetTypeEnum{
	"object_storage":        DatasetDetailsDatasetTypeObjectStorage,
	"data_science_labeling": DatasetDetailsDatasetTypeDataScienceLabeling,
}

// GetDatasetDetailsDatasetTypeEnumValues Enumerates the set of values for DatasetDetailsDatasetTypeEnum
func GetDatasetDetailsDatasetTypeEnumValues() []DatasetDetailsDatasetTypeEnum {
	values := make([]DatasetDetailsDatasetTypeEnum, 0)
	for _, v := range mappingDatasetDetailsDatasetTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatasetDetailsDatasetTypeEnumStringValues Enumerates the set of values in String for DatasetDetailsDatasetTypeEnum
func GetDatasetDetailsDatasetTypeEnumStringValues() []string {
	return []string{
		"OBJECT_STORAGE",
		"DATA_SCIENCE_LABELING",
	}
}

// GetMappingDatasetDetailsDatasetTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatasetDetailsDatasetTypeEnum(val string) (DatasetDetailsDatasetTypeEnum, bool) {
	enum, ok := mappingDatasetDetailsDatasetTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
