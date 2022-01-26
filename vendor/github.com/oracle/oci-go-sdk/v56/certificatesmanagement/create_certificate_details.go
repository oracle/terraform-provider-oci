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

// CreateCertificateDetails The details of the certificate to create.
type CreateCertificateDetails struct {

	// A user-friendly name for the certificate. Names are unique within a compartment. Avoid entering confidential information. Valid characters are uppercase or lowercase letters, numbers, hyphens, underscores, and periods.
	Name *string `mandatory:"true" json:"name"`

	// The OCID of the compartment where you want to create the certificate.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	CertificateConfig CreateCertificateConfigDetails `mandatory:"true" json:"certificateConfig"`

	// A brief description of the certificate. Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// An optional list of rules that control how the certificate is used and managed.
	CertificateRules []CertificateRule `mandatory:"false" json:"certificateRules"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateCertificateDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *CreateCertificateDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description       *string                           `json:"description"`
		CertificateRules  []certificaterule                 `json:"certificateRules"`
		FreeformTags      map[string]string                 `json:"freeformTags"`
		DefinedTags       map[string]map[string]interface{} `json:"definedTags"`
		Name              *string                           `json:"name"`
		CompartmentId     *string                           `json:"compartmentId"`
		CertificateConfig createcertificateconfigdetails    `json:"certificateConfig"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.CertificateRules = make([]CertificateRule, len(model.CertificateRules))
	for i, n := range model.CertificateRules {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.CertificateRules[i] = nn.(CertificateRule)
		} else {
			m.CertificateRules[i] = nil
		}
	}

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.Name = model.Name

	m.CompartmentId = model.CompartmentId

	nn, e = model.CertificateConfig.UnmarshalPolymorphicJSON(model.CertificateConfig.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.CertificateConfig = nn.(CreateCertificateConfigDetails)
	} else {
		m.CertificateConfig = nil
	}

	return
}
