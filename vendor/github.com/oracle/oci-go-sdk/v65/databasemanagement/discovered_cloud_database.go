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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DiscoveredCloudDatabase The details of a cloud Oracle Database discovered in a cloud DB system discovery run.
type DiscoveredCloudDatabase struct {

	// The identifier of the discovered DB system component.
	ComponentId *string `mandatory:"true" json:"componentId"`

	// The user-friendly name for the discovered DB system component. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The name of the discovered DB system component.
	ComponentName *string `mandatory:"true" json:"componentName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The `DB_UNIQUE_NAME` of the cloud database.
	DbUniqueName *string `mandatory:"true" json:"dbUniqueName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the existing OCI resource matching the discovered DB system component.
	ResourceId *string `mandatory:"false" json:"resourceId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the existing Dbaas OCI resource matching the discovered DB system component.
	DbaasId *string `mandatory:"false" json:"dbaasId"`

	// Indicates whether the DB system component should be provisioned as an OCI resource or not.
	IsSelectedForMonitoring *bool `mandatory:"false" json:"isSelectedForMonitoring"`

	// The list of associated components.
	AssociatedComponents []AssociatedCloudComponent `mandatory:"false" json:"associatedComponents"`

	// Indicates whether the Oracle Database is part of a cluster.
	IsCluster *bool `mandatory:"false" json:"isCluster"`

	// The Oracle Database edition.
	DbEdition *string `mandatory:"false" json:"dbEdition"`

	// The Oracle Database ID.
	DbId *string `mandatory:"false" json:"dbId"`

	// The database packs licensed for the cloud Oracle Database.
	DbPacks *string `mandatory:"false" json:"dbPacks"`

	// The Oracle Database version.
	DbVersion *string `mandatory:"false" json:"dbVersion"`

	// The list of Pluggable Databases.
	PluggableDatabases []DiscoveredCloudPluggableDatabase `mandatory:"false" json:"pluggableDatabases"`

	// The list of database instances.
	DbInstances []DiscoveredCloudDbInstance `mandatory:"false" json:"dbInstances"`

	// Indicates whether Diagnostics & Management should be enabled for all the current pluggable databases in the container database.
	CanEnableAllCurrentPdbs *bool `mandatory:"false" json:"canEnableAllCurrentPdbs"`

	// Indicates whether Diagnostics & Management should be enabled automatically for all the pluggable databases in the container database.
	IsAutoEnablePluggableDatabase *bool `mandatory:"false" json:"isAutoEnablePluggableDatabase"`

	Connector CloudDbSystemDiscoveryConnector `mandatory:"false" json:"connector"`

	// The role of the Oracle Database in Oracle Data Guard configuration.
	DbRole DiscoveredCloudDatabaseDbRoleEnum `mandatory:"false" json:"dbRole,omitempty"`

	// The state of the discovered DB system component.
	Status DiscoveredCloudDbSystemComponentStatusEnum `mandatory:"false" json:"status,omitempty"`

	// The type of Oracle Database. Indicates whether the database is a Container Database,
	// Pluggable Database, or a Non-container Database.
	DbType DatabaseSubTypeEnum `mandatory:"false" json:"dbType,omitempty"`
}

// GetComponentId returns ComponentId
func (m DiscoveredCloudDatabase) GetComponentId() *string {
	return m.ComponentId
}

// GetDisplayName returns DisplayName
func (m DiscoveredCloudDatabase) GetDisplayName() *string {
	return m.DisplayName
}

// GetComponentName returns ComponentName
func (m DiscoveredCloudDatabase) GetComponentName() *string {
	return m.ComponentName
}

// GetResourceId returns ResourceId
func (m DiscoveredCloudDatabase) GetResourceId() *string {
	return m.ResourceId
}

// GetDbaasId returns DbaasId
func (m DiscoveredCloudDatabase) GetDbaasId() *string {
	return m.DbaasId
}

// GetIsSelectedForMonitoring returns IsSelectedForMonitoring
func (m DiscoveredCloudDatabase) GetIsSelectedForMonitoring() *bool {
	return m.IsSelectedForMonitoring
}

// GetStatus returns Status
func (m DiscoveredCloudDatabase) GetStatus() DiscoveredCloudDbSystemComponentStatusEnum {
	return m.Status
}

// GetAssociatedComponents returns AssociatedComponents
func (m DiscoveredCloudDatabase) GetAssociatedComponents() []AssociatedCloudComponent {
	return m.AssociatedComponents
}

func (m DiscoveredCloudDatabase) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DiscoveredCloudDatabase) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDiscoveredCloudDatabaseDbRoleEnum(string(m.DbRole)); !ok && m.DbRole != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbRole: %s. Supported values are: %s.", m.DbRole, strings.Join(GetDiscoveredCloudDatabaseDbRoleEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDiscoveredCloudDbSystemComponentStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetDiscoveredCloudDbSystemComponentStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabaseSubTypeEnum(string(m.DbType)); !ok && m.DbType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbType: %s. Supported values are: %s.", m.DbType, strings.Join(GetDatabaseSubTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DiscoveredCloudDatabase) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDiscoveredCloudDatabase DiscoveredCloudDatabase
	s := struct {
		DiscriminatorParam string `json:"componentType"`
		MarshalTypeDiscoveredCloudDatabase
	}{
		"DATABASE",
		(MarshalTypeDiscoveredCloudDatabase)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *DiscoveredCloudDatabase) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ResourceId                    *string                                    `json:"resourceId"`
		DbaasId                       *string                                    `json:"dbaasId"`
		IsSelectedForMonitoring       *bool                                      `json:"isSelectedForMonitoring"`
		Status                        DiscoveredCloudDbSystemComponentStatusEnum `json:"status"`
		AssociatedComponents          []AssociatedCloudComponent                 `json:"associatedComponents"`
		DbType                        DatabaseSubTypeEnum                        `json:"dbType"`
		IsCluster                     *bool                                      `json:"isCluster"`
		DbEdition                     *string                                    `json:"dbEdition"`
		DbId                          *string                                    `json:"dbId"`
		DbPacks                       *string                                    `json:"dbPacks"`
		DbRole                        DiscoveredCloudDatabaseDbRoleEnum          `json:"dbRole"`
		DbVersion                     *string                                    `json:"dbVersion"`
		PluggableDatabases            []DiscoveredCloudPluggableDatabase         `json:"pluggableDatabases"`
		DbInstances                   []DiscoveredCloudDbInstance                `json:"dbInstances"`
		CanEnableAllCurrentPdbs       *bool                                      `json:"canEnableAllCurrentPdbs"`
		IsAutoEnablePluggableDatabase *bool                                      `json:"isAutoEnablePluggableDatabase"`
		Connector                     clouddbsystemdiscoveryconnector            `json:"connector"`
		ComponentId                   *string                                    `json:"componentId"`
		DisplayName                   *string                                    `json:"displayName"`
		ComponentName                 *string                                    `json:"componentName"`
		CompartmentId                 *string                                    `json:"compartmentId"`
		DbUniqueName                  *string                                    `json:"dbUniqueName"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ResourceId = model.ResourceId

	m.DbaasId = model.DbaasId

	m.IsSelectedForMonitoring = model.IsSelectedForMonitoring

	m.Status = model.Status

	m.AssociatedComponents = make([]AssociatedCloudComponent, len(model.AssociatedComponents))
	copy(m.AssociatedComponents, model.AssociatedComponents)
	m.DbType = model.DbType

	m.IsCluster = model.IsCluster

	m.DbEdition = model.DbEdition

	m.DbId = model.DbId

	m.DbPacks = model.DbPacks

	m.DbRole = model.DbRole

	m.DbVersion = model.DbVersion

	m.PluggableDatabases = make([]DiscoveredCloudPluggableDatabase, len(model.PluggableDatabases))
	copy(m.PluggableDatabases, model.PluggableDatabases)
	m.DbInstances = make([]DiscoveredCloudDbInstance, len(model.DbInstances))
	copy(m.DbInstances, model.DbInstances)
	m.CanEnableAllCurrentPdbs = model.CanEnableAllCurrentPdbs

	m.IsAutoEnablePluggableDatabase = model.IsAutoEnablePluggableDatabase

	nn, e = model.Connector.UnmarshalPolymorphicJSON(model.Connector.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Connector = nn.(CloudDbSystemDiscoveryConnector)
	} else {
		m.Connector = nil
	}

	m.ComponentId = model.ComponentId

	m.DisplayName = model.DisplayName

	m.ComponentName = model.ComponentName

	m.CompartmentId = model.CompartmentId

	m.DbUniqueName = model.DbUniqueName

	return
}

// DiscoveredCloudDatabaseDbRoleEnum Enum with underlying type: string
type DiscoveredCloudDatabaseDbRoleEnum string

// Set of constants representing the allowable values for DiscoveredCloudDatabaseDbRoleEnum
const (
	DiscoveredCloudDatabaseDbRoleLogicalStandby  DiscoveredCloudDatabaseDbRoleEnum = "LOGICAL_STANDBY"
	DiscoveredCloudDatabaseDbRolePhysicalStandby DiscoveredCloudDatabaseDbRoleEnum = "PHYSICAL_STANDBY"
	DiscoveredCloudDatabaseDbRoleSnapshotStandby DiscoveredCloudDatabaseDbRoleEnum = "SNAPSHOT_STANDBY"
	DiscoveredCloudDatabaseDbRolePrimary         DiscoveredCloudDatabaseDbRoleEnum = "PRIMARY"
	DiscoveredCloudDatabaseDbRoleFarSync         DiscoveredCloudDatabaseDbRoleEnum = "FAR_SYNC"
)

var mappingDiscoveredCloudDatabaseDbRoleEnum = map[string]DiscoveredCloudDatabaseDbRoleEnum{
	"LOGICAL_STANDBY":  DiscoveredCloudDatabaseDbRoleLogicalStandby,
	"PHYSICAL_STANDBY": DiscoveredCloudDatabaseDbRolePhysicalStandby,
	"SNAPSHOT_STANDBY": DiscoveredCloudDatabaseDbRoleSnapshotStandby,
	"PRIMARY":          DiscoveredCloudDatabaseDbRolePrimary,
	"FAR_SYNC":         DiscoveredCloudDatabaseDbRoleFarSync,
}

var mappingDiscoveredCloudDatabaseDbRoleEnumLowerCase = map[string]DiscoveredCloudDatabaseDbRoleEnum{
	"logical_standby":  DiscoveredCloudDatabaseDbRoleLogicalStandby,
	"physical_standby": DiscoveredCloudDatabaseDbRolePhysicalStandby,
	"snapshot_standby": DiscoveredCloudDatabaseDbRoleSnapshotStandby,
	"primary":          DiscoveredCloudDatabaseDbRolePrimary,
	"far_sync":         DiscoveredCloudDatabaseDbRoleFarSync,
}

// GetDiscoveredCloudDatabaseDbRoleEnumValues Enumerates the set of values for DiscoveredCloudDatabaseDbRoleEnum
func GetDiscoveredCloudDatabaseDbRoleEnumValues() []DiscoveredCloudDatabaseDbRoleEnum {
	values := make([]DiscoveredCloudDatabaseDbRoleEnum, 0)
	for _, v := range mappingDiscoveredCloudDatabaseDbRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetDiscoveredCloudDatabaseDbRoleEnumStringValues Enumerates the set of values in String for DiscoveredCloudDatabaseDbRoleEnum
func GetDiscoveredCloudDatabaseDbRoleEnumStringValues() []string {
	return []string{
		"LOGICAL_STANDBY",
		"PHYSICAL_STANDBY",
		"SNAPSHOT_STANDBY",
		"PRIMARY",
		"FAR_SYNC",
	}
}

// GetMappingDiscoveredCloudDatabaseDbRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDiscoveredCloudDatabaseDbRoleEnum(val string) (DiscoveredCloudDatabaseDbRoleEnum, bool) {
	enum, ok := mappingDiscoveredCloudDatabaseDbRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
