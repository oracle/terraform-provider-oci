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

// CreateCertificateManagedExternallyIssuedByInternalCaConfigDetails The details of the configuration for creating an externally managed certificate which is issued by a private certificate authority (CA).
type CreateCertificateManagedExternallyIssuedByInternalCaConfigDetails struct {

	// The OCID of the private CA.
	IssuerCertificateAuthorityId *string `mandatory:"true" json:"issuerCertificateAuthorityId"`

	// The certificate signing request (in PEM format).
	CsrPem *string `mandatory:"true" json:"csrPem"`

	// A name for the certificate. When the value is not null, a name is unique across versions of a given certificate.
	VersionName *string `mandatory:"false" json:"versionName"`

	Validity *Validity `mandatory:"false" json:"validity"`
}

//GetVersionName returns VersionName
func (m CreateCertificateManagedExternallyIssuedByInternalCaConfigDetails) GetVersionName() *string {
	return m.VersionName
}

func (m CreateCertificateManagedExternallyIssuedByInternalCaConfigDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateCertificateManagedExternallyIssuedByInternalCaConfigDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateCertificateManagedExternallyIssuedByInternalCaConfigDetails CreateCertificateManagedExternallyIssuedByInternalCaConfigDetails
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeCreateCertificateManagedExternallyIssuedByInternalCaConfigDetails
	}{
		"MANAGED_EXTERNALLY_ISSUED_BY_INTERNAL_CA",
		(MarshalTypeCreateCertificateManagedExternallyIssuedByInternalCaConfigDetails)(m),
	}

	return json.Marshal(&s)
}
