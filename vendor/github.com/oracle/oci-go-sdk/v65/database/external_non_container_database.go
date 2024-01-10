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

// ExternalNonContainerDatabase an external Oracle non-container database.
type ExternalNonContainerDatabase struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-friendly name for the external database. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure external database resource.
	Id *string `mandatory:"true" json:"id"`

	// The current state of the Oracle Cloud Infrastructure external database resource.
	LifecycleState ExternalNonContainerDatabaseLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

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
	DatabaseEdition ExternalNonContainerDatabaseDatabaseEditionEnum `mandatory:"false" json:"databaseEdition,omitempty"`

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
	DatabaseConfiguration ExternalNonContainerDatabaseDatabaseConfigurationEnum `mandatory:"false" json:"databaseConfiguration,omitempty"`

	DatabaseManagementConfig *DatabaseManagementConfig `mandatory:"false" json:"databaseManagementConfig"`

	StackMonitoringConfig *StackMonitoringConfig `mandatory:"false" json:"stackMonitoringConfig"`

	OperationsInsightsConfig *OperationsInsightsConfig `mandatory:"false" json:"operationsInsightsConfig"`
}

func (m ExternalNonContainerDatabase) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalNonContainerDatabase) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExternalNonContainerDatabaseLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetExternalNonContainerDatabaseLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingExternalNonContainerDatabaseDatabaseEditionEnum(string(m.DatabaseEdition)); !ok && m.DatabaseEdition != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseEdition: %s. Supported values are: %s.", m.DatabaseEdition, strings.Join(GetExternalNonContainerDatabaseDatabaseEditionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExternalNonContainerDatabaseDatabaseConfigurationEnum(string(m.DatabaseConfiguration)); !ok && m.DatabaseConfiguration != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseConfiguration: %s. Supported values are: %s.", m.DatabaseConfiguration, strings.Join(GetExternalNonContainerDatabaseDatabaseConfigurationEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExternalNonContainerDatabaseLifecycleStateEnum Enum with underlying type: string
type ExternalNonContainerDatabaseLifecycleStateEnum string

// Set of constants representing the allowable values for ExternalNonContainerDatabaseLifecycleStateEnum
const (
	ExternalNonContainerDatabaseLifecycleStateProvisioning ExternalNonContainerDatabaseLifecycleStateEnum = "PROVISIONING"
	ExternalNonContainerDatabaseLifecycleStateNotConnected ExternalNonContainerDatabaseLifecycleStateEnum = "NOT_CONNECTED"
	ExternalNonContainerDatabaseLifecycleStateAvailable    ExternalNonContainerDatabaseLifecycleStateEnum = "AVAILABLE"
	ExternalNonContainerDatabaseLifecycleStateUpdating     ExternalNonContainerDatabaseLifecycleStateEnum = "UPDATING"
	ExternalNonContainerDatabaseLifecycleStateTerminating  ExternalNonContainerDatabaseLifecycleStateEnum = "TERMINATING"
	ExternalNonContainerDatabaseLifecycleStateTerminated   ExternalNonContainerDatabaseLifecycleStateEnum = "TERMINATED"
	ExternalNonContainerDatabaseLifecycleStateFailed       ExternalNonContainerDatabaseLifecycleStateEnum = "FAILED"
)

var mappingExternalNonContainerDatabaseLifecycleStateEnum = map[string]ExternalNonContainerDatabaseLifecycleStateEnum{
	"PROVISIONING":  ExternalNonContainerDatabaseLifecycleStateProvisioning,
	"NOT_CONNECTED": ExternalNonContainerDatabaseLifecycleStateNotConnected,
	"AVAILABLE":     ExternalNonContainerDatabaseLifecycleStateAvailable,
	"UPDATING":      ExternalNonContainerDatabaseLifecycleStateUpdating,
	"TERMINATING":   ExternalNonContainerDatabaseLifecycleStateTerminating,
	"TERMINATED":    ExternalNonContainerDatabaseLifecycleStateTerminated,
	"FAILED":        ExternalNonContainerDatabaseLifecycleStateFailed,
}

var mappingExternalNonContainerDatabaseLifecycleStateEnumLowerCase = map[string]ExternalNonContainerDatabaseLifecycleStateEnum{
	"provisioning":  ExternalNonContainerDatabaseLifecycleStateProvisioning,
	"not_connected": ExternalNonContainerDatabaseLifecycleStateNotConnected,
	"available":     ExternalNonContainerDatabaseLifecycleStateAvailable,
	"updating":      ExternalNonContainerDatabaseLifecycleStateUpdating,
	"terminating":   ExternalNonContainerDatabaseLifecycleStateTerminating,
	"terminated":    ExternalNonContainerDatabaseLifecycleStateTerminated,
	"failed":        ExternalNonContainerDatabaseLifecycleStateFailed,
}

// GetExternalNonContainerDatabaseLifecycleStateEnumValues Enumerates the set of values for ExternalNonContainerDatabaseLifecycleStateEnum
func GetExternalNonContainerDatabaseLifecycleStateEnumValues() []ExternalNonContainerDatabaseLifecycleStateEnum {
	values := make([]ExternalNonContainerDatabaseLifecycleStateEnum, 0)
	for _, v := range mappingExternalNonContainerDatabaseLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalNonContainerDatabaseLifecycleStateEnumStringValues Enumerates the set of values in String for ExternalNonContainerDatabaseLifecycleStateEnum
func GetExternalNonContainerDatabaseLifecycleStateEnumStringValues() []string {
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

// GetMappingExternalNonContainerDatabaseLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalNonContainerDatabaseLifecycleStateEnum(val string) (ExternalNonContainerDatabaseLifecycleStateEnum, bool) {
	enum, ok := mappingExternalNonContainerDatabaseLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExternalNonContainerDatabaseDatabaseEditionEnum Enum with underlying type: string
type ExternalNonContainerDatabaseDatabaseEditionEnum string

// Set of constants representing the allowable values for ExternalNonContainerDatabaseDatabaseEditionEnum
const (
	ExternalNonContainerDatabaseDatabaseEditionStandardEdition                     ExternalNonContainerDatabaseDatabaseEditionEnum = "STANDARD_EDITION"
	ExternalNonContainerDatabaseDatabaseEditionEnterpriseEdition                   ExternalNonContainerDatabaseDatabaseEditionEnum = "ENTERPRISE_EDITION"
	ExternalNonContainerDatabaseDatabaseEditionEnterpriseEditionHighPerformance    ExternalNonContainerDatabaseDatabaseEditionEnum = "ENTERPRISE_EDITION_HIGH_PERFORMANCE"
	ExternalNonContainerDatabaseDatabaseEditionEnterpriseEditionExtremePerformance ExternalNonContainerDatabaseDatabaseEditionEnum = "ENTERPRISE_EDITION_EXTREME_PERFORMANCE"
)

var mappingExternalNonContainerDatabaseDatabaseEditionEnum = map[string]ExternalNonContainerDatabaseDatabaseEditionEnum{
	"STANDARD_EDITION":                       ExternalNonContainerDatabaseDatabaseEditionStandardEdition,
	"ENTERPRISE_EDITION":                     ExternalNonContainerDatabaseDatabaseEditionEnterpriseEdition,
	"ENTERPRISE_EDITION_HIGH_PERFORMANCE":    ExternalNonContainerDatabaseDatabaseEditionEnterpriseEditionHighPerformance,
	"ENTERPRISE_EDITION_EXTREME_PERFORMANCE": ExternalNonContainerDatabaseDatabaseEditionEnterpriseEditionExtremePerformance,
}

var mappingExternalNonContainerDatabaseDatabaseEditionEnumLowerCase = map[string]ExternalNonContainerDatabaseDatabaseEditionEnum{
	"standard_edition":                       ExternalNonContainerDatabaseDatabaseEditionStandardEdition,
	"enterprise_edition":                     ExternalNonContainerDatabaseDatabaseEditionEnterpriseEdition,
	"enterprise_edition_high_performance":    ExternalNonContainerDatabaseDatabaseEditionEnterpriseEditionHighPerformance,
	"enterprise_edition_extreme_performance": ExternalNonContainerDatabaseDatabaseEditionEnterpriseEditionExtremePerformance,
}

// GetExternalNonContainerDatabaseDatabaseEditionEnumValues Enumerates the set of values for ExternalNonContainerDatabaseDatabaseEditionEnum
func GetExternalNonContainerDatabaseDatabaseEditionEnumValues() []ExternalNonContainerDatabaseDatabaseEditionEnum {
	values := make([]ExternalNonContainerDatabaseDatabaseEditionEnum, 0)
	for _, v := range mappingExternalNonContainerDatabaseDatabaseEditionEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalNonContainerDatabaseDatabaseEditionEnumStringValues Enumerates the set of values in String for ExternalNonContainerDatabaseDatabaseEditionEnum
func GetExternalNonContainerDatabaseDatabaseEditionEnumStringValues() []string {
	return []string{
		"STANDARD_EDITION",
		"ENTERPRISE_EDITION",
		"ENTERPRISE_EDITION_HIGH_PERFORMANCE",
		"ENTERPRISE_EDITION_EXTREME_PERFORMANCE",
	}
}

// GetMappingExternalNonContainerDatabaseDatabaseEditionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalNonContainerDatabaseDatabaseEditionEnum(val string) (ExternalNonContainerDatabaseDatabaseEditionEnum, bool) {
	enum, ok := mappingExternalNonContainerDatabaseDatabaseEditionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExternalNonContainerDatabaseDatabaseConfigurationEnum Enum with underlying type: string
type ExternalNonContainerDatabaseDatabaseConfigurationEnum string

// Set of constants representing the allowable values for ExternalNonContainerDatabaseDatabaseConfigurationEnum
const (
	ExternalNonContainerDatabaseDatabaseConfigurationRac            ExternalNonContainerDatabaseDatabaseConfigurationEnum = "RAC"
	ExternalNonContainerDatabaseDatabaseConfigurationSingleInstance ExternalNonContainerDatabaseDatabaseConfigurationEnum = "SINGLE_INSTANCE"
)

var mappingExternalNonContainerDatabaseDatabaseConfigurationEnum = map[string]ExternalNonContainerDatabaseDatabaseConfigurationEnum{
	"RAC":             ExternalNonContainerDatabaseDatabaseConfigurationRac,
	"SINGLE_INSTANCE": ExternalNonContainerDatabaseDatabaseConfigurationSingleInstance,
}

var mappingExternalNonContainerDatabaseDatabaseConfigurationEnumLowerCase = map[string]ExternalNonContainerDatabaseDatabaseConfigurationEnum{
	"rac":             ExternalNonContainerDatabaseDatabaseConfigurationRac,
	"single_instance": ExternalNonContainerDatabaseDatabaseConfigurationSingleInstance,
}

// GetExternalNonContainerDatabaseDatabaseConfigurationEnumValues Enumerates the set of values for ExternalNonContainerDatabaseDatabaseConfigurationEnum
func GetExternalNonContainerDatabaseDatabaseConfigurationEnumValues() []ExternalNonContainerDatabaseDatabaseConfigurationEnum {
	values := make([]ExternalNonContainerDatabaseDatabaseConfigurationEnum, 0)
	for _, v := range mappingExternalNonContainerDatabaseDatabaseConfigurationEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalNonContainerDatabaseDatabaseConfigurationEnumStringValues Enumerates the set of values in String for ExternalNonContainerDatabaseDatabaseConfigurationEnum
func GetExternalNonContainerDatabaseDatabaseConfigurationEnumStringValues() []string {
	return []string{
		"RAC",
		"SINGLE_INSTANCE",
	}
}

// GetMappingExternalNonContainerDatabaseDatabaseConfigurationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalNonContainerDatabaseDatabaseConfigurationEnum(val string) (ExternalNonContainerDatabaseDatabaseConfigurationEnum, bool) {
	enum, ok := mappingExternalNonContainerDatabaseDatabaseConfigurationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
