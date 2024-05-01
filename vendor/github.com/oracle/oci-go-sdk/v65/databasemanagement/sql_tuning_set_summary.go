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

// SqlTuningSetSummary The summary information of a SQL tuning set.
type SqlTuningSetSummary struct {

	// The name of the SQL tuning set.
	Name *string `mandatory:"true" json:"name"`

	// The owner of the SQL tuning set.
	Owner *string `mandatory:"true" json:"owner"`

	// The description of the SQL tuning set.
	Description *string `mandatory:"false" json:"description"`

	// The number of SQL statements in the SQL tuning set.
	StatementCounts *int `mandatory:"false" json:"statementCounts"`

	// The unique Sql tuning set identifier. This is not OCID.
	Id *int `mandatory:"false" json:"id"`

	// The created time of the Sql tuning set.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Last modified time of the Sql tuning set.
	TimeLastModified *common.SDKTime `mandatory:"false" json:"timeLastModified"`

	// Current status of the Sql tuning set.
	Status SqlTuningSetStatusTypesEnum `mandatory:"false" json:"status,omitempty"`

	// Name of the Sql tuning set scheduler job.
	ScheduledJobName *string `mandatory:"false" json:"scheduledJobName"`

	// Latest execution error of the plsql that was submitted as a scheduler job.
	ErrorMessage *string `mandatory:"false" json:"errorMessage"`
}

func (m SqlTuningSetSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlTuningSetSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSqlTuningSetStatusTypesEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetSqlTuningSetStatusTypesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
