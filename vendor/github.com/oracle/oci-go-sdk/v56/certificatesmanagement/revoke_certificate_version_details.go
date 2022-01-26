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

// RevokeCertificateVersionDetails The details for revoking a certificate version.
type RevokeCertificateVersionDetails struct {

	// The reason that the certificate or certificate authority was revoked.
	RevocationReason RevocationReasonEnum `mandatory:"false" json:"revocationReason,omitempty"`
}

func (m RevokeCertificateVersionDetails) String() string {
	return common.PointerString(m)
}
