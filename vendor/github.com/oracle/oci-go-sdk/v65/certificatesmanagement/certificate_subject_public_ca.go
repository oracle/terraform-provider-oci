// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Certificates Service Management API
//
// API for managing certificates.
//

package certificatesmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CertificateSubjectPublicCa The details of the configuration for creating a certificate which is issued by a public certificate authority (CA).
type CertificateSubjectPublicCa struct {

	// Common name or fully-qualified domain name (RDN CN).
	CommonName *string `mandatory:"true" json:"commonName"`

	// Country name (RDN C).
	Country *string `mandatory:"false" json:"country"`

	// Domain component (RDN DC).
	DomainComponent *string `mandatory:"false" json:"domainComponent"`

	// Distinguished name qualifier(RDN DNQ).
	DistinguishedNameQualifier *string `mandatory:"false" json:"distinguishedNameQualifier"`

	// Personal generational qualifier (for example, Sr., Jr. 3rd, or IV).
	GenerationQualifier *string `mandatory:"false" json:"generationQualifier"`

	// Personal given name (RDN G or GN).
	GivenName *string `mandatory:"false" json:"givenName"`

	// Personal initials.
	Initials *string `mandatory:"false" json:"initials"`

	// Locality (RDN L).
	LocalityName *string `mandatory:"false" json:"localityName"`

	// Organization (RDN O).
	Organization *string `mandatory:"false" json:"organization"`

	// Organizational unit (RDN OU).
	OrganizationalUnit *string `mandatory:"false" json:"organizationalUnit"`

	// Subject pseudonym.
	Pseudonym *string `mandatory:"false" json:"pseudonym"`

	// Unique subject identifier, which is not the same as the certificate serial number (RDN SERIALNUMBER).
	SerialNumber *string `mandatory:"false" json:"serialNumber"`

	// State or province name (RDN ST or S).
	StateOrProvinceName *string `mandatory:"false" json:"stateOrProvinceName"`

	// Street address (RDN STREET).
	Street *string `mandatory:"false" json:"street"`

	// Personal surname (RDN SN).
	Surname *string `mandatory:"false" json:"surname"`

	// Title (RDN T or TITLE).
	Title *string `mandatory:"false" json:"title"`

	// User ID (RDN UID).
	UserId *string `mandatory:"false" json:"userId"`

	// DNS zone Name or OCID.
	DnsZoneNameOrId *string `mandatory:"false" json:"dnsZoneNameOrId"`
}

func (m CertificateSubjectPublicCa) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CertificateSubjectPublicCa) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
