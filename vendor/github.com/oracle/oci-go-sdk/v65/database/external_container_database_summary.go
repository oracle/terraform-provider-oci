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

// ExternalContainerDatabaseSummary An Oracle Cloud Infrastructure resource that allows you to manage an external Oracle container database.
type ExternalContainerDatabaseSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-friendly name for the external database. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure external database resource.
	Id *string `mandatory:"true" json:"id"`

	// The current state of the Oracle Cloud Infrastructure external database resource.
	LifecycleState ExternalContainerDatabaseSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

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
	DatabaseEdition ExternalContainerDatabaseSummaryDatabaseEditionEnum `mandatory:"false" json:"databaseEdition,omitempty"`

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
	DatabaseConfiguration ExternalContainerDatabaseSummaryDatabaseConfigurationEnum `mandatory:"false" json:"databaseConfiguration,omitempty"`

	DatabaseManagementConfig *DatabaseManagementConfig `mandatory:"false" json:"databaseManagementConfig"`

	StackMonitoringConfig *StackMonitoringConfig `mandatory:"false" json:"stackMonitoringConfig"`
}

func (m ExternalContainerDatabaseSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalContainerDatabaseSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExternalContainerDatabaseSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetExternalContainerDatabaseSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingExternalContainerDatabaseSummaryDatabaseEditionEnum(string(m.DatabaseEdition)); !ok && m.DatabaseEdition != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseEdition: %s. Supported values are: %s.", m.DatabaseEdition, strings.Join(GetExternalContainerDatabaseSummaryDatabaseEditionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExternalContainerDatabaseSummaryDatabaseConfigurationEnum(string(m.DatabaseConfiguration)); !ok && m.DatabaseConfiguration != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseConfiguration: %s. Supported values are: %s.", m.DatabaseConfiguration, strings.Join(GetExternalContainerDatabaseSummaryDatabaseConfigurationEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExternalContainerDatabaseSummaryLifecycleStateEnum Enum with underlying type: string
type ExternalContainerDatabaseSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for ExternalContainerDatabaseSummaryLifecycleStateEnum
const (
	ExternalContainerDatabaseSummaryLifecycleStateProvisioning ExternalContainerDatabaseSummaryLifecycleStateEnum = "PROVISIONING"
	ExternalContainerDatabaseSummaryLifecycleStateNotConnected ExternalContainerDatabaseSummaryLifecycleStateEnum = "NOT_CONNECTED"
	ExternalContainerDatabaseSummaryLifecycleStateAvailable    ExternalContainerDatabaseSummaryLifecycleStateEnum = "AVAILABLE"
	ExternalContainerDatabaseSummaryLifecycleStateUpdating     ExternalContainerDatabaseSummaryLifecycleStateEnum = "UPDATING"
	ExternalContainerDatabaseSummaryLifecycleStateTerminating  ExternalContainerDatabaseSummaryLifecycleStateEnum = "TERMINATING"
	ExternalContainerDatabaseSummaryLifecycleStateTerminated   ExternalContainerDatabaseSummaryLifecycleStateEnum = "TERMINATED"
	ExternalContainerDatabaseSummaryLifecycleStateFailed       ExternalContainerDatabaseSummaryLifecycleStateEnum = "FAILED"
)

var mappingExternalContainerDatabaseSummaryLifecycleStateEnum = map[string]ExternalContainerDatabaseSummaryLifecycleStateEnum{
	"PROVISIONING":  ExternalContainerDatabaseSummaryLifecycleStateProvisioning,
	"NOT_CONNECTED": ExternalContainerDatabaseSummaryLifecycleStateNotConnected,
	"AVAILABLE":     ExternalContainerDatabaseSummaryLifecycleStateAvailable,
	"UPDATING":      ExternalContainerDatabaseSummaryLifecycleStateUpdating,
	"TERMINATING":   ExternalContainerDatabaseSummaryLifecycleStateTerminating,
	"TERMINATED":    ExternalContainerDatabaseSummaryLifecycleStateTerminated,
	"FAILED":        ExternalContainerDatabaseSummaryLifecycleStateFailed,
}

var mappingExternalContainerDatabaseSummaryLifecycleStateEnumLowerCase = map[string]ExternalContainerDatabaseSummaryLifecycleStateEnum{
	"provisioning":  ExternalContainerDatabaseSummaryLifecycleStateProvisioning,
	"not_connected": ExternalContainerDatabaseSummaryLifecycleStateNotConnected,
	"available":     ExternalContainerDatabaseSummaryLifecycleStateAvailable,
	"updating":      ExternalContainerDatabaseSummaryLifecycleStateUpdating,
	"terminating":   ExternalContainerDatabaseSummaryLifecycleStateTerminating,
	"terminated":    ExternalContainerDatabaseSummaryLifecycleStateTerminated,
	"failed":        ExternalContainerDatabaseSummaryLifecycleStateFailed,
}

// GetExternalContainerDatabaseSummaryLifecycleStateEnumValues Enumerates the set of values for ExternalContainerDatabaseSummaryLifecycleStateEnum
func GetExternalContainerDatabaseSummaryLifecycleStateEnumValues() []ExternalContainerDatabaseSummaryLifecycleStateEnum {
	values := make([]ExternalContainerDatabaseSummaryLifecycleStateEnum, 0)
	for _, v := range mappingExternalContainerDatabaseSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalContainerDatabaseSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for ExternalContainerDatabaseSummaryLifecycleStateEnum
func GetExternalContainerDatabaseSummaryLifecycleStateEnumStringValues() []string {
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

// GetMappingExternalContainerDatabaseSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalContainerDatabaseSummaryLifecycleStateEnum(val string) (ExternalContainerDatabaseSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingExternalContainerDatabaseSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExternalContainerDatabaseSummaryDatabaseEditionEnum Enum with underlying type: string
type ExternalContainerDatabaseSummaryDatabaseEditionEnum string

// Set of constants representing the allowable values for ExternalContainerDatabaseSummaryDatabaseEditionEnum
const (
	ExternalContainerDatabaseSummaryDatabaseEditionStandardEdition                     ExternalContainerDatabaseSummaryDatabaseEditionEnum = "STANDARD_EDITION"
	ExternalContainerDatabaseSummaryDatabaseEditionEnterpriseEdition                   ExternalContainerDatabaseSummaryDatabaseEditionEnum = "ENTERPRISE_EDITION"
	ExternalContainerDatabaseSummaryDatabaseEditionEnterpriseEditionHighPerformance    ExternalContainerDatabaseSummaryDatabaseEditionEnum = "ENTERPRISE_EDITION_HIGH_PERFORMANCE"
	ExternalContainerDatabaseSummaryDatabaseEditionEnterpriseEditionExtremePerformance ExternalContainerDatabaseSummaryDatabaseEditionEnum = "ENTERPRISE_EDITION_EXTREME_PERFORMANCE"
)

var mappingExternalContainerDatabaseSummaryDatabaseEditionEnum = map[string]ExternalContainerDatabaseSummaryDatabaseEditionEnum{
	"STANDARD_EDITION":                       ExternalContainerDatabaseSummaryDatabaseEditionStandardEdition,
	"ENTERPRISE_EDITION":                     ExternalContainerDatabaseSummaryDatabaseEditionEnterpriseEdition,
	"ENTERPRISE_EDITION_HIGH_PERFORMANCE":    ExternalContainerDatabaseSummaryDatabaseEditionEnterpriseEditionHighPerformance,
	"ENTERPRISE_EDITION_EXTREME_PERFORMANCE": ExternalContainerDatabaseSummaryDatabaseEditionEnterpriseEditionExtremePerformance,
}

var mappingExternalContainerDatabaseSummaryDatabaseEditionEnumLowerCase = map[string]ExternalContainerDatabaseSummaryDatabaseEditionEnum{
	"standard_edition":                       ExternalContainerDatabaseSummaryDatabaseEditionStandardEdition,
	"enterprise_edition":                     ExternalContainerDatabaseSummaryDatabaseEditionEnterpriseEdition,
	"enterprise_edition_high_performance":    ExternalContainerDatabaseSummaryDatabaseEditionEnterpriseEditionHighPerformance,
	"enterprise_edition_extreme_performance": ExternalContainerDatabaseSummaryDatabaseEditionEnterpriseEditionExtremePerformance,
}

// GetExternalContainerDatabaseSummaryDatabaseEditionEnumValues Enumerates the set of values for ExternalContainerDatabaseSummaryDatabaseEditionEnum
func GetExternalContainerDatabaseSummaryDatabaseEditionEnumValues() []ExternalContainerDatabaseSummaryDatabaseEditionEnum {
	values := make([]ExternalContainerDatabaseSummaryDatabaseEditionEnum, 0)
	for _, v := range mappingExternalContainerDatabaseSummaryDatabaseEditionEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalContainerDatabaseSummaryDatabaseEditionEnumStringValues Enumerates the set of values in String for ExternalContainerDatabaseSummaryDatabaseEditionEnum
func GetExternalContainerDatabaseSummaryDatabaseEditionEnumStringValues() []string {
	return []string{
		"STANDARD_EDITION",
		"ENTERPRISE_EDITION",
		"ENTERPRISE_EDITION_HIGH_PERFORMANCE",
		"ENTERPRISE_EDITION_EXTREME_PERFORMANCE",
	}
}

// GetMappingExternalContainerDatabaseSummaryDatabaseEditionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalContainerDatabaseSummaryDatabaseEditionEnum(val string) (ExternalContainerDatabaseSummaryDatabaseEditionEnum, bool) {
	enum, ok := mappingExternalContainerDatabaseSummaryDatabaseEditionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExternalContainerDatabaseSummaryDatabaseConfigurationEnum Enum with underlying type: string
type ExternalContainerDatabaseSummaryDatabaseConfigurationEnum string

// Set of constants representing the allowable values for ExternalContainerDatabaseSummaryDatabaseConfigurationEnum
const (
	ExternalContainerDatabaseSummaryDatabaseConfigurationRac            ExternalContainerDatabaseSummaryDatabaseConfigurationEnum = "RAC"
	ExternalContainerDatabaseSummaryDatabaseConfigurationSingleInstance ExternalContainerDatabaseSummaryDatabaseConfigurationEnum = "SINGLE_INSTANCE"
)

var mappingExternalContainerDatabaseSummaryDatabaseConfigurationEnum = map[string]ExternalContainerDatabaseSummaryDatabaseConfigurationEnum{
	"RAC":             ExternalContainerDatabaseSummaryDatabaseConfigurationRac,
	"SINGLE_INSTANCE": ExternalContainerDatabaseSummaryDatabaseConfigurationSingleInstance,
}

var mappingExternalContainerDatabaseSummaryDatabaseConfigurationEnumLowerCase = map[string]ExternalContainerDatabaseSummaryDatabaseConfigurationEnum{
	"rac":             ExternalContainerDatabaseSummaryDatabaseConfigurationRac,
	"single_instance": ExternalContainerDatabaseSummaryDatabaseConfigurationSingleInstance,
}

// GetExternalContainerDatabaseSummaryDatabaseConfigurationEnumValues Enumerates the set of values for ExternalContainerDatabaseSummaryDatabaseConfigurationEnum
func GetExternalContainerDatabaseSummaryDatabaseConfigurationEnumValues() []ExternalContainerDatabaseSummaryDatabaseConfigurationEnum {
	values := make([]ExternalContainerDatabaseSummaryDatabaseConfigurationEnum, 0)
	for _, v := range mappingExternalContainerDatabaseSummaryDatabaseConfigurationEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalContainerDatabaseSummaryDatabaseConfigurationEnumStringValues Enumerates the set of values in String for ExternalContainerDatabaseSummaryDatabaseConfigurationEnum
func GetExternalContainerDatabaseSummaryDatabaseConfigurationEnumStringValues() []string {
	return []string{
		"RAC",
		"SINGLE_INSTANCE",
	}
}

// GetMappingExternalContainerDatabaseSummaryDatabaseConfigurationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalContainerDatabaseSummaryDatabaseConfigurationEnum(val string) (ExternalContainerDatabaseSummaryDatabaseConfigurationEnum, bool) {
	enum, ok := mappingExternalContainerDatabaseSummaryDatabaseConfigurationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
