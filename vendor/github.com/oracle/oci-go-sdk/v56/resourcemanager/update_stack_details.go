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

// UpdateStackDetails Specifies which fields and the data for each to update on the specified stack.
type UpdateStackDetails struct {

	// The name of the stack.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the stack.
	Description *string `mandatory:"false" json:"description"`

	ConfigSource UpdateConfigSourceDetails `mandatory:"false" json:"configSource"`

	// Terraform variables associated with this resource.
	// The maximum number of variables supported is 250.
	// The maximum size of each variable, including both name and value, is 8192 bytes.
	// Example: `{"CompartmentId": "compartment-id-value"}`
	Variables map[string]string `mandatory:"false" json:"variables"`

	// The version of Terraform to use with the stack. Example: `0.12.x`
	TerraformVersion *string `mandatory:"false" json:"terraformVersion"`

	// Free-form tags associated with this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateStackDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateStackDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName      *string                           `json:"displayName"`
		Description      *string                           `json:"description"`
		ConfigSource     updateconfigsourcedetails         `json:"configSource"`
		Variables        map[string]string                 `json:"variables"`
		TerraformVersion *string                           `json:"terraformVersion"`
		FreeformTags     map[string]string                 `json:"freeformTags"`
		DefinedTags      map[string]map[string]interface{} `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	nn, e = model.ConfigSource.UnmarshalPolymorphicJSON(model.ConfigSource.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ConfigSource = nn.(UpdateConfigSourceDetails)
	} else {
		m.ConfigSource = nil
	}

	m.Variables = model.Variables

	m.TerraformVersion = model.TerraformVersion

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}
