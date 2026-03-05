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

// SemanticStore A Semantic Store is a container resource of semantic records, with controllable enrichment refresh and synchronization policy.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator who gives OCI resource access to users. See
// Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm) and Getting Access to Generative AI Resources (https://docs.oracle.com/iaas/Content/generative-ai/iam-policies.htm).
type SemanticStore struct {

	// An OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that uniquely identifies a SemanticStore.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Owning compartment OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for a SemanticStore.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time that the SemanticStore was created in the format of an RFC3339 datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time that the SemanticStore was updated in the format of an RFC3339 datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The lifecycle state of a SemanticStore.
	LifecycleState SemanticStoreLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	DataSource DataSourceDetails `mandatory:"true" json:"dataSource"`

	Schemas SchemasDetails `mandatory:"true" json:"schemas"`

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

	// A message describing the current state in more detail that can provide actionable information.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	RefreshSchedule RefreshScheduleDetails `mandatory:"false" json:"refreshSchedule"`
}

func (m SemanticStore) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SemanticStore) ValidateEnumValue() (bool, error) {
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
func (m *SemanticStore) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description      *string                           `json:"description"`
		LifecycleDetails *string                           `json:"lifecycleDetails"`
		RefreshSchedule  refreshscheduledetails            `json:"refreshSchedule"`
		Id               *string                           `json:"id"`
		DisplayName      *string                           `json:"displayName"`
		CompartmentId    *string                           `json:"compartmentId"`
		TimeCreated      *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated      *common.SDKTime                   `json:"timeUpdated"`
		LifecycleState   SemanticStoreLifecycleStateEnum   `json:"lifecycleState"`
		DataSource       datasourcedetails                 `json:"dataSource"`
		Schemas          schemasdetails                    `json:"schemas"`
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

	m.LifecycleDetails = model.LifecycleDetails

	nn, e = model.RefreshSchedule.UnmarshalPolymorphicJSON(model.RefreshSchedule.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.RefreshSchedule = nn.(RefreshScheduleDetails)
	} else {
		m.RefreshSchedule = nil
	}

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

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

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	return
}

// SemanticStoreLifecycleStateEnum Enum with underlying type: string
type SemanticStoreLifecycleStateEnum string

// Set of constants representing the allowable values for SemanticStoreLifecycleStateEnum
const (
	SemanticStoreLifecycleStateActive   SemanticStoreLifecycleStateEnum = "ACTIVE"
	SemanticStoreLifecycleStateCreating SemanticStoreLifecycleStateEnum = "CREATING"
	SemanticStoreLifecycleStateUpdating SemanticStoreLifecycleStateEnum = "UPDATING"
	SemanticStoreLifecycleStateDeleting SemanticStoreLifecycleStateEnum = "DELETING"
	SemanticStoreLifecycleStateDeleted  SemanticStoreLifecycleStateEnum = "DELETED"
	SemanticStoreLifecycleStateFailed   SemanticStoreLifecycleStateEnum = "FAILED"
)

var mappingSemanticStoreLifecycleStateEnum = map[string]SemanticStoreLifecycleStateEnum{
	"ACTIVE":   SemanticStoreLifecycleStateActive,
	"CREATING": SemanticStoreLifecycleStateCreating,
	"UPDATING": SemanticStoreLifecycleStateUpdating,
	"DELETING": SemanticStoreLifecycleStateDeleting,
	"DELETED":  SemanticStoreLifecycleStateDeleted,
	"FAILED":   SemanticStoreLifecycleStateFailed,
}

var mappingSemanticStoreLifecycleStateEnumLowerCase = map[string]SemanticStoreLifecycleStateEnum{
	"active":   SemanticStoreLifecycleStateActive,
	"creating": SemanticStoreLifecycleStateCreating,
	"updating": SemanticStoreLifecycleStateUpdating,
	"deleting": SemanticStoreLifecycleStateDeleting,
	"deleted":  SemanticStoreLifecycleStateDeleted,
	"failed":   SemanticStoreLifecycleStateFailed,
}

// GetSemanticStoreLifecycleStateEnumValues Enumerates the set of values for SemanticStoreLifecycleStateEnum
func GetSemanticStoreLifecycleStateEnumValues() []SemanticStoreLifecycleStateEnum {
	values := make([]SemanticStoreLifecycleStateEnum, 0)
	for _, v := range mappingSemanticStoreLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSemanticStoreLifecycleStateEnumStringValues Enumerates the set of values in String for SemanticStoreLifecycleStateEnum
func GetSemanticStoreLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"CREATING",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingSemanticStoreLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSemanticStoreLifecycleStateEnum(val string) (SemanticStoreLifecycleStateEnum, bool) {
	enum, ok := mappingSemanticStoreLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
