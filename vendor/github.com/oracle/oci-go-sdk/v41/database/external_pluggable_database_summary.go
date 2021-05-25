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

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the the non-container database that was converted
	// to a pluggable database to create this resource.
	SourceId *string `mandatory:"false" json:"sourceId"`

	OperationsInsightsConfig *OperationsInsightsConfig `mandatory:"false" json:"operationsInsightsConfig"`
}

func (m ExternalPluggableDatabaseSummary) String() string {
	return common.PointerString(m)
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

var mappingExternalPluggableDatabaseSummaryLifecycleState = map[string]ExternalPluggableDatabaseSummaryLifecycleStateEnum{
	"PROVISIONING":  ExternalPluggableDatabaseSummaryLifecycleStateProvisioning,
	"NOT_CONNECTED": ExternalPluggableDatabaseSummaryLifecycleStateNotConnected,
	"AVAILABLE":     ExternalPluggableDatabaseSummaryLifecycleStateAvailable,
	"UPDATING":      ExternalPluggableDatabaseSummaryLifecycleStateUpdating,
	"TERMINATING":   ExternalPluggableDatabaseSummaryLifecycleStateTerminating,
	"TERMINATED":    ExternalPluggableDatabaseSummaryLifecycleStateTerminated,
	"FAILED":        ExternalPluggableDatabaseSummaryLifecycleStateFailed,
}

// GetExternalPluggableDatabaseSummaryLifecycleStateEnumValues Enumerates the set of values for ExternalPluggableDatabaseSummaryLifecycleStateEnum
func GetExternalPluggableDatabaseSummaryLifecycleStateEnumValues() []ExternalPluggableDatabaseSummaryLifecycleStateEnum {
	values := make([]ExternalPluggableDatabaseSummaryLifecycleStateEnum, 0)
	for _, v := range mappingExternalPluggableDatabaseSummaryLifecycleState {
		values = append(values, v)
	}
	return values
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

var mappingExternalPluggableDatabaseSummaryDatabaseEdition = map[string]ExternalPluggableDatabaseSummaryDatabaseEditionEnum{
	"STANDARD_EDITION":                       ExternalPluggableDatabaseSummaryDatabaseEditionStandardEdition,
	"ENTERPRISE_EDITION":                     ExternalPluggableDatabaseSummaryDatabaseEditionEnterpriseEdition,
	"ENTERPRISE_EDITION_HIGH_PERFORMANCE":    ExternalPluggableDatabaseSummaryDatabaseEditionEnterpriseEditionHighPerformance,
	"ENTERPRISE_EDITION_EXTREME_PERFORMANCE": ExternalPluggableDatabaseSummaryDatabaseEditionEnterpriseEditionExtremePerformance,
}

// GetExternalPluggableDatabaseSummaryDatabaseEditionEnumValues Enumerates the set of values for ExternalPluggableDatabaseSummaryDatabaseEditionEnum
func GetExternalPluggableDatabaseSummaryDatabaseEditionEnumValues() []ExternalPluggableDatabaseSummaryDatabaseEditionEnum {
	values := make([]ExternalPluggableDatabaseSummaryDatabaseEditionEnum, 0)
	for _, v := range mappingExternalPluggableDatabaseSummaryDatabaseEdition {
		values = append(values, v)
	}
	return values
}

// ExternalPluggableDatabaseSummaryDatabaseConfigurationEnum Enum with underlying type: string
type ExternalPluggableDatabaseSummaryDatabaseConfigurationEnum string

// Set of constants representing the allowable values for ExternalPluggableDatabaseSummaryDatabaseConfigurationEnum
const (
	ExternalPluggableDatabaseSummaryDatabaseConfigurationRac            ExternalPluggableDatabaseSummaryDatabaseConfigurationEnum = "RAC"
	ExternalPluggableDatabaseSummaryDatabaseConfigurationSingleInstance ExternalPluggableDatabaseSummaryDatabaseConfigurationEnum = "SINGLE_INSTANCE"
)

var mappingExternalPluggableDatabaseSummaryDatabaseConfiguration = map[string]ExternalPluggableDatabaseSummaryDatabaseConfigurationEnum{
	"RAC":             ExternalPluggableDatabaseSummaryDatabaseConfigurationRac,
	"SINGLE_INSTANCE": ExternalPluggableDatabaseSummaryDatabaseConfigurationSingleInstance,
}

// GetExternalPluggableDatabaseSummaryDatabaseConfigurationEnumValues Enumerates the set of values for ExternalPluggableDatabaseSummaryDatabaseConfigurationEnum
func GetExternalPluggableDatabaseSummaryDatabaseConfigurationEnumValues() []ExternalPluggableDatabaseSummaryDatabaseConfigurationEnum {
	values := make([]ExternalPluggableDatabaseSummaryDatabaseConfigurationEnum, 0)
	for _, v := range mappingExternalPluggableDatabaseSummaryDatabaseConfiguration {
		values = append(values, v)
	}
	return values
}
