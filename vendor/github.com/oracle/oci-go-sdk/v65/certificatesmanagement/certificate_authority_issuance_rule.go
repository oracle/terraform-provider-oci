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

// CertificateAuthorityIssuanceRule Issuance rules apply constraints to a certificate authority (CA) to enforce certain conditions regarding the resources it issues. For example, a path length constraint restricts how many subordinate CAs a CA can have. Meanwhile, a name constraint on certificate subject names specifies allowable namespaces for the hierarchical name forms in certificates that any CA in the certificate chain issues. You can't update the issuance rules configured for a CA after you create it.
type CertificateAuthorityIssuanceRule struct {

	// The number of levels of descendants that this certificate authority (CA) can issue. When set to zero, the CA can issue only leaf certificates. There is no limit if the constraint isn't specified.
	PathLengthConstraint *int `mandatory:"false" json:"pathLengthConstraint"`

	NameConstraint *NameConstraint `mandatory:"false" json:"nameConstraint"`

	// A property indicating the maximum validity duration of short-lived certificates issued by this CA, expressed in ISO 8601 (https://en.wikipedia.org/wiki/ISO_8601#Time_intervals) duration format. The maximum allowable value is 7 days. The shortest allowable maximum validity duration you can specify is 1 hour.
	// This is only applicable when issuanceMode is SHORT_LIVED.
	ShortLivedCertificatesMaxValidityDuration *string `mandatory:"false" json:"shortLivedCertificatesMaxValidityDuration"`

	// A property specifying whether the certificate authority (CA) issues short-lived or long-lived certificates. If not provided, the CA defaults to issuing long-lived certificates.
	IssuanceMode CaIssuanceModeEnum `mandatory:"false" json:"issuanceMode,omitempty"`
}

func (m CertificateAuthorityIssuanceRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CertificateAuthorityIssuanceRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCaIssuanceModeEnum(string(m.IssuanceMode)); !ok && m.IssuanceMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IssuanceMode: %s. Supported values are: %s.", m.IssuanceMode, strings.Join(GetCaIssuanceModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CertificateAuthorityIssuanceRule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCertificateAuthorityIssuanceRule CertificateAuthorityIssuanceRule
	s := struct {
		DiscriminatorParam string `json:"ruleType"`
		MarshalTypeCertificateAuthorityIssuanceRule
	}{
		"CERTIFICATE_AUTHORITY_ISSUANCE_RULE",
		(MarshalTypeCertificateAuthorityIssuanceRule)(m),
	}

	return json.Marshal(&s)
}
