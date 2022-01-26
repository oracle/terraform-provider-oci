// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// JobExecutionSummary A summary of a job execution on a Managed Database.
type JobExecutionSummary struct {

	// The identifier of the job execution.
	Id *string `mandatory:"true" json:"id"`

	// The name of the job execution.
	Name *string `mandatory:"true" json:"name"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment in which the parent job resides.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of Managed Database associated with the job execution.
	ManagedDatabaseId *string `mandatory:"true" json:"managedDatabaseId"`

	// The name of the Managed Database associated with the job execution.
	ManagedDatabaseName *string `mandatory:"true" json:"managedDatabaseName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the parent job.
	JobId *string `mandatory:"true" json:"jobId"`

	// The name of the parent job.
	JobName *string `mandatory:"true" json:"jobName"`

	// The status of the job execution.
	Status JobExecutionStatusEnum `mandatory:"true" json:"status"`

	// The date and time when the job execution was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database Group where the parent job has to be executed.
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
