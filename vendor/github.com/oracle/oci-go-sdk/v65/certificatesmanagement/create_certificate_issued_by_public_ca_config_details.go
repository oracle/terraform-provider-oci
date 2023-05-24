// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateCertificateIssuedByPublicCaConfigDetails The details of the configuration for creating a certificate which is issued by a public certificate authority (CA).
type CreateCertificateIssuedByPublicCaConfigDetails struct {
	Subject *CertificateSubject `mandatory:"true" json:"subject"`

	Validity *PublicCaCertificateValidity `mandatory:"true" json:"validity"`

	// A name for the certificate. When the value is not null, a name is unique across versions of a given certificate.
	VersionName *string `mandatory:"false" json:"versionName"`

	// A list of subject alternative names. A subject alternative name specifies the domain names, including subdomains, and IP addresses covered by the certificates issued by this CA.
	SubjectAlternativeNames []CertificateSubjectAlternativeName `mandatory:"false" json:"subjectAlternativeNames"`

	// The name of the profile used to create the certificate, which depends on the type of certificate you need.
	CertificateProfileType CertificateProfileTypeEnum `mandatory:"true" json:"certificateProfileType"`

	// The name of the public CA issuing the certificate.
	PublicCertificateAuthority PublicCertificateAuthorityEnum `mandatory:"false" json:"publicCertificateAuthority,omitempty"`

	// The algorithm to use to create key pairs.
	KeyAlgorithm KeyAlgorithmEnum `mandatory:"false" json:"keyAlgorithm,omitempty"`
}

//GetVersionName returns VersionName
func (m CreateCertificateIssuedByPublicCaConfigDetails) GetVersionName() *string {
	return m.VersionName
}

func (m CreateCertificateIssuedByPublicCaConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateCertificateIssuedByPublicCaConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCertificateProfileTypeEnum(string(m.CertificateProfileType)); !ok && m.CertificateProfileType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CertificateProfileType: %s. Supported values are: %s.", m.CertificateProfileType, strings.Join(GetCertificateProfileTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPublicCertificateAuthorityEnum(string(m.PublicCertificateAuthority)); !ok && m.PublicCertificateAuthority != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PublicCertificateAuthority: %s. Supported values are: %s.", m.PublicCertificateAuthority, strings.Join(GetPublicCertificateAuthorityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingKeyAlgorithmEnum(string(m.KeyAlgorithm)); !ok && m.KeyAlgorithm != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for KeyAlgorithm: %s. Supported values are: %s.", m.KeyAlgorithm, strings.Join(GetKeyAlgorithmEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateCertificateIssuedByPublicCaConfigDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateCertificateIssuedByPublicCaConfigDetails CreateCertificateIssuedByPublicCaConfigDetails
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeCreateCertificateIssuedByPublicCaConfigDetails
	}{
		"ISSUED_BY_PUBLIC_CA",
		(MarshalTypeCreateCertificateIssuedByPublicCaConfigDetails)(m),
	}

	return json.Marshal(&s)
}
