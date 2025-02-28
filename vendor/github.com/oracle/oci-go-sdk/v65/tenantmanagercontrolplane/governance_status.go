// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// GovernanceStatusEnum Enum with underlying type: string
type GovernanceStatusEnum string

// Set of constants representing the allowable values for GovernanceStatusEnum
const (
	GovernanceStatusOptedIn  GovernanceStatusEnum = "OPTED_IN"
	GovernanceStatusOptedOut GovernanceStatusEnum = "OPTED_OUT"
)

var mappingGovernanceStatusEnum = map[string]GovernanceStatusEnum{
	"OPTED_IN":  GovernanceStatusOptedIn,
	"OPTED_OUT": GovernanceStatusOptedOut,
}

var mappingGovernanceStatusEnumLowerCase = map[string]GovernanceStatusEnum{
	"opted_in":  GovernanceStatusOptedIn,
	"opted_out": GovernanceStatusOptedOut,
}

// GetGovernanceStatusEnumValues Enumerates the set of values for GovernanceStatusEnum
func GetGovernanceStatusEnumValues() []GovernanceStatusEnum {
	values := make([]GovernanceStatusEnum, 0)
	for _, v := range mappingGovernanceStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetGovernanceStatusEnumStringValues Enumerates the set of values in String for GovernanceStatusEnum
func GetGovernanceStatusEnumStringValues() []string {
	return []string{
		"OPTED_IN",
		"OPTED_OUT",
	}
}

// GetMappingGovernanceStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGovernanceStatusEnum(val string) (GovernanceStatusEnum, bool) {
	enum, ok := mappingGovernanceStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
