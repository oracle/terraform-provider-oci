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

// UpdateTemplateDetails Updates the specified template.
type UpdateTemplateDetails struct {

	// The template's display name. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the template. Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// Detailed description of the template. This description is displayed in the Console page listing templates when the template is expanded. Avoid entering confidential information.
	LongDescription *string `mandatory:"false" json:"longDescription"`

	// Base64-encoded logo for the template.
	LogoFileBase64Encoded *string `mandatory:"false" json:"logoFileBase64Encoded"`

	TemplateConfigSource UpdateTemplateConfigSourceDetails `mandatory:"false" json:"templateConfigSource"`

	// Free-form tags associated with the resource. Each tag is a key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateTemplateDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateTemplateDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName           *string                           `json:"displayName"`
		Description           *string                           `json:"description"`
		LongDescription       *string                           `json:"longDescription"`
		LogoFileBase64Encoded *string                           `json:"logoFileBase64Encoded"`
		TemplateConfigSource  updatetemplateconfigsourcedetails `json:"templateConfigSource"`
		FreeformTags          map[string]string                 `json:"freeformTags"`
		DefinedTags           map[string]map[string]interface{} `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.LongDescription = model.LongDescription

	m.LogoFileBase64Encoded = model.LogoFileBase64Encoded

	nn, e = model.TemplateConfigSource.UnmarshalPolymorphicJSON(model.TemplateConfigSource.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.TemplateConfigSource = nn.(UpdateTemplateConfigSourceDetails)
	} else {
		m.TemplateConfigSource = nil
	}

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}
