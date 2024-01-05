// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Document Understanding API
//
// Document AI helps customers perform various analysis on their documents. If a customer has lots of documents, they can process them in batch using asynchronous API endpoints.
//

package aidocument

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ModelSummary The metadata about the model.
type ModelSummary struct {

	// A unique identifier that is immutable after creation.
	Id *string `mandatory:"true" json:"id"`

	// The compartment identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The type of the Document model.
	ModelType ModelModelTypeEnum `mandatory:"true" json:"modelType"`

	// The version of the model.
	ModelVersion *string `mandatory:"true" json:"modelVersion"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project that contains the model.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// When the model was created, as an RFC3339 datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the model.
	LifecycleState ModelLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A human-friendly name for the model, which can be changed.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// An optional description of the model.
	Description *string `mandatory:"false" json:"description"`

	// When the model was modified, as an RFC3339 datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail, that can provide actionable information if training failed.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The precision of the trained model.
	Precision *float32 `mandatory:"false" json:"precision"`

	// The tenancy id of the model.
	TenancyId *string `mandatory:"false" json:"tenancyId"`

	// the alias name of the model.
	AliasName *string `mandatory:"false" json:"aliasName"`

	TrainingDataset Dataset `mandatory:"false" json:"trainingDataset"`

	TestingDataset Dataset `mandatory:"false" json:"testingDataset"`

	ValidationDataset Dataset `mandatory:"false" json:"validationDataset"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) list of active custom Key Value models that need to be composed.
	ComponentModels []ComponentModel `mandatory:"false" json:"componentModels"`

	// Set to true when the model is created by using multiple key value extraction models.
	IsComposedModel *bool `mandatory:"false" json:"isComposedModel"`

	// A simple key-value pair that is applied without any predefined name, type, or scope. It exists for cross-compatibility only.
	// For example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// For example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m ModelSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ModelSummary) ValidateEnumValue() (bool, error) {
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
func (m *ModelSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName       *string                           `json:"displayName"`
		Description       *string                           `json:"description"`
		TimeUpdated       *common.SDKTime                   `json:"timeUpdated"`
		LifecycleDetails  *string                           `json:"lifecycleDetails"`
		Precision         *float32                          `json:"precision"`
		TenancyId         *string                           `json:"tenancyId"`
		AliasName         *string                           `json:"aliasName"`
		TrainingDataset   dataset                           `json:"trainingDataset"`
		TestingDataset    dataset                           `json:"testingDataset"`
		ValidationDataset dataset                           `json:"validationDataset"`
		ComponentModels   []ComponentModel                  `json:"componentModels"`
		IsComposedModel   *bool                             `json:"isComposedModel"`
		FreeformTags      map[string]string                 `json:"freeformTags"`
		DefinedTags       map[string]map[string]interface{} `json:"definedTags"`
		SystemTags        map[string]map[string]interface{} `json:"systemTags"`
		Id                *string                           `json:"id"`
		CompartmentId     *string                           `json:"compartmentId"`
		ModelType         ModelModelTypeEnum                `json:"modelType"`
		ModelVersion      *string                           `json:"modelVersion"`
		ProjectId         *string                           `json:"projectId"`
		TimeCreated       *common.SDKTime                   `json:"timeCreated"`
		LifecycleState    ModelLifecycleStateEnum           `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleDetails = model.LifecycleDetails

	m.Precision = model.Precision

	m.TenancyId = model.TenancyId

	m.AliasName = model.AliasName

	nn, e = model.TrainingDataset.UnmarshalPolymorphicJSON(model.TrainingDataset.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.TrainingDataset = nn.(Dataset)
	} else {
		m.TrainingDataset = nil
	}

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

	m.ComponentModels = make([]ComponentModel, len(model.ComponentModels))
	copy(m.ComponentModels, model.ComponentModels)
	m.IsComposedModel = model.IsComposedModel

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.ModelType = model.ModelType

	m.ModelVersion = model.ModelVersion

	m.ProjectId = model.ProjectId

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	return
}
