// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// SELF Service API
//
// Use the SELF Service API to manage Subscriptions in Oracle Cloud Infrastructure Marketplace. For more information, see Overview of Marketplace (https://docs.oracle.com/iaas/Content/Marketplace/Concepts/marketoverview.htm)
//

package self

import (
	"strings"
)

// LifecycleDetailsEnumEnum Enum with underlying type: string
type LifecycleDetailsEnumEnum string

// Set of constants representing the allowable values for LifecycleDetailsEnumEnum
const (
	LifecycleDetailsEnumCreated               LifecycleDetailsEnumEnum = "CREATED"
	LifecycleDetailsEnumPendingActivation     LifecycleDetailsEnumEnum = "PENDING_ACTIVATION"
	LifecycleDetailsEnumProvisioningStarted   LifecycleDetailsEnumEnum = "PROVISIONING_STARTED"
	LifecycleDetailsEnumProvisioningCompleted LifecycleDetailsEnumEnum = "PROVISIONING_COMPLETED"
	LifecycleDetailsEnumProvisioningFailed    LifecycleDetailsEnumEnum = "PROVISIONING_FAILED"
	LifecycleDetailsEnumActive                LifecycleDetailsEnumEnum = "ACTIVE"
	LifecycleDetailsEnumExpired               LifecycleDetailsEnumEnum = "EXPIRED"
	LifecycleDetailsEnumTerminated            LifecycleDetailsEnumEnum = "TERMINATED"
	LifecycleDetailsEnumFailed                LifecycleDetailsEnumEnum = "FAILED"
	LifecycleDetailsEnumDeleting              LifecycleDetailsEnumEnum = "DELETING"
	LifecycleDetailsEnumUpdating              LifecycleDetailsEnumEnum = "UPDATING"
	LifecycleDetailsEnumDeleted               LifecycleDetailsEnumEnum = "DELETED"
)

var mappingLifecycleDetailsEnumEnum = map[string]LifecycleDetailsEnumEnum{
	"CREATED":                LifecycleDetailsEnumCreated,
	"PENDING_ACTIVATION":     LifecycleDetailsEnumPendingActivation,
	"PROVISIONING_STARTED":   LifecycleDetailsEnumProvisioningStarted,
	"PROVISIONING_COMPLETED": LifecycleDetailsEnumProvisioningCompleted,
	"PROVISIONING_FAILED":    LifecycleDetailsEnumProvisioningFailed,
	"ACTIVE":                 LifecycleDetailsEnumActive,
	"EXPIRED":                LifecycleDetailsEnumExpired,
	"TERMINATED":             LifecycleDetailsEnumTerminated,
	"FAILED":                 LifecycleDetailsEnumFailed,
	"DELETING":               LifecycleDetailsEnumDeleting,
	"UPDATING":               LifecycleDetailsEnumUpdating,
	"DELETED":                LifecycleDetailsEnumDeleted,
}

var mappingLifecycleDetailsEnumEnumLowerCase = map[string]LifecycleDetailsEnumEnum{
	"created":                LifecycleDetailsEnumCreated,
	"pending_activation":     LifecycleDetailsEnumPendingActivation,
	"provisioning_started":   LifecycleDetailsEnumProvisioningStarted,
	"provisioning_completed": LifecycleDetailsEnumProvisioningCompleted,
	"provisioning_failed":    LifecycleDetailsEnumProvisioningFailed,
	"active":                 LifecycleDetailsEnumActive,
	"expired":                LifecycleDetailsEnumExpired,
	"terminated":             LifecycleDetailsEnumTerminated,
	"failed":                 LifecycleDetailsEnumFailed,
	"deleting":               LifecycleDetailsEnumDeleting,
	"updating":               LifecycleDetailsEnumUpdating,
	"deleted":                LifecycleDetailsEnumDeleted,
}

// GetLifecycleDetailsEnumEnumValues Enumerates the set of values for LifecycleDetailsEnumEnum
func GetLifecycleDetailsEnumEnumValues() []LifecycleDetailsEnumEnum {
	values := make([]LifecycleDetailsEnumEnum, 0)
	for _, v := range mappingLifecycleDetailsEnumEnum {
		values = append(values, v)
	}
	return values
}

// GetLifecycleDetailsEnumEnumStringValues Enumerates the set of values in String for LifecycleDetailsEnumEnum
func GetLifecycleDetailsEnumEnumStringValues() []string {
	return []string{
		"CREATED",
		"PENDING_ACTIVATION",
		"PROVISIONING_STARTED",
		"PROVISIONING_COMPLETED",
		"PROVISIONING_FAILED",
		"ACTIVE",
		"EXPIRED",
		"TERMINATED",
		"FAILED",
		"DELETING",
		"UPDATING",
		"DELETED",
	}
}

// GetMappingLifecycleDetailsEnumEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLifecycleDetailsEnumEnum(val string) (LifecycleDetailsEnumEnum, bool) {
	enum, ok := mappingLifecycleDetailsEnumEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
