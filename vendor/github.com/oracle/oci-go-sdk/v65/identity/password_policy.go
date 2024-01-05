// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// Use the Identity and Access Management Service API to manage users, groups, identity domains, compartments, policies, tagging, and limits. For information about managing users, groups, compartments, and policies, see Identity and Access Management (without identity domains) (https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm). For information about tagging and service limits, see Tagging (https://docs.cloud.oracle.com/iaas/Content/Tagging/Concepts/taggingoverview.htm) and Service Limits (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/servicelimits.htm). For information about creating, modifying, and deleting identity domains, see Identity and Access Management (with identity domains) (https://docs.cloud.oracle.com/iaas/Content/Identity/home.htm).
//

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PasswordPolicy Password policy, currently set for the given compartment.
type PasswordPolicy struct {

	// Minimum password length required.
	MinimumPasswordLength *int `mandatory:"false" json:"minimumPasswordLength"`

	// At least one uppercase character required.
	IsUppercaseCharactersRequired *bool `mandatory:"false" json:"isUppercaseCharactersRequired"`

	// At least one lower case character required.
	IsLowercaseCharactersRequired *bool `mandatory:"false" json:"isLowercaseCharactersRequired"`

	// At least one numeric character required.
	IsNumericCharactersRequired *bool `mandatory:"false" json:"isNumericCharactersRequired"`

	// At least one special character required.
	IsSpecialCharactersRequired *bool `mandatory:"false" json:"isSpecialCharactersRequired"`

	// User name is allowed to be part of the password.
	IsUsernameContainmentAllowed *bool `mandatory:"false" json:"isUsernameContainmentAllowed"`
}

func (m PasswordPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PasswordPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
