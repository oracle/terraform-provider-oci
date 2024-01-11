// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service API
//
// **Generative AI Service**
// OCI Generative AI is a fully managed service that provides a set of state-of-the-art, customizable LLMs that cover a wide range of use cases for text generation. Use the playground to try out the models out-of-the-box or create and host your own fine-tuned custom models based on your own data on dedicated AI clusters.
//

package generativeai

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Model **Model**
// Properties for a model.
type Model struct {

	// An ID that uniquely identifies a pretrained or fine-tuned model.
	Id *string `mandatory:"true" json:"id"`

	// The compartment OCID for fine-tuned models. For pretrained models, this value is null.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Describes what this model can be used for.
	Capabilities []ModelCapabilityEnum `mandatory:"true" json:"capabilities"`

	// The lifecycle state of a model.
	LifecycleState ModelLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time the model was created in the format of an RFC3339 datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Model type indicating whether this is a pretrained/base model or a custom/fine-tuned model.
	Type ModelTypeEnum `mandatory:"true" json:"type"`

	// An optional description of the model.
	Description *string `mandatory:"false" json:"description"`

	// A message describing the current state in more detail that can provide actionable information.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The provider of the model.
	Vendor *string `mandatory:"false" json:"vendor"`

	// The version of the model.
	Version *string `mandatory:"false" json:"version"`

	// A user-friendly name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The time the model was updated in the format of an RFC3339 datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Base model ID used for fine-tuning. For pretrained models, the value is null.
	BaseModelId *string `mandatory:"false" json:"baseModelId"`

	FineTuneDetails *FineTuneDetails `mandatory:"false" json:"fineTuneDetails"`

	ModelMetrics ModelMetrics `mandatory:"false" json:"modelMetrics"`

	// Whether a model is long-term supported. Only applicable to base models.
	IsLongTermSupported *bool `mandatory:"false" json:"isLongTermSupported"`

	// The time a model is deprecated (i.e. no longer usable for hosting or fine-tuning).
	TimeDeprecated *common.SDKTime `mandatory:"false" json:"timeDeprecated"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
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
	for _, val := range m.Capabilities {
		if _, ok := GetMappingModelCapabilityEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Capabilities: %s. Supported values are: %s.", val, strings.Join(GetModelCapabilityEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingModelLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetModelLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingModelTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetModelTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *Model) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description         *string                           `json:"description"`
		LifecycleDetails    *string                           `json:"lifecycleDetails"`
		Vendor              *string                           `json:"vendor"`
		Version             *string                           `json:"version"`
		DisplayName         *string                           `json:"displayName"`
		TimeUpdated         *common.SDKTime                   `json:"timeUpdated"`
		BaseModelId         *string                           `json:"baseModelId"`
		FineTuneDetails     *FineTuneDetails                  `json:"fineTuneDetails"`
		ModelMetrics        modelmetrics                      `json:"modelMetrics"`
		IsLongTermSupported *bool                             `json:"isLongTermSupported"`
		TimeDeprecated      *common.SDKTime                   `json:"timeDeprecated"`
		FreeformTags        map[string]string                 `json:"freeformTags"`
		DefinedTags         map[string]map[string]interface{} `json:"definedTags"`
		SystemTags          map[string]map[string]interface{} `json:"systemTags"`
		Id                  *string                           `json:"id"`
		CompartmentId       *string                           `json:"compartmentId"`
		Capabilities        []ModelCapabilityEnum             `json:"capabilities"`
		LifecycleState      ModelLifecycleStateEnum           `json:"lifecycleState"`
		TimeCreated         *common.SDKTime                   `json:"timeCreated"`
		Type                ModelTypeEnum                     `json:"type"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.LifecycleDetails = model.LifecycleDetails

	m.Vendor = model.Vendor

	m.Version = model.Version

	m.DisplayName = model.DisplayName

	m.TimeUpdated = model.TimeUpdated

	m.BaseModelId = model.BaseModelId

	m.FineTuneDetails = model.FineTuneDetails

	nn, e = model.ModelMetrics.UnmarshalPolymorphicJSON(model.ModelMetrics.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ModelMetrics = nn.(ModelMetrics)
	} else {
		m.ModelMetrics = nil
	}

	m.IsLongTermSupported = model.IsLongTermSupported

	m.TimeDeprecated = model.TimeDeprecated

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.Capabilities = make([]ModelCapabilityEnum, len(model.Capabilities))
	copy(m.Capabilities, model.Capabilities)
	m.LifecycleState = model.LifecycleState

	m.TimeCreated = model.TimeCreated

	m.Type = model.Type

	return
}

// ModelLifecycleStateEnum Enum with underlying type: string
type ModelLifecycleStateEnum string

// Set of constants representing the allowable values for ModelLifecycleStateEnum
const (
	ModelLifecycleStateActive   ModelLifecycleStateEnum = "ACTIVE"
	ModelLifecycleStateCreating ModelLifecycleStateEnum = "CREATING"
	ModelLifecycleStateDeleting ModelLifecycleStateEnum = "DELETING"
	ModelLifecycleStateDeleted  ModelLifecycleStateEnum = "DELETED"
	ModelLifecycleStateFailed   ModelLifecycleStateEnum = "FAILED"
)

var mappingModelLifecycleStateEnum = map[string]ModelLifecycleStateEnum{
	"ACTIVE":   ModelLifecycleStateActive,
	"CREATING": ModelLifecycleStateCreating,
	"DELETING": ModelLifecycleStateDeleting,
	"DELETED":  ModelLifecycleStateDeleted,
	"FAILED":   ModelLifecycleStateFailed,
}

var mappingModelLifecycleStateEnumLowerCase = map[string]ModelLifecycleStateEnum{
	"active":   ModelLifecycleStateActive,
	"creating": ModelLifecycleStateCreating,
	"deleting": ModelLifecycleStateDeleting,
	"deleted":  ModelLifecycleStateDeleted,
	"failed":   ModelLifecycleStateFailed,
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
		"ACTIVE",
		"CREATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingModelLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModelLifecycleStateEnum(val string) (ModelLifecycleStateEnum, bool) {
	enum, ok := mappingModelLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ModelTypeEnum Enum with underlying type: string
type ModelTypeEnum string

// Set of constants representing the allowable values for ModelTypeEnum
const (
	ModelTypeBase   ModelTypeEnum = "BASE"
	ModelTypeCustom ModelTypeEnum = "CUSTOM"
)

var mappingModelTypeEnum = map[string]ModelTypeEnum{
	"BASE":   ModelTypeBase,
	"CUSTOM": ModelTypeCustom,
}

var mappingModelTypeEnumLowerCase = map[string]ModelTypeEnum{
	"base":   ModelTypeBase,
	"custom": ModelTypeCustom,
}

// GetModelTypeEnumValues Enumerates the set of values for ModelTypeEnum
func GetModelTypeEnumValues() []ModelTypeEnum {
	values := make([]ModelTypeEnum, 0)
	for _, v := range mappingModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetModelTypeEnumStringValues Enumerates the set of values in String for ModelTypeEnum
func GetModelTypeEnumStringValues() []string {
	return []string{
		"BASE",
		"CUSTOM",
	}
}

// GetMappingModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModelTypeEnum(val string) (ModelTypeEnum, bool) {
	enum, ok := mappingModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
