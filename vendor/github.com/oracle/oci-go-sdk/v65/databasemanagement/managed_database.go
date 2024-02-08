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

// ManagedDatabase The details of a Managed Database.
type ManagedDatabase struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the Managed Database.
	Name *string `mandatory:"true" json:"name"`

	// The type of Oracle Database installation.
	DatabaseType DatabaseTypeEnum `mandatory:"true" json:"databaseType"`

	// The subtype of the Oracle Database. Indicates whether the database is a Container Database,
	// Pluggable Database, Non-container Database, Autonomous Database, or Autonomous Container Database.
	DatabaseSubType DatabaseSubTypeEnum `mandatory:"true" json:"databaseSubType"`

	// Indicates whether the Oracle Database is part of a cluster.
	IsCluster *bool `mandatory:"true" json:"isCluster"`

	// The date and time the Managed Database was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The infrastructure used to deploy the Oracle Database.
	DeploymentType DeploymentTypeEnum `mandatory:"false" json:"deploymentType,omitempty"`

	// The management option used when enabling Database Management.
	ManagementOption ManagementOptionEnum `mandatory:"false" json:"managementOption,omitempty"`

	// The workload type of the Autonomous Database.
	WorkloadType WorkloadTypeEnum `mandatory:"false" json:"workloadType,omitempty"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the parent Container Database
	// if Managed Database is a Pluggable Database.
	ParentContainerId *string `mandatory:"false" json:"parentContainerId"`

	// A list of Managed Database Groups that the Managed Database belongs to.
	ManagedDatabaseGroups []ParentGroup `mandatory:"false" json:"managedDatabaseGroups"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external
	// DB system that this Managed Database is part of.
	DbSystemId *string `mandatory:"false" json:"dbSystemId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the storage DB system.
	StorageSystemId *string `mandatory:"false" json:"storageSystemId"`

	// The status of the Oracle Database. Indicates whether the status of the database
	// is UP, DOWN, or UNKNOWN at the current time.
	DatabaseStatus DatabaseStatusEnum `mandatory:"false" json:"databaseStatus,omitempty"`

	// The name of the parent Container Database.
	ParentContainerName *string `mandatory:"false" json:"parentContainerName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment
	// in which the parent Container Database resides, if the Managed Database
	// is a Pluggable Database (PDB).
	ParentContainerCompartmentId *string `mandatory:"false" json:"parentContainerCompartmentId"`

	// The number of Oracle Real Application Clusters (Oracle RAC) database instances.
	InstanceCount *int `mandatory:"false" json:"instanceCount"`

	// The details of the Oracle Real Application Clusters (Oracle RAC) database instances.
	InstanceDetails []InstanceDetails `mandatory:"false" json:"instanceDetails"`

	// The number of PDBs in the Container Database.
	PdbCount *int `mandatory:"false" json:"pdbCount"`

	// The status of the PDB in the Container Database.
	PdbStatus []PdbStatusDetails `mandatory:"false" json:"pdbStatus"`

	// The additional details specific to a type of database defined in `{"key": "value"}` format.
	// Example: `{"bar-key": "value"}`
	AdditionalDetails map[string]string `mandatory:"false" json:"additionalDetails"`
}

func (m ManagedDatabase) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagedDatabase) ValidateEnumValue() (bool, error) {
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
	if _, ok := GetMappingManagementOptionEnum(string(m.ManagementOption)); !ok && m.ManagementOption != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ManagementOption: %s. Supported values are: %s.", m.ManagementOption, strings.Join(GetManagementOptionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingWorkloadTypeEnum(string(m.WorkloadType)); !ok && m.WorkloadType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for WorkloadType: %s. Supported values are: %s.", m.WorkloadType, strings.Join(GetWorkloadTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabaseStatusEnum(string(m.DatabaseStatus)); !ok && m.DatabaseStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseStatus: %s. Supported values are: %s.", m.DatabaseStatus, strings.Join(GetDatabaseStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
