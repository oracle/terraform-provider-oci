// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Secrets Management API
//
// API for managing secrets.
//

package vault

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
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

	// A list of rules to control how the secret is used and managed.
	SecretRules []SecretRule `mandatory:"false" json:"secretRules"`
}

func (m UpdateSecretDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateSecretDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		CurrentVersionNumber *int64                            `json:"currentVersionNumber"`
		DefinedTags          map[string]map[string]interface{} `json:"definedTags"`
		Description          *string                           `json:"description"`
		FreeformTags         map[string]string                 `json:"freeformTags"`
		Metadata             map[string]interface{}            `json:"metadata"`
		SecretContent        secretcontentdetails              `json:"secretContent"`
		SecretRules          []secretrule                      `json:"secretRules"`
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

	return
}
