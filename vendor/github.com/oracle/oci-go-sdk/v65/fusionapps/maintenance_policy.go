// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fusion Applications Environment Management API
//
// Use the Fusion Applications Environment Management API to manage the environments where your Fusion Applications run. For more information, see the Fusion Applications Environment Management documentation (https://docs.cloud.oracle.com/iaas/Content/fusion-applications/home.htm).
//

package fusionapps

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MaintenancePolicy The policy that specifies the maintenance and upgrade preferences for an environment. For more information about the options, see Understanding Environment Maintenance (https://docs.cloud.oracle.com/iaas/Content/fusion-applications/plan-environment-family.htm#about-env-maintenance).
type MaintenancePolicy struct {

	// When "ENABLED", the Fusion environment is patched monthly. When "DISABLED", the Fusion environment is not patched monthly. This setting overrides the environment family setting. When not set, the environment follows the environment family policy.
	MonthlyPatchingOverride MaintenancePolicyMonthlyPatchingOverrideEnum `mandatory:"false" json:"monthlyPatchingOverride,omitempty"`

	// User choice to upgrade both test and prod pods at the same time. Overrides fusion environment families'.
	EnvironmentMaintenanceOverride MaintenancePolicyEnvironmentMaintenanceOverrideEnum `mandatory:"false" json:"environmentMaintenanceOverride,omitempty"`
}

func (m MaintenancePolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MaintenancePolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingMaintenancePolicyMonthlyPatchingOverrideEnum(string(m.MonthlyPatchingOverride)); !ok && m.MonthlyPatchingOverride != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MonthlyPatchingOverride: %s. Supported values are: %s.", m.MonthlyPatchingOverride, strings.Join(GetMaintenancePolicyMonthlyPatchingOverrideEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMaintenancePolicyEnvironmentMaintenanceOverrideEnum(string(m.EnvironmentMaintenanceOverride)); !ok && m.EnvironmentMaintenanceOverride != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EnvironmentMaintenanceOverride: %s. Supported values are: %s.", m.EnvironmentMaintenanceOverride, strings.Join(GetMaintenancePolicyEnvironmentMaintenanceOverrideEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MaintenancePolicyMonthlyPatchingOverrideEnum Enum with underlying type: string
type MaintenancePolicyMonthlyPatchingOverrideEnum string

// Set of constants representing the allowable values for MaintenancePolicyMonthlyPatchingOverrideEnum
const (
	MaintenancePolicyMonthlyPatchingOverrideEnabled  MaintenancePolicyMonthlyPatchingOverrideEnum = "ENABLED"
	MaintenancePolicyMonthlyPatchingOverrideDisabled MaintenancePolicyMonthlyPatchingOverrideEnum = "DISABLED"
	MaintenancePolicyMonthlyPatchingOverrideNone     MaintenancePolicyMonthlyPatchingOverrideEnum = "NONE"
)

var mappingMaintenancePolicyMonthlyPatchingOverrideEnum = map[string]MaintenancePolicyMonthlyPatchingOverrideEnum{
	"ENABLED":  MaintenancePolicyMonthlyPatchingOverrideEnabled,
	"DISABLED": MaintenancePolicyMonthlyPatchingOverrideDisabled,
	"NONE":     MaintenancePolicyMonthlyPatchingOverrideNone,
}

var mappingMaintenancePolicyMonthlyPatchingOverrideEnumLowerCase = map[string]MaintenancePolicyMonthlyPatchingOverrideEnum{
	"enabled":  MaintenancePolicyMonthlyPatchingOverrideEnabled,
	"disabled": MaintenancePolicyMonthlyPatchingOverrideDisabled,
	"none":     MaintenancePolicyMonthlyPatchingOverrideNone,
}

// GetMaintenancePolicyMonthlyPatchingOverrideEnumValues Enumerates the set of values for MaintenancePolicyMonthlyPatchingOverrideEnum
func GetMaintenancePolicyMonthlyPatchingOverrideEnumValues() []MaintenancePolicyMonthlyPatchingOverrideEnum {
	values := make([]MaintenancePolicyMonthlyPatchingOverrideEnum, 0)
	for _, v := range mappingMaintenancePolicyMonthlyPatchingOverrideEnum {
		values = append(values, v)
	}
	return values
}

// GetMaintenancePolicyMonthlyPatchingOverrideEnumStringValues Enumerates the set of values in String for MaintenancePolicyMonthlyPatchingOverrideEnum
func GetMaintenancePolicyMonthlyPatchingOverrideEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
		"NONE",
	}
}

// GetMappingMaintenancePolicyMonthlyPatchingOverrideEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaintenancePolicyMonthlyPatchingOverrideEnum(val string) (MaintenancePolicyMonthlyPatchingOverrideEnum, bool) {
	enum, ok := mappingMaintenancePolicyMonthlyPatchingOverrideEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// MaintenancePolicyEnvironmentMaintenanceOverrideEnum Enum with underlying type: string
type MaintenancePolicyEnvironmentMaintenanceOverrideEnum string

// Set of constants representing the allowable values for MaintenancePolicyEnvironmentMaintenanceOverrideEnum
const (
	MaintenancePolicyEnvironmentMaintenanceOverrideProd    MaintenancePolicyEnvironmentMaintenanceOverrideEnum = "PROD"
	MaintenancePolicyEnvironmentMaintenanceOverrideNonProd MaintenancePolicyEnvironmentMaintenanceOverrideEnum = "NON_PROD"
	MaintenancePolicyEnvironmentMaintenanceOverrideNone    MaintenancePolicyEnvironmentMaintenanceOverrideEnum = "NONE"
)

var mappingMaintenancePolicyEnvironmentMaintenanceOverrideEnum = map[string]MaintenancePolicyEnvironmentMaintenanceOverrideEnum{
	"PROD":     MaintenancePolicyEnvironmentMaintenanceOverrideProd,
	"NON_PROD": MaintenancePolicyEnvironmentMaintenanceOverrideNonProd,
	"NONE":     MaintenancePolicyEnvironmentMaintenanceOverrideNone,
}

var mappingMaintenancePolicyEnvironmentMaintenanceOverrideEnumLowerCase = map[string]MaintenancePolicyEnvironmentMaintenanceOverrideEnum{
	"prod":     MaintenancePolicyEnvironmentMaintenanceOverrideProd,
	"non_prod": MaintenancePolicyEnvironmentMaintenanceOverrideNonProd,
	"none":     MaintenancePolicyEnvironmentMaintenanceOverrideNone,
}

// GetMaintenancePolicyEnvironmentMaintenanceOverrideEnumValues Enumerates the set of values for MaintenancePolicyEnvironmentMaintenanceOverrideEnum
func GetMaintenancePolicyEnvironmentMaintenanceOverrideEnumValues() []MaintenancePolicyEnvironmentMaintenanceOverrideEnum {
	values := make([]MaintenancePolicyEnvironmentMaintenanceOverrideEnum, 0)
	for _, v := range mappingMaintenancePolicyEnvironmentMaintenanceOverrideEnum {
		values = append(values, v)
	}
	return values
}

// GetMaintenancePolicyEnvironmentMaintenanceOverrideEnumStringValues Enumerates the set of values in String for MaintenancePolicyEnvironmentMaintenanceOverrideEnum
func GetMaintenancePolicyEnvironmentMaintenanceOverrideEnumStringValues() []string {
	return []string{
		"PROD",
		"NON_PROD",
		"NONE",
	}
}

// GetMappingMaintenancePolicyEnvironmentMaintenanceOverrideEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaintenancePolicyEnvironmentMaintenanceOverrideEnum(val string) (MaintenancePolicyEnvironmentMaintenanceOverrideEnum, bool) {
	enum, ok := mappingMaintenancePolicyEnvironmentMaintenanceOverrideEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
