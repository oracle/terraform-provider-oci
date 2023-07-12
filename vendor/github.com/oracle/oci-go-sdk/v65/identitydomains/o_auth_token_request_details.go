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

// OAuthTokenRequestDetails Generate the Access Token in JSON Web Token format (JWT).
type OAuthTokenRequestDetails struct {

	// Grant type by which a client requests an Access Token
	GrantType *string `mandatory:"true" json:"grant_type"`

	// Scope for which the Access Token is requested. For the <b>refresh_token</b> grant type, scope is optional.
	Scope *string `mandatory:"true" json:"scope"`

	// Name of the user who wants to access the scope (only when using the <b>Password</b> grant flow)
	Username *string `mandatory:"false" json:"username"`

	// Assertion of the client (only in <b>client assertion</b> cases)
	ClientAssertion *string `mandatory:"false" json:"client_assertion"`

	// Refresh Token that is generated using the <b>offline_access</b> scope (only in the <b>Refresh Token</b> grant flow)
	RefreshToken *string `mandatory:"false" json:"refresh_token"`

	// Client assertion type (only in <b>client assertion</b> cases)
	ClientAssertionType *string `mandatory:"false" json:"client_assertion_type"`

	// Password of the user (only when using the <b>Password</b> grant flow)
	Password *string `mandatory:"false" json:"password"`

	// Assertion of user (only in the <b>assertion</b> grant flow)
	Assertion *string `mandatory:"false" json:"assertion"`

	// Redirect URI where the response is sent (used in the Authorization or Implicit (3-legged) grant flow)
	RedirectUri *string `mandatory:"false" json:"redirect_uri"`

	// Authorization Code that is generated during the call to the Authorize endpoint (only in the Authorization (3-legged) grant flow)
	Code *string `mandatory:"false" json:"code"`

	// Unique identifier for the client (only in <b>client assertion</b> cases)
	ClientId *string `mandatory:"false" json:"client_id"`

	// Subject token representing the subject (only in <b>token exchange</b> cases)
	SubjectToken *string `mandatory:"false" json:"subject_token"`

	// Requested token type (only in <b>token exchange</b> cases)
	RequestedTokenType *string `mandatory:"false" json:"requested_token_type"`
}

func (m OAuthTokenRequestDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OAuthTokenRequestDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
