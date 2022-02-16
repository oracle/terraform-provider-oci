// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// VisionService API
//
// A description of the VisionService API.
//

package aivision

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// Model Machine-learned Model.
type Model struct {

	// Unique identifier that is immutable after creation.
	Id *string `mandatory:"true" json:"id"`

	// Compartment identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// What type of Vision model this is.
	ModelType ModelModelTypeEnum `mandatory:"true" json:"modelType"`

	TrainingDataset Dataset `mandatory:"true" json:"trainingDataset"`

	// The version of the model.
	ModelVersion *string `mandatory:"true" json:"modelVersion"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project which contains the model.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// When the model was created, as an RFC3339 datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Current state of the model.
	LifecycleState ModelLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Human-friendly name for the model, which can be changed.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Optional description of the model.
	Description *string `mandatory:"false" json:"description"`

	// Set to true when experimenting with a new model type or dataset so model training is quick, with a predefined low number of passes through the training data.
	IsQuickMode *bool `mandatory:"false" json:"isQuickMode"`

	// Maximum model training duration in hours, expressed as a decimal fraction.
	MaxTrainingDurationInHours *float64 `mandatory:"false" json:"maxTrainingDurationInHours"`

	// Total hours actually used for model training.
	TrainedDurationInHours *float64 `mandatory:"false" json:"trainedDurationInHours"`

	TestingDataset Dataset `mandatory:"false" json:"testingDataset"`

	ValidationDataset Dataset `mandatory:"false" json:"validationDataset"`

	// When the model was updated, as an RFC3339 datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail which can provide actionable information if training failed.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Precision of the trained model.
	Precision *float32 `mandatory:"false" json:"precision"`

	// Recall of the trained model.
	Recall *float32 `mandatory:"false" json:"recall"`

	// Mean average precision of the trained model.
	AveragePrecision *float32 `mandatory:"false" json:"averagePrecision"`

	// Intersection over union threshold used for calculating precision and recall.
	ConfidenceThreshold *float32 `mandatory:"false" json:"confidenceThreshold"`

	// Number of images in the dataset used to train, validate, and test the model.
	TotalImageCount *int `mandatory:"false" json:"totalImageCount"`

	// Number of images set aside for evaluating model performance metrics after training.
	TestImageCount *int `mandatory:"false" json:"testImageCount"`

	// Complete set of per-label metrics for successfully trained model.
	Metrics *string `mandatory:"false" json:"metrics"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m Model) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Model) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingModelModelTypeEnum(string(m.ModelType)); !ok && m.ModelType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ModelType: %s. Supported values are: %s.", m.ModelType, strings.Join(GetModelModelTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingModelLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetModelLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *Model) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName                *string                           `json:"displayName"`
		Description                *string                           `json:"description"`
		IsQuickMode                *bool                             `json:"isQuickMode"`
		MaxTrainingDurationInHours *float64                          `json:"maxTrainingDurationInHours"`
		TrainedDurationInHours     *float64                          `json:"trainedDurationInHours"`
		TestingDataset             dataset                           `json:"testingDataset"`
		ValidationDataset          dataset                           `json:"validationDataset"`
		TimeUpdated                *common.SDKTime                   `json:"timeUpdated"`
		LifecycleDetails           *string                           `json:"lifecycleDetails"`
		Precision                  *float32                          `json:"precision"`
		Recall                     *float32                          `json:"recall"`
		AveragePrecision           *float32                          `json:"averagePrecision"`
		ConfidenceThreshold        *float32                          `json:"confidenceThreshold"`
		TotalImageCount            *int                              `json:"totalImageCount"`
		TestImageCount             *int                              `json:"testImageCount"`
		Metrics                    *string                           `json:"metrics"`
		FreeformTags               map[string]string                 `json:"freeformTags"`
		DefinedTags                map[string]map[string]interface{} `json:"definedTags"`
		SystemTags                 map[string]map[string]interface{} `json:"systemTags"`
		Id                         *string                           `json:"id"`
		CompartmentId              *string                           `json:"compartmentId"`
		ModelType                  ModelModelTypeEnum                `json:"modelType"`
		TrainingDataset            dataset                           `json:"trainingDataset"`
		ModelVersion               *string                           `json:"modelVersion"`
		ProjectId                  *string                           `json:"projectId"`
		TimeCreated                *common.SDKTime                   `json:"timeCreated"`
		LifecycleState             ModelLifecycleStateEnum           `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.IsQuickMode = model.IsQuickMode

	m.MaxTrainingDurationInHours = model.MaxTrainingDurationInHours

	m.TrainedDurationInHours = model.TrainedDurationInHours

	nn, e = model.TestingDataset.UnmarshalPolymorphicJSON(model.TestingDataset.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.TestingDataset = nn.(Dataset)
	} else {
		m.TestingDataset = nil
	}

	nn, e = model.ValidationDataset.UnmarshalPolymorphicJSON(model.ValidationDataset.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ValidationDataset = nn.(Dataset)
	} else {
		m.ValidationDataset = nil
	}

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleDetails = model.LifecycleDetails

	m.Precision = model.Precision

	m.Recall = model.Recall

	m.AveragePrecision = model.AveragePrecision

	m.ConfidenceThreshold = model.ConfidenceThreshold

	m.TotalImageCount = model.TotalImageCount

	m.TestImageCount = model.TestImageCount

	m.Metrics = model.Metrics

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.ModelType = model.ModelType

	nn, e = model.TrainingDataset.UnmarshalPolymorphicJSON(model.TrainingDataset.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.TrainingDataset = nn.(Dataset)
	} else {
		m.TrainingDataset = nil
	}

	m.ModelVersion = model.ModelVersion

	m.ProjectId = model.ProjectId

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	return
}

// ModelModelTypeEnum Enum with underlying type: string
type ModelModelTypeEnum string

// Set of constants representing the allowable values for ModelModelTypeEnum
const (
	ModelModelTypeImageClassification ModelModelTypeEnum = "IMAGE_CLASSIFICATION"
	ModelModelTypeObjectDetection     ModelModelTypeEnum = "OBJECT_DETECTION"
)

var mappingModelModelTypeEnum = map[string]ModelModelTypeEnum{
	"IMAGE_CLASSIFICATION": ModelModelTypeImageClassification,
	"OBJECT_DETECTION":     ModelModelTypeObjectDetection,
}

// GetModelModelTypeEnumValues Enumerates the set of values for ModelModelTypeEnum
func GetModelModelTypeEnumValues() []ModelModelTypeEnum {
	values := make([]ModelModelTypeEnum, 0)
	for _, v := range mappingModelModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetModelModelTypeEnumStringValues Enumerates the set of values in String for ModelModelTypeEnum
func GetModelModelTypeEnumStringValues() []string {
	return []string{
		"IMAGE_CLASSIFICATION",
		"OBJECT_DETECTION",
	}
}

// GetMappingModelModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModelModelTypeEnum(val string) (ModelModelTypeEnum, bool) {
	mappingModelModelTypeEnumIgnoreCase := make(map[string]ModelModelTypeEnum)
	for k, v := range mappingModelModelTypeEnum {
		mappingModelModelTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingModelModelTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ModelLifecycleStateEnum Enum with underlying type: string
type ModelLifecycleStateEnum string

// Set of constants representing the allowable values for ModelLifecycleStateEnum
const (
	ModelLifecycleStateCreating ModelLifecycleStateEnum = "CREATING"
	ModelLifecycleStateUpdating ModelLifecycleStateEnum = "UPDATING"
	ModelLifecycleStateActive   ModelLifecycleStateEnum = "ACTIVE"
	ModelLifecycleStateDeleting ModelLifecycleStateEnum = "DELETING"
	ModelLifecycleStateDeleted  ModelLifecycleStateEnum = "DELETED"
	ModelLifecycleStateFailed   ModelLifecycleStateEnum = "FAILED"
)

var mappingModelLifecycleStateEnum = map[string]ModelLifecycleStateEnum{
	"CREATING": ModelLifecycleStateCreating,
	"UPDATING": ModelLifecycleStateUpdating,
	"ACTIVE":   ModelLifecycleStateActive,
	"DELETING": ModelLifecycleStateDeleting,
	"DELETED":  ModelLifecycleStateDeleted,
	"FAILED":   ModelLifecycleStateFailed,
}

// GetModelLifecycleStateEnumValues Enumerates the set of values for ModelLifecycleStateEnum
func GetModelLifecycleStateEnumValues() []ModelLifecycleStateEnum {
	values := make([]ModelLifecycleStateEnum, 0)
	for _, v := range mappingModelLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetModelLifecycleStateEnumStringValues Enumerates the set of values in String for ModelLifecycleStateEnum
func GetModelLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingModelLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModelLifecycleStateEnum(val string) (ModelLifecycleStateEnum, bool) {
	mappingModelLifecycleStateEnumIgnoreCase := make(map[string]ModelLifecycleStateEnum)
	for k, v := range mappingModelLifecycleStateEnum {
		mappingModelLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingModelLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
