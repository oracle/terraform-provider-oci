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

// ClassicSubscriptionManagedByEnum Enum with underlying type: string
type ClassicSubscriptionManagedByEnum string

// Set of constants representing the allowable values for ClassicSubscriptionManagedByEnum
const (
	ClassicSubscriptionManagedByAppsmanager         ClassicSubscriptionManagedByEnum = "APPSMANAGER"
	ClassicSubscriptionManagedByServicemanagerproxy ClassicSubscriptionManagedByEnum = "SERVICEMANAGERPROXY"
	ClassicSubscriptionManagedByFusionapps          ClassicSubscriptionManagedByEnum = "FUSIONAPPS"
	ClassicSubscriptionManagedByMyservices          ClassicSubscriptionManagedByEnum = "MYSERVICES"
)

var mappingClassicSubscriptionManagedByEnum = map[string]ClassicSubscriptionManagedByEnum{
	"APPSMANAGER":         ClassicSubscriptionManagedByAppsmanager,
	"SERVICEMANAGERPROXY": ClassicSubscriptionManagedByServicemanagerproxy,
	"FUSIONAPPS":          ClassicSubscriptionManagedByFusionapps,
	"MYSERVICES":          ClassicSubscriptionManagedByMyservices,
}

var mappingClassicSubscriptionManagedByEnumLowerCase = map[string]ClassicSubscriptionManagedByEnum{
	"appsmanager":         ClassicSubscriptionManagedByAppsmanager,
	"servicemanagerproxy": ClassicSubscriptionManagedByServicemanagerproxy,
	"fusionapps":          ClassicSubscriptionManagedByFusionapps,
	"myservices":          ClassicSubscriptionManagedByMyservices,
}

// GetClassicSubscriptionManagedByEnumValues Enumerates the set of values for ClassicSubscriptionManagedByEnum
func GetClassicSubscriptionManagedByEnumValues() []ClassicSubscriptionManagedByEnum {
	values := make([]ClassicSubscriptionManagedByEnum, 0)
	for _, v := range mappingClassicSubscriptionManagedByEnum {
		values = append(values, v)
	}
	return values
}

// GetClassicSubscriptionManagedByEnumStringValues Enumerates the set of values in String for ClassicSubscriptionManagedByEnum
func GetClassicSubscriptionManagedByEnumStringValues() []string {
	return []string{
		"APPSMANAGER",
		"SERVICEMANAGERPROXY",
		"FUSIONAPPS",
		"MYSERVICES",
	}
}

// GetMappingClassicSubscriptionManagedByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingClassicSubscriptionManagedByEnum(val string) (ClassicSubscriptionManagedByEnum, bool) {
	enum, ok := mappingClassicSubscriptionManagedByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
