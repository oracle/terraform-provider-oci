// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// MySqlDatabaseUsageMetrics The list of aggregated metrics for Managed MySQL Databases in the fleet.
type MySqlDatabaseUsageMetrics struct {

	// The OCID of the compartment where the Managed MySQL Database resides.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the Managed MySQL Database.
	DatabaseName *string `mandatory:"true" json:"databaseName"`

	// Indicates MySQL Database type, ONPREMISE or MySQL Database System.
	DatabaseType *string `mandatory:"true" json:"databaseType"`

	// The type of MySQL Database System.
	MdsDeploymentType *string `mandatory:"true" json:"mdsDeploymentType"`

	// The lifecycle state of the MySQL Database System.
	MdslifecycleState *string `mandatory:"true" json:"mdslifecycleState"`

	// The version of the MySQL Database.
	DatabaseVersion *string `mandatory:"true" json:"databaseVersion"`

	// The OCID of the Managed MySQL Database.
	DbId *string `mandatory:"true" json:"dbId"`

	// The status of the MySQL Database. Indicates whether the status of the database
	// is UP, DOWN, or UNKNOWN at the current time.
	DatabaseStatus MySqlDatabaseStatusEnum `mandatory:"true" json:"databaseStatus"`

	// A list of the database health metrics like CPU, Storage, and Memory.
	Metrics []MySqlFleetMetricDefinition `mandatory:"true" json:"metrics"`
}

func (m MySqlDatabaseUsageMetrics) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MySqlDatabaseUsageMetrics) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMySqlDatabaseStatusEnum(string(m.DatabaseStatus)); !ok && m.DatabaseStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseStatus: %s. Supported values are: %s.", m.DatabaseStatus, strings.Join(GetMySqlDatabaseStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
