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

// DedicatedAiClusterSummary Summary information about a dedicated AI cluster.
type DedicatedAiClusterSummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the dedicated AI cluster.
	Id *string `mandatory:"true" json:"id"`

	// The dedicated AI cluster type indicating whether this is a fine-tuning/training processor or hosting/inference processor.
	// Allowed values are:
	// - HOSTING
	// - FINE_TUNING
	Type DedicatedAiClusterTypeEnum `mandatory:"true" json:"type"`

	// The compartment OCID to create the dedicated AI cluster in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time the dedicated AI cluster was created, in the format defined by RFC 3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the dedicated AI cluster.
	// Allowed values are:
	// - CREATING
	// - ACTIVE
	// - UPDATING
	// - DELETING
	// - DELETED
	// - FAILED
	// - NEEDS_ATTENTION
	LifecycleState DedicatedAiClusterLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The number of dedicated units in this AI cluster.
	UnitCount *int `mandatory:"true" json:"unitCount"`

	// The shape of dedicated unit in this AI cluster. The underlying hardware configuration is hidden from customers.
	UnitShape DedicatedAiClusterUnitShapeEnum `mandatory:"true" json:"unitShape"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// An optional description of the dedicated AI cluster.
	Description *string `mandatory:"false" json:"description"`

	// The date and time the dedicated AI cluster was updated, in the format defined by RFC 3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state of the dedicated AI cluster in more detail that can provide actionable information.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	Capacity DedicatedAiClusterCapacity `mandatory:"false" json:"capacity"`

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

func (m DedicatedAiClusterSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DedicatedAiClusterSummary) ValidateEnumValue() (bool, error) {
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
func (m *DedicatedAiClusterSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName      *string                              `json:"displayName"`
		Description      *string                              `json:"description"`
		TimeUpdated      *common.SDKTime                      `json:"timeUpdated"`
		LifecycleDetails *string                              `json:"lifecycleDetails"`
		Capacity         dedicatedaiclustercapacity           `json:"capacity"`
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
