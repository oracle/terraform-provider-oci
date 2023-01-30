// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// UpdateOAuth2ClientCredentialDetails The representation of UpdateOAuth2ClientCredentialDetails
type UpdateOAuth2ClientCredentialDetails struct {

	// Description of the oauth credential to help user differentiate them.
	Description *string `mandatory:"true" json:"description"`

	// Allowed scopes for the given oauth credential.
	Scopes []FullyQualifiedScope `mandatory:"true" json:"scopes"`

	// Indicate if the password to be reset or not in the update.
	IsResetPassword *bool `mandatory:"false" json:"isResetPassword"`
}

func (m UpdateOAuth2ClientCredentialDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateOAuth2ClientCredentialDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
