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

// CertificateRenewalRule A rule that imposes constraints on certificate renewal.
type CertificateRenewalRule struct {

	// A property specifying how often, in days, a certificate should be renewed.
	// Expressed in ISO 8601 (https://en.wikipedia.org/wiki/ISO_8601#Time_intervals) format.
	RenewalInterval *string `mandatory:"true" json:"renewalInterval"`

	// A property specifying the period of time, in days, before the certificate's targeted renewal that the process should occur.
	// Expressed in ISO 8601 (https://en.wikipedia.org/wiki/ISO_8601#Time_intervals) format.
	AdvanceRenewalPeriod *string `mandatory:"true" json:"advanceRenewalPeriod"`
}

func (m CertificateRenewalRule) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CertificateRenewalRule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCertificateRenewalRule CertificateRenewalRule
	s := struct {
		DiscriminatorParam string `json:"ruleType"`
		MarshalTypeCertificateRenewalRule
	}{
		"CERTIFICATE_RENEWAL_RULE",
		(MarshalTypeCertificateRenewalRule)(m),
	}

	return json.Marshal(&s)
}
