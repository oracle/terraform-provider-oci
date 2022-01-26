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

// CreateCertificateAuthorityDetails The details for creating a certificate authority (CA).
type CreateCertificateAuthorityDetails struct {

	// A user-friendly name for the CA. Names are unique within a compartment. Avoid entering confidential information. Valid characters include uppercase or lowercase letters, numbers, hyphens, underscores, and periods.
	Name *string `mandatory:"true" json:"name"`

	// The compartment in which you want to create the CA.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	CertificateAuthorityConfig CreateCertificateAuthorityConfigDetails `mandatory:"true" json:"certificateAuthorityConfig"`

	// The OCID of the Oracle Cloud Infrastructure Vault key used to encrypt the CA.
	KmsKeyId *string `mandatory:"true" json:"kmsKeyId"`

	// A brief description of the CA.
	Description *string `mandatory:"false" json:"description"`

	// A list of rules that control how the CA is used and managed.
	CertificateAuthorityRules []CertificateAuthorityRule `mandatory:"false" json:"certificateAuthorityRules"`

	CertificateRevocationListDetails *CertificateRevocationListDetails `mandatory:"false" json:"certificateRevocationListDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateCertificateAuthorityDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *CreateCertificateAuthorityDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description                      *string                                 `json:"description"`
		CertificateAuthorityRules        []certificateauthorityrule              `json:"certificateAuthorityRules"`
		CertificateRevocationListDetails *CertificateRevocationListDetails       `json:"certificateRevocationListDetails"`
		FreeformTags                     map[string]string                       `json:"freeformTags"`
		DefinedTags                      map[string]map[string]interface{}       `json:"definedTags"`
		Name                             *string                                 `json:"name"`
		CompartmentId                    *string                                 `json:"compartmentId"`
		CertificateAuthorityConfig       createcertificateauthorityconfigdetails `json:"certificateAuthorityConfig"`
		KmsKeyId                         *string                                 `json:"kmsKeyId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

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

	m.CertificateRevocationListDetails = model.CertificateRevocationListDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.Name = model.Name

	m.CompartmentId = model.CompartmentId

	nn, e = model.CertificateAuthorityConfig.UnmarshalPolymorphicJSON(model.CertificateAuthorityConfig.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.CertificateAuthorityConfig = nn.(CreateCertificateAuthorityConfigDetails)
	} else {
		m.CertificateAuthorityConfig = nil
	}

	m.KmsKeyId = model.KmsKeyId

	return
}
