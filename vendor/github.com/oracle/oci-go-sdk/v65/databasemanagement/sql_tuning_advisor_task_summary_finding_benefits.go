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

// SqlTuningAdvisorTaskSummaryFindingBenefits The benefits of the findings in the SQL Tuning Advisor summary report.
type SqlTuningAdvisorTaskSummaryFindingBenefits struct {

	// The actual database time of the SQL statements for which SQL Tuning Advisor recommendations are not implemented.
	DbTimeBeforeRecommended *int `mandatory:"true" json:"dbTimeBeforeRecommended"`

	// The estimated database time of the above SQL statements, if SQL Tuning Advisor recommendations are implemented.
	DbTimeAfterRecommended *int `mandatory:"true" json:"dbTimeAfterRecommended"`

	// The actual database time of the SQL statements for which SQL Tuning Advisor recommendations are implemented.
	DbTimeAfterImplemented *int `mandatory:"true" json:"dbTimeAfterImplemented"`

	// The actual database time of the above SQL statements, before SQL Tuning Advisor recommendations are implemented.
	DbTimeBeforeImplemented *int `mandatory:"true" json:"dbTimeBeforeImplemented"`
}

func (m SqlTuningAdvisorTaskSummaryFindingBenefits) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlTuningAdvisorTaskSummaryFindingBenefits) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
