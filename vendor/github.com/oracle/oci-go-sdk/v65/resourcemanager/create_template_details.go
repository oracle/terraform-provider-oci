// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateTemplateDetails Creation details for a template.
type CreateTemplateDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing this template.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The template's display name. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Description of the template. Avoid entering confidential information.
	Description *string `mandatory:"true" json:"description"`

	TemplateConfigSource CreateTemplateConfigSourceDetails `mandatory:"true" json:"templateConfigSource"`

	// Detailed description of the template. This description is displayed in the Console page listing templates when the template is expanded. Avoid entering confidential information.
	LongDescription *string `mandatory:"false" json:"longDescription"`

	// Base64-encoded logo to use as the template icon.
	// Template icon file requirements: PNG format, 50 KB maximum, 110 x 110 pixels.
	LogoFileBase64Encoded *string `mandatory:"false" json:"logoFileBase64Encoded"`

	// Free-form tags associated with the resource. Each tag is a key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateTemplateDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateTemplateDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateTemplateDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		LongDescription       *string                           `json:"longDescription"`
		LogoFileBase64Encoded *string                           `json:"logoFileBase64Encoded"`
		FreeformTags          map[string]string                 `json:"freeformTags"`
		DefinedTags           map[string]map[string]interface{} `json:"definedTags"`
		CompartmentId         *string                           `json:"compartmentId"`
		DisplayName           *string                           `json:"displayName"`
		Description           *string                           `json:"description"`
		TemplateConfigSource  createtemplateconfigsourcedetails `json:"templateConfigSource"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.LongDescription = model.LongDescription

	m.LogoFileBase64Encoded = model.LogoFileBase64Encoded

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

	m.Description = model.Description

	nn, e = model.TemplateConfigSource.UnmarshalPolymorphicJSON(model.TemplateConfigSource.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.TemplateConfigSource = nn.(CreateTemplateConfigSourceDetails)
	} else {
		m.TemplateConfigSource = nil
	}

	return
}
