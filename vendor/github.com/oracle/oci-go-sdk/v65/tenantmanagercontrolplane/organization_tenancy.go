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

// OrganizationTenancy The information about the organization tenancy.
type OrganizationTenancy struct {

	// OCID of the tenancy.
	TenancyId *string `mandatory:"true" json:"tenancyId"`

	// The governance status of the tenancy.
	GovernanceStatus GovernanceStatusEnum `mandatory:"true" json:"governanceStatus"`

	// Name of the tenancy.
	Name *string `mandatory:"false" json:"name"`

	// Lifecycle state of the organization tenancy.
	LifecycleState OrganizationTenancyLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Role of the organization tenancy.
	Role OrganizationTenancyRoleEnum `mandatory:"false" json:"role,omitempty"`

	// Date and time when the tenancy joined the organization.
	TimeJoined *common.SDKTime `mandatory:"false" json:"timeJoined"`

	// Date and time when the tenancy left the organization.
	TimeLeft *common.SDKTime `mandatory:"false" json:"timeLeft"`

	// Parameter to indicate the tenancy is approved for transfer to another organization.
	IsApprovedForTransfer *bool `mandatory:"false" json:"isApprovedForTransfer"`
}

func (m OrganizationTenancy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OrganizationTenancy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGovernanceStatusEnum(string(m.GovernanceStatus)); !ok && m.GovernanceStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GovernanceStatus: %s. Supported values are: %s.", m.GovernanceStatus, strings.Join(GetGovernanceStatusEnumStringValues(), ",")))
	}

	if _, ok := GetMappingOrganizationTenancyLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOrganizationTenancyLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOrganizationTenancyRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetOrganizationTenancyRoleEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
