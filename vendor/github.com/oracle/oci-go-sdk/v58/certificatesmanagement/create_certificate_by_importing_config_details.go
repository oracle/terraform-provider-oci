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

// CreateCertificateByImportingConfigDetails The details of the configuration for creating a certificate based on the keys from an imported certificate.
type CreateCertificateByImportingConfigDetails struct {

	// The certificate chain (in PEM format) for the imported certificate.
	CertChainPem *string `mandatory:"true" json:"certChainPem"`

	// The private key (in PEM format) for the imported certificate.
	PrivateKeyPem *string `mandatory:"true" json:"privateKeyPem"`

	// The certificate (in PEM format) for the imported certificate.
	CertificatePem *string `mandatory:"true" json:"certificatePem"`

	// A name for the certificate. When the value is not null, a name is unique across versions of a given certificate.
	VersionName *string `mandatory:"false" json:"versionName"`

	// An optional passphrase for the private key.
	PrivateKeyPemPassphrase *string `mandatory:"false" json:"privateKeyPemPassphrase"`
}

//GetVersionName returns VersionName
func (m CreateCertificateByImportingConfigDetails) GetVersionName() *string {
	return m.VersionName
}

func (m CreateCertificateByImportingConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateCertificateByImportingConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateCertificateByImportingConfigDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateCertificateByImportingConfigDetails CreateCertificateByImportingConfigDetails
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeCreateCertificateByImportingConfigDetails
	}{
		"IMPORTED",
		(MarshalTypeCreateCertificateByImportingConfigDetails)(m),
	}

	return json.Marshal(&s)
}
