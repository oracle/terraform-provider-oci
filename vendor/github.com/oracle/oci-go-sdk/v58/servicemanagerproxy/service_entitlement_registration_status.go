// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Manager Proxy API
//
// Use the Service Manager Proxy API to obtain information about SaaS environments provisioned by Service Manager.
// You can get information such as service types and service environment URLs.
//

package servicemanagerproxy

import (
	"strings"
)

// ServiceEntitlementRegistrationStatusEnum Enum with underlying type: string
type ServiceEntitlementRegistrationStatusEnum string

// Set of constants representing the allowable values for ServiceEntitlementRegistrationStatusEnum
const (
	ServiceEntitlementRegistrationStatusInitialized           ServiceEntitlementRegistrationStatusEnum = "INITIALIZED"
	ServiceEntitlementRegistrationStatusBeginActivation       ServiceEntitlementRegistrationStatusEnum = "BEGIN_ACTIVATION"
	ServiceEntitlementRegistrationStatusActive                ServiceEntitlementRegistrationStatusEnum = "ACTIVE"
	ServiceEntitlementRegistrationStatusBeginSoftTermination  ServiceEntitlementRegistrationStatusEnum = "BEGIN_SOFT_TERMINATION"
	ServiceEntitlementRegistrationStatusSoftTerminated        ServiceEntitlementRegistrationStatusEnum = "SOFT_TERMINATED"
	ServiceEntitlementRegistrationStatusBeginTermination      ServiceEntitlementRegistrationStatusEnum = "BEGIN_TERMINATION"
	ServiceEntitlementRegistrationStatusCanceled              ServiceEntitlementRegistrationStatusEnum = "CANCELED"
	ServiceEntitlementRegistrationStatusTerminated            ServiceEntitlementRegistrationStatusEnum = "TERMINATED"
	ServiceEntitlementRegistrationStatusBeginDisabling        ServiceEntitlementRegistrationStatusEnum = "BEGIN_DISABLING"
	ServiceEntitlementRegistrationStatusBeginEnabling         ServiceEntitlementRegistrationStatusEnum = "BEGIN_ENABLING"
	ServiceEntitlementRegistrationStatusBeginMigration        ServiceEntitlementRegistrationStatusEnum = "BEGIN_MIGRATION"
	ServiceEntitlementRegistrationStatusDisabled              ServiceEntitlementRegistrationStatusEnum = "DISABLED"
	ServiceEntitlementRegistrationStatusBeginSuspension       ServiceEntitlementRegistrationStatusEnum = "BEGIN_SUSPENSION"
	ServiceEntitlementRegistrationStatusBeginResumption       ServiceEntitlementRegistrationStatusEnum = "BEGIN_RESUMPTION"
	ServiceEntitlementRegistrationStatusSuspended             ServiceEntitlementRegistrationStatusEnum = "SUSPENDED"
	ServiceEntitlementRegistrationStatusBeginLockRelocation   ServiceEntitlementRegistrationStatusEnum = "BEGIN_LOCK_RELOCATION"
	ServiceEntitlementRegistrationStatusLockedRelocation      ServiceEntitlementRegistrationStatusEnum = "LOCKED_RELOCATION"
	ServiceEntitlementRegistrationStatusBeginRelocation       ServiceEntitlementRegistrationStatusEnum = "BEGIN_RELOCATION"
	ServiceEntitlementRegistrationStatusRelocated             ServiceEntitlementRegistrationStatusEnum = "RELOCATED"
	ServiceEntitlementRegistrationStatusBeginUnlockRelocation ServiceEntitlementRegistrationStatusEnum = "BEGIN_UNLOCK_RELOCATION"
	ServiceEntitlementRegistrationStatusUnlockedRelocation    ServiceEntitlementRegistrationStatusEnum = "UNLOCKED_RELOCATION"
	ServiceEntitlementRegistrationStatusFailedLockRelocation  ServiceEntitlementRegistrationStatusEnum = "FAILED_LOCK_RELOCATION"
	ServiceEntitlementRegistrationStatusFailedActivation      ServiceEntitlementRegistrationStatusEnum = "FAILED_ACTIVATION"
	ServiceEntitlementRegistrationStatusFailedMigration       ServiceEntitlementRegistrationStatusEnum = "FAILED_MIGRATION"
	ServiceEntitlementRegistrationStatusAccessDisabled        ServiceEntitlementRegistrationStatusEnum = "ACCESS_DISABLED"
	ServiceEntitlementRegistrationStatusBeginDisablingAccess  ServiceEntitlementRegistrationStatusEnum = "BEGIN_DISABLING_ACCESS"
	ServiceEntitlementRegistrationStatusBeginEnablingAccess   ServiceEntitlementRegistrationStatusEnum = "BEGIN_ENABLING_ACCESS"
	ServiceEntitlementRegistrationStatusTraUnknown            ServiceEntitlementRegistrationStatusEnum = "TRA_UNKNOWN"
)

var mappingServiceEntitlementRegistrationStatusEnum = map[string]ServiceEntitlementRegistrationStatusEnum{
	"INITIALIZED":             ServiceEntitlementRegistrationStatusInitialized,
	"BEGIN_ACTIVATION":        ServiceEntitlementRegistrationStatusBeginActivation,
	"ACTIVE":                  ServiceEntitlementRegistrationStatusActive,
	"BEGIN_SOFT_TERMINATION":  ServiceEntitlementRegistrationStatusBeginSoftTermination,
	"SOFT_TERMINATED":         ServiceEntitlementRegistrationStatusSoftTerminated,
	"BEGIN_TERMINATION":       ServiceEntitlementRegistrationStatusBeginTermination,
	"CANCELED":                ServiceEntitlementRegistrationStatusCanceled,
	"TERMINATED":              ServiceEntitlementRegistrationStatusTerminated,
	"BEGIN_DISABLING":         ServiceEntitlementRegistrationStatusBeginDisabling,
	"BEGIN_ENABLING":          ServiceEntitlementRegistrationStatusBeginEnabling,
	"BEGIN_MIGRATION":         ServiceEntitlementRegistrationStatusBeginMigration,
	"DISABLED":                ServiceEntitlementRegistrationStatusDisabled,
	"BEGIN_SUSPENSION":        ServiceEntitlementRegistrationStatusBeginSuspension,
	"BEGIN_RESUMPTION":        ServiceEntitlementRegistrationStatusBeginResumption,
	"SUSPENDED":               ServiceEntitlementRegistrationStatusSuspended,
	"BEGIN_LOCK_RELOCATION":   ServiceEntitlementRegistrationStatusBeginLockRelocation,
	"LOCKED_RELOCATION":       ServiceEntitlementRegistrationStatusLockedRelocation,
	"BEGIN_RELOCATION":        ServiceEntitlementRegistrationStatusBeginRelocation,
	"RELOCATED":               ServiceEntitlementRegistrationStatusRelocated,
	"BEGIN_UNLOCK_RELOCATION": ServiceEntitlementRegistrationStatusBeginUnlockRelocation,
	"UNLOCKED_RELOCATION":     ServiceEntitlementRegistrationStatusUnlockedRelocation,
	"FAILED_LOCK_RELOCATION":  ServiceEntitlementRegistrationStatusFailedLockRelocation,
	"FAILED_ACTIVATION":       ServiceEntitlementRegistrationStatusFailedActivation,
	"FAILED_MIGRATION":        ServiceEntitlementRegistrationStatusFailedMigration,
	"ACCESS_DISABLED":         ServiceEntitlementRegistrationStatusAccessDisabled,
	"BEGIN_DISABLING_ACCESS":  ServiceEntitlementRegistrationStatusBeginDisablingAccess,
	"BEGIN_ENABLING_ACCESS":   ServiceEntitlementRegistrationStatusBeginEnablingAccess,
	"TRA_UNKNOWN":             ServiceEntitlementRegistrationStatusTraUnknown,
}

// GetServiceEntitlementRegistrationStatusEnumValues Enumerates the set of values for ServiceEntitlementRegistrationStatusEnum
func GetServiceEntitlementRegistrationStatusEnumValues() []ServiceEntitlementRegistrationStatusEnum {
	values := make([]ServiceEntitlementRegistrationStatusEnum, 0)
	for _, v := range mappingServiceEntitlementRegistrationStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetServiceEntitlementRegistrationStatusEnumStringValues Enumerates the set of values in String for ServiceEntitlementRegistrationStatusEnum
func GetServiceEntitlementRegistrationStatusEnumStringValues() []string {
	return []string{
		"INITIALIZED",
		"BEGIN_ACTIVATION",
		"ACTIVE",
		"BEGIN_SOFT_TERMINATION",
		"SOFT_TERMINATED",
		"BEGIN_TERMINATION",
		"CANCELED",
		"TERMINATED",
		"BEGIN_DISABLING",
		"BEGIN_ENABLING",
		"BEGIN_MIGRATION",
		"DISABLED",
		"BEGIN_SUSPENSION",
		"BEGIN_RESUMPTION",
		"SUSPENDED",
		"BEGIN_LOCK_RELOCATION",
		"LOCKED_RELOCATION",
		"BEGIN_RELOCATION",
		"RELOCATED",
		"BEGIN_UNLOCK_RELOCATION",
		"UNLOCKED_RELOCATION",
		"FAILED_LOCK_RELOCATION",
		"FAILED_ACTIVATION",
		"FAILED_MIGRATION",
		"ACCESS_DISABLED",
		"BEGIN_DISABLING_ACCESS",
		"BEGIN_ENABLING_ACCESS",
		"TRA_UNKNOWN",
	}
}

// GetMappingServiceEntitlementRegistrationStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingServiceEntitlementRegistrationStatusEnum(val string) (ServiceEntitlementRegistrationStatusEnum, bool) {
	mappingServiceEntitlementRegistrationStatusEnumIgnoreCase := make(map[string]ServiceEntitlementRegistrationStatusEnum)
	for k, v := range mappingServiceEntitlementRegistrationStatusEnum {
		mappingServiceEntitlementRegistrationStatusEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingServiceEntitlementRegistrationStatusEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
