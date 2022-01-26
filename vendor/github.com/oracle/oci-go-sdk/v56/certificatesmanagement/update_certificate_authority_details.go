// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Certificates Service Management API
//
// API for managing certificates.
//

package certificatesmanagement

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// UpdateCertificateAuthorityDetails The details for updating a certificate authority (CA).
type UpdateCertificateAuthorityDetails struct {

	// A brief description of the CA.
	Description *string `mandatory:"false" json:"description"`

	// Makes this version the current version. This property cannot be updated in combination with any other properties.
	CurrentVersionNumber *int64 `mandatory:"false" json:"currentVersionNumber"`

	CertificateAuthorityConfig UpdateCertificateAuthorityConfigDetails `mandatory:"false" json:"certificateAuthorityConfig"`

	CertificateRevocationListDetails *CertificateRevocationListDetails `mandatory:"false" json:"certificateRevocationListDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A list of rules that control how the CA is used and managed.
	CertificateAuthorityRules []CertificateAuthorityRule `mandatory:"false" json:"certificateAuthorityRules"`
}

func (m UpdateCertificateAuthorityDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateCertificateAuthorityDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description                      *string                                 `json:"description"`
		CurrentVersionNumber             *int64                                  `json:"currentVersionNumber"`
		CertificateAuthorityConfig       updatecertificateauthorityconfigdetails `json:"certificateAuthorityConfig"`
		CertificateRevocationListDetails *CertificateRevocationListDetails       `json:"certificateRevocationListDetails"`
		FreeformTags                     map[string]string                       `json:"freeformTags"`
		DefinedTags                      map[string]map[string]interface{}       `json:"definedTags"`
		CertificateAuthorityRules        []certificateauthorityrule              `json:"certificateAuthorityRules"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.CurrentVersionNumber = model.CurrentVersionNumber

	nn, e = model.CertificateAuthorityConfig.UnmarshalPolymorphicJSON(model.CertificateAuthorityConfig.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.CertificateAuthorityConfig = nn.(UpdateCertificateAuthorityConfigDetails)
	} else {
		m.CertificateAuthorityConfig = nil
	}

	m.CertificateRevocationListDetails = model.CertificateRevocationListDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.CertificateAuthorityRules = make([]CertificateAuthorityRule, len(model.CertificateAuthorityRules))
	for i, n := range model.CertificateAuthorityRules {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.CertificateAuthorityRules[i] = nn.(CertificateAuthorityRule)
		} else {
			m.CertificateAuthorityRules[i] = nil
		}
	}

	return
}
