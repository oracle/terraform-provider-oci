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

// UpdateCertificateIssuedByInternalCaConfigDetails The details for updating an internally managed certificate which is issued by a private certificate authority (CA).
type UpdateCertificateIssuedByInternalCaConfigDetails struct {

	// A name for the certificate version. When the value is not null, a name is unique across versions of a given certificate.
	VersionName *string `mandatory:"false" json:"versionName"`

	Validity *Validity `mandatory:"false" json:"validity"`

	// The rotation state of the certificate. The default is `CURRENT`, meaning that the certificate is currently in use. A certificate version
	// that you mark as `PENDING` is staged and available for use, but you don't yet want to rotate it into current, active use. For example,
	// you might update a certificate and mark its rotation state as `PENDING` if you haven't yet updated the certificate on the target system.
	Stage UpdateCertificateConfigDetailsStageEnum `mandatory:"false" json:"stage,omitempty"`
}

//GetVersionName returns VersionName
func (m UpdateCertificateIssuedByInternalCaConfigDetails) GetVersionName() *string {
	return m.VersionName
}

//GetStage returns Stage
func (m UpdateCertificateIssuedByInternalCaConfigDetails) GetStage() UpdateCertificateConfigDetailsStageEnum {
	return m.Stage
}

func (m UpdateCertificateIssuedByInternalCaConfigDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateCertificateIssuedByInternalCaConfigDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateCertificateIssuedByInternalCaConfigDetails UpdateCertificateIssuedByInternalCaConfigDetails
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeUpdateCertificateIssuedByInternalCaConfigDetails
	}{
		"ISSUED_BY_INTERNAL_CA",
		(MarshalTypeUpdateCertificateIssuedByInternalCaConfigDetails)(m),
	}

	return json.Marshal(&s)
}
