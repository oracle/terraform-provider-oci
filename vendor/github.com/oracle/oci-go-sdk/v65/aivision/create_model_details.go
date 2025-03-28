// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vision API
//
// Using Vision, you can upload images to detect and classify objects in them. If you have lots of images, you can process them in batch using asynchronous API endpoints. Vision's features are thematically split between Document AI for document-centric images, and Image Analysis for object and scene-based images. Pretrained models and custom models are supported.
//

package aivision

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateModelDetails The information needed to create a new model.
type CreateModelDetails struct {

	// Which type of Vision model this is.
	ModelType ModelModelTypeEnum `mandatory:"true" json:"modelType"`

	// The compartment identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	TrainingDataset Dataset `mandatory:"true" json:"trainingDataset"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project that contains the model.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// A human-friendly name for the model, which can be changed.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// An optional description of the model.
	Description *string `mandatory:"false" json:"description"`

	// The model version
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	// Set to true when experimenting with a new model type or dataset, so the model training is quick, with a predefined low number of passes through the training data.
	IsQuickMode *bool `mandatory:"false" json:"isQuickMode"`

	// The maximum model training duration in hours, expressed as a decimal fraction.
	MaxTrainingDurationInHours *float64 `mandatory:"false" json:"maxTrainingDurationInHours"`

	TestingDataset Dataset `mandatory:"false" json:"testingDataset"`

	ValidationDataset Dataset `mandatory:"false" json:"validationDataset"`

	// A simple key-value pair that is applied without any predefined name, type, or scope. It exists for cross-compatibility only.
	// For example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateModelDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateModelDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingModelModelTypeEnum(string(m.ModelType)); !ok && m.ModelType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ModelType: %s. Supported values are: %s.", m.ModelType, strings.Join(GetModelModelTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateModelDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName                *string                           `json:"displayName"`
		Description                *string                           `json:"description"`
		ModelVersion               *string                           `json:"modelVersion"`
		IsQuickMode                *bool                             `json:"isQuickMode"`
		MaxTrainingDurationInHours *float64                          `json:"maxTrainingDurationInHours"`
		TestingDataset             dataset                           `json:"testingDataset"`
		ValidationDataset          dataset                           `json:"validationDataset"`
		FreeformTags               map[string]string                 `json:"freeformTags"`
		DefinedTags                map[string]map[string]interface{} `json:"definedTags"`
		ModelType                  ModelModelTypeEnum                `json:"modelType"`
		CompartmentId              *string                           `json:"compartmentId"`
		TrainingDataset            dataset                           `json:"trainingDataset"`
		ProjectId                  *string                           `json:"projectId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.ModelVersion = model.ModelVersion

	m.IsQuickMode = model.IsQuickMode

	m.MaxTrainingDurationInHours = model.MaxTrainingDurationInHours

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

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.ModelType = model.ModelType

	m.CompartmentId = model.CompartmentId

	nn, e = model.TrainingDataset.UnmarshalPolymorphicJSON(model.TrainingDataset.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.TrainingDataset = nn.(Dataset)
	} else {
		m.TrainingDataset = nil
	}

	m.ProjectId = model.ProjectId

	return
}
