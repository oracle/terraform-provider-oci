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

// Model Description of the a Model.
type Model struct {

	// Unique identifier model OCID of a model that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly display name for the resource. It does not have to be unique and can be modified. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)  for the model's compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate with the model.
	ProjectId *string `mandatory:"true" json:"projectId"`

	ModelDetails ModelDetails `mandatory:"true" json:"modelDetails"`

	// The time the the model was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The state of the model.
	LifecycleState ModelLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A short description of the Model.
	Description *string `mandatory:"false" json:"description"`

	// The time the model was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	TrainingDataset DatasetDetails `mandatory:"false" json:"trainingDataset"`

	EvaluationResults EvaluationResults `mandatory:"false" json:"evaluationResults"`

	TestStrategy TestStrategy `mandatory:"false" json:"testStrategy"`

	// For pre trained models this will identify model type version used for model creation
	// For custom this will identify model type version used for model creation and custom model on which training has to be done
	// <<service>>::<<service-name>>_<<model-type-version>>::<<custom model on which this training has to be done>>
	// ex: ai-lang::NER_V1::CUSTOM-V0
	Version *string `mandatory:"false" json:"version"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{ "orcl-cloud": { "free-tier-retained": "true" } }`
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
		Description       *string                           `json:"description"`
		TimeUpdated       *common.SDKTime                   `json:"timeUpdated"`
		LifecycleDetails  *string                           `json:"lifecycleDetails"`
		TrainingDataset   datasetdetails                    `json:"trainingDataset"`
		EvaluationResults evaluationresults                 `json:"evaluationResults"`
		TestStrategy      teststrategy                      `json:"testStrategy"`
		Version           *string                           `json:"version"`
		FreeformTags      map[string]string                 `json:"freeformTags"`
		DefinedTags       map[string]map[string]interface{} `json:"definedTags"`
		SystemTags        map[string]map[string]interface{} `json:"systemTags"`
		Id                *string                           `json:"id"`
		DisplayName       *string                           `json:"displayName"`
		CompartmentId     *string                           `json:"compartmentId"`
		ProjectId         *string                           `json:"projectId"`
		ModelDetails      modeldetails                      `json:"modelDetails"`
		TimeCreated       *common.SDKTime                   `json:"timeCreated"`
		LifecycleState    ModelLifecycleStateEnum           `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleDetails = model.LifecycleDetails

	nn, e = model.TrainingDataset.UnmarshalPolymorphicJSON(model.TrainingDataset.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.TrainingDataset = nn.(DatasetDetails)
	} else {
		m.TrainingDataset = nil
	}

	nn, e = model.EvaluationResults.UnmarshalPolymorphicJSON(model.EvaluationResults.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.EvaluationResults = nn.(EvaluationResults)
	} else {
		m.EvaluationResults = nil
	}

	nn, e = model.TestStrategy.UnmarshalPolymorphicJSON(model.TestStrategy.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.TestStrategy = nn.(TestStrategy)
	} else {
		m.TestStrategy = nil
	}

	m.Version = model.Version

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.ProjectId = model.ProjectId

	nn, e = model.ModelDetails.UnmarshalPolymorphicJSON(model.ModelDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ModelDetails = nn.(ModelDetails)
	} else {
		m.ModelDetails = nil
	}

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	return
}

// ModelLifecycleStateEnum Enum with underlying type: string
type ModelLifecycleStateEnum string

// Set of constants representing the allowable values for ModelLifecycleStateEnum
const (
	ModelLifecycleStateDeleting ModelLifecycleStateEnum = "DELETING"
	ModelLifecycleStateDeleted  ModelLifecycleStateEnum = "DELETED"
	ModelLifecycleStateFailed   ModelLifecycleStateEnum = "FAILED"
	ModelLifecycleStateCreating ModelLifecycleStateEnum = "CREATING"
	ModelLifecycleStateActive   ModelLifecycleStateEnum = "ACTIVE"
	ModelLifecycleStateUpdating ModelLifecycleStateEnum = "UPDATING"
)

var mappingModelLifecycleStateEnum = map[string]ModelLifecycleStateEnum{
	"DELETING": ModelLifecycleStateDeleting,
	"DELETED":  ModelLifecycleStateDeleted,
	"FAILED":   ModelLifecycleStateFailed,
	"CREATING": ModelLifecycleStateCreating,
	"ACTIVE":   ModelLifecycleStateActive,
	"UPDATING": ModelLifecycleStateUpdating,
}

var mappingModelLifecycleStateEnumLowerCase = map[string]ModelLifecycleStateEnum{
	"deleting": ModelLifecycleStateDeleting,
	"deleted":  ModelLifecycleStateDeleted,
	"failed":   ModelLifecycleStateFailed,
	"creating": ModelLifecycleStateCreating,
	"active":   ModelLifecycleStateActive,
	"updating": ModelLifecycleStateUpdating,
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
		"DELETING",
		"DELETED",
		"FAILED",
		"CREATING",
		"ACTIVE",
		"UPDATING",
	}
}

// GetMappingModelLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModelLifecycleStateEnum(val string) (ModelLifecycleStateEnum, bool) {
	enum, ok := mappingModelLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
