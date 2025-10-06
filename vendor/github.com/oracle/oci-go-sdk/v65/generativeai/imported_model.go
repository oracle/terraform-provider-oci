// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// ImportedModel Represents a model imported into the system based on an external data source, such as Hugging Face or Object Storage.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator who gives OCI resource access to users. See
// Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm) and Getting Access to Generative AI Resources (https://docs.oracle.com/iaas/Content/generative-ai/iam-policies.htm).
type ImportedModel struct {

	// An OCID that uniquely identifies an imported model.
	Id *string `mandatory:"true" json:"id"`

	// The compartment OCID from which the model is imported.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The lifecycle state of the imported model.
	LifecycleState ImportedModelLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	DataSource ModelDataSource `mandatory:"true" json:"dataSource"`

	// The date and time that the imported model was created in the format of an RFC3339 datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// A user-friendly name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// An optional description of the imported model.
	Description *string `mandatory:"false" json:"description"`

	// Specifies the intended use or supported capabilities of the imported model.
	Capabilities []ImportedModelCapabilityEnum `mandatory:"false" json:"capabilities,omitempty"`

	// Additional information about the current state of the imported model, providing more detailed and actionable context.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The provider of the imported model.
	Vendor *string `mandatory:"false" json:"vendor"`

	// The version of the imported model.
	Version *string `mandatory:"false" json:"version"`

	// The date and time that the imported model was updated in the format of an RFC3339 datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	PreviousState *ImportedModel `mandatory:"false" json:"previousState"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m ImportedModel) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ImportedModel) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingImportedModelLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetImportedModelLifecycleStateEnumStringValues(), ",")))
	}

	for _, val := range m.Capabilities {
		if _, ok := GetMappingImportedModelCapabilityEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Capabilities: %s. Supported values are: %s.", val, strings.Join(GetImportedModelCapabilityEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ImportedModel) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName      *string                           `json:"displayName"`
		Description      *string                           `json:"description"`
		Capabilities     []ImportedModelCapabilityEnum     `json:"capabilities"`
		LifecycleDetails *string                           `json:"lifecycleDetails"`
		Vendor           *string                           `json:"vendor"`
		Version          *string                           `json:"version"`
		TimeUpdated      *common.SDKTime                   `json:"timeUpdated"`
		PreviousState    *ImportedModel                    `json:"previousState"`
		FreeformTags     map[string]string                 `json:"freeformTags"`
		DefinedTags      map[string]map[string]interface{} `json:"definedTags"`
		SystemTags       map[string]map[string]interface{} `json:"systemTags"`
		Id               *string                           `json:"id"`
		CompartmentId    *string                           `json:"compartmentId"`
		LifecycleState   ImportedModelLifecycleStateEnum   `json:"lifecycleState"`
		DataSource       modeldatasource                   `json:"dataSource"`
		TimeCreated      *common.SDKTime                   `json:"timeCreated"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.Capabilities = make([]ImportedModelCapabilityEnum, len(model.Capabilities))
	copy(m.Capabilities, model.Capabilities)
	m.LifecycleDetails = model.LifecycleDetails

	m.Vendor = model.Vendor

	m.Version = model.Version

	m.TimeUpdated = model.TimeUpdated

	m.PreviousState = model.PreviousState

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.LifecycleState = model.LifecycleState

	nn, e = model.DataSource.UnmarshalPolymorphicJSON(model.DataSource.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DataSource = nn.(ModelDataSource)
	} else {
		m.DataSource = nil
	}

	m.TimeCreated = model.TimeCreated

	return
}

// ImportedModelLifecycleStateEnum Enum with underlying type: string
type ImportedModelLifecycleStateEnum string

// Set of constants representing the allowable values for ImportedModelLifecycleStateEnum
const (
	ImportedModelLifecycleStateActive   ImportedModelLifecycleStateEnum = "ACTIVE"
	ImportedModelLifecycleStateCreating ImportedModelLifecycleStateEnum = "CREATING"
	ImportedModelLifecycleStateUpdating ImportedModelLifecycleStateEnum = "UPDATING"
	ImportedModelLifecycleStateDeleting ImportedModelLifecycleStateEnum = "DELETING"
	ImportedModelLifecycleStateDeleted  ImportedModelLifecycleStateEnum = "DELETED"
	ImportedModelLifecycleStateFailed   ImportedModelLifecycleStateEnum = "FAILED"
)

var mappingImportedModelLifecycleStateEnum = map[string]ImportedModelLifecycleStateEnum{
	"ACTIVE":   ImportedModelLifecycleStateActive,
	"CREATING": ImportedModelLifecycleStateCreating,
	"UPDATING": ImportedModelLifecycleStateUpdating,
	"DELETING": ImportedModelLifecycleStateDeleting,
	"DELETED":  ImportedModelLifecycleStateDeleted,
	"FAILED":   ImportedModelLifecycleStateFailed,
}

var mappingImportedModelLifecycleStateEnumLowerCase = map[string]ImportedModelLifecycleStateEnum{
	"active":   ImportedModelLifecycleStateActive,
	"creating": ImportedModelLifecycleStateCreating,
	"updating": ImportedModelLifecycleStateUpdating,
	"deleting": ImportedModelLifecycleStateDeleting,
	"deleted":  ImportedModelLifecycleStateDeleted,
	"failed":   ImportedModelLifecycleStateFailed,
}

// GetImportedModelLifecycleStateEnumValues Enumerates the set of values for ImportedModelLifecycleStateEnum
func GetImportedModelLifecycleStateEnumValues() []ImportedModelLifecycleStateEnum {
	values := make([]ImportedModelLifecycleStateEnum, 0)
	for _, v := range mappingImportedModelLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetImportedModelLifecycleStateEnumStringValues Enumerates the set of values in String for ImportedModelLifecycleStateEnum
func GetImportedModelLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"CREATING",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingImportedModelLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingImportedModelLifecycleStateEnum(val string) (ImportedModelLifecycleStateEnum, bool) {
	enum, ok := mappingImportedModelLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
