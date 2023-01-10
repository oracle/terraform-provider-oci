// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// ExternalContainerDatabase An Oracle Cloud Infrastructure resource that allows you to manage an external container database.
type ExternalContainerDatabase struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-friendly name for the external database. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure external database resource.
	Id *string `mandatory:"true" json:"id"`

	// The current state of the Oracle Cloud Infrastructure external database resource.
	LifecycleState ExternalContainerDatabaseLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

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
	DatabaseEdition ExternalContainerDatabaseDatabaseEditionEnum `mandatory:"false" json:"databaseEdition,omitempty"`

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
	DatabaseConfiguration ExternalContainerDatabaseDatabaseConfigurationEnum `mandatory:"false" json:"databaseConfiguration,omitempty"`

	DatabaseManagementConfig *DatabaseManagementConfig `mandatory:"false" json:"databaseManagementConfig"`

	StackMonitoringConfig *StackMonitoringConfig `mandatory:"false" json:"stackMonitoringConfig"`
}

func (m ExternalContainerDatabase) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalContainerDatabase) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExternalContainerDatabaseLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetExternalContainerDatabaseLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingExternalContainerDatabaseDatabaseEditionEnum(string(m.DatabaseEdition)); !ok && m.DatabaseEdition != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseEdition: %s. Supported values are: %s.", m.DatabaseEdition, strings.Join(GetExternalContainerDatabaseDatabaseEditionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExternalContainerDatabaseDatabaseConfigurationEnum(string(m.DatabaseConfiguration)); !ok && m.DatabaseConfiguration != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseConfiguration: %s. Supported values are: %s.", m.DatabaseConfiguration, strings.Join(GetExternalContainerDatabaseDatabaseConfigurationEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExternalContainerDatabaseLifecycleStateEnum Enum with underlying type: string
type ExternalContainerDatabaseLifecycleStateEnum string

// Set of constants representing the allowable values for ExternalContainerDatabaseLifecycleStateEnum
const (
	ExternalContainerDatabaseLifecycleStateProvisioning ExternalContainerDatabaseLifecycleStateEnum = "PROVISIONING"
	ExternalContainerDatabaseLifecycleStateNotConnected ExternalContainerDatabaseLifecycleStateEnum = "NOT_CONNECTED"
	ExternalContainerDatabaseLifecycleStateAvailable    ExternalContainerDatabaseLifecycleStateEnum = "AVAILABLE"
	ExternalContainerDatabaseLifecycleStateUpdating     ExternalContainerDatabaseLifecycleStateEnum = "UPDATING"
	ExternalContainerDatabaseLifecycleStateTerminating  ExternalContainerDatabaseLifecycleStateEnum = "TERMINATING"
	ExternalContainerDatabaseLifecycleStateTerminated   ExternalContainerDatabaseLifecycleStateEnum = "TERMINATED"
	ExternalContainerDatabaseLifecycleStateFailed       ExternalContainerDatabaseLifecycleStateEnum = "FAILED"
)

var mappingExternalContainerDatabaseLifecycleStateEnum = map[string]ExternalContainerDatabaseLifecycleStateEnum{
	"PROVISIONING":  ExternalContainerDatabaseLifecycleStateProvisioning,
	"NOT_CONNECTED": ExternalContainerDatabaseLifecycleStateNotConnected,
	"AVAILABLE":     ExternalContainerDatabaseLifecycleStateAvailable,
	"UPDATING":      ExternalContainerDatabaseLifecycleStateUpdating,
	"TERMINATING":   ExternalContainerDatabaseLifecycleStateTerminating,
	"TERMINATED":    ExternalContainerDatabaseLifecycleStateTerminated,
	"FAILED":        ExternalContainerDatabaseLifecycleStateFailed,
}

var mappingExternalContainerDatabaseLifecycleStateEnumLowerCase = map[string]ExternalContainerDatabaseLifecycleStateEnum{
	"provisioning":  ExternalContainerDatabaseLifecycleStateProvisioning,
	"not_connected": ExternalContainerDatabaseLifecycleStateNotConnected,
	"available":     ExternalContainerDatabaseLifecycleStateAvailable,
	"updating":      ExternalContainerDatabaseLifecycleStateUpdating,
	"terminating":   ExternalContainerDatabaseLifecycleStateTerminating,
	"terminated":    ExternalContainerDatabaseLifecycleStateTerminated,
	"failed":        ExternalContainerDatabaseLifecycleStateFailed,
}

// GetExternalContainerDatabaseLifecycleStateEnumValues Enumerates the set of values for ExternalContainerDatabaseLifecycleStateEnum
func GetExternalContainerDatabaseLifecycleStateEnumValues() []ExternalContainerDatabaseLifecycleStateEnum {
	values := make([]ExternalContainerDatabaseLifecycleStateEnum, 0)
	for _, v := range mappingExternalContainerDatabaseLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalContainerDatabaseLifecycleStateEnumStringValues Enumerates the set of values in String for ExternalContainerDatabaseLifecycleStateEnum
func GetExternalContainerDatabaseLifecycleStateEnumStringValues() []string {
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

// GetMappingExternalContainerDatabaseLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalContainerDatabaseLifecycleStateEnum(val string) (ExternalContainerDatabaseLifecycleStateEnum, bool) {
	enum, ok := mappingExternalContainerDatabaseLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExternalContainerDatabaseDatabaseEditionEnum Enum with underlying type: string
type ExternalContainerDatabaseDatabaseEditionEnum string

// Set of constants representing the allowable values for ExternalContainerDatabaseDatabaseEditionEnum
const (
	ExternalContainerDatabaseDatabaseEditionStandardEdition                     ExternalContainerDatabaseDatabaseEditionEnum = "STANDARD_EDITION"
	ExternalContainerDatabaseDatabaseEditionEnterpriseEdition                   ExternalContainerDatabaseDatabaseEditionEnum = "ENTERPRISE_EDITION"
	ExternalContainerDatabaseDatabaseEditionEnterpriseEditionHighPerformance    ExternalContainerDatabaseDatabaseEditionEnum = "ENTERPRISE_EDITION_HIGH_PERFORMANCE"
	ExternalContainerDatabaseDatabaseEditionEnterpriseEditionExtremePerformance ExternalContainerDatabaseDatabaseEditionEnum = "ENTERPRISE_EDITION_EXTREME_PERFORMANCE"
)

var mappingExternalContainerDatabaseDatabaseEditionEnum = map[string]ExternalContainerDatabaseDatabaseEditionEnum{
	"STANDARD_EDITION":                       ExternalContainerDatabaseDatabaseEditionStandardEdition,
	"ENTERPRISE_EDITION":                     ExternalContainerDatabaseDatabaseEditionEnterpriseEdition,
	"ENTERPRISE_EDITION_HIGH_PERFORMANCE":    ExternalContainerDatabaseDatabaseEditionEnterpriseEditionHighPerformance,
	"ENTERPRISE_EDITION_EXTREME_PERFORMANCE": ExternalContainerDatabaseDatabaseEditionEnterpriseEditionExtremePerformance,
}

var mappingExternalContainerDatabaseDatabaseEditionEnumLowerCase = map[string]ExternalContainerDatabaseDatabaseEditionEnum{
	"standard_edition":                       ExternalContainerDatabaseDatabaseEditionStandardEdition,
	"enterprise_edition":                     ExternalContainerDatabaseDatabaseEditionEnterpriseEdition,
	"enterprise_edition_high_performance":    ExternalContainerDatabaseDatabaseEditionEnterpriseEditionHighPerformance,
	"enterprise_edition_extreme_performance": ExternalContainerDatabaseDatabaseEditionEnterpriseEditionExtremePerformance,
}

// GetExternalContainerDatabaseDatabaseEditionEnumValues Enumerates the set of values for ExternalContainerDatabaseDatabaseEditionEnum
func GetExternalContainerDatabaseDatabaseEditionEnumValues() []ExternalContainerDatabaseDatabaseEditionEnum {
	values := make([]ExternalContainerDatabaseDatabaseEditionEnum, 0)
	for _, v := range mappingExternalContainerDatabaseDatabaseEditionEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalContainerDatabaseDatabaseEditionEnumStringValues Enumerates the set of values in String for ExternalContainerDatabaseDatabaseEditionEnum
func GetExternalContainerDatabaseDatabaseEditionEnumStringValues() []string {
	return []string{
		"STANDARD_EDITION",
		"ENTERPRISE_EDITION",
		"ENTERPRISE_EDITION_HIGH_PERFORMANCE",
		"ENTERPRISE_EDITION_EXTREME_PERFORMANCE",
	}
}

// GetMappingExternalContainerDatabaseDatabaseEditionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalContainerDatabaseDatabaseEditionEnum(val string) (ExternalContainerDatabaseDatabaseEditionEnum, bool) {
	enum, ok := mappingExternalContainerDatabaseDatabaseEditionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExternalContainerDatabaseDatabaseConfigurationEnum Enum with underlying type: string
type ExternalContainerDatabaseDatabaseConfigurationEnum string

// Set of constants representing the allowable values for ExternalContainerDatabaseDatabaseConfigurationEnum
const (
	ExternalContainerDatabaseDatabaseConfigurationRac            ExternalContainerDatabaseDatabaseConfigurationEnum = "RAC"
	ExternalContainerDatabaseDatabaseConfigurationSingleInstance ExternalContainerDatabaseDatabaseConfigurationEnum = "SINGLE_INSTANCE"
)

var mappingExternalContainerDatabaseDatabaseConfigurationEnum = map[string]ExternalContainerDatabaseDatabaseConfigurationEnum{
	"RAC":             ExternalContainerDatabaseDatabaseConfigurationRac,
	"SINGLE_INSTANCE": ExternalContainerDatabaseDatabaseConfigurationSingleInstance,
}

var mappingExternalContainerDatabaseDatabaseConfigurationEnumLowerCase = map[string]ExternalContainerDatabaseDatabaseConfigurationEnum{
	"rac":             ExternalContainerDatabaseDatabaseConfigurationRac,
	"single_instance": ExternalContainerDatabaseDatabaseConfigurationSingleInstance,
}

// GetExternalContainerDatabaseDatabaseConfigurationEnumValues Enumerates the set of values for ExternalContainerDatabaseDatabaseConfigurationEnum
func GetExternalContainerDatabaseDatabaseConfigurationEnumValues() []ExternalContainerDatabaseDatabaseConfigurationEnum {
	values := make([]ExternalContainerDatabaseDatabaseConfigurationEnum, 0)
	for _, v := range mappingExternalContainerDatabaseDatabaseConfigurationEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalContainerDatabaseDatabaseConfigurationEnumStringValues Enumerates the set of values in String for ExternalContainerDatabaseDatabaseConfigurationEnum
func GetExternalContainerDatabaseDatabaseConfigurationEnumStringValues() []string {
	return []string{
		"RAC",
		"SINGLE_INSTANCE",
	}
}

// GetMappingExternalContainerDatabaseDatabaseConfigurationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalContainerDatabaseDatabaseConfigurationEnum(val string) (ExternalContainerDatabaseDatabaseConfigurationEnum, bool) {
	enum, ok := mappingExternalContainerDatabaseDatabaseConfigurationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
