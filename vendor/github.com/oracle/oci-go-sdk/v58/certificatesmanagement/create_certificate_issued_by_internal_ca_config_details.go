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

// CreateCertificateIssuedByInternalCaConfigDetails The details of the configuration for creating an internally managed certificate which is issued by a private certificate authority (CA).
type CreateCertificateIssuedByInternalCaConfigDetails struct {

	// The OCID of the private CA.
	IssuerCertificateAuthorityId *string `mandatory:"true" json:"issuerCertificateAuthorityId"`

	Subject *CertificateSubject `mandatory:"true" json:"subject"`

	// A name for the certificate. When the value is not null, a name is unique across versions of a given certificate.
	VersionName *string `mandatory:"false" json:"versionName"`

	Validity *Validity `mandatory:"false" json:"validity"`

	// A list of subject alternative names.
	SubjectAlternativeNames []CertificateSubjectAlternativeName `mandatory:"false" json:"subjectAlternativeNames"`

	// The name of the profile used to create the certificate, which depends on the type of certificate you need.
	CertificateProfileType CertificateProfileTypeEnum `mandatory:"true" json:"certificateProfileType"`

	// The algorithm to use to create key pairs.
	KeyAlgorithm KeyAlgorithmEnum `mandatory:"false" json:"keyAlgorithm,omitempty"`

	// The algorithm to use to sign the public key certificate.
	SignatureAlgorithm SignatureAlgorithmEnum `mandatory:"false" json:"signatureAlgorithm,omitempty"`
}

//GetVersionName returns VersionName
func (m CreateCertificateIssuedByInternalCaConfigDetails) GetVersionName() *string {
	return m.VersionName
}

func (m CreateCertificateIssuedByInternalCaConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateCertificateIssuedByInternalCaConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCertificateProfileTypeEnum(string(m.CertificateProfileType)); !ok && m.CertificateProfileType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CertificateProfileType: %s. Supported values are: %s.", m.CertificateProfileType, strings.Join(GetCertificateProfileTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingKeyAlgorithmEnum(string(m.KeyAlgorithm)); !ok && m.KeyAlgorithm != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for KeyAlgorithm: %s. Supported values are: %s.", m.KeyAlgorithm, strings.Join(GetKeyAlgorithmEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSignatureAlgorithmEnum(string(m.SignatureAlgorithm)); !ok && m.SignatureAlgorithm != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SignatureAlgorithm: %s. Supported values are: %s.", m.SignatureAlgorithm, strings.Join(GetSignatureAlgorithmEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateCertificateIssuedByInternalCaConfigDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateCertificateIssuedByInternalCaConfigDetails CreateCertificateIssuedByInternalCaConfigDetails
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeCreateCertificateIssuedByInternalCaConfigDetails
	}{
		"ISSUED_BY_INTERNAL_CA",
		(MarshalTypeCreateCertificateIssuedByInternalCaConfigDetails)(m),
	}

	return json.Marshal(&s)
}
