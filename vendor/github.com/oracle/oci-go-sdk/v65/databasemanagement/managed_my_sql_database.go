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

// ManagedMySqlDatabase The details of the Managed MySQL Database.
type ManagedMySqlDatabase struct {

	// The OCID of the Managed MySQL Database.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the MySQL Database.
	DbName *string `mandatory:"true" json:"dbName"`

	// The version of the MySQL Database.
	DbVersion *string `mandatory:"true" json:"dbVersion"`

	// The date and time the Managed MySQL Database was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The name of the Managed MySQL Database.
	Name *string `mandatory:"true" json:"name"`

	// The customer's selected type for HeatWave management.
	HeatWaveManagementType ManagedMySqlDatabaseHeatWaveManagementTypeEnum `mandatory:"false" json:"heatWaveManagementType,omitempty"`

	// The name of the HeatWave cluster.
	HeatWaveClusterDisplayName *string `mandatory:"false" json:"heatWaveClusterDisplayName"`

	// Indicates whether HeatWave is enabled for the MySQL Database System or not.
	IsHeatWaveEnabled *bool `mandatory:"false" json:"isHeatWaveEnabled"`

	// Indicates whether HeatWave Lakehouse is enabled for the MySQL Database System or not.
	IsLakehouseEnabled *bool `mandatory:"false" json:"isLakehouseEnabled"`

	// The shape of the nodes in the HeatWave cluster.
	HeatWaveNodeShape *string `mandatory:"false" json:"heatWaveNodeShape"`

	// The total memory belonging to the HeatWave cluster in GBs.
	HeatWaveMemorySize *int `mandatory:"false" json:"heatWaveMemorySize"`

	// The information about individual HeatWave nodes in the cluster.
	HeatWaveNodes []HeatWaveNode `mandatory:"false" json:"heatWaveNodes"`

	// Indicates whether the HeatWave cluster is active or not.
	IsHeatWaveActive *bool `mandatory:"false" json:"isHeatWaveActive"`

	// The date and time the Managed MySQL Database was created.
	TimeCreatedHeatWave *common.SDKTime `mandatory:"false" json:"timeCreatedHeatWave"`

	// The type of the MySQL Database. Indicates whether the database
	// is external or MDS.
	DatabaseType MySqlTypeEnum `mandatory:"false" json:"databaseType,omitempty"`

	// Indicates database management status.
	ManagementState ManagementStateEnum `mandatory:"false" json:"managementState,omitempty"`

	// Indicates lifecycle  state of the resource.
	LifecycleState LifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The date and time the Managed MySQL Database was updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m ManagedMySqlDatabase) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagedMySqlDatabase) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingManagedMySqlDatabaseHeatWaveManagementTypeEnum(string(m.HeatWaveManagementType)); !ok && m.HeatWaveManagementType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for HeatWaveManagementType: %s. Supported values are: %s.", m.HeatWaveManagementType, strings.Join(GetManagedMySqlDatabaseHeatWaveManagementTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMySqlTypeEnum(string(m.DatabaseType)); !ok && m.DatabaseType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseType: %s. Supported values are: %s.", m.DatabaseType, strings.Join(GetMySqlTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingManagementStateEnum(string(m.ManagementState)); !ok && m.ManagementState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ManagementState: %s. Supported values are: %s.", m.ManagementState, strings.Join(GetManagementStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStatesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ManagedMySqlDatabaseHeatWaveManagementTypeEnum Enum with underlying type: string
type ManagedMySqlDatabaseHeatWaveManagementTypeEnum string

// Set of constants representing the allowable values for ManagedMySqlDatabaseHeatWaveManagementTypeEnum
const (
	ManagedMySqlDatabaseHeatWaveManagementTypeBasic ManagedMySqlDatabaseHeatWaveManagementTypeEnum = "BASIC"
	ManagedMySqlDatabaseHeatWaveManagementTypeFull  ManagedMySqlDatabaseHeatWaveManagementTypeEnum = "FULL"
)

var mappingManagedMySqlDatabaseHeatWaveManagementTypeEnum = map[string]ManagedMySqlDatabaseHeatWaveManagementTypeEnum{
	"BASIC": ManagedMySqlDatabaseHeatWaveManagementTypeBasic,
	"FULL":  ManagedMySqlDatabaseHeatWaveManagementTypeFull,
}

var mappingManagedMySqlDatabaseHeatWaveManagementTypeEnumLowerCase = map[string]ManagedMySqlDatabaseHeatWaveManagementTypeEnum{
	"basic": ManagedMySqlDatabaseHeatWaveManagementTypeBasic,
	"full":  ManagedMySqlDatabaseHeatWaveManagementTypeFull,
}

// GetManagedMySqlDatabaseHeatWaveManagementTypeEnumValues Enumerates the set of values for ManagedMySqlDatabaseHeatWaveManagementTypeEnum
func GetManagedMySqlDatabaseHeatWaveManagementTypeEnumValues() []ManagedMySqlDatabaseHeatWaveManagementTypeEnum {
	values := make([]ManagedMySqlDatabaseHeatWaveManagementTypeEnum, 0)
	for _, v := range mappingManagedMySqlDatabaseHeatWaveManagementTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetManagedMySqlDatabaseHeatWaveManagementTypeEnumStringValues Enumerates the set of values in String for ManagedMySqlDatabaseHeatWaveManagementTypeEnum
func GetManagedMySqlDatabaseHeatWaveManagementTypeEnumStringValues() []string {
	return []string{
		"BASIC",
		"FULL",
	}
}

// GetMappingManagedMySqlDatabaseHeatWaveManagementTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingManagedMySqlDatabaseHeatWaveManagementTypeEnum(val string) (ManagedMySqlDatabaseHeatWaveManagementTypeEnum, bool) {
	enum, ok := mappingManagedMySqlDatabaseHeatWaveManagementTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
