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

// CertificateAuthoritySummary The metadata details of the certificate authority (CA). This summary object does not contain the CA contents.
type CertificateAuthoritySummary struct {

	// The OCID of the certificate authority (CA).
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name for the CA. Names are unique within a compartment. Avoid entering confidential information. Valid characters include uppercase or lowercase letters, numbers, hyphens, underscores, and periods.
	Name *string `mandatory:"true" json:"name"`

	// A property indicating when the CA was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-04-03T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current lifecycle state of the CA.
	LifecycleState CertificateAuthorityLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID of the compartment under which the CA is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The origin of the CA.
	ConfigType CertificateAuthorityConfigTypeEnum `mandatory:"true" json:"configType"`

	// The OCID of the parent CA which issued this CA. If this is the root CA, then this value is the same as the `id`.
	IssuerCertificateAuthorityId *string `mandatory:"false" json:"issuerCertificateAuthorityId"`

	// A brief description of the CA.
	Description *string `mandatory:"false" json:"description"`

	// An optional property indicating when to delete the CA version, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-04-03T21:10:29.600Z`
	TimeOfDeletion *common.SDKTime `mandatory:"false" json:"timeOfDeletion"`

	// The OCID of the Oracle Cloud Infrastructure Vault key used to encrypt the CA.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// An optional list of rules that control how the CA is used and managed.
	CertificateAuthorityRules []CertificateAuthorityRule `mandatory:"false" json:"certificateAuthorityRules"`

	CurrentVersionSummary *CertificateAuthorityVersionSummary `mandatory:"false" json:"currentVersionSummary"`

	Subject *CertificateSubject `mandatory:"false" json:"subject"`

	// The algorithm used to sign public key certificates that the CA issues.
	SigningAlgorithm SignatureAlgorithmEnum `mandatory:"false" json:"signingAlgorithm,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CertificateAuthoritySummary) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *CertificateAuthoritySummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		IssuerCertificateAuthorityId *string                                `json:"issuerCertificateAuthorityId"`
		Description                  *string                                `json:"description"`
		TimeOfDeletion               *common.SDKTime                        `json:"timeOfDeletion"`
		KmsKeyId                     *string                                `json:"kmsKeyId"`
		CertificateAuthorityRules    []certificateauthorityrule             `json:"certificateAuthorityRules"`
		CurrentVersionSummary        *CertificateAuthorityVersionSummary    `json:"currentVersionSummary"`
		Subject                      *CertificateSubject                    `json:"subject"`
		SigningAlgorithm             SignatureAlgorithmEnum                 `json:"signingAlgorithm"`
		FreeformTags                 map[string]string                      `json:"freeformTags"`
		DefinedTags                  map[string]map[string]interface{}      `json:"definedTags"`
		Id                           *string                                `json:"id"`
		Name                         *string                                `json:"name"`
		TimeCreated                  *common.SDKTime                        `json:"timeCreated"`
		LifecycleState               CertificateAuthorityLifecycleStateEnum `json:"lifecycleState"`
		CompartmentId                *string                                `json:"compartmentId"`
		ConfigType                   CertificateAuthorityConfigTypeEnum     `json:"configType"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.IssuerCertificateAuthorityId = model.IssuerCertificateAuthorityId

	m.Description = model.Description

	m.TimeOfDeletion = model.TimeOfDeletion

	m.KmsKeyId = model.KmsKeyId

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

	m.CurrentVersionSummary = model.CurrentVersionSummary

	m.Subject = model.Subject

	m.SigningAlgorithm = model.SigningAlgorithm

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.Id = model.Id

	m.Name = model.Name

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	m.CompartmentId = model.CompartmentId

	m.ConfigType = model.ConfigType

	return
}
