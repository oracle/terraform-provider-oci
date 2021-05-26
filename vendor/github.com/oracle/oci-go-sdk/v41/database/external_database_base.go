// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/v41/common"
)

// ExternalDatabaseBase A resource that allows you to manage an Oracle Database located outside of Oracle Cloud using Oracle Cloud Infrastructure's Console and APIs.
type ExternalDatabaseBase struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-friendly name for the external database. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure external database resource.
	Id *string `mandatory:"true" json:"id"`

	// The current state of the Oracle Cloud Infrastructure external database resource.
	LifecycleState ExternalDatabaseBaseLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the database was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The `DB_UNIQUE_NAME` of the external database.
	DbUniqueName *string `mandatory:"false" json:"dbUniqueName"`

	// The Oracle Database ID, which identifies an Oracle Database located outside of Oracle Cloud.
	DbId *string `mandatory:"false" json:"dbId"`

	// The Oracle Database version.
	DatabaseVersion *string `mandatory:"false" json:"databaseVersion"`

	// The Oracle Database edition.
	DatabaseEdition ExternalDatabaseBaseDatabaseEditionEnum `mandatory:"false" json:"databaseEdition,omitempty"`

	// The time zone of the external database.
	// It is a time zone offset (a character type in the format '[+|-]TZH:TZM') or a time zone region name,
	// depending on how the time zone value was specified when the database was created / last altered.
	TimeZone *string `mandatory:"false" json:"timeZone"`

	// The character set of the external database.
	CharacterSet *string `mandatory:"false" json:"characterSet"`

	// The national character of the external database.
	NcharacterSet *string `mandatory:"false" json:"ncharacterSet"`

	// The database packs licensed for the external Oracle Database.
	DbPacks *string `mandatory:"false" json:"dbPacks"`

	// The Oracle Database configuration
	DatabaseConfiguration ExternalDatabaseBaseDatabaseConfigurationEnum `mandatory:"false" json:"databaseConfiguration,omitempty"`

	DatabaseManagementConfig *DatabaseManagementConfig `mandatory:"false" json:"databaseManagementConfig"`
}

func (m ExternalDatabaseBase) String() string {
	return common.PointerString(m)
}

// ExternalDatabaseBaseLifecycleStateEnum Enum with underlying type: string
type ExternalDatabaseBaseLifecycleStateEnum string

// Set of constants representing the allowable values for ExternalDatabaseBaseLifecycleStateEnum
const (
	ExternalDatabaseBaseLifecycleStateProvisioning ExternalDatabaseBaseLifecycleStateEnum = "PROVISIONING"
	ExternalDatabaseBaseLifecycleStateNotConnected ExternalDatabaseBaseLifecycleStateEnum = "NOT_CONNECTED"
	ExternalDatabaseBaseLifecycleStateAvailable    ExternalDatabaseBaseLifecycleStateEnum = "AVAILABLE"
	ExternalDatabaseBaseLifecycleStateUpdating     ExternalDatabaseBaseLifecycleStateEnum = "UPDATING"
	ExternalDatabaseBaseLifecycleStateTerminating  ExternalDatabaseBaseLifecycleStateEnum = "TERMINATING"
	ExternalDatabaseBaseLifecycleStateTerminated   ExternalDatabaseBaseLifecycleStateEnum = "TERMINATED"
	ExternalDatabaseBaseLifecycleStateFailed       ExternalDatabaseBaseLifecycleStateEnum = "FAILED"
)

var mappingExternalDatabaseBaseLifecycleState = map[string]ExternalDatabaseBaseLifecycleStateEnum{
	"PROVISIONING":  ExternalDatabaseBaseLifecycleStateProvisioning,
	"NOT_CONNECTED": ExternalDatabaseBaseLifecycleStateNotConnected,
	"AVAILABLE":     ExternalDatabaseBaseLifecycleStateAvailable,
	"UPDATING":      ExternalDatabaseBaseLifecycleStateUpdating,
	"TERMINATING":   ExternalDatabaseBaseLifecycleStateTerminating,
	"TERMINATED":    ExternalDatabaseBaseLifecycleStateTerminated,
	"FAILED":        ExternalDatabaseBaseLifecycleStateFailed,
}

// GetExternalDatabaseBaseLifecycleStateEnumValues Enumerates the set of values for ExternalDatabaseBaseLifecycleStateEnum
func GetExternalDatabaseBaseLifecycleStateEnumValues() []ExternalDatabaseBaseLifecycleStateEnum {
	values := make([]ExternalDatabaseBaseLifecycleStateEnum, 0)
	for _, v := range mappingExternalDatabaseBaseLifecycleState {
		values = append(values, v)
	}
	return values
}

// ExternalDatabaseBaseDatabaseEditionEnum Enum with underlying type: string
type ExternalDatabaseBaseDatabaseEditionEnum string

// Set of constants representing the allowable values for ExternalDatabaseBaseDatabaseEditionEnum
const (
	ExternalDatabaseBaseDatabaseEditionStandardEdition                     ExternalDatabaseBaseDatabaseEditionEnum = "STANDARD_EDITION"
	ExternalDatabaseBaseDatabaseEditionEnterpriseEdition                   ExternalDatabaseBaseDatabaseEditionEnum = "ENTERPRISE_EDITION"
	ExternalDatabaseBaseDatabaseEditionEnterpriseEditionHighPerformance    ExternalDatabaseBaseDatabaseEditionEnum = "ENTERPRISE_EDITION_HIGH_PERFORMANCE"
	ExternalDatabaseBaseDatabaseEditionEnterpriseEditionExtremePerformance ExternalDatabaseBaseDatabaseEditionEnum = "ENTERPRISE_EDITION_EXTREME_PERFORMANCE"
)

var mappingExternalDatabaseBaseDatabaseEdition = map[string]ExternalDatabaseBaseDatabaseEditionEnum{
	"STANDARD_EDITION":                       ExternalDatabaseBaseDatabaseEditionStandardEdition,
	"ENTERPRISE_EDITION":                     ExternalDatabaseBaseDatabaseEditionEnterpriseEdition,
	"ENTERPRISE_EDITION_HIGH_PERFORMANCE":    ExternalDatabaseBaseDatabaseEditionEnterpriseEditionHighPerformance,
	"ENTERPRISE_EDITION_EXTREME_PERFORMANCE": ExternalDatabaseBaseDatabaseEditionEnterpriseEditionExtremePerformance,
}

// GetExternalDatabaseBaseDatabaseEditionEnumValues Enumerates the set of values for ExternalDatabaseBaseDatabaseEditionEnum
func GetExternalDatabaseBaseDatabaseEditionEnumValues() []ExternalDatabaseBaseDatabaseEditionEnum {
	values := make([]ExternalDatabaseBaseDatabaseEditionEnum, 0)
	for _, v := range mappingExternalDatabaseBaseDatabaseEdition {
		values = append(values, v)
	}
	return values
}

// ExternalDatabaseBaseDatabaseConfigurationEnum Enum with underlying type: string
type ExternalDatabaseBaseDatabaseConfigurationEnum string

// Set of constants representing the allowable values for ExternalDatabaseBaseDatabaseConfigurationEnum
const (
	ExternalDatabaseBaseDatabaseConfigurationRac            ExternalDatabaseBaseDatabaseConfigurationEnum = "RAC"
	ExternalDatabaseBaseDatabaseConfigurationSingleInstance ExternalDatabaseBaseDatabaseConfigurationEnum = "SINGLE_INSTANCE"
)

var mappingExternalDatabaseBaseDatabaseConfiguration = map[string]ExternalDatabaseBaseDatabaseConfigurationEnum{
	"RAC":             ExternalDatabaseBaseDatabaseConfigurationRac,
	"SINGLE_INSTANCE": ExternalDatabaseBaseDatabaseConfigurationSingleInstance,
}

// GetExternalDatabaseBaseDatabaseConfigurationEnumValues Enumerates the set of values for ExternalDatabaseBaseDatabaseConfigurationEnum
func GetExternalDatabaseBaseDatabaseConfigurationEnumValues() []ExternalDatabaseBaseDatabaseConfigurationEnum {
	values := make([]ExternalDatabaseBaseDatabaseConfigurationEnum, 0)
	for _, v := range mappingExternalDatabaseBaseDatabaseConfiguration {
		values = append(values, v)
	}
	return values
}
