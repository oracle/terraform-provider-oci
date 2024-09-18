// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service Management API
//
// OCI Generative AI is a fully managed service that provides a set of state-of-the-art, customizable large language models (LLMs) that cover a wide range of use cases for text generation, summarization, and text embeddings.
// Use the Generative AI service management API to create and manage DedicatedAiCluster, Endpoint, Model, and WorkRequest in the Generative AI service. For example, create a custom model by fine-tuning an out-of-the-box model using your own data, on a fine-tuning dedicated AI cluster. Then, create a hosting dedicated AI cluster with an endpoint to host your custom model.
// To access your custom model endpoints, or to try the out-of-the-box models to generate text, summarize, and create text embeddings see the Generative AI Inference API (https://docs.cloud.oracle.com/iaas/api/#/en/generative-ai-inference/latest/).
// To learn more about the service, see the Generative AI documentation (https://docs.cloud.oracle.com/iaas/Content/generative-ai/home.htm).
//

package generativeai

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Dataset The dataset used to fine-tune the model.
// Only one dataset is allowed per custom model, which is split 80-20 for training and validating. You must provide the dataset in a JSON Lines (JSONL) file. Each line in the JSONL file must have the format:
//
// `{"prompt": "<first prompt>", "completion": "<expected completion given first prompt>"}`
type Dataset interface {
}

type dataset struct {
	JsonData    []byte
	DatasetType string `json:"datasetType"`
}

// UnmarshalJSON unmarshals json
func (m *dataset) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdataset dataset
	s := struct {
		Model Unmarshalerdataset
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DatasetType = s.Model.DatasetType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *dataset) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DatasetType {
	case "OBJECT_STORAGE":
		mm := ObjectStorageDataset{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for Dataset: %s.", m.DatasetType)
		return *m, nil
	}
}

func (m dataset) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m dataset) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatasetDatasetTypeEnum Enum with underlying type: string
type DatasetDatasetTypeEnum string

// Set of constants representing the allowable values for DatasetDatasetTypeEnum
const (
	DatasetDatasetTypeObjectStorage DatasetDatasetTypeEnum = "OBJECT_STORAGE"
)

var mappingDatasetDatasetTypeEnum = map[string]DatasetDatasetTypeEnum{
	"OBJECT_STORAGE": DatasetDatasetTypeObjectStorage,
}

var mappingDatasetDatasetTypeEnumLowerCase = map[string]DatasetDatasetTypeEnum{
	"object_storage": DatasetDatasetTypeObjectStorage,
}

// GetDatasetDatasetTypeEnumValues Enumerates the set of values for DatasetDatasetTypeEnum
func GetDatasetDatasetTypeEnumValues() []DatasetDatasetTypeEnum {
	values := make([]DatasetDatasetTypeEnum, 0)
	for _, v := range mappingDatasetDatasetTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatasetDatasetTypeEnumStringValues Enumerates the set of values in String for DatasetDatasetTypeEnum
func GetDatasetDatasetTypeEnumStringValues() []string {
	return []string{
		"OBJECT_STORAGE",
	}
}

// GetMappingDatasetDatasetTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatasetDatasetTypeEnum(val string) (DatasetDatasetTypeEnum, bool) {
	enum, ok := mappingDatasetDatasetTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
