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

// AllowedDomainLicenseTypeSummary (For tenancies that support identity domains) The 'AllowedDomainLicenseTypeSummary' object contains information about the license type of the identity domain.
type AllowedDomainLicenseTypeSummary struct {

	// The license type name.
	// Example: "Oracle Apps Premium"
	Name *string `mandatory:"true" json:"name"`

	// The license type identifier.
	// Example: "oracle-apps-premium"
	LicenseType *string `mandatory:"true" json:"licenseType"`

	// The license type description.
	Description *string `mandatory:"true" json:"description"`
}

func (m AllowedDomainLicenseTypeSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AllowedDomainLicenseTypeSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
