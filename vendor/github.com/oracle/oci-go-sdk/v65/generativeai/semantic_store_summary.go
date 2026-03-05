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

// SemanticStoreSummary Summary information for a SemanticStore.
type SemanticStoreSummary struct {

	// An OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that uniquely identifies a SemanticStore.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Owning compartment OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for a SemanticStore.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	DataSource DataSourceDetails `mandatory:"true" json:"dataSource"`

	Schemas SchemasDetails `mandatory:"true" json:"schemas"`

	// The date and time that the semanticStore was created in the format of an RFC3339 datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time that the semanticStore was updated in the format of an RFC3339 datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the SemanticStore.
	// Allowed values are:
	// - ACTIVE
	// - CREATING
	// - UPDATING
	// - DELETING
	// - DELETED
	// - FAILED
	LifecycleState SemanticStoreLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"true" json:"systemTags"`

	// An optional description of the SemanticStore.
	Description *string `mandatory:"false" json:"description"`

	RefreshSchedule RefreshScheduleDetails `mandatory:"false" json:"refreshSchedule"`

	// A message describing the current state in more detail that can provide actionable information.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m SemanticStoreSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SemanticStoreSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSemanticStoreLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSemanticStoreLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *SemanticStoreSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description      *string                           `json:"description"`
		RefreshSchedule  refreshscheduledetails            `json:"refreshSchedule"`
		LifecycleDetails *string                           `json:"lifecycleDetails"`
		Id               *string                           `json:"id"`
		DisplayName      *string                           `json:"displayName"`
		CompartmentId    *string                           `json:"compartmentId"`
		DataSource       datasourcedetails                 `json:"dataSource"`
		Schemas          schemasdetails                    `json:"schemas"`
		TimeCreated      *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated      *common.SDKTime                   `json:"timeUpdated"`
		LifecycleState   SemanticStoreLifecycleStateEnum   `json:"lifecycleState"`
		FreeformTags     map[string]string                 `json:"freeformTags"`
		DefinedTags      map[string]map[string]interface{} `json:"definedTags"`
		SystemTags       map[string]map[string]interface{} `json:"systemTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	nn, e = model.RefreshSchedule.UnmarshalPolymorphicJSON(model.RefreshSchedule.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.RefreshSchedule = nn.(RefreshScheduleDetails)
	} else {
		m.RefreshSchedule = nil
	}

	m.LifecycleDetails = model.LifecycleDetails

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	nn, e = model.DataSource.UnmarshalPolymorphicJSON(model.DataSource.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DataSource = nn.(DataSourceDetails)
	} else {
		m.DataSource = nil
	}

	nn, e = model.Schemas.UnmarshalPolymorphicJSON(model.Schemas.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Schemas = nn.(SchemasDetails)
	} else {
		m.Schemas = nil
	}

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	return
}
