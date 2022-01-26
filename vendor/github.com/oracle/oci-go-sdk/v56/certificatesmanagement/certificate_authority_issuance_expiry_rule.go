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
