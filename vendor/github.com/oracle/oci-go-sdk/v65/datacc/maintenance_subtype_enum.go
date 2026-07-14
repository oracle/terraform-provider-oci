// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Infrastructure Cloud@Customer Service API
//
// API for Database Infrastructure Cloud@Customer Service. Use this API to manage Database Infrastructure VM clusters, Application VMs, and related resources.
//

package datacc

import (
	"strings"
)

// MaintenanceSubtypeEnumEnum Enum with underlying type: string
type MaintenanceSubtypeEnumEnum string

// Set of constants representing the allowable values for MaintenanceSubtypeEnumEnum
const (
	MaintenanceSubtypeEnumYearly            MaintenanceSubtypeEnumEnum = "YEARLY"
	MaintenanceSubtypeEnumHalfyearly        MaintenanceSubtypeEnumEnum = "HALFYEARLY"
	MaintenanceSubtypeEnumQuarterly         MaintenanceSubtypeEnumEnum = "QUARTERLY"
	MaintenanceSubtypeEnumMonthly           MaintenanceSubtypeEnumEnum = "MONTHLY"
	MaintenanceSubtypeEnumDaily             MaintenanceSubtypeEnumEnum = "DAILY"
	MaintenanceSubtypeEnumHardware          MaintenanceSubtypeEnumEnum = "HARDWARE"
	MaintenanceSubtypeEnumCritical          MaintenanceSubtypeEnumEnum = "CRITICAL"
	MaintenanceSubtypeEnumInfraUpdate       MaintenanceSubtypeEnumEnum = "INFRA_UPDATE"
	MaintenanceSubtypeEnumCpsServicesUpdate MaintenanceSubtypeEnumEnum = "CPS_SERVICES_UPDATE"
	MaintenanceSubtypeEnumCpsVmUpdate       MaintenanceSubtypeEnumEnum = "CPS_VM_UPDATE"
	MaintenanceSubtypeEnumSecurityMonthly   MaintenanceSubtypeEnumEnum = "SECURITY_MONTHLY"
)

var mappingMaintenanceSubtypeEnumEnum = map[string]MaintenanceSubtypeEnumEnum{
	"YEARLY":              MaintenanceSubtypeEnumYearly,
	"HALFYEARLY":          MaintenanceSubtypeEnumHalfyearly,
	"QUARTERLY":           MaintenanceSubtypeEnumQuarterly,
	"MONTHLY":             MaintenanceSubtypeEnumMonthly,
	"DAILY":               MaintenanceSubtypeEnumDaily,
	"HARDWARE":            MaintenanceSubtypeEnumHardware,
	"CRITICAL":            MaintenanceSubtypeEnumCritical,
	"INFRA_UPDATE":        MaintenanceSubtypeEnumInfraUpdate,
	"CPS_SERVICES_UPDATE": MaintenanceSubtypeEnumCpsServicesUpdate,
	"CPS_VM_UPDATE":       MaintenanceSubtypeEnumCpsVmUpdate,
	"SECURITY_MONTHLY":    MaintenanceSubtypeEnumSecurityMonthly,
}

var mappingMaintenanceSubtypeEnumEnumLowerCase = map[string]MaintenanceSubtypeEnumEnum{
	"yearly":              MaintenanceSubtypeEnumYearly,
	"halfyearly":          MaintenanceSubtypeEnumHalfyearly,
	"quarterly":           MaintenanceSubtypeEnumQuarterly,
	"monthly":             MaintenanceSubtypeEnumMonthly,
	"daily":               MaintenanceSubtypeEnumDaily,
	"hardware":            MaintenanceSubtypeEnumHardware,
	"critical":            MaintenanceSubtypeEnumCritical,
	"infra_update":        MaintenanceSubtypeEnumInfraUpdate,
	"cps_services_update": MaintenanceSubtypeEnumCpsServicesUpdate,
	"cps_vm_update":       MaintenanceSubtypeEnumCpsVmUpdate,
	"security_monthly":    MaintenanceSubtypeEnumSecurityMonthly,
}

// GetMaintenanceSubtypeEnumEnumValues Enumerates the set of values for MaintenanceSubtypeEnumEnum
func GetMaintenanceSubtypeEnumEnumValues() []MaintenanceSubtypeEnumEnum {
	values := make([]MaintenanceSubtypeEnumEnum, 0)
	for _, v := range mappingMaintenanceSubtypeEnumEnum {
		values = append(values, v)
	}
	return values
}

// GetMaintenanceSubtypeEnumEnumStringValues Enumerates the set of values in String for MaintenanceSubtypeEnumEnum
func GetMaintenanceSubtypeEnumEnumStringValues() []string {
	return []string{
		"YEARLY",
		"HALFYEARLY",
		"QUARTERLY",
		"MONTHLY",
		"DAILY",
		"HARDWARE",
		"CRITICAL",
		"INFRA_UPDATE",
		"CPS_SERVICES_UPDATE",
		"CPS_VM_UPDATE",
		"SECURITY_MONTHLY",
	}
}

// GetMappingMaintenanceSubtypeEnumEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaintenanceSubtypeEnumEnum(val string) (MaintenanceSubtypeEnumEnum, bool) {
	enum, ok := mappingMaintenanceSubtypeEnumEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
