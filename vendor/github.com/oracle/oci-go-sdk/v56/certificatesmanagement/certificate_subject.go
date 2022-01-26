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

// CertificateSubject The subject of the certificate, which is a distinguished name that identifies the entity that owns the public key in the certificate.
type CertificateSubject struct {

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
}

func (m CertificateSubject) String() string {
	return common.PointerString(m)
}
