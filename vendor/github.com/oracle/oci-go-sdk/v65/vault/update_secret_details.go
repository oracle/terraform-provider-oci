// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vault Secret Management API
//
// Use the Secret Management API to manage secrets and secret versions. For more information, see Managing Secrets (https://docs.cloud.oracle.com/Content/KeyManagement/Tasks/managingsecrets.htm).
//

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateSecretDetails Details for updating a secret.
type UpdateSecretDetails struct {

	// Details to update the secret version of the specified secret. The secret contents,
	// version number, and rules can't be specified at the same time.
	// Updating the secret contents automatically creates a new secret version.
	CurrentVersionNumber *int64 `mandatory:"false" json:"currentVersionNumber"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A brief description of the secret. Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Additional metadata that you can use to provide context about how to use the secret or during rotation or
	// other administrative tasks. For example, for a secret that you use to connect to a database, the additional
	// metadata might specify the connection endpoint and the connection string. Provide additional metadata as key-value pairs.
	Metadata map[string]interface{} `mandatory:"false" json:"metadata"`

	SecretContent SecretContentDetails `mandatory:"false" json:"secretContent"`

	RotationConfig *RotationConfig `mandatory:"false" json:"rotationConfig"`

	// A list of rules to control how the secret is used and managed.
	SecretRules []SecretRule `mandatory:"false" json:"secretRules"`

	SecretGenerationContext SecretGenerationContext `mandatory:"false" json:"secretGenerationContext"`

	// The value of this flag determines whether or not secret content will be generated automatically.
	EnableAutoGeneration *bool `mandatory:"false" json:"enableAutoGeneration"`
}

func (m UpdateSecretDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateSecretDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateSecretDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		CurrentVersionNumber    *int64                            `json:"currentVersionNumber"`
		DefinedTags             map[string]map[string]interface{} `json:"definedTags"`
		Description             *string                           `json:"description"`
		FreeformTags            map[string]string                 `json:"freeformTags"`
		Metadata                map[string]interface{}            `json:"metadata"`
		SecretContent           secretcontentdetails              `json:"secretContent"`
		RotationConfig          *RotationConfig                   `json:"rotationConfig"`
		SecretRules             []secretrule                      `json:"secretRules"`
		SecretGenerationContext secretgenerationcontext           `json:"secretGenerationContext"`
		EnableAutoGeneration    *bool                             `json:"enableAutoGeneration"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.CurrentVersionNumber = model.CurrentVersionNumber

	m.DefinedTags = model.DefinedTags

	m.Description = model.Description

	m.FreeformTags = model.FreeformTags

	m.Metadata = model.Metadata

	nn, e = model.SecretContent.UnmarshalPolymorphicJSON(model.SecretContent.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.SecretContent = nn.(SecretContentDetails)
	} else {
		m.SecretContent = nil
	}

	m.RotationConfig = model.RotationConfig

	m.SecretRules = make([]SecretRule, len(model.SecretRules))
	for i, n := range model.SecretRules {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.SecretRules[i] = nn.(SecretRule)
		} else {
			m.SecretRules[i] = nil
		}
	}
	nn, e = model.SecretGenerationContext.UnmarshalPolymorphicJSON(model.SecretGenerationContext.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.SecretGenerationContext = nn.(SecretGenerationContext)
	} else {
		m.SecretGenerationContext = nil
	}

	m.EnableAutoGeneration = model.EnableAutoGeneration

	return
}
