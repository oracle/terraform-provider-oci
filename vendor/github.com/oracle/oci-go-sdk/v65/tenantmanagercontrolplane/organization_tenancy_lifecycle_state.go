// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// Use the Organizations API to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and organization resources. For more information, see Organization Management Overview (https://docs.oracle.com/iaas/Content/General/Concepts/organization_management_overview.htm).
//

package tenantmanagercontrolplane

import (
	"strings"
)

// OrganizationTenancyLifecycleStateEnum Enum with underlying type: string
type OrganizationTenancyLifecycleStateEnum string

// Set of constants representing the allowable values for OrganizationTenancyLifecycleStateEnum
const (
	OrganizationTenancyLifecycleStateCreating OrganizationTenancyLifecycleStateEnum = "CREATING"
	OrganizationTenancyLifecycleStateActive   OrganizationTenancyLifecycleStateEnum = "ACTIVE"
	OrganizationTenancyLifecycleStateInactive OrganizationTenancyLifecycleStateEnum = "INACTIVE"
	OrganizationTenancyLifecycleStateDeleted  OrganizationTenancyLifecycleStateEnum = "DELETED"
	OrganizationTenancyLifecycleStateFailed   OrganizationTenancyLifecycleStateEnum = "FAILED"
	OrganizationTenancyLifecycleStateDeleting OrganizationTenancyLifecycleStateEnum = "DELETING"
)

var mappingOrganizationTenancyLifecycleStateEnum = map[string]OrganizationTenancyLifecycleStateEnum{
	"CREATING": OrganizationTenancyLifecycleStateCreating,
	"ACTIVE":   OrganizationTenancyLifecycleStateActive,
	"INACTIVE": OrganizationTenancyLifecycleStateInactive,
	"DELETED":  OrganizationTenancyLifecycleStateDeleted,
	"FAILED":   OrganizationTenancyLifecycleStateFailed,
	"DELETING": OrganizationTenancyLifecycleStateDeleting,
}

var mappingOrganizationTenancyLifecycleStateEnumLowerCase = map[string]OrganizationTenancyLifecycleStateEnum{
	"creating": OrganizationTenancyLifecycleStateCreating,
	"active":   OrganizationTenancyLifecycleStateActive,
	"inactive": OrganizationTenancyLifecycleStateInactive,
	"deleted":  OrganizationTenancyLifecycleStateDeleted,
	"failed":   OrganizationTenancyLifecycleStateFailed,
	"deleting": OrganizationTenancyLifecycleStateDeleting,
}

// GetOrganizationTenancyLifecycleStateEnumValues Enumerates the set of values for OrganizationTenancyLifecycleStateEnum
func GetOrganizationTenancyLifecycleStateEnumValues() []OrganizationTenancyLifecycleStateEnum {
	values := make([]OrganizationTenancyLifecycleStateEnum, 0)
	for _, v := range mappingOrganizationTenancyLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOrganizationTenancyLifecycleStateEnumStringValues Enumerates the set of values in String for OrganizationTenancyLifecycleStateEnum
func GetOrganizationTenancyLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"DELETED",
		"FAILED",
		"DELETING",
	}
}

// GetMappingOrganizationTenancyLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOrganizationTenancyLifecycleStateEnum(val string) (OrganizationTenancyLifecycleStateEnum, bool) {
	enum, ok := mappingOrganizationTenancyLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
