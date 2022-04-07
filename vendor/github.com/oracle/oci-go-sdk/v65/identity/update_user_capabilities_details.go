// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, policies, and identity domains.
//

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateUserCapabilitiesDetails The representation of UpdateUserCapabilitiesDetails
type UpdateUserCapabilitiesDetails struct {

	// Indicates if the user can log in to the console.
	CanUseConsolePassword *bool `mandatory:"false" json:"canUseConsolePassword"`

	// Indicates if the user can use API keys.
	CanUseApiKeys *bool `mandatory:"false" json:"canUseApiKeys"`

	// Indicates if the user can use SWIFT passwords / auth tokens.
	CanUseAuthTokens *bool `mandatory:"false" json:"canUseAuthTokens"`

	// Indicates if the user can use SMTP passwords.
	CanUseSmtpCredentials *bool `mandatory:"false" json:"canUseSmtpCredentials"`

	// Indicates if the user can use DB passwords.
	CanUseDBCredentials *bool `mandatory:"false" json:"canUseDBCredentials"`

	// Indicates if the user can use SigV4 symmetric keys.
	CanUseCustomerSecretKeys *bool `mandatory:"false" json:"canUseCustomerSecretKeys"`

	// Indicates if the user can use OAuth2 credentials and tokens.
	CanUseOAuth2ClientCredentials *bool `mandatory:"false" json:"canUseOAuth2ClientCredentials"`
}

func (m UpdateUserCapabilitiesDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateUserCapabilitiesDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
