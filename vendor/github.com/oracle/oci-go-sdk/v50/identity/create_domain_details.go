// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// CreateDomainDetails Create a domain details
type CreateDomainDetails struct {

	// The OCID of the Compartment where domain is created
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The mutable display name of the domain.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Domain entity description
	Description *string `mandatory:"true" json:"description"`

	// The region's name. See Regions and Availability Domains (https://docs.cloud.oracle.com/Content/General/Concepts/regions.htm)
	// for the full list of supported region names.
	// Example: `us-phoenix-1`
	HomeRegion *string `mandatory:"true" json:"homeRegion"`

	// The License type of Domain
	LicenseType *string `mandatory:"true" json:"licenseType"`

	// Indicates whether domain is hidden on login screen or not.
	IsHiddenOnLogin *bool `mandatory:"false" json:"isHiddenOnLogin"`

	// The admin first name
	AdminFirstName *string `mandatory:"false" json:"adminFirstName"`

	// The admin last name
	AdminLastName *string `mandatory:"false" json:"adminLastName"`

	// The admin user name
	AdminUserName *string `mandatory:"false" json:"adminUserName"`

	// The admin email address
	AdminEmail *string `mandatory:"false" json:"adminEmail"`

	// Indicates if admin user created in IDCS stripe would like to receive notification like welcome email
	// or not.
	// Required field only if admin information is provided, otherwise optional.
	IsNotificationBypassed *bool `mandatory:"false" json:"isNotificationBypassed"`

	// Optional field to indicate whether users in the domain are required to have a primary email address or not
	// Defaults to true
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
