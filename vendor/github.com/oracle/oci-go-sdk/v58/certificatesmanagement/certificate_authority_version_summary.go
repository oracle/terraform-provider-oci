// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Certificates Service Management API
//
// API for managing certificates.
//

package certificatesmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// CertificateAuthorityVersionSummary The metadata details of the certificate authority (CA) version. This summary object does not contain the CA contents.
type CertificateAuthorityVersionSummary struct {

	// The OCID of the CA.
	CertificateAuthorityId *string `mandatory:"true" json:"certificateAuthorityId"`

	// A optional property indicating when the CA version was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-04-03T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The version number of the CA.
	VersionNumber *int64 `mandatory:"true" json:"versionNumber"`

	// A list of rotation states for this CA version.
	Stages []VersionStageEnum `mandatory:"true" json:"stages"`

	// The version number of the issuing CA.
	IssuerCaVersionNumber *int64 `mandatory:"false" json:"issuerCaVersionNumber"`

	// A unique certificate identifier used in certificate revocation tracking, formatted as octets.
	// Example: `03 AC FC FA CC B3 CB 02 B8 F8 DE F5 85 E7 7B FF`
	SerialNumber *string `mandatory:"false" json:"serialNumber"`

	// The name of the CA version. When this value is not null, the name is unique across CA versions for a given CA.
	VersionName *string `mandatory:"false" json:"versionName"`

	// An optional property indicating when to delete the CA version, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-04-03T21:10:29.600Z`
	TimeOfDeletion *common.SDKTime `mandatory:"false" json:"timeOfDeletion"`

	Validity *Validity `mandatory:"false" json:"validity"`

	RevocationStatus *RevocationStatus `mandatory:"false" json:"revocationStatus"`
}

func (m CertificateAuthorityVersionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CertificateAuthorityVersionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
