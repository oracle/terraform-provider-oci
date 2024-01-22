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

// TrainingConfig The fine-tuning method and hyperparameters used for fine-tuning a custom model.
type TrainingConfig interface {

	// The maximum number of training epochs to run for.
	GetTotalTrainingEpochs() *int

	// The initial learning rate to be used during training
	GetLearningRate() *float64

	// The batch size used during training.
	GetTrainingBatchSize() *int

	// Stop training if the loss metric does not improve beyond 'early_stopping_threshold' for this many times of evaluation.
	GetEarlyStoppingPatience() *int

	// How much the loss must improve to prevent early stopping.
	GetEarlyStoppingThreshold() *float64

	// Determines how frequently to log model metrics.
	// Every step is logged for the first 20 steps and then follows this parameter for log frequency. Set to 0 to disable logging the model metrics.
	GetLogModelMetricsIntervalInSteps() *int
}

type trainingconfig struct {
	JsonData                       []byte
	TotalTrainingEpochs            *int     `mandatory:"false" json:"totalTrainingEpochs"`
	LearningRate                   *float64 `mandatory:"false" json:"learningRate"`
	TrainingBatchSize              *int     `mandatory:"false" json:"trainingBatchSize"`
	EarlyStoppingPatience          *int     `mandatory:"false" json:"earlyStoppingPatience"`
	EarlyStoppingThreshold         *float64 `mandatory:"false" json:"earlyStoppingThreshold"`
	LogModelMetricsIntervalInSteps *int     `mandatory:"false" json:"logModelMetricsIntervalInSteps"`
	TrainingConfigType             string   `json:"trainingConfigType"`
}

// UnmarshalJSON unmarshals json
func (m *trainingconfig) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertrainingconfig trainingconfig
	s := struct {
		Model Unmarshalertrainingconfig
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.TotalTrainingEpochs = s.Model.TotalTrainingEpochs
	m.LearningRate = s.Model.LearningRate
	m.TrainingBatchSize = s.Model.TrainingBatchSize
	m.EarlyStoppingPatience = s.Model.EarlyStoppingPatience
	m.EarlyStoppingThreshold = s.Model.EarlyStoppingThreshold
	m.LogModelMetricsIntervalInSteps = s.Model.LogModelMetricsIntervalInSteps
	m.TrainingConfigType = s.Model.TrainingConfigType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *trainingconfig) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.TrainingConfigType {
	case "VANILLA_TRAINING_CONFIG":
		mm := VanillaTrainingConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TFEW_TRAINING_CONFIG":
		mm := TFewTrainingConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for TrainingConfig: %s.", m.TrainingConfigType)
		return *m, nil
	}
}

// GetTotalTrainingEpochs returns TotalTrainingEpochs
func (m trainingconfig) GetTotalTrainingEpochs() *int {
	return m.TotalTrainingEpochs
}

// GetLearningRate returns LearningRate
func (m trainingconfig) GetLearningRate() *float64 {
	return m.LearningRate
}

// GetTrainingBatchSize returns TrainingBatchSize
func (m trainingconfig) GetTrainingBatchSize() *int {
	return m.TrainingBatchSize
}

// GetEarlyStoppingPatience returns EarlyStoppingPatience
func (m trainingconfig) GetEarlyStoppingPatience() *int {
	return m.EarlyStoppingPatience
}

// GetEarlyStoppingThreshold returns EarlyStoppingThreshold
func (m trainingconfig) GetEarlyStoppingThreshold() *float64 {
	return m.EarlyStoppingThreshold
}

// GetLogModelMetricsIntervalInSteps returns LogModelMetricsIntervalInSteps
func (m trainingconfig) GetLogModelMetricsIntervalInSteps() *int {
	return m.LogModelMetricsIntervalInSteps
}

func (m trainingconfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m trainingconfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TrainingConfigTrainingConfigTypeEnum Enum with underlying type: string
type TrainingConfigTrainingConfigTypeEnum string

// Set of constants representing the allowable values for TrainingConfigTrainingConfigTypeEnum
const (
	TrainingConfigTrainingConfigTypeTfewTrainingConfig    TrainingConfigTrainingConfigTypeEnum = "TFEW_TRAINING_CONFIG"
	TrainingConfigTrainingConfigTypeVanillaTrainingConfig TrainingConfigTrainingConfigTypeEnum = "VANILLA_TRAINING_CONFIG"
)

var mappingTrainingConfigTrainingConfigTypeEnum = map[string]TrainingConfigTrainingConfigTypeEnum{
	"TFEW_TRAINING_CONFIG":    TrainingConfigTrainingConfigTypeTfewTrainingConfig,
	"VANILLA_TRAINING_CONFIG": TrainingConfigTrainingConfigTypeVanillaTrainingConfig,
}

var mappingTrainingConfigTrainingConfigTypeEnumLowerCase = map[string]TrainingConfigTrainingConfigTypeEnum{
	"tfew_training_config":    TrainingConfigTrainingConfigTypeTfewTrainingConfig,
	"vanilla_training_config": TrainingConfigTrainingConfigTypeVanillaTrainingConfig,
}

// GetTrainingConfigTrainingConfigTypeEnumValues Enumerates the set of values for TrainingConfigTrainingConfigTypeEnum
func GetTrainingConfigTrainingConfigTypeEnumValues() []TrainingConfigTrainingConfigTypeEnum {
	values := make([]TrainingConfigTrainingConfigTypeEnum, 0)
	for _, v := range mappingTrainingConfigTrainingConfigTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTrainingConfigTrainingConfigTypeEnumStringValues Enumerates the set of values in String for TrainingConfigTrainingConfigTypeEnum
func GetTrainingConfigTrainingConfigTypeEnumStringValues() []string {
	return []string{
		"TFEW_TRAINING_CONFIG",
		"VANILLA_TRAINING_CONFIG",
	}
}

// GetMappingTrainingConfigTrainingConfigTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTrainingConfigTrainingConfigTypeEnum(val string) (TrainingConfigTrainingConfigTypeEnum, bool) {
	enum, ok := mappingTrainingConfigTrainingConfigTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
