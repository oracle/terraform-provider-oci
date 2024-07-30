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

// AppExtensionRadiusAppApp This extension defines attributes specific to Apps that represent instances of Radius App.
type AppExtensionRadiusAppApp struct {

	// This is the IP address of the RADIUS Client like Oracle Database server. It can be only IP address and not hostname.
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	ClientIP *string `mandatory:"true" json:"clientIP"`

	// This is the port of RADIUS Proxy which RADIUS client will connect to.
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	Port *string `mandatory:"true" json:"port"`

	// Secret key used to secure communication between RADIUS Proxy and RADIUS client
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	SecretKey *string `mandatory:"true" json:"secretKey"`

	// Indicates to include groups in RADIUS response
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: boolean
	IncludeGroupInResponse *bool `mandatory:"true" json:"includeGroupInResponse"`

	// Secret key used to secure communication between RADIUS Proxy and RADIUS client. This will be available only for few releases for an internal migration requirement. Use secretKey attribute instead of this attribute for all other requirements.
	// **Added In:** 2306131901
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - idcsSensitive: encrypt
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: never
	//  - type: string
	SecretKeyTemporary *string `mandatory:"false" json:"secretKeyTemporary"`

	// If true, capture the client IP address from the RADIUS request packet. IP Address is used for auditing, policy-evaluation and country-code calculation.
	// **Added In:** 2205120021
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	CaptureClientIp *bool `mandatory:"false" json:"captureClientIp"`

	// Value consists of type of RADIUS App. Type can be Oracle Database, VPN etc
	// **Added In:** 2205120021
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	TypeOfRadiusApp *string `mandatory:"false" json:"typeOfRadiusApp"`

	// The name of the attribute that contains the Internet Protocol address of the end-user.
	// **Added In:** 2205120021
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	EndUserIPAttribute *string `mandatory:"false" json:"endUserIPAttribute"`

	// ID used to identify a particular vendor.
	// **Added In:** 2205120021
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	RadiusVendorSpecificId *string `mandatory:"false" json:"radiusVendorSpecificId"`

	// Vendor-specific identifier of the attribute in the RADIUS response that will contain the end-user's country code. This is an integer-value in the range 1 to 255
	// **Added In:** 2205120021
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	CountryCodeResponseAttributeId *string `mandatory:"false" json:"countryCodeResponseAttributeId"`

	// RADIUS attribute that RADIUS-enabled system uses to pass the group membership
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	GroupMembershipRadiusAttribute *string `mandatory:"false" json:"groupMembershipRadiusAttribute"`

	// Configure the responseFormat based on vendor in order to pass it to RADIUS infra
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	ResponseFormat *string `mandatory:"false" json:"responseFormat"`

	// The delimiter used if group membership responseFormat is a delimited list instead of repeating attributes
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	ResponseFormatDelimiter *string `mandatory:"false" json:"responseFormatDelimiter"`

	// Configure the groupNameFormat based on vendor in order to pass it to RADIUS infra
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	GroupNameFormat *string `mandatory:"false" json:"groupNameFormat"`

	// Indicates if password and OTP are passed in the same sign-in request or not.
	// **Added In:** 2205120021
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	PasswordAndOtpTogether *bool `mandatory:"false" json:"passwordAndOtpTogether"`

	// In a successful authentication response, Oracle Identity Cloud Service will pass user's group information restricted to groups persisted in this attribute, in the specified RADIUS attribute.
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [value]
	//  - idcsSearchable: true
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: request
	//  - type: complex
	//  - uniqueness: none
	GroupMembershipToReturn []AppGroupMembershipToReturn `mandatory:"false" json:"groupMembershipToReturn"`
}

func (m AppExtensionRadiusAppApp) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AppExtensionRadiusAppApp) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
