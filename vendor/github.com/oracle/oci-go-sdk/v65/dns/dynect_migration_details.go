// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DNS API
//
// API for the DNS service. Use this API to manage DNS zones, records, and other DNS resources.
// For more information, see Overview of the DNS Service (https://docs.oracle.com/iaas/Content/DNS/Concepts/dnszonemanagement.htm).
//

package dns

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DynectMigrationDetails Details specific to performing a DynECT zone migration.
type DynectMigrationDetails struct {

	// DynECT customer name the zone belongs to.
	CustomerName *string `mandatory:"true" json:"customerName"`

	// DynECT API username to perform the migration with.
	Username *string `mandatory:"true" json:"username"`

	// DynECT API password for the provided username.
	Password *string `mandatory:"true" json:"password"`

	// A map of fully-qualified domain names (FQDNs) to an array of `MigrationReplacement` objects.
	HttpRedirectReplacements map[string][]MigrationReplacement `mandatory:"false" json:"httpRedirectReplacements"`
}

func (m DynectMigrationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DynectMigrationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
