// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// CertificateAuthorityIssuanceExpiryRule A rule that enforces how long certificates or certificate authorities (CAs) issued by this particular CA are valid.
// You must include either or both `leafCertificateMaxValidityDuration` and `certificateAuthorityMaxValidityDuration`.
type CertificateAuthorityIssuanceExpiryRule struct {

	// A property indicating the maximum validity duration, in days, of leaf certificates issued by this CA.
	// Expressed in ISO 8601 (https://en.wikipedia.org/wiki/ISO_8601#Time_intervals) format.
	LeafCertificateMaxValidityDuration *string `mandatory:"false" json:"leafCertificateMaxValidityDuration"`

	// A property indicating the maximum validity duration, in days, of subordinate CA's issued by this CA.
	// Expressed in ISO 8601 (https://en.wikipedia.org/wiki/ISO_8601#Time_intervals) format.
	CertificateAuthorityMaxValidityDuration *string `mandatory:"false" json:"certificateAuthorityMaxValidityDuration"`
}

func (m CertificateAuthorityIssuanceExpiryRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CertificateAuthorityIssuanceExpiryRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CertificateAuthorityIssuanceExpiryRule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCertificateAuthorityIssuanceExpiryRule CertificateAuthorityIssuanceExpiryRule
	s := struct {
		DiscriminatorParam string `json:"ruleType"`
		MarshalTypeCertificateAuthorityIssuanceExpiryRule
	}{
		"CERTIFICATE_AUTHORITY_ISSUANCE_EXPIRY_RULE",
		(MarshalTypeCertificateAuthorityIssuanceExpiryRule)(m),
	}

	return json.Marshal(&s)
}
