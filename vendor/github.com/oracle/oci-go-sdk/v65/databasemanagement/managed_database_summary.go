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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ManagedDatabaseSummary A summary of the Managed Database.
type ManagedDatabaseSummary struct {

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
	// if the Managed Database is a Pluggable Database.
	ParentContainerId *string `mandatory:"false" json:"parentContainerId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external
	// DB system that this Managed Database is part of.
	DbSystemId *string `mandatory:"false" json:"dbSystemId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the storage DB system.
	StorageSystemId *string `mandatory:"false" json:"storageSystemId"`

	// The Oracle Database version.
	DatabaseVersion *string `mandatory:"false" json:"databaseVersion"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// System tags can be viewed by users, but can only be created by the system.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The list of feature configurations
	DbmgmtFeatureConfigs []DatabaseFeatureConfiguration `mandatory:"false" json:"dbmgmtFeatureConfigs"`

	// The operating system of database.
	DatabasePlatformName *string `mandatory:"false" json:"databasePlatformName"`
}

func (m ManagedDatabaseSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagedDatabaseSummary) ValidateEnumValue() (bool, error) {
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
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ManagedDatabaseSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DeploymentType       DeploymentTypeEnum                `json:"deploymentType"`
		ManagementOption     ManagementOptionEnum              `json:"managementOption"`
		WorkloadType         WorkloadTypeEnum                  `json:"workloadType"`
		ParentContainerId    *string                           `json:"parentContainerId"`
		DbSystemId           *string                           `json:"dbSystemId"`
		StorageSystemId      *string                           `json:"storageSystemId"`
		DatabaseVersion      *string                           `json:"databaseVersion"`
		FreeformTags         map[string]string                 `json:"freeformTags"`
		DefinedTags          map[string]map[string]interface{} `json:"definedTags"`
		SystemTags           map[string]map[string]interface{} `json:"systemTags"`
		DbmgmtFeatureConfigs []databasefeatureconfiguration    `json:"dbmgmtFeatureConfigs"`
		DatabasePlatformName *string                           `json:"databasePlatformName"`
		Id                   *string                           `json:"id"`
		CompartmentId        *string                           `json:"compartmentId"`
		Name                 *string                           `json:"name"`
		DatabaseType         DatabaseTypeEnum                  `json:"databaseType"`
		DatabaseSubType      DatabaseSubTypeEnum               `json:"databaseSubType"`
		IsCluster            *bool                             `json:"isCluster"`
		TimeCreated          *common.SDKTime                   `json:"timeCreated"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DeploymentType = model.DeploymentType

	m.ManagementOption = model.ManagementOption

	m.WorkloadType = model.WorkloadType

	m.ParentContainerId = model.ParentContainerId

	m.DbSystemId = model.DbSystemId

	m.StorageSystemId = model.StorageSystemId

	m.DatabaseVersion = model.DatabaseVersion

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.DbmgmtFeatureConfigs = make([]DatabaseFeatureConfiguration, len(model.DbmgmtFeatureConfigs))
	for i, n := range model.DbmgmtFeatureConfigs {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.DbmgmtFeatureConfigs[i] = nn.(DatabaseFeatureConfiguration)
		} else {
			m.DbmgmtFeatureConfigs[i] = nil
		}
	}
	m.DatabasePlatformName = model.DatabasePlatformName

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.Name = model.Name

	m.DatabaseType = model.DatabaseType

	m.DatabaseSubType = model.DatabaseSubType

	m.IsCluster = model.IsCluster

	m.TimeCreated = model.TimeCreated

	return
}
