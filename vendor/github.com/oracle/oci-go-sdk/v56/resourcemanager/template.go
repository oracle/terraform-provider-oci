// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// Use the Resource Manager API to automate deployment and operations for all Oracle Cloud Infrastructure resources.
// Using the infrastructure-as-code (IaC) model, the service is based on Terraform, an open source industry standard that lets DevOps engineers develop and deploy their infrastructure anywhere.
// For more information, see
// the Resource Manager documentation (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/home.htm).
//

package resourcemanager

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// Template The properties that define a template. A template is a pre-built Terraform configuration that provisions a set of resources used in a common scenario.
type Template struct {

	// Unique identifier (OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) for the template.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing this template.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Unique identifier for the category where the template is located.
	// Possible values are `0` (Quick Starts), `1` (Service), `2` (Architecture), and `3` (Private).
	CategoryId *string `mandatory:"false" json:"categoryId"`

	// Human-readable name of the template.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Brief description of the template.
	Description *string `mandatory:"false" json:"description"`

	// Detailed description of the template. This description is displayed in the Console page listing templates when the template is expanded. Avoid entering confidential information.
	LongDescription *string `mandatory:"false" json:"longDescription"`

	// whether the template will work for free tier tenancy.
	IsFreeTier *bool `mandatory:"false" json:"isFreeTier"`

	// The date and time at which the template was created.
	// Format is defined by RFC3339.
	// Example: `2020-11-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	TemplateConfigSource TemplateConfigSource `mandatory:"false" json:"templateConfigSource"`

	// The current lifecycle state of the template.
	LifecycleState TemplateLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Free-form tags associated with the resource. Each tag is a key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m Template) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *Template) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		CompartmentId        *string                           `json:"compartmentId"`
		CategoryId           *string                           `json:"categoryId"`
		DisplayName          *string                           `json:"displayName"`
		Description          *string                           `json:"description"`
		LongDescription      *string                           `json:"longDescription"`
		IsFreeTier           *bool                             `json:"isFreeTier"`
		TimeCreated          *common.SDKTime                   `json:"timeCreated"`
		TemplateConfigSource templateconfigsource              `json:"templateConfigSource"`
		LifecycleState       TemplateLifecycleStateEnum        `json:"lifecycleState"`
		FreeformTags         map[string]string                 `json:"freeformTags"`
		DefinedTags          map[string]map[string]interface{} `json:"definedTags"`
		Id                   *string                           `json:"id"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.CompartmentId = model.CompartmentId

	m.CategoryId = model.CategoryId

	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.LongDescription = model.LongDescription

	m.IsFreeTier = model.IsFreeTier

	m.TimeCreated = model.TimeCreated

	nn, e = model.TemplateConfigSource.UnmarshalPolymorphicJSON(model.TemplateConfigSource.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.TemplateConfigSource = nn.(TemplateConfigSource)
	} else {
		m.TemplateConfigSource = nil
	}

	m.LifecycleState = model.LifecycleState

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.Id = model.Id

	return
}

// TemplateLifecycleStateEnum Enum with underlying type: string
type TemplateLifecycleStateEnum string

// Set of constants representing the allowable values for TemplateLifecycleStateEnum
const (
	TemplateLifecycleStateActive TemplateLifecycleStateEnum = "ACTIVE"
)

var mappingTemplateLifecycleState = map[string]TemplateLifecycleStateEnum{
	"ACTIVE": TemplateLifecycleStateActive,
}

// GetTemplateLifecycleStateEnumValues Enumerates the set of values for TemplateLifecycleStateEnum
func GetTemplateLifecycleStateEnumValues() []TemplateLifecycleStateEnum {
	values := make([]TemplateLifecycleStateEnum, 0)
	for _, v := range mappingTemplateLifecycleState {
		values = append(values, v)
	}
	return values
}
