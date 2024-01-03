// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RoleSummary The details of a role fetched from the database.
type RoleSummary struct {

	// The name of the role.
	RoleName *string `mandatory:"true" json:"roleName"`

	// Type of authentication.
	AuthenticationType *string `mandatory:"true" json:"authenticationType"`

	// Is password required.
	IsPasswordRequired *bool `mandatory:"false" json:"isPasswordRequired"`

	// Is the role common.
	IsCommon *bool `mandatory:"false" json:"isCommon"`

	// Is the role oracle maintained.
	IsOracleMaintained *bool `mandatory:"false" json:"isOracleMaintained"`

	// Is the role inherited.
	IsInherited *bool `mandatory:"false" json:"isInherited"`

	// Is the role implicit.
	IsImplicit *bool `mandatory:"false" json:"isImplicit"`
}

func (m RoleSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RoleSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
