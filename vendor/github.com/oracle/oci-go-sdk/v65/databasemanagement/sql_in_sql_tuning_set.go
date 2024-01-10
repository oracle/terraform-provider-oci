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

// SqlInSqlTuningSet Sql information in the Sql tuning set.
type SqlInSqlTuningSet struct {

	// The unique Sql identifier.
	SqlId *string `mandatory:"true" json:"sqlId"`

	// Plan hash value of the Sql statement.
	PlanHashValue *int64 `mandatory:"true" json:"planHashValue"`

	// Sql text.
	SqlText *string `mandatory:"false" json:"sqlText"`

	// The unique container database identifier.
	ContainerDatabaseId *int64 `mandatory:"false" json:"containerDatabaseId"`

	// The schema name of the Sql.
	Schema *string `mandatory:"false" json:"schema"`

	// The module of the Sql.
	Module *string `mandatory:"false" json:"module"`

	// A list of the Sqls associated with the Sql tuning set.
	Metrics []SqlMetrics `mandatory:"false" json:"metrics"`
}

func (m SqlInSqlTuningSet) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlInSqlTuningSet) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
