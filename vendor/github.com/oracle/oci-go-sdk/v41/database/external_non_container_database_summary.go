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

// ExternalNonContainerDatabaseSummary An Oracle Cloud Infrastructure external non-container database resource. This resource is used to manage a non-container
// database located outside of Oracle Cloud.
type ExternalNonContainerDatabaseSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-friendly name for the external database. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure external database resource.
	Id *string `mandatory:"true" json:"id"`

	// The current state of the Oracle Cloud Infrastructure external database resource.
	LifecycleState ExternalNonContainerDatabaseSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

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
	DatabaseEdition ExternalNonContainerDatabaseSummaryDatabaseEditionEnum `mandatory:"false" json:"databaseEdition,omitempty"`

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
	DatabaseConfiguration ExternalNonContainerDatabaseSummaryDatabaseConfigurationEnum `mandatory:"false" json:"databaseConfiguration,omitempty"`

	DatabaseManagementConfig *DatabaseManagementConfig `mandatory:"false" json:"databaseManagementConfig"`

	OperationsInsightsConfig *OperationsInsightsConfig `mandatory:"false" json:"operationsInsightsConfig"`
}

func (m ExternalNonContainerDatabaseSummary) String() string {
	return common.PointerString(m)
}

// ExternalNonContainerDatabaseSummaryLifecycleStateEnum Enum with underlying type: string
type ExternalNonContainerDatabaseSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for ExternalNonContainerDatabaseSummaryLifecycleStateEnum
const (
	ExternalNonContainerDatabaseSummaryLifecycleStateProvisioning ExternalNonContainerDatabaseSummaryLifecycleStateEnum = "PROVISIONING"
	ExternalNonContainerDatabaseSummaryLifecycleStateNotConnected ExternalNonContainerDatabaseSummaryLifecycleStateEnum = "NOT_CONNECTED"
	ExternalNonContainerDatabaseSummaryLifecycleStateAvailable    ExternalNonContainerDatabaseSummaryLifecycleStateEnum = "AVAILABLE"
	ExternalNonContainerDatabaseSummaryLifecycleStateUpdating     ExternalNonContainerDatabaseSummaryLifecycleStateEnum = "UPDATING"
	ExternalNonContainerDatabaseSummaryLifecycleStateTerminating  ExternalNonContainerDatabaseSummaryLifecycleStateEnum = "TERMINATING"
	ExternalNonContainerDatabaseSummaryLifecycleStateTerminated   ExternalNonContainerDatabaseSummaryLifecycleStateEnum = "TERMINATED"
	ExternalNonContainerDatabaseSummaryLifecycleStateFailed       ExternalNonContainerDatabaseSummaryLifecycleStateEnum = "FAILED"
)

var mappingExternalNonContainerDatabaseSummaryLifecycleState = map[string]ExternalNonContainerDatabaseSummaryLifecycleStateEnum{
	"PROVISIONING":  ExternalNonContainerDatabaseSummaryLifecycleStateProvisioning,
	"NOT_CONNECTED": ExternalNonContainerDatabaseSummaryLifecycleStateNotConnected,
	"AVAILABLE":     ExternalNonContainerDatabaseSummaryLifecycleStateAvailable,
	"UPDATING":      ExternalNonContainerDatabaseSummaryLifecycleStateUpdating,
	"TERMINATING":   ExternalNonContainerDatabaseSummaryLifecycleStateTerminating,
	"TERMINATED":    ExternalNonContainerDatabaseSummaryLifecycleStateTerminated,
	"FAILED":        ExternalNonContainerDatabaseSummaryLifecycleStateFailed,
}

// GetExternalNonContainerDatabaseSummaryLifecycleStateEnumValues Enumerates the set of values for ExternalNonContainerDatabaseSummaryLifecycleStateEnum
func GetExternalNonContainerDatabaseSummaryLifecycleStateEnumValues() []ExternalNonContainerDatabaseSummaryLifecycleStateEnum {
	values := make([]ExternalNonContainerDatabaseSummaryLifecycleStateEnum, 0)
	for _, v := range mappingExternalNonContainerDatabaseSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}

// ExternalNonContainerDatabaseSummaryDatabaseEditionEnum Enum with underlying type: string
type ExternalNonContainerDatabaseSummaryDatabaseEditionEnum string

// Set of constants representing the allowable values for ExternalNonContainerDatabaseSummaryDatabaseEditionEnum
const (
	ExternalNonContainerDatabaseSummaryDatabaseEditionStandardEdition                     ExternalNonContainerDatabaseSummaryDatabaseEditionEnum = "STANDARD_EDITION"
	ExternalNonContainerDatabaseSummaryDatabaseEditionEnterpriseEdition                   ExternalNonContainerDatabaseSummaryDatabaseEditionEnum = "ENTERPRISE_EDITION"
	ExternalNonContainerDatabaseSummaryDatabaseEditionEnterpriseEditionHighPerformance    ExternalNonContainerDatabaseSummaryDatabaseEditionEnum = "ENTERPRISE_EDITION_HIGH_PERFORMANCE"
	ExternalNonContainerDatabaseSummaryDatabaseEditionEnterpriseEditionExtremePerformance ExternalNonContainerDatabaseSummaryDatabaseEditionEnum = "ENTERPRISE_EDITION_EXTREME_PERFORMANCE"
)

var mappingExternalNonContainerDatabaseSummaryDatabaseEdition = map[string]ExternalNonContainerDatabaseSummaryDatabaseEditionEnum{
	"STANDARD_EDITION":                       ExternalNonContainerDatabaseSummaryDatabaseEditionStandardEdition,
	"ENTERPRISE_EDITION":                     ExternalNonContainerDatabaseSummaryDatabaseEditionEnterpriseEdition,
	"ENTERPRISE_EDITION_HIGH_PERFORMANCE":    ExternalNonContainerDatabaseSummaryDatabaseEditionEnterpriseEditionHighPerformance,
	"ENTERPRISE_EDITION_EXTREME_PERFORMANCE": ExternalNonContainerDatabaseSummaryDatabaseEditionEnterpriseEditionExtremePerformance,
}

// GetExternalNonContainerDatabaseSummaryDatabaseEditionEnumValues Enumerates the set of values for ExternalNonContainerDatabaseSummaryDatabaseEditionEnum
func GetExternalNonContainerDatabaseSummaryDatabaseEditionEnumValues() []ExternalNonContainerDatabaseSummaryDatabaseEditionEnum {
	values := make([]ExternalNonContainerDatabaseSummaryDatabaseEditionEnum, 0)
	for _, v := range mappingExternalNonContainerDatabaseSummaryDatabaseEdition {
		values = append(values, v)
	}
	return values
}

// ExternalNonContainerDatabaseSummaryDatabaseConfigurationEnum Enum with underlying type: string
type ExternalNonContainerDatabaseSummaryDatabaseConfigurationEnum string

// Set of constants representing the allowable values for ExternalNonContainerDatabaseSummaryDatabaseConfigurationEnum
const (
	ExternalNonContainerDatabaseSummaryDatabaseConfigurationRac            ExternalNonContainerDatabaseSummaryDatabaseConfigurationEnum = "RAC"
	ExternalNonContainerDatabaseSummaryDatabaseConfigurationSingleInstance ExternalNonContainerDatabaseSummaryDatabaseConfigurationEnum = "SINGLE_INSTANCE"
)

var mappingExternalNonContainerDatabaseSummaryDatabaseConfiguration = map[string]ExternalNonContainerDatabaseSummaryDatabaseConfigurationEnum{
	"RAC":             ExternalNonContainerDatabaseSummaryDatabaseConfigurationRac,
	"SINGLE_INSTANCE": ExternalNonContainerDatabaseSummaryDatabaseConfigurationSingleInstance,
}

// GetExternalNonContainerDatabaseSummaryDatabaseConfigurationEnumValues Enumerates the set of values for ExternalNonContainerDatabaseSummaryDatabaseConfigurationEnum
func GetExternalNonContainerDatabaseSummaryDatabaseConfigurationEnumValues() []ExternalNonContainerDatabaseSummaryDatabaseConfigurationEnum {
	values := make([]ExternalNonContainerDatabaseSummaryDatabaseConfigurationEnum, 0)
	for _, v := range mappingExternalNonContainerDatabaseSummaryDatabaseConfiguration {
		values = append(values, v)
	}
	return values
}
