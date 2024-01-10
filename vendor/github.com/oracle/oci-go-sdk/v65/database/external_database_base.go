// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

	StackMonitoringConfig *StackMonitoringConfig `mandatory:"false" json:"stackMonitoringConfig"`
}

func (m ExternalDatabaseBase) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalDatabaseBase) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExternalDatabaseBaseLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetExternalDatabaseBaseLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingExternalDatabaseBaseDatabaseEditionEnum(string(m.DatabaseEdition)); !ok && m.DatabaseEdition != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseEdition: %s. Supported values are: %s.", m.DatabaseEdition, strings.Join(GetExternalDatabaseBaseDatabaseEditionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExternalDatabaseBaseDatabaseConfigurationEnum(string(m.DatabaseConfiguration)); !ok && m.DatabaseConfiguration != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseConfiguration: %s. Supported values are: %s.", m.DatabaseConfiguration, strings.Join(GetExternalDatabaseBaseDatabaseConfigurationEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingExternalDatabaseBaseLifecycleStateEnum = map[string]ExternalDatabaseBaseLifecycleStateEnum{
	"PROVISIONING":  ExternalDatabaseBaseLifecycleStateProvisioning,
	"NOT_CONNECTED": ExternalDatabaseBaseLifecycleStateNotConnected,
	"AVAILABLE":     ExternalDatabaseBaseLifecycleStateAvailable,
	"UPDATING":      ExternalDatabaseBaseLifecycleStateUpdating,
	"TERMINATING":   ExternalDatabaseBaseLifecycleStateTerminating,
	"TERMINATED":    ExternalDatabaseBaseLifecycleStateTerminated,
	"FAILED":        ExternalDatabaseBaseLifecycleStateFailed,
}

var mappingExternalDatabaseBaseLifecycleStateEnumLowerCase = map[string]ExternalDatabaseBaseLifecycleStateEnum{
	"provisioning":  ExternalDatabaseBaseLifecycleStateProvisioning,
	"not_connected": ExternalDatabaseBaseLifecycleStateNotConnected,
	"available":     ExternalDatabaseBaseLifecycleStateAvailable,
	"updating":      ExternalDatabaseBaseLifecycleStateUpdating,
	"terminating":   ExternalDatabaseBaseLifecycleStateTerminating,
	"terminated":    ExternalDatabaseBaseLifecycleStateTerminated,
	"failed":        ExternalDatabaseBaseLifecycleStateFailed,
}

// GetExternalDatabaseBaseLifecycleStateEnumValues Enumerates the set of values for ExternalDatabaseBaseLifecycleStateEnum
func GetExternalDatabaseBaseLifecycleStateEnumValues() []ExternalDatabaseBaseLifecycleStateEnum {
	values := make([]ExternalDatabaseBaseLifecycleStateEnum, 0)
	for _, v := range mappingExternalDatabaseBaseLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalDatabaseBaseLifecycleStateEnumStringValues Enumerates the set of values in String for ExternalDatabaseBaseLifecycleStateEnum
func GetExternalDatabaseBaseLifecycleStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"NOT_CONNECTED",
		"AVAILABLE",
		"UPDATING",
		"TERMINATING",
		"TERMINATED",
		"FAILED",
	}
}

// GetMappingExternalDatabaseBaseLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalDatabaseBaseLifecycleStateEnum(val string) (ExternalDatabaseBaseLifecycleStateEnum, bool) {
	enum, ok := mappingExternalDatabaseBaseLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
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

var mappingExternalDatabaseBaseDatabaseEditionEnum = map[string]ExternalDatabaseBaseDatabaseEditionEnum{
	"STANDARD_EDITION":                       ExternalDatabaseBaseDatabaseEditionStandardEdition,
	"ENTERPRISE_EDITION":                     ExternalDatabaseBaseDatabaseEditionEnterpriseEdition,
	"ENTERPRISE_EDITION_HIGH_PERFORMANCE":    ExternalDatabaseBaseDatabaseEditionEnterpriseEditionHighPerformance,
	"ENTERPRISE_EDITION_EXTREME_PERFORMANCE": ExternalDatabaseBaseDatabaseEditionEnterpriseEditionExtremePerformance,
}

var mappingExternalDatabaseBaseDatabaseEditionEnumLowerCase = map[string]ExternalDatabaseBaseDatabaseEditionEnum{
	"standard_edition":                       ExternalDatabaseBaseDatabaseEditionStandardEdition,
	"enterprise_edition":                     ExternalDatabaseBaseDatabaseEditionEnterpriseEdition,
	"enterprise_edition_high_performance":    ExternalDatabaseBaseDatabaseEditionEnterpriseEditionHighPerformance,
	"enterprise_edition_extreme_performance": ExternalDatabaseBaseDatabaseEditionEnterpriseEditionExtremePerformance,
}

// GetExternalDatabaseBaseDatabaseEditionEnumValues Enumerates the set of values for ExternalDatabaseBaseDatabaseEditionEnum
func GetExternalDatabaseBaseDatabaseEditionEnumValues() []ExternalDatabaseBaseDatabaseEditionEnum {
	values := make([]ExternalDatabaseBaseDatabaseEditionEnum, 0)
	for _, v := range mappingExternalDatabaseBaseDatabaseEditionEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalDatabaseBaseDatabaseEditionEnumStringValues Enumerates the set of values in String for ExternalDatabaseBaseDatabaseEditionEnum
func GetExternalDatabaseBaseDatabaseEditionEnumStringValues() []string {
	return []string{
		"STANDARD_EDITION",
		"ENTERPRISE_EDITION",
		"ENTERPRISE_EDITION_HIGH_PERFORMANCE",
		"ENTERPRISE_EDITION_EXTREME_PERFORMANCE",
	}
}

// GetMappingExternalDatabaseBaseDatabaseEditionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalDatabaseBaseDatabaseEditionEnum(val string) (ExternalDatabaseBaseDatabaseEditionEnum, bool) {
	enum, ok := mappingExternalDatabaseBaseDatabaseEditionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExternalDatabaseBaseDatabaseConfigurationEnum Enum with underlying type: string
type ExternalDatabaseBaseDatabaseConfigurationEnum string

// Set of constants representing the allowable values for ExternalDatabaseBaseDatabaseConfigurationEnum
const (
	ExternalDatabaseBaseDatabaseConfigurationRac            ExternalDatabaseBaseDatabaseConfigurationEnum = "RAC"
	ExternalDatabaseBaseDatabaseConfigurationSingleInstance ExternalDatabaseBaseDatabaseConfigurationEnum = "SINGLE_INSTANCE"
)

var mappingExternalDatabaseBaseDatabaseConfigurationEnum = map[string]ExternalDatabaseBaseDatabaseConfigurationEnum{
	"RAC":             ExternalDatabaseBaseDatabaseConfigurationRac,
	"SINGLE_INSTANCE": ExternalDatabaseBaseDatabaseConfigurationSingleInstance,
}

var mappingExternalDatabaseBaseDatabaseConfigurationEnumLowerCase = map[string]ExternalDatabaseBaseDatabaseConfigurationEnum{
	"rac":             ExternalDatabaseBaseDatabaseConfigurationRac,
	"single_instance": ExternalDatabaseBaseDatabaseConfigurationSingleInstance,
}

// GetExternalDatabaseBaseDatabaseConfigurationEnumValues Enumerates the set of values for ExternalDatabaseBaseDatabaseConfigurationEnum
func GetExternalDatabaseBaseDatabaseConfigurationEnumValues() []ExternalDatabaseBaseDatabaseConfigurationEnum {
	values := make([]ExternalDatabaseBaseDatabaseConfigurationEnum, 0)
	for _, v := range mappingExternalDatabaseBaseDatabaseConfigurationEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalDatabaseBaseDatabaseConfigurationEnumStringValues Enumerates the set of values in String for ExternalDatabaseBaseDatabaseConfigurationEnum
func GetExternalDatabaseBaseDatabaseConfigurationEnumStringValues() []string {
	return []string{
		"RAC",
		"SINGLE_INSTANCE",
	}
}

// GetMappingExternalDatabaseBaseDatabaseConfigurationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalDatabaseBaseDatabaseConfigurationEnum(val string) (ExternalDatabaseBaseDatabaseConfigurationEnum, bool) {
	enum, ok := mappingExternalDatabaseBaseDatabaseConfigurationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
