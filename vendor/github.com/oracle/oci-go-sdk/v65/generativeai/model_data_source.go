// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service Management API
//
// OCI Generative AI is a fully managed service that provides a set of state-of-the-art, customizable large language models (LLMs) that cover a wide range of use cases for text generation, summarization, and text embeddings.
// Use the Generative AI service management API to create and manage DedicatedAiCluster, Endpoint, Model, and WorkRequest in the Generative AI service. For example, create a custom model by fine-tuning an out-of-the-box model using your own data, on a fine-tuning dedicated AI cluster. Then, create a hosting dedicated AI cluster with an endpoint to host your custom model.
// To access your custom model endpoints, or to try the out-of-the-box models to generate text, summarize, and create text embeddings see the Generative AI Inference API (https://docs.oracle.com/iaas/api/#/en/generative-ai-inference/latest/).
// To learn more about the service, see the Generative AI documentation (https://docs.oracle.com/iaas/Content/generative-ai/home.htm).
//

package generativeai

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ModelDataSource Defines the source location and method used to import the model. Supports importing from Hugging Face,
// an Object Storage location, or by referencing an already imported model.
type ModelDataSource interface {
}

type modeldatasource struct {
	JsonData   []byte
	SourceType string `json:"sourceType"`
}

// UnmarshalJSON unmarshals json
func (m *modeldatasource) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalermodeldatasource modeldatasource
	s := struct {
		Model Unmarshalermodeldatasource
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SourceType = s.Model.SourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *modeldatasource) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SourceType {
	case "OBJECT_STORAGE_OBJECT":
		mm := ObjectStorageObject{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HUGGING_FACE_MODEL":
		mm := HuggingFaceModel{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ModelDataSource: %s.", m.SourceType)
		return *m, nil
	}
}

func (m modeldatasource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m modeldatasource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ModelDataSourceSourceTypeEnum Enum with underlying type: string
type ModelDataSourceSourceTypeEnum string

// Set of constants representing the allowable values for ModelDataSourceSourceTypeEnum
const (
	ModelDataSourceSourceTypeHuggingFaceModel    ModelDataSourceSourceTypeEnum = "HUGGING_FACE_MODEL"
	ModelDataSourceSourceTypeObjectStorageObject ModelDataSourceSourceTypeEnum = "OBJECT_STORAGE_OBJECT"
)

var mappingModelDataSourceSourceTypeEnum = map[string]ModelDataSourceSourceTypeEnum{
	"HUGGING_FACE_MODEL":    ModelDataSourceSourceTypeHuggingFaceModel,
	"OBJECT_STORAGE_OBJECT": ModelDataSourceSourceTypeObjectStorageObject,
}

var mappingModelDataSourceSourceTypeEnumLowerCase = map[string]ModelDataSourceSourceTypeEnum{
	"hugging_face_model":    ModelDataSourceSourceTypeHuggingFaceModel,
	"object_storage_object": ModelDataSourceSourceTypeObjectStorageObject,
}

// GetModelDataSourceSourceTypeEnumValues Enumerates the set of values for ModelDataSourceSourceTypeEnum
func GetModelDataSourceSourceTypeEnumValues() []ModelDataSourceSourceTypeEnum {
	values := make([]ModelDataSourceSourceTypeEnum, 0)
	for _, v := range mappingModelDataSourceSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetModelDataSourceSourceTypeEnumStringValues Enumerates the set of values in String for ModelDataSourceSourceTypeEnum
func GetModelDataSourceSourceTypeEnumStringValues() []string {
	return []string{
		"HUGGING_FACE_MODEL",
		"OBJECT_STORAGE_OBJECT",
	}
}

// GetMappingModelDataSourceSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModelDataSourceSourceTypeEnum(val string) (ModelDataSourceSourceTypeEnum, bool) {
	enum, ok := mappingModelDataSourceSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
