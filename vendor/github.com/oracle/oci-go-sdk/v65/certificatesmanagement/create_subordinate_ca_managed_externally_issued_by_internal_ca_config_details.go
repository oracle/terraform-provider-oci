// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// CreateSubordinateCaManagedExternallyIssuedByInternalCaConfigDetails The details for creating an externally managed subordinate certificate authority (CA) which is issued by an internal private CA.
type CreateSubordinateCaManagedExternallyIssuedByInternalCaConfigDetails struct {

	// The OCID of the issuer private CA.
	IssuerCertificateAuthorityId *string `mandatory:"true" json:"issuerCertificateAuthorityId"`

	// The certificate signing request (in PEM format).
	CsrPem *string `mandatory:"true" json:"csrPem"`

	// The name of the CA version. When the value is not null, a name is unique across versions of a given CA.
	VersionName *string `mandatory:"false" json:"versionName"`

	Validity *Validity `mandatory:"false" json:"validity"`
}

// GetVersionName returns VersionName
func (m CreateSubordinateCaManagedExternallyIssuedByInternalCaConfigDetails) GetVersionName() *string {
	return m.VersionName
}

func (m CreateSubordinateCaManagedExternallyIssuedByInternalCaConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateSubordinateCaManagedExternallyIssuedByInternalCaConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateSubordinateCaManagedExternallyIssuedByInternalCaConfigDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateSubordinateCaManagedExternallyIssuedByInternalCaConfigDetails CreateSubordinateCaManagedExternallyIssuedByInternalCaConfigDetails
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeCreateSubordinateCaManagedExternallyIssuedByInternalCaConfigDetails
	}{
		"SUBORDINATE_CA_MANAGED_EXTERNALLY_ISSUED_BY_INTERNAL_CA",
		(MarshalTypeCreateSubordinateCaManagedExternallyIssuedByInternalCaConfigDetails)(m),
	}

	return json.Marshal(&s)
}
