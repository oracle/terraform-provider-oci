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

// Account The basic information for an account object.
type Account struct {

	// The name of the tenant.
	AccountName *string `mandatory:"true" json:"accountName"`

	// The lifecycle state of the tenant.
	State *string `mandatory:"true" json:"state"`

	// The tenant home region.
	HomeRegion *string `mandatory:"true" json:"homeRegion"`

	// The tenant active regions.
	ActiveRegions []string `mandatory:"true" json:"activeRegions"`

	// The OCID of the tenancy.
	TenancyOcid *string `mandatory:"false" json:"tenancyOcid"`
}

func (m Account) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Account) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
