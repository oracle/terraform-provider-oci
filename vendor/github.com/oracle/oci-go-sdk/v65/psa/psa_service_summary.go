// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PrivateServiceAccess Control Plane API
//
// Use the PrivateServiceAccess Control Plane API to manage privateServiceAccess.
//

package psa

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PsaServiceSummary An OCI service summary, that will be used as a catalog for Private Service Access.
type PsaServiceSummary struct {

	// A unique OCI service identifier.
	// Example: `object-storage-api`
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The public facing service FQDNs, which are going to be used to access the service.
	// Example: `xyz.oraclecloud.com`
	Fqdns []string `mandatory:"true" json:"fqdns"`

	// A description of the OCI service.
	Description *string `mandatory:"false" json:"description"`

	// This optional field will indicate that whether service is IPv6 enabled.
	IsV6Enabled *bool `mandatory:"false" json:"isV6Enabled"`
}

func (m PsaServiceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PsaServiceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
