// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SqlPlanBaselineConfiguration The configuration details of SQL plan baselines. The details include:
// - whether automatic initial plan capture is enabled or disabled
// - whether use of SQL plan baselines is enabled or disabled
// - whether Automatic SPM Evolve Advisor task is enabled or disabled
// - whether high-frequency Automatic SPM Evolve Advisor task is enabled or disabled
// - filters for the automatic initial plan capture
// - parameters for the Automatic SPM Evolve Advisor task
// - plan retention and allocated space for the plan baselines
type SqlPlanBaselineConfiguration struct {

	// Indicates whether the automatic capture of SQL plan baselines is enabled (`true`) or not (`false`).
	IsAutomaticInitialPlanCaptureEnabled *bool `mandatory:"true" json:"isAutomaticInitialPlanCaptureEnabled"`

	// Indicates whether the database uses SQL plan baselines (`true`) or not (`false`).
	IsSqlPlanBaselinesUsageEnabled *bool `mandatory:"true" json:"isSqlPlanBaselinesUsageEnabled"`

	// Indicates whether the Automatic SPM Evolve Advisor task is enabled (`true`) or not (`false`).
	IsAutoSpmEvolveTaskEnabled *bool `mandatory:"true" json:"isAutoSpmEvolveTaskEnabled"`

	// Indicates whether the high frequency Automatic SPM Evolve Advisor task is enabled (`true`) or not (`false`).
	IsHighFrequencyAutoSpmEvolveTaskEnabled *bool `mandatory:"true" json:"isHighFrequencyAutoSpmEvolveTaskEnabled"`

	// The number of weeks to retain unused plans before they are purged.
	PlanRetentionWeeks *int `mandatory:"true" json:"planRetentionWeeks"`

	// The maximum percent of `SYSAUX` space that can be used for SQL Management Base.
	SpaceBudgetPercent *float32 `mandatory:"true" json:"spaceBudgetPercent"`

	// The maximum `SYSAUX` space that can be used for SQL Management Base in MB.
	SpaceBudgetMB *float32 `mandatory:"false" json:"spaceBudgetMB"`

	// The space used by SQL Management Base in MB.
	SpaceUsedMB *float32 `mandatory:"false" json:"spaceUsedMB"`

	// The capture filters used in automatic initial plan capture.
	AutoCaptureFilters []AutomaticCaptureFilter `mandatory:"false" json:"autoCaptureFilters"`

	AutoSpmEvolveTaskParameters *SpmEvolveTaskParameters `mandatory:"false" json:"autoSpmEvolveTaskParameters"`
}

func (m SqlPlanBaselineConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlPlanBaselineConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
