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

// ExternalPluggableDatabaseSummary An Oracle Cloud Infrastructure resource that allows you to manage an external pluggable database.
type ExternalPluggableDatabaseSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-friendly name for the external database. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure external database resource.
	Id *string `mandatory:"true" json:"id"`

	// The current state of the Oracle Cloud Infrastructure external database resource.
	LifecycleState ExternalPluggableDatabaseSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

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
	DatabaseEdition ExternalPluggableDatabaseSummaryDatabaseEditionEnum `mandatory:"false" json:"databaseEdition,omitempty"`

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
	DatabaseConfiguration ExternalPluggableDatabaseSummaryDatabaseConfigurationEnum `mandatory:"false" json:"databaseConfiguration,omitempty"`

	DatabaseManagementConfig *DatabaseManagementConfig `mandatory:"false" json:"databaseManagementConfig"`

	StackMonitoringConfig *StackMonitoringConfig `mandatory:"false" json:"stackMonitoringConfig"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the the non-container database that was converted
	// to a pluggable database to create this resource.
	SourceId *string `mandatory:"false" json:"sourceId"`

	OperationsInsightsConfig *OperationsInsightsConfig `mandatory:"false" json:"operationsInsightsConfig"`
}

func (m ExternalPluggableDatabaseSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalPluggableDatabaseSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExternalPluggableDatabaseSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetExternalPluggableDatabaseSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingExternalPluggableDatabaseSummaryDatabaseEditionEnum(string(m.DatabaseEdition)); !ok && m.DatabaseEdition != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseEdition: %s. Supported values are: %s.", m.DatabaseEdition, strings.Join(GetExternalPluggableDatabaseSummaryDatabaseEditionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExternalPluggableDatabaseSummaryDatabaseConfigurationEnum(string(m.DatabaseConfiguration)); !ok && m.DatabaseConfiguration != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseConfiguration: %s. Supported values are: %s.", m.DatabaseConfiguration, strings.Join(GetExternalPluggableDatabaseSummaryDatabaseConfigurationEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExternalPluggableDatabaseSummaryLifecycleStateEnum Enum with underlying type: string
type ExternalPluggableDatabaseSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for ExternalPluggableDatabaseSummaryLifecycleStateEnum
const (
	ExternalPluggableDatabaseSummaryLifecycleStateProvisioning ExternalPluggableDatabaseSummaryLifecycleStateEnum = "PROVISIONING"
	ExternalPluggableDatabaseSummaryLifecycleStateNotConnected ExternalPluggableDatabaseSummaryLifecycleStateEnum = "NOT_CONNECTED"
	ExternalPluggableDatabaseSummaryLifecycleStateAvailable    ExternalPluggableDatabaseSummaryLifecycleStateEnum = "AVAILABLE"
	ExternalPluggableDatabaseSummaryLifecycleStateUpdating     ExternalPluggableDatabaseSummaryLifecycleStateEnum = "UPDATING"
	ExternalPluggableDatabaseSummaryLifecycleStateTerminating  ExternalPluggableDatabaseSummaryLifecycleStateEnum = "TERMINATING"
	ExternalPluggableDatabaseSummaryLifecycleStateTerminated   ExternalPluggableDatabaseSummaryLifecycleStateEnum = "TERMINATED"
	ExternalPluggableDatabaseSummaryLifecycleStateFailed       ExternalPluggableDatabaseSummaryLifecycleStateEnum = "FAILED"
)

var mappingExternalPluggableDatabaseSummaryLifecycleStateEnum = map[string]ExternalPluggableDatabaseSummaryLifecycleStateEnum{
	"PROVISIONING":  ExternalPluggableDatabaseSummaryLifecycleStateProvisioning,
	"NOT_CONNECTED": ExternalPluggableDatabaseSummaryLifecycleStateNotConnected,
	"AVAILABLE":     ExternalPluggableDatabaseSummaryLifecycleStateAvailable,
	"UPDATING":      ExternalPluggableDatabaseSummaryLifecycleStateUpdating,
	"TERMINATING":   ExternalPluggableDatabaseSummaryLifecycleStateTerminating,
	"TERMINATED":    ExternalPluggableDatabaseSummaryLifecycleStateTerminated,
	"FAILED":        ExternalPluggableDatabaseSummaryLifecycleStateFailed,
}

var mappingExternalPluggableDatabaseSummaryLifecycleStateEnumLowerCase = map[string]ExternalPluggableDatabaseSummaryLifecycleStateEnum{
	"provisioning":  ExternalPluggableDatabaseSummaryLifecycleStateProvisioning,
	"not_connected": ExternalPluggableDatabaseSummaryLifecycleStateNotConnected,
	"available":     ExternalPluggableDatabaseSummaryLifecycleStateAvailable,
	"updating":      ExternalPluggableDatabaseSummaryLifecycleStateUpdating,
	"terminating":   ExternalPluggableDatabaseSummaryLifecycleStateTerminating,
	"terminated":    ExternalPluggableDatabaseSummaryLifecycleStateTerminated,
	"failed":        ExternalPluggableDatabaseSummaryLifecycleStateFailed,
}

// GetExternalPluggableDatabaseSummaryLifecycleStateEnumValues Enumerates the set of values for ExternalPluggableDatabaseSummaryLifecycleStateEnum
func GetExternalPluggableDatabaseSummaryLifecycleStateEnumValues() []ExternalPluggableDatabaseSummaryLifecycleStateEnum {
	values := make([]ExternalPluggableDatabaseSummaryLifecycleStateEnum, 0)
	for _, v := range mappingExternalPluggableDatabaseSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalPluggableDatabaseSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for ExternalPluggableDatabaseSummaryLifecycleStateEnum
func GetExternalPluggableDatabaseSummaryLifecycleStateEnumStringValues() []string {
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

// GetMappingExternalPluggableDatabaseSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalPluggableDatabaseSummaryLifecycleStateEnum(val string) (ExternalPluggableDatabaseSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingExternalPluggableDatabaseSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExternalPluggableDatabaseSummaryDatabaseEditionEnum Enum with underlying type: string
type ExternalPluggableDatabaseSummaryDatabaseEditionEnum string

// Set of constants representing the allowable values for ExternalPluggableDatabaseSummaryDatabaseEditionEnum
const (
	ExternalPluggableDatabaseSummaryDatabaseEditionStandardEdition                     ExternalPluggableDatabaseSummaryDatabaseEditionEnum = "STANDARD_EDITION"
	ExternalPluggableDatabaseSummaryDatabaseEditionEnterpriseEdition                   ExternalPluggableDatabaseSummaryDatabaseEditionEnum = "ENTERPRISE_EDITION"
	ExternalPluggableDatabaseSummaryDatabaseEditionEnterpriseEditionHighPerformance    ExternalPluggableDatabaseSummaryDatabaseEditionEnum = "ENTERPRISE_EDITION_HIGH_PERFORMANCE"
	ExternalPluggableDatabaseSummaryDatabaseEditionEnterpriseEditionExtremePerformance ExternalPluggableDatabaseSummaryDatabaseEditionEnum = "ENTERPRISE_EDITION_EXTREME_PERFORMANCE"
)

var mappingExternalPluggableDatabaseSummaryDatabaseEditionEnum = map[string]ExternalPluggableDatabaseSummaryDatabaseEditionEnum{
	"STANDARD_EDITION":                       ExternalPluggableDatabaseSummaryDatabaseEditionStandardEdition,
	"ENTERPRISE_EDITION":                     ExternalPluggableDatabaseSummaryDatabaseEditionEnterpriseEdition,
	"ENTERPRISE_EDITION_HIGH_PERFORMANCE":    ExternalPluggableDatabaseSummaryDatabaseEditionEnterpriseEditionHighPerformance,
	"ENTERPRISE_EDITION_EXTREME_PERFORMANCE": ExternalPluggableDatabaseSummaryDatabaseEditionEnterpriseEditionExtremePerformance,
}

var mappingExternalPluggableDatabaseSummaryDatabaseEditionEnumLowerCase = map[string]ExternalPluggableDatabaseSummaryDatabaseEditionEnum{
	"standard_edition":                       ExternalPluggableDatabaseSummaryDatabaseEditionStandardEdition,
	"enterprise_edition":                     ExternalPluggableDatabaseSummaryDatabaseEditionEnterpriseEdition,
	"enterprise_edition_high_performance":    ExternalPluggableDatabaseSummaryDatabaseEditionEnterpriseEditionHighPerformance,
	"enterprise_edition_extreme_performance": ExternalPluggableDatabaseSummaryDatabaseEditionEnterpriseEditionExtremePerformance,
}

// GetExternalPluggableDatabaseSummaryDatabaseEditionEnumValues Enumerates the set of values for ExternalPluggableDatabaseSummaryDatabaseEditionEnum
func GetExternalPluggableDatabaseSummaryDatabaseEditionEnumValues() []ExternalPluggableDatabaseSummaryDatabaseEditionEnum {
	values := make([]ExternalPluggableDatabaseSummaryDatabaseEditionEnum, 0)
	for _, v := range mappingExternalPluggableDatabaseSummaryDatabaseEditionEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalPluggableDatabaseSummaryDatabaseEditionEnumStringValues Enumerates the set of values in String for ExternalPluggableDatabaseSummaryDatabaseEditionEnum
func GetExternalPluggableDatabaseSummaryDatabaseEditionEnumStringValues() []string {
	return []string{
		"STANDARD_EDITION",
		"ENTERPRISE_EDITION",
		"ENTERPRISE_EDITION_HIGH_PERFORMANCE",
		"ENTERPRISE_EDITION_EXTREME_PERFORMANCE",
	}
}

// GetMappingExternalPluggableDatabaseSummaryDatabaseEditionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalPluggableDatabaseSummaryDatabaseEditionEnum(val string) (ExternalPluggableDatabaseSummaryDatabaseEditionEnum, bool) {
	enum, ok := mappingExternalPluggableDatabaseSummaryDatabaseEditionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExternalPluggableDatabaseSummaryDatabaseConfigurationEnum Enum with underlying type: string
type ExternalPluggableDatabaseSummaryDatabaseConfigurationEnum string

// Set of constants representing the allowable values for ExternalPluggableDatabaseSummaryDatabaseConfigurationEnum
const (
	ExternalPluggableDatabaseSummaryDatabaseConfigurationRac            ExternalPluggableDatabaseSummaryDatabaseConfigurationEnum = "RAC"
	ExternalPluggableDatabaseSummaryDatabaseConfigurationSingleInstance ExternalPluggableDatabaseSummaryDatabaseConfigurationEnum = "SINGLE_INSTANCE"
)

var mappingExternalPluggableDatabaseSummaryDatabaseConfigurationEnum = map[string]ExternalPluggableDatabaseSummaryDatabaseConfigurationEnum{
	"RAC":             ExternalPluggableDatabaseSummaryDatabaseConfigurationRac,
	"SINGLE_INSTANCE": ExternalPluggableDatabaseSummaryDatabaseConfigurationSingleInstance,
}

var mappingExternalPluggableDatabaseSummaryDatabaseConfigurationEnumLowerCase = map[string]ExternalPluggableDatabaseSummaryDatabaseConfigurationEnum{
	"rac":             ExternalPluggableDatabaseSummaryDatabaseConfigurationRac,
	"single_instance": ExternalPluggableDatabaseSummaryDatabaseConfigurationSingleInstance,
}

// GetExternalPluggableDatabaseSummaryDatabaseConfigurationEnumValues Enumerates the set of values for ExternalPluggableDatabaseSummaryDatabaseConfigurationEnum
func GetExternalPluggableDatabaseSummaryDatabaseConfigurationEnumValues() []ExternalPluggableDatabaseSummaryDatabaseConfigurationEnum {
	values := make([]ExternalPluggableDatabaseSummaryDatabaseConfigurationEnum, 0)
	for _, v := range mappingExternalPluggableDatabaseSummaryDatabaseConfigurationEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalPluggableDatabaseSummaryDatabaseConfigurationEnumStringValues Enumerates the set of values in String for ExternalPluggableDatabaseSummaryDatabaseConfigurationEnum
func GetExternalPluggableDatabaseSummaryDatabaseConfigurationEnumStringValues() []string {
	return []string{
		"RAC",
		"SINGLE_INSTANCE",
	}
}

// GetMappingExternalPluggableDatabaseSummaryDatabaseConfigurationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalPluggableDatabaseSummaryDatabaseConfigurationEnum(val string) (ExternalPluggableDatabaseSummaryDatabaseConfigurationEnum, bool) {
	enum, ok := mappingExternalPluggableDatabaseSummaryDatabaseConfigurationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
