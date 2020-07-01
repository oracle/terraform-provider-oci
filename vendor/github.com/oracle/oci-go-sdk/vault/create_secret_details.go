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

// CreateSecretDetails The details of the secret that you want to create.
type CreateSecretDetails struct {

	// The OCID of the compartment where you want to create the secret.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	SecretContent SecretContentDetails `mandatory:"true" json:"secretContent"`

	// A user-friendly name for the secret. Secret names should be unique within a vault. Avoid entering confidential information. Valid characters are uppercase or lowercase letters, numbers, hyphens, underscores, and periods.
	SecretName *string `mandatory:"true" json:"secretName"`

	// The OCID of the vault where you want to create the secret.
	VaultId *string `mandatory:"true" json:"vaultId"`

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

	// The OCID of the master encryption key that is used to encrypt the secret.
	KeyId *string `mandatory:"false" json:"keyId"`

	// Additional metadata that you can use to provide context about how to use the secret during rotation or
	// other administrative tasks. For example, for a secret that you use to connect to a database, the additional
	// metadata might specify the connection endpoint and the connection string. Provide additional metadata as key-value pairs.
	Metadata map[string]interface{} `mandatory:"false" json:"metadata"`

	// A list of rules to control how the secret is used and managed.
	SecretRules []SecretRule `mandatory:"false" json:"secretRules"`
}

func (m CreateSecretDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *CreateSecretDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DefinedTags   map[string]map[string]interface{} `json:"definedTags"`
		Description   *string                           `json:"description"`
		FreeformTags  map[string]string                 `json:"freeformTags"`
		KeyId         *string                           `json:"keyId"`
		Metadata      map[string]interface{}            `json:"metadata"`
		SecretRules   []secretrule                      `json:"secretRules"`
		CompartmentId *string                           `json:"compartmentId"`
		SecretContent secretcontentdetails              `json:"secretContent"`
		SecretName    *string                           `json:"secretName"`
		VaultId       *string                           `json:"vaultId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DefinedTags = model.DefinedTags

	m.Description = model.Description

	m.FreeformTags = model.FreeformTags

	m.KeyId = model.KeyId

	m.Metadata = model.Metadata

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

	m.CompartmentId = model.CompartmentId

	nn, e = model.SecretContent.UnmarshalPolymorphicJSON(model.SecretContent.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.SecretContent = nn.(SecretContentDetails)
	} else {
		m.SecretContent = nil
	}

	m.SecretName = model.SecretName

	m.VaultId = model.VaultId

	return
}
