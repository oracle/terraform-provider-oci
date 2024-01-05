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

// FamilyMaintenancePolicy The policy that specifies the maintenance and upgrade preferences for an environment. For more information about the options, see Understanding Environment Maintenance (https://docs.cloud.oracle.com/iaas/Content/fusion-applications/plan-environment-family.htm#about-env-maintenance).
type FamilyMaintenancePolicy struct {

	// The quarterly maintenance month group schedule of the Fusion environment family.
	QuarterlyUpgradeBeginTimes *string `mandatory:"false" json:"quarterlyUpgradeBeginTimes"`

	// When True, monthly patching is enabled for the environment family.
	IsMonthlyPatchingEnabled *bool `mandatory:"false" json:"isMonthlyPatchingEnabled"`

	// Option to upgrade both production and non-production environments at the same time. When set to PROD both types of environnments are upgraded on the production schedule. When set to NON_PROD both types of environments are upgraded on the non-production schedule.
	ConcurrentMaintenance FamilyMaintenancePolicyConcurrentMaintenanceEnum `mandatory:"false" json:"concurrentMaintenance,omitempty"`
}

func (m FamilyMaintenancePolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FamilyMaintenancePolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingFamilyMaintenancePolicyConcurrentMaintenanceEnum(string(m.ConcurrentMaintenance)); !ok && m.ConcurrentMaintenance != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConcurrentMaintenance: %s. Supported values are: %s.", m.ConcurrentMaintenance, strings.Join(GetFamilyMaintenancePolicyConcurrentMaintenanceEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FamilyMaintenancePolicyConcurrentMaintenanceEnum Enum with underlying type: string
type FamilyMaintenancePolicyConcurrentMaintenanceEnum string

// Set of constants representing the allowable values for FamilyMaintenancePolicyConcurrentMaintenanceEnum
const (
	FamilyMaintenancePolicyConcurrentMaintenanceProd     FamilyMaintenancePolicyConcurrentMaintenanceEnum = "PROD"
	FamilyMaintenancePolicyConcurrentMaintenanceNonProd  FamilyMaintenancePolicyConcurrentMaintenanceEnum = "NON_PROD"
	FamilyMaintenancePolicyConcurrentMaintenanceDisabled FamilyMaintenancePolicyConcurrentMaintenanceEnum = "DISABLED"
)

var mappingFamilyMaintenancePolicyConcurrentMaintenanceEnum = map[string]FamilyMaintenancePolicyConcurrentMaintenanceEnum{
	"PROD":     FamilyMaintenancePolicyConcurrentMaintenanceProd,
	"NON_PROD": FamilyMaintenancePolicyConcurrentMaintenanceNonProd,
	"DISABLED": FamilyMaintenancePolicyConcurrentMaintenanceDisabled,
}

var mappingFamilyMaintenancePolicyConcurrentMaintenanceEnumLowerCase = map[string]FamilyMaintenancePolicyConcurrentMaintenanceEnum{
	"prod":     FamilyMaintenancePolicyConcurrentMaintenanceProd,
	"non_prod": FamilyMaintenancePolicyConcurrentMaintenanceNonProd,
	"disabled": FamilyMaintenancePolicyConcurrentMaintenanceDisabled,
}

// GetFamilyMaintenancePolicyConcurrentMaintenanceEnumValues Enumerates the set of values for FamilyMaintenancePolicyConcurrentMaintenanceEnum
func GetFamilyMaintenancePolicyConcurrentMaintenanceEnumValues() []FamilyMaintenancePolicyConcurrentMaintenanceEnum {
	values := make([]FamilyMaintenancePolicyConcurrentMaintenanceEnum, 0)
	for _, v := range mappingFamilyMaintenancePolicyConcurrentMaintenanceEnum {
		values = append(values, v)
	}
	return values
}

// GetFamilyMaintenancePolicyConcurrentMaintenanceEnumStringValues Enumerates the set of values in String for FamilyMaintenancePolicyConcurrentMaintenanceEnum
func GetFamilyMaintenancePolicyConcurrentMaintenanceEnumStringValues() []string {
	return []string{
		"PROD",
		"NON_PROD",
		"DISABLED",
	}
}

// GetMappingFamilyMaintenancePolicyConcurrentMaintenanceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFamilyMaintenancePolicyConcurrentMaintenanceEnum(val string) (FamilyMaintenancePolicyConcurrentMaintenanceEnum, bool) {
	enum, ok := mappingFamilyMaintenancePolicyConcurrentMaintenanceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
