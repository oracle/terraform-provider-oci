// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service Management API
//
// OCI Generative AI is a fully managed service that provides a set of state-of-the-art, customizable large language models (LLMs) that cover a wide range of use cases for text generation, summarization, and text embeddings.
// Use the Generative AI service management API to create and manage DedicatedAiCluster, Endpoint, Model, and WorkRequest in the Generative AI service. For example, create a custom model by fine-tuning an out-of-the-box model using your own data, on a fine-tuning dedicated AI cluster. Then, create a hosting dedicated AI cluster with an endpoint to host your custom model.
// To access your custom model endpoints, or to try the out-of-the-box models to generate text, summarize, and create text embeddings see the Generative AI Inference API (https://docs.cloud.oracle.com/#/en/generative-ai-inference/latest/).
// To learn more about the service, see the Generative AI documentation (https://docs.cloud.oracle.com/iaas/Content/generative-ai/home.htm).
//

package generativeai

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ModelMetrics Model metrics during the creation of a new model.
type ModelMetrics interface {
}

type modelmetrics struct {
	JsonData         []byte
	ModelMetricsType string `json:"modelMetricsType"`
}

// UnmarshalJSON unmarshals json
func (m *modelmetrics) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalermodelmetrics modelmetrics
	s := struct {
		Model Unmarshalermodelmetrics
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ModelMetricsType = s.Model.ModelMetricsType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *modelmetrics) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelMetricsType {
	case "TEXT_GENERATION_MODEL_METRICS":
		mm := TextGenerationModelMetrics{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ModelMetrics: %s.", m.ModelMetricsType)
		return *m, nil
	}
}

func (m modelmetrics) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m modelmetrics) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ModelMetricsModelMetricsTypeEnum Enum with underlying type: string
type ModelMetricsModelMetricsTypeEnum string

// Set of constants representing the allowable values for ModelMetricsModelMetricsTypeEnum
const (
	ModelMetricsModelMetricsTypeTextGenerationModelMetrics ModelMetricsModelMetricsTypeEnum = "TEXT_GENERATION_MODEL_METRICS"
)

var mappingModelMetricsModelMetricsTypeEnum = map[string]ModelMetricsModelMetricsTypeEnum{
	"TEXT_GENERATION_MODEL_METRICS": ModelMetricsModelMetricsTypeTextGenerationModelMetrics,
}

var mappingModelMetricsModelMetricsTypeEnumLowerCase = map[string]ModelMetricsModelMetricsTypeEnum{
	"text_generation_model_metrics": ModelMetricsModelMetricsTypeTextGenerationModelMetrics,
}

// GetModelMetricsModelMetricsTypeEnumValues Enumerates the set of values for ModelMetricsModelMetricsTypeEnum
func GetModelMetricsModelMetricsTypeEnumValues() []ModelMetricsModelMetricsTypeEnum {
	values := make([]ModelMetricsModelMetricsTypeEnum, 0)
	for _, v := range mappingModelMetricsModelMetricsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetModelMetricsModelMetricsTypeEnumStringValues Enumerates the set of values in String for ModelMetricsModelMetricsTypeEnum
func GetModelMetricsModelMetricsTypeEnumStringValues() []string {
	return []string{
		"TEXT_GENERATION_MODEL_METRICS",
	}
}

// GetMappingModelMetricsModelMetricsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModelMetricsModelMetricsTypeEnum(val string) (ModelMetricsModelMetricsTypeEnum, bool) {
	enum, ok := mappingModelMetricsModelMetricsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
