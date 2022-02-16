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
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// UpdateCertificateDetails The details of the certificate to update.
type UpdateCertificateDetails struct {

	// A brief description of the certificate. Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// Makes this version the current version. This property cannot be updated in combination with any other properties.
	CurrentVersionNumber *int64 `mandatory:"false" json:"currentVersionNumber"`

	CertificateConfig UpdateCertificateConfigDetails `mandatory:"false" json:"certificateConfig"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// An optional list of rules that control how the certificate is used and managed.
	CertificateRules []CertificateRule `mandatory:"false" json:"certificateRules"`
}

func (m UpdateCertificateDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateCertificateDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateCertificateDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description          *string                           `json:"description"`
		CurrentVersionNumber *int64                            `json:"currentVersionNumber"`
		CertificateConfig    updatecertificateconfigdetails    `json:"certificateConfig"`
		FreeformTags         map[string]string                 `json:"freeformTags"`
		DefinedTags          map[string]map[string]interface{} `json:"definedTags"`
		CertificateRules     []certificaterule                 `json:"certificateRules"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.CurrentVersionNumber = model.CurrentVersionNumber

	nn, e = model.CertificateConfig.UnmarshalPolymorphicJSON(model.CertificateConfig.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.CertificateConfig = nn.(UpdateCertificateConfigDetails)
	} else {
		m.CertificateConfig = nil
	}

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

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

	return
}
