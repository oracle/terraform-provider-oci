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

// CreateSubordinateCaIssuedByInternalCaConfigDetails The details for creating a private subordinate certificate authority (CA) which is issued by a private CA.
type CreateSubordinateCaIssuedByInternalCaConfigDetails struct {

	// The OCID of the private CA.
	IssuerCertificateAuthorityId *string `mandatory:"true" json:"issuerCertificateAuthorityId"`

	Subject *CertificateSubject `mandatory:"true" json:"subject"`

	// The name of the CA version. When the value is not null, a name is unique across versions of a given CA.
	VersionName *string `mandatory:"false" json:"versionName"`

	Validity *Validity `mandatory:"false" json:"validity"`

	// The algorithm used to sign public key certificates that the CA issues.
	SigningAlgorithm SignatureAlgorithmEnum `mandatory:"false" json:"signingAlgorithm,omitempty"`
}

//GetVersionName returns VersionName
func (m CreateSubordinateCaIssuedByInternalCaConfigDetails) GetVersionName() *string {
	return m.VersionName
}

func (m CreateSubordinateCaIssuedByInternalCaConfigDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateSubordinateCaIssuedByInternalCaConfigDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateSubordinateCaIssuedByInternalCaConfigDetails CreateSubordinateCaIssuedByInternalCaConfigDetails
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeCreateSubordinateCaIssuedByInternalCaConfigDetails
	}{
		"SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA",
		(MarshalTypeCreateSubordinateCaIssuedByInternalCaConfigDetails)(m),
	}

	return json.Marshal(&s)
}
