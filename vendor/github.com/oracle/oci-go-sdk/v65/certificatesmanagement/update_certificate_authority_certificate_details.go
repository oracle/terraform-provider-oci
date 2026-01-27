// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// UpdateCertificateAuthorityCertificateDetails The details of the request to update the certificate authority (CA) with a signed certificate for the latest CA version.
type UpdateCertificateAuthorityCertificateDetails struct {

	// The externally signed certificate (in PEM format) for the subordinate certificate authority (CA).
	CertificatePem *string `mandatory:"true" json:"certificatePem"`
}

func (m UpdateCertificateAuthorityCertificateDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateCertificateAuthorityCertificateDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateCertificateAuthorityCertificateDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateCertificateAuthorityCertificateDetails UpdateCertificateAuthorityCertificateDetails
	s := struct {
		DiscriminatorParam string `json:"actionType"`
		MarshalTypeUpdateCertificateAuthorityCertificateDetails
	}{
		"UPDATE_CERTIFICATE",
		(MarshalTypeUpdateCertificateAuthorityCertificateDetails)(m),
	}

	return json.Marshal(&s)
}
