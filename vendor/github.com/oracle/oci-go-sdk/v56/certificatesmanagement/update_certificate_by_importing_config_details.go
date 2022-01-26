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

// UpdateCertificateByImportingConfigDetails The details of the configuration for updating a certificate based on the keys from an imported certificate.
type UpdateCertificateByImportingConfigDetails struct {

	// The certificate chain (in PEM format) for the imported certificate.
	CertChainPem *string `mandatory:"true" json:"certChainPem"`

	// The private key (in PEM format) for the imported certificate.
	PrivateKeyPem *string `mandatory:"true" json:"privateKeyPem"`

	// The certificate (in PEM format) for the imported certificate.
	CertificatePem *string `mandatory:"true" json:"certificatePem"`

	// A name for the certificate version. When the value is not null, a name is unique across versions of a given certificate.
	VersionName *string `mandatory:"false" json:"versionName"`

	// An optional passphrase for the private key.
	PrivateKeyPemPassphrase *string `mandatory:"false" json:"privateKeyPemPassphrase"`

	// The rotation state of the certificate. The default is `CURRENT`, meaning that the certificate is currently in use. A certificate version
	// that you mark as `PENDING` is staged and available for use, but you don't yet want to rotate it into current, active use. For example,
	// you might update a certificate and mark its rotation state as `PENDING` if you haven't yet updated the certificate on the target system.
	Stage UpdateCertificateConfigDetailsStageEnum `mandatory:"false" json:"stage,omitempty"`
}

//GetVersionName returns VersionName
func (m UpdateCertificateByImportingConfigDetails) GetVersionName() *string {
	return m.VersionName
}

//GetStage returns Stage
func (m UpdateCertificateByImportingConfigDetails) GetStage() UpdateCertificateConfigDetailsStageEnum {
	return m.Stage
}

func (m UpdateCertificateByImportingConfigDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateCertificateByImportingConfigDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateCertificateByImportingConfigDetails UpdateCertificateByImportingConfigDetails
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeUpdateCertificateByImportingConfigDetails
	}{
		"IMPORTED",
		(MarshalTypeUpdateCertificateByImportingConfigDetails)(m),
	}

	return json.Marshal(&s)
}
