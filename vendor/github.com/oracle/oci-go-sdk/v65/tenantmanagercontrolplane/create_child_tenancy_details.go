// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// Use the Organizations API to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and organization resources. For more information, see Organization Management Overview (https://docs.oracle.com/iaas/Content/General/Concepts/organization_management_overview.htm).
//

package tenantmanagercontrolplane

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateChildTenancyDetails The parameters for creating a child tenancy.
type CreateChildTenancyDetails struct {

	// The tenancy ID of the parent tenancy.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The tenancy name to use for the child tenancy.
	TenancyName *string `mandatory:"true" json:"tenancyName"`

	// The home region to use for the child tenancy. This must be a region where the parent tenancy is subscribed.
	HomeRegion *string `mandatory:"true" json:"homeRegion"`

	// Email address of the child tenancy administrator.
	AdminEmail *string `mandatory:"true" json:"adminEmail"`

	// The name to use for the administrator policy in the child tenancy. Must contain only letters and underscores.
	PolicyName *string `mandatory:"false" json:"policyName"`

	// The governance status of the child tenancy.
	GovernanceStatus GovernanceStatusEnum `mandatory:"false" json:"governanceStatus,omitempty"`

	// OCID of the subscription that needs to be assigned to the child tenancy.
	SubscriptionId *string `mandatory:"false" json:"subscriptionId"`
}

func (m CreateChildTenancyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateChildTenancyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingGovernanceStatusEnum(string(m.GovernanceStatus)); !ok && m.GovernanceStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GovernanceStatus: %s. Supported values are: %s.", m.GovernanceStatus, strings.Join(GetGovernanceStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
