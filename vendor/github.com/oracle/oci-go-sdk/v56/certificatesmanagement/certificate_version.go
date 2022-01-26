// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Certificates Service Management API
//
// API for managing certificates.
//

package certificatesmanagement

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// CertificateVersion The details of the certificate version. This object does not contain the certificate contents.
type CertificateVersion struct {

	// The OCID of the certificate.
	CertificateId *string `mandatory:"true" json:"certificateId"`

	// A optional property indicating when the certificate version was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-04-03T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The version number of the certificate.
	VersionNumber *int64 `mandatory:"true" json:"versionNumber"`

	// A list of stages of this entity.
	Stages []VersionStageEnum `mandatory:"true" json:"stages"`

	// A unique certificate identifier used in certificate revocation tracking, formatted as octets.
	// Example: `03 AC FC FA CC B3 CB 02 B8 F8 DE F5 85 E7 7B FF`
	SerialNumber *string `mandatory:"false" json:"serialNumber"`

	// The version number of the issuing certificate authority (CA).
	IssuerCaVersionNumber *int64 `mandatory:"false" json:"issuerCaVersionNumber"`

	// The name of the certificate version. When the value is not null, a name is unique across versions of a given certificate.
	VersionName *string `mandatory:"false" json:"versionName"`

	// A list of subject alternative names.
	SubjectAlternativeNames []CertificateSubjectAlternativeName `mandatory:"false" json:"subjectAlternativeNames"`

	// An optional property indicating when to delete the certificate version, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-04-03T21:10:29.600Z`
	TimeOfDeletion *common.SDKTime `mandatory:"false" json:"timeOfDeletion"`

	Validity *Validity `mandatory:"false" json:"validity"`

	RevocationStatus *RevocationStatus `mandatory:"false" json:"revocationStatus"`
}

func (m CertificateVersion) String() string {
	return common.PointerString(m)
}
