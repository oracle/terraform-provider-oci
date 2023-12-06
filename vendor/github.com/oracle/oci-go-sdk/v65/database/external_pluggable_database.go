// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// ExternalPluggableDatabase an external Oracle pluggable database.
type ExternalPluggableDatabase struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-friendly name for the external database. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure external database resource.
	Id *string `mandatory:"true" json:"id"`

	// The current state of the Oracle Cloud Infrastructure external database resource.
	LifecycleState ExternalPluggableDatabaseLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the database was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the
	// CreateExternalContainerDatabaseDetails that contains
	// the specified CreateExternalPluggableDatabaseDetails resource.
	ExternalContainerDatabaseId *string `mandatory:"true" json:"externalContainerDatabaseId"`

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
	DatabaseEdition ExternalPluggableDatabaseDatabaseEditionEnum `mandatory:"false" json:"databaseEdition,omitempty"`

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
	DatabaseConfiguration ExternalPluggableDatabaseDatabaseConfigurationEnum `mandatory:"false" json:"databaseConfiguration,omitempty"`

	DatabaseManagementConfig *DatabaseManagementConfig `mandatory:"false" json:"databaseManagementConfig"`

	StackMonitoringConfig *StackMonitoringConfig `mandatory:"false" json:"stackMonitoringConfig"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the the non-container database that was converted
	// to a pluggable database to create this resource.
	SourceId *string `mandatory:"false" json:"sourceId"`

	OperationsInsightsConfig *OperationsInsightsConfig `mandatory:"false" json:"operationsInsightsConfig"`
}

func (m ExternalPluggableDatabase) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalPluggableDatabase) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExternalPluggableDatabaseLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetExternalPluggableDatabaseLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingExternalPluggableDatabaseDatabaseEditionEnum(string(m.DatabaseEdition)); !ok && m.DatabaseEdition != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseEdition: %s. Supported values are: %s.", m.DatabaseEdition, strings.Join(GetExternalPluggableDatabaseDatabaseEditionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExternalPluggableDatabaseDatabaseConfigurationEnum(string(m.DatabaseConfiguration)); !ok && m.DatabaseConfiguration != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseConfiguration: %s. Supported values are: %s.", m.DatabaseConfiguration, strings.Join(GetExternalPluggableDatabaseDatabaseConfigurationEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExternalPluggableDatabaseLifecycleStateEnum Enum with underlying type: string
type ExternalPluggableDatabaseLifecycleStateEnum string

// Set of constants representing the allowable values for ExternalPluggableDatabaseLifecycleStateEnum
const (
	ExternalPluggableDatabaseLifecycleStateProvisioning ExternalPluggableDatabaseLifecycleStateEnum = "PROVISIONING"
	ExternalPluggableDatabaseLifecycleStateNotConnected ExternalPluggableDatabaseLifecycleStateEnum = "NOT_CONNECTED"
	ExternalPluggableDatabaseLifecycleStateAvailable    ExternalPluggableDatabaseLifecycleStateEnum = "AVAILABLE"
	ExternalPluggableDatabaseLifecycleStateUpdating     ExternalPluggableDatabaseLifecycleStateEnum = "UPDATING"
	ExternalPluggableDatabaseLifecycleStateTerminating  ExternalPluggableDatabaseLifecycleStateEnum = "TERMINATING"
	ExternalPluggableDatabaseLifecycleStateTerminated   ExternalPluggableDatabaseLifecycleStateEnum = "TERMINATED"
	ExternalPluggableDatabaseLifecycleStateFailed       ExternalPluggableDatabaseLifecycleStateEnum = "FAILED"
)

var mappingExternalPluggableDatabaseLifecycleStateEnum = map[string]ExternalPluggableDatabaseLifecycleStateEnum{
	"PROVISIONING":  ExternalPluggableDatabaseLifecycleStateProvisioning,
	"NOT_CONNECTED": ExternalPluggableDatabaseLifecycleStateNotConnected,
	"AVAILABLE":     ExternalPluggableDatabaseLifecycleStateAvailable,
	"UPDATING":      ExternalPluggableDatabaseLifecycleStateUpdating,
	"TERMINATING":   ExternalPluggableDatabaseLifecycleStateTerminating,
	"TERMINATED":    ExternalPluggableDatabaseLifecycleStateTerminated,
	"FAILED":        ExternalPluggableDatabaseLifecycleStateFailed,
}

var mappingExternalPluggableDatabaseLifecycleStateEnumLowerCase = map[string]ExternalPluggableDatabaseLifecycleStateEnum{
	"provisioning":  ExternalPluggableDatabaseLifecycleStateProvisioning,
	"not_connected": ExternalPluggableDatabaseLifecycleStateNotConnected,
	"available":     ExternalPluggableDatabaseLifecycleStateAvailable,
	"updating":      ExternalPluggableDatabaseLifecycleStateUpdating,
	"terminating":   ExternalPluggableDatabaseLifecycleStateTerminating,
	"terminated":    ExternalPluggableDatabaseLifecycleStateTerminated,
	"failed":        ExternalPluggableDatabaseLifecycleStateFailed,
}

// GetExternalPluggableDatabaseLifecycleStateEnumValues Enumerates the set of values for ExternalPluggableDatabaseLifecycleStateEnum
func GetExternalPluggableDatabaseLifecycleStateEnumValues() []ExternalPluggableDatabaseLifecycleStateEnum {
	values := make([]ExternalPluggableDatabaseLifecycleStateEnum, 0)
	for _, v := range mappingExternalPluggableDatabaseLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalPluggableDatabaseLifecycleStateEnumStringValues Enumerates the set of values in String for ExternalPluggableDatabaseLifecycleStateEnum
func GetExternalPluggableDatabaseLifecycleStateEnumStringValues() []string {
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

// GetMappingExternalPluggableDatabaseLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalPluggableDatabaseLifecycleStateEnum(val string) (ExternalPluggableDatabaseLifecycleStateEnum, bool) {
	enum, ok := mappingExternalPluggableDatabaseLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExternalPluggableDatabaseDatabaseEditionEnum Enum with underlying type: string
type ExternalPluggableDatabaseDatabaseEditionEnum string

// Set of constants representing the allowable values for ExternalPluggableDatabaseDatabaseEditionEnum
const (
	ExternalPluggableDatabaseDatabaseEditionStandardEdition                     ExternalPluggableDatabaseDatabaseEditionEnum = "STANDARD_EDITION"
	ExternalPluggableDatabaseDatabaseEditionEnterpriseEdition                   ExternalPluggableDatabaseDatabaseEditionEnum = "ENTERPRISE_EDITION"
	ExternalPluggableDatabaseDatabaseEditionEnterpriseEditionHighPerformance    ExternalPluggableDatabaseDatabaseEditionEnum = "ENTERPRISE_EDITION_HIGH_PERFORMANCE"
	ExternalPluggableDatabaseDatabaseEditionEnterpriseEditionExtremePerformance ExternalPluggableDatabaseDatabaseEditionEnum = "ENTERPRISE_EDITION_EXTREME_PERFORMANCE"
)

var mappingExternalPluggableDatabaseDatabaseEditionEnum = map[string]ExternalPluggableDatabaseDatabaseEditionEnum{
	"STANDARD_EDITION":                       ExternalPluggableDatabaseDatabaseEditionStandardEdition,
	"ENTERPRISE_EDITION":                     ExternalPluggableDatabaseDatabaseEditionEnterpriseEdition,
	"ENTERPRISE_EDITION_HIGH_PERFORMANCE":    ExternalPluggableDatabaseDatabaseEditionEnterpriseEditionHighPerformance,
	"ENTERPRISE_EDITION_EXTREME_PERFORMANCE": ExternalPluggableDatabaseDatabaseEditionEnterpriseEditionExtremePerformance,
}

var mappingExternalPluggableDatabaseDatabaseEditionEnumLowerCase = map[string]ExternalPluggableDatabaseDatabaseEditionEnum{
	"standard_edition":                       ExternalPluggableDatabaseDatabaseEditionStandardEdition,
	"enterprise_edition":                     ExternalPluggableDatabaseDatabaseEditionEnterpriseEdition,
	"enterprise_edition_high_performance":    ExternalPluggableDatabaseDatabaseEditionEnterpriseEditionHighPerformance,
	"enterprise_edition_extreme_performance": ExternalPluggableDatabaseDatabaseEditionEnterpriseEditionExtremePerformance,
}

// GetExternalPluggableDatabaseDatabaseEditionEnumValues Enumerates the set of values for ExternalPluggableDatabaseDatabaseEditionEnum
func GetExternalPluggableDatabaseDatabaseEditionEnumValues() []ExternalPluggableDatabaseDatabaseEditionEnum {
	values := make([]ExternalPluggableDatabaseDatabaseEditionEnum, 0)
	for _, v := range mappingExternalPluggableDatabaseDatabaseEditionEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalPluggableDatabaseDatabaseEditionEnumStringValues Enumerates the set of values in String for ExternalPluggableDatabaseDatabaseEditionEnum
func GetExternalPluggableDatabaseDatabaseEditionEnumStringValues() []string {
	return []string{
		"STANDARD_EDITION",
		"ENTERPRISE_EDITION",
		"ENTERPRISE_EDITION_HIGH_PERFORMANCE",
		"ENTERPRISE_EDITION_EXTREME_PERFORMANCE",
	}
}

// GetMappingExternalPluggableDatabaseDatabaseEditionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalPluggableDatabaseDatabaseEditionEnum(val string) (ExternalPluggableDatabaseDatabaseEditionEnum, bool) {
	enum, ok := mappingExternalPluggableDatabaseDatabaseEditionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExternalPluggableDatabaseDatabaseConfigurationEnum Enum with underlying type: string
type ExternalPluggableDatabaseDatabaseConfigurationEnum string

// Set of constants representing the allowable values for ExternalPluggableDatabaseDatabaseConfigurationEnum
const (
	ExternalPluggableDatabaseDatabaseConfigurationRac            ExternalPluggableDatabaseDatabaseConfigurationEnum = "RAC"
	ExternalPluggableDatabaseDatabaseConfigurationSingleInstance ExternalPluggableDatabaseDatabaseConfigurationEnum = "SINGLE_INSTANCE"
)

var mappingExternalPluggableDatabaseDatabaseConfigurationEnum = map[string]ExternalPluggableDatabaseDatabaseConfigurationEnum{
	"RAC":             ExternalPluggableDatabaseDatabaseConfigurationRac,
	"SINGLE_INSTANCE": ExternalPluggableDatabaseDatabaseConfigurationSingleInstance,
}

var mappingExternalPluggableDatabaseDatabaseConfigurationEnumLowerCase = map[string]ExternalPluggableDatabaseDatabaseConfigurationEnum{
	"rac":             ExternalPluggableDatabaseDatabaseConfigurationRac,
	"single_instance": ExternalPluggableDatabaseDatabaseConfigurationSingleInstance,
}

// GetExternalPluggableDatabaseDatabaseConfigurationEnumValues Enumerates the set of values for ExternalPluggableDatabaseDatabaseConfigurationEnum
func GetExternalPluggableDatabaseDatabaseConfigurationEnumValues() []ExternalPluggableDatabaseDatabaseConfigurationEnum {
	values := make([]ExternalPluggableDatabaseDatabaseConfigurationEnum, 0)
	for _, v := range mappingExternalPluggableDatabaseDatabaseConfigurationEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalPluggableDatabaseDatabaseConfigurationEnumStringValues Enumerates the set of values in String for ExternalPluggableDatabaseDatabaseConfigurationEnum
func GetExternalPluggableDatabaseDatabaseConfigurationEnumStringValues() []string {
	return []string{
		"RAC",
		"SINGLE_INSTANCE",
	}
}

// GetMappingExternalPluggableDatabaseDatabaseConfigurationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalPluggableDatabaseDatabaseConfigurationEnum(val string) (ExternalPluggableDatabaseDatabaseConfigurationEnum, bool) {
	enum, ok := mappingExternalPluggableDatabaseDatabaseConfigurationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
