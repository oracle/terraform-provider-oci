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

// DedicatedAiCluster Dedicated AI clusters are compute resources that you can use for fine-tuning custom models or for hosting endpoints for custom models. The clusters are dedicated to your models and not shared with users in other tenancies.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator who gives OCI resource access to users. See
// Getting Started with Policies (https://docs.cloud.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm) and Getting Access to Generative AI Resouces (https://docs.cloud.oracle.com/iaas/Content/generative-ai/iam-policies.htm).
type DedicatedAiCluster struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the dedicated AI cluster.
	Id *string `mandatory:"true" json:"id"`

	// The dedicated AI cluster type indicating whether this is a fine-tuning/training processor or hosting/inference processor.
	Type DedicatedAiClusterTypeEnum `mandatory:"true" json:"type"`

	// The compartment OCID to create the dedicated AI cluster in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time the dedicated AI cluster was created, in the format defined by RFC 3339
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the dedicated AI cluster.
	LifecycleState DedicatedAiClusterLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The number of dedicated units in this AI cluster.
	UnitCount *int `mandatory:"true" json:"unitCount"`

	// The shape of dedicated unit in this AI cluster. The underlying hardware configuration is hidden from customers.
	UnitShape DedicatedAiClusterUnitShapeEnum `mandatory:"true" json:"unitShape"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// An optional description of the dedicated AI cluster.
	Description *string `mandatory:"false" json:"description"`

	// The date and time the dedicated AI cluster was updated, in the format defined by RFC 3339
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state with detail that can provide actionable information.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	Capacity DedicatedAiClusterCapacity `mandatory:"false" json:"capacity"`

	PreviousState *DedicatedAiCluster `mandatory:"false" json:"previousState"`

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

func (m DedicatedAiCluster) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DedicatedAiCluster) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDedicatedAiClusterTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetDedicatedAiClusterTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDedicatedAiClusterLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDedicatedAiClusterLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDedicatedAiClusterUnitShapeEnum(string(m.UnitShape)); !ok && m.UnitShape != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UnitShape: %s. Supported values are: %s.", m.UnitShape, strings.Join(GetDedicatedAiClusterUnitShapeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *DedicatedAiCluster) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName      *string                              `json:"displayName"`
		Description      *string                              `json:"description"`
		TimeUpdated      *common.SDKTime                      `json:"timeUpdated"`
		LifecycleDetails *string                              `json:"lifecycleDetails"`
		Capacity         dedicatedaiclustercapacity           `json:"capacity"`
		PreviousState    *DedicatedAiCluster                  `json:"previousState"`
		FreeformTags     map[string]string                    `json:"freeformTags"`
		DefinedTags      map[string]map[string]interface{}    `json:"definedTags"`
		SystemTags       map[string]map[string]interface{}    `json:"systemTags"`
		Id               *string                              `json:"id"`
		Type             DedicatedAiClusterTypeEnum           `json:"type"`
		CompartmentId    *string                              `json:"compartmentId"`
		TimeCreated      *common.SDKTime                      `json:"timeCreated"`
		LifecycleState   DedicatedAiClusterLifecycleStateEnum `json:"lifecycleState"`
		UnitCount        *int                                 `json:"unitCount"`
		UnitShape        DedicatedAiClusterUnitShapeEnum      `json:"unitShape"`
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

	nn, e = model.Capacity.UnmarshalPolymorphicJSON(model.Capacity.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Capacity = nn.(DedicatedAiClusterCapacity)
	} else {
		m.Capacity = nil
	}

	m.PreviousState = model.PreviousState

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.Type = model.Type

	m.CompartmentId = model.CompartmentId

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	m.UnitCount = model.UnitCount

	m.UnitShape = model.UnitShape

	return
}

// DedicatedAiClusterTypeEnum Enum with underlying type: string
type DedicatedAiClusterTypeEnum string

// Set of constants representing the allowable values for DedicatedAiClusterTypeEnum
const (
	DedicatedAiClusterTypeHosting    DedicatedAiClusterTypeEnum = "HOSTING"
	DedicatedAiClusterTypeFineTuning DedicatedAiClusterTypeEnum = "FINE_TUNING"
)

var mappingDedicatedAiClusterTypeEnum = map[string]DedicatedAiClusterTypeEnum{
	"HOSTING":     DedicatedAiClusterTypeHosting,
	"FINE_TUNING": DedicatedAiClusterTypeFineTuning,
}

var mappingDedicatedAiClusterTypeEnumLowerCase = map[string]DedicatedAiClusterTypeEnum{
	"hosting":     DedicatedAiClusterTypeHosting,
	"fine_tuning": DedicatedAiClusterTypeFineTuning,
}

// GetDedicatedAiClusterTypeEnumValues Enumerates the set of values for DedicatedAiClusterTypeEnum
func GetDedicatedAiClusterTypeEnumValues() []DedicatedAiClusterTypeEnum {
	values := make([]DedicatedAiClusterTypeEnum, 0)
	for _, v := range mappingDedicatedAiClusterTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDedicatedAiClusterTypeEnumStringValues Enumerates the set of values in String for DedicatedAiClusterTypeEnum
func GetDedicatedAiClusterTypeEnumStringValues() []string {
	return []string{
		"HOSTING",
		"FINE_TUNING",
	}
}

// GetMappingDedicatedAiClusterTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDedicatedAiClusterTypeEnum(val string) (DedicatedAiClusterTypeEnum, bool) {
	enum, ok := mappingDedicatedAiClusterTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DedicatedAiClusterLifecycleStateEnum Enum with underlying type: string
type DedicatedAiClusterLifecycleStateEnum string

// Set of constants representing the allowable values for DedicatedAiClusterLifecycleStateEnum
const (
	DedicatedAiClusterLifecycleStateCreating       DedicatedAiClusterLifecycleStateEnum = "CREATING"
	DedicatedAiClusterLifecycleStateActive         DedicatedAiClusterLifecycleStateEnum = "ACTIVE"
	DedicatedAiClusterLifecycleStateUpdating       DedicatedAiClusterLifecycleStateEnum = "UPDATING"
	DedicatedAiClusterLifecycleStateDeleting       DedicatedAiClusterLifecycleStateEnum = "DELETING"
	DedicatedAiClusterLifecycleStateDeleted        DedicatedAiClusterLifecycleStateEnum = "DELETED"
	DedicatedAiClusterLifecycleStateFailed         DedicatedAiClusterLifecycleStateEnum = "FAILED"
	DedicatedAiClusterLifecycleStateNeedsAttention DedicatedAiClusterLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingDedicatedAiClusterLifecycleStateEnum = map[string]DedicatedAiClusterLifecycleStateEnum{
	"CREATING":        DedicatedAiClusterLifecycleStateCreating,
	"ACTIVE":          DedicatedAiClusterLifecycleStateActive,
	"UPDATING":        DedicatedAiClusterLifecycleStateUpdating,
	"DELETING":        DedicatedAiClusterLifecycleStateDeleting,
	"DELETED":         DedicatedAiClusterLifecycleStateDeleted,
	"FAILED":          DedicatedAiClusterLifecycleStateFailed,
	"NEEDS_ATTENTION": DedicatedAiClusterLifecycleStateNeedsAttention,
}

var mappingDedicatedAiClusterLifecycleStateEnumLowerCase = map[string]DedicatedAiClusterLifecycleStateEnum{
	"creating":        DedicatedAiClusterLifecycleStateCreating,
	"active":          DedicatedAiClusterLifecycleStateActive,
	"updating":        DedicatedAiClusterLifecycleStateUpdating,
	"deleting":        DedicatedAiClusterLifecycleStateDeleting,
	"deleted":         DedicatedAiClusterLifecycleStateDeleted,
	"failed":          DedicatedAiClusterLifecycleStateFailed,
	"needs_attention": DedicatedAiClusterLifecycleStateNeedsAttention,
}

// GetDedicatedAiClusterLifecycleStateEnumValues Enumerates the set of values for DedicatedAiClusterLifecycleStateEnum
func GetDedicatedAiClusterLifecycleStateEnumValues() []DedicatedAiClusterLifecycleStateEnum {
	values := make([]DedicatedAiClusterLifecycleStateEnum, 0)
	for _, v := range mappingDedicatedAiClusterLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDedicatedAiClusterLifecycleStateEnumStringValues Enumerates the set of values in String for DedicatedAiClusterLifecycleStateEnum
func GetDedicatedAiClusterLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingDedicatedAiClusterLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDedicatedAiClusterLifecycleStateEnum(val string) (DedicatedAiClusterLifecycleStateEnum, bool) {
	enum, ok := mappingDedicatedAiClusterLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DedicatedAiClusterUnitShapeEnum Enum with underlying type: string
type DedicatedAiClusterUnitShapeEnum string

// Set of constants representing the allowable values for DedicatedAiClusterUnitShapeEnum
const (
	DedicatedAiClusterUnitShapeLargeCohere    DedicatedAiClusterUnitShapeEnum = "LARGE_COHERE"
	DedicatedAiClusterUnitShapeLargeCohereV2  DedicatedAiClusterUnitShapeEnum = "LARGE_COHERE_V2"
	DedicatedAiClusterUnitShapeSmallCohere    DedicatedAiClusterUnitShapeEnum = "SMALL_COHERE"
	DedicatedAiClusterUnitShapeSmallCohereV2  DedicatedAiClusterUnitShapeEnum = "SMALL_COHERE_V2"
	DedicatedAiClusterUnitShapeEmbedCohere    DedicatedAiClusterUnitShapeEnum = "EMBED_COHERE"
	DedicatedAiClusterUnitShapeLlama270       DedicatedAiClusterUnitShapeEnum = "LLAMA2_70"
	DedicatedAiClusterUnitShapeLargeGeneric   DedicatedAiClusterUnitShapeEnum = "LARGE_GENERIC"
	DedicatedAiClusterUnitShapeLargeCohereV22 DedicatedAiClusterUnitShapeEnum = "LARGE_COHERE_V2_2"
	DedicatedAiClusterUnitShapeLargeGeneric4  DedicatedAiClusterUnitShapeEnum = "LARGE_GENERIC_4"
)

var mappingDedicatedAiClusterUnitShapeEnum = map[string]DedicatedAiClusterUnitShapeEnum{
	"LARGE_COHERE":      DedicatedAiClusterUnitShapeLargeCohere,
	"LARGE_COHERE_V2":   DedicatedAiClusterUnitShapeLargeCohereV2,
	"SMALL_COHERE":      DedicatedAiClusterUnitShapeSmallCohere,
	"SMALL_COHERE_V2":   DedicatedAiClusterUnitShapeSmallCohereV2,
	"EMBED_COHERE":      DedicatedAiClusterUnitShapeEmbedCohere,
	"LLAMA2_70":         DedicatedAiClusterUnitShapeLlama270,
	"LARGE_GENERIC":     DedicatedAiClusterUnitShapeLargeGeneric,
	"LARGE_COHERE_V2_2": DedicatedAiClusterUnitShapeLargeCohereV22,
	"LARGE_GENERIC_4":   DedicatedAiClusterUnitShapeLargeGeneric4,
}

var mappingDedicatedAiClusterUnitShapeEnumLowerCase = map[string]DedicatedAiClusterUnitShapeEnum{
	"large_cohere":      DedicatedAiClusterUnitShapeLargeCohere,
	"large_cohere_v2":   DedicatedAiClusterUnitShapeLargeCohereV2,
	"small_cohere":      DedicatedAiClusterUnitShapeSmallCohere,
	"small_cohere_v2":   DedicatedAiClusterUnitShapeSmallCohereV2,
	"embed_cohere":      DedicatedAiClusterUnitShapeEmbedCohere,
	"llama2_70":         DedicatedAiClusterUnitShapeLlama270,
	"large_generic":     DedicatedAiClusterUnitShapeLargeGeneric,
	"large_cohere_v2_2": DedicatedAiClusterUnitShapeLargeCohereV22,
	"large_generic_4":   DedicatedAiClusterUnitShapeLargeGeneric4,
}

// GetDedicatedAiClusterUnitShapeEnumValues Enumerates the set of values for DedicatedAiClusterUnitShapeEnum
func GetDedicatedAiClusterUnitShapeEnumValues() []DedicatedAiClusterUnitShapeEnum {
	values := make([]DedicatedAiClusterUnitShapeEnum, 0)
	for _, v := range mappingDedicatedAiClusterUnitShapeEnum {
		values = append(values, v)
	}
	return values
}

// GetDedicatedAiClusterUnitShapeEnumStringValues Enumerates the set of values in String for DedicatedAiClusterUnitShapeEnum
func GetDedicatedAiClusterUnitShapeEnumStringValues() []string {
	return []string{
		"LARGE_COHERE",
		"LARGE_COHERE_V2",
		"SMALL_COHERE",
		"SMALL_COHERE_V2",
		"EMBED_COHERE",
		"LLAMA2_70",
		"LARGE_GENERIC",
		"LARGE_COHERE_V2_2",
		"LARGE_GENERIC_4",
	}
}

// GetMappingDedicatedAiClusterUnitShapeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDedicatedAiClusterUnitShapeEnum(val string) (DedicatedAiClusterUnitShapeEnum, bool) {
	enum, ok := mappingDedicatedAiClusterUnitShapeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
