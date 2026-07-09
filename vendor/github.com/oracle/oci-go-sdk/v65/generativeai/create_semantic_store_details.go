// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// CreateSemanticStoreDetails The data to create a SemanticStore.
type CreateSemanticStoreDetails struct {

	// Owning compartment OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for a SemanticStore.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	DataSource CreateDataSourceDetails `mandatory:"true" json:"dataSource"`

	Schemas CreateSchemasDetails `mandatory:"true" json:"schemas"`

	// An optional description of the SemanticStore.
	Description *string `mandatory:"false" json:"description"`

	// Controls which enrichment inputs are enabled for the semantic store.
	// Allowed values are:
	// - COMBINED
	// - METADATA_ONLY
	// - ANNOTATION_ONLY
	// Defaults to COMBINED.
	EnrichmentMode CreateSemanticStoreDetailsEnrichmentModeEnum `mandatory:"false" json:"enrichmentMode,omitempty"`

	ModelSelection SemanticStoreModelSelection `mandatory:"false" json:"modelSelection"`

	RefreshSchedule RefreshScheduleDetails `mandatory:"false" json:"refreshSchedule"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateSemanticStoreDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateSemanticStoreDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateSemanticStoreDetailsEnrichmentModeEnum(string(m.EnrichmentMode)); !ok && m.EnrichmentMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EnrichmentMode: %s. Supported values are: %s.", m.EnrichmentMode, strings.Join(GetCreateSemanticStoreDetailsEnrichmentModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateSemanticStoreDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description     *string                                      `json:"description"`
		EnrichmentMode  CreateSemanticStoreDetailsEnrichmentModeEnum `json:"enrichmentMode"`
		ModelSelection  semanticstoremodelselection                  `json:"modelSelection"`
		RefreshSchedule refreshscheduledetails                       `json:"refreshSchedule"`
		FreeformTags    map[string]string                            `json:"freeformTags"`
		DefinedTags     map[string]map[string]interface{}            `json:"definedTags"`
		CompartmentId   *string                                      `json:"compartmentId"`
		DisplayName     *string                                      `json:"displayName"`
		DataSource      createdatasourcedetails                      `json:"dataSource"`
		Schemas         createschemasdetails                         `json:"schemas"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.EnrichmentMode = model.EnrichmentMode

	nn, e = model.ModelSelection.UnmarshalPolymorphicJSON(model.ModelSelection.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ModelSelection = nn.(SemanticStoreModelSelection)
	} else {
		m.ModelSelection = nil
	}

	nn, e = model.RefreshSchedule.UnmarshalPolymorphicJSON(model.RefreshSchedule.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.RefreshSchedule = nn.(RefreshScheduleDetails)
	} else {
		m.RefreshSchedule = nil
	}

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

	nn, e = model.DataSource.UnmarshalPolymorphicJSON(model.DataSource.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DataSource = nn.(CreateDataSourceDetails)
	} else {
		m.DataSource = nil
	}

	nn, e = model.Schemas.UnmarshalPolymorphicJSON(model.Schemas.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Schemas = nn.(CreateSchemasDetails)
	} else {
		m.Schemas = nil
	}

	return
}

// CreateSemanticStoreDetailsEnrichmentModeEnum Enum with underlying type: string
type CreateSemanticStoreDetailsEnrichmentModeEnum string

// Set of constants representing the allowable values for CreateSemanticStoreDetailsEnrichmentModeEnum
const (
	CreateSemanticStoreDetailsEnrichmentModeCombined       CreateSemanticStoreDetailsEnrichmentModeEnum = "COMBINED"
	CreateSemanticStoreDetailsEnrichmentModeMetadataOnly   CreateSemanticStoreDetailsEnrichmentModeEnum = "METADATA_ONLY"
	CreateSemanticStoreDetailsEnrichmentModeAnnotationOnly CreateSemanticStoreDetailsEnrichmentModeEnum = "ANNOTATION_ONLY"
)

var mappingCreateSemanticStoreDetailsEnrichmentModeEnum = map[string]CreateSemanticStoreDetailsEnrichmentModeEnum{
	"COMBINED":        CreateSemanticStoreDetailsEnrichmentModeCombined,
	"METADATA_ONLY":   CreateSemanticStoreDetailsEnrichmentModeMetadataOnly,
	"ANNOTATION_ONLY": CreateSemanticStoreDetailsEnrichmentModeAnnotationOnly,
}

var mappingCreateSemanticStoreDetailsEnrichmentModeEnumLowerCase = map[string]CreateSemanticStoreDetailsEnrichmentModeEnum{
	"combined":        CreateSemanticStoreDetailsEnrichmentModeCombined,
	"metadata_only":   CreateSemanticStoreDetailsEnrichmentModeMetadataOnly,
	"annotation_only": CreateSemanticStoreDetailsEnrichmentModeAnnotationOnly,
}

// GetCreateSemanticStoreDetailsEnrichmentModeEnumValues Enumerates the set of values for CreateSemanticStoreDetailsEnrichmentModeEnum
func GetCreateSemanticStoreDetailsEnrichmentModeEnumValues() []CreateSemanticStoreDetailsEnrichmentModeEnum {
	values := make([]CreateSemanticStoreDetailsEnrichmentModeEnum, 0)
	for _, v := range mappingCreateSemanticStoreDetailsEnrichmentModeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateSemanticStoreDetailsEnrichmentModeEnumStringValues Enumerates the set of values in String for CreateSemanticStoreDetailsEnrichmentModeEnum
func GetCreateSemanticStoreDetailsEnrichmentModeEnumStringValues() []string {
	return []string{
		"COMBINED",
		"METADATA_ONLY",
		"ANNOTATION_ONLY",
	}
}

// GetMappingCreateSemanticStoreDetailsEnrichmentModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateSemanticStoreDetailsEnrichmentModeEnum(val string) (CreateSemanticStoreDetailsEnrichmentModeEnum, bool) {
	enum, ok := mappingCreateSemanticStoreDetailsEnrichmentModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
