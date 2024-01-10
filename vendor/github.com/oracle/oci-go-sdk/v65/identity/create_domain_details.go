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

// CreateDomainDetails (For tenancies that support identity domains) Details for creating an identity domain.
type CreateDomainDetails struct {

	// The OCID of the compartment where the identity domain is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The mutable display name of the identity domain.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The identity domain description. You can have an empty description.
	Description *string `mandatory:"true" json:"description"`

	// The region's name identifier. See Regions and Availability Domains (https://docs.cloud.oracle.com/Content/General/Concepts/regions.htm)
	// for the full list of supported region names.
	// Example: `us-phoenix-1`
	HomeRegion *string `mandatory:"true" json:"homeRegion"`

	// The license type of the identity domain.
	LicenseType *string `mandatory:"true" json:"licenseType"`

	// Indicates whether the identity domain is hidden on the sign-in screen or not.
	IsHiddenOnLogin *bool `mandatory:"false" json:"isHiddenOnLogin"`

	// The administrator's first name.
	AdminFirstName *string `mandatory:"false" json:"adminFirstName"`

	// The administrator's last name.
	AdminLastName *string `mandatory:"false" json:"adminLastName"`

	// The administrator's user name.
	AdminUserName *string `mandatory:"false" json:"adminUserName"`

	// The administrator's email address.
	AdminEmail *string `mandatory:"false" json:"adminEmail"`

	// Indicates whether or not the administrator user created in the IDCS stripe would like to receive notifications like a welcome email.
	// This field is required only if admin information is provided. This field is otherwise optional.
	IsNotificationBypassed *bool `mandatory:"false" json:"isNotificationBypassed"`

	// Optional field to indicate whether users in the identity domain are required to have a primary email address or not. The default is true.
	IsPrimaryEmailRequired *bool `mandatory:"false" json:"isPrimaryEmailRequired"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateDomainDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDomainDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
