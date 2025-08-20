// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseBackupMetrics The details of the backup metrics for a single database.
type DatabaseBackupMetrics struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
	DbId *string `mandatory:"true" json:"dbId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where the Managed Database resides.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The type of Oracle Database installation.
	DatabaseType DatabaseTypeEnum `mandatory:"true" json:"databaseType"`

	// The subtype of the Oracle Database. Indicates whether the database is a Container Database,
	// Pluggable Database, Non-container Database, Autonomous Database, or Autonomous Container Database.
	DatabaseSubType DatabaseSubTypeEnum `mandatory:"true" json:"databaseSubType"`

	// The infrastructure used to deploy the Oracle Database.
	DeploymentType DeploymentTypeEnum `mandatory:"true" json:"deploymentType"`

	// The display name of the Managed Database.
	DatabaseName *string `mandatory:"true" json:"databaseName"`

	// The database role of the Managed Database.
	DbRole DbRoleEnum `mandatory:"true" json:"dbRole"`

	// Indicates whether Data Guard is enabled for the database.
	IsDataguardEnabled *bool `mandatory:"true" json:"isDataguardEnabled"`

	LastBackupSummary *DatabaseBackupSummary `mandatory:"true" json:"lastBackupSummary"`

	// A list of database health metrics such as BackupDuration and BackupSize.
	Metrics []HaMetricDefinition `mandatory:"true" json:"metrics"`

	// The retention period of the database.
	RetentionPeriodInDays *string `mandatory:"true" json:"retentionPeriodInDays"`

	// Indicates whether automatic backup is enabled or not. This is only applicable for cloud databases.
	IsAutomaticBackupEnabled *bool `mandatory:"true" json:"isAutomaticBackupEnabled"`

	// The array of metric values for different timestamps.
	TimeSeriesMetrics []MetricsAggregationRange `mandatory:"true" json:"timeSeriesMetrics"`

	// The Oracle Database version.
	DatabaseVersion *string `mandatory:"false" json:"databaseVersion"`

	// The workload type of the Autonomous Database.
	WorkloadType WorkloadTypeEnum `mandatory:"false" json:"workloadType,omitempty"`

	// The database ID of the Managed Database. Every database had its own ID and that value is captured here.
	DatabaseId *string `mandatory:"false" json:"databaseId"`

	// The database unique name of the Managed Database.
	DbUniqueName *string `mandatory:"false" json:"dbUniqueName"`
}

func (m DatabaseBackupMetrics) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseBackupMetrics) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDatabaseTypeEnum(string(m.DatabaseType)); !ok && m.DatabaseType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseType: %s. Supported values are: %s.", m.DatabaseType, strings.Join(GetDatabaseTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabaseSubTypeEnum(string(m.DatabaseSubType)); !ok && m.DatabaseSubType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseSubType: %s. Supported values are: %s.", m.DatabaseSubType, strings.Join(GetDatabaseSubTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDeploymentTypeEnum(string(m.DeploymentType)); !ok && m.DeploymentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeploymentType: %s. Supported values are: %s.", m.DeploymentType, strings.Join(GetDeploymentTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDbRoleEnum(string(m.DbRole)); !ok && m.DbRole != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbRole: %s. Supported values are: %s.", m.DbRole, strings.Join(GetDbRoleEnumStringValues(), ",")))
	}

	if _, ok := GetMappingWorkloadTypeEnum(string(m.WorkloadType)); !ok && m.WorkloadType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for WorkloadType: %s. Supported values are: %s.", m.WorkloadType, strings.Join(GetWorkloadTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
