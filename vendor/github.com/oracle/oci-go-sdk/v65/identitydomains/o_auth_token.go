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

// OAuthToken Generate the Access Token in JSON Web Token format (JWT).
type OAuthToken struct {

	// Access Token used to access the scopes
	AccessToken *string `mandatory:"true" json:"access_token"`

	// Expiry time of the Access Token in seconds
	ExpiresIn *float32 `mandatory:"true" json:"expires_in"`

	// Type of Access Token (Bearer)
	TokenType *string `mandatory:"true" json:"token_type"`

	// Refresh Token used to regenerate the Access Token (only when the <b>offline_access</b> scope is used)
	RefreshToken *string `mandatory:"false" json:"refresh_token"`

	// Identity Token generated for the associated client and user (only in 3-legged flows)
	IdToken *string `mandatory:"false" json:"id_token"`
}

func (m OAuthToken) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OAuthToken) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
