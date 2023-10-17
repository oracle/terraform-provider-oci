// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity Domains API
//
// Use the Identity Domains API to manage resources within an identity domain, for example, users, dynamic resource groups, groups, and identity providers. For information about managing resources within identity domains, see Identity and Access Management (with identity domains) (https://docs.oracle.com/iaas/Content/Identity/home.htm). This REST API is SCIM compliant.
// Use the table of contents and search tool to explore the Identity Domains API.
//

package identitydomains

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SettingsCertificateValidation Certificate Validation Config
// **Added In:** 2010242156
// **SCIM++ Properties:**
//   - caseExact: false
//   - multiValued: false
//   - mutability: readWrite
//   - required: false
//   - returned: default
//   - type: complex
//   - uniqueness: none
type SettingsCertificateValidation struct {

	// CRL is enabled Configuration
	// **Added In:** 2010242156
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	CrlEnabled *bool `mandatory:"false" json:"crlEnabled"`

	// Use CRL as Fallback.
	// **Added In:** 2010242156
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	CrlCheckOnOCSPFailureEnabled *bool `mandatory:"false" json:"crlCheckOnOCSPFailureEnabled"`

	// CRL Location.
	// **Added In:** 2010242156
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	CrlLocation *string `mandatory:"false" json:"crlLocation"`

	// The CRL refresh interval in minutes
	// **Added In:** 2010242156
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	CrlRefreshInterval *int `mandatory:"false" json:"crlRefreshInterval"`

	// OCSP is enabled Configuration
	// **Added In:** 2010242156
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	OcspEnabled *bool `mandatory:"false" json:"ocspEnabled"`

	// OCSP Accept unknown response status from ocsp responder.
	// **Added In:** 2010242156
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	OcspUnknownResponseStatusAllowed *bool `mandatory:"false" json:"ocspUnknownResponseStatusAllowed"`

	// OCSP Responder URL
	// **Added In:** 2010242156
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	OcspResponderURL *string `mandatory:"false" json:"ocspResponderURL"`

	// This setting says, OCSP Responder URL present in the issued certificate must be used. Otherwise, OCSP Responder URL from IDP or Settings.
	// **Added In:** 2010242156
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	OcspSettingsResponderURLPreferred *bool `mandatory:"false" json:"ocspSettingsResponderURLPreferred"`

	// The OCSP Timeout duration in minutes
	// **Added In:** 2010242156
	// **SCIM++ Properties:**
	//  - idcsMaxValue: 10
	//  - idcsMinValue: 1
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	OcspTimeoutDuration *int `mandatory:"false" json:"ocspTimeoutDuration"`

	// OCSP Signing Certificate Alias
	// **Added In:** 2010242156
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	OcspSigningCertificateAlias *string `mandatory:"false" json:"ocspSigningCertificateAlias"`
}

func (m SettingsCertificateValidation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SettingsCertificateValidation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
