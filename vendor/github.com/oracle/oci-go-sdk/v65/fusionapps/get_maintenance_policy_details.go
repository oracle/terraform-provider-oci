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

// GetMaintenancePolicyDetails The policy that specifies the maintenance and upgrade preferences for an environment. For more information about the options, see Understanding Environment Maintenance (https://docs.cloud.oracle.com/iaas/Content/fusion-applications/plan-environment-family.htm#about-env-maintenance).
type GetMaintenancePolicyDetails struct {
	QuarterlyUpgradeBeginTimes *QuarterlyUpgradeBeginTimes `mandatory:"false" json:"quarterlyUpgradeBeginTimes"`

	// Whether the Fusion environment will be updated monthly or updated on the quarterly cycle. This setting overrides the monthly patching setting of its Fusion environment family.
	MonthlyPatchingOverride MaintenancePolicyMonthlyPatchingOverrideEnum `mandatory:"false" json:"monthlyPatchingOverride,omitempty"`

	// User choice to upgrade both production and non-production environments at the same time. Overrides the Fusion environment family setting.
	EnvironmentMaintenanceOverride MaintenancePolicyEnvironmentMaintenanceOverrideEnum `mandatory:"false" json:"environmentMaintenanceOverride,omitempty"`
}

func (m GetMaintenancePolicyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GetMaintenancePolicyDetails) ValidateEnumValue() (bool, error) {
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
