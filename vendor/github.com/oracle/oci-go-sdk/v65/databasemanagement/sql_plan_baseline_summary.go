// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SqlPlanBaselineSummary The summary of a SQL plan baseline.
type SqlPlanBaselineSummary struct {

	// The unique plan identifier.
	PlanName *string `mandatory:"true" json:"planName"`

	// The unique SQL identifier.
	SqlHandle *string `mandatory:"true" json:"sqlHandle"`

	// The SQL text (truncated to the first 50 characters).
	SqlText *string `mandatory:"true" json:"sqlText"`

	// The date and time when the plan baseline was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The origin of the SQL plan baseline.
	Origin SqlPlanBaselineOriginEnum `mandatory:"false" json:"origin,omitempty"`

	// The date and time when the plan baseline was last modified.
	TimeLastModified *common.SDKTime `mandatory:"false" json:"timeLastModified"`

	// The date and time when the plan baseline was last executed.
	// **Note:** For performance reasons, database does not update this value
	// immediately after each execution of the plan baseline. Therefore, the plan
	// baseline may have been executed more recently than this value indicates.
	TimeLastExecuted *common.SDKTime `mandatory:"false" json:"timeLastExecuted"`

	// Indicates whether the plan baseline is enabled (`YES`) or disabled (`NO`).
	Enabled *string `mandatory:"false" json:"enabled"`

	// Indicates whether the plan baseline is accepted (`YES`) or not (`NO`).
	Accepted *string `mandatory:"false" json:"accepted"`

	// Indicates whether the plan baseline is fixed (`YES`) or not (`NO`).
	Fixed *string `mandatory:"false" json:"fixed"`

	// Indicates whether the optimizer was able to reproduce the plan (`YES`) or not (`NO`).
	// The value is set to `YES` when a plan is initially added to the plan baseline.
	Reproduced *string `mandatory:"false" json:"reproduced"`

	// Indicates whether the plan baseline is auto-purged (`YES`) or not (`NO`).
	AutoPurge *string `mandatory:"false" json:"autoPurge"`

	// Indicates whether a plan that is automatically captured by SQL plan management is marked adaptive or not.
	// When a new adaptive plan is found for a SQL statement that has an existing SQL plan baseline, that new plan
	// will be added to the SQL plan baseline as an unaccepted plan, and the `ADAPTIVE` property will be marked `YES`.
	// When this new plan is verified (either manually or via the auto evolve task), the plan will be test executed
	// and the final plan determined at execution will become an accepted plan if its performance is better than
	// the existing plan baseline. At this point, the value of the `ADAPTIVE` property is set to `NO` since the plan
	// is no longer adaptive, but resolved.
	Adaptive *string `mandatory:"false" json:"adaptive"`
}

func (m SqlPlanBaselineSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlPlanBaselineSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSqlPlanBaselineOriginEnum(string(m.Origin)); !ok && m.Origin != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Origin: %s. Supported values are: %s.", m.Origin, strings.Join(GetSqlPlanBaselineOriginEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
