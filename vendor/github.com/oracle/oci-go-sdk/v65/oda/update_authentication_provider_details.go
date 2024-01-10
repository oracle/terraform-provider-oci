// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Digital Assistant Service Instance API
//
// API to create and maintain Oracle Digital Assistant service instances.
//

package oda

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateAuthenticationProviderDetails Properties to update an Authentication Provider.
type UpdateAuthenticationProviderDetails struct {

	// The IDPs URL for requesting access tokens.
	TokenEndpointUrl *string `mandatory:"false" json:"tokenEndpointUrl"`

	// The IDPs URL for the page that users authenticate with by entering the user name and password.
	AuthorizationEndpointUrl *string `mandatory:"false" json:"authorizationEndpointUrl"`

	// A shortened version of the authorization URL, which you can get from a URL shortener service (one that allows
	// you to send query parameters).  You might need this because the generated authorization-code-request URL
	// could be too long for SMS and older smart phones.
	ShortAuthorizationCodeRequestUrl *string `mandatory:"false" json:"shortAuthorizationCodeRequestUrl"`

	// If you want to revoke all the refresh tokens and access tokens of the logged-in user from a dialog flow, then
	// you need the IDP's revoke refresh token URL. If you provide this URL, then you can use the System.OAuth2ResetTokens
	// component to revoke the user's tokens for this service.
	RevokeTokenEndpointUrl *string `mandatory:"false" json:"revokeTokenEndpointUrl"`

	// The client ID for the IDP application (OAuth Client) that was registered as described in Identity Provider Registration.
	// With Microsoft identity platform, use the application ID.
	ClientId *string `mandatory:"false" json:"clientId"`

	// The client secret for the IDP application (OAuth Client) that was registered as described in Identity Provider
	// Registration. With Microsoft identity platform, use the application secret.
	ClientSecret *string `mandatory:"false" json:"clientSecret"`

	// A space-separated list of the scopes that must be included when Digital Assistant requests an access token from
	// the provider. Include all the scopes that are required to access the resources. If refresh tokens are enabled,
	// include the scope that’s necessary to get the refresh token (typically offline_access).
	Scopes *string `mandatory:"false" json:"scopes"`

	// The access-token profile claim to use to identify the user.
	SubjectClaim *string `mandatory:"false" json:"subjectClaim"`

	// The number of days to keep the refresh token in the Digital Assistant cache.
	RefreshTokenRetentionPeriodInDays *int `mandatory:"false" json:"refreshTokenRetentionPeriodInDays"`

	// The OAuth Redirect URL.
	RedirectUrl *string `mandatory:"false" json:"redirectUrl"`

	// Simple key-value pair that is applied without any predefined name, type, or scope.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateAuthenticationProviderDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateAuthenticationProviderDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
