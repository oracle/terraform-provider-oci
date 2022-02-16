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

// CertificateSummary The details of the certificate. This object does not contain the certificate contents.
type CertificateSummary struct {

	// The OCID of the certificate.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name for the certificate. Names are unique within a compartment. Avoid entering confidential information. Valid characters are uppercase or lowercase letters, numbers, hyphens, underscores, and periods.
	Name *string `mandatory:"true" json:"name"`

	// A property indicating when the certificate was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-04-03T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current lifecycle state of the certificate.
	LifecycleState CertificateLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID of the compartment that contains the certificate.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The origin of the certificate.
	ConfigType CertificateConfigTypeEnum `mandatory:"true" json:"configType"`

	// The OCID of the certificate authority (CA) that issued the certificate.
	IssuerCertificateAuthorityId *string `mandatory:"false" json:"issuerCertificateAuthorityId"`

	// A brief description of the certificate.
	Description *string `mandatory:"false" json:"description"`

	// An optional property indicating when to delete the certificate version, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-04-03T21:10:29.600Z`
	TimeOfDeletion *common.SDKTime `mandatory:"false" json:"timeOfDeletion"`

	// An optional list of rules that control how the certificate is used and managed.
	CertificateRules []CertificateRule `mandatory:"false" json:"certificateRules"`

	CurrentVersionSummary *CertificateVersionSummary `mandatory:"false" json:"currentVersionSummary"`

	Subject *CertificateSubject `mandatory:"false" json:"subject"`

	// The algorithm used to create key pairs.
	KeyAlgorithm KeyAlgorithmEnum `mandatory:"false" json:"keyAlgorithm,omitempty"`

	// The algorithm used to sign the public key certificate.
	SignatureAlgorithm SignatureAlgorithmEnum `mandatory:"false" json:"signatureAlgorithm,omitempty"`

	// The name of the profile used to create the certificate, which depends on the type of certificate you need.
	CertificateProfileType CertificateProfileTypeEnum `mandatory:"false" json:"certificateProfileType,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CertificateSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CertificateSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCertificateLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCertificateLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCertificateConfigTypeEnum(string(m.ConfigType)); !ok && m.ConfigType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConfigType: %s. Supported values are: %s.", m.ConfigType, strings.Join(GetCertificateConfigTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingKeyAlgorithmEnum(string(m.KeyAlgorithm)); !ok && m.KeyAlgorithm != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for KeyAlgorithm: %s. Supported values are: %s.", m.KeyAlgorithm, strings.Join(GetKeyAlgorithmEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSignatureAlgorithmEnum(string(m.SignatureAlgorithm)); !ok && m.SignatureAlgorithm != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SignatureAlgorithm: %s. Supported values are: %s.", m.SignatureAlgorithm, strings.Join(GetSignatureAlgorithmEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCertificateProfileTypeEnum(string(m.CertificateProfileType)); !ok && m.CertificateProfileType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CertificateProfileType: %s. Supported values are: %s.", m.CertificateProfileType, strings.Join(GetCertificateProfileTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CertificateSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		IssuerCertificateAuthorityId *string                           `json:"issuerCertificateAuthorityId"`
		Description                  *string                           `json:"description"`
		TimeOfDeletion               *common.SDKTime                   `json:"timeOfDeletion"`
		CertificateRules             []certificaterule                 `json:"certificateRules"`
		CurrentVersionSummary        *CertificateVersionSummary        `json:"currentVersionSummary"`
		Subject                      *CertificateSubject               `json:"subject"`
		KeyAlgorithm                 KeyAlgorithmEnum                  `json:"keyAlgorithm"`
		SignatureAlgorithm           SignatureAlgorithmEnum            `json:"signatureAlgorithm"`
		CertificateProfileType       CertificateProfileTypeEnum        `json:"certificateProfileType"`
		FreeformTags                 map[string]string                 `json:"freeformTags"`
		DefinedTags                  map[string]map[string]interface{} `json:"definedTags"`
		Id                           *string                           `json:"id"`
		Name                         *string                           `json:"name"`
		TimeCreated                  *common.SDKTime                   `json:"timeCreated"`
		LifecycleState               CertificateLifecycleStateEnum     `json:"lifecycleState"`
		CompartmentId                *string                           `json:"compartmentId"`
		ConfigType                   CertificateConfigTypeEnum         `json:"configType"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.IssuerCertificateAuthorityId = model.IssuerCertificateAuthorityId

	m.Description = model.Description

	m.TimeOfDeletion = model.TimeOfDeletion

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

	m.CurrentVersionSummary = model.CurrentVersionSummary

	m.Subject = model.Subject

	m.KeyAlgorithm = model.KeyAlgorithm

	m.SignatureAlgorithm = model.SignatureAlgorithm

	m.CertificateProfileType = model.CertificateProfileType

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
