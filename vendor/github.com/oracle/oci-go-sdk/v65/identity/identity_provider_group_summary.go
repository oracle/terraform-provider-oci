// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// Use the Identity and Access Management Service API to manage users, groups, identity domains, compartments, policies, tagging, and limits. For information about managing users, groups, compartments, and policies, see Identity and Access Management (without identity domains) (https://docs.oracle.com/iaas/Content/Identity/Concepts/overview.htm). For information about tagging and service limits, see Tagging (https://docs.oracle.com/iaas/Content/Tagging/Concepts/taggingoverview.htm) and Service Limits (https://docs.oracle.com/iaas/Content/General/Concepts/servicelimits.htm). For information about creating, modifying, and deleting identity domains, see Identity and Access Management (with identity domains) (https://docs.oracle.com/iaas/Content/Identity/home.htm).
//

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// IdentityProviderGroupSummary A group created in an identity provider that can be mapped to a group in OCI
type IdentityProviderGroupSummary struct {

	// The OCID of the `IdentityProviderGroup`.
	Id *string `mandatory:"false" json:"id"`

	// The OCID of the `IdentityProvider` this group belongs to.
	IdentityProviderId *string `mandatory:"false" json:"identityProviderId"`

	// Display name of the group
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Display name of the group
	Name *string `mandatory:"false" json:"name"`

	// Identifier of the group in the identity provider
	ExternalIdentifier *string `mandatory:"false" json:"externalIdentifier"`

	// Date and time the `IdentityProviderGroup` was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Date and time the `IdentityProviderGroup` was last modified, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeModified *common.SDKTime `mandatory:"false" json:"timeModified"`
}

func (m IdentityProviderGroupSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IdentityProviderGroupSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
