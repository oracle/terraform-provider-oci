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

// RevokeCertificateVersionDetails The details for revoking a certificate version.
type RevokeCertificateVersionDetails struct {

	// The reason that the certificate or certificate authority was revoked.
	RevocationReason RevocationReasonEnum `mandatory:"false" json:"revocationReason,omitempty"`
}

func (m RevokeCertificateVersionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RevokeCertificateVersionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRevocationReasonEnum(string(m.RevocationReason)); !ok && m.RevocationReason != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RevocationReason: %s. Supported values are: %s.", m.RevocationReason, strings.Join(GetRevocationReasonEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
