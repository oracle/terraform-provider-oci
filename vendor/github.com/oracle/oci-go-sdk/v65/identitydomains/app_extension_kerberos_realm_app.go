// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity Domains API
//
// Use the Identity Domains API to manage resources within an identity domain, for example, users, dynamic resource groups, groups, and identity providers. For information about managing resources within identity domains, see Identity and Access Management (with identity domains) (https://docs.oracle.com/iaas/Content/Identity/home.htm).
// Use this pattern to construct endpoints for identity domains: `https://<domainURL>/admin/v1/`. See Finding an Identity Domain URL (https://docs.oracle.com/en-us/iaas/Content/Identity/api-getstarted/locate-identity-domain-url.htm) to locate the domain URL you need.
// Use the table of contents and search tool to explore the Identity Domains API.
//

package identitydomains

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AppExtensionKerberosRealmApp Kerberos Realm
type AppExtensionKerberosRealmApp struct {

	// The name of the Kerberos Realm that this App uses for authentication.
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: request
	//  - type: string
	//  - uniqueness: none
	RealmName *string `mandatory:"false" json:"realmName"`

	// The primary key that the system should use to encrypt artifacts that are specific to this Kerberos realm -- for example, to encrypt the Principal Key in each KerberosRealmUser.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - idcsSensitive: none
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: request
	//  - type: string
	//  - uniqueness: none
	MasterKey *string `mandatory:"false" json:"masterKey"`

	// The type of salt that the system will use to encrypt Kerberos-specific artifacts of this App unless another type of salt is specified.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: request
	//  - type: string
	//  - uniqueness: none
	DefaultEncryptionSaltType *string `mandatory:"false" json:"defaultEncryptionSaltType"`

	// The types of salt that are available for the system to use when encrypting Kerberos-specific artifacts for this App.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: request
	//  - type: string
	//  - uniqueness: none
	SupportedEncryptionSaltTypes []string `mandatory:"false" json:"supportedEncryptionSaltTypes"`

	// Ticket Flags
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: request
	//  - type: integer
	//  - uniqueness: none
	TicketFlags *int `mandatory:"false" json:"ticketFlags"`

	// Max Ticket Life in seconds
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: request
	//  - type: integer
	//  - uniqueness: none
	MaxTicketLife *int `mandatory:"false" json:"maxTicketLife"`

	// Max Renewable Age in seconds
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: request
	//  - type: integer
	//  - uniqueness: none
	MaxRenewableAge *int `mandatory:"false" json:"maxRenewableAge"`
}

func (m AppExtensionKerberosRealmApp) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AppExtensionKerberosRealmApp) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
