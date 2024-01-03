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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DiscoveredExternalDatabase The details of an external Oracle Database discovered in an external DB system discovery run.
type DiscoveredExternalDatabase struct {

	// The identifier of the discovered DB system component.
	ComponentId *string `mandatory:"true" json:"componentId"`

	// The user-friendly name for the discovered DB system component. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The name of the discovered DB system component.
	ComponentName *string `mandatory:"true" json:"componentName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The `DB_UNIQUE_NAME` of the external database.
	DbUniqueName *string `mandatory:"true" json:"dbUniqueName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the existing OCI resource matching the discovered DB system component.
	ResourceId *string `mandatory:"false" json:"resourceId"`

	// Indicates whether the DB system component should be provisioned as an OCI resource or not.
	IsSelectedForMonitoring *bool `mandatory:"false" json:"isSelectedForMonitoring"`

	// The list of associated components.
	AssociatedComponents []AssociatedComponent `mandatory:"false" json:"associatedComponents"`

	// Indicates whether the Oracle Database is part of a cluster.
	IsCluster *bool `mandatory:"false" json:"isCluster"`

	// The Oracle Database edition.
	DbEdition *string `mandatory:"false" json:"dbEdition"`

	// The Oracle Database ID.
	DbId *string `mandatory:"false" json:"dbId"`

	// The database packs licensed for the external Oracle Database.
	DbPacks *string `mandatory:"false" json:"dbPacks"`

	// The Oracle Database version.
	DbVersion *string `mandatory:"false" json:"dbVersion"`

	// The list of Pluggable Databases.
	PluggableDatabases []DiscoveredExternalPluggableDatabase `mandatory:"false" json:"pluggableDatabases"`

	Connector ExternalDbSystemDiscoveryConnector `mandatory:"false" json:"connector"`

	// The role of the Oracle Database in Oracle Data Guard configuration.
	DbRole DiscoveredExternalDatabaseDbRoleEnum `mandatory:"false" json:"dbRole,omitempty"`

	// The state of the discovered DB system component.
	Status DiscoveredExternalDbSystemComponentStatusEnum `mandatory:"false" json:"status,omitempty"`

	// The type of Oracle Database. Indicates whether the database is a Container Database,
	// Pluggable Database, or a Non-container Database.
	DbType DatabaseSubTypeEnum `mandatory:"false" json:"dbType,omitempty"`
}

// GetComponentId returns ComponentId
func (m DiscoveredExternalDatabase) GetComponentId() *string {
	return m.ComponentId
}

// GetDisplayName returns DisplayName
func (m DiscoveredExternalDatabase) GetDisplayName() *string {
	return m.DisplayName
}

// GetComponentName returns ComponentName
func (m DiscoveredExternalDatabase) GetComponentName() *string {
	return m.ComponentName
}

// GetResourceId returns ResourceId
func (m DiscoveredExternalDatabase) GetResourceId() *string {
	return m.ResourceId
}

// GetIsSelectedForMonitoring returns IsSelectedForMonitoring
func (m DiscoveredExternalDatabase) GetIsSelectedForMonitoring() *bool {
	return m.IsSelectedForMonitoring
}

// GetStatus returns Status
func (m DiscoveredExternalDatabase) GetStatus() DiscoveredExternalDbSystemComponentStatusEnum {
	return m.Status
}

// GetAssociatedComponents returns AssociatedComponents
func (m DiscoveredExternalDatabase) GetAssociatedComponents() []AssociatedComponent {
	return m.AssociatedComponents
}

func (m DiscoveredExternalDatabase) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DiscoveredExternalDatabase) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDiscoveredExternalDatabaseDbRoleEnum(string(m.DbRole)); !ok && m.DbRole != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbRole: %s. Supported values are: %s.", m.DbRole, strings.Join(GetDiscoveredExternalDatabaseDbRoleEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDiscoveredExternalDbSystemComponentStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetDiscoveredExternalDbSystemComponentStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabaseSubTypeEnum(string(m.DbType)); !ok && m.DbType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbType: %s. Supported values are: %s.", m.DbType, strings.Join(GetDatabaseSubTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DiscoveredExternalDatabase) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDiscoveredExternalDatabase DiscoveredExternalDatabase
	s := struct {
		DiscriminatorParam string `json:"componentType"`
		MarshalTypeDiscoveredExternalDatabase
	}{
		"DATABASE",
		(MarshalTypeDiscoveredExternalDatabase)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *DiscoveredExternalDatabase) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ResourceId              *string                                       `json:"resourceId"`
		IsSelectedForMonitoring *bool                                         `json:"isSelectedForMonitoring"`
		Status                  DiscoveredExternalDbSystemComponentStatusEnum `json:"status"`
		AssociatedComponents    []AssociatedComponent                         `json:"associatedComponents"`
		DbType                  DatabaseSubTypeEnum                           `json:"dbType"`
		IsCluster               *bool                                         `json:"isCluster"`
		DbEdition               *string                                       `json:"dbEdition"`
		DbId                    *string                                       `json:"dbId"`
		DbPacks                 *string                                       `json:"dbPacks"`
		DbRole                  DiscoveredExternalDatabaseDbRoleEnum          `json:"dbRole"`
		DbVersion               *string                                       `json:"dbVersion"`
		PluggableDatabases      []DiscoveredExternalPluggableDatabase         `json:"pluggableDatabases"`
		Connector               externaldbsystemdiscoveryconnector            `json:"connector"`
		ComponentId             *string                                       `json:"componentId"`
		DisplayName             *string                                       `json:"displayName"`
		ComponentName           *string                                       `json:"componentName"`
		CompartmentId           *string                                       `json:"compartmentId"`
		DbUniqueName            *string                                       `json:"dbUniqueName"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ResourceId = model.ResourceId

	m.IsSelectedForMonitoring = model.IsSelectedForMonitoring

	m.Status = model.Status

	m.AssociatedComponents = make([]AssociatedComponent, len(model.AssociatedComponents))
	copy(m.AssociatedComponents, model.AssociatedComponents)
	m.DbType = model.DbType

	m.IsCluster = model.IsCluster

	m.DbEdition = model.DbEdition

	m.DbId = model.DbId

	m.DbPacks = model.DbPacks

	m.DbRole = model.DbRole

	m.DbVersion = model.DbVersion

	m.PluggableDatabases = make([]DiscoveredExternalPluggableDatabase, len(model.PluggableDatabases))
	copy(m.PluggableDatabases, model.PluggableDatabases)
	nn, e = model.Connector.UnmarshalPolymorphicJSON(model.Connector.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Connector = nn.(ExternalDbSystemDiscoveryConnector)
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

// DiscoveredExternalDatabaseDbRoleEnum Enum with underlying type: string
type DiscoveredExternalDatabaseDbRoleEnum string

// Set of constants representing the allowable values for DiscoveredExternalDatabaseDbRoleEnum
const (
	DiscoveredExternalDatabaseDbRoleLogicalStandby  DiscoveredExternalDatabaseDbRoleEnum = "LOGICAL_STANDBY"
	DiscoveredExternalDatabaseDbRolePhysicalStandby DiscoveredExternalDatabaseDbRoleEnum = "PHYSICAL_STANDBY"
	DiscoveredExternalDatabaseDbRoleSnapshotStandby DiscoveredExternalDatabaseDbRoleEnum = "SNAPSHOT_STANDBY"
	DiscoveredExternalDatabaseDbRolePrimary         DiscoveredExternalDatabaseDbRoleEnum = "PRIMARY"
	DiscoveredExternalDatabaseDbRoleFarSync         DiscoveredExternalDatabaseDbRoleEnum = "FAR_SYNC"
)

var mappingDiscoveredExternalDatabaseDbRoleEnum = map[string]DiscoveredExternalDatabaseDbRoleEnum{
	"LOGICAL_STANDBY":  DiscoveredExternalDatabaseDbRoleLogicalStandby,
	"PHYSICAL_STANDBY": DiscoveredExternalDatabaseDbRolePhysicalStandby,
	"SNAPSHOT_STANDBY": DiscoveredExternalDatabaseDbRoleSnapshotStandby,
	"PRIMARY":          DiscoveredExternalDatabaseDbRolePrimary,
	"FAR_SYNC":         DiscoveredExternalDatabaseDbRoleFarSync,
}

var mappingDiscoveredExternalDatabaseDbRoleEnumLowerCase = map[string]DiscoveredExternalDatabaseDbRoleEnum{
	"logical_standby":  DiscoveredExternalDatabaseDbRoleLogicalStandby,
	"physical_standby": DiscoveredExternalDatabaseDbRolePhysicalStandby,
	"snapshot_standby": DiscoveredExternalDatabaseDbRoleSnapshotStandby,
	"primary":          DiscoveredExternalDatabaseDbRolePrimary,
	"far_sync":         DiscoveredExternalDatabaseDbRoleFarSync,
}

// GetDiscoveredExternalDatabaseDbRoleEnumValues Enumerates the set of values for DiscoveredExternalDatabaseDbRoleEnum
func GetDiscoveredExternalDatabaseDbRoleEnumValues() []DiscoveredExternalDatabaseDbRoleEnum {
	values := make([]DiscoveredExternalDatabaseDbRoleEnum, 0)
	for _, v := range mappingDiscoveredExternalDatabaseDbRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetDiscoveredExternalDatabaseDbRoleEnumStringValues Enumerates the set of values in String for DiscoveredExternalDatabaseDbRoleEnum
func GetDiscoveredExternalDatabaseDbRoleEnumStringValues() []string {
	return []string{
		"LOGICAL_STANDBY",
		"PHYSICAL_STANDBY",
		"SNAPSHOT_STANDBY",
		"PRIMARY",
		"FAR_SYNC",
	}
}

// GetMappingDiscoveredExternalDatabaseDbRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDiscoveredExternalDatabaseDbRoleEnum(val string) (DiscoveredExternalDatabaseDbRoleEnum, bool) {
	enum, ok := mappingDiscoveredExternalDatabaseDbRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
