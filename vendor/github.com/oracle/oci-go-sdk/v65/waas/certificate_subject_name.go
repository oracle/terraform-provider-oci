// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CertificateSubjectName The entity to be secured by the certificate.
type CertificateSubjectName struct {

	// ISO 3166-1 alpha-2 code of the country where the organization is located. For a list of codes, see ISO's website (https://www.iso.org/obp/ui/#search/code/).
	Country *string `mandatory:"false" json:"country"`

	// The province where the organization is located.
	StateProvince *string `mandatory:"false" json:"stateProvince"`

	// The city in which the organization is located.
	Locality *string `mandatory:"false" json:"locality"`

	// The organization name.
	Organization *string `mandatory:"false" json:"organization"`

	// The field to differentiate between divisions within an organization.
	OrganizationalUnit *string `mandatory:"false" json:"organizationalUnit"`

	// The fully qualified domain name used for DNS lookups of the server.
	CommonName *string `mandatory:"false" json:"commonName"`

	// The email address of the server's administrator.
	EmailAddress *string `mandatory:"false" json:"emailAddress"`
}

func (m CertificateSubjectName) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CertificateSubjectName) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
