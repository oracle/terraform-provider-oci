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

// HostedDeployment Hosted deployment is designed to support the full spectrum of agent use cases from lightweight, employee-facing assistants and internal workflow automation, to enterprise-grade, large-scale customer-facing workloads.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator who gives OCI resource access to users. See
// Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm) and Getting Access to Generative AI Resources (https://docs.oracle.com/iaas/Content/generative-ai/iam-policies.htm).
type HostedDeployment struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the hosted deployment.
	Id *string `mandatory:"true" json:"id"`

	// The current state of the hosted deployment.
	LifecycleState HostedDeploymentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the application.
	HostedApplicationId *string `mandatory:"true" json:"hostedApplicationId"`

	ActiveArtifact Artifact `mandatory:"true" json:"activeArtifact"`

	// array of Artifacts.
	Artifacts []Artifact `mandatory:"true" json:"artifacts"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The compartment OCID to create the hosted application in.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The date and time the hosted deployment was created, in the format defined by RFC 3339
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the hosted deployment was updated, in the format defined by RFC 3339
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

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

func (m HostedDeployment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HostedDeployment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingHostedDeploymentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetHostedDeploymentLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *HostedDeployment) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName         *string                            `json:"displayName"`
		CompartmentId       *string                            `json:"compartmentId"`
		TimeCreated         *common.SDKTime                    `json:"timeCreated"`
		TimeUpdated         *common.SDKTime                    `json:"timeUpdated"`
		FreeformTags        map[string]string                  `json:"freeformTags"`
		DefinedTags         map[string]map[string]interface{}  `json:"definedTags"`
		SystemTags          map[string]map[string]interface{}  `json:"systemTags"`
		Id                  *string                            `json:"id"`
		LifecycleState      HostedDeploymentLifecycleStateEnum `json:"lifecycleState"`
		HostedApplicationId *string                            `json:"hostedApplicationId"`
		ActiveArtifact      artifact                           `json:"activeArtifact"`
		Artifacts           []artifact                         `json:"artifacts"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.LifecycleState = model.LifecycleState

	m.HostedApplicationId = model.HostedApplicationId

	nn, e = model.ActiveArtifact.UnmarshalPolymorphicJSON(model.ActiveArtifact.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ActiveArtifact = nn.(Artifact)
	} else {
		m.ActiveArtifact = nil
	}

	m.Artifacts = make([]Artifact, len(model.Artifacts))
	for i, n := range model.Artifacts {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Artifacts[i] = nn.(Artifact)
		} else {
			m.Artifacts[i] = nil
		}
	}
	return
}

// HostedDeploymentLifecycleStateEnum Enum with underlying type: string
type HostedDeploymentLifecycleStateEnum string

// Set of constants representing the allowable values for HostedDeploymentLifecycleStateEnum
const (
	HostedDeploymentLifecycleStateCreating HostedDeploymentLifecycleStateEnum = "CREATING"
	HostedDeploymentLifecycleStateActive   HostedDeploymentLifecycleStateEnum = "ACTIVE"
	HostedDeploymentLifecycleStateInactive HostedDeploymentLifecycleStateEnum = "INACTIVE"
	HostedDeploymentLifecycleStateUpdating HostedDeploymentLifecycleStateEnum = "UPDATING"
	HostedDeploymentLifecycleStateDeleting HostedDeploymentLifecycleStateEnum = "DELETING"
	HostedDeploymentLifecycleStateDeleted  HostedDeploymentLifecycleStateEnum = "DELETED"
	HostedDeploymentLifecycleStateFailed   HostedDeploymentLifecycleStateEnum = "FAILED"
)

var mappingHostedDeploymentLifecycleStateEnum = map[string]HostedDeploymentLifecycleStateEnum{
	"CREATING": HostedDeploymentLifecycleStateCreating,
	"ACTIVE":   HostedDeploymentLifecycleStateActive,
	"INACTIVE": HostedDeploymentLifecycleStateInactive,
	"UPDATING": HostedDeploymentLifecycleStateUpdating,
	"DELETING": HostedDeploymentLifecycleStateDeleting,
	"DELETED":  HostedDeploymentLifecycleStateDeleted,
	"FAILED":   HostedDeploymentLifecycleStateFailed,
}

var mappingHostedDeploymentLifecycleStateEnumLowerCase = map[string]HostedDeploymentLifecycleStateEnum{
	"creating": HostedDeploymentLifecycleStateCreating,
	"active":   HostedDeploymentLifecycleStateActive,
	"inactive": HostedDeploymentLifecycleStateInactive,
	"updating": HostedDeploymentLifecycleStateUpdating,
	"deleting": HostedDeploymentLifecycleStateDeleting,
	"deleted":  HostedDeploymentLifecycleStateDeleted,
	"failed":   HostedDeploymentLifecycleStateFailed,
}

// GetHostedDeploymentLifecycleStateEnumValues Enumerates the set of values for HostedDeploymentLifecycleStateEnum
func GetHostedDeploymentLifecycleStateEnumValues() []HostedDeploymentLifecycleStateEnum {
	values := make([]HostedDeploymentLifecycleStateEnum, 0)
	for _, v := range mappingHostedDeploymentLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetHostedDeploymentLifecycleStateEnumStringValues Enumerates the set of values in String for HostedDeploymentLifecycleStateEnum
func GetHostedDeploymentLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingHostedDeploymentLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHostedDeploymentLifecycleStateEnum(val string) (HostedDeploymentLifecycleStateEnum, bool) {
	enum, ok := mappingHostedDeploymentLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
