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

// JobExecutionSummary A summary of a job execution on a Managed Database.
type JobExecutionSummary struct {

	// The identifier of the job execution.
	Id *string `mandatory:"true" json:"id"`

	// The name of the job execution.
	Name *string `mandatory:"true" json:"name"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the parent job resides.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of Managed Database associated with the job execution.
	ManagedDatabaseId *string `mandatory:"true" json:"managedDatabaseId"`

	// The name of the Managed Database associated with the job execution.
	ManagedDatabaseName *string `mandatory:"true" json:"managedDatabaseName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the parent job.
	JobId *string `mandatory:"true" json:"jobId"`

	// The name of the parent job.
	JobName *string `mandatory:"true" json:"jobName"`

	// The status of the job execution.
	Status JobExecutionStatusEnum `mandatory:"true" json:"status"`

	// The date and time when the job execution was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database Group where the parent job has to be executed.
	ManagedDatabaseGroupId *string `mandatory:"false" json:"managedDatabaseGroupId"`

	// The type of Oracle Database installation.
	DatabaseType DatabaseTypeEnum `mandatory:"false" json:"databaseType,omitempty"`

	// The subtype of the Oracle Database. Indicates whether the database is a Container Database, Pluggable Database, or a Non-container Database.
	DatabaseSubType DatabaseSubTypeEnum `mandatory:"false" json:"databaseSubType,omitempty"`

	// A list of the supported infrastructure that can be used to deploy the database.
	DeploymentType DeploymentTypeEnum `mandatory:"false" json:"deploymentType,omitempty"`

	// Indicates whether the Oracle Database is part of a cluster.
	IsCluster *bool `mandatory:"false" json:"isCluster"`

	// The workload type of the Autonomous Database.
	WorkloadType WorkloadTypeEnum `mandatory:"false" json:"workloadType,omitempty"`

	// The date and time when the job execution was completed.
	TimeCompleted *common.SDKTime `mandatory:"false" json:"timeCompleted"`
}

func (m JobExecutionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JobExecutionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingJobExecutionStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetJobExecutionStatusEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDatabaseTypeEnum(string(m.DatabaseType)); !ok && m.DatabaseType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseType: %s. Supported values are: %s.", m.DatabaseType, strings.Join(GetDatabaseTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabaseSubTypeEnum(string(m.DatabaseSubType)); !ok && m.DatabaseSubType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseSubType: %s. Supported values are: %s.", m.DatabaseSubType, strings.Join(GetDatabaseSubTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDeploymentTypeEnum(string(m.DeploymentType)); !ok && m.DeploymentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeploymentType: %s. Supported values are: %s.", m.DeploymentType, strings.Join(GetDeploymentTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingWorkloadTypeEnum(string(m.WorkloadType)); !ok && m.WorkloadType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for WorkloadType: %s. Supported values are: %s.", m.WorkloadType, strings.Join(GetWorkloadTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
