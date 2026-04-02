// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fusion Applications Environment Management API
//
// Use the Fusion Applications Environment Management API to manage the environments where your Fusion Applications run. For more information, see the Fusion Applications Environment Management documentation (https://docs.oracle.com/iaas/Content/fusion-applications/home.htm).
//

package fusionapps

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CertificateInfo Properties of certificate
type CertificateInfo struct {

	// Fully qualified host name
	CommonName *string `mandatory:"false" json:"commonName"`

	// List of subject alternative names, comma separated
	SubjectAltNames []string `mandatory:"false" json:"subjectAltNames"`

	// Company name
	OrganizationName *string `mandatory:"false" json:"organizationName"`

	// Company section
	OrganizationUnit *string `mandatory:"false" json:"organizationUnit"`

	// City
	Locality *string `mandatory:"false" json:"locality"`

	// State or province
	State *string `mandatory:"false" json:"state"`

	// Country name
	Country *string `mandatory:"false" json:"country"`

	// Email address
	EmailAddress *string `mandatory:"false" json:"emailAddress"`

	// Certificate signing request
	OriginCsr *string `mandatory:"false" json:"originCsr"`

	// Akamai Certificate signing request
	AkamaiCsr *string `mandatory:"false" json:"akamaiCsr"`

	// Dv Cert instruction to validate domain, e.g. set DNS token or HTTP token, etc
	OriginDvCertInstruction *string `mandatory:"false" json:"originDvCertInstruction"`

	// Akamai Dv Cert instruction to validate domain, e.g. set DNS token or HTTP token, etc
	AkamaiDvCertInstruction *string `mandatory:"false" json:"akamaiDvCertInstruction"`

	// First name
	FirstName *string `mandatory:"false" json:"firstName"`

	// Last name
	LastName *string `mandatory:"false" json:"lastName"`

	// Phone number
	PhoneNumber *string `mandatory:"false" json:"phoneNumber"`

	// Postal code
	PostalCode *string `mandatory:"false" json:"postalCode"`
}

func (m CertificateInfo) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CertificateInfo) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
