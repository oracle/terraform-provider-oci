// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateCrossConnectLetterOfAuthorityDetails Update the details of the Letter of Authority associated with a cross-connect. Upon successful alterations
// a new updated copy of the LOA will be attached to the cross-connect.
type UpdateCrossConnectLetterOfAuthorityDetails struct {

	// A boolean flag to indicate whether to extend the expiry of the associated LOA with the provided
	// cross-connect. If un-set or set to false, it does not alter the existing expiry of the LOA.
	// On extension an updated copy of the LOA will be provided with the new expiry date. An LOA cannot
	// be extended more than 3 times.
	ShouldExtend *bool `mandatory:"false" json:"shouldExtend"`

	// A boolean flag to indicate whether to remove an attached Authorized Agent to the LOA. If this boolean flag
	// is set, an attempt will be made to remove the attached authorized agent to the LOA, if any, and any value
	// given in the field 'authorizedAgent' will be ignored. In case, of updating an existing Authorized Agent,
	// keep this flag unset and set the expected value in 'authorizedAgent'.
	ShouldRemoveAuthorizedAgent *bool `mandatory:"false" json:"shouldRemoveAuthorizedAgent"`

	// Name of a customer authorized agent which will be appended to the LOA as the field 'Authorized Agent'.
	// If the field is left un-set in the request body, no changes will be done on the LOA for Authorized Agent.
	AuthorizedAgent *string `mandatory:"false" json:"authorizedAgent"`
}

func (m UpdateCrossConnectLetterOfAuthorityDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateCrossConnectLetterOfAuthorityDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
